package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chinahtl/tencent-doc-sdk/constant"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

// ExportDocument 导出文档
/*
  指引地址：https://docs.qq.com/open/document/app/openapi/v2/file/export/async_export.html
*/
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

// GetExportProgress 查询文档导出进度
/*
  指引地址：https://docs.qq.com/open/document/app/openapi/v2/file/export/export_progress.html
*/
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
