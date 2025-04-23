package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chinahtl/tencent-doc-sdk/constant"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

// ListDocuments 获取文档列表
/*
  指引地址：https://docs.qq.com/open/document/app/openapi/v2/file/filter/filter.html
*/
func (c *Client) ListDocuments(ctx context.Context, params *model.ListParams) (*model.ListDocumentsResponse, error) {
	if c.token == nil || c.token.AccessToken == "" {
		return nil, fmt.Errorf("access token is required")
	}

	// 设置默认值
	if params.ListType == "" {
		params.ListType = constant.ListTypeFolder
	}
	if params.SortType == "" {
		params.SortType = constant.SortTypeBrowse
	}
	if params.Limit <= 0 || params.Limit > 20 {
		params.Limit = 20 // 强制限制最大20条
	}
	if params.FolderID == "" {
		params.FolderID = "/" // 默认根目录
	}

	// 构建请求URL
	endpoint := fmt.Sprintf("%s/drive/v2/filter", constant.APIEndpoint)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse endpoint failed: %w", err)
	}

	// 添加查询参数
	q := u.Query()
	q.Add("listType", params.ListType)
	q.Add("sortType", params.SortType)
	q.Add("asc", fmt.Sprintf("%d", params.Asc))
	q.Add("folderID", params.FolderID)
	q.Add("start", fmt.Sprintf("%d", params.Start))
	q.Add("limit", fmt.Sprintf("%d", params.Limit))
	q.Add("isOwner", fmt.Sprintf("%d", params.IsOwner))

	if params.FileType != "" {
		q.Add("fileType", params.FileType)
	}

	u.RawQuery = q.Encode()

	// 发送请求
	var result model.ListDocumentsResponse
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
		return nil, fmt.Errorf("list documents failed: %w", err)
	}

	if result.Ret != 0 {
		return nil, fmt.Errorf("api error: %s (ret=%d)", result.Msg, result.Ret)
	}

	return &result, nil
}

// SearchDocuments 搜索文档
/*
  指引地址：https://docs.qq.com/open/document/app/openapi/v2/file/search/search.html
*/
func (c *Client) SearchDocuments(ctx context.Context, params *model.SearchParams) (*model.SearchDocumentsResponse, error) {
	if c.token == nil || c.token.AccessToken == "" {
		return nil, fmt.Errorf("access token is required")
	}

	// 构建请求URL
	endpoint := fmt.Sprintf("%s/drive/v2/search", constant.APIEndpoint)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse endpoint failed: %w", err)
	}

	// 添加查询参数
	q := u.Query()
	q.Add("searchType", params.SearchType)
	q.Add("searchKey", params.SearchKey)
	q.Add("folderID", params.FolderID)
	q.Add("byOwnership", fmt.Sprintf("%d", params.ByOwnership))
	q.Add("offset", fmt.Sprintf("%d", params.Offset))
	q.Add("size", fmt.Sprintf("%d", params.Size))
	if params.FileTypes != "" {
		q.Add("fileTypes", params.FileTypes)
	}
	if params.SortType != "" {
		q.Add("sortType", params.SortType)
	}
	if params.Asc != 0 {
		q.Add("asc", fmt.Sprintf("%d", params.Asc))
	}

	u.RawQuery = q.Encode()

	// 发送请求
	var result model.SearchDocumentsResponse
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
		return nil, fmt.Errorf("search documents failed: %w", err)
	}

	if result.Ret != 0 {
		return nil, fmt.Errorf("api error: %s (ret=%d)", result.Msg, result.Ret)
	}

	return &result, nil
}

// GetFileMetadata 获取文件元数据
/*
  指引地址：https://docs.qq.com/open/document/app/openapi/v2/file/files/metadata.html
*/
func (c *Client) GetFileMetadata(ctx context.Context, fileID string) (*model.FileMetadataResponse, error) {
	if c.token == nil || c.token.AccessToken == "" {
		return nil, fmt.Errorf("access token is required")
	}

	// 构建请求URL
	endpoint := fmt.Sprintf("%s/drive/v2/files/%s/metadata", constant.APIEndpoint, fileID)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse endpoint failed: %w", err)
	}

	// 发送请求
	var result model.FileMetadataResponse
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
		return nil, fmt.Errorf("get file metadata failed: %w", err)
	}

	if result.Ret != 0 {
		return nil, fmt.Errorf("api error: %s (ret=%d)", result.Msg, result.Ret)
	}

	return &result, nil
}
