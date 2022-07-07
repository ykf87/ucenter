package models

import (
	"strconv"
	"time"
	"ucenter/app/funcs"

	"github.com/gin-gonic/gin"
)

type UserDeviceModel struct {
	Id          int64  `json:"id"`
	Uid         int64  `json:"uid"`
	Deviceid    string `json:"deviceid"`
	Platform    int    `json:"platform"`
	Version     string `json:"version"`
	Brand       string `json:"brand"`
	Devicemodel string `json:"devicemodel"`
	Ip          int64  `json:"ip"`
	Addtime     int64  `json:"addtime"`
}

//记录一次使用环境变动
func AddUserEnvironmentChange(c *gin.Context, user *UserModel) {
	if c.GetHeader("deviceid") == "" {
		return
	}
	pltid, _ := strconv.Atoi(c.GetHeader("platform"))
	if pltid < 1 || pltid > 4 {
		pltid = 3
	}
	m := new(UserDeviceModel)
	m.Uid = user.Id
	m.Deviceid = c.GetHeader("deviceid")
	m.Platform = pltid
	m.Version = c.GetHeader("version")
	m.Brand = c.GetHeader("brand")
	m.Devicemodel = c.GetHeader("devicemodel")
	m.Ip = funcs.InetAtoN(c.ClientIP())
	m.Addtime = time.Now().Unix()
	DB.Table("user_devices").Create(m)
}
