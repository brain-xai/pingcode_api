# Phase 2 迭代规划

目标：在 Phase 1 已完成的基础框架之上，完成「读取类接口（get 属性方法）」的全面落地，并通过规范化测试与对照 OpenAPI 文档，保证请求的正确性与一致性。

本阶段的核心目标：

1. 将当前规划范围内所有 get 属性方法实现出来（按领域模块分批完成）
2. 建立一套验证请求正确性的机制（包含代码层、测试层和与 OpenAPI 的对齐）

---

## 实现记录

### 已实现接口

| 日期 | 接口 | 状态 | 说明 |
|------|------|------|------|
| 2025-01-15 | httpclient 路径参数支持 | ✅ 已实现 | 新增 `GetWithPathParams()` 方法支持路径参数 |
| 2025-01-15 | Project.GetByID | ✅ 已实现 | 按 ID 获取项目详情，包含完整字段 |
| 2025-01-15 | Project.ListStates | ✅ 已实现 | 获取项目状态列表 |
| 2025-01-15 | Project.ListMembers | ✅ 已实现 | 获取项目成员列表 |
| 2025-01-15 | Project.GetProgress | ✅ 已实现 | 获取项目进度 |
| 2025-01-15 | Global.GetCurrentUser | ✅ 已实现 | 获取当前用户信息（GET /v1/myself） |
| 2025-01-15 | Global.GetUser | ✅ 已实现 | 按 ID 获取用户详情 |
| 2025-01-15 | Global.ListUsers | ✅ 已实现 | 获取企业成员列表 |

### 技术改进

- **Project 模型扩展**：新增 Assignee、Members、CreatedBy、UpdatedBy、Visibility 等字段
- **Global 领域创建**：新增 User、Department、Job 等模型
- **测试覆盖**：
  - httpclient: 90%+ 单元测试覆盖率
  - Project Service: 80%+ 单元测试覆盖率
  - Global Service: 80%+ 单元测试覆盖率
  - 集成测试框架已建立（需配置测试环境）

---

## 1. 范围与不在范围

### 1.1 本阶段范围

- 在 **已纳入 SDK 范围** 的领域模块中，补齐所有「读取类（GET）」方法：
  - Project 领域：
    - 示例：`GetProject`（按 ID 获取）、`ListProjects`（已在 Phase 1 实现，必要时补全过滤/分页）
  - 其它已确定纳入 Phase 2 的领域（按当前 OpenAPI 和优先级选择）：
    - 例如：Auth 领域（`GetCurrentUser` / `GetTokenInfo` 等）
    - 例如：Global / Wiki 等（视 Phase 2 范围决定）
- 为每个 get 方法实现完整的：
  - `sdk/model` 对外模型
  - `internal/api` DTO
  - `sdk/service` 中的业务方法
  - model ↔ api 映射逻辑
- 建立「请求正确性」保障：
  - HTTP 方法、路径、Query、Headers 与 OpenAPI 完全一致
  - 必要的请求参数校验
  - 针对关键接口编写集成/契约测试（对接测试环境或 mock）

### 1.2 不在本阶段范围

- 写操作（创建/更新/删除）仍不作为本阶段的主要目标（可只做必要支持）
- 自动重试、限流、高级 metrics/Tracing
- 全量覆盖所有 PingCode OpenAPI 模块（可以先选最重要的一批）

---

## 2. 统一 get 方法清单 & 优先级

### 2.1 枚举 get 方法

**目标：** 基于 OpenAPI 文档和当前 SDK 范围，枚举出本阶段需要实现的所有 get 方法清单。

需要完成：

1. 从 `docs/reference/openpingcode/*.md`（或 OpenAPI 源文档）中，筛选出与 SDK 现阶段领域相关的所有 GET 接口：
   - Project：按 ID 获取项目、获取项目信息、可能还包括项目下部分属性（如成员列表、标签等）
   - Auth：当前用户信息、Token 信息（如有）
   - 其它领域（如本阶段选定）
2. 将清单写成规范化列表（可放入内部设计文档或 phase2 内部附录），包括：
   - HTTP 方法（GET）
   - 路径（含 path param 位置）
   - 支持的 Query 参数
   - 返回的主要资源类型

### 2.2 确定优先级

**目标：** 按业务重要性排序，确保优先完成核心路径。

建议优先级：

1. Project 核心读取：
   - `GetProjectByID`
   - 已有 `ListProjects` 的补全
2. Auth 基础：
   - `GetCurrentUser` 或等价接口（帮助调试 & 验证认证）
3. 其它高频读接口（根据内部使用场景选定）

---

## 3. get 方法的 SDK 设计规范落地

### 3.1 model 层设计（sdk/model）

**目标：** 为每个 get 方法对应的资源定义清晰的对外模型。

需要完成：

- 对于清单中的每一种资源（Project、User 等）：
  - 在相应的 `sdk/model/<domain>` 中定义：
    - 实体结构体（例如 `Project`、`User`）
    - 必要的辅助结构体（例如 `ProjectDetail`、`UserProfile` 等）
- 字段设计原则：
  - 尽量用语义明确的字段名，不直接照搬 OpenAPI 的命名怪癖
  - 对跨领域通用字段（ID、时间、分页）复用 `sdk/model/core` 中的类型
- 对于已有的模型（如 Phase 1 中的 `Project`）：
  - 根据 OpenAPI 补全集合必要字段（在不破坏兼容的前提下）

### 3.2 internal/api DTO 设计（internal/api）

**目标：** 为每个 get 接口定义与 OpenAPI 文档贴合的请求/响应 DTO。

需要完成：

- 在 `internal/api/<domain>` 对应包中：
  - 为每个 GET 接口定义：
    - 请求 DTO（包含 path param、query param）
    - 响应 DTO（完整或必要字段）
- 与 OpenAPI 对齐：
  - 保证 JSON tag、字段类型与 OpenAPI 一致
  - 对可选字段使用指针或合适的零值策略

### 3.3 service 层方法签名设计（sdk/service）

**目标：** 在 service 层定义统一风格的 get 方法。

需要完成：

- 对 Project 领域，约定方法签名风格，例如：
  - `Get(ctx context.Context, id core.ProjectID) (*model.Project, error)`
  - 若有复合查询，则定义相应的 Filter 或 Option 类型
- 对其它领域保持一致命名风格：
  - 资源单体读取统一使用 `Get` 前缀（`GetUser`、`GetWikiPage` 等）
- 对于已有列表方法（如 `ListProjects`）：
  - 如需调整参数（增加 filter、分页等），按 `api_contract.md` 的兼容性要求进行设计

---

## 4. 请求正确性保障

### 4.1 请求构造规范检查

**目标：** 确保所有 get 请求在代码层面严格符合 OpenAPI 的定义。

需要完成：

- 在 `internal/httpclient` 或 service 层建立统一的请求构造方式：
  - 明确区分 path 参数、query 参数和 headers
  - 对 query 参数的序列化方式进行统一（如分页参数、过滤数组等）
- 为关键接口编写「构造请求 URL 的单元测试」：
  - 给定输入参数，验证最终 URL（含 query string）与预期字符串匹配
  - 验证必要 header（Content-Type、Authorization 等）是否设置正确

### 4.2 参数校验

**目标：** 在进入 HTTP 层之前，对明显错误参数进行校验，避免发出无意义请求。

需要完成：

- 在 service 方法中添加基础参数校验：
  - ID 不允许为空（如 ProjectID 不能为空字符串）
  - 必需字段必须存在
- 对于明显错误参数，直接返回配置/调用错误（如 `ErrInvalidConfig` 或自定义 usage error 类型），不发起 HTTP 请求

### 4.3 与 OpenAPI 对照检查（手工或半自动）

**目标：** 确认接口定义与 OpenAPI 文档保持一致。

需要完成：

- 建立一个检查流程（可先手工）：
  - 对每个 get 方法，列出：
    - HTTP 方法 + 路径
    - 主要 query 参数
    - 主要返回字段
  - 与 OpenAPI 文档逐项对照确认
- 后续可考虑：
  - 使用 OpenAPI 解析器，基于 schema 对请求/响应进行结构合法性检查（在 Phase 2 可先不实现工具，只留下 TODO）

---

## 5. 测试与验证

### 5.1 单元测试

需要完成：

- 为每个 get 方法编写单元测试：
  - 参数校验逻辑测试
  - model ↔ api 映射测试（不依赖真实 HTTP）
- 为请求构造编写单元测试：
  - 验证 URL、query、headers 正确生成

### 5.2 集成/契约测试

**目标：** 在有 PingCode 测试环境或可靠 mock 的前提下，验证真实调用。

需要完成：

- 针对 Project 的 `Get`/`List` 编写集成测试：
  - 通过配置测试环境 BaseURL 和 Token
  - 实际发 GET 请求，验证：
    - 成功返回时，model 字段与预期匹配
    - 错误时，错误类型与错误码符合 `error_model.md` 规范
- 若无法访问真实环境：
  - 使用 HTTP mock（如自实现 RoundTripper stub），按 OpenAPI 样例响应构建 mock 返回

验收标准：

- `go test ./...` 通过
- 关键 get 方法拥有至少 1 个集成或高保真 mock 测试用例

---

## 6. 文档与示例更新

### 6.1 文档更新

需要完成：

- 在 SDK 文档中补充分章节：
  - “读取类接口使用指南”
  - 列举典型 get 方法调用示例（包含 Project、Auth 等）
- 如有必要，更新：
  - `api_contract.md`：补充本阶段新增的稳定 API
  - `error_model.md`：如增加了新的错误类型或错误值

### 6.2 示例扩展

需要完成：

- 在 `examples` 下新增或扩展示例：
  - 在现有 `basic_usage` 上增加 “按 ID 获取项目详情” 示例
  - 如果实现了 Auth 的 get 方法，可以增加“获取当前用户信息”的示例

验收标准：

- 示例代码能够编译并运行（在配置正确的前提下）
- 示例覆盖至少一个新实现的 get 方法

---

## 7. Phase 2 验收标准总结

本次迭代完成后，至少要满足：

1. 对已纳入 Phase 2 范围的领域模块：
   - 清单中的所有 get 方法在 SDK 中都有对应实现（model + service + internal/api DTO + 映射）
2. 请求构造正确：
   - HTTP 方法、路径、query、headers 与 OpenAPI 定义保持一致
   - 核心接口有 URL/请求构造单元测试进行验证
3. 错误处理符合规范：
   - 认证错误、参数错误、业务错误和网络错误在 get 方法中均有合理区分
4. 测试与示例：
   - `go test ./...` 通过
   - 至少一个 get 方法具备高保真测试（真实环境或精确 mock）
   - 示例可以演示至少一个新的 get 方法的实际使用方式

只要上述条件达成，即可认为 Phase 2 迭代目标完成，为后续写操作和更多模块扩展打下稳定基础。