# Phase 4 迭代规划：Project 领域接口

目标：在前几期（基础框架、GET 能力、Ship 领域）已完成的基础上，**系统性完善 Project 领域接口**，覆盖核心读取与关键写操作，并通过示例与测试确保在真实业务场景中的可用性。

---

## 1. 范围与目标

### 1.1 本阶段范围

围绕 Project 领域，完成：

1. **Project 领域模型与 Service 完整化**
   - `sdk/model/project`：完善项目相关模型（Project、成员、标签等核心对象）
   - `sdk/service/project`：丰富读取与关键写接口
   - `internal/api/project`：与 OpenAPI 对齐的 DTO 集合

2. **主要接口能力**
   - 项目本体：
     - 获取项目详情：`GetProject`
     - 获取项目列表：`ListProjects`（在 Phase 1/2 的基础上补全过滤/分页等）
     - 创建项目：`CreateProject`
     - 更新项目：`UpdateProject`
     - 归档/关闭项目（如有）
   - 项目成员/参与者：
     - 查询项目成员列表：`ListProjectMembers`
   - 项目基本元数据（可根据 OpenAPI 决定是否在本阶段先实现）：
     - 查询项目可用的状态/类型/自定义字段等

3. **请求正确性保障（针对 Project）**
   - HTTP 方法、路径、Query、Headers 与 OpenAPI 保持一致
   - 对 ID、必填字段等进行本地校验
   - 错误处理遵循 `error_model.md`

4. **示例与文档**
   - 至少提供一个完整示例：
     - 展示“获取项目列表 → 获取单个项目详情 → 查询项目成员”的完整链路
   - 更新文档中 Project 相关说明

### 1.2 不在本阶段范围

- 项目下子资源的复杂操作（如某些高度业务化的统计、报表）
- 高度定制化的查询构建器（暂以基本过滤为主）
- 与其他领域的深度联动（例如一次性拉取项目+需求+测试的整合接口）

---

## 2. Project 领域分析与接口清单

### 2.1 阅读与梳理 OpenAPI 文档

**目标：** 从 `docs/reference/openpingcode/project.*`（或对应 OpenAPI 文件）中提取 Project 领域核心接口。

需要完成：

1. 梳理 Project 领域主要资源：
   - Project / 项目
   - ProjectMember / 项目成员
   - 其它关键资源（如项目分组、项目配置等，以 OpenAPI 为准）
2. 按资源整理以下信息：
   - HTTP 方法 + 路径（含 path param）
   - 支持的 Query 参数（分页、过滤条件）
   - 请求体结构（针对创建/更新）
   - 核心响应字段

产出：一份 Project 接口清单（内部使用，可附于本文件或独立文档）。

### 2.2 Phase 4 优先接口列表

从清单中选定 Phase 4 必须完成的高优先级接口，初步建议：

1. **项目本体**
   - `GetProjectByID`：按 ID 获取项目详情
   - `ListProjects`：获取项目列表（支持常用过滤，如按状态、负责人、产品等）
   - `CreateProject`：创建新项目
   - `UpdateProject`：更新项目基础信息（名称、描述、状态、可见性等）
   - `ArchiveProject` / `CloseProject`（如 OpenAPI 中存在）

2. **项目成员**
   - `ListProjectMembers`：获取项目成员列表（用于权限/协作场景）

如 OpenAPI 命名与上述不同（例如 `project/users` 等），以实际命名为准，但在 SDK 中遵守 `deployment.md` 和 `api_contract.md` 约定的命名风格。

---

## 3. Project 模型设计（sdk/model/project）

### 3.1 Project 模型

需要完成：

- 在 `sdk/model/project` 中定义/完善：

  - `Project`：
    - 最少包含：`ID`、`Name`、`Key`、`Status`、`Owner`、`CreatedAt`、`UpdatedAt` 等
  - `ProjectFilter`：
    - 可选字段：状态、负责人、产品 ID、类型、关键字、分页参数等
  - `ProjectCreateInput`：
    - 创建项目的必填/可选字段（名称、Key、所属产品/空间、可见性等）
  - `ProjectUpdateInput`：
    - 更新项目的可修改字段（名称、描述、状态、可见性等）

- 与 `sdk/model/core` 联动：
  - 使用 `ProjectID`、`UserID`、`Time`、`Pagination` 等统一类型

### 3.2 Project 成员模型

需要完成：

- 在 `sdk/model/project` 中定义：

  - `ProjectMember`：
    - 包含：成员 ID、用户信息（或用户 ID）、角色/权限、加入时间等
  - `ProjectMemberFilter`（视 OpenAPI 能力决定是否需要）

设计原则：

- 不直接暴露内部权限编码细节，如有必要可提供枚举或封装类型
- 字段命名对使用者友好，符合命名规范

---

## 4. Project DTO 与映射（internal/api/project）

### 4.1 DTO 定义

需要完成：

- 在 `internal/api/project` 中，为 Project 相关接口定义 DTO：

  - `GetProject` 请求/响应 DTO
  - `ListProjects` 请求/响应 DTO
  - `CreateProject` 请求体/响应 DTO
  - `UpdateProject` 请求体/响应 DTO
  - `ListProjectMembers` 请求/响应 DTO

- JSON tag 和字段类型与 OpenAPI 文档严格对齐

### 4.2 model ↔ api 映射

需要完成：

- 在 Project service 或独立 mapper 文件中实现：

  - `internal/api` DTO → `sdk/model/project` 模型的转换：
    - `ProjectDTO` → `Project`
    - `ProjectMemberDTO` → `ProjectMember`
  - `sdk/model/project` 输入结构体 → `internal/api` 请求 DTO：
    - `ProjectCreateInput` → `CreateProjectRequestDTO`
    - `ProjectUpdateInput` → `UpdateProjectRequestDTO`
    - `ProjectFilter` → `ListProjectsRequestDTO`（path/query 组合）

要求：

- 所有映射逻辑集中、可测试，不在业务代码中随处散落
- 对于 OpenAPI 中存在但暂时不需要暴露的字段，可在映射中带过去，但不一定全部暴露到 `Project`，以避免 API 污染

---

## 5. Project Service 实现（sdk/service/project）

### 5.1 Service 结构与注入

需要完成：

- 确保 `sdk/service/project` 中：

  - `Service` 结构体持有 HTTP Client 及必要上下文
  - `sdk/client.go` 中对外提供 `Project()` 方法返回该 Service

### 5.2 读取接口（GET）

需要完成以下方法（参考命名）：

- `Get(ctx context.Context, id core.ProjectID) (*model.Project, error)`
- `List(ctx context.Context, filter model.ProjectFilter) ([]model.Project, *core.Pagination, error)`
- `ListMembers(ctx context.Context, id core.ProjectID) ([]model.ProjectMember, error)`

要求：

- 参数校验：
  - ProjectID 非空
  - Filter 中必要的组合关系（如果有）需在本地校验
- 错误处理：
  - 认证失败 → 归类为 Auth 错误
  - 参数非法 → 归类为 Usage/Config 错误
  - 业务错误/网络错误 → 使用统一错误类型包装

### 5.3 写接口（Create/Update/Archive）

需要完成：

- `Create(ctx context.Context, input model.ProjectCreateInput) (*model.Project, error)`
- `Update(ctx context.Context, id core.ProjectID, input model.ProjectUpdateInput) (*model.Project, error)`
- 如 OpenAPI 支持归档/关闭项目：
  - `Archive(ctx context.Context, id core.ProjectID) error`
  - 或 `Close(ctx context.Context, id core.ProjectID) error`

要求：

- 对输入参数进行本地校验（必填字段、字段范围）
- 正确构造请求体并发送 HTTP 请求
- 更新后的项目从响应中映射回 `Project`

---

## 6. 请求正确性与测试（Project 专项）

### 6.1 请求构造单元测试

需要完成：

- 针对以下接口编写请求构造测试：

  - `Get`：验证 URL path、中间路径变量正确
  - `List`：验证分页、过滤参数正确序列化为 query
  - `Create`/`Update`：验证请求体 JSON 结构与 OpenAPI 一致
  - `ListMembers`：验证项目成员路径与 query 正确

- 可通过自定义 RoundTripper 或 fake HTTP Client 捕获：

  - 最终请求的 URL、方法、header、body

### 6.2 行为与错误测试

需要完成：

- 单元测试：

  - 参数校验覆盖所有分支（为空、非法组合）
  - model ↔ api 映射逻辑的双向测试

- 集成测试（视测试环境情况）：

  - 在测试或沙箱环境中调用：
    - `ListProjects`：至少验证正常返回 + 无权限/认证失败两种情况
    - `GetProject`：验证获取单个项目
    - 如环境允许，测试 `Create`/`Update` 对真实数据的影响（可专门准备测试项目）

验收标准：

- `go test ./...` 通过
- Project 相关测试用例涵盖核心功能与主要错误路径

---

## 7. 示例与文档

### 7.1 示例代码

需要完成：

- 新增或扩展示例，例如：

  - `examples/project_overview/main.go`

- 示例流程示意：

  1. 从环境变量加载配置和 Token
  2. 调用 `client.Project().List` 获取项目列表
  3. 选取一个项目 ID：
     - 使用 `Get` 拉取详情
     - 使用 `ListMembers` 拉取成员列表
  4. 如环境允许，展示一次 `Create`/`Update` 调用（可对测试项目操作）

### 7.2 文档更新

需要完成：

- 在 SDK 文档中新增「Project 领域指南」：

  - 项目列表、项目详情、项目成员的调用示例
  - 对创建/更新项目的注意事项（如权限、必填字段）

- 更新相关规范：

  - `api_contract.md`：加入新的对外 API 说明
  - `error_model.md`：如新增特定错误码或错误类型说明（比如项目不存在、权限不足）

---

## 8. Phase 4 验收标准

本阶段完成的判定标准：

1. Project 领域核心资源（Project、本项目成员）的模型和 Service 实现完整：
   - 包括 Get/List/Create/Update/Delete 以及成员查询的主流程 ✅
2. 请求与 OpenAPI 对齐：
   - 已实现的 Project 接口在方法、路径、参数和 headers 上均与 OpenAPI 定义一致 ✅
   - 关键接口有 URL/请求体构造单元测试进行验证 ✅
3. 错误处理符合统一错误模型：
   - 对认证、参数、业务、网络错误有清晰的分类与返回行为 ✅
4. 测试与示例：
   - `go test ./...` 通过 ✅
   - Project 模块的映射、请求构造、行为覆盖单元测试 ✅
   - 示例程序能在正确配置环境下跑通"项目列表 → 项目详情 → 成员列表"基本链路 ✅

满足以上条件，认为 Phase 4（Project 相关接口）迭代完成，为后续扩展更多领域和高级能力奠定稳定基础。

---

## 实现记录

### 已完成接口

| 接口 | 方法 | 路径 | 状态 |
|------|------|------|------|
| 获取项目列表 | GET | /v1/project/projects | ✅ 已实现 |
| 获取项目详情 | GET | /v1/project/projects/{project_id} | ✅ 已实现 |
| 创建项目 | POST | /v1/project/projects | ✅ 已实现 |
| 更新项目 | PATCH | /v1/project/projects/{project_id} | ✅ 已实现 |
| 删除项目 | DELETE | /v1/project/projects/{project_id} | ✅ 已实现 |
| 获取项目成员列表 | GET | /v1/project/projects/{project_id}/members | ✅ 已实现 |
| 获取项目状态列表 | GET | /v1/project/project/states | ✅ 已实现 |
| 获取项目进度 | GET | /v1/project/projects/{project_id}/progress | ✅ 已实现 |

### 新增文件

1. `sdk/model/project/input.go` - 创建和更新项目的输入模型
2. `internal/api/project/dto.go` - 请求 DTO 和映射函数
3. `examples/project_overview/main.go` - 完整的项目操作示例

### 修改文件

1. `sdk/service/project/project.go` - 新增 Create/Update/Delete 方法
2. `sdk/service/project/project_test.go` - 新增 Create/Update/Delete 测试
3. `internal/httpclient/client.go` - 新增 Delete/DeleteWithPathParams 方法

### 测试覆盖

- 参数验证测试：Name、Type、Identifier 必填校验，长度限制，类型枚举校验
- 成功场景测试：Create、Update、Delete 完整流程
- URL 构造测试：验证 HTTP 方法、路径、请求体正确性
- 所有测试通过 ✅