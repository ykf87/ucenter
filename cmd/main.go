package main

import (
	"flag"
	"fmt"
	"log"
	"ucenter/models"
)

var port = flag.Int("port", 18998, "开启的端口号")
var dsn = flag.String("dsn", "", "mysql配置,ex: \"root:111111@tcp(127.0.0.1:3306)/dome\"")
var db = flag.String("db", "mysql", "使用哪种数据库")
var dbfile = flag.String("dbfile", "", "sqlite文件路径,含文件名的完全路径")

func main() {
	flag.Parse()
	if *dsn == "" {
		log.Println("请填写数据库配置")
		return
	}
	err := models.Init(*db, *dsn, *dbfile)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(models.DB)
}
