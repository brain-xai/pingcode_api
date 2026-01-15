package api

import "github.com/brain-xai/pingcode_api/sdk/model/global"

// User API DTO
type User struct {
	ID             string                 `json:"id"`
	URL            string                 `json:"url"`
	Name           string                 `json:"name"`
	DisplayName    string                 `json:"display_name"`
	Avatar         string                 `json:"avatar"`
	Email          string                 `json:"email"`
	Mobile         string                 `json:"mobile"`
	Status         string                 `json:"status"`
	EmployeeNumber string                 `json:"employee_number"`
	Department     *Department            `json:"department,omitempty"`
	Job            *Job                   `json:"job,omitempty"`
	Preferences    map[string]string      `json:"preferences,omitempty"`
	Permissions    map[string]interface{} `json:"permissions,omitempty"`
}

// Department API DTO
type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Job API DTO
type Job struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListUsersResponse 用户列表响应
type ListUsersResponse struct {
	PageSize  int    `json:"page_size"`
	PageIndex int    `json:"page_index"`
	Total     int    `json:"total"`
	Values    []User `json:"values"`
}

// ToModel 将 API DTO 转换为对外模型
func (u *User) ToModel() *global.User {
	user := &global.User{
		ID:             u.ID,
		URL:            u.URL,
		Name:           u.Name,
		DisplayName:    u.DisplayName,
		Avatar:         u.Avatar,
		Email:          u.Email,
		Mobile:         u.Mobile,
		Status:         u.Status,
		EmployeeNumber: u.EmployeeNumber,
	}

	if u.Department != nil {
		user.Department = &global.Department{
			ID:   u.Department.ID,
			Name: u.Department.Name,
		}
	}

	if u.Job != nil {
		user.Job = &global.Job{
			ID:   u.Job.ID,
			Name: u.Job.Name,
		}
	}

	if u.Preferences != nil {
		user.Preferences = make(map[string]string)
		for k, v := range u.Preferences {
			user.Preferences[k] = v
		}
	}

	if u.Permissions != nil {
		user.Permissions = make(map[string]bool)
		for k, v := range u.Permissions {
			user.Permissions[k] = toBool(v)
		}
	}

	return user
}

// toBool 辅助函数：将任意类型转换为 bool
func toBool(v interface{}) bool {
	switch val := v.(type) {
	case bool:
		return val
	case float64:
		return val != 0
	case string:
		return val == "true" || val == "1"
	default:
		return false
	}
}
