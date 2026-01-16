package sprint

// Sprint 迭代模型
type Sprint struct {
	ID                   string              // 迭代 ID
	URL                  string              // 迭代资源地址
	Name                 string              // 迭代名称
	Status               string              // 迭代状态：pending, in_progress, completed
	StartAt              int64               // 开始时间（10位时间戳）
	EndAt                int64               // 结束时间
	Description          string              // 描述
	StartedAt            int64               // 实际开始时间
	CompletedAt          int64               // 实际完成时间
	TotalStoryPoints     float64             // 总故事点
	StartedStoryPoints   float64             // 已开始故事点
	CompletedStoryPoints float64             // 已完成故事点
	CreatedAt            int64               // 创建时间
	UpdatedAt            int64               // 更新时间

	// 关联对象
	Project    *ProjectSummary    // 所属项目
	Assignee   *User              // 负责人
	Categories []SprintCategory   // 迭代类别列表
	CreatedBy  *User              // 创建人
	UpdatedBy  *User              // 更新人
}

// ProjectSummary 项目摘要
type ProjectSummary struct {
	ID         string // 项目 ID
	URL        string // 项目资源地址
	Identifier string // 项目标识
	Name       string // 项目名称
	Type       string // 项目类型
	IsArchived int    // 是否已归档（0/1）
	IsDeleted  int    // 是否已删除（0/1）
}

// User 用户
type User struct {
	ID          string // 用户 ID
	URL         string // 用户资源地址
	Name        string // 用户名
	DisplayName string // 显示名称
	Avatar      string // 头像
}

// SprintCategory 迭代类别摘要
type SprintCategory struct {
	ID   string // 类别 ID
	URL  string // 类别资源地址
	Name string // 类别名称
}

// SprintFilter 迭代列表过滤条件
type SprintFilter struct {
	ProjectID      string // 项目 ID（必填）
	Name           string // 名称过滤（可选）
	Status         string // 状态过滤（可选）：pending, in_progress, completed
	CreatedBetween string // 创建时间范围 "start,end"（可选）
	UpdatedBetween string // 更新时间范围 "start,end"（可选）

	// 分页参数
	PageSize  int // 每页大小（默认 30）
	PageIndex int // 页码（从 0 开始）
}
