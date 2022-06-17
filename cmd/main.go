// linux execute file
// env GOOS=linux GOARCH=amd64 go build
// export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && go build
package main

import (
	"flag"
	"log"
	"ucenter/app"
	"ucenter/app/config"
	"ucenter/app/mails/smtp"
	"ucenter/app/safety/aess"
	"ucenter/app/safety/rsautil"
	"ucenter/models"
)

var configFile = flag.String("c", "config.yaml", "配置文件路径")

func main() {
	flag.Parse()
	if *configFile == "" {
		log.Println("请指定配置文件")
		return
	}
	err := config.Init(*configFile)
	if err != nil {
		log.Println(err)
		return
	}
	err = models.Init(config.Config.DB[0].Type, config.Config.DB[0].Dsn, config.Config.DB[0].Path)
	if err != nil {
		log.Println(err)
		return
	}
	for k, v := range config.Config.Smtp {
		smtp.SetConfig(k, v)
	}
	if config.Config.Aeskey != "" {
		aess.AESKEY = []byte(config.Config.Aeskey)
	}
	rsautil.Generate()
	app.App.Static(config.Config.Static, config.Config.Staticname).Run(config.Config.Port)
}
