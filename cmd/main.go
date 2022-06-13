package main

import (
	"flag"
	// "fmt"
	"log"
	// "net"
	"ucenter/app"
	"ucenter/app/safety/rsautil"
	"ucenter/models"
	// "github.com/oschwald/geoip2-golang"
)

var port = flag.Int("port", 18998, "开启的端口号")
var dsn = flag.String("dsn", "", "mysql配置,ex: \"root:111111@tcp(127.0.0.1:3306)/dome\"")
var db = flag.String("db", "mysql", "使用哪种数据库")
var dbfile = flag.String("dbfile", "", "sqlite文件路径,含文件名的完全路径")
var staticPath = flag.String("path", "", "静态目录路径,不存在则会尝试创建,如: ./static")

func main() {
	// dbs, e := geoip2.Open("./GeoLite2-City.mmdb")
	// if e == nil {
	// 	defer dbs.Close()
	// 	ip := net.ParseIP("112.49.212.193")
	// 	record, _ := dbs.City(ip)
	// 	fmt.Println(record.City)
	// 	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["zh-CN"])
	// 	if len(record.Subdivisions) > 0 {
	// 		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	// 	}
	// 	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	// 	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	// 	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	// 	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	// } else {
	// 	fmt.Println(e)
	// }

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
	rsautil.Generate()
	// aaec, _ := rsautil.RsaEncrypt("aaaaaa")
	// bbec, _ := rsautil.RsaDecrypt(aaec)
	// log.Println(aaec, bbec)

	// ssac, _ := rsautil.Sign("yyyykkf", crypto.MD5)
	// ssbc := rsautil.Verify("aawer", ssac, crypto.MD5)
	// ssbz := rsautil.Verify("yyyykkf", ssac, crypto.MD5)
	// log.Println(ssac, ssbc, ssbz)
	app.App.Static(*staticPath).Run(*port)
	// fmt.Println(models.DB)
}
