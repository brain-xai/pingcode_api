package sprint

import (
	"context"
	"fmt"
	"net/url"

	apisprint "github.com/brain-xai/pingcode_api/internal/api/sprint"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// CreateCategory 创建一个迭代类别
func (s *Service) CreateCategory(ctx context.Context, input sprintmodel.SprintCategoryCreateInput) (*sprintmodel.SprintCategoryFull, error) {
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
	reqDTO := ToCreateCategoryRequestDTO(input)

	// 发送请求
	var resp apisprint.CreateSprintCategoryResponse
	if err := s.client.PostWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_categories", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to create sprint category: %w", err)
	}

	return resp.Category.ToModel(), nil
}

// UpdateCategoryPartially 部分更新一个迭代类别
func (s *Service) UpdateCategoryPartially(ctx context.Context, projectID, categoryID string, input sprintmodel.SprintCategoryUpdateInput) (*sprintmodel.SprintCategoryFull, error) {
	// 参数校验
	if projectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}
	if categoryID == "" {
		return nil, fmt.Errorf("category_id cannot be empty")
	}
	if input.Name == nil && input.GroupID == nil {
		return nil, fmt.Errorf("name or group_id is required for update")
	}

	// 路径参数
	pathParams := map[string]string{
		"project_id":       projectID,
		"sprint_category_id": categoryID,
	}

	// 构建请求 DTO
	reqDTO := ToUpdateCategoryRequestDTO(input)

	// 发送请求
	var resp apisprint.UpdateSprintCategoryResponse
	if err := s.client.PatchWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_categories/{sprint_category_id}", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to update sprint category: %w", err)
	}

	return resp.Category.ToModel(), nil
}

// ListCategories 获取迭代类别列表
func (s *Service) ListCategories(ctx context.Context, filter sprintmodel.SprintCategoryFilter) ([]sprintmodel.SprintCategoryFull, *core.Pagination, error) {
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
	var resp apisprint.ListSprintCategoriesResponse
	if err := s.client.GetWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_categories", pathParams, query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list sprint categories: %w", err)
	}

	// 转换 DTO 为 Model
	categories := make([]sprintmodel.SprintCategoryFull, len(resp.Values))
	for i, c := range resp.Values {
		categories[i] = *c.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return categories, pagination, nil
}

// DeleteCategory 删除一个迭代类别
func (s *Service) DeleteCategory(ctx context.Context, projectID, categoryID string) error {
	// 参数校验
	if projectID == "" {
		return fmt.Errorf("project_id cannot be empty")
	}
	if categoryID == "" {
		return fmt.Errorf("category_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{
		"project_id":        projectID,
		"sprint_category_id": categoryID,
	}

	// 发送请求
	if err := s.client.DeleteWithPathParams(ctx, "/v1/project/projects/{project_id}/sprint_categories/{sprint_category_id}", pathParams, nil); err != nil {
		return fmt.Errorf("failed to delete sprint category: %w", err)
	}

	return nil
}
