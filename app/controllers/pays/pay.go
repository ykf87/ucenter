package pays

import (
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
	orderid, reurl, err := p.Pay("USD", po.Price)
	if err != nil {
		controllers.ErrorNoData(c, "Illegal payments")
		logs.Logger.Error(err, " - 调用支付!")
		return
	}

	od := models.InitNewOrder()
	od.Amount = po.Price
	od.Bi = po.Bi
	od.Orderid = orderid
	od.Pid = po.Id
	od.Uid = user.Id
	res := models.DB.Create(od)
	if res.Error != nil {
		logs.Logger.Error(res.Error)
		controllers.ErrorNoData(c, "Illegal payments")
		return
	}

	controllers.SuccessStr(c, map[string]interface{}{"url": reurl, "orderid": orderid, "id": od.Id}, "")
}

func CheckOrder(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	rs, _ = c.Get("_lang")
	lang := rs.(string)

	id := c.PostForm("id")
	orderid := c.PostForm("orderid")

	od := new(models.Order)
	if id != "" {
		models.DB.Where("id = ?", id).First(od)
	} else if orderid != "" {
		models.DB.Where("orderid = ?", orderid).First(od)
	} else {
		controllers.ErrorNoData(c, "Missing queries")
		return
	}
	if od.Id < 1 || od.Uid != user.Id {
		controllers.ErrorNoData(c, "Order does not exist")
		return
	}
	if od.Status == 1 {
		controllers.SuccessStr(c, od, "Success")
		return
	}

	p, err := payment.Get(lang, "")
	if err != nil {
		controllers.ErrorNoData(c, "Illegal payments")
		logs.Logger.Error(err, " - 调用支付.")
		return
	}

	p.GetOrderDetail(od.Orderid)
}
