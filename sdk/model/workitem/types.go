package workitem

// WorkItemType 工作项类型
type WorkItemType struct {
	ID          string                 // 类型 ID
	URL         string                 // 类型资源地址
	Name        string                 // 类型名称
	Icon        string                 // 类型图标
	Description string                 // 类型描述
	Properties  map[string]string      // 自定义属性
}

// WorkItemTypeScope 工作项类型范围
type WorkItemTypeScope struct {
	ProjectID string // 项目 ID（必填）
}

// WorkItemStatus 工作项状态
type WorkItemStatus struct {
	ID       string // 状态 ID
	URL      string // 状态资源地址
	Name     string // 状态名称
	Type     string // 状态类型：pending, in_progress, completed 等
	Color    string // 状态颜色
	Category string // 状态分类
}

// WorkItemStatusList 工作项状态列表
type WorkItemStatusList struct {
	PageSize  int               // 页面大小
	PageIndex int               // 页码
	Total     int               // 总数
	Values    []WorkItemStatus  // 状态列表
}
