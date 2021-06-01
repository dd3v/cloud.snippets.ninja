package snippet

import (
	"context"
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	"github.com/dd3v/snippets.page.backend/pkg/query"
	dbx "github.com/go-ozzo/ozzo-dbx"
)


type repository struct {
	db *dbcontext.DB
}

//NewMockRepository - ...
func NewRepository(db *dbcontext.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) GetByID(ctx context.Context, id int) (entity.Snippet, error) {
	var snippet entity.Snippet
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"id": id}).One(&snippet)
	return snippet, err
}

func (r repository) QueryByUserID(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
	var snippets []entity.Snippet
	query := r.db.With(ctx).Select().Where(dbx.HashExp{"user_id": userID})
	query.Limit(int64(pagination.GetLimit())).Offset(int64(pagination.GetOffset()))
	for field, value := range filter {
		expression, err := r.buildExpression(field, value)
		if err != nil {
			return snippets, err
		}
		query.AndWhere(expression)
	}
	query.OrderBy(sort.GetSortBy() + " " + sort.GetOrderBy())
	err := query.All(&snippets)
	return snippets, err
}

func (r repository) Create(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	err := r.db.With(ctx).Model(&snippet).Insert()
	return snippet, err
}

func (r repository) Update(ctx context.Context, snippet entity.Snippet) error {
	return r.db.With(ctx).Model(&snippet).Exclude("UserID", "CreatedAt").Update()
}

func (r repository) Delete(ctx context.Context, snippet entity.Snippet) error {
	return r.db.With(ctx).Model(&snippet).Delete()
}

func (r repository) CountByUserID(ctx context.Context, userID int, filter map[string]string) (int, error) {
	var count int
	query := r.db.With(ctx).Select("COUNT(*)").From("snippets").Where(dbx.HashExp{"user_id": userID})
	for field, value := range filter {
		expression, err := r.buildExpression(field, value)
		if err != nil {
			return 0, err
		}
		query.AndWhere(expression)
	}
	err := query.Row(&count)
	return count, err
}

func (r repository) buildExpression(key string, value string) (dbx.Expression, error) {
	var expression dbx.Expression
	var err error
	switch key {
	case "favorite", "access_level", "language":
		expression = dbx.HashExp{key: value}
		break
	case "title":
		expression = dbx.NewExp("MATCH (title,content) AGAINST ({:keywords} IN BOOLEAN MODE)", dbx.Params{"keywords": value + "*"})
		break
	default:
		err = errors.New("Undefined filter key")
		break
	}
	return expression, err
}
