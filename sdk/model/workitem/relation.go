package workitem

// WorkItemLinkType 工作项关联类型
type WorkItemLinkType struct {
	ID          string                 // 关联类型 ID
	Name        string                 // 关联类型名称
	Category    string                 // 关联类型分类：duplicate, relate 等
	IsSystem    bool                   // 是否系统内置
	Properties  map[string]string      // 自定义属性
}

// WorkItemLink 工作项关联
type WorkItemLink struct {
	ID           string              // 关联 ID
	URL          string              // 关联资源地址
	RelationType string              // 关联类型
	Source       *WorkItemSummary    // 源工作项
	Target       *WorkItemSummary    // 目标工作项
	Properties   map[string]string   // 自定义属性
}
