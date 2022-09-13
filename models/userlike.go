package models

import (
	"time"
	"ucenter/app/config"
)

type UserLikeModel struct {
	Id      int64                  `json:"id"`
	Likeid  int64                  `json:"likeid"`
	Mutual  int64                  `json:"mutual"`
	Addtime int64                  `json:"addtime"`
	Info    map[string]interface{} `json:"info" gorm:"-"`
}

//检查是否已经喜欢了对方
func CheckUserIsLiked(uid, likeid int64) *UserLikeModel {
	lks := new(UserLikeModel)
	rs := DB.Table("user_likes").Where("id = ? and likeid = ?", uid, likeid).Find(lks)
	if rs.Error != nil || lks == nil || lks.Id < 1 {
		return nil
	}
	return lks
}

//喜欢一个人
func UserLikeAdd(uid, likeid int64) bool {
	liked := CheckUserIsLiked(uid, likeid)
	if liked == nil {
		ddt := new(UserLikeModel)
		ddt.Id = uid
		ddt.Likeid = likeid
		ddt.Addtime = time.Now().Unix()

		faleback := CheckUserIsLiked(likeid, uid)
		if faleback != nil {
			ddt.Mutual = 1
		}
		rs := DB.Table("user_likes").Create(ddt)
		if rs.Error == nil {
			if faleback != nil {
				DB.Table("user_likes").Where("id = ? and likeid = ?", likeid, uid).Update("mutual", 1)
			}
		} else {
			return false
		}
	}
	return true
}

//获取用户id喜欢的用户
func GetUserLikedList(uid int64, inids []int64, page, limit int) []*UserLikeModel {
	var lks []*UserLikeModel
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = config.Config.Limit
	}
	rs := DB.Table("user_likes").Where("id = ?", uid).Offset((page - 1) * limit).Limit(limit)
	if inids != nil && len(inids) > 0 {
		rs = rs.Where("likeid in ?", inids)
	}
	rrs := rs.Find(&lks)
	if rrs.Error != nil {
		return nil
	}
	return lks
}
