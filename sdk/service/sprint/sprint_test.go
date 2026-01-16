package sprint

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brain-xai/pingcode_api/internal/httpclient"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSprintService_Create_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	tests := []struct {
		name        string
		input       sprintmodel.SprintCreateInput
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty project id",
			input:       sprintmodel.SprintCreateInput{Name: "Sprint 1", StartAt: 100, EndAt: 200, AssigneeID: "user-123"},
			expectError: true,
			errorMsg:    "project_id cannot be empty",
		},
		{
			name:        "empty name",
			input:       sprintmodel.SprintCreateInput{ProjectID: "proj-123", StartAt: 100, EndAt: 200, AssigneeID: "user-123"},
			expectError: true,
			errorMsg:    "name is required",
		},
		{
			name:        "empty start_at",
			input:       sprintmodel.SprintCreateInput{ProjectID: "proj-123", Name: "Sprint 1", EndAt: 200, AssigneeID: "user-123"},
			expectError: true,
			errorMsg:    "start_at is required",
		},
		{
			name:        "empty end_at",
			input:       sprintmodel.SprintCreateInput{ProjectID: "proj-123", Name: "Sprint 1", StartAt: 100, AssigneeID: "user-123"},
			expectError: true,
			errorMsg:    "end_at is required",
		},
		{
			name:        "empty assignee_id",
			input:       sprintmodel.SprintCreateInput{ProjectID: "proj-123", Name: "Sprint 1", StartAt: 100, EndAt: 200},
			expectError: true,
			errorMsg:    "assignee_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Create(context.Background(), tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSprintService_Create_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/project/projects/proj-123/sprints", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// 验证请求体
		var req map[string]interface{}
		json.NewDecoder(r.Body).Decode(&req)
		assert.Equal(t, "Sprint 1", req["name"])
		assert.Equal(t, float64(1589791860), req["start_at"])
		assert.Equal(t, float64(1589878260), req["end_at"])
		assert.Equal(t, "user-123", req["assignee_id"])

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"sprint": {
				"id": "sprint-123",
				"url": "https://api.pingcode.com/v1/project/projects/proj-123/sprints/sprint-123",
				"name": "Sprint 1",
				"status": "pending",
				"start_at": 1589791860,
				"end_at": 1589878260,
				"description": "Test sprint",
				"started_at": 0,
				"completed_at": 0,
				"total_story_points": 0,
				"started_story_points": 0,
				"completed_story_points": 0,
				"created_at": 1676454024,
				"updated_at": 1676454024,
				"project": {
					"id": "proj-123",
					"url": "https://api.pingcode.com/v1/project/projects/proj-123",
					"identifier": "SCR",
					"name": "Test Project",
					"type": "scrum",
					"is_archived": 0,
					"is_deleted": 0
				},
				"assignee": {
					"id": "user-123",
					"url": "https://api.pingcode.com/v1/directory/users/user-123",
					"name": "john",
					"display_name": "John",
					"avatar": "https://example.com/avatar.png"
				},
				"categories": [],
				"created_by": {
					"id": "user-123",
					"name": "john",
					"display_name": "John"
				},
				"updated_by": {
					"id": "user-123",
					"name": "john",
					"display_name": "John"
				}
			}
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	input := sprintmodel.SprintCreateInput{
		ProjectID:  "proj-123",
		Name:       "Sprint 1",
		StartAt:    1589791860,
		EndAt:      1589878260,
		AssigneeID: "user-123",
		Description: "Test sprint",
		Status:     "pending",
	}

	result, err := service.Create(context.Background(), input)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "sprint-123", result.ID)
	assert.Equal(t, "Sprint 1", result.Name)
	assert.Equal(t, "pending", result.Status)
}

func TestSprintService_BatchCreate_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	t.Run("empty inputs", func(t *testing.T) {
		result, err := service.BatchCreate(context.Background(), []sprintmodel.SprintCreateInput{})
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "inputs cannot be empty")
	})

	t.Run("more than 100 sprints", func(t *testing.T) {
		inputs := make([]sprintmodel.SprintCreateInput, 101)
		result, err := service.BatchCreate(context.Background(), inputs)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "cannot create more than 100 sprints")
	})
}

func TestSprintService_UpdatePartially_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	t.Run("empty project_id", func(t *testing.T) {
		result, err := service.UpdatePartially(context.Background(), "", "sprint-123", sprintmodel.SprintUpdateInput{})
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "project_id cannot be empty")
	})

	t.Run("empty sprint_id", func(t *testing.T) {
		result, err := service.UpdatePartially(context.Background(), "proj-123", "", sprintmodel.SprintUpdateInput{})
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "sprint_id cannot be empty")
	})
}

func TestSprintService_List_ParameterValidation(t *testing.T) {
	client, _ := httpclient.NewClient("http://test", 5*time.Second)
	service := NewService(client, "http://test")

	result, _, err := service.List(context.Background(), sprintmodel.SprintFilter{})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "project_id cannot be empty")
}

func TestSprintService_List_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/project/projects/proj-123/sprints", r.URL.Path)

		// 验证查询参数
		query := r.URL.Query()
		assert.Equal(t, "Sprint 1", query.Get("name"))
		assert.Equal(t, "in_progress", query.Get("status"))

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_size": 30,
			"page_index": 0,
			"total": 1,
			"values": [
				{
					"id": "sprint-123",
					"url": "https://api.pingcode.com/v1/project/projects/proj-123/sprints/sprint-123",
					"name": "Sprint 1",
					"status": "in_progress",
					"start_at": 1589791860,
					"end_at": 1589878260,
					"description": "Test sprint",
					"started_at": 1589791860,
					"completed_at": 0,
					"total_story_points": 0,
					"started_story_points": 0,
					"completed_story_points": 0,
					"created_at": 1676454024,
					"updated_at": 1676454024,
					"project": {
						"id": "proj-123",
						"url": "https://api.pingcode.com/v1/project/projects/proj-123",
						"identifier": "SCR",
						"name": "Test Project",
						"type": "scrum",
						"is_archived": 0,
						"is_deleted": 0
					},
					"assignee": {
						"id": "user-123",
						"name": "john",
						"display_name": "John"
					},
					"categories": [],
					"created_by": {
						"id": "user-123",
						"name": "john"
					},
					"updated_by": {
						"id": "user-123",
						"name": "john"
					}
				}
			]
		}`))
	}))
	defer server.Close()

	client, err := httpclient.NewClient(server.URL, 10*time.Second)
	require.NoError(t, err)
	client.SetToken("test-token")

	service := NewService(client, server.URL)

	filter := sprintmodel.SprintFilter{
		ProjectID: "proj-123",
		Name:      "Sprint 1",
		Status:    "in_progress",
	}

	result, pagination, err := service.List(context.Background(), filter)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "sprint-123", result[0].ID)
	assert.NotNil(t, pagination)
	assert.Equal(t, 30, pagination.PageSize)
	assert.Equal(t, 0, pagination.PageIndex)
	assert.Equal(t, 1, pagination.Total)
}
