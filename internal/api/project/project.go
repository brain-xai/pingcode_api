package api

import "github.com/brain-xai/pingcode_api/sdk/model/project"

// ListProjectsResponse 项目列表响应
type ListProjectsResponse struct {
	PageSize  int       `json:"page_size"`
	PageIndex int       `json:"page_index"`
	Total     int       `json:"total"`
	Values    []Project `json:"values"`
}

// Project API 项目 DTO（完整结构）
type Project struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	State       State  `json:"state"`
	StartAt     int64  `json:"start_at"`
	EndAt       int64  `json:"end_at"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// State 项目状态
type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ToModel 将 API DTO 转换为对外模型
func (p *Project) ToModel() *project.Project {
	return &project.Project{
		ID:          p.ID,
		Identifier:  p.Identifier,
		Name:        p.Name,
		Type:        p.Type,
		Description: p.Description,
		State:       p.State.Name,
		StartAt:     p.StartAt,
		EndAt:       p.EndAt,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
