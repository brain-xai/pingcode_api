package workitem

import (
	"context"
	"fmt"
	"net/url"

	apiworkitem "github.com/brain-xai/pingcode_api/internal/api/workitem"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// ListLinkTypes 获取关联类型列表
// 注意：projectID 参数当前未被后端使用，保留是为了未来扩展
func (s *Service) ListLinkTypes(ctx context.Context, projectID string) ([]workitemmodel.WorkItemLinkType, error) {
	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemRelationTypesResponse
	if err := s.client.Get(ctx, "/v1/project/work_item/relation_types", nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list work item relation types: %w", err)
	}

	// 转换 DTO 为 Model
	types := make([]workitemmodel.WorkItemLinkType, len(resp.Values))
	for i, t := range resp.Values {
		types[i] = *t.ToModel()
	}

	return types, nil
}

// Link 关联工作项
func (s *Service) Link(ctx context.Context, workItemID string, targetID string, linkTypeID string) (*workitemmodel.WorkItemLink, error) {
	// 参数校验
	if workItemID == "" {
		return nil, fmt.Errorf("work_item_id cannot be empty")
	}
	if targetID == "" {
		return nil, fmt.Errorf("target_work_item_id cannot be empty")
	}
	if linkTypeID == "" {
		return nil, fmt.Errorf("relation_type cannot be empty")
	}

	// 构建请求 DTO
	reqDTO := apiworkitem.CreateWorkItemRelationRequest{
		TargetWorkItemID: targetID,
		RelationType:     linkTypeID,
	}

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID}

	// 发送 POST 请求
	var resp apiworkitem.WorkItemRelation
	if err := s.client.PostWithPathParams(ctx, "/v1/project/work_items/{work_item_id}/relations", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to link work items: %w", err)
	}

	return resp.ToModel(), nil
}

// ListLinks 获取关联列表
func (s *Service) ListLinks(ctx context.Context, workItemID string, filter workitemmodel.WorkItemLinkFilter) ([]workitemmodel.WorkItemLink, *core.Pagination, error) {
	// 参数校验
	if workItemID == "" {
		return nil, nil, fmt.Errorf("work_item_id cannot be empty")
	}

	// 构建查询参数
	query := url.Values{}
	if filter.TypeID != "" {
		query.Set("relation_type", filter.TypeID)
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID}

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemRelationsResponse
	if err := s.client.GetWithPathParams(ctx, "/v1/project/work_items/{work_item_id}/relations", pathParams, query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list work item relations: %w", err)
	}

	// 转换 DTO 为 Model
	links := make([]workitemmodel.WorkItemLink, len(resp.Values))
	for i, l := range resp.Values {
		links[i] = *l.ToModel()
	}

	// 构建分页信息
	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return links, pagination, nil
}

// Unlink 取消关联（已废弃，请使用 UnlinkWithWorkItem）
// Deprecated: 此方法已废弃，因为需要 work_item_id 才能正确删除关联
// 请使用 UnlinkWithWorkItem(ctx, workItemID, linkID) 替代
func (s *Service) Unlink(ctx context.Context, linkID string) error {
	return fmt.Errorf("Unlink is deprecated, please use UnlinkWithWorkItem(ctx, workItemID, linkID) instead")
}

// UnlinkWithWorkItem 取消工作项关联
func (s *Service) UnlinkWithWorkItem(ctx context.Context, workItemID, linkID string) error {
	// 参数校验
	if workItemID == "" {
		return fmt.Errorf("work_item_id cannot be empty")
	}
	if linkID == "" {
		return fmt.Errorf("link_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{
		"work_item_id": workItemID,
		"relation_id":  linkID,
	}

	// 发送 DELETE 请求
	if err := s.client.DeleteWithPathParams(ctx, "/v1/project/work_items/{work_item_id}/relations/{relation_id}", pathParams, nil); err != nil {
		return fmt.Errorf("failed to unlink work items: %w", err)
	}

	return nil
}
