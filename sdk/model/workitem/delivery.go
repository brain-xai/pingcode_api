package workitem

// WorkItemDeliveryTarget 工作项交付目标
type WorkItemDeliveryTarget struct {
	ID          string                 // 交付目标 ID
	URL         string                 // 交付目标资源地址
	Name        string                 // 交付目标名称
	Description string                 // 交付目标描述
	ContentType string                 // 交付物类型：link, attachment
	Content     map[string]string      // 交付物内容（如 {"name": "xxx", "href": "xxx"}）
	WorkItemID  string                 // 工作项 ID
	WorkItem    *WorkItemSummary       // 工作项
	Project     *ProjectSummary        // 项目
	CreatedAt   int64                  // 创建时间
	UpdatedAt   int64                  // 更新时间
	CreatedBy   *User                  // 创建人
	UpdatedBy   *User                  // 更新人
	IsArchived  bool                   // 是否已归档
	IsDeleted   bool                   // 是否已删除
}

// WorkItemDeliveryTargetCreateInput 创建交付目标输入
type WorkItemDeliveryTargetCreateInput struct {
	WorkItemID  string                 // 工作项 ID（必填）
	Name        string                 // 交付目标名称（必填）
	Description string                 // 交付目标描述
	ContentType string                 // 交付物类型：link
	Content     map[string]interface{} // 交付物内容
}

// WorkItemDeliveryTargetUpdateInput 更新交付目标输入
type WorkItemDeliveryTargetUpdateInput struct {
	Name        *string                // 交付目标名称
	Description *string                // 交付目标描述
	ContentType *string                // 交付物类型
	Content     map[string]interface{} // 交付物内容
}

// WorkItemDeliveryTargetList 交付目标列表
type WorkItemDeliveryTargetList struct {
	PageSize  int                       // 页面大小
	PageIndex int                       // 页码
	Total     int                       // 总数
	Values    []WorkItemDeliveryTarget  // 交付目标列表
}
