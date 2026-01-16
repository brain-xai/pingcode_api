package sprint

import (
	"context"
	"fmt"
	"net/url"

	apisprint "github.com/brain-xai/pingcode_api/internal/api/sprint"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// Create 创建一个迭代
func (s *Service) Create(ctx context.Context, input sprintmodel.SprintCreateInput) (*sprintmodel.Sprint, error) {
	// 参数校验
	if input.ProjectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if input.StartAt == 0 {
		return nil, fmt.Errorf("start_at is required")
	}
	if input.EndAt == 0 {
		return nil, fmt.Errorf("end_at is required")
	}
	if input.AssigneeID == "" {
		return nil, fmt.Errorf("assignee_id is required")
	}

	// 路径参数
	pathParams := map[string]string{"project_id": input.ProjectID}

	// 构建请求 DTO
	reqDTO := ToCreateRequestDTO(input)

	// 发送请求
	var resp apisprint.CreateSprintResponse
	if err := s.client.PostWithPathParams(ctx, "/v1/project/projects/{project_id}/sprints", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to create sprint: %w", err)
	}

	return resp.Sprint.ToModel(), nil
}

// BatchCreate 批量创建迭代
func (s *Service) BatchCreate(ctx context.Context, inputs []sprintmodel.SprintCreateInput) ([]sprintmodel.Sprint, error) {
	// 参数校验
	if len(inputs) == 0 {
		return nil, fmt.Errorf("inputs cannot be empty")
	}
	if len(inputs) > 100 {
		return nil, fmt.Errorf("cannot create more than 100 sprints at once")
	}

	// 校验每个输入
	for i, input := range inputs {
		if input.ProjectID == "" {
			return nil, fmt.Errorf("inputs[%d]: project_id cannot be empty", i)
		}
		if input.Name == "" {
			return nil, fmt.Errorf("inputs[%d]: name is required", i)
		}
		if input.StartAt == 0 {
			return nil, fmt.Errorf("inputs[%d]: start_at is required", i)
		}
		if input.EndAt == 0 {
			return nil, fmt.Errorf("inputs[%d]: end_at is required", i)
		}
		if input.AssigneeID == "" {
			return nil, fmt.Errorf("inputs[%d]: assignee_id is required", i)
		}
	}

	// 构建请求 DTO
	reqDTO := ToBatchCreateRequestDTO(inputs)

	// 发送请求
	var resp []apisprint.BatchCreateSprintsResponse
	if err := s.client.Post(ctx, "/v1/project/sprints/bulk", nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to batch create sprints: %w", err)
	}

	// 转换响应
	sprints := make([]sprintmodel.Sprint, len(resp))
	for i, item := range resp {
		sprints[i] = *item.Sprint.ToModel()
	}

	return sprints, nil
}

// UpdatePartially 部分更新一个迭代
func (s *Service) UpdatePartially(ctx context.Context, projectID, sprintID string, input sprintmodel.SprintUpdateInput) (*sprintmodel.Sprint, error) {
	// 参数校验
	if projectID == "" {
		return nil, fmt.Errorf("project_id cannot be empty")
	}
	if sprintID == "" {
		return nil, fmt.Errorf("sprint_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{
		"project_id": projectID,
		"sprint_id":  sprintID,
	}

	// 构建请求 DTO
	reqDTO := ToUpdateRequestDTO(input)

	// 发送请求
	var resp apisprint.UpdateSprintResponse
	if err := s.client.PatchWithPathParams(ctx, "/v1/project/projects/{project_id}/sprints/{sprint_id}", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to update sprint: %w", err)
	}

	return resp.Sprint.ToModel(), nil
}

// List 获取迭代列表
func (s *Service) List(ctx context.Context, filter sprintmodel.SprintFilter) ([]sprintmodel.Sprint, *core.Pagination, error) {
	// 参数校验
	if filter.ProjectID == "" {
		return nil, nil, fmt.Errorf("project_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{"project_id": filter.ProjectID}

	// 构建查询参数
	query := url.Values{}
	if filter.Name != "" {
		query.Set("name", filter.Name)
	}
	if filter.Status != "" {
		query.Set("status", filter.Status)
	}
	if filter.CreatedBetween != "" {
		query.Set("created_between", filter.CreatedBetween)
	}
	if filter.UpdatedBetween != "" {
		query.Set("updated_between", filter.UpdatedBetween)
	}

	// 分页参数
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送请求
	var resp apisprint.ListSprintsResponse
	if err := s.client.GetWithPathParams(ctx, "/v1/project/projects/{project_id}/sprints", pathParams, query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list sprints: %w", err)
	}

	// 转换 DTO 为 Model
	sprints := make([]sprintmodel.Sprint, len(resp.Values))
	for i, s := range resp.Values {
		sprints[i] = *s.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return sprints, pagination, nil
}
