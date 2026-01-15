package project

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brain-xai/pingcode_api/internal/httpclient"
	projectmodel "github.com/brain-xai/pingcode_api/sdk/model/project"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectService_Get_ParameterValidation(t *testing.T) {
	tests := []struct {
		name        string
		projectID   string
		expectError bool
	}{
		{
			name:        "empty project id",
			projectID:   "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, _ := httpclient.NewClient("http://test", 5*time.Second)
			service := NewService(client, "http://test")

			result, err := service.Get(context.Background(), tt.projectID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestProjectService_Get_Success(t *testing.T) {
	// 创建 mock 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/project/projects/test-project-123", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "test-project-123",
			"name": "Test Project",
			"identifier": "TEST",
			"type": "scrum",
			"description": "Test project description",
			"state": {
				"id": "state-123",
				"name": "正常",
				"type": "in_progress"
			},
			"start_at": 1680278400,
			"end_at": 1682870399,
			"created_at": 1583290300,
			"updated_at": 1583290300,
			"assignee": {
				"id": "user-123",
				"name": "john",
				"display_name": "John Doe",
				"avatar": "https://example.com/avatar.png"
			},
			"scope_type": "user_group",
			"scope_id": "group-123",
			"visibility": "private",
			"color": "#F693E7",
			"is_archived": 0,
			"is_deleted": 0
		}`))
	}))
	defer server.Close()

	// 创建客户端
	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	// 调用方法
	result, err := service.Get(context.Background(), "test-project-123")

	// 验证结果
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-project-123", result.ID)
	assert.Equal(t, "Test Project", result.Name)
	assert.Equal(t, "TEST", result.Identifier)
	assert.Equal(t, "scrum", result.Type)
	assert.Equal(t, "正常", result.State)
	assert.NotNil(t, result.Assignee)
	assert.Equal(t, "user-123", result.Assignee.ID)
	assert.Equal(t, "John Doe", result.Assignee.DisplayName)
}

func TestProjectService_ListStates_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, _, err := service.ListStates(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestProjectService_ListMembers_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, _, err := service.ListMembers(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestProjectService_GetProgress_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetProgress(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestProjectService_List_URLConstruction(t *testing.T) {
	// 创建 mock 服务器来验证 URL 构造
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证 URL
		assert.Equal(t, "/v1/project/projects", r.URL.Path)

		// 验证查询参数
		query := r.URL.Query()
		assert.Equal(t, "TEST", query.Get("identifier"))
		assert.Equal(t, "scrum", query.Get("type"))
		assert.Equal(t, "true", query.Get("include_deleted"))
		assert.Equal(t, "true", query.Get("include_archived"))

		// 返回响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 1,
			"values": [{
				"id": "project-123",
				"name": "Test Project",
				"identifier": "TEST",
				"type": "scrum",
				"state": {"id": "s1", "name": "正常", "type": "in_progress"}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := httpclient.NewClient(server.URL, 10*time.Second)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	filter := projectmodel.ProjectFilter{
		Identifier:      "TEST",
		Type:            "scrum",
		IncludeDeleted:  true,
		IncludeArchived: true,
	}

	projects, pagination, err := service.List(context.Background(), filter)

	require.NoError(t, err)
	assert.Len(t, projects, 1)
	assert.Equal(t, "project-123", projects[0].ID)
	assert.Equal(t, 1, pagination.Total)
}
