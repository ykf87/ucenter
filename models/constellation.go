package models

import (
	"strings"
)

type ConstellationModel struct {
	Id   int64
	Name string
}

var ConstellationMap GlobalMapStruct = make(GlobalMapStruct)

//初始化星座表到map
func SetConstellationMap() error {
	langs, err := GetAllLanguages(false)
	if err != nil {
		return err
	}
	for code, _ := range langs {
		SetConstellationMapByLang(code, false)
	}
	return nil
}

func SetConstellationMapByLang(lang string, reset bool) error {
	lang = strings.ToLower(lang)
	_, ok := ConstellationMap[lang]
	if ok && reset == false {
		return nil
	}
	var dts []*ConstellationModel
	rs := DB.Table(lang + "_constellations").Find(&dts)
	if rs.Error != nil {
		return rs.Error
	}
	cl := make(map[int64]string)
	for _, v := range dts {
		cl[v.Id] = v.Name
	}
	ConstellationMap[lang] = cl
	return nil
}

//获取所有星座
func GetAllConstellations(lang, kv string) (dt interface{}, err error) {
	tbName := strings.ToLower(lang + "_constellations")
	dbObject := DB.Table(tbName)

	var vser []*ConstellationModel
	rs := dbObject.Find(&vser)
	if rs.Error != nil {
		err = rs.Error
		return
	}

	if kv != "" {
		bbds := make(map[int64]string)
		for _, v := range vser {
			bbds[v.Id] = v.Name
		}
		dt = bbds
	} else {
		dt = vser
	}
	return
}
