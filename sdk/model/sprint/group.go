package sprint

// SprintGroup 迭代分组模型（API 中称为 SprintSection）
type SprintGroup struct {
	ID       string            // 分组 ID
	URL      string            // 分组资源地址
	Name     string            // 分组名称
	Project  *ProjectSummary   // 所属项目
}

// SprintGroupCreateInput 创建迭代分组输入
type SprintGroupCreateInput struct {
	// 必填字段
	ProjectID string // 项目 ID
	Name      string // 分组名称
}

// SprintGroupUpdateInput 更新迭代分组输入（所有字段可选）
type SprintGroupUpdateInput struct {
	Name *string // 分组名称
}

// SprintGroupFilter 迭代分组列表过滤条件
type SprintGroupFilter struct {
	ProjectID string // 项目 ID（必填）

	// 分页参数
	PageSize  int // 每页大小（默认 30）
	PageIndex int // 页码（从 0 开始）
}
