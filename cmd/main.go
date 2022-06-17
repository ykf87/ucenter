// linux execute file
// env GOOS=linux GOARCH=amd64 go build
// export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && go build
package main

import (
	"flag"
	"log"
	"ucenter/app/launch"
)

var configFile = flag.String("c", "config.yaml", "配置文件路径")

func main() {
	flag.Parse()
	if *configFile == "" {
		log.Println("请指定配置文件")
		return
	}
	launch.Start(*configFile)
}
