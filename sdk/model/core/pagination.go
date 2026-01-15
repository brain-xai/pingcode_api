package core

// Pagination 分页信息
type Pagination struct {
	PageSize  int // 每页大小
	PageIndex int // 当前页码（从0开始）
	Total     int // 总记录数
}
