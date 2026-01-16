# Phase 6 迭代规划：工作项（Work Item）领域接口

目标：在 Project / Requirement 能力的基础上，**系统完成“工作项（Work Item）相关接口”**，覆盖从创建、更新、删除，到属性、标签、关联、流转记录、交付目标的完整读写闭环，使 SDK 能够作为工作项管理的一站式访问层。

---

## 1. 范围与目标

### 1.1 本阶段范围

围绕“工作项”领域完成以下接口（与 OpenAPI 一一对应）：

**基础操作**

- 创建一个工作项
- 部分更新一个工作项
- 批量部分更新工作项属性
- 获取工作项列表
- 删除一个工作项

**属性与分类管理**

- 获取工作项类型列表
- 获取工作项状态列表
- 获取工作项属性列表
- 获取工作项优先级列表
- 获取工作项标签列表

**标签管理**

- 向工作项中添加一个标签
- 在工作项中移除一个标签

**关联管理**

- 获取工作项关联类型列表
- 关联一个工作项
- 获取关联的工作项列表
- 取消关联一个工作项

**流程与交付目标**

- 获取工作项流转记录列表
- 创建一个工作项交付目标
- 部分更新一个工作项交付目标
- 获取工作项交付目标列表
- 删除一个工作项交付目标

### 1.2 不在本阶段范围

- 工作项下更细粒度的子资源（评论、附件、时间记录等）
- 高级功能：批量导入导出、复杂报表、统计分析
- 跨系统同步/集成等更高层的业务封装

---

## 2. 领域拆分与接口分组

### 2.1 领域划分

在 SDK 中，建议将“工作项”作为独立子领域组织：

- `sdk/model/workitem`
- `sdk/service/workitem`
- `internal/api/workitem`

内部再按子功能分组（以文件或子类型区分，不强制子包）：

- 基础操作：`workitem_core.go`
- 属性与分类：`workitem_meta.go`
- 标签管理：`workitem_tags.go`
- 关联管理：`workitem_links.go`
- 流程与交付目标：`workitem_flow.go` / `workitem_delivery.go`

### 2.2 对外接口清单（SDK 视角）

结合 OpenAPI，Phase 6 在 SDK 级计划提供大致如下方法（命名示例）：

**基础操作**

- `Create(ctx, input WorkItemCreateInput) (*WorkItem, error)`
- `UpdatePartially(ctx, id WorkItemID, input WorkItemPartialUpdateInput) (*WorkItem, error)`
- `BatchUpdatePartially(ctx, req WorkItemBatchPartialUpdateInput) (*WorkItemBatchPartialUpdateResult, error)`
- `List(ctx, filter WorkItemFilter) ([]WorkItem, *Pagination, error)`
- `Delete(ctx, id WorkItemID) error`

**属性与分类**

- `ListTypes(ctx, scope WorkItemTypeScope) ([]WorkItemType, error)`
- `ListStatuses(ctx, filter WorkItemStatusFilter) ([]WorkItemStatus, error)`
- `ListFields(ctx, filter WorkItemFieldFilter) ([]WorkItemField, error)`
- `ListPriorities(ctx) ([]WorkItemPriority, error)`
- `ListTags(ctx, filter WorkItemTagFilter) ([]WorkItemTag, error)`

**标签管理**

- `AddTag(ctx, id WorkItemID, tagID WorkItemTagID) error`
- `RemoveTag(ctx, id WorkItemID, tagID WorkItemTagID) error`

**关联管理**

- `ListLinkTypes(ctx) ([]WorkItemLinkType, error)`
- `Link(ctx, fromID WorkItemID, toID WorkItemID, linkTypeID WorkItemLinkTypeID) error`
- `ListLinkedWorkItems(ctx, id WorkItemID, filter WorkItemLinkFilter) ([]WorkItemLink, error)`
- `Unlink(ctx, linkID WorkItemLinkID) error`  
  或基于（fromID, toID, linkTypeID）组合视 OpenAPI 定义

**流程与交付目标**

- `ListTransitions(ctx, id WorkItemID) ([]WorkItemTransition, error)`
- `CreateDeliveryTarget(ctx, id WorkItemID, input WorkItemDeliveryTargetCreateInput) (*WorkItemDeliveryTarget, error)`
- `UpdateDeliveryTargetPartially(ctx, targetID WorkItemDeliveryTargetID, input WorkItemDeliveryTargetPartialUpdateInput) (*WorkItemDeliveryTarget, error)`
- `ListDeliveryTargets(ctx, id WorkItemID, filter WorkItemDeliveryTargetFilter) ([]WorkItemDeliveryTarget, error)`
- `DeleteDeliveryTarget(ctx, targetID WorkItemDeliveryTargetID) error`

具体签名需根据 OpenAPI 细节微调，但整体风格在本阶段固定下来。

---

## 3. 模型设计（sdk/model/workitem）

### 3.1 核心工作项模型

需要完成：

- 在 `sdk/model/workitem` 中定义：

  - `WorkItem`：
    - 核心字段：
      - `ID` (`WorkItemID`)
      - `Title`
      - `Description`
      - `Type`（引用 `WorkItemType` 或 `WorkItemTypeID`）
      - `Status`（`WorkItemStatus`）
      - `Priority`（`WorkItemPriority`）
      - `AssigneeID` (`UserID`)
      - `ReporterID` (`UserID`)
      - 关联项目/产品/需求等 ID
      - 时间字段：`CreatedAt`、`UpdatedAt`、`StartedAt`、`CompletedAt` 等（视 OpenAPI）
  - `WorkItemFilter`：
      - 过滤条件：项目/产品/需求、类型、状态、优先级、标签、负责人、关键字、时间范围、分页等
  - `WorkItemCreateInput`：
      - 必填：标题、项目/产品/类型等
      - 选填：描述、优先级、负责人、初始状态、标签等
  - `WorkItemPartialUpdateInput`：
      - 可更新的字段集合，支持“部分更新”（指针字段或显式 patch 语义）

### 3.2 属性与分类模型

需要完成：

- `WorkItemType`, `WorkItemStatus`, `WorkItemField`, `WorkItemPriority`, `WorkItemTag` 等结构体：
  - 统一放在 `sdk/model/workitem`，命名清晰、语义稳定
  - 使用 `WorkItemTypeID`、`WorkItemStatusID` 等 ID 类型（位于 `sdk/model/core`）

### 3.3 标签与关联模型

需要完成：

- 标签：
  - `WorkItemTagID`
  - `WorkItemTag`（名称、颜色、描述等）
  - `WorkItemTagFilter`（按项目/产品/工作项类型等过滤）

- 关联：
  - `WorkItemLinkType`、`WorkItemLinkTypeID`
  - `WorkItemLink`：包含 from/to 工作项、类型、方向、创建信息等
  - `WorkItemLinkFilter`：按类型/方向过滤关联项

### 3.4 流转记录与交付目标模型

需要完成：

- 流转记录：
  - `WorkItemTransition`：包含从状态、到状态、操作者、时间等

- 交付目标：
  - `WorkItemDeliveryTarget` / `WorkItemDeliveryTargetID`
  - `WorkItemDeliveryTargetCreateInput`
  - `WorkItemDeliveryTargetPartialUpdateInput`
  - `WorkItemDeliveryTargetFilter`

模型设计需遵守：

- 命名规范（见 `deployment.md`）
- 对外不暴露 OpenAPI 的“内部字段命名”，而是封装为业务友好的字段

---

## 4. DTO 与映射（internal/api/workitem）

### 4.1 DTO 定义

需要完成：

- 在 `internal/api/workitem` 中，为上文提到的每类接口定义请求/响应 DTO：

  - 基础操作 DTO（Create / Patch / BatchPatch / List / Delete）
  - 属性与分类相关 DTO（Type/Status/Field/Priority/Tag）
  - 标签增删 DTO
  - 关联类型与关联管理 DTO
  - 流转记录与交付目标 DTO

- DTO 字段和 JSON tag 必须与 OpenAPI 文档保持一致，适合被生成工具接管。

### 4.2 model ↔ api 映射

需要完成：

- 在 `sdk/service/workitem` 或独立 mapper 文件中实现映射：

  - DTO → 模型：
    - `WorkItemDTO` → `WorkItem`
    - 各种 “meta” DTO → 对应模型（Type/Status/Field/Priority/Tag/Link/Transition/DeliveryTarget）
  - 模型/输入 → DTO：
    - `WorkItemCreateInput` → `CreateWorkItemRequestDTO`
    - `WorkItemPartialUpdateInput` → `PatchWorkItemRequestDTO`
    - `WorkItemBatchPartialUpdateInput` → 对应批量 DTO
    - 各 Filter → 相应 List 请求 DTO

要求：

- 映射集中管理、可单元测试
- 保证不静默丢掉重要字段（除非明确设计为隐藏字段）

---

## 5. Service 实现（sdk/service/workitem）

### 5.1 Service 结构与 Client 集成

需要完成：

- 在 `sdk/service/workitem` 中定义：

  - `Service` 结构体（持有 HTTP Client）
  - 分文件实现不同子功能（基础操作、meta、标签、关联、交付目标）

- 在 `sdk/client.go` 中：
  - 增加 `WorkItem() *workitem.Service` 方法，将 Service 挂载到顶层 Client

### 5.2 核心方法实现

按分组实现对应的方法（见 2.2），要求：

- 使用 `context.Context`
- 参数在进入 HTTP 层前进行校验（ID/必填字段）
- 错误通过统一错误类型返回（与 `error_model.md` 一致）

注意：

- 部分更新与批量更新方法需要清晰区分：
  - 单个部分更新：明确哪些字段会被覆盖
  - 批量更新：要求调用方显式指定更新范围/目标集合，避免歧义

---

## 6. 请求正确性与测试（工作项专项）

### 6.1 请求构造单元测试

需要完成：

- 对所有核心工作项接口编写请求构造测试，包括但不限于：

  - 创建/部分更新/批量更新/删除工作项
  - 获取列表（含过滤与分页）
  - 标签增删
  - 关联增删与关联列表查询
  - 流转记录与交付目标相关接口

测试验证：

- HTTP 方法正确（GET/POST/PATCH/DELETE 等）
- URL path 正确替换 path 参数
- Query 参数正确序列化
- Body JSON 结构符合 OpenAPI 要求（字段名、嵌套结构）

### 6.2 行为与错误测试

需要完成：

- 单元测试：
  - 参数校验分支（空 ID、缺字段、非法组合）
  - model ↔ DTO 映射逻辑
- 集成测试（视环境能力）：
  - 至少覆盖以下场景：
    1. 创建一个工作项
    2. 部分更新该工作项
    3. 为其添加标签、创建关联、创建交付目标
    4. 获取流转记录和交付目标列表
    5. 删除交付目标或工作项（视业务允许）

验收标准：

- `go test ./...` 通过
- 工作项模块主要路径和错误路径都有覆盖

---

## 7. 示例与文档

### 7.1 示例代码

需要完成：

- 新增示例：`examples/workitems/main.go`（或类似）

示例流程：

1. 从环境变量读取 BaseURL 与 Token
2. 选取一个项目/产品上下文（可通过 Project/Ship Service 获取）
3. 创建一个工作项
4. 列出工作项列表，找到刚创建的项
5. 部分更新其若干字段（如优先级）
6. 添加标签、关联另一个工作项
7. 查询流转记录和交付目标列表（如 OpenAPI 支持）
8. 清理：可选删除测试工作项/交付目标

### 7.2 文档更新

需要完成：

- 在 SDK 文档中新增「工作项管理」章节：
  - 如何创建/更新/删除工作项
  - 如何管理属性、标签、关联、交付目标
  - 示例调用代码片段

- 更新相关规范：
  - `api_contract.md`：新增工作项相关对外 API
  - `error_model.md`：如有工作项特有错误（例如类型/状态非法、不可删除等），补充说明

---

## 8. Phase 6 验收标准

本阶段完成的判定标准：

1. 工作项领域主要接口全部有 SDK 封装：
   - 基础 CRUD / 列表 / 批量部分更新
   - 类型/状态/属性/优先级/标签列表
   - 标签增删
   - 关联类型、关联增删与关联列表
   - 流转记录与交付目标 CRUD/列表
2. 请求行为正确：
   - 方法、路径、query、body 与 OpenAPI 定义一致（通过单元测试验证）
3. 错误处理符合统一规范：
   - 参数错误、认证错误、业务错误、网络错误可区分
4. 测试与示例：
   - `go test ./...` 通过
   - 工作项模块有覆盖映射、请求构造、行为的单元测试
   - 示例程序在正确配置下，能跑通“创建 → 更新 → 标签/关联/交付目标 → 查询”的关键场景

满足上述条件，即认为 Phase 6（工作项相关接口）迭代目标完成，为后续扩展子资源和高级能力打下坚实基础。