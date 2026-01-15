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

## 2. 目录与模块规范

> 下面描述的是与部署/构建强相关的目录与模块约定，而不是完整项目结构。

- 根模块：`github.com/brain-xai/pingcode_api`
  - 不允许再创建第二个 Go module，避免多模块构建混乱。
- SDK 相关目录（建议）：
  - `sdk/`：对外暴露的 SDK 入口与各 Service 实现
    - `sdk/client.go`：Client 定义与初始化逻辑
    - `sdk/config.go`：Config 定义与默认值
    - `sdk/auth/`、`sdk/project/`、`sdk/ship/`、`sdk/testhub/`、`sdk/wiki/` …
  - `internal/`：仅供 SDK 内部使用的工具
    - `internal/httpclient/`：HTTP 封装、重试、日志、metrics Hook
    - `internal/errors/`：错误类型定义与映射
    - `internal/encoding/`：JSON 编解码、时间格式等
  - `cmd/`（可选）：与 SDK 相关的命令行工具，如 SDK 生成器或调试工具
  - `examples/`：示例代码，必须可编译运行

目录约束：

- `internal/` 下的任何包不得被 SDK 使用者直接引用。
- 对外导出的类型与函数只能位于：
  - `sdk/` 及其子目录
  - `cmd/` 下仅用于可执行程序，不对外形成稳定 API。

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
- 主干分支任何 commit 必须保持“随时可发布”的状态。

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