package models

import (
	"strings"
	"ucenter/app/config"
)

var ConstellationMap GlobalMapStruct = make(GlobalMapStruct)

//初始化星座表到map
func SetConstellationMap() error {
	langs, err := GetAllLanguages(false)
	if err != nil {
		return err
	}
	ttmmp := make(GlobalMapStruct)
	for code, _ := range langs {
		code = strings.ToLower(code)
		tmp, err := setConstellationMapByLang(code)
		if err == nil {
			ttmmp[code] = tmp
		}
	}
	ConstellationMap = ttmmp
	return nil
}

func setConstellationMapByLang(lang string) (map[int64]string, error) {
	var dts []*IdNameModel
	rs := DB.Table(lang + "_constellations").Find(&dts)
	if rs.Error != nil {
		return nil, rs.Error
	}
	cl := make(map[int64]string)
	for _, v := range dts {
		cl[v.Id] = v.Name
	}
	return cl, nil
}

//获取所有星座
func GetAllConstellations(lang, filter, kv string) (ddt interface{}) {
	ors, ok := ConstellationMap[lang]
	if !ok {
		ors, ok = ConstellationMap[strings.ToLower(config.Config.Lang)]
		if !ok {
			return
		}
	}
	if kv != "" {
		if filter != "" {
			dt := make(map[int64]string)
			for k, v := range ors {
				if strings.Contains(v, filter) {
					dt[k] = v
				}
			}
			ddt = dt
		} else {
			ddt = ors
		}
	} else {
		var dt []map[string]interface{}
		for k, v := range ors {
			if filter != "" && strings.Contains(v, filter) == false {
				continue
			}
			lls := make(map[string]interface{})
			lls["id"] = k
			lls["name"] = v
			dt = append(dt, lls)
		}
		ddt = dt
	}
	return
	// tbName := strings.ToLower(lang + "_constellations")
	// dbObject := DB.Table(tbName)

	// var vser []*IdNameModel
	// rs := dbObject.Find(&vser)
	// if rs.Error != nil {
	// 	err = rs.Error
	// 	return
	// }

	// if kv != "" {
	// 	bbds := make(map[int64]string)
	// 	for _, v := range vser {
	// 		bbds[v.Id] = v.Name
	// 	}
	// 	dt = bbds
	// } else {
	// 	dt = vser
	// }
	// return
}
