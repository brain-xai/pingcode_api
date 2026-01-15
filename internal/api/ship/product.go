package ship

import (
	shipmodel "github.com/brain-xai/pingcode_api/sdk/model/ship"
)

// ListProductsResponse 产品列表响应
type ListProductsResponse struct {
	PageSize  int       `json:"page_size"`
	PageIndex int       `json:"page_index"`
	Total     int       `json:"total"`
	Values    []Product `json:"values"`
}

// Product API 产品 DTO（与 OpenAPI 一致）
type Product struct {
	ID          string         `json:"id"`
	URL         string         `json:"url"`
	Identifier  string         `json:"identifier"`
	Name        string         `json:"name"`
	ScopeType   string         `json:"scope_type"`
	ScopeID     string         `json:"scope_id"`
	Visibility  string         `json:"visibility"`
	Color       string         `json:"color"`
	Description string         `json:"description"`
	Members     []ProductMember `json:"members"`
	CreatedAt   int64          `json:"created_at"`
	UpdatedAt   int64          `json:"updated_at"`
	CreatedBy   *User          `json:"created_by"`
	UpdatedBy   *User          `json:"updated_by"`
	IsArchived  int            `json:"is_archived"` // API 返回 0/1
	IsDeleted   int            `json:"is_deleted"`  // API 返回 0/1
}

// ProductMember 产品成员 DTO
type ProductMember struct {
	ID        string     `json:"id"`
	URL       string     `json:"url"`
	Type      string     `json:"type"`
	User      *User      `json:"user"`
	UserGroup *UserGroup `json:"user_group"`
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

// ToModel 将 Product DTO 转换为 Model
func (p *Product) ToModel() *shipmodel.Product {
	model := &shipmodel.Product{
		ID:          p.ID,
		URL:         p.URL,
		Identifier:  p.Identifier,
		Name:        p.Name,
		ScopeType:   p.ScopeType,
		ScopeID:     p.ScopeID,
		Visibility:  p.Visibility,
		Color:       p.Color,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		IsArchived:  p.IsArchived == 1,
		IsDeleted:   p.IsDeleted == 1,
	}

	// 转换 Members
	if len(p.Members) > 0 {
		model.Members = make([]shipmodel.ProductMember, len(p.Members))
		for i, m := range p.Members {
			model.Members[i] = *m.ToModel()
		}
	}

	// 转换 CreatedBy
	if p.CreatedBy != nil {
		model.CreatedBy = p.CreatedBy.ToModel()
	}

	// 转换 UpdatedBy
	if p.UpdatedBy != nil {
		model.UpdatedBy = p.UpdatedBy.ToModel()
	}

	return model
}

// ToModel 将 ProductMember DTO 转换为 Model
func (m *ProductMember) ToModel() *shipmodel.ProductMember {
	model := &shipmodel.ProductMember{
		ID:   m.ID,
		URL:  m.URL,
		Type: m.Type,
	}

	if m.User != nil {
		model.User = m.User.ToModel()
	}

	if m.UserGroup != nil {
		model.UserGroup = &shipmodel.UserGroup{
			ID:   m.UserGroup.ID,
			Name: m.UserGroup.Name,
		}
	}

	return model
}

// ToModel 将 User DTO 转换为 Model
func (u *User) ToModel() *shipmodel.User {
	return &shipmodel.User{
		ID:          u.ID,
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Avatar:      u.Avatar,
	}
}
