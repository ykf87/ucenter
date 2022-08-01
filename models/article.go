package models

import (
	"encoding/json"
	"html/template"
	"log"
	"ucenter/app/config"

	carbon "github.com/golang-module/carbon/v2"
	"github.com/tidwall/gjson"
)

type ArticleModel struct {
	Id      int64  `json:"id"`
	Key     string `json:"key"`
	Title   string `json:"title"`
	Keyword string `json:"keyword"`
	Desc    string `json:"desc"`
	Cont    string `json:"cont"`
	Addtime int64  `json:"addtime"`
	pin     int64  `json:"pin"`
	status  int64  `json:"status"`
	Sort    int64  `json:"sort"`
	Views   int64  `json:"views"`
}

func GetArticleRow(id int64, key, lang string) *ArticleModel {
	tableName := "`" + lang + "_articles`"
	dbob := DB.Table(tableName)
	if id > 0 {
		dbob = dbob.Where("id = ?", id)
	} else if key != "" {
		dbob = dbob.Where("`key` = ?", key)
	} else {
		return nil
	}
	art := new(ArticleModel)
	dbob.Where("status = 1").Find(art)
	if art == nil || art.Id < 1 {
		if config.Config.Lang != lang {
			tableName = "`" + config.Config.Lang + "_articles`"
			dbob.Table(tableName)
			dbob.Find(art)
		} else {
			return nil
		}
		if art == nil || art.Id < 1 {
			return nil
		}
	}
	go DB.Table(tableName).Where("id = ?", art.Id).Update("views", (art.Views + 1)) //增加浏览量
	return art
}

//格式化内容
func (this *ArticleModel) Fmt(lang, timezone string) map[string]interface{} {
	var fmt string
	fmts, ok := config.Config.Timefmts[lang]
	if ok {
		fmt = fmts.Datetimefmt
	} else {
		fmt = config.Config.Datetimefmt
	}

	b, err := json.Marshal(this)
	if err != nil {
		log.Println(err, "格式化 article 时出错, models/article.go")
		return nil
	}

	dt := make(map[string]interface{})

	for k, v := range gjson.ParseBytes(b).Map() {
		if k == "addtime" {
			dt[k] = carbon.CreateFromTimestamp(v.Int()).SetTimezone(timezone).Carbon2Time().Format(fmt)
		} else {
			dt[k] = template.HTML(v.String())
		}
	}
	return dt
}
