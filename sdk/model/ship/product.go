package ship

// Product 产品模型
type Product struct {
	ID          string         // 产品 ID
	URL         string         // 产品资源地址
	Identifier  string         // 产品标识（如 "SLC"）
	Name        string         // 产品名称
	ScopeType   string         // 所属类型：organization, user_group
	ScopeID     string         // 所属 ID
	Visibility  string         // 可见性：private, public
	Color       string         // 主题色
	Description string         // 产品描述
	CreatedAt   int64          // 创建时间（10位时间戳）
	UpdatedAt   int64          // 更新时间

	// 关联对象
	Members   []ProductMember // 成员列表
	CreatedBy *User           // 创建人
	UpdatedBy *User           // 更新人

	// 状态字段
	IsArchived bool // 是否已归档
	IsDeleted  bool // 是否已删除
}

// ProductMember 产品成员
type ProductMember struct {
	ID        string     // 成员 ID
	URL       string     // 成员资源地址
	Type      string     // 类型：user, user_group
	User      *User      // 当 type=user 时
	UserGroup *UserGroup // 当 type=user_group 时
}

// User 用户引用
type User struct {
	ID          string // 用户 ID
	Name        string // 用户名
	DisplayName string // 显示名称
	Avatar      string // 头像 URL
}

// UserGroup 团队引用
type UserGroup struct {
	ID   string // 团队 ID
	Name string // 团队名称
}

// ProductFilter 产品列表过滤条件
type ProductFilter struct {
	Query     string // 查询关键字（可选）
	PageSize  int    // 每页大小（可选，默认30）
	PageIndex int    // 页码（可选，从0开始）
}
