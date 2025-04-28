package client

import (
	"context"
	"fmt"

	"github.com/chinahtl/tencent-doc-sdk/constant"
	"github.com/chinahtl/tencent-doc-sdk/model"
	"github.com/chinahtl/tencent-doc-sdk/util"
)

// GetUserInfo 获取当前用户信息。
//
// ctx 用于控制请求的上下文，可用于超时控制和取消。
//
// 返回用户信息，包括：
//   - OpenID: 用户的唯一标识
//   - Nick: 用户昵称
//   - Avatar: 用户头像URL
//   - Source: 用户来源
//   - UnionID: 用户的UnionID
//
// 可能返回的错误：
//   - access token未设置
//   - API调用失败
//   - 服务端返回错误
//
// API参考：https://docs.qq.com/oauth/v2/userinfo
func (c *Client) GetUserInfo(ctx context.Context) (*model.UserInfo, error) {
	if c.token.AccessToken == "" {
		return nil, fmt.Errorf("access token not set")
	}

	url := fmt.Sprintf("%s?access_token=%s", constant.UserInfoEndpoint, c.token.AccessToken)

	var resp model.UserInfoResponse
	if err := util.HTTPGet(ctx, url, &resp); err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	if resp.Ret != 0 {
		return nil, fmt.Errorf("get user info failed with ret %d: %s", resp.Ret, resp.Msg)
	}

	return &resp.Data, nil
}
