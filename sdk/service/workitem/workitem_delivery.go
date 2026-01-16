package workitem

import (
	"context"
	"fmt"
	"net/url"

	apiworkitem "github.com/brain-xai/pingcode_api/internal/api/workitem"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// CreateDeliveryTarget 创建工作项交付目标
func (s *Service) CreateDeliveryTarget(ctx context.Context, input workitemmodel.WorkItemDeliveryTargetCreateInput) (*workitemmodel.WorkItemDeliveryTarget, error) {
	// 参数校验
	if input.WorkItemID == "" {
		return nil, fmt.Errorf("work_item_id is required")
	}
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	// 构建请求 DTO
	reqDTO := apiworkitem.ToCreateDeliverableRequestDTO(input)

	// 发送 POST 请求
	var resp apiworkitem.WorkItemDeliverable
	if err := s.client.Post(ctx, "/v1/project/deliverables", nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to create delivery target: %w", err)
	}

	return resp.ToModel(), nil
}

// UpdateDeliveryTarget 部分更新工作项交付目标
func (s *Service) UpdateDeliveryTarget(ctx context.Context, targetID string, input workitemmodel.WorkItemDeliveryTargetUpdateInput) (*workitemmodel.WorkItemDeliveryTarget, error) {
	// 参数校验
	if targetID == "" {
		return nil, fmt.Errorf("delivery_target_id cannot be empty")
	}

	// 构建请求 DTO
	reqDTO := apiworkitem.ToUpdateDeliverableRequestDTO(input)

	// 构建路径参数
	pathParams := map[string]string{"deliverable_target_id": targetID}

	// 发送 PATCH 请求
	var resp apiworkitem.WorkItemDeliverable
	if err := s.client.PatchWithPathParams(ctx, "/v1/project/deliverables/{deliverable_target_id}", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to update delivery target: %w", err)
	}

	return resp.ToModel(), nil
}

// ListDeliveryTargets 获取工作项交付目标列表
func (s *Service) ListDeliveryTargets(ctx context.Context, filter workitemmodel.WorkItemDeliveryTargetFilter) ([]workitemmodel.WorkItemDeliveryTarget, *core.Pagination, error) {
	// 参数校验
	if filter.WorkItemID == "" {
		return nil, nil, fmt.Errorf("work_item_id is required")
	}

	// 构建查询参数
	query := url.Values{}
	query.Set("work_item_id", filter.WorkItemID)

	if filter.Status != "" {
		query.Set("status", filter.Status)
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemDeliverablesResponse
	if err := s.client.Get(ctx, "/v1/project/deliverables", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list delivery targets: %w", err)
	}

	// 转换 DTO 为 Model
	targets := make([]workitemmodel.WorkItemDeliveryTarget, len(resp.Values))
	for i, t := range resp.Values {
		targets[i] = *t.ToModel()
	}

	// 构建分页信息
	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return targets, pagination, nil
}

// DeleteDeliveryTarget 删除工作项交付目标
func (s *Service) DeleteDeliveryTarget(ctx context.Context, targetID string) error {
	// 参数校验
	if targetID == "" {
		return fmt.Errorf("delivery_target_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{"deliverable_target_id": targetID}

	// 发送 DELETE 请求
	if err := s.client.DeleteWithPathParams(ctx, "/v1/project/deliverables/{deliverable_target_id}", pathParams, nil); err != nil {
		return fmt.Errorf("failed to delete delivery target: %w", err)
	}

	return nil
}
