package sprint

import (
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// CreateSprintCategoryRequest 创建迭代类别请求
type CreateSprintCategoryRequest struct {
	Name      string `json:"name"`
	GroupID   string `json:"section_id,omitempty"` // API 字段名为 section_id
}

// CreateSprintCategoryResponse 创建迭代类别响应
type CreateSprintCategoryResponse struct {
	Category SprintCategory `json:"sprint_category"` // API 字段名为 sprint_category
}

// UpdateSprintCategoryRequest 更新迭代类别请求
type UpdateSprintCategoryRequest struct {
	Name    string `json:"name,omitempty"`
	GroupID string `json:"section_id,omitempty"` // API 字段名为 section_id
}

// UpdateSprintCategoryResponse 更新迭代类别响应
type UpdateSprintCategoryResponse struct {
	Category SprintCategory `json:"sprint_category"` // API 字段名为 sprint_category
}

// ListSprintCategoriesResponse 迭代类别列表响应
type ListSprintCategoriesResponse struct {
	PageSize  int             `json:"page_size"`
	PageIndex int             `json:"page_index"`
	Total     int             `json:"total"`
	Values    []SprintCategory `json:"values"`
}

// DeleteSprintCategoryResponse 删除迭代类别响应
type DeleteSprintCategoryResponse struct {
	Category SprintCategory `json:"sprint_category"` // API 字段名为 sprint_category
}

// SprintCategory 迭代类别 DTO
type SprintCategory struct {
	ID      string              `json:"id"`
	URL     string              `json:"url"`
	Name    string              `json:"name"`
	Project ProjectSummary      `json:"project"`
	Group   *SprintGroupSummary `json:"section"` // API 字段名为 section
}

// SprintGroupSummary 迭代分组摘要 DTO
type SprintGroupSummary struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}

// ToModel 转换为模型
func (c *SprintCategory) ToModel() *sprintmodel.SprintCategoryFull {
	model := &sprintmodel.SprintCategoryFull{
		ID:      c.ID,
		URL:     c.URL,
		Name:    c.Name,
		GroupID: "",
	}

	// 转换 Project
	model.Project = c.Project.ToModel()

	// 转换 Group
	if c.Group != nil {
		model.Group = c.Group.ToModel()
		model.GroupID = c.Group.ID
	}

	return model
}

// ToModel 转换为模型
func (g *SprintGroupSummary) ToModel() *sprintmodel.SprintGroup {
	return &sprintmodel.SprintGroup{
		ID:   g.ID,
		URL:  g.URL,
		Name: g.Name,
	}
}
