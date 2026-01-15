package global

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brain-xai/pingcode_api/internal/httpclient"
	globalmodel "github.com/brain-xai/pingcode_api/sdk/model/global"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGlobalService_GetCurrentUser_Success(t *testing.T) {
	// 创建 mock 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/myself", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "user-123",
			"url": "https://example.com/users/user-123",
			"name": "john",
			"display_name": "John Doe",
			"avatar": "https://example.com/avatar.png",
			"email": "john@example.com",
			"mobile": "15000000000",
			"status": "enabled",
			"employee_number": "EMP001",
			"preferences": {
				"locale": "zh-cn",
				"timezone": "Asia/Shanghai"
			},
			"permissions": {
				"agile_create_project": true,
				"agile_configuration": false
			}
		}`))
	}))
	defer server.Close()

	// 创建客户端
	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	// 调用方法
	result, err := service.GetCurrentUser(context.Background())

	// 验证结果
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "user-123", result.ID)
	assert.Equal(t, "john", result.Name)
	assert.Equal(t, "John Doe", result.DisplayName)
	assert.Equal(t, "john@example.com", result.Email)
	assert.Equal(t, "enabled", result.Status)
	assert.Equal(t, "EMP001", result.EmployeeNumber)
	assert.NotNil(t, result.Preferences)
	assert.Equal(t, "zh-cn", result.Preferences["locale"])
	assert.NotNil(t, result.Permissions)
	assert.True(t, result.Permissions["agile_create_project"])
	assert.False(t, result.Permissions["agile_configuration"])
}

func TestGlobalService_GetUser_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, err := service.GetUser(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGlobalService_GetUser_Success(t *testing.T) {
	// 创建 mock 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/directory/users/user-456", r.URL.Path)

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "user-456",
			"name": "jane",
			"display_name": "Jane Smith",
			"email": "jane@example.com",
			"department": {
				"id": "dept-123",
				"name": "Engineering"
			},
			"job": {
				"id": "job-123",
				"name": "Engineer"
			}
		}`))
	}))
	defer server.Close()

	client, _ := httpclient.NewClient(server.URL, 10*time.Second)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	result, err := service.GetUser(context.Background(), "user-456")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "user-456", result.ID)
	assert.Equal(t, "Jane Smith", result.DisplayName)
	assert.NotNil(t, result.Department)
	assert.Equal(t, "Engineering", result.Department.Name)
	assert.NotNil(t, result.Job)
	assert.Equal(t, "Engineer", result.Job.Name)
}

func TestGlobalService_ListUsers_URLConstruction(t *testing.T) {
	// 创建 mock 服务器来验证 URL 构造
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证 URL
		assert.Equal(t, "/v1/directory/users", r.URL.Path)

		// 验证查询参数
		query := r.URL.Query()
		assert.Equal(t, "john", query.Get("name"))
		assert.Equal(t, "john smith", query.Get("keywords"))
		assert.Equal(t, "dept1,dept2", query.Get("department_ids"))

		// 返回响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 1,
			"values": [{
				"id": "user-789",
				"name": "john",
				"display_name": "John Smith"
			}]
		}`))
	}))
	defer server.Close()

	client, _ := httpclient.NewClient(server.URL, 10*time.Second)
	client.SetToken("test-token")
	service := NewService(client, server.URL)

	filter := globalmodel.UserFilter{
		Name:          "john",
		Keywords:      "john smith",
		DepartmentIDs: "dept1,dept2",
	}

	users, pagination, err := service.ListUsers(context.Background(), filter)

	require.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "user-789", users[0].ID)
	assert.Equal(t, 1, pagination.Total)
}
