package models

import (
	"errors"
)

type ProvinceModel struct {
	Id        int64  `json:"id"`
	CountryId int64  `json:"country_id"`
	Name      string `json:"name"`
}

func GetProvinceByNameAndCountry(country_id int64, name, lang string) (*ProvinceModel, error) {
	if lang == "" {
		lang = TableLang
	}
	rs := new(ProvinceModel)
	DB.Table("provinces_"+lang).Where("country_id = ? and name = ?", country_id, name).First(rs)
	if rs.Id > 0 {
		return rs, nil
	}
	return nil, errors.New("City not found")
}
