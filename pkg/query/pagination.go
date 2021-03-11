package query

type Pagination interface {
	GetPage() int
	GetLimit() int
	GetOffset() int
}

type pagination struct {
	page  int
	limit int
}

func NewPagination(page int, limit int) pagination {
	return pagination{page: page, limit: limit}
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
