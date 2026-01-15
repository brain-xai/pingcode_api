package ship

import (
	"context"
	"fmt"
	"net/url"

	apiship "github.com/brain-xai/pingcode_api/internal/api/ship"
	"github.com/brain-xai/pingcode_api/internal/httpclient"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
	shipmodel "github.com/brain-xai/pingcode_api/sdk/model/ship"
)

// Service Ship 服务
type Service struct {
	client  *httpclient.Client
	baseURL string
}

// NewService 创建 Ship 服务
func NewService(client *httpclient.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}

// ListProducts 获取产品列表
func (s *Service) ListProducts(ctx context.Context, filter shipmodel.ProductFilter) ([]shipmodel.Product, *core.Pagination, error) {
	// 构建查询参数
	query := url.Values{}

	if filter.Query != "" {
		query.Set("query", filter.Query)
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送 GET 请求
	var resp apiship.ListProductsResponse
	if err := s.client.Get(ctx, "/v1/ship/products", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list products: %w", err)
	}

	// 转换 DTO 为 Model
	products := make([]shipmodel.Product, len(resp.Values))
	for i, p := range resp.Values {
		products[i] = *p.ToModel()
	}

	// 构建分页信息
	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return products, pagination, nil
}

// GetProduct 按 ID 获取产品详情
func (s *Service) GetProduct(ctx context.Context, productID string) (*shipmodel.Product, error) {
	// 参数校验
	if productID == "" {
		return nil, fmt.Errorf("product_id cannot be empty")
	}

	// 构建路径参数
	pathParams := map[string]string{"product_id": productID}

	// 发送 GET 请求
	var resp apiship.Product
	if err := s.client.GetWithPathParams(ctx, "/v1/ship/products/{product_id}", pathParams, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return resp.ToModel(), nil
}

// ListRequirements 获取需求列表
func (s *Service) ListRequirements(ctx context.Context, filter shipmodel.RequirementFilter) ([]shipmodel.Requirement, *core.Pagination, error) {
	// 构建查询参数
	query := url.Values{}

	if filter.ProductID != "" {
		query.Set("product_id", filter.ProductID)
	}
	if filter.StateID != "" {
		query.Set("state_id", filter.StateID)
	}
	if filter.PriorityID != "" {
		query.Set("priority_id", filter.PriorityID)
	}
	if filter.Keywords != "" {
		query.Set("keywords", filter.Keywords)
	}
	if filter.CreatedBetween != "" {
		query.Set("created_between", filter.CreatedBetween)
	}
	if filter.UpdatedBetween != "" {
		query.Set("updated_between", filter.UpdatedBetween)
	}
	if filter.IncludePublicImageToken != "" {
		query.Set("include_public_image_token", filter.IncludePublicImageToken)
	}
	if filter.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	}
	if filter.PageIndex > 0 {
		query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	}

	// 发送 GET 请求
	var resp apiship.ListRequirementsResponse
	if err := s.client.Get(ctx, "/v1/ship/ideas", query, &resp); err != nil {
		return nil, nil, fmt.Errorf("failed to list requirements: %w", err)
	}

	// 转换 DTO 为 Model
	requirements := make([]shipmodel.Requirement, len(resp.Values))
	for i, r := range resp.Values {
		requirements[i] = *r.ToModel()
	}

	// 构建分页信息
	pagination := &core.Pagination{
		PageSize:  resp.PageSize,
		PageIndex: resp.PageIndex,
		Total:     resp.Total,
	}

	return requirements, pagination, nil
}

// GetRequirement 按 ID 获取需求详情
func (s *Service) GetRequirement(ctx context.Context, requirementID string) (*shipmodel.Requirement, error) {
	// 参数校验
	if requirementID == "" {
		return nil, fmt.Errorf("requirement_id cannot be empty")
	}

	// 构建路径参数（注意：API 使用 idea_id）
	pathParams := map[string]string{"idea_id": requirementID}

	// 发送 GET 请求
	var resp apiship.Requirement
	if err := s.client.GetWithPathParams(ctx, "/v1/ship/ideas/{idea_id}", pathParams, nil, &resp); err != nil {
		return nil, fmt.Errorf("failed to get requirement: %w", err)
	}

	return resp.ToModel(), nil
}

// CreateRequirement 创建需求
func (s *Service) CreateRequirement(ctx context.Context, input shipmodel.RequirementCreateInput) (*shipmodel.Requirement, error) {
	// 参数校验
	if input.ProductID == "" {
		return nil, fmt.Errorf("product_id is required")
	}
	if input.Title == "" {
		return nil, fmt.Errorf("title is required")
	}
	if len(input.Title) > 255 {
		return nil, fmt.Errorf("title must be less than 255 characters")
	}

	// 构建请求 DTO
	reqDTO := &apiship.CreateRequirementRequest{
		ProductID:   input.ProductID,
		Title:       input.Title,
		AssigneeID:  input.AssigneeID,
		Description: input.Description,
		SuiteID:     input.SuiteID,
		Properties:  input.Properties,
	}

	// 发送 POST 请求
	var resp apiship.Requirement
	if err := s.client.Post(ctx, "/v1/ship/ideas", nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to create requirement: %w", err)
	}

	return resp.ToModel(), nil
}

// UpdateRequirement 更新需求
func (s *Service) UpdateRequirement(ctx context.Context, requirementID string, input shipmodel.RequirementUpdateInput) (*shipmodel.Requirement, error) {
	// 参数校验
	if requirementID == "" {
		return nil, fmt.Errorf("requirement_id cannot be empty")
	}

	// 构建请求 DTO
	reqDTO := &apiship.UpdateRequirementRequest{
		Title:       input.Title,
		Description: input.Description,
		StateID:     input.StateID,
		PriorityID:  input.PriorityID,
		AssigneeID:  input.AssigneeID,
		Progress:    input.Progress,
		PlanID:      input.PlanID,
		SuiteID:     input.SuiteID,
		Properties:  input.Properties,
	}

	// 转换 TimeRange
	if input.PlanAt != nil {
		reqDTO.PlanAt = &apiship.TimeRange{
			From:        input.PlanAt.From,
			To:          input.PlanAt.To,
			Granularity: input.PlanAt.Granularity,
		}
	}
	if input.RealAt != nil {
		reqDTO.RealAt = &apiship.TimeRange{
			From:        input.RealAt.From,
			To:          input.RealAt.To,
			Granularity: input.RealAt.Granularity,
		}
	}

	// 构建路径参数（注意：API 使用 idea_id）
	pathParams := map[string]string{"idea_id": requirementID}

	// 发送 PATCH 请求
	var resp apiship.Requirement
	if err := s.client.PatchWithPathParams(ctx, "/v1/ship/ideas/{idea_id}", pathParams, nil, reqDTO, &resp); err != nil {
		return nil, fmt.Errorf("failed to update requirement: %w", err)
	}

	return resp.ToModel(), nil
}
