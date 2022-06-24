package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"ucenter/app/config"
)

type CityModel struct {
	Id         int64  `json:"id"`
	CountryId  int64  `json:"country_id"`
	ProvinceId int64  `json:"province_id"`
	Name       string `json:"name"`
}

func GetCityById(id int64, lang string) (*CityModel, error) {
	if lang == "" {
		lang = config.Config.Lang
	}
	lang = strings.ToLower(lang)
	tbs := new(CityModel)
	rs := DB.Table(lang+"_cities").Where("id = ?", id).First(tbs)
	if rs.Error == nil {
		return tbs, nil
	}
	return nil, errors.New("City not found")
}

func GetCityByNameAndCountryId(name string, countryId int64, lang string) (*CityModel, error) {
	if lang == "" {
		lang = config.Config.Lang
	}
	rs := new(CityModel)
	DB.Table(lang+"_cities").Where("country_id = ? and name = ?", countryId, name).First(rs)
	if rs.Id > 0 {
		return rs, nil
	}
	return nil, errors.New("City not found")
}

//获取国家下的城市
func GetCityByCountryId(lang string, countryid int64) ([]*CityModel, error) {
	if lang == "" || countryid < 1 {
		return nil, errors.New("Missing queries")
	}

	var lrs []*CityModel
	rs := DB.Table(lang+"_cities").Where("country_id = ?", countryid).Find(&lrs)
	if rs.Error != nil {
		log.Println(rs.Error, " - 获取国家城市列表出错!")
		return nil, errors.New("No city list found")
	}
	if lrs == nil {
		return nil, errors.New("No city list found")
	}
	return lrs, nil
}

func GetCityFilterAndPage(lang, filter string, countryid int64, provinceid, page, limit int, kv string) (dts interface{}, err error) {
	if countryid < 1 && provinceid < 1 {
		err = errors.New("Missing queries")
		return
	}
	if page < 1 {
		page = 1
	}
	if limit == 0 {
		limit = config.Config.Limit
	}

	dbs := DB.Table(lang + "_cities")
	if countryid > 0 {
		dbs = dbs.Where("country_id = ?", countryid)
	}
	if provinceid > 0 {
		dbs = dbs.Where("province_id = ?", provinceid)
	}
	if limit > 0 {
		dbs = dbs.Limit(limit).Offset((page - 1) * limit)
	}
	if filter != "" {
		dbs = dbs.Where("name like ?", "%"+filter+"%")
	}
	dbs = dbs.Order("name ASC")
	var ctmds []*CityModel
	rs := dbs.Find(&ctmds)
	if rs.Error != nil {
		err = rs.Error
		dts = nil
	} else {
		if len(ctmds) < 1 {
			return nil, errors.New("No results found")
		}
		if kv != "" {
			ddzzs := make(map[string]interface{})
			for _, v := range ctmds {
				ddzzs[fmt.Sprintf("%d", v.Id)] = v.Name
			}
			dts = ddzzs
		} else {
			dts = ctmds
		}
	}
	return
}
