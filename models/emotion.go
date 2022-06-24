//情感状态
package models

import (
	"strings"
	"ucenter/app/config"
)

var EmotionMap GlobalMapStruct = make(GlobalMapStruct)

//初始化情感状态到map
func SetEmotionMap() error {
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
	EmotionMap = ttmmp
	return nil
}

func setEmotionMapByLang(lang string) (map[int64]string, error) {
	var dts []*IdNameModel
	rs := DB.Table(lang + "_emotions").Find(&dts)
	if rs.Error != nil {
		return nil, rs.Error
	}
	cl := make(map[int64]string)
	for _, v := range dts {
		cl[v.Id] = v.Name
	}
	return cl, nil
}

//返回情感状态列表
func EmotionList(lang, filter, kv string) (ddt interface{}) {
	ors, ok := EmotionMap[lang]
	if !ok {
		ors, ok = EmotionMap[strings.ToLower(config.Config.Lang)]
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
