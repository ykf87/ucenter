package index

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var err error
	var rs []map[string]interface{}

	if pc != "" {
		if strings.Contains(pc, "/") {
			ids := strings.Split(pc, "/")
			countryId, _ := strconv.Atoi(ids[0])
			provId, _ := strconv.Atoi(ids[1])
			rs, err = models.GetCityFilterAndPage(lang, filter, countryId, provId, page, limit)
		} else {
			countryId, _ := strconv.Atoi(pc)
			rs, err = models.GetProvinceByFilterAndPage(lang, filter, countryId, page, limit)
		}
	} else {
		rs, err = models.GetCountryByFilterAndPage(lang, filter, page, limit)
	}

	if err != nil || len(rs) < 1 {
		controllers.Resp(c, nil, &controllers.Msg{Str: "No results found"}, 404)
	} else {
		controllers.Success(c, rs, nil)
	}
}
