# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 用户原则

1. **改动后必须保证测试通过**
2. **改动后必须更新相关文档**：更新API文档、数据模型文档、测试指南和部署指南
3. **根据phase*.md文件的要求迭代后，需要将你的改动提交到对应的phase*.md文件中**
4. **改动后必须保证编译通过**
5. **改动后必须更新相关的依赖**：如果有新增的依赖项，必须在 `go.mod` 中添加并运行 `go mod tidy`。
6. **使用中文回答**

## 产品文档
参考文档: [prd](docs/prd.md)

## 技术规范
参考文档: [deployment](docs/deployment.md)

## pingcode API文档
参考文档: [api](docs/reference/openpingcode/Readme.md)

## 迭代原则
敏捷迭代的核心原则是持续交付有价值的、可工作的软件，通过短周期迭代（如 Sprint 冲刺）来快速响应变化，强调客户满意度、跨职能团队协作（业务与开发每日沟通）、频繁交付、持续改进，以及重视个人互动和可工作软件胜过流程和文档。 