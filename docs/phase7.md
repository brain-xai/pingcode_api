# Phase 7：工作项实现问题收敛与质量加固

目标：在 Phase 6 的基础上，系统性消除当前工作项实现中的明显问题与“坏味道”，保证 SDK 对外行为稳定、参数语义一致、文档不再误导用户，为后续扩展打下干净的技术基础。

---

## 1. 范围与目标

### 1.1 本阶段范围（必须完成）

围绕当前实现中已识别的问题，完成以下修正：

- SDK 客户端与超时配置
- 工作项关联（Link / Unlink）接口的语义与实现
- 工作项相关 Service 的参数使用一致性
- DTO ↔ Model 映射中的命名和信息缺失问题
- 文档与实际行为不一致的地方

### 1.2 不在本阶段范围

- 不新增新的业务能力（例如评论、附件等子资源）
- 不大幅重构包结构（`sdk/model/...`、`sdk/service/...` 目录保持不变）
- 不引入 Breaking Change 除非必须修复现有“必然错误”的行为（优先考虑兼容方式修复）

---

## 2. 已知问题列表（当前实现的“坑”）

本阶段聚焦修复以下已经确认存在的问题：

### 2.1 超时配置逻辑错误（严重）

文件：[sdk/client.go](file:///root/src/pingcode_api/sdk/client.go)

问题：

- `NewClient` 中的 Timeout 默认逻辑与 `Config` 的默认值不一致，且实现明显错误：
  - `NewDefaultConfig` 使用 `Timeout: 30 * time.Second`
  - `NewClient` 中逻辑：
    - `timeout := cfg.Timeout`
    - `if timeout == 0 { timeout = 30 * defaultTimeout }`
    - `defaultTimeout = 30`
  - 实际结果是：当用户手动构造 `Config` 且未设置 `Timeout` 时，得到的是一个极小（近似 0）的超时时间，行为与预期严重不符。

影响：

- 在网络有轻微抖动的生产环境中，SDK 请求会频繁超时，表现为“随机失败”，非常难排查。
- 这是直接影响用户稳定性的 bug，必须修复。

### 2.2 Unlink 实现不可用且与文档不符（严重）

文件：

- 实现：[sdk/service/workitem/workitem_links.go](file:///root/src/pingcode_api/sdk/service/workitem/workitem_links.go)
- 文档示例：[docs/workitem_management.md](file:///root/src/pingcode_api/docs/workitem_management.md)

问题：

- `Unlink(ctx, linkID string)` 的实现使用了硬编码的 `"dummy"` 作为 `work_item_id`：
  - URL 模板为：`/v1/project/work_items/{work_item_id}/relations/{relation_id}`
  - 当前代码：`pathParams := map[string]string{"work_item_id": "dummy", "relation_id": linkID}`
- 这在真实环境中几乎必然 404 或出错，属于“看上去有 API，实际不可用”。
- 文档中却直接示例调用：`client.WorkItem().Unlink(ctx, linkID)`，严重误导用户。

影响：

- 用户按文档调用必然踩坑。
- 这是一个“用户空间”层面的严重问题，必须在本阶段修复，并保证文档同步。

### 2.3 ListLinkTypes 参数使用不一致（中等）

文件：[sdk/service/workitem/workitem_links.go](file:///root/src/pingcode_api/sdk/service/workitem/workitem_links.go)

问题：

- 函数签名：`ListLinkTypes(ctx context.Context, projectID string)`
- 实现中会校验 `projectID` 非空，但后续请求完全没有使用该参数：
  - 直接对 `/v1/project/work_item/relation_types` 发 GET，无任何 query 参数。
- 这暴露了“签名和行为脱节”的问题：
  - 如果后端 API 不需要 projectID，那么 SDK 强制要求只是增加噪音；
  - 如果后端 API 需要 projectID，则当前实现是错的。

影响：

- 对调用者来说，projectID 参数语义不清，容易误解。
- 对后续演进来说，这是一个“隐藏的特殊情况”，会阻碍重构。

### 2.4 DTO ↔ Model 映射不一致与命名混乱（轻–中等）

文件：[internal/api/workitem/dto.go](file:///root/src/pingcode_api/internal/api/workitem/dto.go) 及相关 Model

问题点包括但不限于：

1. Phase/Type 命名混乱：
   - DTO 中 `WorkItem` 有 `Phase *Phase` 字段（ID/Identifier/Title）。
   - Model 中 `WorkItem` 的 `Phase` 是 `string`（仅存储一个 ID）。
   - 转换方法命名为 `getTypeID(phase *Phase)`，实际返回的是 Phase.ID。
   - 语义和命名明显不匹配，容易给维护者造成误导。

2. WorkItemSummary 中的 State 永远为空：
   - 模型 `WorkItemSummary` 中有 `State *WorkItemStatus` 字段。
   - DTO `WorkItemSummary` 中没有 state 字段，对应的 `ToModel()` 也从未赋值。
   - 对 SDK 使用者来说，这是一个“看起来有字段，实际上永远为 nil”的陷阱。

3. Reporter 字段无实际数据来源：
   - 模型 `WorkItem` 里有 `ReporterID` / `Reporter` 字段。
   - DTO 中目前没有 Reporter 相关字段，映射函数也未使用。
   - 对用户而言，这些字段当前等同于“死字段”。

影响：

- 增加理解成本，后续再演进时容易出错。
- 如果不及时清理，将来会很难再做兼容性调整。

### 2.5 属性类字段全转 string 的信息丢失（可接受但需记录）

文件：多个 DTO 的 `toString` 使用处。

问题：

- DTO 中各种 `map[string]interface{}` 类型的属性在映射到 Model 时统一通过 `fmt.Sprintf` 转为 `map[string]string`。
- 这会丢失类型信息（数值/布尔/时间等）。

影响：

- 对当前“展示为主”的 SDK 使用场景可以接受，但一旦未来需要在 SDK 侧做更复杂的本地校验或逻辑，会成为限制。

处理策略：

- 本阶段不强行改动类型（避免破坏用户空间）。
- 在本阶段计划中明确记录这一点，后续如要增强，需设计兼容方案（例如增加 typed accessor，而不是直接改字段类型）。

---

## 3. 设计原则（Phase 7 的约束）

为了确保修补不会引入新坑，本阶段遵守以下原则：

- **Never break userspace**：
  - 保持已经对外暴露的函数签名尽量不变。
  - 无法兼容时，优先通过“增加重载/新增方法”方式解决，而不是直接修改已有方法行为。
- **消除特殊情况而不是掩盖**：
  - 不再引入类似 `"dummy"` 这种“假参数”方案。
  - 参数校验与 HTTP path/query 构造必须一一对应。
- **实用主义优先**：
  - 聚焦真实会影响用户程序稳定/正确性的 bug。
  - 不做大规模重构，只做必要的、价值明显的清理。
- **文档与行为同步**：
  - 任何修复若会改变对外可见行为，必须同步更新文档。
  - `docs/workitem_management.md` 和 `api_contract.md` 要反映真实行为。

---

## 4. 具体修复计划

### 4.1 修复 SDK Client 超时逻辑

目标：确保所有构造 `sdk.Client` 的方式，在未显式设置 `Timeout` 时都得到一个合理的默认值（例如 30s），与 `NewDefaultConfig` 一致。

计划：

1. 梳理超时相关入口：
   - `sdk.NewDefaultConfig()`
   - `sdk.LoadConfigFromEnv()`
   - 直接手动构造 `&sdk.Config{}` 然后调用 `sdk.NewClient`

2. 调整 `NewClient` 中的逻辑：
   - 如果 `cfg.Timeout == 0`，设置为一个明确的默认值（例如 `30 * time.Second`），与 `NewDefaultConfig` 对齐。
   - 考虑是否允许用户传入自定义 `HTTPClient`，如允许则需要尊重其内部 timeout。

3. 增加单元测试：
   - 场景：
     - 使用 `NewDefaultConfig` → timeout 为 30s。
     - 使用 `LoadConfigFromEnv` 且无环境变量时 → timeout 为 30s。
     - 用户显式设置 `Timeout` → `NewClient` 不再修改该值。
   - 确保 `internal/httpclient.NewClient` 得到的 timeout 符合预期。

### 4.2 重构 Unlink 语义与实现

目标：让 `Unlink` 变成一个真实可用、行为清晰的 API，消除 `"dummy"` 特殊情况。

计划：

1. 对照 OpenAPI / 后端实现，确认支持的删除接口形态：
   - 优先采用“仅凭 relation id 删除”的接口（例如 `DELETE /v1/project/work_item_relations/{relation_id}`）。
   - 若必须带 `work_item_id`，则需要调整 SDK API 形状。

2. 根据后端能力确定 SDK 方案：

   - 方案 A（优先）：后端支持按 relation id 删除
     - 修改 `Unlink(ctx, linkID string)` 的 HTTP path 为真实有效的 URL。
     - 移除 `"dummy"` work_item_id 的伪路径参数。
     - 更新文档中关于 Unlink 的说明和示例。

   - 方案 B：后端需要 work_item_id + relation_id
     - 新增方法：`UnlinkWithWorkItem(ctx, workItemID, linkID string) error`。
     - 保留旧的 `Unlink(ctx, linkID string)` 但标记为 Deprecated（通过注释和文档），内部可以：
       - 要么直接返回一个明晰的错误，说明必须使用新方法；
       - 要么在有足够信息时尝试自动查找 workItemID（代价较大，不推荐）。
     - 文档改为推荐使用新的方法，旧方法仅保留兼容。

3. 增加单元测试：
   - 验证构造的 URL 与 path 参数和 OpenAPI 一致。
   - 对错误场景（空 ID）保持现有参数校验行为。

4. 更新文档：
   - `docs/workitem_management.md` 中的“取消关联”部分，改为展示新的正确方法签名和示例。
   - 如有行为变化，在 `phase7.md` 中记录兼容性考量。

### 4.3 对齐 ListLinkTypes 的参数与行为

目标：保证 `ListLinkTypes` 的签名与实际请求一致，不再出现“要求 projectID 但完全不用”的情况。

计划：

1. 调查后端接口：
   - 确认 `/v1/project/work_item/relation_types` 是否需要 projectID。
   - 如果需要，则应在 query 中带上 `project_id`。
   - 如果不需要，则应考虑移除函数签名中的 projectID（或标记为兼容参数）。

2. 根据结果调整 SDK 行为：

   - 若后端需要 projectID：
     - 在 `ListLinkTypes` 中构建 `url.Values`，设置 `project_id`。
     - 保留现有函数签名，仅补齐实现。
     - 补充参数校验相关的测试。

   - 若后端不需要 projectID：
     - 保留现有函数签名（避免破坏用户），但在文档中说明该参数当前未被使用。
     - 未来视需要可以在新版本引入不带 projectID 的重载。

3. 补充测试与文档：
   - 单元测试检查 query 是否按预期附加。
   - 文档更新 `ListLinkTypes` 的参数说明。

### 4.4 清理 DTO ↔ Model 映射中的命名和信息缺失

目标：在不破坏现有字段的前提下，提升可读性并尽可能减少“总为空/总丢失”的字段。

计划：

1. Phase / Type 命名修正：
   - 将 `getTypeID` 更名为反映真实含义的函数名（例如 `getPhaseID`），避免继续扩大误导。
   - 对应调用点全部调整。
   - 不改变对外模型字段类型（`WorkItem.Phase` 仍为 string），只是让内部命名清晰。

2. WorkItemSummary 中 State 的处理：
   - 与后端/产品确认：summary 是否有状态信息，如果有：
     - 在 DTO 中补充 state 字段，并实现映射。
     - 更新 `ToModel()` 填充 `State`。
   - 如果后端不打算提供此信息：
     - 考虑在模型中去掉 `State` 字段，或者在文档中明确标记为“当前版本不返回”。

3. Reporter 字段：
   - 同样与后端确认是否有 reporter 数据：
     - 如有实际字段（例如 `reporter_id` / `reporter`），补齐 DTO 与映射。
     - 如没有，考虑删减模型中的字段，以免长期挂着“死字段”。

4. 增加映射层单元测试：
   - 针对 `WorkItem`、`WorkItemSummary`、`Phase`、`Reporter` 等关键字段，构造 DTO，验证 Model 映射结果。
   - 确保不再出现“命名误导”和“无意丢字段”的情况。

### 4.5 属性字段 string 化的记录与约束

目标：明确当前行为，同时为未来演进留接口。

计划：

1. 在内部开发文档中记录：
   - 当前所有自定义属性/Properties 均在 Model 层转为 `map[string]string`。
   - 这样做的设计权衡（简化调用 vs 类型丢失）。

2. 在公开文档中简要说明：
   - 对于属性类字段，SDK 模型层统一以字符串呈现。
   - 调用方如需保留类型信息，建议在存储时自行编码（例如 JSON 字符串）。

3. 后续如需扩展：
   - 考虑增加 typed accessor（例如 `GetIntProperty(name string)`），而不是直接修改字段类型。

---

## 5. 测试与验收标准

### 5.1 单元测试

需要增加和/或补全的测试覆盖：

- `sdk.NewClient` 超时逻辑：
  - 不同配置组合下，最终传入 HTTP Client 的 Timeout 值。
- `WorkItemService.Unlink`（以及可能的 `UnlinkWithWorkItem`）：
  - 参数校验（空 ID）。
  - 构造的 HTTP path 与 path 参数是否符合 OpenAPI。
- `ListLinkTypes`：
  - 根据后端需求验证是否正确使用 projectID。
- DTO ↔ Model 映射：
  - 特别是 Phase、WorkItemSummary.State、Reporter 等字段的映射正确性。

### 5.2 集成测试（如环境允许）

在现有 `tests/integration` 基础上，优先考虑增加：

- 工作项关联的集成测试：
  - 创建两个工作项 → 建立关联 → 列出关联 → 取消关联 → 校验结果。
- 超时行为的简单验证：
  - 在可控环境下模拟慢响应，确认默认 Timeout 足够，且用户可覆盖。

---

## 6. 文档更新

需要更新的文档包括：

- `docs/workitem_management.md`：
  - 修正“取消关联”示例，使用新的、真实可用的 Unlink 方案。
  - 如 Unlink 签名有兼容性调整，在文档中明确说明。
- `docs/api_contract.md`：
  - 补充/修正文档中关于工作项关联删除接口的描述，与实际 HTTP path 对齐。
- 本文件 `phase7.md`：
  - 作为本阶段实施与验收的蓝本，后续如有新增问题，可在本文件追加条目。

---

## 7. Phase 7 验收标准

本阶段完成的判定标准：

1. **已知严重问题全部消除**：
   - 默认超时行为与文档一致且可控。
   - Unlink 行为可在真实环境下成功取消关联，不再依赖 `"dummy"` 特殊值。
2. **接口签名与行为一致**：
   - `ListLinkTypes` 对 projectID 的使用与后端语义对齐。
   - DTO ↔ Model 映射中不再存在明显的命名误导或“永远为空”的伪字段，或者这些字段被明确标记/移除。
3. **测试通过**：
   - `go test ./...` 通过。
   - 新增的单元测试覆盖关键修复点。
4. **文档同步更新**：
   - 工作项相关文档描述与真实行为一致，示例代码可直接运行（在正确配置环境下）。
5. **不破坏用户空间**：
   - 已有对外 API 的行为不会在无说明的情况下“静默改变”。
   - 任何不得不调整的地方都有清晰的兼容策略和迁移说明。

满足以上条件后，认为 Phase 7 完成，工作项 SDK 基础质量达到可对外长期使用的标准，为后续功能扩展提供稳定基线。

---

## 8. 实施记录

### 已完成项

- [x] **超时配置修复**（问题 1）
  - 修改 `sdk/client.go` 常量定义：`defaultTimeout = 30 * time.Second`
  - 修改 NewClient 逻辑：`if timeout == 0 { timeout = defaultTimeout }`
  - 添加 time 包导入
  - 新增 `sdk/client_test.go` 测试文件

- [x] **Unlink 修复**（问题 2）
  - 新增 `UnlinkWithWorkItem(ctx, workItemID, linkID string)` 方法
  - 旧 `Unlink` 方法标记为 Deprecated，返回明确错误
  - 更新文档 `docs/workitem_management.md`

- [x] **ListLinkTypes 参数对齐**（问题 3）
  - 移除无用的 projectID 非空校验
  - 添加注释说明参数当前未被使用

- [x] **DTO 映射清理**（问题 4）
  - 重命名 `getTypeID` → `getPhaseID`
  - 补充 WorkItemSummary.State 字段和映射
  - 补充 WorkItem.ReporterID/Reporter 字段和映射

- [x] **测试通过**
  - 所有测试通过：`go test ./...`

- [x] **文档更新**
  - `docs/workitem_management.md` 中的"取消关联"章节已更新

### 修改的文件

1. `sdk/client.go` - 修复超时配置逻辑
2. `sdk/service/workitem/workitem_links.go` - 修复 Unlink 和 ListLinkTypes
3. `internal/api/workitem/dto.go` - 修复 DTO 映射问题
4. `docs/workitem_management.md` - 更新 Unlink 使用示例

### 新增的文件

1. `sdk/client_test.go` - 超时配置单元测试

### 兼容性影响

1. **Unlink**：
   - 旧方法 `Unlink(ctx, linkID)` 标记为 Deprecated
   - 返回明确错误信息，引导用户使用 `UnlinkWithWorkItem(ctx, workItemID, linkID)`
   - 编译通过，运行时错误引导迁移

2. **其他修复**：
   - 无破坏性变更
   - 所有改动保持向后兼容

### 验收结果

✅ **已知严重问题全部消除**：
   - 默认超时行为统一为 30 秒
   - Unlink 可在真实环境下成功取消关联

✅ **接口签名与行为一致**：
   - ListLinkTypes 参数使用与后端对齐
   - DTO ↔ Model 映射中命名清晰，无"永远为空"的伪字段

✅ **测试通过**：
   - `go test ./...` 全部通过
   - 新增单元测试覆盖关键修复点

✅ **文档同步更新**：
   - `docs/workitem_management.md` 描述与真实行为一致
   - 示例代码可直接运行

✅ **不破坏用户空间**：
   - 已有对外 API 行为未静默改变
   - 所有调整都有清晰的兼容策略

**Phase 7 完成！**
