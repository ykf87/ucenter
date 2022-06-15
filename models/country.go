package models

import (
	"errors"
	"strings"
	"ucenter/app/config"
)

type CountryModel struct {
	Id        int64   `json:"id"`
	Iso3      string  `json:"iso3"`
	Iso2      string  `json:"iso2"`
	Phonecode string  `json:"phonecode"`
	Currency  string  `json:"currency"`
	Region    int64   `json:"region"`
	Subregion int64   `json:"subregion"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Emoji     string  `json:"emoji"`
}

func GetCountryByIso(iso string) (*CountryModel, error) {
	r := new(CountryModel)
	DB.Table("countries").Where("iso", iso).First(r)
	if r.Id > 0 {
		return r, nil
	}
	return nil, errors.New("Country not found")
}

func GetCountryByFilterAndPage(lang, filter string, page, limit int) (dts []map[string]interface{}, err error) {
	if page < 1 {
		page = 1
	}
	if limit == 0 {
		limit = config.Config.Limit
	}

	lang = strings.ToLower(lang)
	dbs := DB.Table("countries_" + lang).Order("name ASC")
	if limit > 0 {
		dbs = dbs.Limit(limit).Offset((page - 1) * limit)
	}
	if filter != "" {
		dbs = dbs.Where("name like ?", "%"+filter+"%")
	}
	rs := dbs.Find(&dts)
	if rs.Error != nil {
		// if lang != "en" {
		// 	dbs = DB.Table("countries_en").Limit(limit).Offset((page - 1) * limit).Order("name DESC")
		// 	if filter != "" {
		// 		dbs = dbs.Where("name like ?", "%"+filter+"%")
		// 	}
		// 	rs = dbs.Find(&dts)
		// 	if rs.Error != nil {
		// 		err = rs.Error
		// 		dts = nil
		// 	}
		// } else {
		err = rs.Error
		dts = nil
		// }
	}
	return
}
