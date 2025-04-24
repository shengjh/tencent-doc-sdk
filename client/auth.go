package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chinahtl/tencent-doc-sdk/constant"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

// GetAuthURL 获取OAuth授权URL
/*
  功能：构造用于用户授权的URL地址
  文档指引：https://docs.qq.com/open/document/app/oauth2/authorize.html

  参数说明：
  - 无

  返回值：
  - string: 完整的授权URL字符串

  注意：
  - 会自动生成state参数用于防止CSRF攻击
  - 包含client_id, redirect_uri等必要参数
*/
func (c *Client) GetAuthURL() string {
	u, _ := url.Parse(constant.AuthEndpoint)
	query := url.Values{}
	query.Set("client_id", c.config.ClientID)
	query.Set("redirect_uri", c.config.RedirectURI)
	query.Set("response_type", "code")
	query.Set("scope", constant.AllScope)

	if c.config.RandomState != "" {
		query.Set("state", c.config.RandomState)
	} else {
		query.Set("state", util.GenerateRandomString(16))
	}

	u.RawQuery = query.Encode()
	return u.String()
}

// ExchangeToken 使用授权码换取访问令牌
/*
  功能：通过OAuth授权码交换access_token和refresh_token
  文档指引：https://docs.qq.com/open/document/app/oauth2/access_token.html

  参数：
  - ctx: 上下文，用于超时控制
  - code: 从授权回调获取的授权码

  返回值：
  - *model.TokenResponse: 令牌响应，包含access_token等信息
  - error: 错误信息

  特殊处理：
  - 如果配置了InitialToken，则直接返回初始令牌
*/
func (c *Client) ExchangeToken(ctx context.Context, code string) (*model.TokenResponse, error) {

	if c.config.InitialToken != nil && c.config.InitialToken.AccessToken != "" {
		return &model.TokenResponse{
			Token: model.Token{
				AccessToken:  c.config.InitialToken.AccessToken,
				RefreshToken: c.config.InitialToken.RefreshToken,
				UserID:       c.config.InitialToken.UserID,
			},
		}, nil
	}

	params := url.Values{}
	params.Set("client_id", c.config.ClientID)
	params.Set("client_secret", c.config.ClientSecret)
	params.Set("code", code)
	params.Set("grant_type", "authorization_code")
	params.Set("redirect_uri", c.config.RedirectURI)

	var result model.TokenResponse
	err := util.PostForm(ctx, c.httpClient, constant.TokenEndpoint, params, &result)
	if err != nil {
		return nil, fmt.Errorf("exchange token failed: %w", err)
	}

	return &result, nil
}

// RefreshToken 使用刷新令牌获取新的访问令牌
//
// 实现 OAuth 2.0 刷新令牌流程，官方文档参考：
// https://docs.qq.com/open/document/app/oauth2/refresh_token.html
//
// 参数:
//   - ctx: 上下文对象，用于超时和取消控制
//   - refreshToken: 之前获取的刷新令牌
//
// 返回值:
//   - *model.TokenResponse: 新的令牌响应，包含新的访问令牌和刷新令牌
//   - error: 操作失败时返回的错误信息
//
// 使用示例:
//
//	newToken, err := client.RefreshToken(context.Background(), "旧的刷新令牌")
//	if err != nil {
//	    // 处理错误
//	}
//	accessToken := newToken.Token.AccessToken
//
// 注意事项:
//   - 刷新令牌本身也有有效期，过期后将无法使用
//   - 新的刷新令牌可能与原令牌相同或不同
//   - 建议在访问令牌过期前主动刷新
func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (*model.TokenResponse, error) {
	params := url.Values{}
	params.Set("client_id", c.config.ClientID)
	params.Set("client_secret", c.config.ClientSecret)
	params.Set("refresh_token", refreshToken)
	params.Set("grant_type", "refresh_token")

	var result model.TokenResponse
	err := util.PostForm(ctx, c.httpClient, constant.TokenEndpoint, params, &result)
	if err != nil {
		return nil, fmt.Errorf("refresh token failed: %w", err)
	}

	return &result, nil
}
