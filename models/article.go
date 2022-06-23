package models

import (
	"strings"
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
	tableName := strings.ToLower(lang + "_articles")
	dbob := DB.Table(tableName)
	if id > 0 {
		dbob = dbob.Where("id = ?", id)
	} else if key != "" {
		dbob = dbob.Where("key = ?", key)
	} else {
		return nil
	}
	art := new(ArticleModel)
	dbob.Find(art)
	if art == nil || art.Id < 1 {
		return nil
	}
	return art
}
