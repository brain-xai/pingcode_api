package workitem

import (
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// ============= 工作项类型 =============

// ListWorkItemTypesResponse 工作项类型列表响应
type ListWorkItemTypesResponse struct {
	PageSize  int              `json:"page_size"`
	PageIndex int              `json:"page_index"`
	Total     int              `json:"total"`
	Values    []WorkItemType   `json:"values"`
}

// WorkItemType 工作项类型 DTO
type WorkItemType struct {
	ID          string                 `json:"id"`
	URL         string                 `json:"url"`
	Name        string                 `json:"name"`
	Icon        string                 `json:"icon"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
}

// ToModel 将 WorkItemType DTO 转换为 Model
func (w *WorkItemType) ToModel() *workitemmodel.WorkItemType {
	model := &workitemmodel.WorkItemType{
		ID:          w.ID,
		URL:         w.URL,
		Name:        w.Name,
		Icon:        w.Icon,
		Description: w.Description,
	}

	if w.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range w.Properties {
			model.Properties[k] = toString(v)
		}
	}

	return model
}

// ============= 工作项状态 =============

// ListWorkItemStatesResponse 工作项状态列表响应
type ListWorkItemStatesResponse struct {
	PageSize  int               `json:"page_size"`
	PageIndex int               `json:"page_index"`
	Total     int               `json:"total"`
	Values    []WorkItemStatus  `json:"values"`
}

// WorkItemStatus 工作项状态 DTO
type WorkItemStatus struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Color    string `json:"color"`
	Category string `json:"category"`
}

// ToModel 将 WorkItemStatus DTO 转换为 Model
func (w *WorkItemStatus) ToModel() *workitemmodel.WorkItemStatus {
	return &workitemmodel.WorkItemStatus{
		ID:       w.ID,
		URL:      w.URL,
		Name:     w.Name,
		Type:     w.Type,
		Color:    w.Color,
		Category: w.Category,
	}
}

// ============= 工作项属性 =============

// ListWorkItemFieldsResponse 工作项属性列表响应
type ListWorkItemFieldsResponse struct {
	PageSize  int              `json:"page_size"`
	PageIndex int              `json:"page_index"`
	Total     int              `json:"total"`
	Values    []WorkItemField  `json:"values"`
}

// WorkItemField 工作项属性 DTO
type WorkItemField struct {
	ID          string                 `json:"id"`
	URL         string                 `json:"url"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Options     []WorkItemFieldOption  `json:"options"`
	Properties  map[string]interface{} `json:"properties"`
}

// WorkItemFieldOption 属性选项 DTO
type WorkItemFieldOption struct {
	ID   string `json:"_id"`
	Text string `json:"text"`
}

// ToModel 将 WorkItemField DTO 转换为 Model
func (w *WorkItemField) ToModel() *workitemmodel.WorkItemField {
	model := &workitemmodel.WorkItemField{
		ID:          w.ID,
		URL:         w.URL,
		Name:        w.Name,
		Type:        w.Type,
		Description: w.Description,
	}

	if len(w.Options) > 0 {
		model.Options = make([]workitemmodel.WorkItemFieldOption, len(w.Options))
		for i, opt := range w.Options {
			model.Options[i] = *opt.ToModel()
		}
	}

	if w.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range w.Properties {
			model.Properties[k] = toString(v)
		}
	}

	return model
}

// ToModel 将 WorkItemFieldOption DTO 转换为 Model
func (w *WorkItemFieldOption) ToModel() *workitemmodel.WorkItemFieldOption {
	return &workitemmodel.WorkItemFieldOption{
		ID:   w.ID,
		Text: w.Text,
	}
}

// ============= 工作项优先级 =============

// ListWorkItemPrioritiesResponse 工作项优先级列表响应
type ListWorkItemPrioritiesResponse struct {
	PageSize  int                    `json:"page_size"`
	PageIndex int                    `json:"page_index"`
	Total     int                    `json:"total"`
	Values    []WorkItemPriority     `json:"values"`
}

// WorkItemPriority 工作项优先级 DTO
type WorkItemPriority struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Icon   string `json:"icon"`
}

// ToModel 将 WorkItemPriority DTO 转换为 Model
func (w *WorkItemPriority) ToModel() *workitemmodel.WorkItemPriority {
	return &workitemmodel.WorkItemPriority{
		ID:    w.ID,
		Name:  w.Name,
		Color: w.Color,
		Icon:  w.Icon,
	}
}

// ============= 工作项标签 =============

// ListWorkItemTagsResponse 工作项标签列表响应
type ListWorkItemTagsResponse struct {
	PageSize  int              `json:"page_size"`
	PageIndex int              `json:"page_index"`
	Total     int              `json:"total"`
	Values    []WorkItemTag    `json:"values"`
}

// WorkItemTag 工作项标签 DTO
type WorkItemTag struct {
	ID          string                 `json:"id"`
	URL         string                 `json:"url"`
	Name        string                 `json:"name"`
	Color       string                 `json:"color"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
}

// ToModel 将 WorkItemTag DTO 转换为 Model
func (w *WorkItemTag) ToModel() *workitemmodel.WorkItemTag {
	model := &workitemmodel.WorkItemTag{
		ID:          w.ID,
		URL:         w.URL,
		Name:        w.Name,
		Color:       w.Color,
		Description: w.Description,
	}

	if w.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range w.Properties {
			model.Properties[k] = toString(v)
		}
	}

	return model
}
