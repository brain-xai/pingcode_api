package project

import (
	"context"
	"fmt"
	"net/url"

	apiproject "github.com/brain-xai/pingcode_api/internal/api/project"
	"github.com/brain-xai/pingcode_api/internal/httpclient"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	projectmodel "github.com/brain-xai/pingcode_api/sdk/model/project"
)

// Service 项目服务
type Service struct {
	client  *httpclient.Client
	baseURL string
}

// NewService 创建项目服务
func NewService(client *httpclient.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}

// List 获取项目列表
func (s *Service) List(ctx context.Context, filter projectmodel.ProjectFilter) ([]projectmodel.Project, *core.Pagination, error) {
	// 构建查询参数
	query := url.Values{}

	if filter.Identifier != "" {
		query.Set("identifier", filter.Identifier)
	}
	if filter.Type != "" {
		query.Set("type", filter.Type)
	}
	if filter.IncludeDeleted {
		query.Set("include_deleted", "true")
	}
	if filter.IncludeArchived {
		query.Set("include_archived", "true")
	}

	// 发送请求
	var resp apiproject.ListProjectsResponse
	if err := s.client.Get(ctx, "/v1/project/projects", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list projects: %w", err)
	}

	// 转换 DTO 为 Model
	projects := make([]projectmodel.Project, len(resp.Values))
	for i, p := range resp.Values {
		projects[i] = *p.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return projects, pagination, nil
}

// Get 按 ID 获取项目详情
func (s *Service) Get(ctx context.Context, projectID string) (*projectmodel.Project, error) {
	// 参数校验
	if projectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{"project_id": projectID}

	// 发送请求
	var resp apiproject.Project
	if err := s.client.GetWithPathParams(ctx, "/v1/project/projects/{project_id}", pathParams, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	return resp.ToModel(), nil
}

// ListStates 获取项目状态列表
func (s *Service) ListStates(ctx context.Context, projectID string) ([]projectmodel.State, *core.Pagination, error) {
	// 参数校验
	if projectID == "" {
		return nil, nil, fmt.Errorf("project_id cannot be empty")
	}

	// 查询参数
	query := url.Values{}
	query.Set("project_id", projectID)

	// 发送请求
	var resp apiproject.ListProjectStatesResponse
	if err := s.client.Get(ctx, "/v1/project/project/states", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list project states: %w", err)
	}

	// 转换 DTO 为 Model
	states := make([]projectmodel.State, len(resp.Values))
	for i, s := range resp.Values {
		states[i] = *s.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return states, pagination, nil
}

// ListMembers 获取项目成员列表
func (s *Service) ListMembers(ctx context.Context, projectID string) ([]projectmodel.ProjectMember, *core.Pagination, error) {
	// 参数校验
	if projectID == "" {
		return nil, nil, fmt.Errorf("project_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{"project_id": projectID}

	// 发送请求
	var resp apiproject.ListProjectMembersResponse
	if err := s.client.GetWithPathParams(ctx, "/v1/project/projects/{project_id}/members", pathParams, nil, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list project members: %w", err)
	}

	// 转换 DTO 为 Model
	members := make([]projectmodel.ProjectMember, len(resp.Values))
	for i, m := range resp.Values {
		members[i] = *m.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return members, pagination, nil
}

// GetProgress 获取项目进度
func (s *Service) GetProgress(ctx context.Context, projectID string) (*projectmodel.Progress, error) {
	// 参数校验
	if projectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{"project_id": projectID}

	// 发送请求
	var resp apiproject.GetProjectProgressResponse
	if err := s.client.GetWithPathParams(ctx, "/v1/project/projects/{project_id}/progress", pathParams, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get project progress: %w", err)
	}

	return &projectmodel.Progress{
		CompletedCount:  resp.CompletedCount,
		TotalCount:      resp.TotalCount,
		CompletionRate:  resp.CompletionRate,
	}, nil
}

// Create 创建项目
func (s *Service) Create(ctx context.Context, input projectmodel.ProjectCreateInput) (*projectmodel.Project, error) {
	// 参数校验
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if len(input.Name) > 255 {
		return nil, fmt.Errorf("name must be less than 255 characters")
	}
	if input.Type == "" {
		return nil, fmt.Errorf("type is required")
	}
	if !isValidProjectType(input.Type) {
		return nil, fmt.Errorf("type must be one of: scrum, kanban, waterfall, hybrid")
	}
	if input.Identifier == "" {
		return nil, fmt.Errorf("identifier is required")
	}
	if len(input.Identifier) > 15 {
		return nil, fmt.Errorf("identifier must be less than 15 characters")
	}
	if input.ScopeType == "user_group" && input.ScopeID == "" {
		return nil, fmt.Errorf("scope_id is required when scope_type is user_group")
	}

	// 构建请求 DTO
	reqDTO := apiproject.ToCreateRequestDTO(input)

	// 发送 POST 请求
	var resp apiproject.Project
	if err := s.client.Post(ctx, "/v1/project/projects", nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return resp.ToModel(), nil
}

// Update 更新项目
func (s *Service) Update(ctx context.Context, projectID string, input projectmodel.ProjectUpdateInput) (*projectmodel.Project, error) {
	if projectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}

	reqDTO := apiproject.ToUpdateRequestDTO(input)
	pathParams := map[string]string{"project_id": projectID}

	var resp apiproject.Project
	if err := s.client.PatchWithPathParams(ctx, "/v1/project/projects/{project_id}", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	return resp.ToModel(), nil
}

// Delete 删除项目
func (s *Service) Delete(ctx context.Context, projectID string) error {
	if projectID == "" {
		return fmt.Errorf("project_id cannot be empty")
	}

	pathParams := map[string]string{"project_id": projectID}
	if err := s.client.DeleteWithPathParams(ctx, "/v1/project/projects/{project_id}", pathParams, nil); err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}

	return nil
}

// isValidProjectType 验证项目类型是否有效
func isValidProjectType(t string) bool {
	return t == "scrum" || t == "kanban" || t == "waterfall" || t == "hybrid"
}
