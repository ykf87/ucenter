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
	Id   int64  `json:"id"`
	Name string `json:"name"`
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

func GetCountryByFilterAndPage(lang, filter string, page, limit int, kv string) (dts interface{}, err error) {
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

	var nngdfg []*CountryNameModel
	rs := dbs.Find(&nngdfg)
	if rs.Error != nil {
		err = rs.Error
		dts = nil
	} else {
		if len(nngdfg) < 1 {
			return nil, errors.New("No results found")
		}
		if kv != "" {
			dtszz := make(map[string]interface{})
			for _, v := range nngdfg {
				for iso, b := range Countries {
					if v.Id == b.Id {
						dtszz[iso] = v.Name
						break
					}
				}
			}
			dts = dtszz
		} else {
			dts = nngdfg
		}
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

//返回phonecode
func CountryPhoneCode(lang, kv, filter string, page, limit int) (dts interface{}, err error) {
	lang = strings.ToLower(lang)
	if page < 1 {
		page = 1
	}
	if limit == 0 {
		limit = config.Config.Limit
	}
	type sdsds struct {
		Id        int64  `json:"id"`
		Iso       string `json:"iso"`
		Phonecode string `json:"phonecode"`
		Name      string `json:"name"`
	}
	var nds []*sdsds
	dbs := DB.Select("a.id, a.iso, a.phonecode, b.name").Table("countries as a").Joins("left join `" + lang + "_countries` as b on a.id = b.id")
	if filter != "" {
		dbs = dbs.Where("b.name like ?", "%"+filter+"%")
	}
	if limit > 0 {
		dbs = dbs.Limit(limit).Offset((page - 1) * limit)
	}
	rs := dbs.Find(&nds)
	if rs.Error != nil {
		return nil, errors.New("No results found")
	}
	if kv != "" {
		mps := make(map[string]interface{})
		for _, v := range nds {
			mps[v.Iso] = v.Phonecode
		}
		return mps, nil
	}
	return nds, nil
}
