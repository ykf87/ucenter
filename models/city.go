package models

import (
	"errors"
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
	return nil, rs.Error
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

func GetCityFilterAndPage(lang, filter string, countryid int64, provinceid, page, limit int) (dts map[string]interface{}, err error) {
	if countryid < 1 || provinceid < 1 {
		err = errors.New("Missing queries")
		return
	}
	if page < 1 {
		page = 1
	}
	if limit == 0 {
		limit = config.Config.Limit
	}

	lang = strings.ToLower(lang)
	dbs := DB.Table(lang+"_cities").Where("country_id = ? and province_id = ?", countryid, provinceid).Order("name ASC")
	if limit > 0 {
		dbs = dbs.Limit(limit).Offset((page - 1) * limit)
	}
	if filter != "" {
		dbs = dbs.Where("name like ?", "%"+filter+"%")
	}
	rs := dbs.Find(&dts)
	if rs.Error != nil {
		err = rs.Error
		dts = nil
	}
	return
}
