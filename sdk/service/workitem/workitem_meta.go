package workitem

import (
	"context"
	"fmt"
	"net/url"

	apiworkitem "github.com/brain-xai/pingcode_api/internal/api/workitem"
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// ListTypes 获取工作项类型列表
func (s *Service) ListTypes(ctx context.Context, scope workitemmodel.WorkItemTypeScope) ([]workitemmodel.WorkItemType, error) {
	// 参数校验
	if scope.ProjectID == "" {
		return nil, fmt.Errorf("project_id is required")
	}

	// 构建查询参数
	query := url.Values{}
	query.Set("project_id", scope.ProjectID)

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemTypesResponse
	if err := s.client.Get(ctx, "/v1/project/work_item_types", query, &resp); err != nil {
		return nil, fmt.Errorf("failed to list work item types: %w", err)
	}

	// 转换 DTO 为 Model
	types := make([]workitemmodel.WorkItemType, len(resp.Values))
	for i, t := range resp.Values {
		types[i] = *t.ToModel()
	}

	return types, nil
}

// ListStatuses 获取工作项状态列表
func (s *Service) ListStatuses(ctx context.Context, filter workitemmodel.WorkItemStatusFilter) (*workitemmodel.WorkItemStatusList, error) {
	// 参数校验
	if filter.ProjectID == "" {
		return nil, fmt.Errorf("project_id is required")
	}

	// 构建查询参数
	query := url.Values{}
	query.Set("project_id", filter.ProjectID)

	if filter.TypeID != "" {
		query.Set("type_id", filter.TypeID)
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemStatesResponse
	if err := s.client.Get(ctx, "/v1/project/work_item_states", query, &resp); err != nil {
		return nil, fmt.Errorf("failed to list work item states: %w", err)
	}

	// 转换 DTO 为 Model
	statuses := make([]workitemmodel.WorkItemStatus, len(resp.Values))
	for i, s := range resp.Values {
		statuses[i] = *s.ToModel()
	}

	return &workitemmodel.WorkItemStatusList{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
		Values:    statuses,
	}, nil
}

// ListFields 获取工作项属性列表
func (s *Service) ListFields(ctx context.Context, filter workitemmodel.WorkItemFieldFilter) (*workitemmodel.WorkItemFieldList, error) {
	// 参数校验
	if filter.ProjectID == "" {
		return nil, fmt.Errorf("project_id is required")
	}

	// 构建查询参数
	query := url.Values{}
	query.Set("project_id", filter.ProjectID)

	if filter.TypeID != "" {
		query.Set("type_id", filter.TypeID)
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemFieldsResponse
	if err := s.client.Get(ctx, "/v1/project/work_item_fields", query, &resp); err != nil {
		return nil, fmt.Errorf("failed to list work item fields: %w", err)
	}

	// 转换 DTO 为 Model
	fields := make([]workitemmodel.WorkItemField, len(resp.Values))
	for i, f := range resp.Values {
		fields[i] = *f.ToModel()
	}

	return &workitemmodel.WorkItemFieldList{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
		Values:    fields,
	}, nil
}

// ListPriorities 获取工作项优先级列表
func (s *Service) ListPriorities(ctx context.Context, projectID string) (*workitemmodel.WorkItemPriorityList, error) {
	// 参数校验
	if projectID == "" {
		return nil, fmt.Errorf("project_id is required")
	}

	// 构建查询参数
	query := url.Values{}
	query.Set("project_id", projectID)

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemPrioritiesResponse
	if err := s.client.Get(ctx, "/v1/project/work_item/priorities", query, &resp); err != nil {
		return nil, fmt.Errorf("failed to list work item priorities: %w", err)
	}

	// 转换 DTO 为 Model
	priorities := make([]workitemmodel.WorkItemPriority, len(resp.Values))
	for i, p := range resp.Values {
		priorities[i] = *p.ToModel()
	}

	return &workitemmodel.WorkItemPriorityList{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
		Values:    priorities,
	}, nil
}

// ListTags 获取工作项标签列表
func (s *Service) ListTags(ctx context.Context, filter workitemmodel.WorkItemTagFilter) (*workitemmodel.WorkItemTagList, error) {
	// 参数校验
	if filter.ProjectID == "" {
		return nil, fmt.Errorf("project_id is required")
	}

	// 构建查询参数
	query := url.Values{}
	query.Set("project_id", filter.ProjectID)

	if filter.TypeID != "" {
		query.Set("type_id", filter.TypeID)
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemTagsResponse
	if err := s.client.Get(ctx, "/v1/project/tags", query, &resp); err != nil {
		return nil, fmt.Errorf("failed to list work item tags: %w", err)
	}

	// 转换 DTO 为 Model
	tags := make([]workitemmodel.WorkItemTag, len(resp.Values))
	for i, t := range resp.Values {
		tags[i] = *t.ToModel()
	}

	return &workitemmodel.WorkItemTagList{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
		Values:    tags,
	}, nil
}
