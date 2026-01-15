package global

import (
	"context"
	"fmt"
	"net/url"

	apiglobal "github.com/brain-xai/pingcode_api/internal/api/global"
	"github.com/brain-xai/pingcode_api/internal/httpclient"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	"github.com/brain-xai/pingcode_api/sdk/model/global"
)

// Service 全局服务
type Service struct {
	client  *httpclient.Client
	baseURL string
}

// NewService 创建全局服务
func NewService(client *httpclient.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}

// GetCurrentUser 获取当前用户信息（GET /v1/myself）
func (s *Service) GetCurrentUser(ctx context.Context) (*global.User, error) {
	// 发送请求
	var resp apiglobal.User
	if err := s.client.Get(ctx, "/v1/myself", nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	return resp.ToModel(), nil
}

// GetUser 按 ID 获取用户详情（GET /v1/directory/users/{user_id}）
func (s *Service) GetUser(ctx context.Context, userID string) (*global.User, error) {
	// 参数校验
	if userID == "" {
		return nil, fmt.Errorf("user_id cannot be empty")
	}

	// 路径参数
	pathParams := map[string]string{"user_id": userID}

	// 发送请求
	var resp apiglobal.User
	if err := s.client.GetWithPathParams(ctx, "/v1/directory/users/{user_id}", pathParams, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return resp.ToModel(), nil
}

// ListUsers 获取企业成员列表（GET /v1/directory/users）
func (s *Service) ListUsers(ctx context.Context, filter global.UserFilter) ([]global.User, *core.Pagination, error) {
	// 构建查询参数
	query := url.Values{}

	if filter.Name != "" {
		query.Set("name", filter.Name)
	}
	if filter.Keywords != "" {
		query.Set("keywords", filter.Keywords)
	}
	if filter.DepartmentIDs != "" {
		query.Set("department_ids", filter.DepartmentIDs)
	}

	// 发送请求
	var resp apiglobal.ListUsersResponse
	if err := s.client.Get(ctx, "/v1/directory/users", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list users: %w", err)
	}

	// 转换 DTO 为 Model
	users := make([]global.User, len(resp.Values))
	for i, u := range resp.Values {
		users[i] = *u.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return users, pagination, nil
}
