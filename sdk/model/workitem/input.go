package workitem

// WorkItemCreateInput 创建工作项输入
type WorkItemCreateInput struct {
	// 必填字段
	ProjectID string // 项目 ID（必填）
	TypeID    string // 工作项类型 ID（必填）
	Title     string // 标题（必填，最大255字符）

	// 可选字段
	Description       string                 // 描述
	StartAt           *int64                 // 开始时间（可选）
	EndAt             *int64                 // 结束时间
	PriorityID        string                 // 优先级 ID
	StateID           string                 // 状态 ID
	AssigneeID        string                 // 负责人 ID
	ParentID          string                 // 父工作项 ID
	SprintID          string                 // Sprint ID
	VersionID         string                 // 版本 ID
	BoardID           string                 // 看板 ID
	EntryID           string                 // 看板栏 ID
	SwimlaneID        string                 // 泳道 ID
	StoryPoints       float64                // 故事点
	EstimatedWorkload float64                // 预估工时
	RemainingWorkload float64                // 剩余工时
	Properties        map[string]interface{} // 自定义属性
	ParticipantIDs    []string               // 参与人 ID 列表
}

// WorkItemUpdateInput 部分更新工作项输入（所有字段可选，使用指针）
type WorkItemUpdateInput struct {
	Title             *string                // 标题
	Description       *string                // 描述
	StartAt           *int64                 // 开始时间
	EndAt             *int64                 // 结束时间
	PriorityID        *string                // 优先级 ID
	StateID           *string                // 状态 ID
	AssigneeID        *string                // 负责人 ID
	ParentID          *string                // 父工作项 ID
	SprintID          *string                // Sprint ID
	VersionID         *string                // 版本 ID
	BoardID           *string                // 看板 ID
	EntryID           *string                // 看板栏 ID
	SwimlaneID        *string                // 泳道 ID
	StoryPoints       *float64               // 故事点
	EstimatedWorkload *float64               // 预估工时
	RemainingWorkload *float64               // 剩余工时
	Properties        map[string]interface{} // 自定义属性
	ParticipantIDs    []string               // 参与人 ID 列表
}

// WorkItemBatchUpdateInput 批量部分更新工作项输入
type WorkItemBatchUpdateInput struct {
	WorkItemIDs []string              // 要更新的工作项 ID 列表
	Updates     WorkItemUpdateInput   // 要更新的字段
}

// WorkItemBatchUpdateResult 批量更新结果
type WorkItemBatchUpdateResult struct {
	SuccessIDs []string // 成功更新的 ID
	FailedIDs  []string // 失败的 ID
}
