package models

import (
	"strings"
	"ucenter/app/config"
)

type TemperamentModel struct {
	Id   int64
	Sex  int64
	Name string
}

var TemperamentMap GlobalMapStruct = make(GlobalMapStruct)
var TemperamentSexMap map[string]map[int64]map[int64]string = make(map[string]map[int64]map[int64]string)

//初始化性格表到map
func SetTemperamentMap() error {
	langs, err := GetAllLanguages(false)
	if err != nil {
		return err
	}
	tmmpp := make(GlobalMapStruct)
	tsmp := make(map[string]map[int64]map[int64]string)
	for code, _ := range langs {
		rrs, err := setTemperamentMapByLang(code)
		if err == nil {
			cl := make(map[int64]string)
			c2 := make(map[int64]map[int64]string)
			for _, v := range rrs {
				cl[v.Id] = v.Name
				if _, ok := c2[v.Sex]; !ok {
					c2[v.Sex] = make(map[int64]string)
				}
				c2[v.Sex][v.Id] = v.Name
			}
			tmmpp[code] = cl
			tsmp[code] = c2
		}
	}
	TemperamentMap = tmmpp
	TemperamentSexMap = tsmp
	return nil
}

func setTemperamentMapByLang(lang string) (dts []*TemperamentModel, err error) {
	rs := DB.Table(lang + "_temperaments").Find(&dts)
	if rs.Error != nil {
		err = rs.Error
		return
	}
	// cl := make(map[int64]string)
	// for _, v := range dts {
	// 	cl[v.Id] = v.Name
	// }
	return
}

//获取所有性格
func GetAllTemperaments(lang, filter, kv string, sex int64) (ddt interface{}) {
	ors, ok := TemperamentSexMap[lang]
	flang := lang
	if !ok {
		flang = strings.ToLower(config.Config.Lang)
		ors, ok = TemperamentSexMap[flang]
		if !ok {
			return nil
		}
	}
	if sex > 0 {
		bbqt, ok := ors[sex]
		if !ok {
			return nil
		}
		if kv == "" {
			var bfdgdf []*IdNameModel
			for id, val := range bbqt {
				if filter != "" && strings.Contains(val, filter) == false {
					continue
				}
				ccdp := new(IdNameModel)
				ccdp.Id = id
				ccdp.Name = val
				bfdgdf = append(bfdgdf, ccdp)
			}
			ddt = bfdgdf
		} else {
			if filter != "" {
				werf := make(map[int64]string)
				for id, val := range bbqt {
					if strings.Contains(val, filter) {
						werf[id] = val
					}
				}
				ddt = werf
			} else {
				ddt = bbqt
			}
		}
		return
	}
	if kv != "" {
		dt := make(map[string]map[int64]string)
		for sxid, v := range ors {
			var strr string
			if sxid == 2 {
				strr = "female"
			} else {
				strr = "male"
			}
			if _, ojbk := dt[strr]; !ojbk {
				dt[strr] = make(map[int64]string)
			}
			for iid, zv := range v {
				if filter != "" && strings.Contains(zv, filter) == false {
					continue
				}
				dt[strr][iid] = zv
			}
		}
		ddt = dt
	} else {
		dt := make(map[string][]*IdNameModel)
		for sxid, v := range ors {
			var strr string
			if sxid == 2 {
				strr = "female"
			} else {
				strr = "male"
			}
			for iid, zv := range v {
				if filter != "" && strings.Contains(zv, filter) == false {
					continue
				}
				ssass := new(IdNameModel)
				ssass.Id = iid
				ssass.Name = zv
				dt[strr] = append(dt[strr], ssass)
			}
			ddt = dt
			// if filter != "" && strings.Contains(v, filter) == false {
			// 	continue
			// }
			// lls := make(map[string]interface{})
			// lls["id"] = k
			// lls["name"] = v
			// dt = append(dt, lls)
		}
	}
	return
	// tbName := strings.ToLower(lang + "_temperaments")
	// dbObject := DB.Table(tbName)
	// if filter != "" {
	// 	dbObject = dbObject.Where("name like ?", "%"+filter+"%")
	// }
	// if sex > 0 {
	// 	dbObject = dbObject.Where("sex = ?", sex)
	// }

	// var vser []*TemperamentModel
	// rs := dbObject.Find(&vser)
	// if rs.Error != nil {
	// 	err = rs.Error
	// 	return
	// }

	// if kv != "" {
	// 	bbds := make(map[int64]map[int64]string)
	// 	for _, v := range vser {
	// 		bbds[v.Sex][v.Id] = v.Name
	// 	}
	// 	dt = bbds
	// } else {
	// 	type pertzcsd struct {
	// 		Id   int64  `json:"id"`
	// 		Name string `json:"name"`
	// 	}
	// 	bbds := make(map[string][]*pertzcsd)
	// 	for _, v := range vser {
	// 		var kstr string
	// 		if v.Sex == 1 {
	// 			kstr = "male"
	// 		} else {
	// 			kstr = "female"
	// 		}
	// 		bbds[kstr] = append(bbds[kstr], &pertzcsd{Id: v.Id, Name: v.Name})
	// 	}
	// 	dt = bbds
	// }
	// return
}
