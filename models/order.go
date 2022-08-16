package models

import (
	"time"
)

type Order struct {
	Id      int64   `json:"id" gorm:"primaryKey"`
	Uid     int64   `json:"uid" gorm:"not null;index"`
	Pid     int64   `json:"pid" gorm:"not null;index"`
	Orderid string  `json:"orderid" gorm:"not null"`
	Amount  float64 `json:"amount" gorm:"not null"`
	Bi      int     `json:"bi" gorm:"default:0"`
	Addtime int64   `json:"addtime" gorm:"not null"`
	Status  int     `json:"status" gorm:"default:0;not null"`
	PayWay  int     `json:"pay_way" gorm:"default:1"`
	Paytime int64   `json:"paytime"`
}

func InitNewOrder() *Order {
	od := new(Order)
	od.Addtime = time.Now().Unix()
	return od
}
