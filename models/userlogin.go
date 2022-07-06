package models

import (
	"time"
	"ucenter/app/config"
)

type UserLoginModel struct {
	Uid     int64 `json:"uid"`
	Addtime int64 `json:"addtime"`
}

//增加一条登录数据
func AddUserLoginRow(user *UserModel) {
	r := new(UserLoginModel)
	r.Uid = user.Id
	r.Addtime = time.Now().Unix()
	DB.Table("user_logins").Create(r)
}

//获取活跃用户列表
func GetPositiveUserList(page, limit int, notin []int64) ([]*UserModel, int64) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = config.Config.Limit
	} else if limit > 100 {
		limit = config.Config.Limit
	}
	var list []*UserModel
	var total int64

	dbs := DB.Select("a.*").Table("users as a").Joins("left join user_logins as b on a.id = b.uid")
	if len(notin) > 0 {
		dbs = dbs.Where("a.id not in ?", notin)
	}

	// if page == 1 && len(list) > 0 {
	// 	DB.Table("user_logins").Count(&total)
	// }

	dbs = dbs.Limit(limit).Offset((page - 1) * limit).Order("b.addtime DESC")
	dbs.Find(&list)
	return list, total
}
