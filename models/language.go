package models

import (
	"strings"
	"ucenter/app/config"

	"github.com/gin-gonic/gin"
)

type LanguageModel struct {
	Id   int64  `json:"id"`
	Iso  string `json:"iso"`
	Name string `json:"name"`
}

var LangLists map[string]*LanguageModel

func GetAllLanguages(reget bool) (list map[string]*LanguageModel, err error) {
	if LangLists != nil && reget == false {
		list = LangLists
		return
	}
	var langs []*LanguageModel
	rs := DB.Table("languages").Where("status = 1").Order("sort DESC").Find(&langs)
	if rs.Error != nil {
		err = rs.Error
		return
	}
	list = make(map[string]*LanguageModel)
	for _, v := range langs {
		list[strings.ToLower(v.Iso)] = v
	}
	LangLists = list
	return
}

func GetAllLanguagesResKv() (list map[string]string) {
	list = make(map[string]string)
	for _, v := range LangLists {
		list[v.Iso] = v.Name
	}
	return
}

//获取客户端语言
func GetClientLang(c *gin.Context) string {
	var lang string
	if cc, err := c.GetQuery("lang"); err == true {
		lang = cc
	} else if c.GetHeader("lang") != "" {
		lang = c.GetHeader("lang")
	} else if cc, err := c.Cookie("lang"); err == nil {
		lang = cc
	} else if c.GetHeader("Accept-Language") != "" {
		langs := strings.Split(c.GetHeader("Accept-Language"), ",")
		lang = langs[0]
	} else {
		lang = config.Config.Lang
	}
	lang = strings.ToLower(lang)
	if _, ok := LangLists[lang]; !ok {
		lang = strings.ToLower(config.Config.Lang)
	}

	return lang
}
