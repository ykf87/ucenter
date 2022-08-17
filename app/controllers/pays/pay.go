package pays

import (
	"strconv"
	"time"
	"ucenter/app/controllers"
	"ucenter/app/logs"

	// "ucenter/app/payment"
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

	p := models.Paypal
	p.SetLang(lang.(string))
	// p, err := payment.Get(lang.(string), "")
	// if err != nil {
	// 	controllers.ErrorNoData(c, "Illegal payments")
	// 	logs.Logger.Error(err, " - 调用支付.")
	// 	return
	// }
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
	// go od.ListenOrder()

	controllers.SuccessStr(c, map[string]interface{}{"url": reurl, "orderid": orderid, "id": od.Id}, "")
}

func CheckOrder(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	id := c.PostForm("id")
	orderid := c.PostForm("orderid")

	rightNow := c.PostForm("rightnow")

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
	// if od.Status == 1 {
	// 	controllers.SuccessStr(c, od, "Success")
	// 	return
	// }

	// p, err := payment.Get(lang, "")
	// if err != nil {
	// 	controllers.ErrorNoData(c, "Illegal payments")
	// 	logs.Logger.Error(err, " - 调用支付.")
	// 	return
	// }
	if rightNow != "" {
		od.FollowerStatus()
	} else if od.Status == 0 && time.Now().Unix()-od.Addtime >= 420 { //订单未支付,七分钟后开启用户查询
		od.FollowerStatus()
		// rs, _ = c.Get("_lang")
		// lang := rs.(string)
		// p := models.Paypal
		// p.SetLang(lang)

		// odt, err := p.GetOrderDetail(od.Orderid)
		// if err != nil {
		// 	controllers.ErrorNoData(c, "Illegal payments")
		// 	logs.Logger.Error(err, " - 调用支付.")
		// 	return
		// } else {
		// 	if stts, ok := models.OrderStatusMap[odt.Status]; ok {
		// 		if od.Status != stts {
		// 			od.ChangeOrderStatus(stts)
		// 		}
		// 	}
		// }
	}

	controllers.SuccessStr(c, od, "")
}

//扣费计费
func Billing(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	interid := c.PostForm("interid")
	isend := c.PostForm("ended")
	tp := c.PostForm("type") //1-语音, 2-视频
	tttp, _ := strconv.Atoi(tp)

	if (tttp != 1 && tttp != 2) || interid == "" {
		controllers.ErrorNoData(c, "Error")
		return
	}
	var onecode int64
	onecode = 10
	if tttp == 1 {
		onecode = 5
	}

	cos := new(models.Consume)
	models.DB.Where("uid = ?", user.Id).Where("connect_id = ?", interid).First(cos)
	if cos.Id > 0 {
		now := time.Now().Unix()
		usetime := now - cos.Start
		cost := ((usetime / 60) + 1) * cos.Seccost
		data := map[string]interface{}{
			"uptime":  time.Now().Unix(),
			"cost":    cost,
			"usetime": usetime,
		}
		if isend == "1" {
			data["status"] = 1
		}
		models.DB.Model(&models.Consume{}).Where("id = ?", cos.Id).Updates(data)
	} else {
		cos = models.OpenConsume(user.Id, interid, tttp, onecode)
	}

	balance := user.GetUserBalance()
	msg := ""
	if balance < onecode {
		msg = "Insufficient balance"
	}
	cos.Balance = balance
	controllers.SuccessStr(c, cos, msg)

	// if cos != nil {
	// 	balance := user.GetUserBalance()
	// 	if balance < onecode {
	// 		controllers.ErrorNoData(c, "Insufficient balance")
	// 		return
	// 	}
	// 	controllers.SuccessStr(c, rrrr, "")
	// } else {
	// 	controllers.ErrorNoData(c, "Failed to create")
	// }

}

//查询余额
func Balance(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	controllers.SuccessStr(c, map[string]interface{}{"balance": user.GetUserBalance()}, "")
}
