package workitem

// WorkItemTransition 工作项流转记录
type WorkItemTransition struct {
	ID         string              // 流转记录 ID
	URL        string              // 流转记录资源地址
	WorkItemID string              // 工作项 ID
	FromState  *WorkItemStatus     // 从状态
	ToState    *WorkItemStatus     // 到状态
	CreatedAt  int64               // 创建时间
	CreatedBy  *User               // 创建人
	Comment    string              // 备注
}

// WorkItemTransitionList 流转记录列表
type WorkItemTransitionList struct {
	PageSize  int                    // 页面大小
	PageIndex int                    // 页码
	Total     int                    // 总数
	Values    []WorkItemTransition   // 流转记录列表
}
