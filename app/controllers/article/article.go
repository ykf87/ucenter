package article

import (
	"strconv"
	"ucenter/app/controllers"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	langob, _ := c.Get("_lang")
	lang := langob.(string)
	timezones, _ := c.Get("_timezone")
	timezone := timezones.(string)

	key := c.Param("key")
	if key != "" {
		row := models.GetArticleRow(0, key, lang)
		if row != nil && row.Id > 0 {
			dt := row.Fmt(lang, timezone)
			controllers.SuccessStr(c, dt, "Success")
			return
		}
	} else if idstr := c.Query("id"); idstr != "" {
		id32, _ := strconv.Atoi(idstr)
		if id32 > 0 {
			row := models.GetArticleRow(int64(id32), "", lang)
			if row != nil && row.Id > 0 {
				dt := row.Fmt(lang, timezone)
				controllers.SuccessStr(c, dt, "Success")
				return
			}
		}
	}
	controllers.ErrorNotFound(c)
}
