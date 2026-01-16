# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.8.0] - 2025-01-16 (Phase 8: 对外发布前的兼容性与发布体系完善)

### Added
- 引入 Apache License 2.0 开源许可
- 新增 GitHub Actions CI 配置，自动化构建和测试
- README 新增"兼容性与版本策略"章节，明确 SemVer 和 API 兼容承诺
- README 新增示例索引表，提升文档可读性

### Changed
- 完善 deployment.md 文档，补充 CI 使用说明
- 优化 README 的 License 章节，提供详细的许可说明
- 更新示例代码章节，添加所有 6 个示例的索引表

### Documentation
- 明确对外 API 兼容性承诺
- 补充 CI/CD 标准流程说明

## [0.7.0] - 2025-01-16 (Phase 7: 质量问题收敛)

### Added
- 客户端超时配置单元测试
- UnlinkWithWorkItem方法，支持正确的关联取消

### Fixed
- 修复默认超时配置错误，统一为30秒
- 修复Unlink实现不可用问题（使用dummy work_item_id）
- 修复ListLinkTypes无用的projectID校验
- 修复DTO映射中的命名混乱（getTypeID → getPhaseID）
- 补充WorkItemSummary.State字段和映射
- 补充WorkItem.ReporterID/Reporter字段和映射

### Deprecated
- Unlink方法标记为废弃，引导使用UnlinkWithWorkItem

## [0.6.0] - 2025-01-15 (Phase 6: 工作项领域接口)

### Added
- 实现工作项管理完整功能（21个接口）
  - 基础操作：Create、Update、BatchUpdate、List、Get、Delete
  - 属性与分类：ListTypes、ListStatuses、ListFields、ListPriorities、ListTags
  - 标签管理：AddTag、RemoveTag
  - 关联管理：ListLinkTypes、Link、ListLinks、Unlink
  - 流程与交付目标：ListTransitions、CreateDeliveryTarget、UpdateDeliveryTarget、ListDeliveryTargets、DeleteDeliveryTarget
- 新增sdk/model/workitem完整数据模型
- 新增sdk/service/workitem服务层实现
- 新增internal/api/workitem DTO映射层
- 新增workitem_management.md使用指南

## [0.5.0] - 2025-01-15 (Phase 5: 需求辅助接口)

### Added
- 新增需求辅助接口（5个）用于UI场景
  - GetRequirementStates - 获取需求状态列表
  - GetRequirementPriorities - 获取优先级列表
  - GetRequirementProperties - 获取自定义属性列表
  - GetRequirementSuites - 获取模块列表
  - GetRequirementPlans - 获取排期列表
- 完善需求状态、优先级、属性等数据模型

## [0.4.0] - 2025-01-15 (Phase 4: Project领域接口)

### Added
- 新增Project领域CRUD完整能力
  - CreateProject - 创建项目
  - UpdateProject - 更新项目
  - DeleteProject - 删除项目
- 新增Global服务用户信息查询
  - GetCurrentUser - 获取当前用户
  - GetUser - 按ID获取用户
  - ListUsers - 获取企业成员列表
- 新增项目成员、进度等查询能力
- 新增examples/project_overview完整示例

## [0.3.0] - 2025-01-15 (Phase 3: Ship领域接口)

### Added
- 新增Ship产品管理接口
  - ListProducts - 获取产品列表
  - GetProduct - 获取产品详情
- 新增Ship需求管理核心接口
  - ListRequirements - 获取需求列表
  - GetRequirement - 获取需求详情
  - CreateRequirement - 创建需求
  - UpdateRequirement - 更新需求
- 新增HTTP客户端PATCH方法支持
- 新增examples/ship_products和examples/ship_requirements示例

## [0.2.0] - 2025-01-15 (Phase 2: 读取类接口完善)

### Added
- 完善Project领域读取接口
  - GetProject - 按ID获取项目详情
  - ListStates - 获取项目状态列表
  - ListMembers - 获取项目成员列表
  - GetProgress - 获取项目进度
- 新增Global领域服务
  - GetCurrentUser - 获取当前用户信息
  - GetUser - 按ID获取用户详情
  - ListUsers - 获取企业成员列表
- 新增httpclient路径参数支持（GetWithPathParams）
- 完善单元测试覆盖率（Project 80%+, Global 80%+）

## [0.1.0] - 2025-01-15 (Phase 1: 基础框架搭建)

### Added
- 初始化SDK基础框架
  - sdk/client - 客户端核心
  - sdk/config - 配置管理
  - sdk/errors - 统一错误模型
  - sdk/model/core/pagination - 分页模型
- 实现认证模块
  - 支持Client Credentials授权方式
  - Token自动过期检测（提前5分钟判定为过期）
- 实现Project基础接口
  - ListProjects - 获取项目列表
- 新增examples/basic_usage基础示例
- 建立deployment.md技术规范体系
- 建立api_contract.md契约规范体系

[Unreleased]: https://github.com/brain-xai/pingcode_api/compare/v0.7.0...HEAD
[0.8.0]: https://github.com/brain-xai/pingcode_api/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/brain-xai/pingcode_api/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/brain-xai/pingcode_api/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/brain-xai/pingcode_api/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/brain-xai/pingcode_api/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/brain-xai/pingcode_api/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/brain-xai/pingcode_api/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/brain-xai/pingcode_api/releases/tag/v0.1.0
