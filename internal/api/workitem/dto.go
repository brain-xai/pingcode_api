package workitem

import (
	"fmt"
	workitemmodel "github.com/brain-xai/pingcode_api/sdk/model/workitem"
)

// ListWorkItemsResponse 工作项列表响应
type ListWorkItemsResponse struct {
	PageSize  int        `json:"page_size"`
	PageIndex int        `json:"page_index"`
	Total     int        `json:"total"`
	Values    []WorkItem `json:"values"`
}

// WorkItem API 工作项 DTO（与 OpenAPI 一致）
type WorkItem struct {
	ID               string                 `json:"id"`
	URL              string                 `json:"url"`
	Identifier       string                 `json:"identifier"`
	ShortID          string                 `json:"short_id"`
	HTMLURL          string                 `json:"html_url"`
	Title            string                 `json:"title"`
	Description      string                 `json:"description"`
	Project          *ProjectSummary        `json:"project"`
	Type             string                 `json:"type"`
	State            *WorkItemStatus        `json:"state"`
	Priority         *WorkItemPriority      `json:"priority"`
	ParentID         string                 `json:"parent_id"`
	Parent           *WorkItemSummary       `json:"parent"`
	AssigneeID       string                 `json:"assignee_id"`
	Assignee         *User                  `json:"assignee"`
	ParticipantIDs   []string               `json:"participant_ids"`
	Participants     []Participant          `json:"participants"`
	StartAt          int64                  `json:"start_at"`
	EndAt            int64                  `json:"end_at"`
	CompletedAt      *int64                 `json:"completed_at"`
	SprintID         string                 `json:"sprint_id"`
	Sprint           *Sprint                `json:"sprint"`
	VersionID        string                 `json:"version_id"`
	Version          *Version               `json:"version"`
	BoardID          string                 `json:"board_id"`
	Board            *Board                 `json:"board"`
	EntryID          string                 `json:"entry_id"`
	Entry            *Entry                 `json:"entry"`
	SwimlaneID       string                 `json:"swimlane_id"`
	Swimlane         *Swimlane              `json:"swimlane"`
	Phase            *Phase                 `json:"phase"`
	StoryPoints      float64                `json:"story_points"`
	EstimatedWorkload float64               `json:"estimated_workload"`
	RemainingWorkload float64               `json:"remaining_workload"`
	Properties       map[string]interface{} `json:"properties"`
	Tags             []WorkItemTag          `json:"tags"`
	PublicImageToken string                 `json:"public_image_token"`
	CreatedAt        int64                  `json:"created_at"`
	UpdatedAt        int64                  `json:"updated_at"`
	CreatedBy        *User                  `json:"created_by"`
	UpdatedBy        *User                  `json:"updated_by"`
	IsArchived       int                    `json:"is_archived"`
	IsDeleted        int                    `json:"is_deleted"`
}

// ProjectSummary 项目简要信息 DTO
type ProjectSummary struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	IsArchived int    `json:"is_archived"`
	IsDeleted  int    `json:"is_deleted"`
}

// WorkItemSummary 工作项简要信息 DTO（用于嵌套引用）
type WorkItemSummary struct {
	ID         string                 `json:"id"`
	Identifier string                 `json:"identifier"`
	Title      string                 `json:"title"`
	Type       string                 `json:"type"`
	StartAt    int64                  `json:"start_at"`
	EndAt      int64                  `json:"end_at"`
	ParentID   string                 `json:"parent_id"`
	Properties map[string]interface{} `json:"properties"`
}

// Sprint Sprint DTO
type Sprint struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Version 版本 DTO
type Version struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	StartAt int64  `json:"start_at"`
	EndAt   int64  `json:"end_at"`
	Stage  *Stage `json:"stage"`
}

// Stage 阶段 DTO
type Stage struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Color string `json:"color"`
}

// Board 看板 DTO
type Board struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	WorkItemTypes []string `json:"work_item_types"`
}

// Entry 看板栏 DTO
type Entry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Swimlane 泳道 DTO
type Swimlane struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Phase 阶段 DTO
type Phase struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
}

// User 用户 DTO
type User struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

// UserGroup 团队 DTO
type UserGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Participant 参与人 DTO
type Participant struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	User      *User      `json:"user"`
	UserGroup *UserGroup `json:"user_group"`
}

// CreateWorkItemRequest 创建工作项请求 DTO
type CreateWorkItemRequest struct {
	ProjectID         string                 `json:"project_id"`
	TypeID            string                 `json:"type_id"`
	Title             string                 `json:"title"`
	Description       string                 `json:"description,omitempty"`
	StartAt           *int64                 `json:"start_at,omitempty"`
	EndAt             *int64                 `json:"end_at,omitempty"`
	PriorityID        string                 `json:"priority_id,omitempty"`
	StateID           string                 `json:"state_id,omitempty"`
	AssigneeID        string                 `json:"assignee_id,omitempty"`
	ParentID          string                 `json:"parent_id,omitempty"`
	SprintID          string                 `json:"sprint_id,omitempty"`
	VersionID         string                 `json:"version_id,omitempty"`
	BoardID           string                 `json:"board_id,omitempty"`
	EntryID           string                 `json:"entry_id,omitempty"`
	SwimlaneID        string                 `json:"swimlane_id,omitempty"`
	StoryPoints       float64                `json:"story_points,omitempty"`
	EstimatedWorkload float64                `json:"estimated_workload,omitempty"`
	RemainingWorkload float64                `json:"remaining_workload,omitempty"`
	Properties        map[string]interface{} `json:"properties,omitempty"`
	ParticipantIDs    []string               `json:"participant_ids,omitempty"`
}

// UpdateWorkItemRequest 更新工作项请求 DTO
type UpdateWorkItemRequest struct {
	Title             *string                `json:"title,omitempty"`
	Description       *string                `json:"description,omitempty"`
	StartAt           *int64                 `json:"start_at,omitempty"`
	EndAt             *int64                 `json:"end_at,omitempty"`
	PriorityID        *string                `json:"priority_id,omitempty"`
	StateID           *string                `json:"state_id,omitempty"`
	AssigneeID        *string                `json:"assignee_id,omitempty"`
	ParentID          *string                `json:"parent_id,omitempty"`
	SprintID          *string                `json:"sprint_id,omitempty"`
	VersionID         *string                `json:"version_id,omitempty"`
	BoardID           *string                `json:"board_id,omitempty"`
	EntryID           *string                `json:"entry_id,omitempty"`
	SwimlaneID        *string                `json:"swimlane_id,omitempty"`
	StoryPoints       *float64               `json:"story_points,omitempty"`
	EstimatedWorkload *float64               `json:"estimated_workload,omitempty"`
	RemainingWorkload *float64               `json:"remaining_workload,omitempty"`
	Properties        map[string]interface{} `json:"properties,omitempty"`
	ParticipantIDs    []string               `json:"participant_ids,omitempty"`
}

// BatchUpdateWorkItemsRequest 批量更新工作项请求 DTO
type BatchUpdateWorkItemsRequest struct {
	WorkItemIDs []string              `json:"work_item_ids"`
	Updates     UpdateWorkItemRequest `json:"updates"`
}

// BatchUpdateWorkItemsResponse 批量更新响应 DTO
type BatchUpdateWorkItemsResponse struct {
	SuccessIDs []string `json:"success_ids"`
	FailedIDs  []string `json:"failed_ids"`
}

// ToModel 将 WorkItem DTO 转换为 Model
func (w *WorkItem) ToModel() *workitemmodel.WorkItem {
	model := &workitemmodel.WorkItem{
		ID:                w.ID,
		URL:               w.URL,
		Identifier:        w.Identifier,
		ShortID:           w.ShortID,
		HTMLURL:           w.HTMLURL,
		Title:             w.Title,
		Description:       w.Description,
		ParentID:          w.ParentID,
		AssigneeID:        w.AssigneeID,
		ParticipantIDs:    w.ParticipantIDs,
		StartAt:           w.StartAt,
		EndAt:             w.EndAt,
		CompletedAt:       w.CompletedAt,
		SprintID:          w.SprintID,
		VersionID:         w.VersionID,
		BoardID:           w.BoardID,
		EntryID:           w.EntryID,
		SwimlaneID:        w.SwimlaneID,
		Phase:             w.getTypeID(w.Phase),
		StoryPoints:       w.StoryPoints,
		EstimatedWorkload: w.EstimatedWorkload,
		RemainingWorkload: w.RemainingWorkload,
		PublicImageToken:  w.PublicImageToken,
		CreatedAt:         w.CreatedAt,
		UpdatedAt:         w.UpdatedAt,
		IsArchived:        w.IsArchived == 1,
		IsDeleted:         w.IsDeleted == 1,
	}

	// 转换 Project
	if w.Project != nil {
		model.Project = w.Project.ToModel()
	}

	// 转换 Type
	model.Type = &workitemmodel.WorkItemType{
		ID:   w.Type,
		Name: w.Type,
	}

	// 转换 State
	if w.State != nil {
		model.State = w.State.ToModel()
	}

	// 转换 Priority
	if w.Priority != nil {
		model.Priority = w.Priority.ToModel()
	}

	// 转换 Parent
	if w.Parent != nil {
		model.Parent = w.Parent.ToModel()
	}

	// 转换 Assignee
	if w.Assignee != nil {
		model.Assignee = w.Assignee.ToModel()
	}

	// 转换 Sprint
	if w.Sprint != nil {
		model.Sprint = w.Sprint.ToModel()
	}

	// 转换 Version
	if w.Version != nil {
		model.Version = w.Version.ToModel()
	}

	// 转换 Board
	if w.Board != nil {
		model.Board = w.Board.ToModel()
	}

	// 转换 Entry
	if w.Entry != nil {
		model.Entry = w.Entry.ToModel()
	}

	// 转换 Swimlane
	if w.Swimlane != nil {
		model.Swimlane = w.Swimlane.ToModel()
	}

	// 转换 Properties
	if w.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range w.Properties {
			model.Properties[k] = toString(v)
		}
	}

	// 转换 Tags
	if len(w.Tags) > 0 {
		model.Tags = make([]workitemmodel.WorkItemTag, len(w.Tags))
		for i, t := range w.Tags {
			model.Tags[i] = *t.ToModel()
		}
	}

	// 转换 Participants
	if len(w.Participants) > 0 {
		model.Participants = make([]workitemmodel.Participant, len(w.Participants))
		for i, p := range w.Participants {
			model.Participants[i] = *p.ToModel()
		}
	}

	// 转换 CreatedBy
	if w.CreatedBy != nil {
		model.CreatedBy = w.CreatedBy.ToModel()
	}

	// 转换 UpdatedBy
	if w.UpdatedBy != nil {
		model.UpdatedBy = w.UpdatedBy.ToModel()
	}

	return model
}

// ToModel 将 ProjectSummary DTO 转换为 Model
func (p *ProjectSummary) ToModel() *workitemmodel.ProjectSummary {
	return &workitemmodel.ProjectSummary{
		ID:         p.ID,
		Identifier: p.Identifier,
		Name:       p.Name,
		Type:       p.Type,
	}
}

// ToModel 将 WorkItemSummary DTO 转换为 Model
func (w *WorkItemSummary) ToModel() *workitemmodel.WorkItemSummary {
	model := &workitemmodel.WorkItemSummary{
		ID:         w.ID,
		Identifier: w.Identifier,
		Title:      w.Title,
	}

	// 转换 Type
	model.Type = &workitemmodel.WorkItemType{
		ID:   w.Type,
		Name: w.Type,
	}

	// 转换 Properties
	if w.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range w.Properties {
			model.Properties[k] = toString(v)
		}
	}

	return model
}

// ToModel 将 Sprint DTO 转换为 Model
func (s *Sprint) ToModel() *workitemmodel.Sprint {
	return &workitemmodel.Sprint{
		ID:   s.ID,
		Name: s.Name,
	}
}

// ToModel 将 Version DTO 转换为 Model
func (v *Version) ToModel() *workitemmodel.Version {
	return &workitemmodel.Version{
		ID:   v.ID,
		Name: v.Name,
	}
}

// ToModel 将 Board DTO 转换为 Model
func (b *Board) ToModel() *workitemmodel.Board {
	return &workitemmodel.Board{
		ID:   b.ID,
		Name: b.Name,
	}
}

// ToModel 将 Entry DTO 转换为 Model
func (e *Entry) ToModel() *workitemmodel.Entry {
	return &workitemmodel.Entry{
		ID:   e.ID,
		Name: e.Name,
	}
}

// ToModel 将 Swimlane DTO 转换为 Model
func (s *Swimlane) ToModel() *workitemmodel.Swimlane {
	return &workitemmodel.Swimlane{
		ID:   s.ID,
		Name: s.Name,
	}
}

// ToModel 将 User DTO 转换为 Model
func (u *User) ToModel() *workitemmodel.User {
	return &workitemmodel.User{
		ID:          u.ID,
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Avatar:      u.Avatar,
	}
}

// ToModel 将 Participant DTO 转换为 Model
func (p *Participant) ToModel() *workitemmodel.Participant {
	model := &workitemmodel.Participant{
		ID:   p.ID,
		Type: p.Type,
	}

	if p.User != nil {
		model.User = p.User.ToModel()
	}

	if p.UserGroup != nil {
		model.UserGroup = &workitemmodel.UserGroup{
			ID:   p.UserGroup.ID,
			Name: p.UserGroup.Name,
		}
	}

	return model
}

// getTypeID 从 Phase 获取 ID
func (w *WorkItem) getTypeID(phase *Phase) string {
	if phase == nil {
		return ""
	}
	return phase.ID
}

// toString 辅助函数：将任意类型转换为字符串
func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%v", val)
	case bool:
		return fmt.Sprintf("%v", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// ToCreateRequestDTO 将 CreateInput 转换为请求 DTO
func ToCreateRequestDTO(input workitemmodel.WorkItemCreateInput) CreateWorkItemRequest {
	return CreateWorkItemRequest{
		ProjectID:         input.ProjectID,
		TypeID:            input.TypeID,
		Title:             input.Title,
		Description:       input.Description,
		StartAt:           input.StartAt,
		EndAt:             input.EndAt,
		PriorityID:        input.PriorityID,
		StateID:           input.StateID,
		AssigneeID:        input.AssigneeID,
		ParentID:          input.ParentID,
		SprintID:          input.SprintID,
		VersionID:         input.VersionID,
		BoardID:           input.BoardID,
		EntryID:           input.EntryID,
		SwimlaneID:        input.SwimlaneID,
		StoryPoints:       input.StoryPoints,
		EstimatedWorkload: input.EstimatedWorkload,
		RemainingWorkload: input.RemainingWorkload,
		Properties:        input.Properties,
		ParticipantIDs:    input.ParticipantIDs,
	}
}

// ToUpdateRequestDTO 将 UpdateInput 转换为请求 DTO
func ToUpdateRequestDTO(input workitemmodel.WorkItemUpdateInput) UpdateWorkItemRequest {
	return UpdateWorkItemRequest{
		Title:             input.Title,
		Description:       input.Description,
		StartAt:           input.StartAt,
		EndAt:             input.EndAt,
		PriorityID:        input.PriorityID,
		StateID:           input.StateID,
		AssigneeID:        input.AssigneeID,
		ParentID:          input.ParentID,
		SprintID:          input.SprintID,
		VersionID:         input.VersionID,
		BoardID:           input.BoardID,
		EntryID:           input.EntryID,
		SwimlaneID:        input.SwimlaneID,
		StoryPoints:       input.StoryPoints,
		EstimatedWorkload: input.EstimatedWorkload,
		RemainingWorkload: input.RemainingWorkload,
		Properties:        input.Properties,
		ParticipantIDs:    input.ParticipantIDs,
	}
}

// ToBatchUpdateRequestDTO 将 BatchUpdateInput 转换为请求 DTO
func ToBatchUpdateRequestDTO(input workitemmodel.WorkItemBatchUpdateInput) BatchUpdateWorkItemsRequest {
	return BatchUpdateWorkItemsRequest{
		WorkItemIDs: input.WorkItemIDs,
		Updates:     ToUpdateRequestDTO(input.Updates),
	}
}
