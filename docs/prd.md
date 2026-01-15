# PingCode OpenAPI Golang SDK PRD

## 1. 产品背景与目标

### 1.1 背景

- PingCode 已提供较为完整的 OpenAPI 文档（参考 `docs/reference/openpingcode` 目录下的各模块：概览、认证、产品、项目、全局、测试、知识、DevOps 等）。
- 目前其他 Go 项目若要接入 PingCode：
  - 需要自行拼装 HTTP 请求、处理认证、错误码、重试等逻辑；
  - 容易出现重复实现、风格不统一、错误处理不规范等问题。
- 缺少一套**符合 Go 开发习惯、稳定、易用**的官方/准官方 SDK。

### 1.2 产品愿景

> 提供一套基于 PingCode OpenAPI 的 Golang SDK，屏蔽 HTTP 细节和认证流程，提供 type-safe、模块化的接口封装，让 Go 项目能像调用本地库一样使用 PingCode 能力。

### 1.3 目标

- 短期目标（V1）：
  - 支持核心模块的稳定访问：认证、项目、产品、测试、知识（最常用的几个）。
  - 提供统一的 Client 初始化方式、认证方式和错误处理规范。
  - 覆盖常见场景的示例代码，降低接入门槛。
- 中期目标（V2+）：
  - 覆盖 OpenAPI 文档中大部分/全部模块。
  - 增加自动生成部分代码的能力（基于 OpenAPI），减少维护成本。
  - 增强高阶能力：重试策略、限流、链路追踪（metrics/log/trace）等。

---

## 2. 使用人群与使用场景

### 2.1 目标用户

- 内部 Go 服务开发者：需要在微服务中访问 PingCode 数据或触发操作。
- 外部集成方：有较强 Go 能力，希望集成 PingCode 功能到自身产品中。
- 工具/脚本作者：通过 Go 编写 CLI、定时任务与 PingCode 交互。

### 2.2 典型使用场景

- 从项目管理系统中拉取项目信息、需求列表、测试结果等。
- 自动创建/更新任务、需求、测试用例，进行 CI/CD 集成。
- 将 PingCode 中的知识库内容与其他系统同步。
- 自定义报表/看板服务，从 PingCode 抽数做 BI 分析。

---

## 3. 产品范围（Scope）

### 3.1 范围内（Must Have）

- 基础能力：
  - Client 初始化与配置（BaseURL、自定义 HTTP Client、日志、超时等）。
  - 支持至少两种认证方式（具体以 OpenAPI 文档为准，例如：Token / OAuth2）。
  - 全局统一错误类型（比如 `type Error struct { Code, Message, Raw *http.Response }`）。
- 核心模块 SDK 封装（基于已有 OpenAPI 文档）：
  - 认证模块（auth）：认证相关接口（如获取 token / 刷新 token 等）。
  - 项目模块（project）：查询项目信息、列表、成员等。
  - 产品模块（ship）：需求、版本等操作。
  - 测试模块（testhub）：测试用例、测试计划、执行结果等。
  - 知识模块（wiki）：知识空间、文档获取等。
- 基础开发体验：
  - 统一的包结构和命名规范，符合 Go 通用习惯。
  - 清晰的错误处理模式，避免使用者再去解析原始 HTTP Response。
  - 样例代码：至少提供若干 end-to-end 示例（例如：创建项目、拉取测试结果等）。

### 3.2 可选范围（Nice to Have）

- 其它模块封装（global、devops 等）。
- 自动生成部分 Model 和 API Client 的工具（根据 OpenAPI）。
- 更高级别封装（例如：批量操作、分页迭代器、通用查询封装等）。

### 3.3 不在范围内（Out of Scope at V1）

- 多语言 SDK（本 PRD 仅面向 Golang）。
- 完整 GUI 配置工具（如图形化配置 Client 的桌面应用）。
- 跨进程分布式缓存/限流系统（只提供可扩展接口，具体实现由业务方负责）。

---

## 4. 功能设计

### 4.1 整体架构与包结构

#### 4.1.1 架构原则

- 单一职责：每个模块（service）只负责对应领域逻辑。
- 高内聚、低耦合：HTTP 底层和业务 service 分层，避免耦合。
- 无魔法：避免复杂反射和隐式行为，使用显式类型和清晰 API。

#### 4.1.2 典型包结构（建议）

- `openpingcode/`
  - `client.go`：核心 Client 类型与初始化逻辑。
  - `config.go`：配置结构体与默认配置。
  - `auth/`：认证相关 service、model。
  - `project/`
  - `ship/`
  - `testhub/`
  - `wiki/`
  - `internal/`
    - `httpclient/`：底层 HTTP 请求封装、重试、日志等。
    - `encoding/`：JSON 编解码、时间类型处理等。
    - `errors/`：内部错误定义与转换。
  - `examples/`：示例代码。

> 注意：具体结构应与当前仓库的 Go 模块组织方式保持一致，以避免破坏现有使用方式。

### 4.2 Client 与配置

#### 4.2.1 Client 初始化

- 目标：
  - 提供简单、可读的初始化方式。
  - 保留足够扩展点（自定义 HTTP Client、Logger、Middlewares 等）。

- 需求：
  - 支持通过函数 `NewClient(config Config) (*Client, error)` 创建 Client。
  - `Config` 中至少包括：
    - `BaseURL`：PingCode OpenAPI 网关地址。
    - `Auth` 配置：如 Token / ClientID / ClientSecret / RefreshToken 等。
    - `HTTPClient`：可选，自定义 HTTP Client（不传则使用默认）。
    - `Timeout`：默认超时。
  - 支持 Option 模式（可选），便于扩展：
    - `WithHTTPClient(...)`
    - `WithLogger(...)`
    - `WithRetryPolicy(...)` 等。

#### 4.2.2 认证

- 根据 `auth.md` 中的 OpenAPI 描述，支持至少一种稳定可用的认证方式。
- 提供认证中间层，自动在请求中注入 Token / 头部，不要求业务方重复设置。
- 若 Token 过期，支持：
  - 明确的错误返回，让使用方自行决定重试；
  - 或（V2 之后）自动刷新机制，可通过配置启用。

### 4.3 模块 Service 设计

以 `Project` 模块为例：

- 提供接口风格：
  - `client.Project.Get(ctx, projectID)`  
  - `client.Project.List(ctx, filter)`  
- `filter` 类型使用强类型结构体，而不是 `map[string]interface{}`，保证类型安全。
- 返回值为结构体 + 错误：
  - `(*Project, error)` 或 `(ListProjectsResponse, error)`。
- 对分页接口：
  - V1：直接暴露 API 的分页参数与响应结构；
  - V2：可以提供分页迭代器 `ProjectIterator`，简化调用。

其它模块（Ship、TestHub、Wiki 等）按类似模式设计，所有方法都围绕以下原则：

- 使用上下文 `context.Context`，支持取消和超时控制。
- 不在 SDK 中引入业务强绑定逻辑，只做 API 级封装。

### 4.4 错误处理与返回值

- 定义统一错误类型，例如：

  - `type APIError struct { StatusCode int; Code string; Message string; RawBody []byte }`

- 所有 HTTP 请求失败（非 2xx）返回 `APIError` 类型。
- JSON 解析错误、网络错误等返回标准 `error`，但可通过 `errors.Is/As` 判断。
- 保证调用者可以通过错误类型判断是：
  - 请求参数问题；
  - 认证问题；
  - 服务器内部错误；
  - 网络/解析错误。

### 4.5 日志与调试

- V1：不强依赖具体日志框架，提供简单 Hook：
  - 在 Config/Option 中可以注入 `Logger` 接口。
  - SDK 内部在关键行为（如请求失败）调用 Logger。
- 兼容无日志场景：默认 Logger 为空实现，不影响业务。

---

## 5. 非功能需求

### 5.1 性能

- 针对典型调用场景，SDK 本身不能成为明显瓶颈：
  - 使用高效 JSON 库（默认可使用标准库，若项目已有统一选择则遵从）。
  - 连接复用交由 `http.Client` 控制，避免每次创建新 Client。
- 避免过度分配和拷贝大对象。

### 5.2 稳定性与向后兼容（Never break userspace）

- SDK 版本需要与 OpenAPI 版本明确对应（例如语义化版本号）。
- 对外公开 API 一旦发布：
  - 禁止破坏性变更（签名变更、行为改变）；
  - 如必须变更，采用新增 API + 标记旧 API deprecated 的方式。
- 所有 breaking change 必须：
  - 明确在 release note 中说明；
  - 给出迁移路径。

### 5.3 安全性

- 不在日志中打印敏感信息（Token、密码、密钥等）。
- 合理设置默认超时，避免调用挂死。
- 提供 TLS/HTTPS 支持（默认必须是 HTTPS）。

---

## 6. 版本规划与里程碑

### 6.1 V0.1（PoC）

- 完成基础 Client 结构和 1~2 个模块的最小实现（例如：auth + project）。
- 能在内部 demo 服务中跑通真实调用流程。
- 提供基础示例代码和简单 README。

### 6.2 V1.0（可用于生产）

- 覆盖至少：auth、project、ship、testhub、wiki 核心接口。
- 对所有暴露 API 写单元测试 + 集成测试（在允许范围内）。
- 接入现有 CI（lint、test、build）并全部通过。
- 文档完善：
  - 使用说明；
  - 典型示例；
  - 版本兼容策略说明。

### 6.3 后续版本（V1.x+）

- 覆盖剩余模块（global、devops 等）。
- 引入生成工具，自动同步 OpenAPI 更新。
- 增强健壮性（重试策略、限流、metrics 等）。

---

## 7. 测试与验收标准

### 7.1 测试范围

- 单元测试：
  - Config、Client 初始化；
  - 各模块 Service 的输入校验、错误处理。
- 集成测试：
  - 在测试环境调用真实 PingCode OpenAPI（可通过配置开关控制）。
  - 验证主要 API 调用路径正确工作。

### 7.2 验收标准

- 所有单元测试、集成测试、lint、typecheck 全部通过。
- 使用 SDK 的示例应用可以稳定完成以下操作：
  - 使用配置初始化 Client；
  - 通过认证；
  - 拉取项目列表；
  - 对至少一个模块完成创建/查询/更新的闭环操作。
- 对外 API 接口 freeze，并在文档中记录。

---

## 8. 文档与示例

### 8.1 SDK 文档

- README：
  - SDK 简介；
  - 安装方式；
  - 快速开始示例。
- 参考文档：
  - 每个模块的 API 列表与调用示例；
  - 错误类型说明；
  - 配置项说明。

> 注意：修改或新增影响用户使用的行为时，必须同步更新相关文档，特别是初始化方式、认证方式、错误处理策略等。

### 8.2 示例代码

- `examples/basic_usage`：最小可运行例子：
  - 读取环境变量中的 Token；
  - 初始化 Client；
  - 拉取项目列表并打印。
- `examples/ci_integration`：展示如何在 CI 中使用 SDK 获取测试结果/构建状态等。

---

## 9. 风险与对策

### 9.1 OpenAPI 变更带来的破坏风险

- 对策：
  - 建立 OpenAPI 版本与 SDK 版本映射；
  - 增加简易的“变更检测”（例如在生成阶段检查 diff）。

### 9.2 过度封装导致的调试困难

- 对策：
  - 始终保留查看原始请求/响应信息的能力（受安全控制）；
  - 错误类型中包含必要的原始信息（状态码、部分 body）。
