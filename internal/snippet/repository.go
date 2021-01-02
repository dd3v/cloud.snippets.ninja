package snippet

import (
	"context"
	"fmt"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

//Repository - ...
type Repository interface {
	GetByUserID(context context.Context, userID int, limit int, offset int, conditions map[string]interface{}) ([]entity.Snippet, error)
	GetByIDAndUserID(context context.Context, ID int, UserID int) (entity.Snippet, error)
	Query(context context.Context, limit int, offset int, conditions map[string]interface{}) ([]entity.Snippet, error)
	FindByID(context context.Context, id int) (entity.Snippet, error)
	Create(context context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Update(context context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Delete(context context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Count(context context.Context) (int, error)
}

type repository struct {
	db *dbcontext.DB
}

//NewRepository - ...
func NewRepository(db *dbcontext.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) GetByIDAndUserID(context context.Context, ID int, UserID int) (entity.Snippet, error) {
	var snippet entity.Snippet
	err := r.db.With(context).Select().Where(dbx.HashExp{"id": ID, "user_id": UserID}).One(&snippet)
	return snippet, err
}

func (r repository) GetByUserID(ctx context.Context, userID int, limit int, offset int, conditions map[string]interface{}) ([]entity.Snippet, error) {
	var snippets []entity.Snippet
	query := r.db.With(ctx).Select().Limit(int64(limit)).Offset(int64(offset)).Where(dbx.HashExp{"user_id": userID})
	if access, exists := conditions["access"]; exists {
		query.AndWhere(dbx.HashExp{"access": access})
	}
	if favorite, exists := conditions["favorite"]; exists {
		query.AndWhere(dbx.HashExp{"favorite": favorite})
	}
	if title, exists := conditions["title"]; exists {
		query.AndWhere(dbx.Like("title", title.(string)))
	}
	err := query.All(&snippets)
	return snippets, err
}

func (r repository) Query(context context.Context, limit int, offset int, conditions map[string]interface{}) ([]entity.Snippet, error) {
	var snippets []entity.Snippet

	fmt.Println(conditions)

	err := r.db.With(context).Select().Limit(int64(limit)).Offset(int64(offset)).All(&snippets)

	return snippets, err
}

func (r repository) FindByID(context context.Context, id int) (entity.Snippet, error) {
	var snippet entity.Snippet
	err := r.db.With(context).Select().Model(id, &snippet)
	return snippet, err
}

func (r repository) Create(context context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	err := r.db.With(context).Model(&snippet).Insert()
	return snippet, err
}

func (r repository) Update(context context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	err := r.db.With(context).Model(&snippet).Update()
	return snippet, err
}

func (r repository) Delete(context context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	err := r.db.With(context).Model(&snippet).Delete()
	return snippet, err
}

func (r repository) Count(context context.Context) (int, error) {
	var count int
	err := r.db.With(context).Select("COUNT(*)").From("snippets").Row(&count)
	return count, err
}
