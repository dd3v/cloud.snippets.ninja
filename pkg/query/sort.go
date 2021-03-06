package query

type Sort interface {
	GetSortBy() string
	GetOrderBy() string
}

type sort struct {
	sortBy  string
	orderBy string
}

func NewSort(sortBy string, orderBy string) Sort {
	return sort{sortBy, orderBy}
}

func (s sort) GetSortBy() string {
	return s.sortBy
}

func (s sort) GetOrderBy() string {
	return s.orderBy
}
