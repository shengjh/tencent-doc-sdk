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
	docClient := createClient()
	authURL := getAuthURL(docClient)
	fmt.Printf("Auth URL: %s\n", authURL)

	code := "4QF_AON2MMIXN72NC5D_JQ"
	tokenResp := exchangeToken(docClient, code)
	fmt.Printf("Access Token: %s\n", tokenResp.AccessToken)

	setToken(docClient, tokenResp)

	listRootDocuments(docClient)
	listFolderDocuments(docClient, "XCoWJIgNtEZj")
	searchDocuments(docClient, "测试")

	docID := "300000000$XZaWRcPddZFp"
	exportDocument(docClient, docID)
}

// 创建并返回一个客户端实例
func createClient() *client.Client {
	existingToken := &model.Token{
		AccessToken:  "",
		RefreshToken: "",
		ExpiresIn:    108000,
		UserID:       "",
	}

	return client.NewClient(
		config.WithClientID(""),
		config.WithClientSecret(""),
		config.WithRedirectURI(""),
		config.WithTimeout(time.Second*30),
		config.WithRandomState(util.GenerateRandomString(20)),
		config.WithInitialToken(existingToken),
	)
}

// 获取授权URL
func getAuthURL(docClient *client.Client) string {
	return docClient.GetAuthURL()
}

// 用授权码换取Token
func exchangeToken(docClient *client.Client, code string) *model.TokenResponse {
	tokenResp, err := docClient.ExchangeToken(context.Background(), code)
	if err != nil {
		log.Fatalf("Exchange token failed: %v", err)
	}
	return tokenResp
}

// 设置Token
func setToken(docClient *client.Client, tokenResp *model.TokenResponse) {
	docClient.WithToken(&model.Token{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		ExpiresIn:    tokenResp.ExpiresIn,
		UserID:       tokenResp.UserID,
	})
}

// 列出根目录下的文档
func listRootDocuments(docClient *client.Client) {
	params := &model.ListParams{}
	resp, err := docClient.ListDocuments(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

// 列出指定文件夹下的文档
func listFolderDocuments(docClient *client.Client, folderID string) {
	params := &model.ListParams{FolderID: folderID}
	resp, err := docClient.ListDocuments(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

// 搜索文档
func searchDocuments(docClient *client.Client, searchKey string) {
	params := &model.SearchParams{
		SearchType: "title",
		SearchKey:  searchKey,
		Offset:     0,
		Size:       10,
	}
	resp, err := docClient.SearchDocuments(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

// 导出文档
func exportDocument(docClient *client.Client, docID string) {
	req := &model.ExportRequest{}
	resp, err := docClient.ExportDocument(context.Background(), docID, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	progressResp := pollExportProgress(docClient, docID, resp.Data.OperationID)
	if progressResp.Data.Progress == 100 {
		downloadExportedFile(progressResp.Data.URL)
	} else {
		fmt.Println("导出超时，请稍后重试")
	}
}

// 轮询导出进度
func pollExportProgress(docClient *client.Client, docID, operationID string) *model.ExportProgressResponse {
	var progressResp *model.ExportProgressResponse
	for i := 0; i < 100; i++ {
		time.Sleep(2 * time.Second)

		var err error
		progressResp, err = docClient.GetExportProgress(context.Background(), docID, operationID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("当前进度: %d%%\n", progressResp.Data.Progress)
		if progressResp.Data.Progress == 100 {
			break
		}
	}
	return progressResp
}

// 下载导出的文件
func downloadExportedFile(url string) {
	fmt.Println("导出完成，下载URL:", url)
	cos, err := util.DownloadFromCOS(url, "./tmp/cos")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("下载完成，文件路径:", cos)
}
