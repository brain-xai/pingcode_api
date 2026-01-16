package workitem

import (
	"github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// ============= 工作项流转记录 =============

// ListWorkItemTransitionsResponse 流转记录列表响应
type ListWorkItemTransitionsResponse struct {
	PageSize  int                     `json:"page_size"`
	PageIndex int                     `json:"page_index"`
	Total     int                     `json:"total"`
	Values    []WorkItemTransition    `json:"values"`
}

// WorkItemTransition 流转记录 DTO
type WorkItemTransition struct {
	ID         string             `json:"id"`
	URL        string             `json:"url"`
	WorkItem   *WorkItemSummary   `json:"work_item"`
	FromState  *WorkItemStatus    `json:"from_state"`
	ToState    *WorkItemStatus    `json:"to_state"`
	CreatedAt  int64              `json:"created_at"`
	CreatedBy  *User              `json:"created_by"`
}

// ToModel 将 WorkItemTransition DTO 转换为 Model
func (w *WorkItemTransition) ToModel() *workitem.WorkItemTransition {
	model := &workitem.WorkItemTransition{
		ID:        w.ID,
		URL:       w.URL,
		CreatedAt: w.CreatedAt,
	}

	if w.WorkItem != nil {
		model.WorkItemID = w.WorkItem.ID
	}

	if w.FromState != nil {
		model.FromState = w.FromState.ToModel()
	}

	if w.ToState != nil {
		model.ToState = w.ToState.ToModel()
	}

	if w.CreatedBy != nil {
		model.CreatedBy = w.CreatedBy.ToModel()
	}

	return model
}

// ============= 工作项交付目标 =============

// ListWorkItemDeliverablesResponse 交付目标列表响应
type ListWorkItemDeliverablesResponse struct {
	PageSize  int                        `json:"page_size"`
	PageIndex int                        `json:"page_index"`
	Total     int                        `json:"total"`
	Values    []WorkItemDeliverable      `json:"values"`
}

// WorkItemDeliverable 交付目标 DTO
type WorkItemDeliverable struct {
	ID          string                 `json:"id"`
	URL         string                 `json:"url"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	ContentType string                 `json:"content_type"`
	Content     map[string]interface{} `json:"content"`
	Project     *ProjectSummary        `json:"project"`
	WorkItem    *WorkItemSummary       `json:"work_item"`
	CreatedAt   int64                  `json:"created_at"`
	UpdatedAt   int64                  `json:"updated_at"`
	CreatedBy   *User                  `json:"created_by"`
	UpdatedBy   *User                  `json:"updated_by"`
	IsArchived  int                    `json:"is_archived"`
	IsDeleted   int                    `json:"is_deleted"`
}

// CreateWorkItemDeliverableRequest 创建交付目标请求 DTO
type CreateWorkItemDeliverableRequest struct {
	WorkItemID  string                 `json:"work_item_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Content     map[string]interface{} `json:"content,omitempty"`
}

// UpdateWorkItemDeliverableRequest 更新交付目标请求 DTO
type UpdateWorkItemDeliverableRequest struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	ContentType *string                `json:"content_type,omitempty"`
	Content     map[string]interface{} `json:"content,omitempty"`
}

// ToModel 将 WorkItemDeliverable DTO 转换为 Model
func (w *WorkItemDeliverable) ToModel() *workitem.WorkItemDeliveryTarget {
	model := &workitem.WorkItemDeliveryTarget{
		ID:          w.ID,
		URL:         w.URL,
		Name:        w.Name,
		Description: w.Description,
		ContentType: w.ContentType,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
		IsArchived:  w.IsArchived == 1,
		IsDeleted:   w.IsDeleted == 1,
	}

	if w.Content != nil {
		model.Content = make(map[string]string)
		for k, v := range w.Content {
			model.Content[k] = toString(v)
		}
	}

	if w.Project != nil {
		model.Project = w.Project.ToModel()
	}

	if w.WorkItem != nil {
		model.WorkItemID = w.WorkItem.ID
		model.WorkItem = w.WorkItem.ToModel()
	}

	if w.CreatedBy != nil {
		model.CreatedBy = w.CreatedBy.ToModel()
	}

	if w.UpdatedBy != nil {
		model.UpdatedBy = w.UpdatedBy.ToModel()
	}

	return model
}

// ToCreateRequestDTO 将 CreateInput 转换为请求 DTO
func ToCreateDeliverableRequestDTO(input workitem.WorkItemDeliveryTargetCreateInput) CreateWorkItemDeliverableRequest {
	return CreateWorkItemDeliverableRequest{
		WorkItemID:  input.WorkItemID,
		Name:        input.Name,
		Description: input.Description,
		ContentType: input.ContentType,
		Content:     input.Content,
	}
}

// ToUpdateRequestDTO 将 UpdateInput 转换为请求 DTO
func ToUpdateDeliverableRequestDTO(input workitem.WorkItemDeliveryTargetUpdateInput) UpdateWorkItemDeliverableRequest {
	return UpdateWorkItemDeliverableRequest{
		Name:        input.Name,
		Description: input.Description,
		ContentType: input.ContentType,
		Content:     input.Content,
	}
}
