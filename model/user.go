package model

// UserInfo 用户信息
type UserInfo struct {
	OpenID  string `json:"openID"`
	Nick    string `json:"nick"`
	Avatar  string `json:"avatar"`
	Source  string `json:"source"`
	UnionID string `json:"unionID"`
}

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	Ret  int      `json:"ret"`
	Msg  string   `json:"msg"`
	Data UserInfo `json:"data"`
}
