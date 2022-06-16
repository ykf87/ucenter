package models

import (
	"errors"
	"net"
	"strings"
	"ucenter/app/config"

	"github.com/oschwald/geoip2-golang"
)

type CountryModel struct {
	Id        int64  `json:"id"`
	Iso3      string `json:"iso3"`
	Iso       string `json:"iso"`
	Phonecode string `json:"phonecode"`
	Currency  string `json:"currency"`
	Timezone  string `json:"timezone"`
	lat       string `json:"lat"`
	lon       string `json:"lon"`
	Emoji     string `json:"emoji"`
}

type CountryNameModel struct {
	Id   int64
	Name string
}

var CountryMap GlobalMapStruct = make(GlobalMapStruct)
var Countries = make(map[string]*CountryModel)

func InitCountry() error {
	langs, err := GetAllLanguages(false)
	if err != nil {
		return err
	}
	for code, _ := range langs {
		SetCountryMapByLang(code, false)
	}

	var countrys []*CountryModel
	DB.Table("countries").Find(&countrys)
	for _, v := range countrys {
		Countries[v.Iso] = v
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

func GetCountryByIso(iso string) (ct *CountryModel, err error) {
	rs, ok := Countries[iso]
	if !ok {
		err = errors.New("Country not found")
		return
	}
	ct = rs
	return
	// r := new(CountryModel)
	// DB.Table("countries").Where("iso", iso).First(r)
	// if r.Id > 0 {
	// 	return r, nil
	// }
	// return nil, errors.New("Country not found")
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

//根据ip获取国家
func GetCountryByIp(ip string) (ct *CountryModel, err error) {
	geodb, errs := geoip2.Open("./GeoLite2-City.mmdb")
	if errs == nil {
		defer geodb.Close()
		ip := net.ParseIP(ip)
		record, errs := geodb.City(ip)
		var iso string
		if errs == nil {
			if record.Country.IsoCode == "" {
				iso = config.Config.Country
			} else {
				iso = record.Country.IsoCode
			}
			ct, err = GetCountryByIso(iso)
		} else {
			err = errs
		}
	} else {
		err = errs
	}
	return
}
