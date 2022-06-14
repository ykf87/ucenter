package main

import (
	"flag"
	"log"
	"ucenter/app"
	"ucenter/app/config"
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

	// s := smtps.Client(config.Config.Smtp.Host, config.Config.Smtp.Email, config.Config.Smtp.Pass, config.Config.APPName, config.Config.Smtp.Port)
	// s.SetSender("blandal@foxmail.com").SetAtta("1.txt").SetSubject("subject111").SetGeter("blandal.com@gmail.com").SetGeter("1603601628@qq.com").SetMessage("收到反馈尽快答复").Send()

	rsautil.Generate()
	app.App.Static(config.Config.Static).Run(config.Config.Port)
}
