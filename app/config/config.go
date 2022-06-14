package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	APPName string `default:"Dome"`
	Port    int    `required:"true"`
	Lang    string `default:"en"` //默认语言
	Static  string

	DB []struct {
		Type string
		Dsn  string
		Path string
	}

	Smtp struct {
		Host  string
		Port  int
		Email string
		Pass  string
	}
}{}

func Init(path string) (err error) {
	err = configor.Load(&Config, path)
	return
}
