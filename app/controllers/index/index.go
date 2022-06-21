package index

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"ucenter/app/config"
	"ucenter/app/controllers"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Media(c *gin.Context) {
	uri := c.Param("path")
	path := "./media/" + uri
	content, err := os.ReadFile(path)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	fmt.Println(uri, string(content))
	controllers.Success(c, nil, nil)
}

func Country(c *gin.Context) {
	filter := c.Query("q")
	pc := strings.Trim(c.Param("procity"), "/")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	langs, _ := c.Get("_lang")
	lang := langs.(string)
	kv := c.Query("kv")
	var err error
	var rs interface{}

	if pc != "" {
		if strings.Contains(pc, "/") {
			ids := strings.Split(pc, "/")
			iso := strings.ToUpper(ids[0])
			provId, _ := strconv.Atoi(ids[1])
			country, ok := models.Countries[iso]
			if ok {
				rs, err = models.GetCityFilterAndPage(lang, filter, country.Id, provId, page, limit, kv)
			} else {
				err = errors.New("No results found")
			}
		} else {
			pc = strings.ToUpper(pc)
			country, ok := models.Countries[pc]
			if ok {
				rs, err = models.GetProvinceByFilterAndPage(lang, filter, country.Id, page, limit, kv)
			} else {
				err = errors.New("No results found")
			}
		}
	} else {
		rs, err = models.GetCountryByFilterAndPage(lang, filter, page, limit, kv)
	}

	if err != nil {
		controllers.Resp(c, nil, &controllers.Msg{Str: "No results found"}, 404)
	} else {
		controllers.Success(c, rs, nil)
	}
}

//获取一些列表
func Lists(c *gin.Context) {
	filter := strings.Trim(c.Query("q"), " ")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	tb := strings.Trim(c.Param("table"), "/ ")
	langobj, _ := c.Get("_lang")
	lang := langobj.(string)
	kv := c.Query("kv")
	if limit == 0 {
		limit = config.Config.Limit
	}
	if page < 0 {
		page = 1
	}

	if tb == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "No results found"})
		return
	}
	tbName := lang + "_" + tb
	tbName = strings.ToLower(tbName)

	dbObject := models.DB.Table(tbName)
	if tb == "temperaments" {
		sex := c.Query("sex")
		if sex != "" {
			dbObject = dbObject.Where("sex = ?", sex)
		}
	}
	if filter != "" {
		dbObject = dbObject.Where("name like ?", "%"+filter+"%")
	}

	type sspfd struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
	var dts []*sspfd
	if limit > 0 {
		dbObject = dbObject.Limit(limit).Offset((page - 1) * limit)
	}
	rs := dbObject.Find(&dts)
	if rs.Error != nil {
		log.Println(rs.Error)
		controllers.Error(c, nil, &controllers.Msg{Str: "No results found"})
		return
	}
	if kv != "" {
		ngst := make(map[int64]string)
		for _, v := range dts {
			ngst[v.Id] = v.Name
		}
		controllers.Success(c, ngst, &controllers.Msg{Str: "Success"})
	} else {
		controllers.Success(c, dts, &controllers.Msg{Str: "Success"})
	}
}

//获取系统支持的语言列表
func Languages(c *gin.Context) {
	kv := c.Query("kv")
	l, err := models.GetAllLanguages(false)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: "No results found"})
	} else {
		if kv != "" {
			mp := make(map[string]interface{})
			for _, v := range l {
				mp[v.Iso] = v.Name
			}
			controllers.Success(c, mp, nil)
		} else {
			var mp []map[string]interface{}
			for _, v := range l {
				nns := make(map[string]interface{})
				nns["id"] = v.Id
				nns["iso"] = v.Iso
				nns["name"] = v.Name
				mp = append(mp, nns)
			}
			controllers.Success(c, mp, nil)
		}
	}
}

//获取国家手机区号
func CountryPhoneCode(c *gin.Context) {
	filter := strings.Trim(c.Query("q"), " ")
	kv := c.Query("kv")
	iso := strings.Trim(c.Param("iso"), "/")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	langobj, _ := c.Get("_lang")
	lang := langobj.(string)

	if iso != "" {
		rs, err := models.GetCountryByIso(iso)
		if err == nil {
			if kv != "" {
				controllers.Success(c, map[string]string{rs.Iso: rs.Phonecode}, &controllers.Msg{Str: "Success"})
				return
			} else {
				controllers.Success(c, map[string]string{"iso": rs.Iso, "code": rs.Phonecode}, &controllers.Msg{Str: "Success"})
				return
			}
		}
	} else {
		rs, err := models.CountryPhoneCode(lang, kv, filter, page, limit)
		if err == nil {
			controllers.Success(c, rs, &controllers.Msg{Str: "Success"})
			return
		}
	}
	controllers.Error(c, nil, &controllers.Msg{Str: "No results found"})
	return
}

//所有数据集合
func Totals(c *gin.Context) {
	langobj, _ := c.Get("_lang")
	lang := langobj.(string)

	ddt := make(map[string]interface{})
	ddt["countrycode"], _ = models.CountryPhoneCode(lang, "", "", 0, -1)
	ddt["languages"], _ = models.GetAllLanguages(false)
	ddt["temperaments"], _ = models.GetAllTemperaments(lang, "", "", 0)
	ddt["constellations"], _ = models.GetAllConstellations(lang, "")

	controllers.Success(c, ddt, &controllers.Msg{Str: "Success"})
}

//强行停止服务器
func Panics(c *gin.Context) {
	config.Och <- false
}
