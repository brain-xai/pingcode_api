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

## 功能清单

### 核心能力
- ✅ **认证管理** - 支持 Client Credentials 和 Access Token 两种认证方式
- ✅ **自动 Token 管理** - 自动获取和刷新访问令牌
- ✅ **统一错误处理** - 类型安全的错误模型和错误处理
- ✅ **分页支持** - 内置分页查询能力

### 已实现领域

#### Project - 项目管理
- 获取项目列表、项目详情
- 创建、更新、删除项目
- 获取项目成员、状态、进度

#### Ship - 产品与需求管理
- 产品管理：获取产品列表、产品详情
- 需求管理：获取需求列表、需求详情、创建需求、更新需求
- 需求辅助接口（UI场景）：状态、优先级、属性、模块、排期

#### WorkItem - 工作项管理
- 基础操作：创建、更新、批量更新、查询、删除工作项
- 属性与分类：类型、状态、属性、优先级、标签
- 标签管理：添加、移除标签
- 关联管理：创建、查询、取消工作项关联
- 交付目标：创建、更新、查询、删除交付目标
- 流转记录：查询工作项流转历史

#### Global - 全局服务
- 获取当前用户信息
- 获取用户详情、企业成员列表

## 示例代码

项目提供完整的示例代码，位于 `examples/` 目录：

| 示例目录 | 功能说明 | 主要演示 |
|---------|---------|---------|
| `basic_usage` | 最小可运行示例 | 获取项目列表、环境变量配置 |
| `project_overview` | 项目管理整体验证 | 项目列表、详情、成员、进度查询 |
| `ship_products` | 产品管理 | 获取产品列表和详情 |
| `ship_requirements` | 需求管理 | 需求查询、创建、更新 |
| `list_ideas` | 创意列表查询 | 获取产品创意列表 |
| `workitems` | 工作项管理 | 创建、更新、查询、关联工作项 |

### 快速运行示例

所有示例均使用环境变量进行配置：

```bash
cd examples/basic_usage  # 可替换为其他示例目录
export PINGCODE_BASE_URL=https://open.pingcode.com
export PINGCODE_CLIENT_ID=your_client_id
export PINGCODE_CLIENT_SECRET=your_client_secret
go run main.go
```

## 兼容性与版本策略

### 语义化版本

本项目遵循 [语义化版本 (SemVer)](https://semver.org/lang/zh-CN/) 规范：

- **MAJOR** (主版本号) - 不兼容的 API 变更
- **MINOR** (次版本号) - 向后兼容的功能新增
- **PATCH** (修订号) - 向后兼容的问题修复

### API 兼容承诺

我们对 `sdk/` 包下的公开 API 做以下承诺：

1. **稳定性保证** - 一旦某个导出类型、方法进入已发布版本，即视为对外契约
2. **静默变更禁止** - 不会在 MINOR/PATCH 版本中无声改变已发布 API 的语义或行为
3. **废弃流程** - 需要移除某个 API 时，会：
   - 标记为 `Deprecated`
   - 在文档中说明替代方案
   - 保留至少一个 MINOR 版本周期
   - 在 CHANGELOG 中记录

### 破坏性变更处理

所有破坏性变更仅在新 MAJOR 版本中进行，且必须：

- 在 CHANGELOG.md 中明确记录
- 提供详细的迁移指南
- 通过合理的版本号升级路径

详细规则请参考 [API 契约规范](docs/api_contract.md)。

### 当前版本

当前版本处于 **0.x** 阶段，以快速迭代为主，但依然遵守基本的兼容性约束。

- 公开 API 一旦发布，尽量避免修改
- 必须修改时，会在文档中明确说明并提供迁移路径

## 文档

### 核心文档
- [变更日志 (CHANGELOG)](CHANGELOG.md) - 版本变更历史
- [产品需求文档 (PRD)](docs/prd.md) - 产品背景与目标
- [技术规范](docs/deployment.md) - 架构设计、包结构、命名规范
- [API 契约规范](docs/api_contract.md) - 对外 API 列表和兼容性规则
- [错误模型](docs/error_model.md) - 错误类型设计规范
- [OpenAPI 集成](docs/openapi_integration.md) - OpenAPI 集成策略

### 使用指南
- [工作项管理指南](docs/workitem_management.md) - 工作项完整使用指南

### API 参考
- [PingCode OpenAPI 参考](docs/reference/openpingcode/Readme.md) - 完整的 OpenAPI 文档

## License

本项目采用 [Apache License 2.0](LICENSE) 开源许可。

### 许可概要

- ✅ **商用使用** - 可以用于商业目的
- ✅ **修改** - 可以修改源代码
- ✅ **分发** - 可以分发原版或修改版
- ✅ **私用** - 可以私下使用和修改
- ⚠️ **责任** - 软件按"原样"提供，不提供任何担保
- 📄 **要求** - 必须保留版权和许可证声明

详细条款请参见 [LICENSE](LICENSE) 文件。
