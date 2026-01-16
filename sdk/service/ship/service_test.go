package ship

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brain-xai/pingcode_api/internal/httpclient"
	shipmodel "github.com/brain-xai/pingcode_api/sdk/model/ship"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestShipService_GetProduct_ParameterValidation 测试参数验证
func TestShipService_GetProduct_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetProduct(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_GetProduct_Success 测试获取产品成功
func TestShipService_GetProduct_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/ship/products/test-product-123", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "test-product-123",
			"url": "https://example.com/products/test-product-123",
			"name": "Test Product",
			"identifier": "TEST",
			"scope_type": "user_group",
			"scope_id": "group-123",
			"visibility": "private",
			"color": "#F693E7",
			"description": "Test product description",
			"created_at": 1583290300,
			"updated_at": 1583290300,
			"is_archived": 0,
			"is_deleted": 0,
			"members": [{
				"id": "member-123",
				"url": "https://example.com/members/member-123",
				"type": "user",
				"user": {
					"id": "user-123",
					"name": "john",
					"display_name": "John Doe",
					"avatar": "https://example.com/avatar.png"
				}
			}],
			"created_by": {
				"id": "user-123",
				"name": "john",
				"display_name": "John Doe",
				"avatar": "https://example.com/avatar.png"
			}
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	result, err := service.GetProduct(context.Background(), "test-product-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-product-123", result.ID)
	assert.Equal(t, "Test Product", result.Name)
	assert.Equal(t, "TEST", result.Identifier)
	assert.Equal(t, "user_group", result.ScopeType)
	assert.False(t, result.IsArchived)
	assert.False(t, result.IsDeleted)
	assert.Len(t, result.Members, 1)
	assert.NotNil(t, result.CreatedBy)
}

// TestShipService_ListProducts_URLConstruction 测试产品列表 URL 构造
func TestShipService_ListProducts_URLConstruction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/ship/products", r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, "test-keyword", query.Get("query"))
		assert.Equal(t, "50", query.Get("page_size"))
		assert.Equal(t, "1", query.Get("page_index"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 50,
			"page_index": 1,
			"total": 1,
			"values": [{
				"id": "product-123",
				"name": "Test Product",
				"identifier": "TEST",
				"is_archived": 0,
				"is_deleted": 0
			}]
		}`))
	}))
	defer server.Close()

	client, _ := httpclient.NewClient(server.URL, 10*time.Second)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	filter := shipmodel.ProductFilter{
		Query:     "test-keyword",
		PageSize:  50,
		PageIndex: 1,
	}

	products, pagination, err := service.ListProducts(context.Background(), filter)

	require.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "product-123", products[0].ID)
	assert.Equal(t, 1, pagination.Total)
}

// TestShipService_GetRequirement_ParameterValidation 测试需求参数验证
func TestShipService_GetRequirement_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetRequirement(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_GetRequirement_Success 测试获取需求成功
func TestShipService_GetRequirement_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/ship/ideas/req-123", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "req-123",
			"url": "https://example.com/ideas/req-123",
			"title": "Test Requirement",
			"identifier": "TEST-1",
			"short_id": "abc123",
			"html_url": "https://example.com/ideas/req-123/html",
			"product": {
				"id": "product-123",
				"identifier": "TEST",
				"name": "Test Product",
				"is_archived": 0,
				"is_deleted": 0
			},
			"assignee": {
				"id": "user-123",
				"name": "john",
				"display_name": "John Doe",
				"avatar": "https://example.com/avatar.png"
			},
			"state": {
				"id": "state-123",
				"name": "进行中",
				"type": "in_progress"
			},
			"priority": {
				"id": "priority-123",
				"name": "P1"
			},
			"score": 100,
			"progress": 0.5,
			"description": "Test description",
			"created_at": 1583290300,
			"updated_at": 1583290300,
			"is_archived": 0,
			"is_deleted": 0
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	result, err := service.GetRequirement(context.Background(), "req-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "req-123", result.ID)
	assert.Equal(t, "Test Requirement", result.Title)
	assert.Equal(t, "TEST-1", result.Identifier)
	assert.NotNil(t, result.Product)
	assert.Equal(t, "Test Product", result.Product.Name)
	assert.NotNil(t, result.Assignee)
	assert.Equal(t, "John Doe", result.Assignee.DisplayName)
	assert.NotNil(t, result.State)
	assert.Equal(t, "进行中", result.State.Name)
	assert.Equal(t, 100.0, result.Score)
	assert.Equal(t, 0.5, result.Progress)
}

// TestShipService_ListRequirements_URLConstruction 测试需求列表 URL 构造
func TestShipService_ListRequirements_URLConstruction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/ship/ideas", r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, "product-123", query.Get("product_id"))
		assert.Equal(t, "state-123", query.Get("state_id"))
		assert.Equal(t, "priority-123", query.Get("priority_id"))
		assert.Equal(t, "test keyword", query.Get("keywords"))
		assert.Equal(t, "123456,789012", query.Get("created_between"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 1,
			"values": [{
				"id": "req-123",
				"title": "Test Requirement",
				"identifier": "TEST-1",
				"is_archived": 0,
				"is_deleted": 0
			}]
		}`))
	}))
	defer server.Close()

	client, _ := httpclient.NewClient(server.URL, 10*time.Second)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	filter := shipmodel.RequirementFilter{
		ProductID:      "product-123",
		StateID:        "state-123",
		PriorityID:     "priority-123",
		Keywords:       "test keyword",
		CreatedBetween: "123456,789012",
	}

	requirements, pagination, err := service.ListRequirements(context.Background(), filter)

	require.NoError(t, err)
	assert.Len(t, requirements, 1)
	assert.Equal(t, "req-123", requirements[0].ID)
	assert.Equal(t, 1, pagination.Total)
}

// TestShipService_CreateRequirement_ParameterValidation 测试创建需求参数验证
func TestShipService_CreateRequirement_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	tests := []struct {
		name        string
		input       shipmodel.RequirementCreateInput
		expectError bool
	}{
		{
			name: "missing product_id",
			input: shipmodel.RequirementCreateInput{
				Title: "Test Requirement",
			},
			expectError: true,
		},
		{
			name: "missing title",
			input: shipmodel.RequirementCreateInput{
				ProductID: "product-123",
			},
			expectError: true,
		},
		{
			name: "title too long",
			input: shipmodel.RequirementCreateInput{
				ProductID: "product-123",
				Title:     string(make([]byte, 256)),
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.CreateRequirement(context.Background(), tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

// TestShipService_CreateRequirement_Success 测试创建需求成功
func TestShipService_CreateRequirement_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/ship/ideas", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)

		assert.Equal(t, "product-123", reqBody["product_id"])
		assert.Equal(t, "Test Requirement", reqBody["title"])

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "req-123",
			"title": "Test Requirement",
			"identifier": "TEST-1",
			"product": {
				"id": "product-123",
				"identifier": "TEST",
				"name": "Test Product",
				"is_archived": 0,
				"is_deleted": 0
			},
			"is_archived": 0,
			"is_deleted": 0
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	input := shipmodel.RequirementCreateInput{
		ProductID: "product-123",
		Title:     "Test Requirement",
	}

	result, err := service.CreateRequirement(context.Background(), input)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "req-123", result.ID)
	assert.Equal(t, "Test Requirement", result.Title)
}

// TestShipService_UpdateRequirement_ParameterValidation 测试更新需求参数验证
func TestShipService_UpdateRequirement_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.UpdateRequirement(context.Background(), "", shipmodel.RequirementUpdateInput{})

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_UpdateRequirement_Success 测试更新需求成功
func TestShipService_UpdateRequirement_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/v1/ship/ideas/req-123", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)

		assert.Equal(t, "Updated Title", reqBody["title"])
		assert.Equal(t, float64(0.75), reqBody["progress"])

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "req-123",
			"title": "Updated Title",
			"identifier": "TEST-1",
			"progress": 0.75,
			"is_archived": 0,
			"is_deleted": 0
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	updatedTitle := "Updated Title"
	progress := 0.75

	input := shipmodel.RequirementUpdateInput{
		Title:    &updatedTitle,
		Progress: &progress,
	}

	result, err := service.UpdateRequirement(context.Background(), "req-123", input)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "req-123", result.ID)
	assert.Equal(t, "Updated Title", result.Title)
	assert.Equal(t, 0.75, result.Progress)
}

// ============= 需求辅助接口测试 =============

// TestShipService_GetRequirementStates_ParameterValidation 测试获取需求状态列表参数验证
func TestShipService_GetRequirementStates_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetRequirementStates(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_GetRequirementStates_Success 测试获取需求状态列表成功
func TestShipService_GetRequirementStates_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/ship/idea/states", r.URL.Path)
		assert.Equal(t, "product-123", r.URL.Query().Get("product_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 2,
			"values": [
				{
					"id": "state-1",
					"url": "https://example.com/states/state-1",
					"name": "待评审",
					"type": "pending"
				},
				{
					"id": "state-2",
					"url": "https://example.com/states/state-2",
					"name": "进行中",
					"type": "in_progress"
				}
			]
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	result, err := service.GetRequirementStates(context.Background(), "product-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 30, result.PageSize)
	assert.Equal(t, 0, result.PageIndex)
	assert.Equal(t, 2, result.Total)
	assert.Len(t, result.Values, 2)
	assert.Equal(t, "待评审", result.Values[0].Name)
	assert.Equal(t, "pending", result.Values[0].Type)
}

// TestShipService_GetRequirementPriorities_ParameterValidation 测试获取需求优先级列表参数验证
func TestShipService_GetRequirementPriorities_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetRequirementPriorities(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_GetRequirementPriorities_Success 测试获取需求优先级列表成功
func TestShipService_GetRequirementPriorities_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/ship/idea/priorities", r.URL.Path)
		assert.Equal(t, "product-123", r.URL.Query().Get("product_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 3,
			"values": [
				{
					"id": "priority-1",
					"name": "P0"
				},
				{
					"id": "priority-2",
					"name": "P1"
				}
			]
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	result, err := service.GetRequirementPriorities(context.Background(), "product-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, result.Total)
	assert.Len(t, result.Values, 2)
	assert.Equal(t, "P0", result.Values[0].Name)
}

// TestShipService_GetRequirementProperties_ParameterValidation 测试获取需求属性列表参数验证
func TestShipService_GetRequirementProperties_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetRequirementProperties(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_GetRequirementProperties_Success 测试获取需求属性列表成功
func TestShipService_GetRequirementProperties_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/ship/idea/properties", r.URL.Path)
		assert.Equal(t, "product-123", r.URL.Query().Get("product_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 2,
			"values": [
				{
					"id": "backlog_type",
					"url": "https://example.com/properties/backlog_type",
					"name": "需求类型",
					"type": "select",
					"options": [
						{
							"_id": "opt-1",
							"text": "功能需求"
						},
						{
							"_id": "opt-2",
							"text": "体验优化"
						}
					]
				},
				{
					"id": "identifier",
					"url": "https://example.com/properties/identifier",
					"name": "编号",
					"type": "text",
					"options": null
				}
			]
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	result, err := service.GetRequirementProperties(context.Background(), "product-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.Total)
	assert.Len(t, result.Values, 2)
	assert.Equal(t, "需求类型", result.Values[0].Name)
	assert.Equal(t, "select", result.Values[0].Type)
	assert.Len(t, result.Values[0].Options, 2)
	assert.Equal(t, "功能需求", result.Values[0].Options[0].Text)
}

// TestShipService_GetRequirementSuites_ParameterValidation 测试获取需求模块列表参数验证
func TestShipService_GetRequirementSuites_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetRequirementSuites(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_GetRequirementSuites_Success 测试获取需求模块列表成功
func TestShipService_GetRequirementSuites_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/ship/idea/suites", r.URL.Path)
		assert.Equal(t, "product-123", r.URL.Query().Get("product_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 1,
			"values": [
				{
					"id": "suite-1",
					"url": "https://example.com/suites/suite-1",
					"name": "需求模块",
					"type": "module"
				}
			]
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	result, err := service.GetRequirementSuites(context.Background(), "product-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Total)
	assert.Len(t, result.Values, 1)
	assert.Equal(t, "需求模块", result.Values[0].Name)
}

// TestShipService_GetRequirementPlans_ParameterValidation 测试获取需求排期列表参数验证
func TestShipService_GetRequirementPlans_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetRequirementPlans(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestShipService_GetRequirementPlans_Success 测试获取需求排期列表成功
func TestShipService_GetRequirementPlans_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/ship/idea/plans", r.URL.Path)
		assert.Equal(t, "product-123", r.URL.Query().Get("product_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 1,
			"values": [
				{
					"id": "plan-1",
					"url": "https://example.com/plans/plan-1",
					"name": "账号管理"
				}
			]
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	result, err := service.GetRequirementPlans(context.Background(), "product-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Total)
	assert.Len(t, result.Values, 1)
	assert.Equal(t, "账号管理", result.Values[0].Name)
}
