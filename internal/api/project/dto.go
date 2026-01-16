package api

import projectmodel "github.com/brain-xai/pingcode_api/sdk/model/project"

// CreateProjectRequest 创建项目请求 DTO
type CreateProjectRequest struct {
	Name        string                    `json:"name"`
	Type        string                    `json:"type"`
	Identifier  string                    `json:"identifier"`
	ScopeType   string                    `json:"scope_type,omitempty"`
	ScopeID     string                    `json:"scope_id,omitempty"`
	Visibility  string                    `json:"visibility,omitempty"`
	Description string                    `json:"description,omitempty"`
	StartAt     *int64                    `json:"start_at,omitempty"`
	EndAt       *int64                    `json:"end_at,omitempty"`
	AssigneeID  string                    `json:"assignee_id,omitempty"`
	Members     []ProjectMemberInputDTO   `json:"members,omitempty"`
}

// ProjectMemberInputDTO 项目成员输入 DTO
type ProjectMemberInputDTO struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// UpdateProjectRequest 更新项目请求 DTO
type UpdateProjectRequest struct {
	Name        *string                `json:"name,omitempty"`
	Identifier  *string                `json:"identifier,omitempty"`
	Description *string                `json:"description,omitempty"`
	StartAt     *int64                 `json:"start_at,omitempty"`
	EndAt       *int64                 `json:"end_at,omitempty"`
	AssigneeID  *string                `json:"assignee_id,omitempty"`
	StateID     *string                `json:"state_id,omitempty"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
}

// ToCreateRequestDTO 将 ProjectCreateInput 转换为 CreateProjectRequest
func ToCreateRequestDTO(input projectmodel.ProjectCreateInput) *CreateProjectRequest {
	dto := &CreateProjectRequest{
		Name:        input.Name,
		Type:        input.Type,
		Identifier:  input.Identifier,
		ScopeType:   input.ScopeType,
		ScopeID:     input.ScopeID,
		Visibility:  input.Visibility,
		Description: input.Description,
		AssigneeID:  input.AssigneeID,
		StartAt:     input.StartAt,
		EndAt:       input.EndAt,
	}

	if len(input.Members) > 0 {
		dto.Members = make([]ProjectMemberInputDTO, len(input.Members))
		for i, m := range input.Members {
			dto.Members[i] = ProjectMemberInputDTO{ID: m.ID, Type: m.Type}
		}
	}

	return dto
}

// ToUpdateRequestDTO 将 ProjectUpdateInput 转换为 UpdateProjectRequest
func ToUpdateRequestDTO(input projectmodel.ProjectUpdateInput) *UpdateProjectRequest {
	// 转换 Properties: map[string]string -> map[string]interface{}
	var props map[string]interface{}
	if input.Properties != nil {
		props = make(map[string]interface{}, len(input.Properties))
		for k, v := range input.Properties {
			props[k] = v
		}
	}

	return &UpdateProjectRequest{
		Name:        input.Name,
		Identifier:  input.Identifier,
		Description: input.Description,
		StartAt:     input.StartAt,
		EndAt:       input.EndAt,
		AssigneeID:  input.AssigneeID,
		StateID:     input.StateID,
		Properties:  props,
	}
}
