package sprint

import (
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// CreateSprintRequest 创建迭代请求
type CreateSprintRequest struct {
	Name        string   `json:"name"`
	StartAt     int64    `json:"start_at"`
	EndAt       int64    `json:"end_at"`
	AssigneeID  string   `json:"assignee_id"`
	Description string   `json:"description,omitempty"`
	Status      string   `json:"status,omitempty"`
	CategoryIDs []string `json:"category_ids,omitempty"`
}

// CreateSprintResponse 创建迭代响应
type CreateSprintResponse struct {
	Sprint Sprint `json:"sprint"`
}

// BatchCreateSprintsRequest 批量创建迭代请求
type BatchCreateSprintsRequest struct {
	Sprints []CreateSprintRequest `json:"sprints"`
}

// BatchCreateSprintsResponse 批量创建迭代响应项
type BatchCreateSprintsResponse struct {
	State  string  `json:"state"`
	Sprint Sprint  `json:"sprint"`
}

// UpdateSprintRequest 更新迭代请求
type UpdateSprintRequest struct {
	Name        *string  `json:"name,omitempty"`
	StartAt     *int64   `json:"start_at,omitempty"`
	EndAt       *int64   `json:"end_at,omitempty"`
	AssigneeID  *string  `json:"assignee_id,omitempty"`
	Description *string  `json:"description,omitempty"`
	Status      *string  `json:"status,omitempty"`
	CategoryIDs []string `json:"category_ids,omitempty"`
}

// UpdateSprintResponse 更新迭代响应
type UpdateSprintResponse struct {
	Sprint Sprint `json:"sprint"`
}

// ListSprintsResponse 迭代列表响应
type ListSprintsResponse struct {
	PageSize  int      `json:"page_size"`
	PageIndex int      `json:"page_index"`
	Total     int      `json:"total"`
	Values    []Sprint `json:"values"`
}

// Sprint 迭代 DTO
type Sprint struct {
	ID                   string                 `json:"id"`
	URL                  string                 `json:"url"`
	Name                 string                 `json:"name"`
	Status               string                 `json:"status"`
	StartAt              int64                  `json:"start_at"`
	EndAt                int64                  `json:"end_at"`
	Description          string                 `json:"description"`
	StartedAt            int64                  `json:"started_at"`
	CompletedAt          int64                  `json:"completed_at"`
	TotalStoryPoints     float64                `json:"total_story_points"`
	StartedStoryPoints   float64                `json:"started_story_points"`
	CompletedStoryPoints float64                `json:"completed_story_points"`
	CreatedAt            int64                  `json:"created_at"`
	UpdatedAt            int64                  `json:"updated_at"`
	Project              ProjectSummary         `json:"project"`
	Assignee             *User                  `json:"assignee"`
	Categories           []SprintCategoryRef    `json:"categories"`
	CreatedBy            *User                  `json:"created_by"`
	UpdatedBy            *User                  `json:"updated_by"`
}

// ProjectSummary 项目摘要 DTO
type ProjectSummary struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	IsArchived int    `json:"is_archived"`
	IsDeleted  int    `json:"is_deleted"`
}

// User 用户 DTO
type User struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

// SprintCategoryRef 迭代类别摘要 DTO
type SprintCategoryRef struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}

// ToModel 转换为模型
func (s *Sprint) ToModel() *sprintmodel.Sprint {
	model := &sprintmodel.Sprint{
		ID:                   s.ID,
		URL:                  s.URL,
		Name:                 s.Name,
		Status:               s.Status,
		StartAt:              s.StartAt,
		EndAt:                s.EndAt,
		Description:          s.Description,
		StartedAt:            s.StartedAt,
		CompletedAt:          s.CompletedAt,
		TotalStoryPoints:     s.TotalStoryPoints,
		StartedStoryPoints:   s.StartedStoryPoints,
		CompletedStoryPoints: s.CompletedStoryPoints,
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
	}

	// 转换 Project
	model.Project = s.Project.ToModel()

	// 转换 Assignee
	if s.Assignee != nil {
		model.Assignee = s.Assignee.ToModel()
	}

	// 转换 Categories
	if len(s.Categories) > 0 {
		model.Categories = make([]sprintmodel.SprintCategory, len(s.Categories))
		for i, cat := range s.Categories {
			model.Categories[i] = *cat.ToRefModel()
		}
	}

	// 转换 CreatedBy
	if s.CreatedBy != nil {
		model.CreatedBy = s.CreatedBy.ToModel()
	}

	// 转换 UpdatedBy
	if s.UpdatedBy != nil {
		model.UpdatedBy = s.UpdatedBy.ToModel()
	}

	return model
}

// ToModel 转换为模型
func (p *ProjectSummary) ToModel() *sprintmodel.ProjectSummary {
	return &sprintmodel.ProjectSummary{
		ID:         p.ID,
		URL:        p.URL,
		Identifier: p.Identifier,
		Name:       p.Name,
		Type:       p.Type,
		IsArchived: p.IsArchived,
		IsDeleted:  p.IsDeleted,
	}
}

// ToModel 转换为模型
func (u *User) ToModel() *sprintmodel.User {
	return &sprintmodel.User{
		ID:          u.ID,
		URL:         u.URL,
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Avatar:      u.Avatar,
	}
}

// ToRefModel 转换为类别摘要模型
func (c *SprintCategoryRef) ToRefModel() *sprintmodel.SprintCategory {
	return &sprintmodel.SprintCategory{
		ID:   c.ID,
		URL:  c.URL,
		Name: c.Name,
	}
}
