package model

// Document 文档/文件夹信息
type Document struct {
	ID             string `json:"ID"`
	Title          string `json:"title"`
	Type           string `json:"type"` // folder/doc/sheet等
	URL            string `json:"url"`
	Status         string `json:"status"`
	FileSource     string `json:"fileSource"` // personal/external
	IsCreator      bool   `json:"isCreator"`
	CreatorName    string `json:"creatorName"`
	IsOwner        bool   `json:"isOwner"`
	OwnerName      string `json:"ownerName"`
	CreateTime     int64  `json:"createTime"`     // 时间戳
	LastModifyTime int64  `json:"lastModifyTime"` // 时间戳
	LastBrowseTime int64  `json:"lastBrowseTime"` // 时间戳
	Starred        bool   `json:"starred,omitempty"`
	Pinned         bool   `json:"pinned,omitempty"`
	IsCollaborated bool   `json:"isCollaborated"`
}

// ListDocumentsResponse 文档列表响应
type ListDocumentsResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Next int         `json:"next"` // 下一次请求的起始位置
		List []*Document `json:"list"` // 文档列表
	} `json:"data"`
}

// ListParams 列表参数
type ListParams struct {
	ListType string `json:"listType"` // folder/file/all
	SortType string `json:"sortType"` // browse/time/name
	Asc      int    `json:"asc"`      // 1:正序 0:倒序
	FolderID string `json:"folderID"` // 文件夹ID
	Start    int    `json:"start"`    // 起始位置
	Limit    int    `json:"limit"`    // 每页数量
	IsOwner  int    `json:"isOwner"`  // 1:所有文件 2:仅自己拥有的
	FileType string `json:"fileType"` // 文件类型过滤(多个用-分隔)
}

// SearchDocument 搜索到的文档信息
type SearchDocument struct {
	ID             string `json:"ID"`
	Title          string `json:"title"`
	Type           string `json:"type"` // doc/sheet/slide等
	URL            string `json:"url"`
	Status         string `json:"status"`
	OwnerName      string `json:"ownerName"`
	FileSource     string `json:"fileSource"` // enterprise/personal
	Highlight      string `json:"highlight,omitempty"`
	LastModifyTime int64  `json:"lastModifyTime"` // 时间戳
	LastModifyName string `json:"lastModifyName"`
	CreateTime     int64  `json:"createTime"` // 时间戳
}

// SearchDocumentsResponse 文档搜索响应
type SearchDocumentsResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Next    int               `json:"next"`    // 下一次请求的偏移量
		Total   int               `json:"total"`   // 总匹配数
		HasMore bool              `json:"hasMore"` // 是否还有更多
		List    []*SearchDocument `json:"list"`    // 文档列表
	} `json:"data"`
}

// SearchParams 搜索参数
type SearchParams struct {
	SearchType  string `json:"searchType"`  // title/content
	SearchKey   string `json:"searchKey"`   // 搜索关键词
	FolderID    string `json:"folderID"`    // 文件夹ID
	ByOwnership int    `json:"byOwnership"` // 按所有权过滤
	Offset      int    `json:"offset"`      // 偏移量
	Size        int    `json:"size"`        // 每页大小
}
