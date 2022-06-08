package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(dbtype, dsn string) (err error) {
	DB, err = sql.Open(dbtype, dsn)
	return
}
