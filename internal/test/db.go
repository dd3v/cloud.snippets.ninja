package test

import (
	"os"
	"testing"

	"github.com/dd3v/cloud.snippets.ninja/internal/config"
	"github.com/dd3v/cloud.snippets.ninja/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

var db *dbcontext.DB

//Database - ...
func Database(t *testing.T) *dbcontext.DB {
	if db != nil {
		return db
	}
	cfg, err := config.Load("../../cfg/local.yml")
	if err != nil {
		t.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}
	mysql, err := dbx.MustOpen("mysql", cfg.TestDatabaseDNS)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	db := dbcontext.New(mysql)
	return db
}

//TruncateTable - ...
func TruncateTable(t *testing.T, db *dbcontext.DB, table string) {
	_, err := db.DB().TruncateTable(table).Execute()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = db.DB().NewQuery("ALTER TABLE " + table + " AUTO_INCREMENT = 1").Execute()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
