package test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/dd3v/snippets.page.backend/internal/config"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

var db *dbcontext.DB

//Database - ...
func Database(t *testing.T) *dbcontext.DB {
	if db != nil {
		return db
	}
	config := config.NewConfig()
	_, err := toml.DecodeFile("../../config/app.toml", config)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	pgsql, err := dbx.MustOpen("postgres", config.TestDatabaseDNS)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	db := dbcontext.New(pgsql)
	return db
}

//TruncateTable - ...
func TruncateTable(t *testing.T, db *dbcontext.DB, table string) {
	_, err := db.DB().TruncateTable(table).Execute()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = db.DB().NewQuery("ALTER SEQUENCE " + table + "_id_seq RESTART").Execute()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
