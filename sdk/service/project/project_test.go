package project

import (
	"context"
	"encoding/json"
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

// TestProjectService_Create_ParameterValidation 测试创建项目的参数验证
func TestProjectService_Create_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	tests := []struct {
		name        string
		input       projectmodel.ProjectCreateInput
		expectError bool
		errorMsg    string
	}{
		{
			name: "missing name",
			input: projectmodel.ProjectCreateInput{
				Type:       "scrum",
				Identifier: "TEST",
			},
			expectError: true,
			errorMsg:    "name is required",
		},
		{
			name: "missing type",
			input: projectmodel.ProjectCreateInput{
				Name:       "Test Project",
				Identifier: "TEST",
			},
			expectError: true,
			errorMsg:    "type is required",
		},
		{
			name: "invalid type",
			input: projectmodel.ProjectCreateInput{
				Name:       "Test Project",
				Type:       "invalid",
				Identifier: "TEST",
			},
			expectError: true,
			errorMsg:    "type must be one of",
		},
		{
			name: "missing identifier",
			input: projectmodel.ProjectCreateInput{
				Name: "Test Project",
				Type: "scrum",
			},
			expectError: true,
			errorMsg:    "identifier is required",
		},
		{
			name: "name too long",
			input: projectmodel.ProjectCreateInput{
				Name:       string(make([]byte, 256)),
				Type:       "scrum",
				Identifier: "TEST",
			},
			expectError: true,
			errorMsg:    "name must be less than 255 characters",
		},
		{
			name: "identifier too long",
			input: projectmodel.ProjectCreateInput{
				Name:       "Test Project",
				Type:       "scrum",
				Identifier: "VERYLONGIDENTIFIER",
			},
			expectError: true,
			errorMsg:    "identifier must be less than 15 characters",
		},
		{
			name: "scope_type=user_group but missing scope_id",
			input: projectmodel.ProjectCreateInput{
				Name:       "Test Project",
				Type:       "scrum",
				Identifier: "TEST",
				ScopeType:  "user_group",
			},
			expectError: true,
			errorMsg:    "scope_id is required when scope_type is user_group",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Create(context.Background(), tt.input)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), tt.errorMsg)
			}
		})
	}
}

// TestProjectService_Create_Success 测试创建项目成功
func TestProjectService_Create_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/project/projects", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)

		assert.Equal(t, "Test Project", reqBody["name"])
		assert.Equal(t, "scrum", reqBody["type"])
		assert.Equal(t, "TEST", reqBody["identifier"])

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "project-123",
			"url": "https://example.com/projects/project-123",
			"identifier": "TEST",
			"name": "Test Project",
			"type": "scrum",
			"state": {
				"id": "state-123",
				"name": "正常",
				"type": "in_progress"
			},
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

	input := projectmodel.ProjectCreateInput{
		Name:       "Test Project",
		Type:       "scrum",
		Identifier: "TEST",
	}

	result, err := service.Create(context.Background(), input)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "project-123", result.ID)
	assert.Equal(t, "Test Project", result.Name)
	assert.Equal(t, "TEST", result.Identifier)
}

// TestProjectService_Update_ParameterValidation 测试更新项目的参数验证
func TestProjectService_Update_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	// 测试 project_id 为空
	result, err := service.Update(context.Background(), "", projectmodel.ProjectUpdateInput{})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "project_id cannot be empty")
}

// TestProjectService_Update_Success 测试更新项目成功
func TestProjectService_Update_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/v1/project/projects/project-123", r.URL.Path)

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)

		assert.Equal(t, "Updated Name", reqBody["name"])

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "project-123",
			"url": "https://example.com/projects/project-123",
			"identifier": "TEST",
			"name": "Updated Name",
			"type": "scrum",
			"state": {
				"id": "state-123",
				"name": "正常",
				"type": "in_progress"
			},
			"created_at": 1583290300,
			"updated_at": 1583290300,
			"is_archived": 0,
			"is_deleted": 0
		}`))
	}))
	defer server.Close()

	client, _ := httpclient.NewClient(server.URL, 10*time.Second)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	updatedName := "Updated Name"
	input := projectmodel.ProjectUpdateInput{
		Name: &updatedName,
	}

	result, err := service.Update(context.Background(), "project-123", input)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Updated Name", result.Name)
}

// TestProjectService_Delete_ParameterValidation 测试删除项目的参数验证
func TestProjectService_Delete_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	// 测试 project_id 为空
	err := service.Delete(context.Background(), "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "project_id cannot be empty")
}

// TestProjectService_Delete_Success 测试删除项目成功
func TestProjectService_Delete_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/project/projects/project-123", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := httpclient.NewClient(server.URL, 10*time.Second)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	err := service.Delete(context.Background(), "project-123")

	require.NoError(t, err)
}
