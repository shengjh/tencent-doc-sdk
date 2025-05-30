package client

import (
	"context"
	"net/http"

	"github.com/chinahtl/tencent-doc-sdk/config"
	"github.com/chinahtl/tencent-doc-sdk/model"
)

// TencentDocClient 腾讯文档客户端接口
type TencentDocClient interface {
	// 授权相关
	GetAuthURL() string
	ExchangeToken(ctx context.Context, code string) (*model.TokenResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*model.TokenResponse, error)

	// 用户相关
	GetUserInfo(ctx context.Context) (*model.UserInfo, error)
	// 文档操作
	ListDocuments(ctx context.Context, params *model.ListParams) (*model.ListDocumentsResponse, error)
	SearchDocuments(ctx context.Context, params *model.SearchParams) (*model.SearchDocumentsResponse, error)
	GetFileMetadata(ctx context.Context, fileID string) (*model.FileMetadataResponse, error)

	// 导出相关
	ExportDocument(ctx context.Context, docID string, req *model.ExportRequest) (*model.ExportResponse, error)
	GetExportProgress(ctx context.Context, docID string, operationID string) (*model.ExportProgressResponse, error)
}

// Client 实现 TencentDocClient 接口
type Client struct {
	config     *config.Config
	httpClient *http.Client
	token      *model.Token
}

// 确保 Client 实现 TencentDocClient 接口
var _ TencentDocClient = (*Client)(nil)

// WithToken 设置访问令牌
func (c *Client) WithToken(token *model.Token) *Client {
	c.token = token
	return c
}

// NewClient 创建新的客户端实例
func NewClient(opts ...config.Option) *Client {
	cfg := config.DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	httpClient := &http.Client{Timeout: cfg.Timeout}

	// 如果配置了自定义 Transport，则使用它
	if cfg.Transport != nil {
		httpClient.Transport = cfg.Transport
	}

	client := &Client{
		config:     cfg,
		httpClient: httpClient,
	}

	// 如果提供了初始 Token，则设置它
	if cfg.InitialToken != nil {
		client.token = cfg.InitialToken
	}

	return client
}
