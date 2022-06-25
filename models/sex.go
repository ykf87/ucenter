package models

import (
	"ucenter/app/i18n"
)

var SEXBASELIST = []string{
	"Confidential",
	"Male",
	"Female",
	"Transgender",
}

//获取性别列表
func GetSexList(lang string) []*IdNameModel {
	var reslist []*IdNameModel
	for k, v := range SEXBASELIST {
		rr := new(IdNameModel)
		rr.Id = int64(k)
		rr.Name = string(i18n.T(lang, v))
		reslist = append(reslist, rr)
	}
	return reslist
}

//获取性别列表,返回键值对
func GetSexListByKv(lang string) (dt map[int]string) {
	dt = make(map[int]string)
	for k, v := range SEXBASELIST {
		dt[k] = string(i18n.T(lang, v))
	}
	return
}

//获取性别文本
func GetSexString(lang string, id int) string {
	if id < 0 || id >= len(SEXBASELIST) {
		return ""
	}
	str := string(i18n.T(lang, SEXBASELIST[id]))
	return str
}
