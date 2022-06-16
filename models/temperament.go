package models

import (
	"strings"
)

type TemperamentModel struct {
	Id   int64
	Sex  int64
	Name string
}

var TemperamentMap GlobalMapStruct = make(GlobalMapStruct)

//初始化星座表到map
func SetTemperamentMap() error {
	langs, err := GetAllLanguages(false)
	if err != nil {
		return err
	}
	for code, _ := range langs {
		SetTemperamentMapByLang(code, false)
	}
	return nil
}

func SetTemperamentMapByLang(lang string, reset bool) error {
	lang = strings.ToLower(lang)
	_, ok := TemperamentMap[lang]
	if ok && reset == false {
		return nil
	}
	var dts []*TemperamentModel
	rs := DB.Table(lang + "_temperaments").Find(&dts)
	if rs.Error != nil {
		return rs.Error
	}
	cl := make(map[int64]string)
	for _, v := range dts {
		cl[v.Id] = v.Name
	}
	TemperamentMap[lang] = cl
	return nil
}
