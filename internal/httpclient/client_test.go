package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildPath(t *testing.T) {
	tests := []struct {
		name        string
		template    string
		pathParams  map[string]string
		expected    string
		expectError bool
	}{
		{
			name:       "simple substitution",
			template:   "/v1/project/projects/{project_id}",
			pathParams: map[string]string{"project_id": "123"},
			expected:   "/v1/project/projects/123",
		},
		{
			name:       "multiple substitutions",
			template:   "/v1/project/projects/{project_id}/members/{member_id}",
			pathParams: map[string]string{"project_id": "123", "member_id": "456"},
			expected:   "/v1/project/projects/123/members/456",
		},
		{
			name:       "URL encoding",
			template:   "/v1/project/projects/{project_id}",
			pathParams: map[string]string{"project_id": "abc/def"},
			expected:   "/v1/project/projects/abc%2Fdef",
		},
		{
			name:       "URL encoding with space",
			template:   "/v1/project/projects/{project_id}",
			pathParams: map[string]string{"project_id": "hello world"},
			expected:   "/v1/project/projects/hello%20world",
		},
		{
			name:       "special characters - question mark",
			template:   "/v1/project/projects/{project_id}",
			pathParams: map[string]string{"project_id": "test?query"},
			expected:   "/v1/project/projects/test%3Fquery",
		},
		{
			name:       "special characters - slash",
			template:   "/v1/project/projects/{project_id}",
			pathParams: map[string]string{"project_id": "test/value"},
			expected:   "/v1/project/projects/test%2Fvalue",
		},
		{
			name:        "missing placeholder",
			template:    "/v1/project/projects/{project_id}",
			pathParams:  map[string]string{"wrong_id": "123"},
			expectError: true,
		},
		{
			name:       "empty path params",
			template:   "/v1/project/projects",
			pathParams: map[string]string{},
			expected:   "/v1/project/projects",
		},
		{
			name:       "partial substitution",
			template:   "/v1/project/projects/{project_id}/work_items/{work_item_id}",
			pathParams: map[string]string{"project_id": "123"},
			expected:   "/v1/project/projects/123/work_items/{work_item_id}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{}
			result, err := client.buildPath(tt.template, tt.pathParams)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetWithPathParams(t *testing.T) {
	tests := []struct {
		name           string
		pathTemplate   string
		pathParams     map[string]string
		query          url.Values
		expectedPath   string
		expectedQuery  string
		responseBody   string
		responseStatus int
		expectError    bool
	}{
		{
			name:         "success with path params",
			pathTemplate: "/v1/project/projects/{project_id}",
			pathParams:   map[string]string{"project_id": "123"},
			query:        url.Values{"test": []string{"value"}},
			expectedPath: "/v1/project/projects/123",
			expectedQuery: "test=value",
			responseBody: `{"id": "123", "name": "Test Project"}`,
			responseStatus: http.StatusOK,
		},
		{
			name:         "success with multiple path params",
			pathTemplate: "/v1/project/projects/{project_id}/members/{member_id}",
			pathParams:   map[string]string{"project_id": "123", "member_id": "456"},
			query:        nil,
			expectedPath: "/v1/project/projects/123/members/456",
			expectedQuery: "",
			responseBody: `{"id": "456", "name": "John"}`,
			responseStatus: http.StatusOK,
		},
		{
			name:         "error - invalid path placeholder",
			pathTemplate: "/v1/project/projects/{project_id}",
			pathParams:   map[string]string{"wrong_key": "123"},
			query:        nil,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建 mock 服务器
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// 验证路径
				assert.Equal(t, tt.expectedPath, r.URL.Path)

				// 验证查询参数
				if tt.expectedQuery != "" {
					assert.Equal(t, tt.expectedQuery, r.URL.Query().Encode())
				}

				// 验证 Authorization 头
				assert.Contains(t, r.Header.Get("Authorization"), "Bearer")

				// 返回响应
				w.WriteHeader(tt.responseStatus)
				if tt.responseBody != "" {
					w.Write([]byte(tt.responseBody))
				}
			}))
			defer server.Close()

			// 创建客户端
			client, err := NewClient(server.URL, 5*time.Second)
			require.NoError(t, err)
			client.SetToken("test-token")

			// 发送请求
			var result map[string]interface{}
			err = client.GetWithPathParams(context.Background(), tt.pathTemplate, tt.pathParams, tt.query, &result)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestGetWithPathParams_WithRealHTTPClient(t *testing.T) {
	// 测试实际 HTTP 客户端的行为
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/project/projects/test-123", r.URL.Path)
		assert.Equal(t, "filter=value", r.URL.Query().Encode())
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "test-123", "name": "Test Project", "type": "scrum"}`))
	}))
	defer server.Close()

	client, err := NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	pathParams := map[string]string{"project_id": "test-123"}
	query := url.Values{"filter": []string{"value"}}

	var result map[string]interface{}
	err = client.GetWithPathParams(context.Background(), "/v1/project/projects/{project_id}", pathParams, query, &result)

	require.NoError(t, err)
	assert.Equal(t, "test-123", result["id"])
	assert.Equal(t, "Test Project", result["name"])
	assert.Equal(t, "scrum", result["type"])
}

func TestGetWithPathParams_ErrorHandling(t *testing.T) {
	tests := []struct {
		name           string
		pathTemplate   string
		pathParams     map[string]string
		responseStatus int
		responseBody   string
		expectAPIError bool
	}{
		{
			name:           "404 not found",
			pathTemplate:   "/v1/project/projects/{project_id}",
			pathParams:     map[string]string{"project_id": "nonexistent"},
			responseStatus: http.StatusNotFound,
			responseBody:   `{"code": "resource_not_found", "message": "Project not found"}`,
			expectAPIError: true,
		},
		{
			name:           "401 unauthorized",
			pathTemplate:   "/v1/project/projects/{project_id}",
			pathParams:     map[string]string{"project_id": "123"},
			responseStatus: http.StatusUnauthorized,
			responseBody:   `{"code": "authentication_failed", "message": "Invalid token"}`,
			expectAPIError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.responseStatus)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client, _ := NewClient(server.URL, 5*time.Second)
			client.SetToken("test-token")

			var result map[string]interface{}
			err := client.GetWithPathParams(context.Background(), tt.pathTemplate, tt.pathParams, nil, &result)

			require.Error(t, err)
			// 验证错误发生即可
			assert.Error(t, err)
		})
	}
}
