package query

type Pagination interface {
	GetPage() int
	GetLimit() int
	GetOffset() int
	GetTotalItems() int
	GetTotalPages() int
}

type pagination struct {
	page       int
	limit      int
	totalItem  int
	totalPages int
}

func NewPagination(page int, limit int, totalItems int) pagination {
	return pagination{page: page, limit: limit, totalItem: totalItems}
}

func (p pagination) GetPage() int {
	return p.page
}

func (p pagination) GetLimit() int {
	return p.limit
}

func (p pagination) GetOffset() int {
	return (p.page - 1) * p.limit
}

func (p pagination) GetTotalItems() int {
	return p.totalItem
}

func (p pagination) GetTotalPages() int {
	return (p.totalItem + p.limit - 1) / p.limit
}
