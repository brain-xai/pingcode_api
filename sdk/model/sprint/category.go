package sprint

// SprintCategoryFull 迭代类别完整模型
type SprintCategoryFull struct {
	ID        string          // 类别 ID
	URL       string          // 类别资源地址
	Name      string          // 类别名称
	Project   *ProjectSummary // 所属项目
	Group     *SprintGroup    // 所属分组（API 中称为 section）
	GroupID   string          // 所属分组 ID
}

// SprintCategoryCreateInput 创建迭代类别输入
type SprintCategoryCreateInput struct {
	// 必填字段
	ProjectID string // 项目 ID
	Name      string // 类别名称

	// 可选字段
	GroupID string // 所属分组 ID（API 中称为 section_id）
}

// SprintCategoryUpdateInput 更新迭代类别输入（所有字段可选）
type SprintCategoryUpdateInput struct {
	Name    *string // 类别名称
	GroupID *string // 所属分组 ID
}

// SprintCategoryFilter 迭代类别列表过滤条件
type SprintCategoryFilter struct {
	ProjectID string // 项目 ID（必填）

	// 分页参数
	PageSize  int // 每页大小（默认 30）
	PageIndex int // 页码（从 0 开始）
}
