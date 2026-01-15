package main

import (
	"context"
	"fmt"
	"log"

	"github.com/brain-xai/pingcode_api/sdk"
	"github.com/brain-xai/pingcode_api/sdk/model/project"
)

func main() {
	// 方式1：从环境变量自动加载配置
	config, err := sdk.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 方式2：手动配置（取消注释使用）
	// config := sdk.NewDefaultConfig()
	// config.BaseURL = "https://open.pingcode.com"
	// config.Auth = &sdk.AuthConfig{
	// 	ClientID:     "your_client_id",
	// 	ClientSecret: "your_client_secret",
	// }

	// 创建客户端（会自动获取 access_token）
	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 获取项目列表
	projects, pagination, err := client.Project().List(context.Background(), project.ProjectFilter{})
	if err != nil {
		log.Fatalf("Failed to list projects: %v", err)
	}

	// 输出结果
	fmt.Printf("Found %d projects:\n", len(projects))
	for _, p := range projects {
		fmt.Printf("- [%s] %s (%s)\n", p.Identifier, p.Name, p.Type)
	}
	fmt.Printf("Total: %d, Page: %d/%d\n", pagination.Total, pagination.PageIndex, pagination.PageSize)
}
