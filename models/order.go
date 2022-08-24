package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	_ "ucenter/app/config"
	"ucenter/app/logs"
	"ucenter/app/payment"
)

type Order struct {
	Id      int64   `json:"id" gorm:"primaryKey"`
	Uid     int64   `json:"-" gorm:"not null;index"`
	Mail    string  `json:"mail" gorm:"not null;index"`
	Pid     int64   `json:"-" gorm:"not null;index"`
	Orderid string  `json:"orderid" gorm:"not null"`
	Amount  float64 `json:"amount" gorm:"not null"`
	Bi      int     `json:"bi" gorm:"default:0"`
	Addtime int64   `json:"-" gorm:"not null"`
	Status  int     `json:"status" gorm:"default:0;not null"`
	PayWay  int     `json:"-" gorm:"default:1"`
	Paytime int64   `json:"paytime"`
}

var OrderStatusMap map[string]int

var Paypal payment.Payment

func init() {
	p, err := payment.Get("en", "")
	if err != nil {
		logs.Logger.Error(err)
		panic(err)
	} else {
		Paypal = p
	}

	OrderStatusMap = map[string]int{
		"CREATED":               0,
		"SAVED":                 2,
		"APPROVED":              -2,
		"VOIDED":                -1,
		"COMPLETED":             1,
		"PAYER_ACTION_REQUIRED": 3,
	}
}

func InitNewOrder() *Order {
	od := new(Order)
	od.Addtime = time.Now().Unix()
	return od
}

func (this *Order) ListenOrder() {
	time.Sleep(time.Second * 30) //30秒后开始查询订单状态
	now := time.Now().Unix()
	for {
		if this.FollowerStatus() == nil {
			break
		}
		// o, err := Paypal.GetOrderDetail(this.Orderid)
		// if err == nil {
		// 	if stts, ok := OrderStatusMap[o.Status]; ok {
		// 		if stts != this.Status {
		// 			ut, _ := strconv.Atoi(o.UpdateTime)
		// 			this.Paytime = int64(ut)
		// 			this.ChangeOrderStatus(stts)
		// 			break
		// 		}
		// 	}
		// }
		if time.Now().Unix()-now >= 300 {
			break
		}
		time.Sleep(time.Second * 30)
	}
}

//同步订单状态
func (this *Order) FollowerStatus() error {
	o, err := Paypal.GetOrderDetail(this.Orderid)
	if err == nil {
		ods, _ := json.Marshal(o)
		if stts, ok := OrderStatusMap[o.Status]; ok {
			if stts != this.Status {
				ut, _ := strconv.Atoi(o.UpdateTime)
				this.Paytime = int64(ut)
				logs.Logger.Println(string(ods))
				return this.ChangeOrderStatus(stts)
			}
		} else {
			logs.Logger.Error(fmt.Sprintf("订单状态不存在!%s - %s", o.Status, string(ods)))
		}
	} else {
		logs.Logger.Error(err)
		return err
	}
	return errors.New("no update")
}

//统一的修改订单状态,禁止在其他修改
func (this *Order) ChangeOrderStatus(status int) error {
	if this.Status == status {
		return errors.New("Order status not change")
	}

	paytime := this.Paytime
	if paytime == 0 {
		paytime = time.Now().Unix()
	}
	data := map[string]interface{}{
		"status":  status,
		"paytime": this.Paytime,
	}
	if DB.Table("orders").Where("id = ?", this.Id).Updates(data).Error == nil {
		this.Status = status
		this.Paytime = paytime
	}
	return nil
}

//获取用户成功充值金额
func Recharged(uid int64) int64 {
	var total int64
	DB.Model(&Order{}).Select("sum(bi) as total").Where("uid = ?", uid).Where("status = 1").First(&total)
	return total
}
