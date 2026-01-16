# PingCode OpenAPI Golang SDK Deployment 技术规范

> 目标：定义 SDK 在开发、构建、发布、兼容性与运维层面的统一技术规范，避免“每个服务各搞一套”，同时确保 **Never break userspace**。

模块路径：`github.com/brain-xai/pingcode_api`  
语言版本：`go 1.24.x`（保持与 go.mod 一致）

---

## 1. 总体设计原则

1. **不破坏用户空间**
   - 任何改动不得无声地改变已有公开 API 的语义或行为。
   - 若必须做破坏性变更，必须通过新版本号、Deprecated 标记和迁移文档显式暴露。

2. **简单优先**
   - 构建与发布流程保持最少步骤，避免引入多套脚本体系。
   - 不依赖复杂的生成工具链作为**必需条件**，生成只作为加速开发手段。

3. **可重复构建**
   - 同一 commit 在任何环境（本地/CI）执行规范化的构建命令，输出结果应一致。
   - 所有构建参数、环境变量有明确文档，禁止“只存在于某个人脑子里”。

4. **清晰分层**
   - 区分 SDK 源码、生成代码、示例、测试和发布工件。
   - 内部工具与对外 API 严格隔离（internal/）。

---

## 2. 目录、包结构与数据模型规范

> 本章节是整个 SDK 技术实现的“骨架规范”，定义包结构和数据结构的唯一来源。任何实现都必须严格遵守本节约定。

### 2.1 顶层模块与目录

- 根模块：`github.com/brain-xai/pingcode_api`
  - 不允许再创建第二个 Go module，避免多模块构建混乱。

- 顶层目录：
  - `sdk/`：对外暴露的 SDK 入口与各 Service 实现
    - `sdk/client.go`：Client 定义与初始化逻辑
    - `sdk/config.go`：Config 定义与默认值
    - `sdk/errors.go`：SDK 统一错误类型和映射
    - `sdk/model/`：所有对外可见的业务数据结构（统一数据结构树）
    - `sdk/service/`：各领域的 Service（API 调用逻辑）
  - `internal/`：仅供 SDK 内部使用的工具和 API 映射层
    - `internal/api/`：与 OpenAPI 一一对应的原始 DTO（通常由工具生成）
    - `internal/httpclient/`：HTTP 封装、重试、日志、metrics Hook
    - `internal/encoding/`：JSON 编解码、时间格式等
    - `internal/errors/`：内部错误定义与转换
  - `cmd/`（可选）：与 SDK 相关的命令行工具，如 SDK 生成器或调试工具
  - `examples/`：示例代码，必须可编译运行

目录约束：

- `internal/` 下的任何包不得被 SDK 使用者直接引用。
- 对外导出的类型与函数只能位于：
  - `sdk/` 及其子目录
  - `cmd/` 下仅用于可执行程序，不对外形成稳定 API。

### 2.2 sdk/model：统一对外数据结构树

- 所有对 SDK 使用者可见的业务数据结构必须放在 `sdk/model` 目录树下，禁止在 service 或其它目录随意新增对外结构体。
- 统一结构：

  - `sdk/model/core/`：跨领域复用的基础类型
    - 各种 ID 类型：`ProjectID`、`IssueID`、`UserID` 等
    - 分页结构：`Pagination`、`PageRequest` 等
    - 通用时间、状态和枚举类型

  - `sdk/model/auth/`：认证领域模型
  - `sdk/model/project/`：项目领域模型
  - `sdk/model/ship/`：产品/需求领域模型
  - `sdk/model/testhub/`：测试领域模型
  - `sdk/model/wiki/`：知识库领域模型
  - `sdk/model/devops/`：DevOps 领域模型
  - `sdk/model/global/`：全局配置与跨模块通用模型

约束：

- 只有真正跨多个模块的基础类型才放入 `sdk/model/core`，避免成为“垃圾场”。
- 一旦某个 model 类型被用于对外 API 的入参或返回值，即视为公共 API 的一部分，需遵守语义化版本管理与向后兼容规则（见第 6 章）。

### 2.3 sdk/service：按领域划分的 Service 层

- 结构：

  - `sdk/service/auth/`
  - `sdk/service/project/`
  - `sdk/service/ship/`
  - `sdk/service/testhub/`
  - `sdk/service/wiki/`
  - `sdk/service/devops/`
  - `sdk/service/global/`

- 规范：

  - Service 层只负责“如何调用 API”：HTTP 方法、URL 路径、分页与重试策略等。
  - Service 对外暴露的入参与返回值类型必须全部来自 `sdk/model/...`，禁止直接暴露 `internal/api` 类型。
  - Service 中不得使用 `map[string]interface{}` 作为对外请求参数类型，必须使用明确的 Filter/Input 结构体，这些结构体放在对应的 `sdk/model` 包中。
  - Service 方法必须使用 `context.Context` 控制生命周期，禁止无 context 的网络调用。

### 2.4 internal/api：OpenAPI 映射层

- 结构：

  - `internal/api/auth/`
  - `internal/api/project/`
  - `internal/api/ship/`
  - `internal/api/testhub/`
  - `internal/api/wiki/`
  - `internal/api/devops/`
  - `internal/api/global/`

- 职责：

  - 与 PingCode OpenAPI 文档保持尽可能一一对应的 DTO 定义：字段名、JSON 标签等紧跟文档。
  - 可以通过代码生成工具从 OpenAPI 描述生成，禁止手动调整生成后的代码逻辑（仅允许在模板或上游描述中修改）。

- 约束：

  - `internal/api` 仅供 SDK 内部使用，绝不能从 `sdk/` 包直接导出或被使用方依赖。
  - 当 OpenAPI 文档发生变更时，优先修改 `internal/api` 生成逻辑和 model/api 的映射关系，不直接对外暴力修改 `sdk/model` 结构，从而减少对使用者的破坏。

### 2.5 model 与 internal/api 的映射规范

- 所有对外暴露的 `sdk/model` 类型与 `internal/api` 类型之间必须有清晰的转换关系：

  - 输入（调用方 → SDK）：`sdk/model` → `internal/api` → HTTP JSON
  - 输出（服务端 → SDK）：HTTP JSON → `internal/api` → `sdk/model`

- 转换逻辑：

  - 建议在各 `sdk/service` 包内或单独的 `mapper` 文件中实现，禁止在随机位置散落转换代码。
  - 禁止在转换过程中静默吞掉字段或错误，所有丢弃/默认行为必须在代码中可见并可审计。

该映射规范是保证“OpenAPI 可以演进，而 SDK 使用者不被轻易破坏”的关键机制。

---

## 3. 构建与编译规范

### 3.1 本地构建命令

统一使用 Go 工具链：

- 构建 SDK 相关可执行工具（如有）：

```bash
go build ./cmd/...
```

- 验证 SDK 编译通过（不构建可执行）：

```bash
go build ./...
```

要求：

- 所有提交到主干分支的代码必须保证 `go build ./...` 能通过。
- 不引入非必要的 Makefile 复杂封装，如果有 Makefile：
  - `make build` 等价或简单包装 `go build ./...`。
  - Makefile 中不得引入额外语言的构建依赖作为“强依赖”（例如必须安装 Node 才能编译 SDK 是垃圾设计）。

### 3.2 生成代码（如有）

若存在根据 OpenAPI 生成的代码（model、client 等）：

- 生成代码必须写入 `internal/gen/` 或 `sdk/gen/` 之类易识别目录。
- 生成命令示例：

```bash
go run ./cmd/openapi-gen --input docs/reference/openpingcode/overview.md --output internal/gen
```

规范：

- 生成工具必须是 go 程序或标准工具，避免依赖私有二进制。
- 生成代码不得手动编辑，所有修改通过源定义（OpenAPI/模板）完成。
- CI 中必须有“防漂移”检查（可在后续 CI 配置中实现）：重新生成后不能产生 diff。

---

## 4. 测试与质量控制

### 4.1 测试命令

统一测试入口：

```bash
go test ./...
```

要求：

- 所有 SDK 对外导出 API 必须有单元测试或集成测试覆盖。
- 对外行为稳定的模块（如 auth、project 等）需重点保障：
  - 入参校验
  - 错误映射（HTTP -> SDK 错误类型）
  - 分页/重试等行为

### 4.2 静态检查与格式化

规范（若项目已有现成工具，则遵从；否则建议如下）：

- 格式化：使用 `gofmt` 或 `goimports`
- 静态检查（建议）：

```bash
go vet ./...
```

或者引入 `golangci-lint` 时：

```bash
golangci-lint run ./...
```

要求：

- 不允许引入和项目整体不兼容的 lint 规则（例如突然要求全局重命名大量符号）。
- 所有新代码必须通过已有 lint 配置，不在 PR 中引入“先合并再慢慢修”的技术债。

---

## 5. 配置与环境变量规范

### 5.1 SDK Config 结构

SDK 的配置结构 `Config` 是部署行为的核心入口，规范如下：

- 必须包含并明确文档化的字段：
  - `BaseURL`：PingCode OpenAPI 网关地址
  - `Auth`：认证配置（Token、ClientID/Secret 等）
  - `HTTPClient`：可选自定义 HTTP Client
  - `Timeout`：请求超时时间
- 必须提供安全的默认值策略：
  - 未显式设置 `Timeout`，采用合理默认，比如 10s。
  - 未设置 `BaseURL` 时立即返回配置错误，而不是随便拼一个假地址。

### 5.2 环境变量约定（建议）

为方便部署环境注入配置，约定以下环境变量（SDK 不强依赖，但提供辅助方法）：

- `PINGCODE_BASE_URL`
- `PINGCODE_ACCESS_TOKEN` 或 `PINGCODE_CLIENT_ID` / `PINGCODE_CLIENT_SECRET`
- `PINGCODE_REQUEST_TIMEOUT`

可提供辅助函数：

- `LoadConfigFromEnv() (Config, error)`  
  - 仅从环境变量构建 Config，不直接创建 Client。
  - 一旦解析失败返回明确错误，禁止“静默降级”。

---

## 6. 版本管理与发布规范

### 6.1 语义化版本

采用语义化版本号：`vMAJOR.MINOR.PATCH`，对应规则：

- **MAJOR**：存在破坏性变更（对公开 API 的签名或语义造成不兼容）。
- **MINOR**：向后兼容的新特性、新接口。
- **PATCH**：向后兼容的 bugfix 或内部实现优化。

约束：

- 不允许在 PATCH 版本中改变任何用户可观察到的行为（除非修复明确的 bug）。
- 不允许“偷偷把一个 breaking change 放在 MINOR 里”，这是对用户的背刺。

### 6.2 Tag 与发布流程（Git）

- 每次发布 SDK 版本必须打 Git tag，例如：`v1.2.3`。
- Tag 对应的 commit 必须为：
  - `go test ./...` 全量通过；
  - 构建通过；
  - 文档更新完毕（包括 README、示例、变更日志）。

发布步骤建议：

1. 更新版本号（若在代码中硬编码需要更新）。
2. 更新 CHANGELOG / 发布说明。
3. 确保 CI 全绿。
4. 打 tag：`git tag vX.Y.Z && git push origin vX.Y.Z`。

---

## 7. CI/CD 集成规范

> 这里规定的是行为，不绑死具体 CI 实现（GitHub Actions/GitLab CI 等）。

### 7.1 CI 阶段划分

最小 CI 流程包含以下阶段：

1. **Lint**（可选视项目现状而定）
   - `go vet ./...`
   - 或 `golangci-lint run ./...`
2. **Test**
   - `go test ./...`
3. **Build**
   - `go build ./...`（确保所有包可编译）

约束：

- PR 必须通过上述阶段才能合并到主干。
- 主干分支任何 commit 必须保持"随时可发布"的状态。

### 7.1.1 当前 CI 实现

项目使用 GitHub Actions 实现 CI，配置文件位于 `.github/workflows/ci.yml`。

**触发条件**：
- push 到 main 分支
- 针对 main 分支的 pull request

**检查步骤**：
1. `go mod download` - 下载依赖
2. `go build ./...` - 编译检查
3. `go test ./... -v -race` - 测试并检测竞态条件

**标准命令**：
开发者在本地应使用以下命令确保代码质量：
```bash
# 编译检查
go build ./...

# 运行测试
go test ./... -v

# 竞态检测
go test ./... -race
```

### 7.2 发布流水线（可选自动化）

- 在打 tag 时自动触发发布流程：
  - 重新执行 Lint/Test/Build，保证与合入时一致。
  - 若将来需要发布到私有包仓库 / Proxy，可在此阶段完成。

---

## 8. 兼容性与废弃策略

### 8.1 向后兼容策略

- 所有对外暴露的类型、函数、字段一旦发布：
  - 不得随意删除或更改含义。
- 如果某个 API 需要废弃：
  - 在代码中标记为 `Deprecated`（注释说明替代方案）。
  - 在至少一个 MINOR 版本周期内保留。
  - 在删除前必须更新文档和迁移指南。

### 8.2 OpenAPI 变更同步

当 PingCode OpenAPI 发生变更时：

1. 先更新生成工具/模型，不直接改动 SDK API。
2. 分析变更类型：
   - 仅新增字段/接口 → MINOR 版本。
   - 删除字段/接口或行为改变 → MAJOR 版本。
3. 对破坏性变更，提供适配层：
   - 通过新的方法/字段暴露新行为；
   - 保留旧接口一段时间，但在内部适配。

---

## 9. 运行时行为与观测性

### 9.1 日志

- SDK 内部不直接依赖具体日志库，使用接口抽象：
  - 如定义简单 `Logger` 接口：`Debugf/Infof/Warnf/Errorf`。
- 默认实现为 no-op，业务方可以按需注入。
- 不允许输出敏感信息（Token、密码等）。

### 9.2 超时与重试

- 所有请求必须受 `context.Context` 约束。
- 默认超时由 SDK 自己控制（通过 Config 或 HTTP Client），不得依赖业务方“刚好写对”。
- 重试策略（如实现）必须：
  - 明确哪些错误会触发重试（网络错误/5xx 等）。
  - 默认重试次数与退避策略可配置。
  - 禁止无限重试。

### 9.3 Metrics/Tracing Hook（可选）

- 通过配置注入 Hook 函数（例如：`OnRequest`, `OnResponse`, `OnError`）。
- SDK 自身不直接绑定 Prometheus/OpenTelemetry 等具体实现。

---

## 10. 文档与示例的技术要求

- 所有示例代码必须能通过 `go build`。
- 示例中使用的导入路径必须与实际模块路径一致：  
  `github.com/brain-xai/pingcode_api/sdk/...`
- 每个对外 Service 至少有一个完整示例：
  - 初始化 Client；
  - 调用真实 API；
  - 处理错误与结果。

---

## 11. 命名规范

本章节统一约束 SDK 的命名风格，避免出现“同一概念多种命名”的混乱情况。所有新代码必须遵守本节约定，旧代码逐步收敛。

### 11.1 包命名

- 包名一律使用小写，禁止下划线和驼峰，例如：`project`、`testhub`、`httpclient`。
- 业务领域包名直接使用领域名：
  - `sdk/model/project`、`sdk/service/project`、`internal/api/project` 等。
- 缩写统一：
  - `DevOps` 统一为 `devops`；
  - `TestHub` 统一为 `testhub`；
  - 不允许同一概念出现多种大小写变体（例如 `testHub`、`testHubService` 这类写法视为错误）。

### 11.2 类型、结构体与接口命名

- 所有导出类型首字母大写，使用驼峰命名：`Project`、`TestCase`、`WikiPage`。
- 结构体命名必须体现含义，不使用无信息缩写：
  - 实体：`Project`、`Requirement`、`TestPlan` 等；
  - 输入：`ProjectCreateInput`、`ProjectUpdateInput`；
  - 过滤：`ProjectFilter`、`TestCaseFilter`；
  - 响应包装（如有需要）：`ListProjectsResult`。
- 接口命名：
  - 行为接口使用 `er` 结尾或领域名：`Logger`、`HTTPClient`、`AuthProvider`；
  - Service 类型统一命名为 `Service`，放在对应的 `sdk/service/<domain>` 包，例如 `type Service struct { ... }`，通过包名区分领域。
- ID 类型：
  - 使用 `XXXID` 形式，例如：`ProjectID`、`IssueID`、`UserID`；
  - 这些类型统一放在 `sdk/model/core` 中，避免各处重复定义。

### 11.3 函数与方法命名

- 对外导出方法使用动词开头、语义清晰的驼峰命名：
  - 查询：`GetProject`、`ListProjects`、`GetTestCase`；
  - 创建：`CreateProject`、`CreateRequirement`；
  - 更新：`UpdateProject`、`UpdateTestPlan`；
  - 删除/归档：`DeleteProject`、`ArchiveRequirement`。
- 同一领域内相同语义必须使用相同动词：
  - 获取单个资源统一用 `Get`，禁止混用 `Fetch` / `Find` / `Load`；
  - 创建统一用 `Create`，禁止某些地方用 `New` 表示远端创建。
- Service 方法签名必须包含 `context.Context` 作为第一个参数：
  - `func (s *Service) GetProject(ctx context.Context, id core.ProjectID) ...`

### 11.4 变量与参数命名

- 变量和参数使用小驼峰命名，尽量短而不失语义：`projectID`、`userID`、`filter`、`ctx`。
- 上下文变量统一命名为 `ctx`，不要使用 `c`、`context` 等其他形式。
- 避免无信息命名：`data`、`info`、`res` 等仅在非常局部、含义明确的情况下使用，禁止在对外 API 中出现这类名称。

### 11.5 错误与常量命名

- 错误变量：
  - 使用 `errXxx` 形式，例如：`errInvalidConfig`、`errUnauthorized`；
  - 对外可见的错误类型统一集中在 `sdk/errors.go` 或 `sdk/model/core`，避免各处散落。
- 常量：
  - 使用大写 + 下划线：`DefaultTimeout` 这类导出常量可以使用驼峰；
  - 仅当确实需要暴露给使用者时才导出常量，否则保持在内部包。

### 11.6 缩写与专有名词

- 通用缩写遵循 Go 官方惯例：
  - `ID`、`URL`、`HTTP` 等缩写整体视为一个单词：`ProjectID`、`BaseURL`、`HTTPClient`；
  - 不允许出现 `ProjectId`、`HttpClient` 这类大小写错误。
- PingCode 领域内的专有名词，需要在首次使用处固定写法，并在后续保持一致，例如：
  - `TestHub` → 包名 `testhub`，类型名 `TestHubConfig`（如确有需要）。

### 11.7 命名冲突与演进

- 若 OpenAPI 中的命名与现有 SDK 命名冲突：
  - 以 SDK 对外命名为准，通过 `internal/api` 映射层做适配；
  - 禁止为了“对齐 OpenAPI”而直接破坏已有 SDK 的公开命名（违反向后兼容）。
- 新增 API 或模型时，应优先复用现有命名模式；如果必须引入新概念，先在本章节补充命名约定再落地实现。
