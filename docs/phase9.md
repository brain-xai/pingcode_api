# Phase 9 迭代规划：Scrum 迭代管理（Sprint Management）

状态：✅ 已完成

目标：在已有 Project / Work Item 能力基础上，补齐 Scrum 迭代管理相关接口，使 SDK 能够围绕"迭代（Sprint）"这一核心概念进行完整的读写操作，包括迭代本体、迭代分组和迭代类别。

## 实现摘要

### 已完成功能

**1. 模型层 (sdk/model/sprint)**
- ✅ Sprint 模型（sprint.go）
- ✅ Sprint 输入模型（input.go）
- ✅ SprintGroup 模型（group.go）
- ✅ SprintCategory 模型（category.go）

**2. DTO 层 (internal/api/sprint)**
- ✅ Sprint DTO（sprint.go）
- ✅ SprintGroup DTO（group.go）
- ✅ SprintCategory DTO（category.go）

**3. Service 层 (sdk/service/sprint)**
- ✅ Service 结构（service.go）
- ✅ Sprint 管理方法（sprint.go + sprint_dto_mapper.go）
  - Create - 创建迭代
  - BatchCreate - 批量创建迭代
  - UpdatePartially - 部分更新迭代
  - List - 获取迭代列表
- ✅ SprintGroup 管理方法（group.go + group_dto_mapper.go）
  - CreateGroup - 创建分组
  - UpdateGroupPartially - 部分更新分组
  - ListGroups - 获取分组列表
  - DeleteGroup - 删除分组
- ✅ SprintCategory 管理方法（category.go + category_dto_mapper.go）
  - CreateCategory - 创建类别
  - UpdateCategoryPartially - 部分更新类别
  - ListCategories - 获取类别列表
  - DeleteCategory - 删除类别

**4. 客户端集成 (sdk/client.go)**
- ✅ 添加 Sprint() 方法

**5. 测试 (sdk/service/sprint)**
- ✅ sprint_test.go - Sprint 测试
- ✅ group_test.go - SprintGroup 测试
- ✅ category_test.go - SprintCategory 测试

**6. 示例代码**
- ✅ examples/sprint_management/main.go

---

---

## 1. 范围与目标

### 1.1 本阶段范围

围绕 Scrum 迭代管理，SDK 需要完成以下能力：

**迭代（Sprint）管理**

- 创建一个迭代
- 批量创建迭代
- 部分更新一个迭代
- 获取迭代列表

**迭代分组管理**

- 创建一个迭代分组
- 部分更新一个迭代分组
- 获取迭代分组列表
- 删除一个迭代分组

**迭代类别管理**

- 创建一个迭代类别
- 部分更新一个迭代类别
- 获取迭代类别列表
- 删除一个迭代类别

### 1.2 不在本阶段范围

- 迭代与工作项/交付目标等的高级联动行为（如自动计算燃尽图、智能排期等）
- 迭代下详细统计/报表接口（如有）
- 非 Scrum 类型项目的专用迭代行为（如与 Kanban 专属概念的绑定）

---

## 2. 领域划分与包结构

### 2.1 模块位置

结合现有项目结构，迭代管理相关代码建议组织为：

- `sdk/model/sprint`：与迭代相关的对外模型
  - Sprint、SprintGroup、SprintCategory
- `sdk/service/sprint`：迭代管理 Service
  - Sprint 管理
  - SprintGroup 管理
  - SprintCategory 管理
- `internal/api/sprint`：与 OpenAPI 对应的 DTO
  - 请求/响应结构

如 OpenAPI 中将迭代归类到 Project 模块，也依旧在 SDK 中使用 `sprint` 为独立领域，避免 `project` 包过度膨胀。

### 2.2 对外 Service 入口

在 `Client` 上提供统一入口，例如：

- `client.Sprint()`：返回 `*sprint.Service`

---

## 3. 模型设计（sdk/model/sprint）

### 3.1 Sprint 模型

需要定义的核心结构体：

- `Sprint`：
  - 字段示例（以 OpenAPI 为准）：
    - `ID`（SprintID）
    - `Name`
    - `ProjectID`
    - `CategoryID`（迭代类别）
    - `GroupID`（迭代分组）
    - `Goal`（迭代目标说明）
    - 时间：`StartAt`、`EndAt`
    - 状态：`Status`（planned/active/closed 等）
- `SprintFilter`：
  - 按项目、状态、时间范围等过滤
  - 支持分页参数（PageIndex、PageSize 等）

- `SprintCreateInput`：
  - 必填：名称、所属项目、时间范围
  - 选填：类别、分组、目标说明等

- `SprintPartialUpdateInput`：
  - 可更新字段：名称、时间、状态、类别/分组等
  - 使用指针字段或局部更新语义，避免与 Create 混淆

### 3.2 SprintGroup 模型

需要定义：

- `SprintGroup`：
  - `ID`（SprintGroupID）
  - `Name`
  - `ProjectID`
  - 可选说明字段
- `SprintGroupCreateInput`
- `SprintGroupPartialUpdateInput`
- `SprintGroupFilter`（如需要过滤特定项目下的分组）

### 3.3 SprintCategory 模型

需要定义：

- `SprintCategory`：
  - `ID`（SprintCategoryID）
  - `Name`
  - `ProjectID` 或适用范围
  - 可选说明字段
- `SprintCategoryCreateInput`
- `SprintCategoryPartialUpdateInput`
- `SprintCategoryFilter`

模型设计要求：

- ID 类型统一放入 core（如 `SprintID`、`SprintGroupID`、`SprintCategoryID`）
- 命名遵守统一命名规范（全部大写缩写等）

---

## 4. DTO 设计（internal/api/sprint）

### 4.1 Sprint 相关 DTO

在 `internal/api/sprint` 中定义与 OpenAPI 对齐的：

- `CreateSprintRequestDTO` / `CreateSprintResponseDTO`
- `BatchCreateSprintsRequestDTO` / `BatchCreateSprintsResponseDTO`
- `PatchSprintRequestDTO` / `PatchSprintResponseDTO`
- `ListSprintsRequestDTO` / `ListSprintsResponseDTO`

字段完全贴合 OpenAPI 描述：JSON tag、字段名、类型一致。

### 4.2 SprintGroup DTO

- `CreateSprintGroupRequestDTO` / `CreateSprintGroupResponseDTO`
- `PatchSprintGroupRequestDTO` / `PatchSprintGroupResponseDTO`
- `ListSprintGroupsRequestDTO` / `ListSprintGroupsResponseDTO`
- `DeleteSprintGroupRequestDTO` / `DeleteSprintGroupResponseDTO`（如 OpenAPI 对 DELETE 有 body，就定义请求 DTO）

### 4.3 SprintCategory DTO

- `CreateSprintCategoryRequestDTO` / `CreateSprintCategoryResponseDTO`
- `PatchSprintCategoryRequestDTO` / `PatchSprintCategoryResponseDTO`
- `ListSprintCategoriesRequestDTO` / `ListSprintCategoriesResponseDTO`
- `DeleteSprintCategoryRequestDTO` / `DeleteSprintCategoryResponseDTO`

---

## 5. model ↔ DTO 映射

### 5.1 映射原则

- DTO ↔ model 映射集中在 `sdk/service/sprint` 包下，单独 mapper 文件处理：
  - `sprint_dto_mapper.go`、`sprint_group_dto_mapper.go`、`sprint_category_dto_mapper.go`
- 不在业务函数里散写 JSON 字段映射。
- 所有新增字段在映射层都必须显式处理或明确注释忽略原因。

### 5.2 具体映射任务

- Sprint：
  - `CreateInput` → `CreateSprintRequestDTO`
  - `PartialUpdateInput` → `PatchSprintRequestDTO`
  - `Filter` → `ListSprintsRequestDTO`（包括分页/过滤参数）
  - `SprintDTO` → `Sprint`

- SprintGroup：
  - `SprintGroupCreateInput` → `CreateSprintGroupRequestDTO`
  - `SprintGroupPartialUpdateInput` → `PatchSprintGroupRequestDTO`
  - `SprintGroupDTO` → `SprintGroup`

- SprintCategory：
  - 类似上面映射关系

---

## 6. Service 接口设计与实现（sdk/service/sprint）

### 6.1 Service 结构

定义：

- `type Service struct { client *internalhttp.Client }`
- 在 `sdk/client.go`：
  - `func (c *Client) Sprint() *sprint.Service`

### 6.2 Sprint 管理方法

对外方法示例：

- `Create(ctx context.Context, input sprintmodel.SprintCreateInput) (*sprintmodel.Sprint, error)`
- `BatchCreate(ctx context.Context, inputs []sprintmodel.SprintCreateInput) ([]sprintmodel.Sprint, error)`
- `UpdatePartially(ctx context.Context, id core.SprintID, input sprintmodel.SprintPartialUpdateInput) (*sprintmodel.Sprint, error)`
- `List(ctx context.Context, filter sprintmodel.SprintFilter) ([]sprintmodel.Sprint, *core.Pagination, error)`

要求：

- 所有方法第一个参数为 `ctx context.Context`
- ID 为空、必填字段缺失等直接在本地报 Usage/Config 错误，不发送 HTTP 请求
- 错误返回遵循统一错误处理规范

### 6.3 SprintGroup 管理方法

对外方法示例：

- `CreateGroup(ctx context.Context, input sprintmodel.SprintGroupCreateInput) (*sprintmodel.SprintGroup, error)`
- `UpdateGroupPartially(ctx context.Context, id core.SprintGroupID, input sprintmodel.SprintGroupPartialUpdateInput) (*sprintmodel.SprintGroup, error)`
- `ListGroups(ctx context.Context, filter sprintmodel.SprintGroupFilter) ([]sprintmodel.SprintGroup, error)`
- `DeleteGroup(ctx context.Context, id core.SprintGroupID) error`

### 6.4 SprintCategory 管理方法

对外方法示例：

- `CreateCategory(ctx context.Context, input sprintmodel.SprintCategoryCreateInput) (*sprintmodel.SprintCategory, error)`
- `UpdateCategoryPartially(ctx context.Context, id core.SprintCategoryID, input sprintmodel.SprintCategoryPartialUpdateInput) (*sprintmodel.SprintCategory, error)`
- `ListCategories(ctx context.Context, filter sprintmodel.SprintCategoryFilter) ([]sprintmodel.SprintCategory, error)`
- `DeleteCategory(ctx context.Context, id core.SprintCategoryID) error`

方法命名风格需与现有 Project/Ship/WorkItem 模块保持一致（Create/UpdatePartially/List/Delete）。

---

## 7. 请求正确性与测试

### 7.1 请求构造测试

需要为以下方法编写请求构造单元测试：

- Sprint：
  - `Create` / `BatchCreate`：
    - 验证 URL、HTTP 方法、Body JSON 字段
  - `UpdatePartially`：
    - 验证 PATCH/PUT 路径正确、Body 仅包含部分字段
  - `List`：
    - 验证 Query 参数（项目、状态、时间范围、分页）正确序列化

- SprintGroup / SprintCategory：
  - 创建/更新/列表/删除全部需要测试：
    - URL path、HTTP 方法、Body/Query 正确

测试方式：

- 使用自定义 RoundTripper 捕获请求
- 断言请求的 Method、URL、Headers、Body JSON 与期望匹配

### 7.2 行为与错误测试

- 单元测试：
  - 参数校验（空 ID、缺少必填字段）
  - model ↔ DTO 映射逻辑
- 如有测试环境：
  - 集成测试：
    - 在测试项目中创建一个迭代
    - 批量创建几个迭代
    - 修改其中一个迭代的时间或状态
    - 查询迭代列表，确认数据正确
    - 创建迭代分组、类别并联动使用

---

## 8. 示例与文档

### 8.1 示例代码

新增示例：

- `examples/sprint_management/main.go`

示例流程建议：

1. 从环境变量读取 BaseURL、Token 和 ProjectID
2. 使用 `client.Sprint()`：
   - 创建一个迭代类别
   - 创建一个迭代分组
   - 在该分组/类别下创建一个迭代
   - 查询迭代列表并打印关键信息
   - 部分更新迭代（例如修改时间或目标）

要求：

- 示例可编译运行（配置正确的前提下）
- 错误处理清晰，便于调试

### 8.2 文档更新

- 在 SDK 文档中增加「迭代管理」章节：
  - 描述迭代、迭代分组、迭代类别的概念和典型使用场景
  - 提供对应示例代码片段
- 若有公共说明文档（如规范/契约文档），可补充 Sprint 模块相关章节（非必须本阶段完成）

---

## 9. Phase 9 验收标准

本阶段完成的判定标准：

1. 迭代（Sprint）、迭代分组、迭代类别的模型和 Service 方法全部实现：
   - 对应“创建 / 批量创建 / 部分更新 / 列表 / 删除”（删除仅适用于分组和类别）的主流程
2. 所有相关 HTTP 请求与 OpenAPI 保持一致：
   - 方法、路径、query、body 通过单元测试验证
3. 错误处理合理：
   - 对参数错误进行本地校验
   - 认证错误、业务错误、网络错误使用统一方式返回
4. 测试：
   - `go test ./...` 通过
   - Sprint 模块有覆盖映射、请求构造、核心行为的单元测试
5. 示例：
   - Sprint 管理示例能够在配置正确的环境中成功跑通“创建迭代类别/分组 → 创建迭代 → 列表查询 → 部分更新”链路

满足以上条件，即认为 Phase 9（Scrum 迭代管理相关接口）迭代目标完成。