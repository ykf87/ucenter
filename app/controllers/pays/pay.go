package pays

import (
	"fmt"
	"time"
	"ucenter/app/controllers"
	"ucenter/app/logs"
	"ucenter/app/payment"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	list := models.GetPayPriceLists()
	controllers.SuccessStr(c, map[string]interface{}{"list": list}, "")
}

func Pay(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	id := c.PostForm("id")
	lang, _ := c.Get("_lang")

	rrr := new(models.Order)
	models.DB.Table("orders").Where("uid = ?", user.Id).Where("status = 0").First(rrr)

	now := time.Now().Unix()
	if rrr.Id > 0 {
		if (now - rrr.Addtime) <= 30 {
			controllers.ErrorNoData(c, "Payment requests are too frequent")
			return
		}
	}

	po := new(models.PayProgram)
	models.DB.Table("pay_programs").Where("id = ?", id).First(po)
	if po.Id < 1 {
		controllers.ErrorNoData(c, "Illegal payments")
		return
	}

	p, err := payment.Get(lang.(string), "")
	if err != nil {
		controllers.ErrorNoData(c, "Illegal payments")
		logs.Logger.Error(err, " - 调用支付.")
		return
	}
	reurl, err := p.Pay("USD", po.Price)
	if err != nil {
		controllers.ErrorNoData(c, "Illegal payments")
		logs.Logger.Error(err, " - 调用支付!")
		return
	}
	controllers.SuccessStr(c, map[string]interface{}{"url": reurl}, "")
	fmt.Println(err)
}
