package test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/dd3v/snippets.ninja/internal/config"
	"github.com/dd3v/snippets.ninja/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"

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
	mysql, err := dbx.MustOpen("mysql", config.TestDatabaseDNS)
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
