package models

import (
	"strings"
)

var IncomesMap map[int64]string

func GetAllIncomes(reget bool) (list map[int64]string, err error) {
	if IncomesMap != nil && reget == false {
		list = IncomesMap
		return
	}
	var incomes []*IdNameModel
	rs := DB.Table("incomes").Find(&incomes)
	if rs.Error != nil {
		err = rs.Error
		return
	}
	list = make(map[int64]string)
	for _, v := range incomes {
		list[v.Id] = v.Name
	}
	IncomesMap = list
	return
}

//返回收入列表
func IncomesList(filter, kv string) (ddt interface{}) {
	if kv != "" {
		if filter != "" {
			dt := make(map[int64]string)
			for k, v := range IncomesMap {
				if strings.Contains(v, filter) {
					dt[k] = v
				}
			}
			ddt = dt
		} else {
			ddt = IncomesMap
		}
	} else {
		var dt []map[string]interface{}
		for k, v := range IncomesMap {
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
