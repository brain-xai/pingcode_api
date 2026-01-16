package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brain-xai/pingcode_api/sdk"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

func main() {
	// 从环境变量加载配置
	config, err := sdk.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	projectID := os.Getenv("PINGCODE_PROJECT_ID")
	// 创建客户端
	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("=== Sprint 管理示例 ===\n")

	// 1. 创建迭代分组
	fmt.Println("1. 创建迭代分组...")
	groupName := fmt.Sprintf("测试分组 - %d", time.Now().Unix())
	group, err := client.Sprint().CreateGroup(ctx, sprintmodel.SprintGroupCreateInput{
		ProjectID: projectID,
		Name:      groupName,
	})
	if err != nil {
		log.Printf("创建分组失败: %v", err)
	} else {
		fmt.Printf("   创建成功: %s (ID: %s)\n", group.Name, group.ID)
	}

	// 如果创建分组失败，使用默认项目 ID 继续演示
	groupID := ""
	if group != nil {
		groupID = group.ID
	}

	// 2. 创建迭代类别
	fmt.Println("\n2. 创建迭代类别...")
	categoryName := fmt.Sprintf("测试类别 - %d", time.Now().Unix())
	category, err := client.Sprint().CreateCategory(ctx, sprintmodel.SprintCategoryCreateInput{
		ProjectID: projectID,
		Name:      categoryName,
		GroupID:   groupID, // 可选：关联到分组
	})
	if err != nil {
		log.Printf("创建类别失败: %v", err)
	} else {
		fmt.Printf("   创建成功: %s (ID: %s)\n", category.Name, category.ID)
		if category.Group != nil {
			fmt.Printf("   所属分组: %s\n", category.Group.Name)
		}
	}

	// 类别 ID 用于创建迭代
	categoryIDs := []string{}
	if category != nil {
		categoryIDs = append(categoryIDs, category.ID)
	}

	// 3. 创建迭代
	fmt.Println("\n3. 创建迭代...")
	now := time.Now()
	sprintName := fmt.Sprintf("Sprint %d", now.Unix())
	sprint, err := client.Sprint().Create(ctx, sprintmodel.SprintCreateInput{
		ProjectID:   projectID,
		Name:        sprintName,
		StartAt:     now.Add(24 * time.Hour).Unix(),
		EndAt:       now.Add(14 * 24 * time.Hour).Unix(),
		AssigneeID:  os.Getenv("PINGCODE_USER_ID"), // 可选：设置负责人
		Description: "这是一个测试迭代",
		Status:      "pending",
		CategoryIDs: categoryIDs, // 可选：关联类别
	})
	if err != nil {
		log.Printf("创建迭代失败: %v", err)
	} else {
		fmt.Printf("   创建成功: %s (ID: %s)\n", sprint.Name, sprint.ID)
		fmt.Printf("   状态: %s\n", sprint.Status)
		fmt.Printf("   时间: %d - %d\n", sprint.StartAt, sprint.EndAt)
		if len(sprint.Categories) > 0 {
			fmt.Printf("   类别: ")
			for _, cat := range sprint.Categories {
				fmt.Printf("%s ", cat.Name)
			}
			fmt.Println()
		}
	}

	// 4. 批量创建迭代
	fmt.Println("\n4. 批量创建迭代...")
	batchInputs := []sprintmodel.SprintCreateInput{
		{
			ProjectID:  projectID,
			Name:       fmt.Sprintf("Sprint Batch 1 - %d", now.Unix()),
			StartAt:    now.Add(15 * 24 * time.Hour).Unix(),
			EndAt:      now.Add(29 * 24 * time.Hour).Unix(),
			AssigneeID: os.Getenv("PINGCODE_USER_ID"),
			Status:     "pending",
		},
		{
			ProjectID:  projectID,
			Name:       fmt.Sprintf("Sprint Batch 2 - %d", now.Unix()),
			StartAt:    now.Add(30 * 24 * time.Hour).Unix(),
			EndAt:      now.Add(44 * 24 * time.Hour).Unix(),
			AssigneeID: os.Getenv("PINGCODE_USER_ID"),
			Status:     "pending",
		},
	}
	batchSprints, err := client.Sprint().BatchCreate(ctx, batchInputs)
	if err != nil {
		log.Printf("批量创建迭代失败: %v", err)
	} else {
		fmt.Printf("   批量创建成功: %d 个迭代\n", len(batchSprints))
		for _, s := range batchSprints {
			fmt.Printf("   - %s (ID: %s)\n", s.Name, s.ID)
		}
	}

	// 5. 查询迭代列表
	fmt.Println("\n5. 查询迭代列表...")
	sprints, pagination, err := client.Sprint().List(ctx, sprintmodel.SprintFilter{
		ProjectID: projectID,
		PageSize:  10,
		PageIndex: 0,
	})
	if err != nil {
		log.Printf("查询迭代列表失败: %v", err)
	} else {
		fmt.Printf("   查询成功: 共 %d 个迭代，当前页显示 %d 个\n", pagination.Total, len(sprints))
		for i, s := range sprints {
			fmt.Printf("   %d. %s (状态: %s)\n", i+1, s.Name, s.Status)
		}
	}

	// 6. 部分更新迭代
	if sprint != nil {
		fmt.Println("\n6. 部分更新迭代...")
		newStatus := "in_progress"
		updatedSprint, err := client.Sprint().UpdatePartially(ctx, projectID, sprint.ID, sprintmodel.SprintUpdateInput{
			Status: &newStatus,
		})
		if err != nil {
			log.Printf("更新迭代失败: %v", err)
		} else {
			fmt.Printf("   更新成功: %s (新状态: %s)\n", updatedSprint.Name, updatedSprint.Status)
		}
	}

	// 7. 查询迭代分组列表
	fmt.Println("\n7. 查询迭代分组列表...")
	groups, _, err := client.Sprint().ListGroups(ctx, sprintmodel.SprintGroupFilter{
		ProjectID: projectID,
	})
	if err != nil {
		log.Printf("查询分组列表失败: %v", err)
	} else {
		fmt.Printf("   查询成功: %d 个分组\n", len(groups))
		for _, g := range groups {
			fmt.Printf("   - %s (ID: %s)\n", g.Name, g.ID)
		}
	}

	// 8. 查询迭代类别列表
	fmt.Println("\n8. 查询迭代类别列表...")
	categories, _, err := client.Sprint().ListCategories(ctx, sprintmodel.SprintCategoryFilter{
		ProjectID: projectID,
	})
	if err != nil {
		log.Printf("查询类别列表失败: %v", err)
	} else {
		fmt.Printf("   查询成功: %d 个类别\n", len(categories))
		for _, c := range categories {
			fmt.Printf("   - %s (ID: %s)", c.Name, c.ID)
			if c.Group != nil {
				fmt.Printf(" [分组: %s]", c.Group.Name)
			}
			fmt.Println()
		}
	}

	fmt.Println("\n=== 示例完成 ===")
}
