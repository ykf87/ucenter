package models

import (
	"errors"
	"strings"
	"ucenter/app/config"
)

type ProvinceModel struct {
	Id        int64  `json:"id"`
	CountryId int64  `json:"country_id"`
	Name      string `json:"name"`
}

func GetProvinceById(id int64, lang string) (*ProvinceModel, error) {
	if lang == "" {
		lang = config.Config.Lang
	}
	lang = strings.ToLower(lang)
	tbs := new(ProvinceModel)
	rs := DB.Table(lang+"_provinces").Where("id = ?", id).First(tbs)
	if rs.Error == nil {
		return tbs, nil
	}
	return nil, rs.Error
}

func GetProvinceByNameAndCountry(country_id int64, name, lang string) (*ProvinceModel, error) {
	if lang == "" {
		lang = config.Config.Lang
	}
	rs := new(ProvinceModel)
	DB.Table(lang+"_provinces").Where("country_id = ? and name = ?", country_id, name).First(rs)
	if rs.Id > 0 {
		return rs, nil
	}
	return nil, errors.New("City not found")
}

func GetProvinceByFilterAndPage(lang, filter string, countryid, page, limit int) (dts []map[string]interface{}, err error) {
	if countryid < 1 {
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
	dbs := DB.Table(lang+"_provinces").Where("country_id = ?", countryid).Order("name ASC")
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
