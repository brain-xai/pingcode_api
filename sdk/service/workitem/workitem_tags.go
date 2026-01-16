package workitem

import (
	"context"
	"fmt"
)

// AddTag 向工作项添加标签
func (s *Service) AddTag(ctx context.Context, workItemID string, tagID string) error {
	// 参数校验
	if workItemID == "" {
		return fmt.Errorf("work_item_id cannot be empty")
	}
	if tagID == "" {
		return fmt.Errorf("tag_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID, "tag_id": tagID}

	// 发送 POST 请求
	if err := s.client.PostWithPathParams(ctx, "/v1/project/work_items/{work_item_id}/tags/{tag_id}", pathParams, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to add tag to work item: %w", err)
	}

	return nil
}

// RemoveTag 从工作项移除标签
func (s *Service) RemoveTag(ctx context.Context, workItemID string, tagID string) error {
	// 参数校验
	if workItemID == "" {
		return fmt.Errorf("work_item_id cannot be empty")
	}
	if tagID == "" {
		return fmt.Errorf("tag_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{"work_item_id": workItemID, "tag_id": tagID}

	// 发送 DELETE 请求
	if err := s.client.DeleteWithPathParams(ctx, "/v1/project/work_items/{work_item_id}/tags/{tag_id}", pathParams, nil); err != nil {
		return fmt.Errorf("failed to remove tag from work item: %w", err)
	}

	return nil
}
