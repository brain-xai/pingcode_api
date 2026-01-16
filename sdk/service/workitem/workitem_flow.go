package workitem

import (
	"context"
	"fmt"

	apiworkitem "github.com/brain-xai/pingcode_api/internal/api/workitem"
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// ListTransitions 获取工作项流转记录列表
func (s *Service) ListTransitions(ctx context.Context, workItemID string) (*workitemmodel.WorkItemTransitionList, error) {
	// 参数校验
	if workItemID == "" {
		return nil, fmt.Errorf("work_item_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID}

	// 发送 GET 请求
	var resp apiworkitem.ListWorkItemTransitionsResponse
	if err := s.client.GetWithPathParams(ctx, "/v1/project/work_items/{work_item_id}/transition_histories", pathParams, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to list work item transitions: %w", err)
	}

	// 转换 DTO 为 Model
	transitions := make([]workitemmodel.WorkItemTransition, len(resp.Values))
	for i, t := range resp.Values {
		transitions[i] = *t.ToModel()
	}

	return &workitemmodel.WorkItemTransitionList{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
		Values:    transitions,
	}, nil
}
