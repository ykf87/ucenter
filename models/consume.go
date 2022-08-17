package models

import (
	"time"
	"ucenter/app/logs"
)

type Consume struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	Uid       int64  `json:"uid" gorm:"index;not null"`
	ConnectId string `json:"connect_id" gorm:"index;not null"`
	Voice     int    `json:"voice" gorm:"not null;index"`
	Start     int64  `json:"start" gorm:"not null"`
	Uptime    int64  `json:"uptime"`
	End       int64  `json:"end"`
	Usetime   int64  `json:"usetime"`
	Seccost   int64  `json:"seccost"`
	Cost      int64  `json:"cost" gorm:"index"`
	Status    int    `json:"status" gorm:"index"`
	Balance   int64  `json:"balance" gorm:"-"`
}

func IUsed(uid int64) int64 {
	var total int64
	DB.Model(&Consume{}).Select("sum(cost) as total").Where("uid = ?", uid).First(&total)
	return total
}

func OpenConsume(uid int64, cid string, tp int, cost int64) *Consume {
	r := new(Consume)
	r.Uid = uid
	r.ConnectId = cid
	r.Voice = tp
	r.Start = time.Now().Unix()
	r.Seccost = cost
	r.Uptime = r.Start + 1
	r.Cost = cost

	rs := DB.Create(r)
	if rs.Error == nil {
		return r
	}
	logs.Logger.Error(rs.Error)
	return nil
}
