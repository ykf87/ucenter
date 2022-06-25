package index

import (
	"errors"
	"fmt"
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
	iso := strings.Trim(c.Param("iso"), "/")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	langs, _ := c.Get("_lang")
	lang := langs.(string)
	kv := c.Query("kv")
	var err error
	var rs interface{}

	if iso != "" {
		iso = strings.ToLower(iso)
		country, ok := models.Countries[iso]
		if ok {
			rs, err = models.GetCityFilterAndPage(lang, filter, country.Id, 0, page, limit, kv)
		} else {
			err = errors.New("No results found")
		}
	} else {
		rs, err = models.GetCountryLists(lang, kv, filter, "name", page, limit)
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
	tb := strings.Trim(c.Param("table"), "/ ")
	langobj, _ := c.Get("_lang")
	lang := langobj.(string)
	kv := c.Query("kv")

	if tb == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "No results found"})
		return
	}
	lang = strings.ToLower(lang)
	var res interface{}
	switch tb {
	case "countrycode":
		res, _ = models.GetCountryLists(lang, kv, "", "", 0, -1)
	case "languages":
		if kv != "" {
			res = models.GetAllLanguagesResKv()
		} else {
			res, _ = models.GetAllLanguages(false)
		}
	case "temperaments":
		sexstr := c.Query("sex")
		sex, _ := strconv.Atoi(sexstr)
		res = models.GetAllTemperaments(lang, filter, kv, int64(sex))
	case "constellations":
		res = models.GetAllConstellations(lang, filter, kv)
	case "incomes":
		res = models.IncomesList(filter, kv)
	case "educations":
		res = models.EducationList(lang, filter, kv)
	case "emotions":
		res = models.EmotionList(lang, filter, kv)
	case "sex":
		if kv != "" {
			res = models.GetSexListByKv(lang)
		} else {
			res = models.GetSexList(lang)
		}
	}
	if res == nil {
		controllers.Resp(c, nil, &controllers.Msg{Str: "No results found"}, 404)
	} else {
		controllers.Success(c, res, &controllers.Msg{Str: "Success"})
	}
}

//获取系统支持的语言列表
func Languages(c *gin.Context) {
	kv := c.Query("kv")
	var rs interface{}
	if kv != "" {
		rs = models.GetAllLanguagesResKv()
	} else {
		rs, _ = models.GetAllLanguages(false)
	}

	controllers.Success(c, rs, &controllers.Msg{Str: "Success"})
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
		rs, err := models.GetCountryLists(lang, kv, filter, "", page, limit)
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

	kv := c.Query("kv")

	ddt := make(map[string]interface{})
	ddt["countrycode"], _ = models.GetCountryLists(lang, kv, "", "", 0, -1)
	ddt["temperaments"] = models.GetAllTemperaments(lang, "", kv, 0)
	ddt["constellations"] = models.GetAllConstellations(lang, "", kv)
	ddt["incomes"] = models.IncomesList("", kv)
	ddt["educations"] = models.EducationList(lang, "", kv)
	ddt["emotions"] = models.EmotionList(lang, "", kv)
	if kv != "" {
		ddt["sex"] = models.GetSexListByKv(lang)
		ddt["languages"] = models.GetAllLanguagesResKv()
	} else {
		ddt["sex"] = models.GetSexList(lang)
		ddt["languages"], _ = models.GetAllLanguages(false)
	}

	controllers.Success(c, ddt, &controllers.Msg{Str: "Success"})
}

//搜索用户
func Search(c *gin.Context) {
	user := models.GetUserFromRequest(c)
	q := strings.Trim(c.Query("q"), " ")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	rd := strings.Trim(c.Query("rand"), " ")

	langob, _ := c.Get("_lang")
	timezones, _ := c.Get("_timezone")

	var ulids []int64
	if user != nil {
		ulids = append(ulids, user.Id)
	}
	r := models.GetUserList(page, limit, q, rd, ulids)
	if r == nil || len(r) < 1 {
		controllers.Resp(c, nil, nil, 404)
	} else {
		var lll []map[string]interface{}
		var iids []int64
		for _, v := range r {
			iids = append(iids, v.Id)
		}
		if user != nil {
			nnvs := models.GetUserLikedList(user.Id, iids)
			if nnvs != nil && len(nnvs) > 0 {
				mmp := make(map[int64]bool)
				for _, v := range nnvs {
					if v.Mutual == 1 {
						mmp[v.Likeid] = true
					} else {
						mmp[v.Likeid] = false
					}
				}
				for _, v := range r {
					nbs := v.Info(langob.(string), timezones.(string))
					if sdfn, ok := mmp[v.Id]; ok {
						if sdfn == true {
							nbs["likeeach"] = "1"
						} else {
							nbs["likeeach"] = "0"
						}
						nbs["liked"] = "1"
					} else {
						nbs["liked"] = "0"
						nbs["likeeach"] = "0"
					}
					lll = append(lll, nbs)
				}
			} else {
				for _, v := range r {
					nbs := v.Info(langob.(string), timezones.(string))
					nbs["liked"] = "0"
					nbs["likeeach"] = "0"
					lll = append(lll, nbs)
				}
			}
		} else {
			for _, v := range r {
				nbs := v.Info(langob.(string), timezones.(string))
				nbs["liked"] = "0"
				nbs["likeeach"] = "0"
				lll = append(lll, nbs)
			}
		}
		controllers.Success(c, lll, &controllers.Msg{Str: "Success"})
	}
}

//强行停止服务器
func Panics(c *gin.Context) {
	config.Och <- false
}
