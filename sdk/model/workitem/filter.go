package workitem

// WorkItemFilter 工作项列表过滤条件
type WorkItemFilter struct {
	// 基础过滤
	ProjectID string // 项目 ID（常用过滤）
	Query     string // 搜索关键字

	// 分类过滤
	TypeID     string // 工作项类型 ID
	StateID    string // 状态 ID
	PriorityID string // 优先级 ID

	// 人员过滤
	AssigneeID string // 负责人 ID
	ReporterID string // 报告人 ID

	// 层级过滤
	ParentID string // 父工作项 ID

	// 时间过滤
	StartAtFrom  int64 // 开始时间-起（10位时间戳）
	StartAtTo    int64 // 开始时间-止
	EndAtFrom    int64 // 结束时间-起
	EndAtTo      int64 // 结束时间-止
	CreatedAfter int64 // 创建时间之后
	UpdatedAfter int64 // 更新时间之后

	// Sprint 和版本
	SprintID  string // Sprint ID
	VersionID string // 版本 ID

	// 标签过滤
	TagIDs []string // 标签 ID 列表

	// 状态标记
	IncludeArchived bool // 是否包含已归档
	IncludeDeleted  bool // 是否包含已删除

	// 分页
	PageSize  int // 每页大小（可选，默认30）
	PageIndex int // 页码（从0开始）
}

// WorkItemTagFilter 标签列表过滤条件
type WorkItemTagFilter struct {
	ProjectID string // 项目 ID
	TypeID    string // 工作项类型 ID（可选）
	PageSize  int    // 每页大小
	PageIndex int    // 页码
}

// WorkItemStatusFilter 状态列表过滤条件
type WorkItemStatusFilter struct {
	ProjectID string // 项目 ID（必填）
	TypeID    string // 工作项类型 ID（可选）
	PageSize  int    // 每页大小
	PageIndex int    // 页码
}

// WorkItemFieldFilter 属性列表过滤条件
type WorkItemFieldFilter struct {
	ProjectID string // 项目 ID（必填）
	TypeID    string // 工作项类型 ID（可选）
	PageSize  int    // 每页大小
	PageIndex int    // 页码
}

// WorkItemLinkFilter 关联列表过滤条件
type WorkItemLinkFilter struct {
	TypeID    string // 关联类型 ID（可选）
	PageSize  int    // 每页大小
	PageIndex int    // 页码
}

// WorkItemDeliveryTargetFilter 交付目标列表过滤条件
type WorkItemDeliveryTargetFilter struct {
	WorkItemID string // 工作项 ID
	Status     string // 状态（可选）
	PageSize   int    // 每页大小
	PageIndex  int    // 页码
}
