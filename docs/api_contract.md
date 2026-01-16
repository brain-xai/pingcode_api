# PingCode OpenAPI Golang SDK 对外 API 契约规范

> 目标：明确什么算“对外契约”，哪些行为可以改，哪些不能改，避免后期因为不同人对“兼容性”的理解不一致而产生争议。

---

## 1. 范围说明

- 本文档约束的是 **SDK 使用者可见的 API**，包括但不限于：
  - 模块路径：`github.com/brain-xai/pingcode_api/sdk/...`。
  - 导出类型：结构体、接口、别名等。
  - 导出函数与方法：`NewClient`、各领域 Service 方法等。
  - 对外可见的错误类型与错误值。
- 以下内容明确不属于对外契约：
  - `internal/` 目录下的所有实现细节。
  - 生成代码的具体结构，只要对外行为不变。
  - `cmd/` 目录下的调试工具实现细节（除非文档明确标记为对外稳定接口）。

---

## 2. 稳定 API 范围

### 2.1 模块与包路径

- 所有 `sdk/` 下的包路径视为对外稳定命名空间：
  - 例如：`github.com/brain-xai/pingcode_api/sdk/model/project`、`sdk/service/project`。
- 一旦某个包在公开文档或示例中出现，即视为对外可用包，不得随意重命名或移动。

### 2.2 导出类型与字段

- 所有导出类型（首字母大写）一旦发布到稳定版本：
  - 类型名视为契约，禁止重命名。
  - 字段的含义、类型、JSON 标签如有变更，必须按照语义化版本规则处理。
- 新增字段：
  - 允许在 MINOR 版本中新增导出字段，但必须保证旧代码在编译与运行时都不被破坏。

### 2.3 导出函数与方法

- 对外导出的函数/方法一旦进入稳定版本：
  - 函数名固定，不得更改。
  - 参数列表和返回值类型一旦变更，即视为 breaking change。
- 需要调整行为时优先策略：
  - 新增新方法（例如 `CreateProjectV2`），并在旧方法上标记 Deprecated。

### 2.4 错误类型与错误值

- SDK 对外暴露的错误类型（如统一的 `APIError` 或类似类型）属于稳定契约：
  - 字段必须有清晰含义，变更时遵守版本规则。
  - 错误类型的行为（例如是否实现 `errors.Is/As`）必须保持稳定。
- 约定的“语义性错误值”（如 `ErrUnauthorized`、`ErrTimeout`）一旦公开，视为契约的一部分。

---

## 3. Breaking Change 判定标准

以下改动一律视为 **Breaking Change**：

- 删除任意对外导出类型、字段、函数、方法、常量。
- 更改导出字段的类型或语义（例如从 `string` 改为 `int`，或者改变单位）。
- 更改函数或方法的参数列表、返回值类型。
- 更改已有方法或函数的行为，使其在同样的输入下产生不同的关键结果（例如：原本创建资源改成只做校验）。
- 更改错误分类规则，导致调用方基于错误类型的分支逻辑失效。

以下改动 **不视为 Breaking Change**（前提是行为不变）：

- 增加新的导出类型、方法、字段（在调用方不需要修改现有代码的前提下）。
- 修正文档、注释、拼写错误。
- 优化内部实现（包括 `internal/` 目录）而不改变对外行为。

---

## 4. API 演进策略

### 4.1 新功能添加

- 新功能应通过新增类型/方法的方式加入现有包结构，避免修改已有 API 的语义。
- 若新功能是已有功能的严格超集：
  - 可以提供 `XxxV2`、`XxxWithOptions` 等新方法；
  - 在文档中引导用户迁移到新方法。

### 4.2 功能废弃（Deprecation）

- 当需要弃用某个 API 时，必须：
  - 在注释中使用 `Deprecated:` 说明，并给出替代方案；
  - 在 CHANGELOG 中记录废弃信息；
  - 至少在一个 MAJOR 版本周期内保留该 API，避免立即移除。

### 4.3 API 删除

- 仅在发布新的 MAJOR 版本时，才允许删除已标记为 Deprecated 且保留了足够时间的 API。
- 删除前必须给出明确的迁移路径和示例。

---

## 5. 文档与示例的一致性

- 所有对外文档（README、参考文档、示例代码）必须与本契约保持一致：
  - 一旦在对外文档中出现的调用方式，即视为被支持的稳定用法。
  - 禁止在文档中出现依赖 `internal/` 或其它内部实现细节的示例。
- 文档更新被视为发版的一部分：
  - 新增/废弃 API 时必须同步更新文档和示例。

---

## 6. 冲突处理与例外

- 当 OpenAPI 文档与现有 SDK 契约发生冲突时：
  - 以 SDK 已发布的对外契约为第一优先级；
  - 通过内部映射和适配层（`internal/api` → `sdk/model`）处理差异；
  - 必要时通过新版本 API 暴露新行为，而不是直接改变旧 API。
- 如确有必须违背以上规则的极端情况：
  - 必须在设计评审中记录例外原因和影响范围；
  - 在发布说明中明确标注为"破坏性变更"。

---

## 7. 对外 API 列表

### 7.1 Ship 领域 - 产品管理

#### 查询类
- `ListProducts(ctx, filter) ([]Product, *Pagination, error)`
  - 获取产品列表
  - 参数：filter（包含分页和查询条件）
  - 返回：产品列表和分页信息

- `GetProduct(ctx, productID) (*Product, error)`
  - 按 ID 获取产品详情
  - 参数：productID（必填）
  - 返回：产品详情

### 7.2 Ship 领域 - 需求管理

#### 查询类
- `ListRequirements(ctx, filter) ([]Requirement, *Pagination, error)`
  - 获取需求列表
  - 参数：filter（包含产品 ID、状态 ID、优先级 ID、关键字等过滤条件）
  - 返回：需求列表和分页信息

- `GetRequirement(ctx, requirementID) (*Requirement, error)`
  - 按 ID 获取需求详情
  - 参数：requirementID（必填）
  - 返回：需求详情

#### 写入类
- `CreateRequirement(ctx, input) (*Requirement, error)`
  - 创建新需求
  - 参数：input（包含产品 ID、标题等必填字段）
  - 返回：创建的需求详情

- `UpdateRequirement(ctx, requirementID, input) (*Requirement, error)`
  - 更新需求信息
  - 参数：requirementID（必填）、input（包含标题、描述、状态 ID、优先级 ID 等可选字段）
  - 返回：更新后的需求详情
  - 注意：状态变更通过传入 `StateID` 字段实现

### 7.3 Ship 领域 - 需求辅助接口

#### 配置查询类（用于 UI 场景）

- `GetRequirementStates(ctx, productID) (*RequirementStateList, error)`
  - 获取产品下的需求状态列表
  - 参数：productID（必填）
  - 返回：分页的状态列表（包含 ID、名称、类型）

- `GetRequirementPriorities(ctx, productID) (*RequirementPriorityList, error)`
  - 获取产品下的优先级列表
  - 参数：productID（必填）
  - 返回：分页的优先级列表（包含 ID、名称）

- `GetRequirementProperties(ctx, productID) (*RequirementPropertyList, error)`
  - 获取产品下的自定义属性列表
  - 参数：productID（必填）
  - 返回：分页的属性列表（包含 ID、名称、类型、选项）

- `GetRequirementSuites(ctx, productID) (*RequirementSuiteList, error)`
  - 获取产品下的模块列表
  - 参数：productID（必填）
  - 返回：分页的模块列表（包含 ID、名称、类型）

- `GetRequirementPlans(ctx, productID) (*RequirementPlanList, error)`
  - 获取产品下的排期列表
  - 参数：productID（必填）
  - 返回：分页的排期列表（包含 ID、名称、URL）

**用途说明**：
- 这些接口主要用于 UI 场景
- 在创建/编辑需求时提供下拉选项
- 所有接口都需要 `productID` 参数
- 返回结果包含分页信息

### 7.4 Project 领域 - 项目管理

#### 查询类
- `ListProjects(ctx, filter) ([]Project, *Pagination, error)`
  - 获取项目列表

- `GetProject(ctx, projectID) (*Project, error)`
  - 按 ID 获取项目详情

#### 写入类
- `CreateProject(ctx, input) (*Project, error)`
  - 创建新项目

- `UpdateProject(ctx, projectID, input) (*Project, error)`
  - 更新项目信息

- `DeleteProject(ctx, projectID) error`
  - 删除项目

### 7.5 Global 领域 - 全局服务

- `GetUser(ctx, userID) (*User, error)`
  - 获取用户信息
  - 参数：userID（必填）
  - 返回：用户详情

### 7.6 WorkItem 领域 - 工作项管理

#### 基础操作类

- `Create(ctx, input) (*WorkItem, error)`
  - 创建工作项
  - 参数：input（包含项目 ID、类型 ID、标题等必填字段）
  - 返回：创建的工作项详情

- `Update(ctx, workItemID, input) (*WorkItem, error)`
  - 部分更新工作项
  - 参数：workItemID（必填）、input（包含所有可选字段，使用指针类型表示部分更新）
  - 返回：更新后的工作项详情

- `BatchUpdate(ctx, input) (*WorkItemBatchUpdateResult, error)`
  - 批量部分更新工作项属性
  - 参数：input（包含工作项 ID 列表和更新字段）
  - 返回：批量更新结果（成功 ID 列表和失败 ID 列表）

- `List(ctx, filter) ([]WorkItem, *Pagination, error)`
  - 获取工作项列表
  - 参数：filter（包含项目 ID、类型 ID、状态 ID、负责人 ID、标签 ID 等过滤条件）
  - 返回：工作项列表和分页信息

- `Get(ctx, workItemID) (*WorkItem, error)`
  - 按 ID 获取工作项详情
  - 参数：workItemID（必填）
  - 返回：工作项详情

- `Delete(ctx, workItemID) error`
  - 删除工作项
  - 参数：workItemID（必填）

#### 属性与分类类

- `ListTypes(ctx, scope) ([]WorkItemType, error)`
  - 获取工作项类型列表
  - 参数：scope（包含项目 ID）
  - 返回：工作项类型列表

- `ListStatuses(ctx, filter) (*WorkItemStatusList, error)`
  - 获取工作项状态列表
  - 参数：filter（包含项目 ID、可选的类型 ID）
  - 返回：分页的状态列表

- `ListFields(ctx, filter) (*WorkItemFieldList, error)`
  - 获取工作项属性列表
  - 参数：filter（包含项目 ID、可选的类型 ID）
  - 返回：分页的属性列表

- `ListPriorities(ctx, projectID) (*WorkItemPriorityList, error)`
  - 获取工作项优先级列表
  - 参数：projectID（必填）
  - 返回：分页的优先级列表

- `ListTags(ctx, filter) (*WorkItemTagList, error)`
  - 获取工作项标签列表
  - 参数：filter（包含项目 ID、可选的类型 ID）
  - 返回：分页的标签列表

#### 标签管理类

- `AddTag(ctx, workItemID, tagID) error`
  - 向工作项添加标签
  - 参数：workItemID（必填）、tagID（必填）

- `RemoveTag(ctx, workItemID, tagID) error`
  - 从工作项移除标签
  - 参数：workItemID（必填）、tagID（必填）

#### 关联管理类

- `ListLinkTypes(ctx, projectID) ([]WorkItemLinkType, error)`
  - 获取关联类型列表
  - 参数：projectID（必填）
  - 返回：关联类型列表

- `Link(ctx, workItemID, targetID, linkTypeID) (*WorkItemLink, error)`
  - 关联工作项
  - 参数：workItemID（源工作项 ID）、targetID（目标工作项 ID）、linkTypeID（关联类型 ID）
  - 返回：创建的关联详情

- `ListLinks(ctx, workItemID, filter) ([]WorkItemLink, *Pagination, error)`
  - 获取关联的工作项列表
  - 参数：workItemID（必填）、filter（可选的关联类型 ID）
  - 返回：关联列表和分页信息

- `Unlink(ctx, linkID) error`
  - 取消关联
  - 参数：linkID（必填）

#### 流程与交付目标类

- `ListTransitions(ctx, workItemID) (*WorkItemTransitionList, error)`
  - 获取工作项流转记录列表
  - 参数：workItemID（必填）
  - 返回：分页的流转记录列表

- `CreateDeliveryTarget(ctx, input) (*WorkItemDeliveryTarget, error)`
  - 创建工作项交付目标
  - 参数：input（包含工作项 ID、名称等必填字段）
  - 返回：创建的交付目标详情

- `UpdateDeliveryTarget(ctx, targetID, input) (*WorkItemDeliveryTarget, error)`
  - 部分更新工作项交付目标
  - 参数：targetID（必填）、input（包含所有可选字段）
  - 返回：更新后的交付目标详情

- `ListDeliveryTargets(ctx, filter) ([]WorkItemDeliveryTarget, *Pagination, error)`
  - 获取工作项交付目标列表
  - 参数：filter（包含工作项 ID、可选的状态）
  - 返回：交付目标列表和分页信息

- `DeleteDeliveryTarget(ctx, targetID) error`
  - 删除工作项交付目标
  - 参数：targetID（必填）

---

## 8. 版本历史

详细的版本变更记录请查看 [CHANGELOG.md](../CHANGELOG.md)。

以下为主要版本概览：

- **v0.7.0** (Phase 7) - 质量问题收敛
- **v0.6.0** (Phase 6) - 工作项领域接口（21个）
- **v0.5.0** (Phase 5) - 需求辅助接口（5个）
- **v0.4.0** (Phase 4) - Project领域CRUD + Global服务
- **v0.3.0** (Phase 3) - Ship领域接口
- **v0.2.0** (Phase 2) - 读取类接口完善
- **v0.1.0** (Phase 1) - 基础框架搭建

