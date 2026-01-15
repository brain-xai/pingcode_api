package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/brain-xai/pingcode_api/sdk"
	"github.com/brain-xai/pingcode_api/sdk/model/core"
)

// IdeaFilter 需求筛选条件
type IdeaFilter struct {
	ProductID                string // 产品 ID (必填)
	StateID                  string // 需求状态 ID (可选)
	PriorityID               string // 需求优先级 ID (可选)
	Keywords                 string // 搜索关键字 (可选)
	CreatedBetween           string // 创建时间范围，逗号分隔 (可选)
	UpdatedBetween           string // 更新时间范围，逗号分隔 (可选)
	IncludePublicImageToken  string // 包含图片资源 token (可选)
	PageSize                 int    // 每页数量 (可选)
	PageIndex                int    // 页码 (可选)
}

// Idea 需求
type Idea struct {
	ID         string      `json:"id"`
	URL        string      `json:"url"`
	Identifier string      `json:"identifier"`
	Title      string      `json:"title"`
	ShortID    string      `json:"short_id"`
	HTMLURL    string      `json:"html_url"`
	Product    interface{} `json:"product"`
	Assignee   interface{} `json:"assignee"`
	State      interface{} `json:"state"`
	Priority   interface{} `json:"priority"`
	Plan       interface{} `json:"plan"`
	Suite      interface{} `json:"suite"`
	PlanAt     interface{} `json:"plan_at"`
	RealAt     interface{} `json:"real_at"`
	Score      float64     `json:"score"`
	Progress   float64     `json:"progress"`
	Desc       string      `json:"description"`
	Properties interface{} `json:"properties"`
	CreatedAt  int64       `json:"created_at"`
	UpdatedAt  int64       `json:"updated_at"`
	IsArchived int         `json:"is_archived"`
	IsDeleted  int         `json:"is_deleted"`
}

// IdeasResponse 需求列表 API 响应
type IdeasResponse struct {
	Values    []Idea `json:"values"`
	PageSize  int    `json:"page_size"`
	PageIndex int    `json:"page_index"`
	Total     int    `json:"total"`
}

// ListIdeas 获取需求列表
// 这是一个临时函数，用于演示如何直接调用 API
// 正式使用时建议实现完整的 Ship Service
func ListIdeas(client *sdk.Client, ctx context.Context, filter IdeaFilter) ([]Idea, *core.Pagination, error) {
	// 获取底层 HTTP 客户端
	// 注意: 这需要访问内部实现，建议在 SDK 中添加 Ship Service
	return nil, nil, fmt.Errorf("请使用 Ship Service API: client.Ship().ListIdeas()")
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

	// 示例：获取需求列表的参数
	filter := IdeaFilter{
		ProductID: "your_product_id", // 替换为实际的产品 ID
		PageSize:  50,
		PageIndex: 1,
		// 可选参数:
		// StateID:    "state_id",
		// PriorityID: "priority_id",
		// Keywords:   "搜索关键字",
	}

	fmt.Println("获取需求列表示例")
	fmt.Println("================")
	fmt.Printf("产品 ID: %s\n", filter.ProductID)
	fmt.Printf("每页数量: %d\n", filter.PageSize)
	fmt.Printf("页码: %d\n\n", filter.PageIndex)

	// 构建查询参数示例
	query := url.Values{}
	query.Set("product_id", filter.ProductID)
	query.Set("page_size", fmt.Sprintf("%d", filter.PageSize))
	query.Set("page_index", fmt.Sprintf("%d", filter.PageIndex))
	if filter.StateID != "" {
		query.Set("state_id", filter.StateID)
	}
	if filter.PriorityID != "" {
		query.Set("priority_id", filter.PriorityID)
	}
	if filter.Keywords != "" {
		query.Set("keywords", filter.Keywords)
	}

	fmt.Println("API 端点: GET /v1/ship/ideas")
	fmt.Println("查询参数:")
	for key, values := range query {
		fmt.Printf("  %s: %v\n", key, values)
	}

	// 尝试调用（会失败，因为 Ship Service 尚未实现）
	_, _, err = ListIdeas(client, ctx, filter)
	if err != nil {
		fmt.Printf("\n当前状态: %v\n", err)
		fmt.Println("\n建议操作:")
		fmt.Println("1. 等待 Ship Service SDK 实现完成")
		fmt.Println("2. 或参考此示例直接调用 HTTP API")
	}

	// 打印响应示例
	fmt.Println("\n响应示例:")
	exampleResponse := IdeasResponse{
		Values: []Idea{
			{
				ID:         "64b4d70ba368e6594360ea24",
				URL:        "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
				Identifier: "PRD-1",
				Title:      "用户登录功能",
				ShortID:    "AbCd1234",
				HTMLURL:    "https://your-company.pingcode.com/ship/ideas/AbCd1234",
				CreatedAt:  1234567890,
				UpdatedAt:  1234567890,
				IsArchived: 0,
				IsDeleted:  0,
			},
		},
		PageSize:  50,
		PageIndex: 1,
		Total:     1,
	}

	jsonBytes, _ := json.MarshalIndent(exampleResponse, "", "  ")
	fmt.Println(string(jsonBytes))

	// 打印使用说明
	fmt.Println("\n================")
	fmt.Println("使用说明:")
	fmt.Println("1. 设置环境变量:")
	fmt.Println("   export PINGCODE_BASE_URL=https://open.pingcode.com")
	fmt.Println("   export PINGCODE_CLIENT_ID=your_client_id")
	fmt.Println("   export PINGCODE_CLIENT_SECRET=your_client_secret")
	fmt.Println("2. 替换 product_id 为实际值")
	fmt.Println("3. 运行: go run main.go")
}

