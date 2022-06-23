package config

import (
	"github.com/jinzhu/configor"
)

var Och chan bool = make(chan bool)

type SmtpConf struct {
	Host   string
	Port   int
	Email  string
	Pass   string
	Sender string
}

type ImConf struct {
	Id  string
	Key string
}

type ConfigStruct struct {
	APPName          string `required:"true"`
	Domain           string `required:"true"`
	Port             int    `required:"true"`
	Lang             string `default:"en"` //默认语言
	Auther           string `default:"blandal.com@gmail.com"`
	Copyright        string
	Logo             string
	Static           string
	Staticname       string
	Aeskey           string
	Universalcaptcha string //万能验证码
	Limit            int    `default:20`
	Country          string `default:"US"`
	Timezone         string `default:"America/Adak"`
	Datetimefmt      string `default:"2006-01-02 15:04:05"`
	Datefmt          string `default:"2006-01-02"`
	Timefmt          string `default:"15:04:05"`
	Useim            string `required:"true"` //使用的im

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

	Im map[string]*ImConf `required:"true"`
}

var Config = new(ConfigStruct)
var Cpath string

func Init(path string) (err error) {
	cc := new(ConfigStruct)
	err = configor.Load(cc, path)
	if err == nil {
		Config = cc
	}
	return
}
