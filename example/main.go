package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chinahtl/tencent-doc-sdk/client"
	"github.com/chinahtl/tencent-doc-sdk/config"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

func main() {
	existingToken := &model.Token{
		AccessToken:  "",
		RefreshToken: "",
		ExpiresIn:    108000,
		UserID:       "",
	}

	// 创建客户端
	docClient := client.NewClient(
		config.WithClientID(""),
		config.WithClientSecret(""),
		config.WithRedirectURI(""),
		config.WithTimeout(time.Second*30),
		config.WithRandomState(util.GenerateRandomString(20)),
		config.WithInitialToken(existingToken),
	)

	// 获取授权URL
	authURL := docClient.GetAuthURL()
	fmt.Printf("Auth URL: %s\n", authURL)

	// 假设用户授权后返回了code
	code := "4QF_AON2MMIXN72NC5D_JQ"

	// 用code换取token
	tokenResp, err := docClient.ExchangeToken(context.Background(), code)
	if err != nil {
		log.Fatalf("Exchange token failed: %v", err)
	}

	fmt.Printf("Access Token: %s\n", tokenResp.AccessToken)

	// 设置token
	docClient.WithToken(&model.Token{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		ExpiresIn:    tokenResp.ExpiresIn,
		UserID:       tokenResp.UserID,
	})

	// 构建根目录搜索
	params := &model.ListParams{}

	resp, err := docClient.ListDocuments(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	//---------------------------------

	// 构建指定目录搜索
	params1 := &model.ListParams{FolderID: "XCoWJIgNtEZj"}

	resp2, err := docClient.ListDocuments(context.Background(), params1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp2)

	//---------------------------------

	// 构建搜索参数
	params2 := &model.SearchParams{
		SearchType: "title", // 按标题搜索
		SearchKey:  "测试",    // 搜索关键词
		Offset:     0,       // 从第一条开始
		Size:       10,      // 每页10条
	}

	// 执行搜索
	resp3, err := docClient.SearchDocuments(context.Background(), params2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp3)

	//---------------------------------

	// 文档ID和导出参数
	docID := "300000000$XZaWRcPddZFp"
	req := &model.ExportRequest{
		//ExportType: constant.ExportTypePDF, // 导出为PDF
	}

	// 执行导出
	resp4, err := docClient.ExportDocument(context.Background(), docID, req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp4)

	//---------------------------------

	// 轮询导出进度
	var progressResp *model.ExportProgressResponse
	for i := 0; i < 100; i++ {
		time.Sleep(2 * time.Second) // 每2秒查询一次

		progressResp, err = docClient.GetExportProgress(
			context.Background(),
			docID,
			resp4.Data.OperationID,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("当前进度: %d%%\n", progressResp.Data.Progress)
		if progressResp.Data.Progress == 100 {
			break
		}
	}

	// 下载导出文件
	if progressResp.Data.Progress == 100 {
		fmt.Println("导出完成，下载URL:", progressResp.Data.URL)

		cos, err := util.DownloadFromCOS(progressResp.Data.URL, "./tmp/cos")
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("下载完成，文件路径:", cos)

	} else {
		fmt.Println("导出超时，请稍后重试")
	}

}
