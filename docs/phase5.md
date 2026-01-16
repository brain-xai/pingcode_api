# Phase 5 迭代规划：需求（Requirement）领域接口

目标：在 Ship 领域整体能力的基础上，**专注完成“需求（Requirement）相关接口”**，使 SDK 能够在真实业务中完整管理需求的生命周期（查询、创建、更新、状态变更等），并与产品/项目上下文正确关联。

---

## 1. 范围与目标

### 1.1 本阶段范围

围绕需求（Requirement）领域，完成：

1. **需求领域模型与 Service 完整化**
   - `sdk/model/ship`：完善需求相关模型（Requirement、状态、优先级等）
   - `sdk/service/ship`：补齐需求相关的读写接口
   - `internal/api/ship`：对应需求相关 OpenAPI DTO 的定义和对齐

2. **主要接口能力（Requirement 核心）**
   - 获取需求：
     - 按 ID 获取需求详情：`GetRequirement`
     - 按条件获取需求列表：`ListRequirements`
   - 管理需求：
     - 创建需求：`CreateRequirement`
     - 更新需求：`UpdateRequirement`
     - 变更需求状态（如：打开、进行中、已完成、关闭等）：
       - 通过 `UpdateRequirement` 接口传入 `StateID` 实现
   - 关联上下文：
     - 按产品 / 项目 / 迭代等维度过滤需求（具体以 OpenAPI 能力为准）

3. **辅助接口能力（用于 UI 场景）**
   - 获取需求配置信息：
     - `GetRequirementStates` - 获取产品的需求状态列表
     - `GetRequirementPriorities` - 获取产品的优先级列表
     - `GetRequirementProperties` - 获取产品的自定义属性列表
     - `GetRequirementSuites` - 获取产品的模块列表
     - `GetRequirementPlans` - 获取产品的排期列表
   - 用途：在创建/编辑需求时提供下拉选项

4. **请求正确性与一致性**
   - 所有需求相关 HTTP 请求的方法、路径、Query 参数、Body 和 Headers 与 OpenAPI 定义一致
   - 参数校验合理，错误映射符合 `error_model.md`

5. **示例与文档**
   - 提供至少一个完整示例：
     - “在某个产品/项目下查询需求列表，并创建一个新需求”
   - 文档中补充“需求管理”使用指南

### 1.2 不在本阶段范围

- 需求下细粒度子资源（如附件、评论、历史记录等），可留到后续阶段
- 高级功能：批量导入需求、复杂报表/统计、智能排序/推荐
- 与测试/DevOps 等其他领域的深度联动接口

---

## 2. Requirement 领域分析与接口清单

### 2.1 阅读与梳理 OpenAPI 文档

**目标：** 从 `docs/reference/openpingcode/ship.*`（或对应 OpenAPI 源文件中 Ship 模块）中提取需求相关接口。

需要完成：

1. 识别需求资源的命名：
   - 可能叫：`Requirement` / `Story` / `Issue` / `Idea` 等，以实际接口命名为准
2. 整理以下信息：
   - 获取需求详情的接口（GET + path）
   - 获取需求列表的接口（GET + path + query）
   - 创建需求接口（POST）
   - 更新需求接口（PUT/PATCH）
   - 状态变更接口（若有）、批量操作接口（如在本阶段决定支持）
3. 为每个接口记录：
   - HTTP 方法与路径（含 path param 模板）
   - Query 参数（过滤条件、分页、排序等）
   - 请求体结构（创建/更新）
   - 返回的主要字段及结构

产出：一份 Requirement 接口清单，作为实现和测试的基准。

---

## 3. Requirement 模型设计（sdk/model/ship）

### 3.1 Requirement 实体模型

需要完成：

- 在 `sdk/model/ship` 中定义/完善：

  - `Requirement`：
    - 核心字段（以 OpenAPI 为准，至少包括）：
      - `ID`（建议使用 `core.RequirementID`）
      - `Title`
      - `Description`
      - `Status`
      - `Priority`
      - `Assignee`（负责人，可能为 `core.UserID` 或嵌套结构）
      - `Reporter`/`Creator`
      - `ProductID` / `ProjectID`
      - 时间相关字段：`CreatedAt`、`UpdatedAt`、`StartAt`、`DueAt`（如有）
  - 状态与优先级：
      - `RequirementStatus`（枚举型别，内部使用 string）
      - `RequirementPriority`（枚举型别）

- 将跨领域通用字段统一放在 `sdk/model/core`（如 ID、时间类型），Requirement 只组合使用。

### 3.2 Requirement 入参与过滤模型

需要完成：

- 在 `sdk/model/ship` 中定义：

  - `RequirementFilter`：
    - 典型过滤条件：
      - `ProductID` / `ProjectID`
      - `Status` 列表
      - `AssigneeID`
      - 关键字（标题/描述）
      - 分页参数（引用 `core.PageRequest` 或类似）
  - `RequirementCreateInput`：
    - 必填字段：标题、所属产品/项目、类型等
    - 选填字段：描述、优先级、初始状态、负责人等
  - `RequirementUpdateInput`：
    - 可更新字段：标题、描述、优先级、负责人等（不包含状态变更时，可根据 OpenAPI 设计）
  - `RequirementStatusChangeInput`（若 OpenAPI 状态变更接口需要复杂信息）：
    - 目标状态、新状态原因等字段

设计要求：

- 输入结构体要对使用者语义清晰，避免暴露后端内部字段名
- 默认值和可选值处理清晰，避免“零值即有效”引发歧义

---

## 4. Requirement DTO 与映射（internal/api/ship）

### 4.1 DTO 定义

需要完成：

- 在 `internal/api/ship` 中，为需求相关接口定义 DTO：

  - `GetRequirementRequestDTO` / `GetRequirementResponseDTO`
  - `ListRequirementsRequestDTO` / `ListRequirementsResponseDTO`
  - `CreateRequirementRequestDTO` / `CreateRequirementResponseDTO`
  - `UpdateRequirementRequestDTO` / `UpdateRequirementResponseDTO`
  - `ChangeRequirementStatusRequestDTO` / `ChangeRequirementStatusResponseDTO`（如 OpenAPI 有对应接口）

要求：

- DTO 字段、JSON tag 与 OpenAPI 文档严格一致
- 对可选字段使用指针或合适零值，避免歧义

### 4.2 model ↔ api 映射实现

需要完成：

- 在 Ship service 或独立 mapper 文件中实现双向映射：

  - DTO → 模型：
    - `RequirementDTO` → `Requirement`
  - 模型/输入 → DTO：
    - `RequirementCreateInput` → `CreateRequirementRequestDTO`
    - `RequirementUpdateInput` → `UpdateRequirementRequestDTO`
    - `RequirementFilter` → `ListRequirementsRequestDTO`（包含 query/分页映射）
    - `RequirementStatusChangeInput` → `ChangeRequirementStatusRequestDTO`

要求：

- 映射逻辑集中管理（例如 `mapper_requirement.go`），不在各个函数中散写
- 对状态/优先级等枚举提供安全转换逻辑（未知值应有兜底策略）
- 对映射逻辑编写单元测试，覆盖常见字段组合

---

## 5. Requirement Service 实现（sdk/service/ship）

### 5.1 Service 方法签名

需要完成：

- 在 `sdk/service/ship` 的 `Service` 中增加需求相关方法（示例）：

  - 读取：
    - `GetRequirement(ctx context.Context, id core.RequirementID) (*model.Requirement, error)`
    - `ListRequirements(ctx context.Context, filter model.RequirementFilter) ([]model.Requirement, *core.Pagination, error)`
  - 写入：
    - `CreateRequirement(ctx context.Context, input model.RequirementCreateInput) (*model.Requirement, error)`
    - `UpdateRequirement(ctx context.Context, id core.RequirementID, input model.RequirementUpdateInput) (*model.Requirement, error)`
    - `ChangeRequirementStatus(ctx context.Context, id core.RequirementID, status model.RequirementStatus, opts *model.RequirementStatusChangeInput) (*model.Requirement, error)`（具体签名根据 OpenAPI 做微调）

要求：

- 所有方法第一个参数为 `ctx context.Context`
- 对外类型全部来自 `sdk/model/...`，不暴露 `internal/api` 类型
- 命名遵守 `deployment.md` 中的命名规范（Get/List/Create/Update/Change）

### 5.2 参数校验与错误处理

需要完成：

- 参数校验：

  - `RequirementID` 非空
  - `RequirementCreateInput` 必填字段不为空（如 Title、所属对象等）
  - `RequirementUpdateInput` 在允许空字段情况下的语义（例如是否支持部分更新）
  - 状态变更时，合法状态集合校验（如在 SDK 侧保留可选的状态枚举）

- 错误处理：

  - 参数错误 → 明确为 Usage/Config 错误（可返回带特定错误码的 `APIError` 或预定义错误）
  - 认证错误 → 与现有错误模型一致
  - 业务错误 → 保持 HTTP 状态码 + 业务错误码 + 消息，方便调用方区分
  - 网络/系统错误 → 透传关键信息，利于排查

---

## 6. 请求正确性与测试（需求专项）

### 6.1 请求构造测试

需要完成：

- 使用 mock HTTP Client / RoundTripper 为以下接口编写请求构造测试：

  - `GetRequirement`：
    - 验证 URL path 中 ID 替换正确
    - 验证 HTTP 方法为 GET
  - `ListRequirements`：
    - 验证分页、过滤条件正确序列化为 query 参数
  - `CreateRequirement` / `UpdateRequirement`：
    - 验证请求体 JSON 结构与 OpenAPI 示例一致
  - `ChangeRequirementStatus`：
    - 验证路径/请求体中状态字段正确映射

### 6.2 行为与错误测试

需要完成：

- 单元测试：

  - 映射逻辑（model ↔ DTO）覆盖主要路径与错误路径
  - 参数校验（空 ID、缺少必填字段、非法状态等）

- 集成测试（视实际环境能力）：

  - 在测试环境中执行至少一个完整流程（可选方案）：
    1. 列出某产品/项目下的需求
    2. 创建一个新需求
    3. 更新该需求的部分字段
    4. 修改需求状态
  - 验证每一步的返回数据符合预期，且错误模型合理

验收标准：

- `go test ./...` 依然通过
- Requirement 相关测试不会破坏已有模块测试

---

## 7. 示例与文档

### 7.1 示例代码

需要完成：

- 新增或扩展示例，例如：

  - `examples/requirements/main.go`

- 示例流程建议：

  1. 从环境变量读取 BaseURL 与 Token
  2. 选择一个 Product / Project（可通过 Ship / Project Service 获取）
  3. 调用 `Ship().ListRequirements` 拉取该上下文下的需求列表
  4. 调用 `Ship().CreateRequirement` 创建一个新需求
  5. 调用 `Ship().ChangeRequirementStatus` 将需求状态修改为指定状态

要求：

- 示例可编译运行（在正确配置下）
- 示例中有基础的错误处理与打印，便于调试

### 7.2 文档更新

需要完成：

- 在 SDK 文档中新增「需求管理」小节：

  - 如何查询需求列表
  - 如何创建和更新需求
  - 如何变更需求状态
  - 示例代码片段（可引用 `examples/requirements`）

- 更新规范文档（如有需要）：

  - `api_contract.md`：加入需求相关对外 API 列表
  - `error_model.md`：如针对需求领域有特定错误码（例如状态非法、父级对象不存在），补充说明

---

## 8. Phase 5 验收标准

本阶段完成的判定标准：

1. Requirement 领域主要读写接口完成：
   - 按 ID 获取需求详情
   - 按条件列表查询需求
   - 创建需求
   - 更新需求基础信息
   - 状态变更（如 OpenAPI 支持）
2. 请求行为正确：
   - 方法、路径、query、body 与 OpenAPI 定义一致（通过单元测试验证）
3. 错误处理符合统一规范：
   - 能区分参数错误、认证错误、业务错误与网络错误
4. 测试与示例：
   - `go test ./...` 通过
   - Requirement 模块有覆盖映射、请求构造、核心行为的单元测试
   - 示例程序在正确配置下，可以跑通“查询 + 创建 + 状态变更”的典型需求流程

满足上述条件，即可认为 Phase 5（需求相关接口）迭代目标完成，为后续扩展需求子资源（评论、附件、关联关系等）打下基础。