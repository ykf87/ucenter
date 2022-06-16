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

type CountryNameModel struct {
	Id   int64
	Name string
}

var CountryMap GlobalMapStruct = make(GlobalMapStruct)

func InitCountry() error {
	langs, err := GetAllLanguages(false)
	if err != nil {
		return err
	}
	for code, _ := range langs {
		SetCountryMapByLang(code, false)
	}
	return nil
}

func SetCountryMapByLang(lang string, reset bool) error {
	lang = strings.ToLower(lang)
	_, ok := CountryMap[lang]
	if ok && reset == false {
		return nil
	}
	var dts []*CountryNameModel
	rs := DB.Table(lang + "_countries").Find(&dts)
	if rs.Error != nil {
		return rs.Error
	}
	cl := make(map[int64]string)
	for _, v := range dts {
		cl[v.Id] = v.Name
	}
	CountryMap[lang] = cl
	return nil
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
	dbs := DB.Table(lang + "_countries").Order("name ASC")
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
