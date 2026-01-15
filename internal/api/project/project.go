package api

import (
	"fmt"

	projectmodel "github.com/brain-xai/pingcode_api/sdk/model/project"
)

// ListProjectsResponse 项目列表响应
type ListProjectsResponse struct {
	PageSize  int       `json:"page_size"`
	PageIndex int       `json:"page_index"`
	Total     int       `json:"total"`
	Values    []Project `json:"values"`
}

// Project API 项目 DTO（完整结构）
type Project struct {
	ID          string                `json:"id"`
	URL         string                `json:"url"`
	Identifier  string                `json:"identifier"`
	Name        string                `json:"name"`
	Type        string                `json:"type"`
	Description string                `json:"description"`
	State       State                 `json:"state"`
	StartAt     int64                 `json:"start_at"`
	EndAt       int64                 `json:"end_at"`
	CreatedAt   int64                 `json:"created_at"`
	UpdatedAt   int64                 `json:"updated_at"`
	// 扩展字段（与 OpenAPI 一致）
	Assignee    *User                 `json:"assignee"`
	ScopeType   string                `json:"scope_type"`
	ScopeID     string                `json:"scope_id"`
	Visibility  string                `json:"visibility"`
	Color       string                `json:"color"`
	Properties  map[string]interface{} `json:"properties"`
	Members     []ProjectMember       `json:"members"`
	CreatedBy   *User                 `json:"created_by"`
	UpdatedBy   *User                 `json:"updated_by"`
	IsArchived  int                   `json:"is_archived"` // API 返回 0/1
	IsDeleted   int                   `json:"is_deleted"`  // API 返回 0/1
}

// State 项目状态
type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// User 用户 DTO
type User struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

// UserGroup 团队 DTO
type UserGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProjectMember 项目成员 DTO
type ProjectMember struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`      // user, user_group
	User      *User      `json:"user"`      // 当 type=user 时
	UserGroup *UserGroup `json:"user_group"` // 当 type=user_group 时
}

// ListProjectStatesResponse 项目状态列表响应
type ListProjectStatesResponse struct {
	PageSize  int     `json:"page_size"`
	PageIndex int     `json:"page_index"`
	Total     int     `json:"total"`
	Values    []State `json:"values"`
}

// ListProjectMembersResponse 项目成员列表响应
type ListProjectMembersResponse struct {
	PageSize  int             `json:"page_size"`
	PageIndex int             `json:"page_index"`
	Total     int             `json:"total"`
	Values    []ProjectMember `json:"values"`
}

// GetProjectProgressResponse 项目进度响应
type GetProjectProgressResponse struct {
	CompletedCount int     `json:"completed_count"`
	TotalCount     int     `json:"total_count"`
	CompletionRate float64 `json:"completion_rate"`
}

// ToModel 将 API DTO 转换为对外模型
func (p *Project) ToModel() *projectmodel.Project {
	model := &projectmodel.Project{
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
		ScopeType:   p.ScopeType,
		ScopeID:     p.ScopeID,
		Visibility:  p.Visibility,
		Color:       p.Color,
		IsArchived:  p.IsArchived == 1,
		IsDeleted:   p.IsDeleted == 1,
	}

	// 转换 Assignee
	if p.Assignee != nil {
		model.Assignee = &projectmodel.User{
			ID:          p.Assignee.ID,
			Name:        p.Assignee.Name,
			DisplayName: p.Assignee.DisplayName,
			Avatar:      p.Assignee.Avatar,
		}
	}

	// 转换 Properties
	if p.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range p.Properties {
			model.Properties[k] = toString(v)
		}
	}

	// 转换 Members
	if len(p.Members) > 0 {
		model.Members = make([]projectmodel.ProjectMember, len(p.Members))
		for i, m := range p.Members {
			model.Members[i] = *m.ToModel()
		}
	}

	// 转换 CreatedBy
	if p.CreatedBy != nil {
		model.CreatedBy = &projectmodel.User{
			ID:          p.CreatedBy.ID,
			Name:        p.CreatedBy.Name,
			DisplayName: p.CreatedBy.DisplayName,
			Avatar:      p.CreatedBy.Avatar,
		}
	}

	// 转换 UpdatedBy
	if p.UpdatedBy != nil {
		model.UpdatedBy = &projectmodel.User{
			ID:          p.UpdatedBy.ID,
			Name:        p.UpdatedBy.Name,
			DisplayName: p.UpdatedBy.DisplayName,
			Avatar:      p.UpdatedBy.Avatar,
		}
	}

	return model
}

// ToModel 将 ProjectMember DTO 转换为对外模型
func (m *ProjectMember) ToModel() *projectmodel.ProjectMember {
	model := &projectmodel.ProjectMember{
		ID:   m.ID,
		Type: m.Type,
	}

	if m.User != nil {
		model.User = &projectmodel.User{
			ID:          m.User.ID,
			Name:        m.User.Name,
			DisplayName: m.User.DisplayName,
			Avatar:      m.User.Avatar,
		}
	}

	if m.UserGroup != nil {
		model.UserGroup = &projectmodel.UserGroup{
			ID:   m.UserGroup.ID,
			Name: m.UserGroup.Name,
		}
	}

	return model
}

// ToModel 将 State DTO 转换为对外模型
func (s *State) ToModel() *projectmodel.State {
	return &projectmodel.State{
		ID:   s.ID,
		Name: s.Name,
		Type: s.Type,
	}
}

// toString 辅助函数：将任意类型转换为字符串
func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%v", val)
	case bool:
		return fmt.Sprintf("%v", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}
