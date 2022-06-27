package launch

import (
	"log"
	"ucenter/app"
	"ucenter/app/config"
	"ucenter/app/mails/smtp"
	"ucenter/app/safety/aess"
	"ucenter/app/safety/rsautil"
	"ucenter/models"
)

func Start(filename string) {
	err := config.Init(filename)
	if err != nil {
		log.Println(err)
		return
	}
	config.Cpath = filename
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

	go func() {
		app.App.Static(config.Config.Static, config.Config.Staticname).Template("templates/*").Run(config.Config.Port)
	}()
	<-config.Och
	log.Println("Panic from post!")
}
