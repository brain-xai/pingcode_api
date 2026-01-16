package sprint

import (
	"context"
	"fmt"
	"net/url"

	apisprint "github.com/brain-xai/pingcode_api/internal/api/sprint"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// CreateGroup 创建一个迭代分组
func (s *Service) CreateGroup(ctx context.Context, input sprintmodel.SprintGroupCreateInput) (*sprintmodel.SprintGroup, error) {
	// 参数校验
	if input.ProjectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	// 路径参数
	pathParams := map[string]string{"project_id": input.ProjectID}

	// 构建请求 DTO
	reqDTO := ToCreateGroupRequestDTO(input)

	// 发送请求
	var resp apisprint.CreateSprintGroupResponse
	if err := s.client.PostWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_sections", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to create sprint group: %w", err)
	}

	return resp.Group.ToModel(), nil
}

// UpdateGroupPartially 部分更新一个迭代分组
func (s *Service) UpdateGroupPartially(ctx context.Context, projectID, groupID string, input sprintmodel.SprintGroupUpdateInput) (*sprintmodel.SprintGroup, error) {
	// 参数校验
	if projectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}
	if groupID == "" {
		return nil, fmt.Errorf("group_id cannot be empty")
	}
	if input.Name == nil {
		return nil, fmt.Errorf("name is required for update")
	}

	// 路径参数
	pathParams := map[string]string{
		"project_id":  projectID,
		"section_id":  groupID, // API 中称为 section_id
	}

	// 构建请求 DTO
	reqDTO := ToUpdateGroupRequestDTO(input)

	// 发送请求
	var resp apisprint.UpdateSprintGroupResponse
	if err := s.client.PatchWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_sections/{section_id}", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to update sprint group: %w", err)
	}

	return resp.Group.ToModel(), nil
}

// ListGroups 获取迭代分组列表
func (s *Service) ListGroups(ctx context.Context, filter sprintmodel.SprintGroupFilter) ([]sprintmodel.SprintGroup, *core.Pagination, error) {
	// 参数校验
	if filter.ProjectID == "" {
		return nil, nil, fmt.Errorf("project_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{"project_id": filter.ProjectID}

	// 构建查询参数
	query := url.Values{}
	// 分页参数
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送请求
	var resp apisprint.ListSprintGroupsResponse
	if err := s.client.GetWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_sections", pathParams, query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list sprint groups: %w", err)
	}

	// 转换 DTO 为 Model
	groups := make([]sprintmodel.SprintGroup, len(resp.Values))
	for i, g := range resp.Values {
		groups[i] = *g.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return groups, pagination, nil
}

// DeleteGroup 删除一个迭代分组
func (s *Service) DeleteGroup(ctx context.Context, projectID, groupID string) error {
	// 参数校验
	if projectID == "" {
		return fmt.Errorf("project_id cannot be empty")
	}
	if groupID == "" {
		return fmt.Errorf("group_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{
		"project_id":  projectID,
		"section_id":  groupID, // API 中称为 section_id
	}

	// 发送请求
	if err := s.client.DeleteWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_sections/{section_id}", pathParams, nil); err != nil {
		return fmt.Errorf("failed to delete sprint group: %w", err)
	}

	return nil
}
