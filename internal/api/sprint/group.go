package sprint

import (
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// CreateSprintGroupRequest 创建迭代分组请求（API 中称为 SprintSection）
type CreateSprintGroupRequest struct {
	Name string `json:"name"`
}

// CreateSprintGroupResponse 创建迭代分组响应
type CreateSprintGroupResponse struct {
	Group SprintGroup `json:"sprint_section"` // API 字段名为 sprint_section
}

// UpdateSprintGroupRequest 更新迭代分组请求
type UpdateSprintGroupRequest struct {
	Name string `json:"name"`
}

// UpdateSprintGroupResponse 更新迭代分组响应
type UpdateSprintGroupResponse struct {
	Group SprintGroup `json:"sprint_section"` // API 字段名为 sprint_section
}

// ListSprintGroupsResponse 迭代分组列表响应
type ListSprintGroupsResponse struct {
	PageSize  int          `json:"page_size"`
	PageIndex int          `json:"page_index"`
	Total     int          `json:"total"`
	Values    []SprintGroup `json:"values"`
}

// DeleteSprintGroupResponse 删除迭代分组响应
type DeleteSprintGroupResponse struct {
	Group SprintGroup `json:"sprint_section"` // API 字段名为 sprint_section
}

// SprintGroup 迭代分组 DTO（API 中称为 SprintSection）
type SprintGroup struct {
	ID      string         `json:"id"`
	URL     string         `json:"url"`
	Name    string         `json:"name"`
	Project ProjectSummary `json:"project"`
}

// ToModel 转换为模型
func (g *SprintGroup) ToModel() *sprintmodel.SprintGroup {
	model := &sprintmodel.SprintGroup{
		ID:   g.ID,
		URL:  g.URL,
		Name: g.Name,
	}

	// 转换 Project
	model.Project = g.Project.ToModel()

	return model
}
