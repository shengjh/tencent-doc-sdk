package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chinahtl/tencent-doc-sdk/constant"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

// ExportDocument 异步导出腾讯文档到指定格式。
//
// ctx 用于控制请求的上下文，可用于超时控制和取消。
//
// docID 是要导出的文档ID，不能为空。
//
// req 包含导出相关的配置参数：
//   - ExportType: 导出文件类型，支持的格式取决于文档类型
//
// 返回导出任务的响应信息，包括：
//   - OperationID: 导出任务ID，用于后续查询导出进度
//   - Status: 任务状态
//
// 可能返回的错误：
//   - access token未设置
//   - 文档ID为空
//   - API调用失败
//   - 服务端返回错误
//
// API参考：https://docs.qq.com/open/document/app/openapi/v2/file/export/async_export.html
func (c *Client) ExportDocument(ctx context.Context, docID string, req *model.ExportRequest) (*model.ExportResponse, error) {
	if c.token == nil || c.token.AccessToken == "" {
		return nil, fmt.Errorf("access token is required")
	}
	if docID == "" {
		return nil, fmt.Errorf("document ID cannot be empty")
	}

	// 构建请求URL
	endpoint := fmt.Sprintf("%s/drive/v2/files/%s/async-export", constant.APIEndpoint, docID)

	// 准备表单数据
	form := url.Values{}
	if req.ExportType != "" {
		form.Add("exportType", req.ExportType)
	}

	// 发送请求
	var result model.ExportResponse
	err := util.PostFormWithHeaders(
		ctx,
		c.httpClient,
		endpoint,
		form,
		map[string]string{
			"Access-Token": c.token.AccessToken,
			"Client-Id":    c.config.ClientID,
			"Open-Id":      c.token.UserID,
			"Content-Type": "application/x-www-form-urlencoded",
		},
		&result,
	)
	if err != nil {
		return nil, fmt.Errorf("export document failed: %w", err)
	}

	if result.Ret != 0 {
		return nil, fmt.Errorf("api error: %s (ret=%d)", result.Msg, result.Ret)
	}

	return &result, nil
}

// GetExportProgress 查询腾讯文档的导出进度。
//
// ctx 用于控制请求的上下文，可用于超时控制和取消。
//
// docID 是正在导出的文档ID，不能为空。
//
// operationID 是导出任务的操作ID，来自ExportDocument的返回结果，不能为空。
//
// 返回导出进度信息，包括：
//   - Status: 导出状态
//   - URL: 导出成功时的文件下载地址
//   - Progress: 当前导出进度（百分比）
//   - Message: 状态描述或错误信息
//
// 可能返回的错误：
//   - access token未设置
//   - 文档ID为空
//   - 操作ID为空
//   - API调用失败
//   - 服务端返回错误
//
// API参考：https://docs.qq.com/open/document/app/openapi/v2/file/export/export_progress.html
func (c *Client) GetExportProgress(
	ctx context.Context,
	docID string,
	operationID string,
) (*model.ExportProgressResponse, error) {
	if c.token == nil || c.token.AccessToken == "" {
		return nil, fmt.Errorf("access token is required")
	}
	if docID == "" {
		return nil, fmt.Errorf("document ID cannot be empty")
	}
	if operationID == "" {
		return nil, fmt.Errorf("operation ID cannot be empty")
	}

	// 构建请求URL
	endpoint := fmt.Sprintf("%s/drive/v2/files/%s/export-progress", constant.APIEndpoint, docID)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse endpoint failed: %w", err)
	}

	// 添加查询参数
	q := u.Query()
	q.Add("operationID", operationID)
	u.RawQuery = q.Encode()

	// 发送请求
	var result model.ExportProgressResponse
	err = util.GetWithCustomHeaders(
		ctx,
		c.httpClient,
		u.String(),
		map[string]string{
			"Access-Token": c.token.AccessToken,
			"Client-Id":    c.config.ClientID,
			"Open-Id":      c.token.UserID,
		},
		&result,
	)
	if err != nil {
		return nil, fmt.Errorf("get export progress failed: %w", err)
	}

	if result.Ret != 0 {
		return nil, fmt.Errorf("api error: %s (ret=%d)", result.Msg, result.Ret)
	}

	return &result, nil
}
