package global

// User 用户模型（复用于 myself 和 directory/users）
type User struct {
	ID             string            `json:"id"`
	URL            string            `json:"url"`
	Name           string            `json:"name"`
	DisplayName    string            `json:"display_name"`
	Avatar         string            `json:"avatar"`
	Email          string            `json:"email"`
	Mobile         string            `json:"mobile"`
	Status         string            `json:"status"`
	EmployeeNumber string            `json:"employee_number"`
	Department     *Department       `json:"department,omitempty"`
	Job            *Job              `json:"job,omitempty"`
	Preferences    map[string]string `json:"preferences,omitempty"` // myself 特有
	Permissions    map[string]bool   `json:"permissions,omitempty"` // myself 特有
}

// Department 部门
type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Job 职位
type Job struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserFilter 用户列表过滤条件
type UserFilter struct {
	Name          string // 用户名（精确匹配）
	Keywords      string // 关键词模糊搜索
	DepartmentIDs string // 部门ID列表，逗号分隔
}
