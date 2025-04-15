package model

// ExportRequest 文档导出请求参数
type ExportRequest struct {
	ExportType string `json:"exportType"` // pdf/docx/xlsx等
}

// ExportResponse 文档导出响应
type ExportResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		OperationID string `json:"operationID"` // 异步操作ID
	} `json:"data"`
}

// ExportProgressResponse 导出进度响应
type ExportProgressResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		URL      string `json:"url"`      // 下载URL(进度100%时返回)
		Progress int    `json:"progress"` // 当前进度(0-100)
	} `json:"data"`
}
