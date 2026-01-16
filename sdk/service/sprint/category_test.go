package sprint

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brain-xai/pingcode_api/internal/httpclient"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSprintService_CreateCategory_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	tests := []struct {
		name        string
		input       sprintmodel.SprintCategoryCreateInput
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty project id",
			input:       sprintmodel.SprintCategoryCreateInput{Name: "Category 1"},
			expectError: true,
			errorMsg:    "project_id cannot be empty",
		},
		{
			name:        "empty name",
			input:       sprintmodel.SprintCategoryCreateInput{ProjectID: "proj-123"},
			expectError: true,
			errorMsg:    "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.CreateCategory(context.Background(), tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSprintService_CreateCategory_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/project/projects/proj-123/sprint_categories", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"sprint_category": {
				"id": "category-123",
				"url": "https://api.pingcode.com/v1/project/projects/proj-123/sprint_categories/category-123",
				"name": "Category 1",
				"project": {
					"id": "proj-123",
					"url": "https://api.pingcode.com/v1/project/projects/proj-123",
					"identifier": "SCR",
					"name": "Test Project",
					"type": "scrum",
					"is_archived": 0,
					"is_deleted": 0
				},
				"section": {
					"id": "group-123",
					"url": "https://api.pingcode.com/v1/project/projects/proj-123/sprint_sections/group-123",
					"name": "Group 1"
				}
			}
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	input := sprintmodel.SprintCategoryCreateInput{
		ProjectID: "proj-123",
		Name:      "Category 1",
		GroupID:   "group-123",
	}

	result, err := service.CreateCategory(context.Background(), input)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "category-123", result.ID)
	assert.Equal(t, "Category 1", result.Name)
	assert.Equal(t, "group-123", result.GroupID)
}

func TestSprintService_UpdateCategoryPartially_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	name := "Updated Name"

	tests := []struct {
		name        string
		projectID   string
		categoryID  string
		input       sprintmodel.SprintCategoryUpdateInput
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty project_id",
			projectID:   "",
			categoryID:  "category-123",
			input:       sprintmodel.SprintCategoryUpdateInput{Name: &name},
			expectError: true,
			errorMsg:    "project_id cannot be empty",
		},
		{
			name:        "empty category_id",
			projectID:   "proj-123",
			categoryID:  "",
			input:       sprintmodel.SprintCategoryUpdateInput{Name: &name},
			expectError: true,
			errorMsg:    "category_id cannot be empty",
		},
		{
			name:        "nil name and group_id",
			projectID:   "proj-123",
			categoryID:  "category-123",
			input:       sprintmodel.SprintCategoryUpdateInput{},
			expectError: true,
			errorMsg:    "name or group_id is required for update",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.UpdateCategoryPartially(context.Background(), tt.projectID, tt.categoryID, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSprintService_ListCategories_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, _, err := service.ListCategories(context.Background(), sprintmodel.SprintCategoryFilter{})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "project_id cannot be empty")
}

func TestSprintService_DeleteCategory_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	tests := []struct {
		name        string
		projectID   string
		categoryID  string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty project_id",
			projectID:   "",
			categoryID:  "category-123",
			expectError: true,
			errorMsg:    "project_id cannot be empty",
		},
		{
			name:        "empty category_id",
			projectID:   "proj-123",
			categoryID:  "",
			expectError: true,
			errorMsg:    "category_id cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.DeleteCategory(context.Background(), tt.projectID, tt.categoryID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			}
		})
	}
}

func TestSprintService_DeleteCategory_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/project/projects/proj-123/sprint_categories/category-123", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 返回成功响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"sprint_category": {
				"id": "category-123",
				"url": "https://api.pingcode.com/v1/project/projects/proj-123/sprint_categories/category-123",
				"name": "Category 1",
				"project": {
					"id": "proj-123",
					"url": "https://api.pingcode.com/v1/project/projects/proj-123",
					"identifier": "SCR",
					"name": "Test Project",
					"type": "scrum",
					"is_archived": 0,
					"is_deleted": 0
				}
			}
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	err = service.DeleteCategory(context.Background(), "proj-123", "category-123")

	require.NoError(t, err)
}
