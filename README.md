# Tencent Doc SDK for Go

腾讯文档 SDK Go 语言版本，提供简单易用的腾讯文档 API 调用接口。支持文档授权、文档操作、文档导出等功能。

## 功能特性

- ✨ 腾讯文档 API 支持
- 🔐 灵活的认证流程
- 📝 文档操作（列表、搜索）
- 📤 文档导出功能
- 🔄 自动 Token 刷新
- 🔧 可配置的 HTTP 客户端

## 安装

```bash
go get github.com/chinahtl/tencent-doc-sdk
```

## 快速开始

### 1. 初始化客户端

```go
import (
    "github.com/chinahtl/tencent-doc-sdk/client"
    "github.com/chinahtl/tencent-doc-sdk/config"
    "github.com/chinahtl/tencent-doc-sdk/util"
)

// 创建客户端（方式一：基本配置）
docClient := client.NewClient(
    config.WithClientID("your-client-id"),
    config.WithClientSecret("your-client-secret"),
    config.WithRedirectURI("your-redirect-uri"),
    config.WithTimeout(time.Second*30),
    config.WithRandomState(util.GenerateRandomString(20)),
)

// 创建客户端（方式二：使用已有 Token）
existingToken := &model.Token{
    AccessToken:  "your-access-token",
    RefreshToken: "your-refresh-token",
    ExpiresIn:    3600,
    UserID:       "user-id",
}

docClient := client.NewClient(
    config.WithClientID("your-client-id"),
    config.WithClientSecret("your-client-secret"),
    config.WithRedirectURI("your-redirect-uri"),
    config.WithInitialToken(existingToken),  // 设置初始 token
)
```

### 2. 授权认证

```go
// 获取授权 URL
authURL := docClient.GetAuthURL()

// 使用授权码交换 Token
tokenResp, err := docClient.ExchangeToken(context.Background(), "authorization-code")
if err != nil {
    log.Fatal(err)
}

// 刷新 Token
newTokenResp, err := docClient.RefreshToken(context.Background())
if err != nil {
    log.Fatal(err)
}
```

### 3. 文档操作

```go
// 列出文档
filter := &model.ListParams{
    // 设置过滤参数
}
docs, err := docClient.ListDocuments(context.Background(), filter)
if err != nil {
    log.Fatal(err)
}

// 搜索文档
searchFilter := &model.SearchParams{
    // 设置搜索参数
}
searchResults, err := docClient.SearchDocuments(context.Background(), searchFilter)
if err != nil {
    log.Fatal(err)
}
```

### 4. 导出功能

```go
// 获取导出进度
progress, err := docClient.GetExportProgress(context.Background(), "doc_id", "operation_id")
if err != nil {
    log.Fatal(err)
}
```

## 高级配置

### 自定义 HTTP 客户端

```go
customHTTPClient := &http.Client{
    Timeout: time.Second * 60,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 100,
    },
}
docClient.WithHTTPClient(customHTTPClient)
```


## 配置选项

| 配置项 | 说明 | 必填 | 默认值 |
|--------|------|------|--------|
| ClientID | 应用 ID | 是 | - |
| ClientSecret | 应用密钥 | 是 | - |
| RedirectURI | 重定向 URI | 是 | - |
| Timeout | HTTP 请求超时时间 | 否 | 30s |
| RandomState | 随机状态值 | 否 | 自动生成 |
| InitialToken | 初始 Token | 否 | nil |

## 注意事项

1. 请妥善保管 ClientID 和 ClientSecret，不要泄露
2. Token 有效期有限
3. 建议在生产环境中使用自定义的 HTTP 客户端配置
4. 错误处理建议使用 type assertion 来处理不同类型的错误

## 接口实现

本SDK已完整实现以下核心接口功能：

### 认证授权接口
- `GetAuthURL() string` - 获取授权URL
- `ExchangeToken(ctx context.Context, code string)` - 通过授权码交换Token
- `RefreshToken(ctx context.Context, refreshToken string)` - 刷新Access Token

### 文档操作接口
- `ListDocuments(ctx context.Context, params *model.ListParams)` - 列出用户文档
- `SearchDocuments(ctx context.Context, params *model.SearchParams)` - 搜索文档
- `GetFileMetadata(ctx context.Context, fileID string)` - 获取文件元数据

### 文档导出接口
- `ExportDocument(ctx context.Context, docID string, req *model.ExportRequest)` - 导出文档
- `GetExportProgress(ctx context.Context, docID string, operationID string)` - 查询导出进度


## 示例代码

完整的示例代码可以在 [example](./example) 目录下找到。

## 贡献指南

欢迎提交 Issue 和 Pull Request。在提交 PR 之前，请确保：

1. 代码符合 Go 语言规范
2. 添加了必要的测试用例
3. 更新了相关文档

## 许可证

MIT License

## 相关链接

- [腾讯文档开发者中心](https://docs.qq.com/open/)





