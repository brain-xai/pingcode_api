package project

// Project 项目模型（简化版，仅包含核心字段）
type Project struct {
	ID          string // 项目 ID
	Identifier  string // 项目标识（如 "SCR"）
	Name        string // 项目名称
	Type        string // 项目类型：scrum, kanban, waterfall, hybrid
	Description string // 项目描述
	State       string // 项目状态名称
	StartAt     int64  // 开始时间（10位时间戳）
	EndAt       int64  // 结束时间
	CreatedAt   int64  // 创建时间
	UpdatedAt   int64  // 更新时间
}

// ProjectFilter 项目列表过滤条件
type ProjectFilter struct {
	Identifier      string // 项目标识（可选）
	Type            string // 项目类型（可选）
	IncludeDeleted  bool   // 是否包含已删除
	IncludeArchived bool   // 是否包含已归档
}
