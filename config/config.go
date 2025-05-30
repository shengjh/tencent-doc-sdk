// config/config.go
package config

import (
	"net/http"
	"time"

	"github.com/chinahtl/tencent-doc-sdk/model"
)

// Config 客户端配置
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	RandomState  string
	Timeout      time.Duration
	InitialToken *model.Token      // 新增初始 Token 字段
	Transport    http.RoundTripper // 自定义 HTTP Transport
}

// Option 定义配置选项函数类型
type Option func(*Config)

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Timeout: 30 * time.Second,
	}
}

// WithClientID 设置客户端ID
func WithClientID(clientID string) Option {
	return func(c *Config) {
		c.ClientID = clientID
	}
}

// WithClientSecret 设置客户端密钥
func WithClientSecret(clientSecret string) Option {
	return func(c *Config) {
		c.ClientSecret = clientSecret
	}
}

// WithRedirectURI 设置重定向URI
func WithRedirectURI(redirectURI string) Option {
	return func(c *Config) {
		c.RedirectURI = redirectURI
	}
}

// WithTimeout 设置超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithRandomState 设置state
func WithRandomState(randomState string) Option {
	return func(c *Config) {
		c.RandomState = randomState
	}
}

// WithInitialToken 设置初始 Token
func WithInitialToken(token *model.Token) Option {
	return func(c *Config) {
		c.InitialToken = token
	}
}

// WithHttpTransport 设置自定义的 HTTP Transport
func WithHttpTransport(transport http.RoundTripper) Option {
	return func(c *Config) {
		c.Transport = transport
	}
}
