package model

// Token 访问令牌
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	UserID       string `json:"user_id"`
	Scope        string `json:"scope"`
}

// TokenResponse Token响应
type TokenResponse struct {
	Token
	Scope string `json:"scope"`
}
