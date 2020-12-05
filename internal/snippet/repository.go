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
	List(context context.Context, limit int, offset int) ([]entity.Snippet, error)
	ListByUserID(context context.Context, conditions map[string]interface{}, limit int, offset int) ([]entity.Snippet, error)
	Query(context context.Context, conditions map[string]interface{}, limit int, offset int) ([]entity.Snippet, error)
	FindByID(context context.Context, id int) (entity.Snippet, error)
	Create(context context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Update(context context.Context, snippet entity.Snippet) error
	Delete(context context.Context, id int) error
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

func (r repository) GetByUserID(ctx context.Context, userID int, limit int, offset int, conditions map[string]interface{}) ([]entity.Snippet, error) {

	var snippets []entity.Snippet

	fmt.Println("repository call")
	fmt.Println("FAVORITE VALUES")
	fmt.Println(conditions["favorite"])
	if conditions["favorite"] == false {
		fmt.Print("favorite is false")
	}
	if conditions["favorite"] == true {
		fmt.Print("favorite is true")
	}

	query := r.db.With(ctx).Select().Limit(int64(limit))

	query.All(&snippets)

	fmt.Println(snippets)

	return snippets, nil
}

func (r repository) Query(context context.Context, conditions map[string]interface{}, limit int, offset int) ([]entity.Snippet, error) {
	var snippets []entity.Snippet

	fmt.Println(conditions)

	err := r.db.With(context).Select().Limit(int64(limit)).Offset(int64(offset)).All(&snippets)

	return snippets, err
}

func (r repository) List(context context.Context, limit int, offset int) ([]entity.Snippet, error) {
	var snippets []entity.Snippet
	err := r.db.With(context).Select().Limit(int64(limit)).Offset(int64(offset)).All(&snippets)
	return snippets, err
}

func (r repository) ListByUserID(context context.Context, conditions map[string]interface{}, limit int, offset int) ([]entity.Snippet, error) {

	var snippets []entity.Snippet

	c := dbx.HashExp(conditions)

	err := r.db.With(context).Select().Where(c).All(&snippets)

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

func (r repository) Update(context context.Context, snippet entity.Snippet) error {
	return r.db.With(context).Model(&snippet).Update()
}

func (r repository) Delete(context context.Context, id int) error {
	snippet, err := r.FindByID(context, id)
	if err != nil {
		return err
	}
	return r.db.With(context).Model(&snippet).Delete()
}

func (r repository) Count(context context.Context) (int, error) {
	var count int
	err := r.db.With(context).Select("COUNT(*)").From("snippets").Row(&count)
	return count, err
}
