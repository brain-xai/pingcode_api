# Phase 3 迭代规划：Ship 领域接口

目标：在 Phase 1、Phase 2 已完成基础框架和 GET 能力的基础上，**专注完成 Ship 领域（产品/需求等）的主要接口实现**，并通过示例和测试验证可用性。

---

## 1. 范围与目标

### 1.1 本阶段范围

围绕 Ship 领域，完成：

1. **Ship 领域模型与 Service 完整化**
   - `sdk/model/ship`：核心模型（Product、Requirement 等）
   - `sdk/service/ship`：核心读取与关键写接口
   - `internal/api/ship`：与 OpenAPI 对齐的 DTO

2. **主要接口能力**
   - 产品（Product）相关：
     - 获取单个产品详情（Get）
     - 获取产品列表（List）
   - 需求（Requirement / Story 等）相关（具体名称以 OpenAPI 为准）：
     - 获取单个需求详情（Get）
     - 获取需求列表（List）
     - 创建需求（Create）
     - 更新需求的关键字段（例如状态/标题/描述等）

3. **请求正确性保障（针对 Ship）**
   - Ship 相关所有 HTTP 请求的：
     - 方法、路径、Query、Headers 与 OpenAPI 保持一致
   - 参数校验和错误模型集成

4. **示例与文档**
   - 至少提供一个完整示例：
     - “为某个产品拉取需求列表，并创建一个新需求”
   - 更新文档中和 Ship 领域相关的章节

### 1.2 不在本阶段范围

- Ship 领域中非常长尾的接口（如极少使用的统计、报表类接口）
- 高级能力：批量导入、复杂查询构建器等
- 其它领域（Project 以外的新增写接口）仅作为下一阶段内容

---

## 2. Ship 领域分析与接口清单

### 2.1 阅读与梳理 OpenAPI 文档

**目标：** 从 `docs/reference/openpingcode/ship.*`（或对应 OpenAPI 文件）中提取 Ship 领域核心接口。

需要完成：

1. 梳理 Ship 领域主要资源：
   - Product / 产品
   - Requirement / 需求（故事/用例，具体以接口命名为准）
   - 版本 / Release（如有）
2. 按资源整理以下信息：
   - 资源名称（如 Product）
   - 相关的 GET/POST/PUT 等接口路径
   - 必需请求参数、常用过滤条件
   - 核心响应字段

产出：一份内部用的 Ship 接口清单（可附在本文件末尾或单独文档）。

### 2.2 Phase 3 优先接口列表

在所有 Ship 接口中，优先实现：

1. **产品 Product**
   - `GetProductByID`
   - `ListProducts`
2. **需求 Requirement**
   - `GetRequirementByID`
   - `ListRequirements`（按产品/项目等维度过滤）
   - `CreateRequirement`
   - `UpdateRequirement`（至少支持更新状态和基础信息）

如资源命名在 OpenAPI 中不同（如 Idea/Epic 等），以实际名称为准，但遵循 `deployment.md` 中的命名规范。

---

## 3. Ship 模型设计（sdk/model/ship）

### 3.1 Product 模型

需要完成：

- 在 `sdk/model/ship` 中定义 Product 相关结构体：
  - `Product`：包含 ID、名称、Key、状态、所属空间/项目等核心字段
  - 视需要增加：
    - `ProductSummary`（轻量信息）
    - `ProductFilter`（用于列表查询）
- 复用 `sdk/model/core` 中的 ID、时间等类型

### 3.2 Requirement 模型

需要完成：

- 在 `sdk/model/ship` 中定义 Requirement 相关结构体：
  - `Requirement`：包含 ID、标题、描述、状态、优先级、所属产品等
  - `RequirementFilter`：用于列表过滤（按产品、状态、负责人等）
  - `RequirementCreateInput`：创建需求入参
  - `RequirementUpdateInput`：更新需求入参（至少覆盖 Phase 3 关注字段）

设计原则：

- 字段命名贴合业务含义，避免直接照搬 API 的内部术语
- 所有对外类型遵守 `deployment.md` 中的命名规范和 `api_contract.md` 的契约要求

---

## 4. Ship DTO 与映射（internal/api/ship）

### 4.1 DTO 定义

需要完成：

- 在 `internal/api/ship` 中，为 Ship 领域接口定义 DTO：
  - Product 请求/响应 DTO
  - Requirement 请求/响应 DTO（列表、详情、创建、更新）
- 保证：
  - JSON tag 与 OpenAPI 文档一致
  - 可选字段用指针或合适的零值处理

### 4.2 model ↔ api 映射

需要完成：

- 在 Ship service 或专门的 mapper 中实现：
  - `internal/api` DTO → `sdk/model/ship` 模型的转换
  - `sdk/model/ship` 输入结构体 → `internal/api` 请求 DTO

要求：

- 不静默丢字段，任何“丢弃/默认”的行为在代码中清晰可见
- 映射逻辑有单元测试覆盖

---

## 5. Ship Service 实现（sdk/service/ship）

### 5.1 Service 结构与注入

需要完成：

- 在 `sdk/service/ship` 中：
  - 定义 `Service` 结构体（持有 HTTP Client 等）
  - 在 `sdk/client.go` 中提供 `Ship()` 方法返回该 Service

### 5.2 读取接口（GET）

需要完成以下方法（命名示例，以实际命名为准）：

- 产品：
  - `GetProduct(ctx context.Context, id core.ProductID) (*model.Product, error)`
  - `ListProducts(ctx context.Context, filter model.ProductFilter) ([]model.Product, *core.Pagination, error)`
- 需求：
  - `GetRequirement(ctx context.Context, id core.RequirementID) (*model.Requirement, error)`
  - `ListRequirements(ctx context.Context, filter model.RequirementFilter) ([]model.Requirement, *core.Pagination, error)`

要求：

- 使用 `context.Context`
- 参数校验（ID 不为空，必填过滤条件非空等）
- 错误使用统一错误模型（参考 `error_model.md`）

### 5.3 写接口（Create/Update）

需要完成：

- `CreateRequirement(ctx context.Context, input model.RequirementCreateInput) (*model.Requirement, error)`
- `UpdateRequirement(ctx context.Context, id core.RequirementID, input model.RequirementUpdateInput) (*model.Requirement, error)`

要求：

- 对入参进行本地校验（必填字段）
- 正确映射为 `internal/api` DTO，并调用对应的 API
- 返回创建/更新后的最新 Requirement

---

## 6. 请求正确性与测试

### 6.1 请求构造校验

需要完成：

- 为 Ship 相关 Service 方法编写单元测试，验证：
  - HTTP 方法为 GET/POST/PUT 等与 OpenAPI 一致
  - 路径（含 path param）正确
  - Query 参数序列化正确（分页、过滤）
  - 必要 headers（如 Authorization）正确设置

### 6.2 行为测试

需要完成：

- 单元测试：
  - 对每个 Ship 方法的参数校验、映射逻辑进行测试
- 集成测试（如有测试环境或可靠 mock）：
  - 至少为以下方法编写集成测试：
    - `ListProducts`
    - `ListRequirements`
    - `CreateRequirement`（可在测试环境创建后立即删除或标记）

---

## 7. 示例与文档

### 7.1 示例代码

需要完成：

- 新增或扩展示例目录，例如：
  - `examples/ship_requirements/main.go`
- 示例流程：
  1. 从环境变量读取 BaseURL 与 Token
  2. 使用 SDK 初始化 Client
  3. 使用 `Ship().ListProducts` 获取一个产品
  4. 使用 `Ship().ListRequirements` 获取该产品下的需求列表
  5. 使用 `Ship().CreateRequirement` 为该产品创建一个新需求（可选）

### 7.2 文档更新

需要完成：

- 在 SDK 文档中新增「Ship 领域指南」小节：
  - 如何获取产品列表、需求列表
  - 如何创建与更新需求
- 若新增 API 进入稳定契约范围：
  - 更新 `api_contract.md`（列出新的对外 API）
  - 根据实际错误行为更新 `error_model.md`（如新增特定错误码或错误类型说明）

---

## 8. Phase 3 验收标准

本阶段完成的判定标准：

1. Ship 领域核心资源（Product、Requirement）的模型和 Service 实现完整：
   - 包含 Get/List/Create/Update 的主流程
2. 所有 Ship 方法的请求与 OpenAPI 定义保持一致：
   - HTTP 方法、路径、参数、headers 经测试验证
3. 错误处理符合统一错误模型：
   - 参数错误、认证错误、业务错误和网络错误可区分
4. 测试：
   - `go test ./...` 通过
   - Ship 模块的映射、请求构造、关键行为有单元测试覆盖
5. 示例：
   - Ship 示例能够在配置正确的情况下运行，并展示获取产品/需求列表（以及可选的创建需求）

满足以上条件，即可认为 Phase 3 Ship 相关接口迭代完成，可以进入后续阶段（如其它领域写接口、自动生成工具完善等）。

---

## 9. 实施记录

### 9.1 已完成内容

#### 9.1.1 Model 层
- ✅ `sdk/model/ship/product.go`：Product、ProductMember、ProductFilter 模型
- ✅ `sdk/model/ship/requirement.go`：Requirement、ProductSummary、RequirementState、RequirementPriority、Plan、Suite、TimeRange、Participant、RequirementFilter、RequirementCreateInput、RequirementUpdateInput 模型

#### 9.1.2 DTO 层
- ✅ `internal/api/ship/product.go`：Product DTO 和 ToModel 转换方法
- ✅ `internal/api/ship/requirement.go`：Requirement DTO 和 ToModel 转换方法，包括 CreateRequirementRequest 和 UpdateRequirementRequest

#### 9.1.3 HTTP Client 扩展
- ✅ `internal/httpclient/client.go`：添加 Patch 和 PatchWithPathParams 方法

#### 9.1.4 Service 层
- ✅ `sdk/service/ship/service.go`：实现以下方法
  - ListProducts
  - GetProduct
  - ListRequirements
  - GetRequirement
  - CreateRequirement
  - UpdateRequirement

#### 9.1.5 Client 集成
- ✅ `sdk/client.go`：添加 Ship() 方法

#### 9.1.6 测试
- ✅ `sdk/service/ship/service_test.go`：完整单元测试
  - 参数验证测试
  - HTTP 请求正确性测试
  - DTO 转换测试

#### 9.1.7 示例代码
- ✅ `examples/ship_products/main.go`：产品操作示例
- ✅ `examples/ship_requirements/main.go`：需求操作完整示例

### 9.2 测试结果

```
ok  	github.com/brain-xai/pingcode_api/sdk/service/ship	0.014s
```

所有测试通过，包括：
- TestShipService_GetProduct_ParameterValidation
- TestShipService_GetProduct_Success
- TestShipService_ListProducts_URLConstruction
- TestShipService_GetRequirement_ParameterValidation
- TestShipService_GetRequirement_Success
- TestShipService_ListRequirements_URLConstruction
- TestShipService_CreateRequirement_ParameterValidation
- TestShipService_CreateRequirement_Success
- TestShipService_UpdateRequirement_ParameterValidation
- TestShipService_UpdateRequirement_Success

### 9.3 验收确认

- ✅ Ship 领域核心资源（Product、Requirement）的模型和 Service 实现完整
- ✅ 所有 Ship 方法的请求与 OpenAPI 定义保持一致
- ✅ 错误处理符合统一错误模型
- ✅ `go test ./...` 通过
- ✅ Ship 示例能够在配置正确的情况下运行

**Phase 3 已于 2025-01-15 完成**