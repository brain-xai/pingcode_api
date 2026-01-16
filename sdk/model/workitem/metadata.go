package workitem

// WorkItemField 工作项属性定义
type WorkItemField struct {
	ID          string                 // 属性 ID
	URL         string                 // 属性资源地址
	Name        string                 // 属性名称
	Type        string                 // 属性类型：text, select, number, date 等
	Description string                 // 属性描述
	Options     []WorkItemFieldOption  // 选项列表（当 type=select 时）
	Properties  map[string]string      // 自定义属性
}

// WorkItemFieldOption 属性选项
type WorkItemFieldOption struct {
	ID   string // 选项 ID
	Text string // 选项文本
}

// WorkItemFieldList 工作项属性列表
type WorkItemFieldList struct {
	PageSize  int             // 页面大小
	PageIndex int             // 页码
	Total     int             // 总数
	Values    []WorkItemField // 属性列表
}

// WorkItemPriority 工作项优先级
type WorkItemPriority struct {
	ID     string // 优先级 ID
	Name   string // 优先级名称：P0, P1, P2, P3, P4
	Color  string // 优先级颜色
	Icon   string // 优先级图标
	Level  int    // 优先级级别
}

// WorkItemPriorityList 优先级列表
type WorkItemPriorityList struct {
	PageSize  int                  // 页面大小
	PageIndex int                  // 页码
	Total     int                  // 总数
	Values    []WorkItemPriority   // 优先级列表
}

// WorkItemTag 工作项标签
type WorkItemTag struct {
	ID          string                 // 标签 ID
	URL         string                 // 标签资源地址
	Name        string                 // 标签名称
	Color       string                 // 标签颜色
	Description string                 // 标签描述
	Properties  map[string]string      // 自定义属性
}

// WorkItemTagList 标签列表
type WorkItemTagList struct {
	PageSize  int             // 页面大小
	PageIndex int             // 页码
	Total     int             // 总数
	Values    []WorkItemTag   // 标签列表
}
