package ship

import (
	"fmt"
	shipmodel "github.com/brain-xai/pingcode_api/sdk/model/ship"
)

// ListRequirementsResponse 需求列表响应
type ListRequirementsResponse struct {
	PageSize  int           `json:"page_size"`
	PageIndex int           `json:"page_index"`
	Total     int           `json:"total"`
	Values    []Requirement `json:"values"`
}

// Requirement API 需求 DTO（与 OpenAPI 一致）
type Requirement struct {
	ID          string                 `json:"id"`
	URL         string                 `json:"url"`
	Identifier  string                 `json:"identifier"`
	Title       string                 `json:"title"`
	ShortID     string                 `json:"short_id"`
	HTMLURL     string                 `json:"html_url"`
	Product     *ProductSummary        `json:"product"`
	Assignee    *User                  `json:"assignee"`
	State       *RequirementState      `json:"state"`
	Priority    *RequirementPriority   `json:"priority"`
	Plan        *Plan                  `json:"plan"`
	Suite       *Suite                 `json:"suite"`
	PlanAt      *TimeRange             `json:"plan_at"`
	RealAt      *TimeRange             `json:"real_at"`
	Score       float64                `json:"score"`
	Progress    float64                `json:"progress"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
	Participants []Participant         `json:"participants"`
	CompletedAt *int64                 `json:"completed_at"`
	CompletedBy *User                  `json:"completed_by"`
	CreatedAt   int64                  `json:"created_at"`
	UpdatedAt   int64                  `json:"updated_at"`
	CreatedBy   *User                  `json:"created_by"`
	UpdatedBy   *User                  `json:"updated_by"`
	IsArchived  int                    `json:"is_archived"`
	IsDeleted   int                    `json:"is_deleted"`
}

// ProductSummary 产品简要信息 DTO
type ProductSummary struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	IsArchived int    `json:"is_archived"`
	IsDeleted  int    `json:"is_deleted"`
}

// RequirementState 需求状态 DTO
type RequirementState struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// RequirementPriority 需求优先级 DTO
type RequirementPriority struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Plan 产品排期 DTO
type Plan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Suite 产品模块 DTO
type Suite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

// TimeRange 时间范围 DTO
type TimeRange struct {
	From        int64  `json:"from"`
	To          int64  `json:"to"`
	Granularity string `json:"granularity"`
}

// Participant 参与人 DTO
type Participant struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	User      *User      `json:"user"`
	UserGroup *UserGroup `json:"user_group"`
}

// CreateRequirementRequest 创建需求请求 DTO
type CreateRequirementRequest struct {
	ProductID   string                 `json:"product_id"`
	Title       string                 `json:"title"`
	AssigneeID  string                 `json:"assignee_id,omitempty"`
	Description string                 `json:"description,omitempty"`
	SuiteID     string                 `json:"suite_id,omitempty"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
}

// UpdateRequirementRequest 更新需求请求 DTO
type UpdateRequirementRequest struct {
	Title       *string                `json:"title,omitempty"`
	Description *string                `json:"description,omitempty"`
	StateID     *string                `json:"state_id,omitempty"`
	PriorityID  *string                `json:"priority_id,omitempty"`
	AssigneeID  *string                `json:"assignee_id,omitempty"`
	Progress    *float64               `json:"progress,omitempty"`
	PlanAt      *TimeRange             `json:"plan_at,omitempty"`
	RealAt      *TimeRange             `json:"real_at,omitempty"`
	PlanID      *string                `json:"plan_id,omitempty"`
	SuiteID     *string                `json:"suite_id,omitempty"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
}

// ToModel 将 Requirement DTO 转换为 Model
func (r *Requirement) ToModel() *shipmodel.Requirement {
	model := &shipmodel.Requirement{
		ID:          r.ID,
		URL:         r.URL,
		Identifier:  r.Identifier,
		Title:       r.Title,
		ShortID:     r.ShortID,
		HTMLURL:     r.HTMLURL,
		Score:       r.Score,
		Progress:    r.Progress,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		IsArchived:  r.IsArchived == 1,
		IsDeleted:   r.IsDeleted == 1,
	}

	// 转换 Product
	if r.Product != nil {
		model.Product = r.Product.ToModel()
	}

	// 转换 Assignee
	if r.Assignee != nil {
		model.Assignee = r.Assignee.ToModel()
	}

	// 转换 State
	if r.State != nil {
		model.State = r.State.ToModel()
	}

	// 转换 Priority
	if r.Priority != nil {
		model.Priority = r.Priority.ToModel()
	}

	// 转换 Plan
	if r.Plan != nil {
		model.Plan = r.Plan.ToModel()
	}

	// 转换 Suite
	if r.Suite != nil {
		model.Suite = r.Suite.ToModel()
	}

	// 转换 PlanAt
	if r.PlanAt != nil {
		model.PlanAt = r.PlanAt.ToModel()
	}

	// 转换 RealAt
	if r.RealAt != nil {
		model.RealAt = r.RealAt.ToModel()
	}

	// 转换 Properties
	if r.Properties != nil {
		model.Properties = make(map[string]string)
		for k, v := range r.Properties {
			model.Properties[k] = toString(v)
		}
	}

	// 转换 Participants
	if len(r.Participants) > 0 {
		model.Participants = make([]shipmodel.Participant, len(r.Participants))
		for i, p := range r.Participants {
			model.Participants[i] = *p.ToModel()
		}
	}

	// 转换 CompletedAt
	if r.CompletedAt != nil {
		model.CompletedAt = r.CompletedAt
	}

	// 转换 CompletedBy
	if r.CompletedBy != nil {
		model.CompletedBy = r.CompletedBy.ToModel()
	}

	// 转换 CreatedBy
	if r.CreatedBy != nil {
		model.CreatedBy = r.CreatedBy.ToModel()
	}

	// 转换 UpdatedBy
	if r.UpdatedBy != nil {
		model.UpdatedBy = r.UpdatedBy.ToModel()
	}

	return model
}

// ToModel 将 ProductSummary DTO 转换为 Model
func (p *ProductSummary) ToModel() *shipmodel.ProductSummary {
	return &shipmodel.ProductSummary{
		ID:         p.ID,
		Identifier: p.Identifier,
		Name:       p.Name,
		IsArchived: p.IsArchived == 1,
		IsDeleted:  p.IsDeleted == 1,
	}
}

// ToModel 将 RequirementState DTO 转换为 Model
func (s *RequirementState) ToModel() *shipmodel.RequirementState {
	return &shipmodel.RequirementState{
		ID:   s.ID,
		Name: s.Name,
		Type: s.Type,
	}
}

// ToModel 将 RequirementPriority DTO 转换为 Model
func (p *RequirementPriority) ToModel() *shipmodel.RequirementPriority {
	return &shipmodel.RequirementPriority{
		ID:   p.ID,
		Name: p.Name,
	}
}

// ToModel 将 Plan DTO 转换为 Model
func (p *Plan) ToModel() *shipmodel.Plan {
	return &shipmodel.Plan{
		ID:   p.ID,
		Name: p.Name,
		URL:  p.URL,
	}
}

// ToModel 将 Suite DTO 转换为 Model
func (s *Suite) ToModel() *shipmodel.Suite {
	return &shipmodel.Suite{
		ID:   s.ID,
		Name: s.Name,
		Type: s.Type,
		URL:  s.URL,
	}
}

// ToModel 将 TimeRange DTO 转换为 Model
func (t *TimeRange) ToModel() *shipmodel.TimeRange {
	return &shipmodel.TimeRange{
		From:        t.From,
		To:          t.To,
		Granularity: t.Granularity,
	}
}

// ToModel 将 Participant DTO 转换为 Model
func (p *Participant) ToModel() *shipmodel.Participant {
	model := &shipmodel.Participant{
		ID:   p.ID,
		Type: p.Type,
	}

	if p.User != nil {
		model.User = p.User.ToModel()
	}

	if p.UserGroup != nil {
		model.UserGroup = &shipmodel.UserGroup{
			ID:   p.UserGroup.ID,
			Name: p.UserGroup.Name,
		}
	}

	return model
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
