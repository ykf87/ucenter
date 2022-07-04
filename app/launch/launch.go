package launch

import (
	"fmt"
	"ucenter/app"
	"ucenter/app/config"
	"ucenter/app/logs"
	"ucenter/app/mails/smtp"
	"ucenter/app/safety/aess"
	"ucenter/app/safety/rsautil"
	"ucenter/models"
)

func Start(filename string, port int) {
	err := config.Init(filename)
	if err != nil {
		logs.Logger.Error(err)
		return
	}
	config.Cpath = filename
	err = models.Init(config.Config.DB[0].Type, config.Config.DB[0].Dsn, config.Config.DB[0].Path)
	if err != nil {
		logs.Logger.Error(err)
		return
	}
	for k, v := range config.Config.Smtp {
		smtp.SetConfig(k, v)
	}
	if config.Config.Aeskey != "" {
		aess.AESKEY = []byte(config.Config.Aeskey)
	}
	rsautil.Generate()

	var runport int
	if port > 0 {
		runport = port
	} else {
		runport = config.Config.Port
	}

	if runport < 1000 {
		logs.Logger.Error(fmt.Sprintf("端口号不能小于1000，当前端口号为：%d", runport))
		return
	}

	go func() {
		app.App.Static(config.Config.Static, config.Config.Staticname).Template("templates/*").Run(runport)
	}()
	<-config.Och
	logs.Logger.Info("Panic from post!")
}
