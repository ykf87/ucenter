package models

import (
	"time"
)

type UserLikeModel struct {
	Id      int64 `json:"id"`
	Likeid  int64 `json:"likeid"`
	Mutual  int64 `json:"mutual"`
	Addtime int64 `json:"addtime"`
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
func GetUserLikedList(uid int64, inids []int64) []*UserLikeModel {
	var lks []*UserLikeModel
	rs := DB.Table("user_likes").Where("id = ?", uid)
	if inids != nil && len(inids) > 0 {
		rs = rs.Where("likeid in ?", inids)
	}
	rrs := rs.Find(&lks)
	if rrs.Error != nil {
		return nil
	}
	return lks
}
