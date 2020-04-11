package dbcontext

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type DB struct {
	db *dbx.DB
}

type TransactionFunc func(ctx context.Context, f func(ctx context.Context) error) error

type contextKey int

const (
	txKey contextKey = iota
)

func New(db *dbx.DB) *DB {
	return &DB{db}
}

func (db *DB) DB() *dbx.DB {
	return db.db
}

func (db *DB) With(ctx context.Context) dbx.Builder {
	if tx, ok := ctx.Value(txKey).(*dbx.Tx); ok {
		return tx
	}
	return db.db.WithContext(ctx)
}
