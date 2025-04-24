package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chinahtl/tencent-doc-sdk/constant"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

// ListDocuments 获取腾讯文档列表。
//
// params 包含以下字段：
//   - ListType: 列表类型，默认为folder（文件夹）
//   - SortType: 排序类型，默认为browse（浏览）
//   - Asc: 是否升序排序
//   - FolderID: 文件夹ID，默认为根目录"/"
//   - Start: 起始位置
//   - Limit: 每页数量，默认20，最大20
//   - IsOwner: 是否仅显示自己创建的文档
//   - FileType: 文件类型过滤（可选）
//
// 返回文档列表响应，包含文档信息列表及相关元数据。
// 如果发生错误，可能的错误类型包括：
//   - access token未设置
//   - API调用失败
//   - 服务端返回错误
//
// API参考：https://docs.qq.com/open/document/app/openapi/v2/file/filter/filter.html
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

// SearchDocuments 在腾讯文档中搜索文档。
//
// params 包含以下字段：
//   - SearchType: 搜索类型
//   - SearchKey: 搜索关键词
//   - FolderID: 文件夹ID
//   - ByOwnership: 是否按所有权搜索
//   - Offset: 分页起始位置
//   - Size: 每页数量
//   - FileTypes: 文件类型过滤（可选）
//   - SortType: 排序类型（可选）
//   - Asc: 是否升序排序（可选）
//
// 返回搜索结果响应，包含匹配的文档列表及相关元数据。
// 如果发生错误，可能的错误类型包括：
//   - access token未设置
//   - API调用失败
//   - 服务端返回错误
//
// API参考：https://docs.qq.com/open/document/app/openapi/v2/file/search/search.html
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

// GetFileMetadata 获取腾讯文档文件的元数据信息。
//
// fileID 是要获取元数据的文件ID，必填参数。
//
// 返回文件的元数据信息，包括：
//   - 文件基本信息（ID、名称、类型等）
//   - 创建和修改时间
//   - 所有者信息
//   - 文件大小
//   - 其他元数据属性
//
// 如果发生错误，可能的错误类型包括：
//   - access token未设置
//   - 文件ID无效
//   - API调用失败
//   - 服务端返回错误
//
// API参考：https://docs.qq.com/open/document/app/openapi/v2/file/files/metadata.html
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
