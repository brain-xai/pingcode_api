package workitem

import (
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// ============= 工作项关联类型 =============

// ListWorkItemRelationTypesResponse 关联类型列表响应
type ListWorkItemRelationTypesResponse struct {
	PageSize  int                     `json:"page_size"`
	PageIndex int                     `json:"page_index"`
	Total     int                     `json:"total"`
	Values    []WorkItemRelationType  `json:"values"`
}

// WorkItemRelationType 关联类型 DTO
type WorkItemRelationType struct {
	ID          string                 `json:"id"`
	URL         string                 `json:"url"`
	Name        string                 `json:"name"`
	Category    string                 `json:"category"`
	IsSystem    int                    `json:"is_system"`
	Properties  map[string]interface{} `json:"properties"`
}

// ToModel 将 WorkItemRelationType DTO 转换为 Model
func (w *WorkItemRelationType) ToModel() *workitemmodel.WorkItemLinkType {
	model := &workitemmodel.WorkItemLinkType{
		ID:       w.ID,
		Name:     w.Name,
		Category: w.Category,
		IsSystem: w.IsSystem == 1,
	}

	if w.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range w.Properties {
			model.Properties[k] = toString(v)
		}
	}

	return model
}

// ============= 工作项关联 =============

// ListWorkItemRelationsResponse 关联列表响应
type ListWorkItemRelationsResponse struct {
	PageSize  int                `json:"page_size"`
	PageIndex int                `json:"page_index"`
	Total     int                `json:"total"`
	Values    []WorkItemRelation `json:"values"`
}

// WorkItemRelation 关联 DTO
type WorkItemRelation struct {
	ID             string             `json:"id"`
	URL            string             `json:"url"`
	RelationType   string             `json:"relation_type"`
	OriginWorkItem *WorkItemSummary   `json:"origin_work_item"`
	TargetWorkItem *WorkItemSummary   `json:"target_work_item"`
}

// ToModel 将 WorkItemRelation DTO 转换为 Model
func (w *WorkItemRelation) ToModel() *workitemmodel.WorkItemLink {
	model := &workitemmodel.WorkItemLink{
		ID:           w.ID,
		URL:          w.URL,
		RelationType: w.RelationType,
	}

	if w.OriginWorkItem != nil {
		model.Source = w.OriginWorkItem.ToModel()
	}

	if w.TargetWorkItem != nil {
		model.Target = w.TargetWorkItem.ToModel()
	}

	return model
}

// CreateWorkItemRelationRequest 创建关联请求 DTO
type CreateWorkItemRelationRequest struct {
	TargetWorkItemID string `json:"target_work_item_id"`
	RelationType     string `json:"relation_type"`
}
