package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brain-xai/pingcode_api/sdk"
	"github.com/brain-xai/pingcode_api/sdk/model/project"
)

func main() {
	// 从环境变量加载配置
	config, err := sdk.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建客户端
	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	// 1. 获取项目列表
	fmt.Println("=== 1. 获取项目列表 ===")
	projects, pagination, err := client.Project().List(ctx, project.ProjectFilter{
		Type:            "scrum",
		IncludeArchived: false,
		IncludeDeleted:  false,
	})
	if err != nil {
		log.Fatalf("获取项目列表失败: %v", err)
	}

	fmt.Printf("找到 %d 个项目（总共 %d 个）：\n", len(projects), pagination.Total)
	for _, p := range projects {
		fmt.Printf("  - [%s] %s (类型: %s, 状态: %s)\n", p.Identifier, p.Name, p.Type, p.State)
	}

	// 如果没有项目，创建一个测试项目
	if len(projects) == 0 {
		fmt.Println("\n没有找到项目，创建一个测试项目...")

		// 2. 创建新项目
		fmt.Println("\n=== 2. 创建新项目 ===")
		startAt := time.Now().Unix()
		endAt := time.Now().AddDate(0, 3, 0).Unix()

		newProject, err := client.Project().Create(ctx, project.ProjectCreateInput{
			Name:        "SDK 测试项目",
			Type:        "scrum",
			Identifier:  fmt.Sprintf("SDK%d", time.Now().Unix()),
			Visibility:  "private",
			Description: "这是一个通过 SDK 创建的测试项目",
			StartAt:     &startAt,
			EndAt:       &endAt,
		})
		if err != nil {
			log.Fatalf("创建项目失败: %v", err)
		}

		fmt.Printf("创建成功：项目 ID=%s, 标识=%s, 名称=%s\n", newProject.ID, newProject.Identifier, newProject.Name)

		// 使用新创建的项目进行后续操作
		projects = []project.Project{*newProject}
	}

	projectID := projects[0].ID

	// 3. 获取项目详情
	fmt.Println("\n=== 3. 获取项目详情 ===")
	projectDetail, err := client.Project().Get(ctx, projectID)
	if err != nil {
		log.Fatalf("获取项目详情失败: %v", err)
	}

	fmt.Printf("项目详情：\n")
	fmt.Printf("  ID: %s\n", projectDetail.ID)
	fmt.Printf("  名称: %s\n", projectDetail.Name)
	fmt.Printf("  标识: %s\n", projectDetail.Identifier)
	fmt.Printf("  类型: %s\n", projectDetail.Type)
	fmt.Printf("  状态: %s\n", projectDetail.State)
	fmt.Printf("  可见性: %s\n", projectDetail.Visibility)
	fmt.Printf("  描述: %s\n", projectDetail.Description)
	if projectDetail.Assignee != nil {
		fmt.Printf("  负责人: %s (%s)\n", projectDetail.Assignee.DisplayName, projectDetail.Assignee.Name)
	}
	fmt.Printf("  创建时间: %s\n", time.Unix(projectDetail.CreatedAt, 0).Format("2006-01-02 15:04:05"))

	// 4. 获取项目成员列表
	fmt.Println("\n=== 4. 获取项目成员列表 ===")
	members, _, err := client.Project().ListMembers(ctx, projectID)
	if err != nil {
		log.Printf("获取项目成员失败: %v", err)
	} else {
		fmt.Printf("项目成员（%d 人）：\n", len(members))
		for _, m := range members {
			if m.Type == "user" && m.User != nil {
				fmt.Printf("  - %s (%s) [用户]\n", m.User.DisplayName, m.User.Name)
			} else if m.Type == "user_group" && m.UserGroup != nil {
				fmt.Printf("  - %s [团队]\n", m.UserGroup.Name)
			}
		}
	}

	// 5. 获取项目进度
	fmt.Println("\n=== 5. 获取项目进度 ===")
	progress, err := client.Project().GetProgress(ctx, projectID)
	if err != nil {
		log.Printf("获取项目进度失败: %v", err)
	} else {
		fmt.Printf("项目进度：\n")
		fmt.Printf("  已完成: %d\n", progress.CompletedCount)
		fmt.Printf("  总数: %d\n", progress.TotalCount)
		fmt.Printf("  完成率: %.2f%%\n", progress.CompletionRate*100)
	}

	// 6. 更新项目（仅对测试项目）
	if len(projects) > 0 && len(projects[0].Identifier) >= 3 && projects[0].Identifier[:3] == "SDK" {
		fmt.Println("\n=== 6. 更新项目 ===")
		newDescription := "更新后的项目描述 - " + time.Now().Format("2006-01-02 15:04:05")

		updatedProject, err := client.Project().Update(ctx, projectID, project.ProjectUpdateInput{
			Description: &newDescription,
		})
		if err != nil {
			log.Printf("更新项目失败: %v", err)
		} else {
			fmt.Printf("更新成功：新描述=%s\n", updatedProject.Description)
		}
	}

	fmt.Println("\n=== 完成 ===")
}
