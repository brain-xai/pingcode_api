package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brain-xai/pingcode_api/sdk"
	projectmodel "github.com/brain-xai/pingcode_api/sdk/model/project"
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

func main() {
	// 从环境变量加载配置
	config, err := sdk.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建 SDK 客户端
	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	fmt.Println("========================================")
	fmt.Println("       工作项管理示例")
	fmt.Println("========================================")

	// 1. 获取项目列表
	fmt.Println("\n1. 获取项目列表...")
	projects, _, err := client.Project().List(ctx, projectmodel.ProjectFilter{})
	if err != nil {
		log.Fatalf("获取项目列表失败: %v", err)
	}

	if len(projects) == 0 {
		fmt.Println("没有找到项目，请先创建项目")
		return
	}

	project := projects[0]
	fmt.Printf("使用项目: %s (ID: %s)\n", project.Name, project.ID)

	// 2. 获取工作项类型列表
	fmt.Println("\n2. 获取工作项类型列表...")
	types, err := client.WorkItem().ListTypes(ctx, workitemmodel.WorkItemTypeScope{
		ProjectID: project.ID,
	})
	if err != nil {
		log.Printf("获取工作项类型失败: %v", err)
		return
	}

	if len(types) == 0 {
		fmt.Println("没有找到工作项类型")
		return
	}

	workItemType := types[2]
	fmt.Printf("使用工作项类型: %s (ID: %s)\n", workItemType.Name, workItemType.ID)

	// 3. 获取该项目的现有工作项列表
	fmt.Println("\n3. 获取工作项列表...")
	workItems, pagination, err := client.WorkItem().List(ctx, workitemmodel.WorkItemFilter{
		ProjectID: project.ID,
		PageSize:  10,
		PageIndex: 0,
	})
	if err != nil {
		log.Fatalf("获取工作项列表失败: %v", err)
	}

	fmt.Printf("找到 %d 个工作项:\n", pagination.Total)
	for _, wi := range workItems {
		fmt.Printf("  - [%s] %s (类型: %s, 状态: %s)\n",
			wi.Identifier,
			wi.Title,
			getTypeName(wi.Type),
			getStateName(wi.State))
	}

	// 4. 创建一个新工作项
	fmt.Println("\n4. 创建新工作项...")
	newWI, err := client.WorkItem().Create(ctx, workitemmodel.WorkItemCreateInput{
		ProjectID:   project.ID,
		TypeID:      workItemType.ID,
		Title:       fmt.Sprintf("测试工作项 - %s", time.Now().Format("2006-01-02 15:04:05")),
		Description: "这是一个通过 SDK 创建的测试工作项",
	})
	if err != nil {
		log.Fatalf("创建工作项失败: %v", err)
	}

	fmt.Printf("创建成功！\n")
	fmt.Printf("  ID: %s\n", newWI.ID)
	fmt.Printf("  编号: %s\n", newWI.Identifier)
	fmt.Printf("  标题: %s\n", newWI.Title)
	fmt.Printf("  类型: %s\n", newWI.Type.Name)

	// 5. 部分更新工作项
	fmt.Println("\n5. 部分更新工作项...")
	updatedTitle := "已更新的工作项标题"
	description := "已更新的工作项描述"

	updatedWI, err := client.WorkItem().Update(ctx, newWI.ID, workitemmodel.WorkItemUpdateInput{
		Title:       &updatedTitle,
		Description: &description,
	})
	if err != nil {
		log.Fatalf("更新工作项失败: %v", err)
	}

	fmt.Printf("更新成功！\n")
	fmt.Printf("  新标题: %s\n", updatedWI.Title)
	fmt.Printf("  新描述: %s\n", updatedWI.Description)

	// 6. 获取工作项详情
	fmt.Println("\n6. 获取工作项详情...")
	wiDetail, err := client.WorkItem().Get(ctx, newWI.ID)
	if err != nil {
		log.Fatalf("获取工作项详情失败: %v", err)
	}

	fmt.Printf("工作项详情:\n")
	fmt.Printf("  编号: %s\n", wiDetail.Identifier)
	fmt.Printf("  标题: %s\n", wiDetail.Title)
	fmt.Printf("  描述: %s\n", wiDetail.Description)
	fmt.Printf("  类型: %s\n", getTypeName(wiDetail.Type))
	fmt.Printf("  状态: %s\n", getStateName(wiDetail.State))

	// 7. 获取工作项配置信息（用于 UI 场景）
	fmt.Println("\n7. 获取工作项配置信息...")

	// 获取可用状态
	states, err := client.WorkItem().ListStatuses(ctx, workitemmodel.WorkItemStatusFilter{
		ProjectID: project.ID,
		TypeID:    workItemType.ID,
		PageSize:  100,
	})
	if err != nil {
		log.Printf("获取工作项状态失败: %v", err)
	} else {
		fmt.Printf("可用状态 (%d 个):\n", states.Total)
		for _, state := range states.Values {
			fmt.Printf("  - %s (%s)\n", state.Name, state.Type)
		}
	}

	// 获取可用优先级
	priorities, err := client.WorkItem().ListPriorities(ctx, project.ID)
	if err != nil {
		log.Printf("获取工作项优先级失败: %v", err)
	} else {
		fmt.Printf("可用优先级 (%d 个):\n", priorities.Total)
		for _, priority := range priorities.Values {
			fmt.Printf("  - %s\n", priority.Name)
		}
	}

	// 获取可用标签
	tags, err := client.WorkItem().ListTags(ctx, workitemmodel.WorkItemTagFilter{
		ProjectID: project.ID,
		PageSize:  100,
	})
	if err != nil {
		log.Printf("获取工作项标签失败: %v", err)
	} else {
		fmt.Printf("可用标签 (%d 个):\n", tags.Total)
		for _, tag := range tags.Values {
			fmt.Printf("  - %s\n", tag.Name)
		}
	}
	
	// 8. 获取流转记录
	fmt.Println("\n8. 获取工作项流转记录...")
	transitions, err := client.WorkItem().ListTransitions(ctx, newWI.ID)
	if err != nil {
		log.Printf("获取流转记录失败: %v", err)
	} else {
		fmt.Printf("流转记录 (%d 条):\n", transitions.Total)
		for _, transition := range transitions.Values {
			fromState := ""
			toState := ""
			if transition.FromState != nil {
				fromState = transition.FromState.Name
			}
			if transition.ToState != nil {
				toState = transition.ToState.Name
			}
			createdBy := ""
			if transition.CreatedBy != nil {
				createdBy = transition.CreatedBy.DisplayName
			}
			fmt.Printf("  - %s -> %s (操作人: %s)\n", fromState, toState, createdBy)
		}
	}

	fmt.Println("\n========================================")
	fmt.Println("示例执行完成！")
	fmt.Println("========================================")
}

// 辅助函数：获取类型名称
func getTypeName(workItemType *workitemmodel.WorkItemType) string {
	if workItemType == nil {
		return "未设置"
	}
	return workItemType.Name
}

// 辅助函数：获取状态名称
func getStateName(state *workitemmodel.WorkItemStatus) string {
	if state == nil {
		return "未设置"
	}
	return state.Name
}
