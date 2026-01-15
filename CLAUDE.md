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