package main

import (
	"context"
	"fmt"
	"log"

	"github.com/brain-xai/pingcode_api/sdk"
	shipmodel "github.com/brain-xai/pingcode_api/sdk/model/ship"
)

// formatTime 格式化时间戳为可读字符串
func formatTime(ts int64) string {
	if ts == 0 {
		return "-"
	}
	return fmt.Sprintf("%d", ts)
}

// formatVisibility 格式化可见性
func formatVisibility(v string) string {
	switch v {
	case "private":
		return "私有"
	case "public":
		return "公开"
	default:
		return v
	}
}

// printProduct 打印产品信息
func printProduct(p shipmodel.Product) {
	fmt.Printf("\n========================================\n")
	fmt.Printf("产品名称: %s\n", p.Name)
	fmt.Printf("产品标识: %s\n", p.Identifier)
	fmt.Printf("产品 ID: %s\n", p.ID)
	fmt.Printf("可见性: %s\n", formatVisibility(p.Visibility))
	fmt.Printf("主题色: %s\n", p.Color)
	fmt.Printf("描述: %s\n", p.Description)
	fmt.Printf("创建时间: %s\n", formatTime(p.CreatedAt))
	fmt.Printf("更新时间: %s\n", formatTime(p.UpdatedAt))
	if p.CreatedBy != nil {
		fmt.Printf("创建人: %s (%s)\n", p.CreatedBy.DisplayName, p.CreatedBy.Name)
	}

	if len(p.Members) > 0 {
		fmt.Printf("\n成员列表 (%d 人):\n", len(p.Members))
		for i, m := range p.Members {
			if m.Type == "user" && m.User != nil {
				fmt.Printf("  %d. %s (%s) - 用户\n", i+1, m.User.DisplayName, m.User.Name)
			} else if m.Type == "user_group" && m.UserGroup != nil {
				fmt.Printf("  %d. %s - 团队\n", i+1, m.UserGroup.Name)
			}
		}
	}
	fmt.Printf("========================================\n")
}

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
	fmt.Println("       获取产品列表示例")
	fmt.Println("========================================")

	// 获取产品列表
	filter := shipmodel.ProductFilter{
		Query:     "",
		PageSize:  50,
		PageIndex: 0,
	}

	products, pagination, err := client.Ship().ListProducts(ctx, filter)
	if err != nil {
		log.Fatalf("获取产品列表失败: %v", err)
	}

	fmt.Printf("\n找到 %d 个产品:\n", pagination.Total)
	for _, p := range products {
		printProduct(p)
	}

	fmt.Printf("\n分页信息: 第 %d 页 (每页 %d 条，共 %d 条)\n",
		pagination.PageIndex,
		pagination.PageSize,
		pagination.Total,
	)

	// 如果有产品，获取第一个产品的详情
	if len(products) > 0 {
		firstProductID := products[0].ID
		fmt.Printf("\n获取产品详情: %s\n", firstProductID)

		product, err := client.Ship().GetProduct(ctx, firstProductID)
		if err != nil {
			log.Fatalf("获取产品详情失败: %v", err)
		}

		fmt.Printf("\n产品详情:\n")
		printProduct(*product)
	}

	// 使用说明
	fmt.Println("\n========================================")
	fmt.Println("使用说明:")
	fmt.Println("========================================")
	fmt.Println("1. 设置环境变量:")
	fmt.Println("   export PINGCODE_BASE_URL=https://open.pingcode.com")
	fmt.Println("   export PINGCODE_CLIENT_ID=your_client_id")
	fmt.Println("   export PINGCODE_CLIENT_SECRET=your_client_secret")
	fmt.Println("")
	fmt.Println("2. 运行此示例:")
	fmt.Println("   cd examples/ship_products")
	fmt.Println("   go run main.go")
}
