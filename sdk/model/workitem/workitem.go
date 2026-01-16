package workitem

// WorkItem 工作项模型
type WorkItem struct {
	// 基础字段
	ID          string // 工作项 ID
	URL         string // 工作项资源地址
	Identifier  string // 工作项编号（如 "SCR-1"）
	ShortID     string // 短 ID（如 "eqWqLmTO"）
	HTMLURL     string // HTML 链接
	Title       string // 标题
	Description string // 描述

	// 项目关联
	Project *ProjectSummary // 所属项目

	// 分类字段
	Type     *WorkItemType     // 工作项类型
	State    *WorkItemStatus   // 工作项状态
	Priority *WorkItemPriority // 优先级

	// 层级关系
	ParentID string          // 父工作项 ID
	Parent   *WorkItemSummary // 父工作项

	// 人员字段
	AssigneeID     string // 负责人 ID
	Assignee       *User  // 负责人
	ReporterID     string // 报告人 ID
	Reporter       *User  // 报告人
	ParticipantIDs []string // 参与人 ID 列表
	Participants   []Participant // 参与人列表

	// 时间字段
	StartAt     int64  // 开始时间（10位时间戳）
	EndAt       int64  // 结束时间
	CompletedAt *int64 // 完成时间

	// Sprint 和版本
	SprintID  string  // Sprint ID
	Sprint    *Sprint // Sprint
	VersionID string  // 版本 ID
	Version   *Version // 版本

	// 看板和泳道
	BoardID    string    // 看板 ID
	Board      *Board    // 看板
	EntryID    string    // 看板栏 ID
	Entry      *Entry    // 看板栏
	SwimlaneID string    // 泳道 ID
	Swimlane   *Swimlane // 泳道
	Phase      string    // 所属阶段 ID

	// 数值字段
	StoryPoints        float64 // 故事点
	EstimatedWorkload  float64 // 预估工时
	RemainingWorkload  float64 // 剩余工时

	// 自定义属性
	Properties map[string]string // 自定义属性

	// 标签
	Tags              []WorkItemTag // 标签列表
	PublicImageToken  string        // 图片资源 token

	// 创建/更新信息
	CreatedAt int64  // 创建时间
	UpdatedAt int64  // 更新时间
	CreatedBy *User  // 创建人
	UpdatedBy *User  // 更新人

	// 状态
	IsArchived bool // 是否已归档
	IsDeleted  bool // 是否已删除
}

// ProjectSummary 项目简要信息
type ProjectSummary struct {
	ID         string // 项目 ID
	Identifier string // 项目标识
	Name       string // 项目名称
	Type       string // 项目类型
}

// WorkItemSummary 工作项简要信息（用于嵌套引用）
type WorkItemSummary struct {
	ID         string           // 工作项 ID
	Identifier string           // 工作项编号
	Title      string           // 标题
	Type       *WorkItemType    // 工作项类型
	State      *WorkItemStatus  // 工作项状态
	Properties map[string]string // 自定义属性
}

// Sprint Sprint 引用
type Sprint struct {
	ID   string // Sprint ID
	Name string // Sprint 名称
}

// Version 版本引用
type Version struct {
	ID   string // 版本 ID
	Name string // 版本名称
}

// Board 看板引用
type Board struct {
	ID   string // 看板 ID
	Name string // 看板名称
}

// Entry 看板栏引用
type Entry struct {
	ID   string // 看板栏 ID
	Name string // 看板栏名称
}

// Swimlane 泳道引用
type Swimlane struct {
	ID   string // 泳道 ID
	Name string // 泳道名称
}

// User 用户引用
type User struct {
	ID          string // 用户 ID
	Name        string // 用户名
	DisplayName string // 显示名称
	Avatar      string // 头像
}

// UserGroup 团队引用
type UserGroup struct {
	ID   string // 团队 ID
	Name string // 团队名称
}

// Participant 参与人
type Participant struct {
	ID        string     // 参与人 ID
	Type      string     // 类型：user, user_group
	User      *User      // 用户
	UserGroup *UserGroup // 团队
}
