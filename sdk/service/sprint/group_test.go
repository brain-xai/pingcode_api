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

func TestSprintService_CreateGroup_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	tests := []struct {
		name        string
		input       sprintmodel.SprintGroupCreateInput
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty project id",
			input:       sprintmodel.SprintGroupCreateInput{Name: "Group 1"},
			expectError: true,
			errorMsg:    "project_id cannot be empty",
		},
		{
			name:        "empty name",
			input:       sprintmodel.SprintGroupCreateInput{ProjectID: "proj-123"},
			expectError: true,
			errorMsg:    "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.CreateGroup(context.Background(), tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSprintService_CreateGroup_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/project/projects/proj-123/sprint_sections", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"sprint_section": {
				"id": "group-123",
				"url": "https://api.pingcode.com/v1/project/projects/proj-123/sprint_sections/group-123",
				"name": "Group 1",
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

	input := sprintmodel.SprintGroupCreateInput{
		ProjectID: "proj-123",
		Name:      "Group 1",
	}

	result, err := service.CreateGroup(context.Background(), input)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "group-123", result.ID)
	assert.Equal(t, "Group 1", result.Name)
}

func TestSprintService_UpdateGroupPartially_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	name := "Updated Name"

	tests := []struct {
		name        string
		projectID   string
		groupID     string
		input       sprintmodel.SprintGroupUpdateInput
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty project_id",
			projectID:   "",
			groupID:     "group-123",
			input:       sprintmodel.SprintGroupUpdateInput{Name: &name},
			expectError: true,
			errorMsg:    "project_id cannot be empty",
		},
		{
			name:        "empty group_id",
			projectID:   "proj-123",
			groupID:     "",
			input:       sprintmodel.SprintGroupUpdateInput{Name: &name},
			expectError: true,
			errorMsg:    "group_id cannot be empty",
		},
		{
			name:        "nil name",
			projectID:   "proj-123",
			groupID:     "group-123",
			input:       sprintmodel.SprintGroupUpdateInput{Name: nil},
			expectError: true,
			errorMsg:    "name is required for update",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.UpdateGroupPartially(context.Background(), tt.projectID, tt.groupID, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSprintService_ListGroups_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, _, err := service.ListGroups(context.Background(), sprintmodel.SprintGroupFilter{})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "project_id cannot be empty")
}

func TestSprintService_DeleteGroup_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	tests := []struct {
		name        string
		projectID   string
		groupID     string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty project_id",
			projectID:   "",
			groupID:     "group-123",
			expectError: true,
			errorMsg:    "project_id cannot be empty",
		},
		{
			name:        "empty group_id",
			projectID:   "proj-123",
			groupID:     "",
			expectError: true,
			errorMsg:    "group_id cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.DeleteGroup(context.Background(), tt.projectID, tt.groupID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			}
		})
	}
}

func TestSprintService_DeleteGroup_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/project/projects/proj-123/sprint_sections/group-123", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 返回成功响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"sprint_section": {
				"id": "group-123",
				"url": "https://api.pingcode.com/v1/project/projects/proj-123/sprint_sections/group-123",
				"name": "Group 1",
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

	err = service.DeleteGroup(context.Background(), "proj-123", "group-123")

	require.NoError(t, err)
}
