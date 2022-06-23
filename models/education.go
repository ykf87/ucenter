//学历
package models

import (
	"strings"
	"ucenter/app/config"
)

var EducationMap GlobalMapStruct = make(GlobalMapStruct)

//初始化星座表到map
func SetEducationMap() error {
	langs, err := GetAllLanguages(false)
	if err != nil {
		return err
	}
	ttmmp := make(GlobalMapStruct)
	for code, _ := range langs {
		tmp, err := setEducationMapByLang(code)
		if err == nil {
			ttmmp[code] = tmp
		}
	}
	EducationMap = ttmmp
	return nil
}

func setEducationMapByLang(lang string) (map[int64]string, error) {
	var dts []*IdNameModel
	rs := DB.Table(lang + "_educations").Find(&dts)
	if rs.Error != nil {
		return nil, rs.Error
	}
	cl := make(map[int64]string)
	for _, v := range dts {
		cl[v.Id] = v.Name
	}
	return cl, nil
}

//返回学历列表
func EducationList(lang, filter, kv string) (ddt interface{}) {
	ors, ok := EducationMap[lang]
	if !ok {
		ors, ok = EducationMap[strings.ToLower(config.Config.Lang)]
		if !ok {
			return nil
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
}
