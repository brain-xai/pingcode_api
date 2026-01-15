# Phase 1 迭代规划

目标：在现有文档体系和技术规范（`prd.md`、`deployment.md`、`api_contract.md`、`error_model.md`、`openapi_integration.md`）的基础上，完成首个可用版本的最小闭环能力：

1. 基础技术框架搭建（模块骨架与 Client）
2. 用户认证能力
3. 获取项目列表能力
4. 提供一个可运行的示例，演示“获取项目列表”

本阶段不追求功能面广度，而是验证整体架构与规范在真实代码中的可行性。

---

## 1. 范围与不在范围

### 1.1 本阶段范围

- 实现最小可用 SDK 框架：
  - `sdk/client`
  - `sdk/config`
  - 基础错误模型落地
- 支持一种稳定可用的认证方式（基于当前 PingCode OpenAPI 支持的方式）：
  - 例如：Token 认证（以实际 API 文档为准）
- 实现 Project 模块的“获取项目列表”能力：
  - 包括 `sdk/model/project` 的核心结构体
  - `sdk/service/project` 的列表接口
  - 对应的 `internal/api/project` DTO 与映射
- 提供一个可以 `go run` 的示例：
  - 从环境变量读取配置和 Token
  - 初始化 Client
  - 调用“项目列表”接口并输出结果

### 1.2 不在本阶段范围

- 其它领域（测试、知识、DevOps 等）的业务能力
- 自动生成工具的完备实现（只做最小可行方案）
- 高级能力：自动重试策略、限流、metrics、链路追踪
- 多种认证方式的并行支持（先选定一种跑通）

---

## 2. 基础技术框架

### 2.1 SDK 顶层骨架创建

**目标：** 按 `deployment.md` 规范，创建最小目录与文件结构，使项目可编译。

需要完成：

- 目录与文件：
  - `sdk/client.go`
  - `sdk/config.go`
  - `sdk/errors.go`
  - `sdk/model/core/`
  - `sdk/model/project/`
  - `sdk/service/project/`
  - `internal/httpclient/`
  - `internal/api/project/`（先可手写最小 DTO，后续再对接生成）
- 基本内容：
  - `Config` 结构体骨架（仅包含 Phase1 必需字段）
  - `Client` 结构体骨架及 `NewClient` 函数
  - `Client` 内挂载 Project Service 的入口（例如 `Project()` 方法）
  - 基础 HTTP Client 包装（`internal/httpclient`）：
    - 持有 `*http.Client`
    - 支持设置 BaseURL 和超时
    - 预留添加认证头、中间件的 Hook

验收标准：

- `go build ./...` 能通过
- 目录结构符合 `deployment.md` 的 2.x 章节要求

### 2.2 错误模型初始实现

**目标：** 基于 `error_model.md`，实现统一错误类型的最小集合。

需要完成：

- 在 `sdk/errors.go` 或 `sdk/model/core` 中：
  - 定义统一错误类型（例如 `APIError`）：
    - 包含 HTTP 状态码、业务错误码、消息、请求 ID 等字段（可先用最小子集）
  - 定义若干语义性错误值（可从 Phase1 需要的开始）：
    - `ErrInvalidConfig`
    - `ErrUnauthorized`（若认证失败时使用）
- 实现与 `internal/httpclient` 的连接：
  - 在 HTTP 请求失败时构造 `APIError`
  - 保证后续可通过 `errors.As` 获取 `APIError`

验收标准：

- Project 列表接口出现错误时，能够返回统一的错误类型
- 错误类型与字段命名符合 `deployment.md` 命名规范

---

## 3. 用户认证实现

### 3.1 认证配置与加载

**目标：** 确定 Phase1 使用的认证方式，并在 Config 中加入相应字段。

需要完成：

- 在 `sdk/config.go` 中：
  - 为当前选定的认证方式增加配置字段（例如：`Token` 或 `ClientID`/`ClientSecret`）
  - 定义 `AuthConfig` 结构体（如果认证信息较复杂）
- 提供基础配置加载辅助函数（可选）：
  - 如 `LoadConfigFromEnv()`：
    - 读取环境变量：`PINGCODE_BASE_URL`、`PINGCODE_ACCESS_TOKEN` 等
    - 构造 `Config`，在缺失必需字段时返回 `ErrInvalidConfig`

验收标准：

- 可在示例中仅通过设置环境变量+少量代码就完成初始配置
- 不泄露 Token 等敏感信息到日志

### 3.2 认证中间层与 HTTP 请求集成

**目标：** 在所有 HTTP 请求中自动注入认证信息。

需要完成：

- 在 `internal/httpclient` 或独立位置：
  - 实现一个负责在请求头中注入认证信息的中间层：
    - 例如：在每个请求加上 `Authorization: Bearer <token>` 或其他约定头
- 在 `Client` 初始化时：
  - 根据 `Config` 检查认证配置是否完整
  - 初始化 HTTP 客户端并携带认证配置
- 暂不实现自动刷新 Token（如不在 Phase1 目标中），但需要：
  - 清晰地在错误路径中返回认证失败错误

验收标准：

- 调用 Project 列表接口时，认证信息自动附加，无需调用方手工设置头部
- 认证失败时，返回可识别的错误（如 `ErrUnauthorized` 或 `APIError` 中的特定码）

---

## 4. 获取项目列表能力

### 4.1 对外模型（sdk/model/project）

**目标：** 定义对使用者友好的 Project 列表模型，不暴露底层 DTO。

需要完成：

- 在 `sdk/model/project` 中定义：
  - `Project` 基本结构体（Phase1 只需核心字段，比如 ID、Name、Key、Status）
  - `ProjectFilter`（如暂不复杂，可先定义最小过滤条件，甚至可为空）
  - 如需分页，则从 `sdk/model/core` 引用 `Pagination` 等类型
- 确保字段命名与含义稳定、易懂，符合命名规范

验收标准：

- 示例代码中对 `Project` 类型的使用自然、直观
- 不出现 `map[string]interface{}` 这类不透明类型

### 4.2 internal/api DTO（internal/api/project）

**目标：** 定义与实际 OpenAPI 文档对应的 Project 列表请求/响应 DTO。

需要完成：

- 在 `internal/api/project` 中：
  - 定义请求 DTO（如有需要，包含分页参数、过滤条件）
  - 定义响应 DTO，完整映射 OpenAPI 中的字段（可在 Phase1 只实现正在使用的子集）
- DTO 的 JSON tag 与字段命名尽量贴近 OpenAPI 文档

验收标准：

- DTO 能够与 PingCode 实际返回的 JSON 正确编解码
- 不直接被 `sdk` 对外导出

### 4.3 Service 方法实现（sdk/service/project）

**目标：** 提供一个对外可用的“获取项目列表”方法。

需要完成：

- 在 `sdk/service/project` 中：
  - 定义 `Service` 结构体，持有内部 HTTP Client
  - 实现方法，如：
    - `List(ctx context.Context, filter model.ProjectFilter) ([]model.Project, *core.Pagination, error)`
- 方法内部逻辑：
  - 将 `ProjectFilter` 转换为 `internal/api` 请求 DTO
  - 使用 `internal/httpclient` 发送请求
  - 将响应 DTO 映射为 `[]model.Project` 和分页信息

验收标准：

- 通过示例程序可成功调用 `List` 并获得项目列表
- 错误路径走统一错误模型

---

## 5. 示例：获取项目列表

### 5.1 示例目录与文件

**目标：** 提供一个可直接运行的例子，演示如何使用 SDK 获取项目列表。

需要完成：

- 新目录与文件：
  - `examples/basic_usage/main.go`（或等价路径）
- 示例内容：
  - 从环境变量读取：
    - `PINGCODE_BASE_URL`
    - `PINGCODE_ACCESS_TOKEN`（或其他认证信息）
  - 调用 SDK：
    - 构造 `Config`（或使用 `LoadConfigFromEnv`）
    - `NewClient(config)`
    - `client.Project().List(ctx, filter)`
  - 打印项目列表的关键信息（ID、Name）

验收标准：

- 在正确配置环境变量后，运行：
  - `go run ./examples/basic_usage`
- 能在终端输出项目列表或清晰的错误信息

---

## 6. 测试与质量保障

### 6.1 单元测试

需要完成：

- 对以下内容编写单元测试（至少覆盖核心逻辑）：
  - `Config` 校验逻辑（包括认证配置）
  - `Project` Service 的参数处理和映射逻辑（不依赖真实 HTTP 服务，可使用 mock）
  - 错误模型的基本行为（`errors.As` 等）

验收标准：

- `go test ./...` 通过
- 关键路径（Client 初始化 + Project 列表）具备基本单测

### 6.2 集成测试（可选）

如有测试环境，可在后续子阶段补充：

- 针对真实 PingCode 测试环境编写集成测试：
  - 实际发起“获取项目列表”请求
  - 校验成功路径与常见错误路径

---

## 7. 文档与对齐

需要同步更新的文档与说明：

- 在 `README` 或合适位置：
  - 增加“快速开始：获取项目列表”的简短说明，指向示例代码
- 如有必要，在 `deployment.md` 中补充：
  - Phase1 中已落地的部分（例如具体认证方式的现状说明）
- 确保 `api_contract.md` 和 `error_model.md` 中对于本阶段实现的部分都是一致的：
  - 对外暴露的 API（Client、Project Service）满足契约定义
  - 错误模型行为与文档一致

---

## 8. Phase 1 验收标准总结

本次迭代完成后，至少要满足：

- 从干净环境开始，执行：
  - `go build ./...` 成功
  - `go test ./...` 成功
- 在设置正确环境变量后：
  - `go run ./examples/basic_usage` 可以成功调用 PingCode OpenAPI，获取并打印项目列表
- 对外暴露的 API：
  - 包结构符合 `deployment.md`
  - 行为与 `api_contract.md`、`error_model.md` 一致
- 所有变更均不违反 "Never break userspace" 原则（本阶段主要是新增，不应有破坏性改动）

---

## 9. 实现状态

### 已完成 ✅

- **基础技术框架**
  - ✅ `sdk/client.go` - Client 结构体和 NewClient 函数
  - ✅ `sdk/config.go` - Config 结构体，支持 client_id/client_secret 和 access_token
  - ✅ `sdk/errors.go` - 统一错误类型
  - ✅ `sdk/model/core/pagination.go` - 分页模型

- **认证模块**
  - ✅ `sdk/model/auth/token.go` - 认证令牌模型
  - ✅ `internal/api/auth/auth.go` - 认证 API DTO
  - ✅ `sdk/service/auth/auth.go` - 认证服务（获取 token）
  - ✅ 支持客户端凭据（Client Credentials）授权方式
  - ✅ Token 过期检测（提前 5 分钟判定为过期）

- **Project 模块**
  - ✅ `sdk/model/project/project.go` - 项目模型（核心字段）
  - ✅ `internal/api/project/project.go` - 项目 API DTO
  - ✅ `sdk/service/project/project.go` - 项目服务（列表接口）

- **HTTP 客户端**
  - ✅ `internal/httpclient/client.go` - HTTP 客户端封装
  - ✅ 自动添加 Authorization 头（除获取 token 请求外）
  - ✅ 统一错误处理

- **示例代码**
  - ✅ `examples/basic_usage/main.go` - 完整的使用示例
  - ✅ 支持从环境变量加载配置

- **文档**
  - ✅ `README.md` - 项目说明文档

### 编译与测试状态

- ✅ `go build ./...` 通过
- ✅ 所有包结构符合 `deployment.md` 规范
- ⏸️ 单元测试（可在后续补充）

### 待运行验证

需要配置真实的环境变量来验证实际 API 调用：

```bash
export PINGCODE_BASE_URL=https://open.pingcode.com
export PINGCODE_CLIENT_ID=your_client_id
export PINGCODE_CLIENT_SECRET=your_client_secret
go run ./examples/basic_usage
```