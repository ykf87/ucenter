package models

import (
	"strings"
)

type ApplangModel struct {
	Key  string `json:"key"`
	Path string `json:"path"`
	Val  string `json:"val"`
}

func GetAppLangs(lang string) map[string]interface{} {
	table := lang + "applangs"

	rrs := make(map[string]interface{})
	paths := make(map[string]map[string]string)
	var list []*ApplangModel
	if DB.Table(table).Find(&list).Error == nil {
		for _, v := range list {
			if v.Key == "" || v.Val == "" {
				continue
			}
			ks := strings.Split(v.Key, " ")
			ks[0] = strings.ToUpper(ks[0])
			k := strings.Join(ks, " ")
			if v.Path != "" {
				path := strings.ToLower(v.Path)
				if _, ok := paths[path]; !ok {
					paths[path] = make(map[string]string)
					paths[path][k] = v.Val
				}
			} else {
				rrs[k] = v.Val
			}
		}
		if len(paths) > 0 {
			for j, l := range paths {
				rrs[j] = l
			}
		}
	}
	return rrs
}
