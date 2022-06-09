package models

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var DefaultSqliteFile = "db.db"

func Init(dbtype, dsn, dbfile string) (err error) {
	switch dbtype {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
	case "sqlite":
		if dbfile == "" {
			dbfile = DefaultSqliteFile
		}
		DB, err = gorm.Open(sqlite.Open(dbfile), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
	default:
		err = errors.New("数据库格式不支持!")
	}

	return
}
