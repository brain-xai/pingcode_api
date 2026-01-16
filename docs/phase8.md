# Phase 8: 对外发布前的兼容性与发布体系完善

> 目标：在不改变现有对外 API 行为的前提下，补齐「License、Go 版本声明、CI 和兼容性对外承诺」，让 `github.com/brain-xai/pingcode_api` 作为一个公开 SDK 可以被互联网上的开发者放心集成和长期依赖。

---

## 1. License：采用 Apache 2.0

### 1.1 新增 LICENSE 文件

- 在仓库根目录新增 `LICENSE` 文件，内容为标准 **Apache License 2.0** 正文。
- 要求：
  - 不做任何私货修改，直接使用官方文本。
  - 年份与版权所有者信息按实际情况填写（例如：`Copyright [2025] [PingCode]`）。

### 1.2 更新 README 中的 License 小节

- 将当前 README 中的「License: 待定」改为明确说明：
  - 本项目使用 **Apache License 2.0**。
  - 简单一句话描述许可条件（允许商用、修改、分发，需保留版权和许可证说明等）。
- 确保 README 末尾的 License 小节与根目录 `LICENSE` 文件严格一致，不出现冲突说法。

---

## 2. Go 版本与构建要求回到现实世界

### 2.1 调整 go.mod 中的 Go 版本

- 当前：
  - `go.mod`：`go 1.24.2`
- 调整方案：
  - 将 `go 1.24.2` 改为实际 **已经 GA 且被支持的最低 Go 版本**，例如：
    - `go 1.21`（或 1.20，根据实际支持策略决定），确保：
      - 当前代码在该版本上可通过 `go build ./...` 与 `go test ./...`。
- 原则：
  - 使用「真实存在且稳定」的版本号，避免给外部开发者制造不必要的使用门槛。
  - 如将来确实需要更高版本（例如使用新语法），再通过 MINOR 版本提升进行声明。

### 2.2 同步更新 deployment 文档中的语言版本

- 当前：  
  - [docs/deployment.md](docs/deployment.md) 中写有：
    - `语言版本：go 1.24.x`
- 调整方案：
  - 将文档中的语言版本改为与 `go.mod` 一致的版本号（例如：`go 1.21.x`）。
- 要求：
  - 文档必须和 go.mod 保持同步，否则视为文档 bug。

---

## 3. CI：最小但必要的自动化守门人

### 3.1 引入基础 CI 流程

- 目标：保证所有对外发布前，至少满足：
  - `go test ./...` 通过。
  - `go build ./...` 通过。
- 建议方案（以 GitHub Actions 为例）：
  - 新增 `.github/workflows/ci.yml`（文件名可根据团队规范调整）。
  - 流程内容：
    - 触发条件：
      - push 到主分支。
      - 对主分支的 pull request。
    - 运行环境：
      - 至少一个 Go 版本（与 `go.mod` 中声明的版本一致）。
    - 关键步骤：
      - `go mod tidy`（可选，根据你的策略）。
      - `go build ./...`
      - `go test ./...`
- 原则：
  - Phase 8 不追求复杂矩阵、覆盖率门槛等高级规则，只确保**任何改动不会破坏基本可用性**。
  - CI 报红即视为不允许合入主干。

### 3.2 CI 与文档的对应关系

- 在 `docs/deployment.md` 中补充一小节：
  - 说明标准测试命令：`go test ./...`。
  - 说明 CI 的最低要求与入口（例如 GitHub Actions 配置位置）。
- 保证「文档写的构建/测试命令」与「CI 实际执行的命令」完全一致。

---

## 4. 对外兼容性与版本策略的显式说明

### 4.1 README 中增加“兼容性与版本策略”小节

- 在 README 中新增小节，内容包括但不限于：
  - 使用 **语义化版本（SemVer）**：
    - 当前 0.x 阶段，以快速迭代为主，但依然遵守基本的兼容性约束。
  - 对 `sdk/` 下公开 API 的兼容承诺：
    - 一旦某个导出类型、方法进入已发布版本，就视为对外契约。
    - 不会在 MINOR/PATCH 版本中无声改变已发布 API 的语义或行为。
  - 对 Breaking Change 的处理：
    - 仅在 MAJOR 版本中移除已长期标记为 Deprecated 的 API。
    - 所有 Breaking Change 必须在 `CHANGELOG.md` 中明确记录。
  - 指向详细契约文档：
    - 链接到 `docs/api_contract.md`，说明详细规则。

### 4.2 确认 CHANGELOG 与实际 tags 对齐

- 检查 `CHANGELOG.md` 中已有的版本记录（`v0.1.0` ~ `v0.7.0`）是否都有对应的 Git tag。
- Phase 8 不必大改 CHANGELOG，只需要在下一版本发布时：
  - 新增 `[0.8.0] - Phase 8: 发布前准备` 章节。
  - 记录本 phase 实施的变更：
    - 引入 Apache 2.0 License。
    - 调整 Go 版本。
    - 引入基础 CI。
    - 补充 README 的兼容性说明。

---

## 5. 文档与示例：轻量整理

### 5.1 README 中的示例索引整理

- 在 README 的“示例代码”部分，增加一个简单的示例索引表，将已有示例映射清晰：
  - `examples/basic_usage` - 获取项目列表（最小可运行示例）
  - `examples/project_overview` - 项目管理整体验证
  - `examples/ship_products` - 产品管理
  - `examples/ship_requirements` - 需求管理
  - `examples/workitems` - 工作项管理常用操作
- 目的：让新用户在 10 秒内找到自己关心的场景。

### 5.2 pkg.go.dev 文档可读性提升（可选）

- 在 `sdk` 顶层包（如 `sdk/client.go` 或单独文件）增加 package 级注释（英文/双语皆可）：
  - 简要说明：
    - 这是 PingCode OpenAPI 的 Go SDK。
    - 如何初始化客户端（LoadConfigFromEnv + NewClient）。
    - 主入口服务：Project、Ship、WorkItem、Global。
- 这样在 `pkg.go.dev` 上的入口页面不会空白，有利于新用户快速理解。

---

## 6. 验收标准

Phase 8 完成的判定标准：

1. 仓库根目录存在 `LICENSE` 文件，内容为 Apache License 2.0，README 中 License 小节与之保持一致。
2. `go.mod` 中的 `go` 版本为真实存在且支持的版本，`docs/deployment.md` 中的语言版本与之相同。
3. 主分支上存在可运行的 CI 配置，能在至少一个 Go 版本上执行 `go build ./...` 和 `go test ./...` 并通过。
4. README 中新增了“兼容性与版本策略”小节，明确说明 SemVer 和对 `sdk/` API 的兼容承诺，并链接到 `docs/api_contract.md`。
5. 示例索引清晰，外部开发者可以通过 README 快速定位 `examples/*` 中的示例。
6. 为本阶段实际发布的版本（如 `v0.8.0`）在 `CHANGELOG.md` 中新增条目，记录上述变更。

---