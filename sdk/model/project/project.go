package project

// Project 项目模型（完整版）
type Project struct {
	ID          string       // 项目 ID
	Identifier  string       // 项目标识（如 "SCR"）
	Name        string       // 项目名称
	Type        string       // 项目类型：scrum, kanban, waterfall, hybrid
	Description string       // 项目描述
	State       string       // 项目状态名称
	StartAt     int64        // 开始时间（10位时间戳）
	EndAt       int64        // 结束时间
	CreatedAt   int64        // 创建时间
	UpdatedAt   int64        // 更新时间

	// 扩展字段
	Assignee    *User           `json:"assignee"`     // 项目负责人
	ScopeType   string          `json:"scope_type"`   // 所属类型：organization, user_group
	ScopeID     string          `json:"scope_id"`     // 所属 ID
	Visibility  string          `json:"visibility"`   // 可见性：private, public
	Color       string          `json:"color"`        // 主题色
	Properties  map[string]string `json:"properties"` // 自定义属性
	Members     []ProjectMember `json:"members"`      // 成员列表
	CreatedBy   *User           `json:"created_by"`   // 创建人
	UpdatedBy   *User           `json:"updated_by"`   // 更新人
	IsArchived  bool            `json:"is_archived"`  // 是否已归档
	IsDeleted   bool            `json:"is_deleted"`   // 是否已删除
}

// User 用户引用
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

// UserGroup 团队引用
type UserGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProjectMember 项目成员
type ProjectMember struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`      // user, user_group
	User      *User       `json:"user"`      // 当 type=user 时
	UserGroup *UserGroup  `json:"user_group"` // 当 type=user_group 时
}

// State 项目状态
type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Progress 项目进度
type Progress struct {
	CompletedCount int     `json:"completed_count"`
	TotalCount     int     `json:"total_count"`
	CompletionRate float64 `json:"completion_rate"`
}

// ProjectFilter 项目列表过滤条件
type ProjectFilter struct {
	Identifier      string // 项目标识（可选）
	Type            string // 项目类型（可选）
	IncludeDeleted  bool   // 是否包含已删除
	IncludeArchived bool   // 是否包含已归档
}
