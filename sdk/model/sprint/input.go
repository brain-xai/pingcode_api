package sprint

// SprintCreateInput 创建迭代输入
type SprintCreateInput struct {
	// 必填字段
	ProjectID  string   // 项目 ID
	Name       string   // 迭代名称
	StartAt    int64    // 开始时间（10位时间戳）
	EndAt      int64    // 结束时间
	AssigneeID string   // 负责人 ID

	// 可选字段
	Description string  // 描述
	Status      string  // 状态：pending（默认）, in_progress, completed
	CategoryIDs []string // 类别 ID 列表
}

// SprintUpdateInput 更新迭代输入（所有字段可选）
type SprintUpdateInput struct {
	Name        *string   // 迭代名称
	StartAt     *int64    // 开始时间
	EndAt       *int64    // 结束时间
	AssigneeID  *string   // 负责人 ID
	Description *string   // 描述
	Status      *string   // 状态
	CategoryIDs []string  // 类别 ID 列表
}

// SprintBatchCreateInput 批量创建迭代输入
type SprintBatchCreateInput struct {
	Sprints []SprintCreateInput // 迭代列表（最多 100 个）
}
