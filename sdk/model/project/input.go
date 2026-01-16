package project

// ProjectCreateInput 创建项目输入
type ProjectCreateInput struct {
	// 必填字段
	Name       string // 项目名称（最大255字符）
	Type       string // 项目类型：scrum, kanban, waterfall, hybrid
	Identifier string // 项目标识（唯一，不超过15字符）

	// 可选字段
	ScopeType  string // 所属类型：organization, user_group（默认 organization）
	ScopeID    string // 所属 ID（当 scope_type=user_group 时必填）
	Visibility string // 可见性：private, public（默认 private）
	Description string // 项目描述

	// 时间相关
	StartAt    *int64 // 开始时间（10位时间戳，可选）
	EndAt      *int64 // 结束时间（10位时间戳，可选）

	// 人员相关
	AssigneeID string // 项目负责人 ID（可选）
	Members    []ProjectMemberInput // 项目成员列表（可选）
}

// ProjectMemberInput 项目成员输入
type ProjectMemberInput struct {
	ID   string // 成员 ID（用户或团队 ID）
	Type string // 类型：user, user_group
}

// ProjectUpdateInput 更新项目输入（所有字段可选）
type ProjectUpdateInput struct {
	Name        *string  // 项目名称
	Identifier  *string  // 项目标识
	Description *string  // 项目描述
	StartAt     *int64   // 开始时间
	EndAt       *int64   // 结束时间
	AssigneeID  *string  // 项目负责人 ID
	StateID     *string  // 项目状态 ID
	Properties  map[string]string // 自定义属性
}
