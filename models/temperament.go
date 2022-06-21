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

//获取所有性格
func GetAllTemperaments(lang, filter, kv string, sex int64) (dt interface{}, err error) {
	tbName := strings.ToLower(lang + "_temperaments")
	dbObject := DB.Table(tbName)
	if filter != "" {
		dbObject = dbObject.Where("name like ?", "%"+filter+"%")
	}
	if sex > 0 {
		dbObject = dbObject.Where("sex = ?", sex)
	}

	var vser []*TemperamentModel
	rs := dbObject.Find(&vser)
	if rs.Error != nil {
		err = rs.Error
		return
	}

	if kv != "" {
		bbds := make(map[int64]map[int64]string)
		for _, v := range vser {
			bbds[v.Sex][v.Id] = v.Name
		}
		dt = bbds
	} else {
		type pertzcsd struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		}
		bbds := make(map[string][]*pertzcsd)
		for _, v := range vser {
			var kstr string
			if v.Sex == 1 {
				kstr = "male"
			} else {
				kstr = "female"
			}
			bbds[kstr] = append(bbds[kstr], &pertzcsd{Id: v.Id, Name: v.Name})
		}
		dt = bbds
	}
	return
}
