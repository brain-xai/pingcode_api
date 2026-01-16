# 工作项管理

本文档介绍如何使用 PingCode SDK 进行工作项的完整管理，包括创建、更新、删除、查询，以及属性、标签、关联、流转记录和交付目标的管理。

## 概述

工作项（Work Item）是 PingCode 中用于跟踪和管理工作的核心实体，包括史诗、特性、用户故事、任务、缺陷等多种类型。

## 初始化客户端

```go
import "github.com/brain-xai/pingcode_api/sdk"

// 从环境变量加载配置
config, err := sdk.LoadConfigFromEnv()
if err != nil {
    log.Fatal(err)
}

// 创建 SDK 客户端
client, err := sdk.NewClient(config)
if err != nil {
    log.Fatal(err)
}
```

## 基础操作

### 创建工作项

```go
ctx := context.Background()

workItem, err := client.WorkItem().Create(ctx, workitemmodel.WorkItemCreateInput{
    ProjectID:  "project_id",     // 必填：项目 ID
    TypeID:     "story",          // 必填：工作项类型 ID
    Title:      "实现用户登录功能", // 必填：标题
    Description: "实现用户登录功能，包括用户名密码登录和第三方登录", // 可选
    PriorityID: "priority_id",    // 可选：优先级 ID
    StateID:    "state_id",       // 可选：状态 ID
    AssigneeID: "user_id",        // 可选：负责人 ID
    SprintID:   "sprint_id",      // 可选：Sprint ID（Scrum 项目）
    BoardID:    "board_id",       // 可选：看板 ID（Kanban 项目）
    StoryPoints: 5.0,              // 可选：故事点
})
if err != nil {
    log.Fatal(err)
}
```

### 部分更新工作项

```go
updatedTitle := "更新后的标题"
updatedDesc := "更新后的描述"

workItem, err := client.WorkItem().Update(ctx, workItemID, workitemmodel.WorkItemUpdateInput{
    Title:       &updatedTitle,  // 使用指针表示可选字段
    Description: &updatedDesc,
    StoryPoints: pointerTo(8.0), // 更新故事点
})
if err != nil {
    log.Fatal(err)
}
```

### 批量更新工作项

```go
result, err := client.WorkItem().BatchUpdate(ctx, workitemmodel.WorkItemBatchUpdateInput{
    WorkItemIDs: []string{"id1", "id2", "id3"}, // 要更新的工作项 ID 列表
    Updates: workitemmodel.WorkItemUpdateInput{
        PriorityID: pointerTo("new_priority_id"),
    },
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("成功: %v, 失败: %v\n", result.SuccessIDs, result.FailedIDs)
```

### 获取工作项列表

```go
workItems, pagination, err := client.WorkItem().List(ctx, workitemmodel.WorkItemFilter{
    ProjectID: "project_id", // 必填：项目 ID
    Query:     "登录",        // 可选：搜索关键字
    TypeID:    "story",       // 可选：按类型过滤
    StateID:   "in_progress", // 可选：按状态过滤
    AssigneeID: "user_id",    // 可选：按负责人过滤
    TagIDs:    []string{"tag1", "tag2"}, // 可选：按标签过滤
    PageSize:  30,           // 可选：每页大小（默认 30）
    PageIndex: 0,            // 可选：页码（从 0 开始）
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("共 %d 个工作项，当前页 %d 个\n", pagination.Total, len(workItems))
for _, wi := range workItems {
    fmt.Printf("- [%s] %s\n", wi.Identifier, wi.Title)
}
```

### 获取单个工作项

```go
workItem, err := client.WorkItem().Get(ctx, workItemID)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("工作项: %s\n", workItem.Title)
fmt.Printf("状态: %s\n", workItem.State.Name)
fmt.Printf("负责人: %s\n", workItem.Assignee.DisplayName)
```

### 删除工作项

```go
err := client.WorkItem().Delete(ctx, workItemID)
if err != nil {
    log.Fatal(err)
}
```

## 属性与分类管理

### 获取工作项类型列表

```go
types, err := client.WorkItem().ListTypes(ctx, workitemmodel.WorkItemTypeScope{
    ProjectID: "project_id",
})
if err != nil {
    log.Fatal(err)
}

for _, t := range types {
    fmt.Printf("- %s (%s)\n", t.Name, t.ID)
}
```

### 获取工作项状态列表

```go
statusList, err := client.WorkItem().ListStatuses(ctx, workitemmodel.WorkItemStatusFilter{
    ProjectID: "project_id", // 必填
    TypeID:    "story",      // 可选：按类型过滤
    PageSize:  100,
})
if err != nil {
    log.Fatal(err)
}

for _, s := range statusList.Values {
    fmt.Printf("- %s (%s)\n", s.Name, s.Type)
}
```

### 获取工作项属性列表

```go
fields, err := client.WorkItem().ListFields(ctx, workitemmodel.WorkItemFieldFilter{
    ProjectID: "project_id", // 必填
    TypeID:    "story",      // 可选：按类型过滤
})
if err != nil {
    log.Fatal(err)
}

for _, f := range fields.Values {
    fmt.Printf("- %s (%s)\n", f.Name, f.Type)
}
```

### 获取工作项优先级列表

```go
priorities, err := client.WorkItem().ListPriorities(ctx, "project_id")
if err != nil {
    log.Fatal(err)
}

for _, p := range priorities.Values {
    fmt.Printf("- %s\n", p.Name)
}
```

### 获取工作项标签列表

```go
tags, err := client.WorkItem().ListTags(ctx, workitemmodel.WorkItemTagFilter{
    ProjectID: "project_id", // 必填
    TypeID:    "story",      // 可选：按类型过滤
})
if err != nil {
    log.Fatal(err)
}

for _, t := range tags.Values {
    fmt.Printf("- %s (%s)\n", t.Name, t.Color)
}
```

## 标签管理

### 向工作项添加标签

```go
err := client.WorkItem().AddTag(ctx, workItemID, tagID)
if err != nil {
    log.Fatal(err)
}
```

### 从工作项移除标签

```go
err := client.WorkItem().RemoveTag(ctx, workItemID, tagID)
if err != nil {
    log.Fatal(err)
}
```

## 关联管理

### 获取关联类型列表

```go
linkTypes, err := client.WorkItem().ListLinkTypes(ctx, "project_id")
if err != nil {
    log.Fatal(err)
}

for _, lt := range linkTypes {
    fmt.Printf("- %s (%s)\n", lt.Name, lt.Category)
}
```

### 关联工作项

```go
link, err := client.WorkItem().Link(ctx, fromWorkItemID, toWorkItemID, "relate")
if err != nil {
    log.Fatal(err)
}
```

### 获取关联列表

```go
links, _, err := client.WorkItem().ListLinks(ctx, workItemID, workitemmodel.WorkItemLinkFilter{
    TypeID: "relate", // 可选：按关联类型过滤
})
if err != nil {
    log.Fatal(err)
}

for _, l := range links {
    fmt.Printf("- %s -> %s (%s)\n", l.Source.Title, l.Target.Title, l.RelationType)
}
```

### 取消关联

> 注意：当前版本需要提供工作项 ID 和关联 ID

```go
err := client.WorkItem().UnlinkWithWorkItem(ctx, workItemID, linkID)
if err != nil {
    log.Fatal(err)
}
```

旧方法 `Unlink(ctx, linkID)` 已废弃，请使用新方法 `UnlinkWithWorkItem`。


## 流转记录

### 获取工作项流转记录

```go
transitions, err := client.WorkItem().ListTransitions(ctx, workItemID)
if err != nil {
    log.Fatal(err)
}

for _, t := range transitions.Values {
    fromState := ""
    toState := ""
    if t.FromState != nil {
        fromState = t.FromState.Name
    }
    if t.ToState != nil {
        toState = t.ToState.Name
    }
    fmt.Printf("- %s -> %s (操作人: %s)\n", fromState, toState, t.CreatedBy.DisplayName)
}
```

## 交付目标

### 创建交付目标

```go
deliveryTarget, err := client.WorkItem().CreateDeliveryTarget(ctx, workitemmodel.WorkItemDeliveryTargetCreateInput{
    WorkItemID:  "work_item_id", // 必填：工作项 ID
    Name:        "API 文档",      // 必填：交付目标名称
    Description: "完成 API 接口文档编写",
    ContentType: "link",
    Content: map[string]interface{}{
        "name": "API 文档链接",
        "href": "https://docs.example.com/api",
    },
})
if err != nil {
    log.Fatal(err)
}
```

### 更新交付目标

```go
updatedName := "更新后的交付目标名称"
deliveryTarget, err := client.WorkItem().UpdateDeliveryTarget(ctx, targetID, workitemmodel.WorkItemDeliveryTargetUpdateInput{
    Name: &updatedName,
})
if err != nil {
    log.Fatal(err)
}
```

### 获取交付目标列表

```go
deliveryTargets, _, err := client.WorkItem().ListDeliveryTargets(ctx, workitemmodel.WorkItemDeliveryTargetFilter{
    WorkItemID: "work_item_id", // 必填：工作项 ID
    Status:     "pending",       // 可选：按状态过滤
})
if err != nil {
    log.Fatal(err)
}

for _, dt := range deliveryTargets {
    fmt.Printf("- %s (%s)\n", dt.Name, dt.ContentType)
}
```

### 删除交付目标

```go
err := client.WorkItem().DeleteDeliveryTarget(ctx, targetID)
if err != nil {
    log.Fatal(err)
}
```

## 辅助函数

```go
// 创建指针的辅助函数
func pointerTo[T any](v T) *T {
    return &v
}
```

## 完整示例

请参考 `examples/workitems/main.go` 查看完整的使用示例。

运行示例：

```bash
cd examples/workitems
export PINGCODE_BASE_URL=https://open.pingcode.com
export PINGCODE_ACCESS_TOKEN=your_token_here
go run main.go
```

## 错误处理

所有方法在失败时都会返回错误。建议根据错误类型进行适当的处理：

```go
import "errors"

workItem, err := client.WorkItem().Create(ctx, input)
if err != nil {
    if errors.Is(err, sdk.ErrUnauthorized) {
        // 处理认证错误
        log.Fatal("认证失败，请检查 Token")
    } else if apiErr, ok := err.(*sdk.APIError); ok {
        // 处理 API 错误
        log.Printf("API 错误: %s (代码: %s)\n", apiErr.Message, apiErr.Code)
    } else {
        // 处理其他错误
        log.Fatal(err)
    }
}
```

## 注意事项

1. **项目类型兼容性**：某些字段只对特定项目类型有效：
   - `SprintID` 只对 Scrum 和 Hybrid 项目有效
   - `BoardID`、`EntryID`、`SwimlaneID` 只对 Kanban 和 Hybrid 项目有效

2. **状态流转**：设置 `StateID` 时必须确保该状态符合工作项类型的状态流转规则

3. **父子关系**：设置 `ParentID` 时需要确保父工作项类型与当前工作项类型的父子关系配置兼容

4. **批量操作**：批量更新操作需要确保所有目标工作项 ID 属于同一个项目
