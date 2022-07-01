package models

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
	"ucenter/app/config"
	"ucenter/app/logs"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IdNameModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GlobalMapStruct map[string]map[int64]string
type GlobalMapString map[string]string

var DB *gorm.DB
var DefaultSqliteFile = "db.db"

func Init(dbtype, dsn, dbfile string) (err error) {
	logName := "./logs/db.log"
	src, errs := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if errs != nil {
		logs.Logger.Error(errs)
		err = errs
		return
	}
	newLogger := logger.New(
		log.New(src, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second * time.Duration(2), // 慢 SQL 阈值
			LogLevel:      logger.Error,                   // Log level
			Colorful:      false,                          // 禁用彩色打印
		},
	)

	switch dbtype {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			// Logger: logger.Default.LogMode(logger.Error),
			Logger: newLogger,
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
	return PrevLoadDB()
}

//数据库数据预加载到内存
func PrevLoadDB() (err error) {
	err = DBToCache(nil)
	if err != nil {
		return
	}

	go func() {
		ch := make(chan error)
		for {
			time.Sleep(time.Minute * 360)
			GetAllLanguages(true)
			DBToCache(ch)
			r := <-ch
			if r != nil {
				logs.Logger.Error("系统需要重启 - auto DBToCache fail: ", r)
				break
			}
		}
	}()
	return
}

//数据库持久化到内存
func DBToCache(ch chan error) (err error) {
	err = InitCountry()
	if err != nil {
		if ch != nil {
			ch <- err
		}
		return
	}
	err = SetConstellationMap()
	if err != nil {
		if ch != nil {
			ch <- err
		}
		return
	}
	err = SetTemperamentMap()
	if err != nil {
		if ch != nil {
			ch <- err
		}
		return
	}
	err = SetEducationMap()
	if err != nil {
		if ch != nil {
			ch <- err
		}
		return
	}
	err = SetEmotionMap()
	if err != nil {
		if ch != nil {
			ch <- err
		}
		return
	}
	_, err = GetAllIncomes(true)
	if err != nil {
		if ch != nil {
			ch <- err
		}
		return
	}
	if ch != nil {
		ch <- nil
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

func (this *GlobalMapStruct) Get(lang string, id int64) (name string) {
	dflang := strings.ToLower(config.Config.Lang)
	if lang == "" {
		lang = dflang
	}
	mp := *this

	v, ok := mp[lang]
	if !ok {
		if lang == dflang {
			return
		}
		v, ok = mp[dflang]
		if !ok {
			return
		}
	}
	nn, ok := v[id]
	if ok {
		name = nn
	}
	return
}
