package ship

// Requirement 需求模型（API 中称为 Idea）
type Requirement struct {
	ID         string // 需求 ID
	URL        string // 需求资源地址
	Identifier string // 需求编号（如 "SLC-1"）
	Title      string // 需求标题
	ShortID    string // 短 ID（如 "Ogf1EYey"）
	HTMLURL    string // HTML 链接

	// 关联对象
	Product   *ProductSummary      // 所属产品（引用）
	Assignee  *User                // 负责人
	State     *RequirementState    // 状态
	Priority  *RequirementPriority // 优先级
	Plan      *Plan                // 产品排期
	Suite     *Suite               // 产品模块

	// 时间相关
	PlanAt *TimeRange // 计划时间
	RealAt *TimeRange // 实际时间

	// 数值字段
	Score    float64 // 分数
	Progress float64 // 进度（0-1的小数）

	// 内容字段
	Description string            // 描述
	Properties  map[string]string // 自定义属性

	// 参与人
	Participants []Participant // 参与人列表

	// 完成信息
	CompletedAt *int64 // 完成时间
	CompletedBy *User  // 完成人

	// 创建/更新信息
	CreatedAt int64  // 创建时间
	UpdatedAt int64  // 更新时间
	CreatedBy *User  // 创建人
	UpdatedBy *User  // 更新人

	// 状态
	IsArchived bool // 是否已归档
	IsDeleted  bool // 是否已删除
}

// ProductSummary 产品简要信息（用于需求中的引用）
type ProductSummary struct {
	ID         string // 产品 ID
	Identifier string // 产品标识
	Name       string // 产品名称
	IsArchived bool   // 是否已归档
	IsDeleted  bool   // 是否已删除
}

// RequirementState 需求状态
type RequirementState struct {
	ID   string // 状态 ID
	Name string // 状态名称
	Type string // 状态类型：pending, in_progress, completed 等
}

// RequirementPriority 需求优先级
type RequirementPriority struct {
	ID   string // 优先级 ID
	Name string // 优先级名称：P0, P1, P2, P3, P4
}

// Plan 产品排期
type Plan struct {
	ID   string // 排期 ID
	Name string // 排期名称
	URL  string // 排期资源地址
}

// Suite 产品模块
type Suite struct {
	ID   string // 模块 ID
	Name string // 模块名称
	Type string // 模块类型：module, category 等
	URL  string // 模块资源地址
}

// TimeRange 时间范围
type TimeRange struct {
	From        int64  // 开始时间（10位时间戳）
	To          int64  // 结束时间
	Granularity string // 时间单位：year, quarter, month, day
}

// Participant 参与人
type Participant struct {
	ID        string     // 参与人 ID
	Type      string     // 类型：user, user_group
	User      *User      // 用户
	UserGroup *UserGroup // 团队
}

// RequirementFilter 需求列表过滤条件
type RequirementFilter struct {
	ProductID       string // 产品 ID（常用过滤）
	StateID         string // 需求状态 ID（可选）
	PriorityID      string // 需求优先级 ID（可选）
	Keywords        string // 搜索关键字（可选）
	CreatedBetween  string // 创建时间范围，逗号分隔（可选）
	UpdatedBetween  string // 更新时间范围，逗号分隔（可选）
	IncludePublicImageToken string // 包含图片资源token（可选）
	PageSize         int    // 每页大小（可选，默认30）
	PageIndex        int    // 页码（可选，从0开始）
}

// RequirementCreateInput 创建需求入参
type RequirementCreateInput struct {
	ProductID   string                 // 产品 ID（必填）
	Title       string                 // 标题（必填，最大255）
	AssigneeID  string                 // 负责人 ID（可选）
	Description string                 // 描述（可选）
	SuiteID     string                 // 产品模块 ID（可选）
	Properties  map[string]interface{} // 自定义属性（可选）
}

// RequirementUpdateInput 更新需求入参（所有字段可选）
type RequirementUpdateInput struct {
	Title       *string                // 标题（可选）
	Description *string                // 描述（可选）
	StateID     *string                // 状态 ID（可选）
	PriorityID  *string                // 优先级 ID（可选）
	AssigneeID  *string                // 负责人 ID（可选）
	Progress    *float64               // 进度 0-1（可选）
	PlanAt      *TimeRange             // 计划时间（可选，整体更新）
	RealAt      *TimeRange             // 实际时间（可选，整体更新）
	PlanID      *string                // 产品排期 ID（可选）
	SuiteID     *string                // 产品模块 ID（可选）
	Properties  map[string]interface{} // 自定义属性（可选）
}

// RequirementStateList 需求状态列表
type RequirementStateList struct {
	PageSize  int                 // 页面大小
	PageIndex int                 // 页码
	Total     int                 // 总数
	Values    []RequirementState  // 状态列表
}

// RequirementPriorityList 需求优先级列表
type RequirementPriorityList struct {
	PageSize  int                    // 页面大小
	PageIndex int                    // 页码
	Total     int                    // 总数
	Values    []RequirementPriority  // 优先级列表
}

// RequirementPropertyOption 需求属性选项
type RequirementPropertyOption struct {
	ID   string // 选项 ID
	Text string // 选项文本
}

// RequirementProperty 需求属性定义
type RequirementProperty struct {
	ID      string                       // 属性 ID
	URL     string                       // 属性资源地址
	Name    string                       // 属性名称
	Type    string                       // 属性类型：text, select, number 等
	Options []RequirementPropertyOption  // 选项列表（当 type=select 时）
}

// RequirementPropertyList 需求属性列表
type RequirementPropertyList struct {
	PageSize  int                    // 页面大小
	PageIndex int                    // 页码
	Total     int                    // 总数
	Values    []RequirementProperty  // 属性列表
}

// RequirementSuiteList 需求模块列表
type RequirementSuiteList struct {
	PageSize  int     // 页面大小
	PageIndex int     // 页码
	Total     int     // 总数
	Values    []Suite // 模块列表
}

// RequirementPlanList 需求排期列表
type RequirementPlanList struct {
	PageSize  int    // 页面大小
	PageIndex int    // 页码
	Total     int    // 总数
	Values    []Plan // 排期列表
}
