package models

import (
	"errors"
	"fmt"
	"math/big"
	"net"

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

func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}
