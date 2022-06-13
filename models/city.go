package models

import (
	"errors"
)

type CityModel struct {
	Id         int64  `json:"id"`
	CountryId  int64  `json:"country_id"`
	ProvinceId int64  `json:"province_id"`
	Name       string `json:"name"`
}

var TableLang = "en"

func GetCityByNameAndCountryId(name string, countryId int64, lang string) (*CityModel, error) {
	if lang == "" {
		lang = TableLang
	}
	rs := new(CityModel)
	DB.Table("cities_"+lang).Where("country_id = ? and name = ?", countryId, name).First(rs)
	if rs.Id > 0 {
		return rs, nil
	}
	return nil, errors.New("City not found")
}
