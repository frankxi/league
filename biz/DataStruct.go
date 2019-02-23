package biz

// 分页对象
type Page struct {
	PageNo   int `json:"pageNo" binding:"required"`   // 当前页码
	PageSize int `json:"pageSize" binding:"required"` // 每页条数
	Total    int `json:"total"`                       // 总条数
}
