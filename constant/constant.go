package constant

const (
	// AuthEndpoint 授权端点
	AuthEndpoint = "https://docs.qq.com/oauth/v2/authorize"
	// TokenEndpoint Token端点
	TokenEndpoint = "https://docs.qq.com/oauth/v2/token"
	// APIEndpoint API端点
	APIEndpoint = "https://docs.qq.com/openapi"
	// UserInfoEndpoint API端点
	UserInfoEndpoint = "https://docs.qq.com/oauth/v2/userinfo"
	// AllScope 全部权限
	AllScope = "all"
)

const (
	ListTypeFolder = "folder"
	ListTypeFile   = "file"
	ListTypeAll    = "all"

	SortTypeBrowse = "browse"
	SortTypeTime   = "time"
	SortTypeName   = "name"

	FileTypeDoc   = "doc"
	FileTypeSheet = "sheet"
	FileTypeSlide = "slide"

	ExportTypePDF  = "pdf"
	ExportTypeDocx = "docx"
	ExportTypeXlsx = "xlsx"
	ExportTypePptx = "pptx"
)
