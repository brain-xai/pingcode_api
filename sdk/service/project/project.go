package project

import (
	"context"
	"fmt"
	"net/url"

	apiproject "github.com/brain-xai/pingcode_api/internal/api/project"
	"github.com/brain-xai/pingcode_api/internal/httpclient"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	"github.com/brain-xai/pingcode_api/sdk/model/project"
)

// Service 项目服务
type Service struct {
	client  *httpclient.Client
	baseURL string
}

// NewService 创建项目服务
func NewService(client *httpclient.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}

// List 获取项目列表
func (s *Service) List(ctx context.Context, filter project.ProjectFilter) ([]project.Project, *core.Pagination, error) {
	// 构建查询参数
	query := url.Values{}

	if filter.Identifier != "" {
		query.Set("identifier", filter.Identifier)
	}
	if filter.Type != "" {
		query.Set("type", filter.Type)
	}
	if filter.IncludeDeleted {
		query.Set("include_deleted", "true")
	}
	if filter.IncludeArchived {
		query.Set("include_archived", "true")
	}

	// 发送请求
	var resp apiproject.ListProjectsResponse
	if err := s.client.Get(ctx, "/v1/project/projects", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list projects: %w", err)
	}

	// 转换 DTO 为 Model
	projects := make([]project.Project, len(resp.Values))
	for i, p := range resp.Values {
		projects[i] = *p.ToModel()
	}

	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return projects, pagination, nil
}
