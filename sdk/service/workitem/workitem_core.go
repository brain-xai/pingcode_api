package workitem

import (
	"context"
	"fmt"
	"net/url"

	apiworkitem "github.com/brain-xai/pingcode_api/internal/api/workitem"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// Create 创建工作项
func (s *Service) Create(ctx context.Context, input workitemmodel.WorkItemCreateInput) (*workitemmodel.WorkItem, error) {
	// 参数校验
	if input.ProjectID == "" {
		return nil, fmt.Errorf("project_id is required")
	}
	if input.TypeID == "" {
		return nil, fmt.Errorf("type_id is required")
	}
	if input.Title == "" {
		return nil, fmt.Errorf("title is required")
	}

	// 构建请求 DTO
	reqDTO := apiworkitem.ToCreateRequestDTO(input)

	// 发送 POST 请求
	var resp apiworkitem.WorkItem
	if err := s.client.Post(ctx, "/v1/project/work_items", nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to create work item: %w", err)
	}

	return resp.ToModel(), nil
}

// Update 部分更新工作项
func (s *Service) Update(ctx context.Context, workItemID string, input workitemmodel.WorkItemUpdateInput) (*workitemmodel.WorkItem, error) {
	// 参数校验
	if workItemID == "" {
		return nil, fmt.Errorf("work_item_id cannot be empty")
	}

	// 构建请求 DTO
	reqDTO := apiworkitem.ToUpdateRequestDTO(input)

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID}

	// 发送 PATCH 请求
	var resp apiworkitem.WorkItem
	if err := s.client.PatchWithPathParams(ctx, "/v1/project/work_items/{work_item_id}", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to update work item: %w", err)
	}

	return resp.ToModel(), nil
}

// BatchUpdate 批量部分更新工作项
func (s *Service) BatchUpdate(ctx context.Context, input workitemmodel.WorkItemBatchUpdateInput) (*workitemmodel.WorkItemBatchUpdateResult, error) {
	// 参数校验
	if len(input.WorkItemIDs) == 0 {
		return nil, fmt.Errorf("work_item_ids cannot be empty")
	}

	// 构建请求 DTO
	reqDTO := apiworkitem.ToBatchUpdateRequestDTO(input)

	// 发送 PATCH 请求
	var resp apiworkitem.BatchUpdateWorkItemsResponse
	if err := s.client.Patch(ctx, "/v1/project/work_items/batch", nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to batch update work items: %w", err)
	}

	return &workitemmodel.WorkItemBatchUpdateResult{
		SuccessIDs: resp.SuccessIDs,
		FailedIDs:  resp.FailedIDs,
	}, nil
}

// List 获取工作项列表
func (s *Service) List(ctx context.Context, filter workitemmodel.WorkItemFilter) ([]workitemmodel.WorkItem, *core.Pagination, error) {
	// 参数校验
	if filter.ProjectID == "" {
		return nil, nil, fmt.Errorf("project_id is required")
	}

	// 构建查询参数
	query := url.Values{}
	query.Set("project_id", filter.ProjectID)

	if filter.Query != "" {
		query.Set("query", filter.Query)
	}
	if filter.TypeID != "" {
		query.Set("type_id", filter.TypeID)
	}
	if filter.StateID != "" {
		query.Set("state_id", filter.StateID)
	}
	if filter.PriorityID != "" {
		query.Set("priority_id", filter.PriorityID)
	}
	if filter.AssigneeID != "" {
		query.Set("assignee_id", filter.AssigneeID)
	}
	if filter.ReporterID != "" {
		query.Set("reporter_id", filter.ReporterID)
	}
	if filter.ParentID != "" {
		query.Set("parent_id", filter.ParentID)
	}
	if filter.SprintID != "" {
		query.Set("sprint_id", filter.SprintID)
	}
	if filter.VersionID != "" {
		query.Set("version_id", filter.VersionID)
	}
	if filter.StartAtFrom > 0 {
		query.Set("start_at_from", fmt.Sprintf("%d", filter.StartAtFrom))
	}
	if filter.StartAtTo > 0 {
		query.Set("start_at_to", fmt.Sprintf("%d", filter.StartAtTo))
	}
	if filter.EndAtFrom > 0 {
		query.Set("end_at_from", fmt.Sprintf("%d", filter.EndAtFrom))
	}
	if filter.EndAtTo > 0 {
		query.Set("end_at_to", fmt.Sprintf("%d", filter.EndAtTo))
	}
	if filter.CreatedAfter > 0 {
		query.Set("created_after", fmt.Sprintf("%d", filter.CreatedAfter))
	}
	if filter.UpdatedAfter > 0 {
		query.Set("updated_after", fmt.Sprintf("%d", filter.UpdatedAfter))
	}
	if len(filter.TagIDs) > 0 {
		for _, tagID := range filter.TagIDs {
			query.Add("tag_ids", tagID)
		}
	}
	if filter.IncludeArchived {
		query.Set("include_archived", "1")
	}
	if filter.IncludeDeleted {
		query.Set("include_deleted", "1")
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemsResponse
	if err := s.client.Get(ctx, "/v1/project/work_items", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list work items: %w", err)
	}

	// 转换 DTO 为 Model
	workItems := make([]workitemmodel.WorkItem, len(resp.Values))
	for i, w := range resp.Values {
		workItems[i] = *w.ToModel()
	}

	// 构建分页信息
	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return workItems, pagination, nil
}

// Get 按 ID 获取工作项详情
func (s *Service) Get(ctx context.Context, workItemID string) (*workitemmodel.WorkItem, error) {
	// 参数校验
	if workItemID == "" {
		return nil, fmt.Errorf("work_item_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID}

	// 发送 GET 请求
	var resp apiworkitem.WorkItem
	if err := s.client.GetWithPathParams(ctx, "/v1/project/work_items/{work_item_id}", pathParams, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get work item: %w", err)
	}

	return resp.ToModel(), nil
}

// Delete 删除工作项
func (s *Service) Delete(ctx context.Context, workItemID string) error {
	// 参数校验
	if workItemID == "" {
		return fmt.Errorf("work_item_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID}

	// 发送 DELETE 请求
	if err := s.client.DeleteWithPathParams(ctx, "/v1/project/work_items/{work_item_id}", pathParams, nil); err != nil {
		return fmt.Errorf("failed to delete work item: %w", err)
	}

	return nil
}
