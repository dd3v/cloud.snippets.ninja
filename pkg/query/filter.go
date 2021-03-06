package query

//Filter - query condition stack
type Filter interface {
	Push(condition string, field string, value string)
	List() map[string]expression
	Size() int
}

type expression struct {
	Condition string
	Field     string
	Value     string
}

type filter map[string]expression

func NewFilter() Filter {
	return filter{}
}

func (f filter) List() map[string]expression {
	return f
}

func (f filter) Push(condition string, field string, value string) {
	f[field] = expression{condition, field, value}
}

func (f filter) Size() int {
	return len(f)
}
