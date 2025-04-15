package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chinahtl/tencent-doc-sdk/constant"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

// GetAuthURL 获取授权URL
/*
  指引地址：https://docs.qq.com/open/document/app/oauth2/authorize.html
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

// ExchangeToken 用授权码换取Token
/*
  指引地址：https://docs.qq.com/open/document/app/oauth2/access_token.html
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

// RefreshToken 刷新Token
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
