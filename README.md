# PingCode OpenAPI Golang SDK

PingCode OpenAPI 的 Go 语言 SDK，提供类型安全、模块化的接口封装。

## 快速开始

### 安装

```bash
go get github.com/brain-xai/pingcode_api
```

### 获取项目列表

```go
package main

import (
	"context"
	"log"

	"github.com/brain-xai/pingcode_api/sdk"
)

func main() {
	// 方式1：使用环境变量（推荐）
	config, err := sdk.LoadConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// 方式2：手动配置
	// config := &sdk.Config{
	// 	BaseURL: "https://open.pingcode.com",
	// 	Auth: &sdk.AuthConfig{
	// 		ClientID:     "your_client_id",
	// 		ClientSecret: "your_client_secret",
	// 	},
	// }

	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// SDK 会自动获取 access_token
	projects, _, err := client.Project().List(context.Background(), sdk.ProjectFilter{})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range projects {
		log.Printf("%s: %s", p.Identifier, p.Name)
	}
}
```

### 环境变量

| 环境变量 | 说明 | 默认值 |
|---------|------|--------|
| `PINGCODE_BASE_URL` | API 基础 URL | `https://open.pingcode.com` |
| `PINGCODE_CLIENT_ID` | PingCode 应用的 Client ID | - |
| `PINGCODE_CLIENT_SECRET` | PingCode 应用的 Secret | - |
| `PINGCODE_ACCESS_TOKEN` | 直接提供访问令牌（可选） | - |

## 认证方式

SDK 支持两种认证配置方式：

### 1. 客户端凭据（推荐）

使用 Client ID 和 Client Secret，SDK 会自动获取访问令牌：

```go
config := &sdk.Config{
    BaseURL: "https://open.pingcode.com",
    Auth: &sdk.AuthConfig{
        ClientID:     "your_client_id",
        ClientSecret: "your_client_secret",
    },
}
client, _ := sdk.NewClient(config)
```

### 2. 直接提供访问令牌

如果您已经有访问令牌，可以直接使用：

```go
config := &sdk.Config{
    BaseURL: "https://open.pingcode.com",
    Auth: &sdk.AuthConfig{
        AccessToken: "your_access_token",
    },
}
client, _ := sdk.NewClient(config)
```

## 示例

查看 `examples/basic_usage` 获取完整示例。

```bash
cd examples/basic_usage
export PINGCODE_BASE_URL=https://open.pingcode.com
export PINGCODE_CLIENT_ID=your_client_id
export PINGCODE_CLIENT_SECRET=your_client_secret
go run main.go
```

### 获取项目详情

```go
project, err := client.Project().Get(ctx, "project-id")
if err != nil {
    log.Fatal(err)
}
log.Printf("项目: %s (%s)", project.Name, project.Identifier)
if project.Assignee != nil {
    log.Printf("负责人: %s", project.Assignee.DisplayName)
}
```

### 获取当前用户信息

```go
user, err := client.Global().GetCurrentUser(ctx)
if err != nil {
    log.Fatal(err)
}
log.Printf("当前用户: %s (%s)", user.DisplayName, user.Email)
```

### 获取企业成员列表

```go
users, _, err := client.Global().ListUsers(ctx, global.UserFilter{
    Keywords: "john",
})
for _, user := range users {
    log.Printf("- %s (%s)", user.DisplayName, user.Email)
}
```

### 获取项目成员列表

```go
members, _, err := client.Project().ListMembers(ctx, projectID)
for _, member := range members {
    if member.Type == "user" && member.User != nil {
        log.Printf("- 用户: %s", member.User.DisplayName)
    } else if member.Type == "user_group" && member.UserGroup != nil {
        log.Printf("- 团队: %s", member.UserGroup.Name)
    }
}
```

## API 端点

| 端点 | 方法 | 说明 |
|------|------|------|
| `/v1/auth/token` | GET | 获取访问令牌（不需要认证） |
| `/v1/project/projects` | GET | 获取项目列表 |
| `/v1/project/projects/{project_id}` | GET | 获取项目详情 |
| `/v1/project/projects/{project_id}/members` | GET | 获取项目成员列表 |
| `/v1/project/project/states` | GET | 获取项目状态列表 |
| `/v1/project/projects/{project_id}/progress` | GET | 获取项目进度 |
| `/v1/myself` | GET | 获取当前用户信息 |
| `/v1/directory/users/{user_id}` | GET | 获取用户详情 |
| `/v1/directory/users` | GET | 获取企业成员列表 |

## 文档

- [产品需求文档](docs/prd.md)
- [技术规范](docs/deployment.md)
- [API 契约规范](docs/api_contract.md)
- [错误模型](docs/error_model.md)
- [OpenAPI 集成](docs/openapi_integration.md)
- [PingCode API 参考](docs/reference/openpingcode/Readme.md)

## License

[待定]
