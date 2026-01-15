// +build integration

package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/brain-xai/pingcode_api/sdk"
	projectmodel "github.com/brain-xai/pingcode_api/sdk/model/project"
	globalmodel "github.com/brain-xai/pingcode_api/sdk/model/global"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestClient 创建用于集成测试的 SDK 客户端
// 需要设置环境变量：
// - PINGCODE_TEST_BASE_URL: PingCode API 的基础 URL
// - PINGCODE_TEST_TOKEN: 访问令牌（access_token）
func setupTestClient(t *testing.T) *sdk.Client {
	baseURL := os.Getenv("PINGCODE_TEST_BASE_URL")
	token := os.Getenv("PINGCODE_TEST_TOKEN")

	if baseURL == "" || token == "" {
		t.Skip("PINGCODE_TEST_BASE_URL or PINGCODE_TEST_TOKEN not set")
	}

	cfg := &sdk.Config{
		BaseURL: baseURL,
		Auth: &sdk.AuthConfig{
			AccessToken: token,
		},
	}

	client, err := sdk.NewClient(cfg)
	require.NoError(t, err)
	return client
}

func TestProjectIntegration_ListAndGet(t *testing.T) {
	client := setupTestClient(t)
	ctx := context.Background()

	// 先获取项目列表
	t.Run("ListProjects", func(t *testing.T) {
		projects, pagination, err := client.Project().List(ctx, projectmodel.ProjectFilter{})
		require.NoError(t, err)
		assert.NotNil(t, pagination)

		if len(projects) == 0 {
			t.Skip("no projects found in test environment")
		}

		t.Logf("Found %d projects (total: %d)", len(projects), pagination.Total)

		// 验证项目列表不为空
		assert.Greater(t, len(projects), 0)

		// 验证第一个项目的字段
		firstProject := projects[0]
		assert.NotEmpty(t, firstProject.ID)
		assert.NotEmpty(t, firstProject.Name)
		assert.NotEmpty(t, firstProject.Identifier)
	})

	// 获取第一个项目的详情
	t.Run("GetProject", func(t *testing.T) {
		projects, _, err := client.Project().List(ctx, projectmodel.ProjectFilter{})
		require.NoError(t, err)
		require.Greater(t, len(projects), 0)

		projectID := projects[0].ID
		t.Logf("Getting project details for: %s", projectID)

		project, err := client.Project().Get(ctx, projectID)
		require.NoError(t, err)
		assert.Equal(t, projectID, project.ID)
		assert.NotEmpty(t, project.Name)
		assert.NotEmpty(t, project.Type)

		t.Logf("Project: %s (%s) - %s", project.Name, project.Identifier, project.Type)
	})
}

func TestProjectIntegration_ListMembers(t *testing.T) {
	client := setupTestClient(t)
	ctx := context.Background()

	// 先获取一个项目
	projects, _, err := client.Project().List(ctx, projectmodel.ProjectFilter{})
	require.NoError(t, err)
	if len(projects) == 0 {
		t.Skip("no projects found in test environment")
	}

	projectID := projects[0].ID

	// 获取项目成员
	t.Run("ListProjectMembers", func(t *testing.T) {
		members, pagination, err := client.Project().ListMembers(ctx, projectID)
		require.NoError(t, err)
		assert.NotNil(t, pagination)

		t.Logf("Found %d members for project %s (total: %d)", len(members), projectID, pagination.Total)

		// 打印成员信息
		for _, member := range members {
			if member.Type == "user" && member.User != nil {
				t.Logf("  - User: %s (%s)", member.User.DisplayName, member.User.Name)
			} else if member.Type == "user_group" && member.UserGroup != nil {
				t.Logf("  - Group: %s", member.UserGroup.Name)
			}
		}
	})
}

func TestGlobalIntegration_GetCurrentUser(t *testing.T) {
	client := setupTestClient(t)
	ctx := context.Background()

	t.Run("GetCurrentUser", func(t *testing.T) {
		user, err := client.Global().GetCurrentUser(ctx)
		require.NoError(t, err)
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.DisplayName)

		t.Logf("Current user: %s (%s)", user.DisplayName, user.Name)
		t.Logf("Email: %s", user.Email)
		t.Logf("Status: %s", user.Status)

		if user.Department != nil {
			t.Logf("Department: %s", user.Department.Name)
		}

		if len(user.Preferences) > 0 {
			t.Logf("Preferences: %v", user.Preferences)
		}

		if len(user.Permissions) > 0 {
			t.Logf("Permissions: %d permissions", len(user.Permissions))
		}
	})
}

func TestGlobalIntegration_ListUsers(t *testing.T) {
	client := setupTestClient(t)
	ctx := context.Background()

	t.Run("ListUsers", func(t *testing.T) {
		users, pagination, err := client.Global().ListUsers(ctx, globalmodel.UserFilter{})
		require.NoError(t, err)
		assert.NotNil(t, pagination)

		t.Logf("Found %d users (total: %d)", len(users), pagination.Total)

		// 验证用户列表不为空
		assert.Greater(t, len(users), 0)

		// 打印前几个用户
		for i, user := range users {
			if i >= 5 {
				break
			}
			t.Logf("  - %s (%s) - %s", user.DisplayName, user.Name, user.Email)
		}
	})

	t.Run("ListUsersWithFilter", func(t *testing.T) {
		// 测试带过滤条件的查询
		users, pagination, err := client.Global().ListUsers(ctx, globalmodel.UserFilter{
			Keywords: "test",
		})
		require.NoError(t, err)
		assert.NotNil(t, pagination)

		t.Logf("Found %d users matching 'test' (total: %d)", len(users), pagination.Total)
	})
}

func TestProjectIntegration_ErrorHandling(t *testing.T) {
	client := setupTestClient(t)
	ctx := context.Background()

	t.Run("GetNonExistentProject", func(t *testing.T) {
		// 尝试获取不存在的项目
		_, err := client.Project().Get(ctx, "nonexistent-project-id")
		assert.Error(t, err)
		t.Logf("Expected error when getting non-existent project: %v", err)
	})

	t.Run("GetUserWithEmptyID", func(t *testing.T) {
		// 尝试使用空 ID 获取用户
		_, err := client.Global().GetUser(ctx, "")
		assert.Error(t, err)
		t.Logf("Expected error when getting user with empty ID: %v", err)
	})
}

// TestEndToEnd 端到端测试：完整的业务流程
func TestEndToEnd(t *testing.T) {
	client := setupTestClient(t)
	ctx := context.Background()

	t.Log("=== Running End-to-End Test ===")

	// 1. 获取当前用户信息
	currentUser, err := client.Global().GetCurrentUser(ctx)
	require.NoError(t, err)
	t.Logf("Step 1: Current user is %s (%s)", currentUser.DisplayName, currentUser.Email)

	// 2. 获取项目列表
	projects, _, err := client.Project().List(ctx, projectmodel.ProjectFilter{})
	require.NoError(t, err)
	require.Greater(t, len(projects), 0, "need at least one project")
	t.Logf("Step 2: Found %d projects", len(projects))

	// 3. 获取第一个项目的详情
	firstProject := projects[0]
	projectDetail, err := client.Project().Get(ctx, firstProject.ID)
	require.NoError(t, err)
	t.Logf("Step 3: Project details - %s (%s), type: %s",
		projectDetail.Name, projectDetail.Identifier, projectDetail.Type)

	// 4. 获取项目成员列表
	members, _, err := client.Project().ListMembers(ctx, firstProject.ID)
	require.NoError(t, err)
	t.Logf("Step 4: Project has %d members", len(members))

	// 5. 获取项目状态列表
	states, _, err := client.Project().ListStates(ctx, firstProject.ID)
	require.NoError(t, err)
	t.Logf("Step 5: Project has %d states", len(states))

	// 6. 如果有负责人，获取负责人信息
	if projectDetail.Assignee != nil {
		assigneeID := projectDetail.Assignee.ID
		assignee, err := client.Global().GetUser(ctx, assigneeID)
		require.NoError(t, err)
		t.Logf("Step 6: Project assignee is %s (%s)", assignee.DisplayName, assignee.Email)
	}

	t.Log("=== End-to-End Test Completed Successfully ===")
}

// Example 示例：展示如何使用 SDK
func Example_integration() {
	// 注意：这是一个示例，不会实际运行
	_ = func() {
		// 创建 SDK 客户端
		cfg := &sdk.Config{
			BaseURL: "https://your-pingcode-instance.com",
			Auth: &sdk.AuthConfig{
				ClientID:     "your-client-id",
				ClientSecret: "your-client-secret",
			},
		}

		client, err := sdk.NewClient(cfg)
		if err != nil {
			panic(err)
		}

		ctx := context.Background()

		// 获取当前用户
		user, err := client.Global().GetCurrentUser(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello, %s!\n", user.DisplayName)

		// 获取项目列表
		projects, _, err := client.Project().List(ctx, projectmodel.ProjectFilter{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("You have %d projects\n", len(projects))

		// 获取第一个项目的详情
		if len(projects) > 0 {
			project, err := client.Project().Get(ctx, projects[0].ID)
			if err != nil {
				panic(err)
			}
			fmt.Printf("First project: %s\n", project.Name)
		}
	}
}
