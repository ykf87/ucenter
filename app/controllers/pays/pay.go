package pays

import (
	"fmt"
	"strconv"
	"time"
	"ucenter/app/controllers"
	"ucenter/app/logs"
	"ucenter/app/payment/apple"

	// "ucenter/app/payment"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	list := models.GetPayPriceLists()
	controllers.SuccessStr(c, map[string]interface{}{"list": list}, "")
}

func Pay(c *gin.Context) {
	// body, _ := ioutil.ReadAll(c.Request.Body)
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	id := c.PostForm("id")
	lang, _ := c.Get("_lang")
	if lang == "zh-cn" {
		lang = "zh"
	}
	// fmt.Println(string(body), "\r\n---", id)

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
	od.Mail = user.Mail
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
	od.CaptureOrder(orderid)
	if rightNow != "" {
		if od.FollowerStatus() == nil {
			go user.ChangeRecharge(0.0)
		}
		// go user.ChangeRecharge(0.0)
	} else if od.Status == 0 && time.Now().Unix()-od.Addtime >= 420 { //订单未支付,七分钟后开启用户查询
		if od.FollowerStatus() == nil {
			go user.ChangeRecharge(0.0)
		}
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

	interid := c.PostForm("id") //对方用户id
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
	now := time.Now().Unix()
	lt := now - 120
	models.DB.Where("uid = ?", user.Id).Where("connect_id = ?", interid).Where("status = 0").Where("uptime >= ?", lt).First(cos)
	if cos.Id > 0 {
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
	if isend == "1" {
		bbs, _ := strconv.ParseFloat(fmt.Sprintf("%d", balance), 64)
		go user.ChangeUsed(0.0, bbs)
	}
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

//苹果端
func ApplePay(c *gin.Context) {
	platform := c.GetHeader("platform")
	if platform != "2" {
		controllers.ErrorNotFound(c)
		return
	}

	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	programs := c.PostForm("programs") //充值方案
	prices := c.PostForm("price")      //充值金额
	price, _ := strconv.ParseFloat(prices, 64)
	appleOrderId := c.PostForm("orderid") //苹果充值id
	if appleOrderId == "" || programs == "" || price <= 0 {
		controllers.ErrorNotFound(c)
		return
	}

	odid, err := apple.VeryOrder(appleOrderId)
	if err != nil && odid == "" {
		controllers.ErrorNoData(c, "Error")
		return
	}

	po := new(models.PayProgram)
	models.DB.Table("pay_programs").Where("id = ?", programs).First(po)
	if po.Id < 1 {
		controllers.ErrorNotFound(c)
		return
	}

	ood := new(models.Order)
	models.DB.Model(&models.Order{}).Where("orderid = ?", odid).First(ood)
	if ood.Id > 0 {
		controllers.ErrorNoData(c, "The order already exists")
		return
	}

	order := new(models.Order)
	now := time.Now().Unix()
	order.Addtime = now
	order.Amount = po.Price
	order.Bi = po.Bi
	order.Orderid = odid
	order.Paytime = now
	order.PayWay = 2
	order.Pid = po.Id
	order.Uid = user.Id
	order.Mail = user.Mail
	rbs := models.DB.Create(order)
	if rbs.Error != nil {
		controllers.ErrorNoData(c, "Failed to recharge")
		return
	}
	if order.ChangeOrderStatus(1) == nil {
		go user.ChangeRecharge(price)
	}
	controllers.SuccessStr(c, map[string]interface{}{"balance": user.GetUserBalance()}, "")
}

func GooglePay(c *gin.Context) {

}
