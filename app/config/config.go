package config

import (
	"github.com/jinzhu/configor"
)

type SmtpConf struct {
	Host  string
	Port  int
	Email string
	Pass  string
}

var Config = struct {
	APPName     string `required:"true"`
	Domain      string `required:"true"`
	Port        int    `required:"true"`
	Lang        string `default:"en"` //默认语言
	Auther      string `default:"blandal.com@gmail.com"`
	Copyright   string
	Static      string
	Staticname  string
	Aeskey      string
	Limit       int    `default:20`
	Country     string `default:"US"`
	Timezone    string `default:"America/Adak"`
	Datetimefmt string `default:"2006-01-02 15:04:05"`
	Datefmt     string `default:"2006-01-02"`
	Timefmt     string `default:"15:04:05"`

	DB []struct {
		Type string
		Dsn  string
		Path string
	} `required:"true"`

	Smtp map[string]*SmtpConf `required:"true"`

	Timefmts map[string]struct {
		Datetimefmt string
		Datefmt     string
		Timefmt     string
	}
}{}

func Init(path string) (err error) {
	err = configor.Load(&Config, path)
	return
}
