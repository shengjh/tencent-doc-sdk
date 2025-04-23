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
	SearchKey   string `json:"searchKey"`   // 搜索关键字(必填)
	SearchType  string `json:"searchType"`  // 搜索类型(必填): title-标题搜索 owner-所有者搜索
	ResultType  string `json:"resultType"`  // 返回结果类型: all-所有 folder-只返回文件夹(默认all)
	FolderID    string `json:"folderID"`    // 搜索范围文件夹ID，空表示所有文件
	Offset      int    `json:"offset"`      // 起始偏移量(默认0)
	Size        int    `json:"size"`        // 返回条目数量(默认20，最大50)
	SortType    string `json:"sortType"`    // 排序规则: modify-修改时间(默认) create-创建时间 browse-访问时间
	Asc         int    `json:"asc"`         // 排序顺序: 1-正序 0-倒序(默认)
	ByOwnership int    `json:"byOwnership"` // 所有者过滤: 1-是所有者 0-否(默认)
	FileTypes   string `json:"fileTypes"`   // 文件类型过滤，多种用"-"分隔，如"doc-sheet"
}

// FileMetadataResponse 文件元数据响应结构
type FileMetadataResponse struct {
	Ret  int    `json:"ret"` // 返回码
	Msg  string `json:"msg"` // 返回信息
	Data struct {
		ID             string `json:"ID"`             // 文件唯一标识
		Title          string `json:"title"`          // 文件标题
		Type           string `json:"type"`           // 文件类型(doc/sheet/ppt等)
		URL            string `json:"url"`            // 文件访问URL
		Status         string `json:"status"`         // 文件状态(normal/deleted等)
		IsCreator      bool   `json:"isCreator"`      // 当前用户是否为创建者
		CreateTime     int64  `json:"createTime"`     // 创建时间戳(秒级)
		CreatorName    string `json:"creatorName"`    // 创建者名称
		IsOwner        bool   `json:"isOwner"`        // 当前用户是否为所有者
		OwnerName      string `json:"ownerName"`      // 所有者名称
		LastModifyTime int64  `json:"lastModifyTime"` // 最后修改时间戳(秒级)
		LastModifyName string `json:"lastModifyName"` // 最后修改者名称
		OwnerID        string `json:"ownerID"`        // 所有者ID
	} `json:"data"` // 文件元数据详情
}
