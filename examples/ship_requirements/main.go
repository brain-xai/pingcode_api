package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brain-xai/pingcode_api/sdk"
	shipmodel "github.com/brain-xai/pingcode_api/sdk/model/ship"
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
	fmt.Println("       Ship 需求管理示例")
	fmt.Println("========================================")

	// 1. 获取产品列表
	fmt.Println("\n1. 获取产品列表...")
	products, _, err := client.Ship().ListProducts(ctx, shipmodel.ProductFilter{
		PageSize:  10,
		PageIndex: 0,
	})
	if err != nil {
		log.Fatalf("获取产品列表失败: %v", err)
	}

	if len(products) == 0 {
		fmt.Println("没有找到产品，请先创建产品")
		return
	}

	product := products[0]
	fmt.Printf("使用产品: %s (ID: %s)\n", product.Name, product.ID)

	// 2. 获取该产品的需求列表
	fmt.Println("\n2. 获取需求列表...")
	requirements, pagination, err := client.Ship().ListRequirements(ctx, shipmodel.RequirementFilter{
		ProductID: product.ID,
		PageSize:  10,
		PageIndex: 0,
	})
	if err != nil {
		log.Fatalf("获取需求列表失败: %v", err)
	}

	fmt.Printf("找到 %d 个需求:\n", pagination.Total)
	for _, req := range requirements {
		fmt.Printf("  - [%s] %s (状态: %s, 负责人: %s)\n",
			req.Identifier,
			req.Title,
			getStateName(req.State),
			getUserName(req.Assignee))
	}

	// 3. 创建一个新需求
	fmt.Println("\n3. 创建新需求...")
	newReq, err := client.Ship().CreateRequirement(ctx, shipmodel.RequirementCreateInput{
		ProductID:  product.ID,
		Title:      fmt.Sprintf("测试需求 - %s", time.Now().Format("2006-01-02 15:04:05")),
		Description: "这是一个通过 SDK 创建的测试需求",
	})
	if err != nil {
		log.Fatalf("创建需求失败: %v", err)
	}

	fmt.Printf("创建成功！\n")
	fmt.Printf("  ID: %s\n", newReq.ID)
	fmt.Printf("  编号: %s\n", newReq.Identifier)
	fmt.Printf("  标题: %s\n", newReq.Title)

	// 4. 更新需求
	fmt.Println("\n4. 更新需求...")
	updatedTitle := "已更新的需求标题"
	progress := 0.5

	updatedReq, err := client.Ship().UpdateRequirement(ctx, newReq.ID, shipmodel.RequirementUpdateInput{
		Title:    &updatedTitle,
		Progress: &progress,
	})
	if err != nil {
		log.Fatalf("更新需求失败: %v", err)
	}

	fmt.Printf("更新成功！\n")
	fmt.Printf("  新标题: %s\n", updatedReq.Title)
	fmt.Printf("  新进度: %.0f%%\n", updatedReq.Progress*100)

	// 5. 获取需求详情
	fmt.Println("\n5. 获取需求详情...")
	reqDetail, err := client.Ship().GetRequirement(ctx, newReq.ID)
	if err != nil {
		log.Fatalf("获取需求详情失败: %v", err)
	}

	fmt.Printf("需求详情:\n")
	fmt.Printf("  编号: %s\n", reqDetail.Identifier)
	fmt.Printf("  标题: %s\n", reqDetail.Title)
	fmt.Printf("  描述: %s\n", reqDetail.Description)
	fmt.Printf("  状态: %s\n", getStateName(reqDetail.State))
	fmt.Printf("  进度: %.0f%%\n", reqDetail.Progress*100)

	// 6. 获取产品的需求配置信息（用于 UI 场景）
	fmt.Println("\n6. 获取需求配置信息...")

	// 获取可用状态
	states, err := client.Ship().GetRequirementStates(ctx, product.ID)
	if err != nil {
		log.Printf("获取需求状态失败: %v", err)
	} else {
		fmt.Printf("可用状态 (%d 个):\n", states.Total)
		for _, state := range states.Values {
			fmt.Printf("  - %s (%s)\n", state.Name, state.Type)
		}
	}

	// 获取可用优先级
	priorities, err := client.Ship().GetRequirementPriorities(ctx, product.ID)
	if err != nil {
		log.Printf("获取需求优先级失败: %v", err)
	} else {
		fmt.Printf("可用优先级 (%d 个):\n", priorities.Total)
		for _, priority := range priorities.Values {
			fmt.Printf("  - %s\n", priority.Name)
		}
	}

	// 获取可用模块
	suites, err := client.Ship().GetRequirementSuites(ctx, product.ID)
	if err != nil {
		log.Printf("获取需求模块失败: %v", err)
	} else {
		fmt.Printf("可用模块 (%d 个):\n", suites.Total)
		for _, suite := range suites.Values {
			fmt.Printf("  - %s (%s)\n", suite.Name, suite.Type)
		}
	}

	fmt.Println("\n========================================")
	fmt.Println("示例执行完成！")
	fmt.Println("========================================")
}

// 辅助函数：获取状态名称
func getStateName(state *shipmodel.RequirementState) string {
	if state == nil {
		return "未设置"
	}
	return state.Name
}

// 辅助函数：获取用户名称
func getUserName(user *shipmodel.User) string {
	if user == nil {
		return "未分配"
	}
	return user.DisplayName
}
