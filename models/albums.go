package models

import (
	"time"
	"ucenter/app/config"
	"ucenter/app/uploadfile/images"

	carbon "github.com/golang-module/carbon/v2"
)

const (
	ALBUMSAVEPATH = "static/user/album"
)

type AlbumModel struct {
	Id         int64  `json:"id"`
	Uid        int64  `json:"uid"`
	Src        string `json:"src"`
	Private    int    `json:"private"`
	Likes      int64  `json:"likes"`
	Addtime    int64  `json:"addtime"`
	Addtimefmt string `json:"addtimefmt"`
}

func GetAlbumList(uid int64, page, limit int, private bool) []*AlbumModel {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = config.Config.Limit
	}
	obj := DB.Table("user_albums").Where("uid = ?", uid)
	if private == true {
		obj = obj.Where("private = 1")
	} else {
		obj = obj.Where("private = 0")
	}
	obj = obj.Limit(limit).Offset((page - 1) * limit).Order("id DESC")
	var list []*AlbumModel
	rs := obj.Find(&list)
	if rs.Error != nil {
		return nil
	}
	return list
}

//添加相册
func AddAlbumList(uid int64, srcs []string, private int) ([]*AlbumModel, error) {
	var albs []*AlbumModel
	now := time.Now().Unix()
	for _, v := range srcs {
		if v != "" {
			row := new(AlbumModel)
			row.Uid = uid
			row.Src = v
			row.Private = private
			row.Addtime = now
			albs = append(albs, row)
		}
	}
	rs := DB.Table("user_albums").Create(&albs)
	return albs, rs.Error
}

//通过id获取列表
func GetAlbumListByIds(uid int64, ids []string) []*AlbumModel {
	var list []*AlbumModel
	DB.Table("user_albums").Where("uid = ?", uid).Where("id in ?", ids).Find(&list)
	return list
}

//通过id删除
func RemoveAlbumByIds(uid int64, ids []int64) error {
	rs := DB.Table("user_albums").Where("uid = ?", uid).Where("id in ?", ids).Delete(&AlbumModel{})
	return rs.Error
}

func (this *AlbumModel) Fmt(timezone, lang string) {
	var fmt string
	fmts, ok := config.Config.Timefmts[lang]
	if ok {
		fmt = fmts.Datetimefmt
	} else {
		fmt = config.Config.Datetimefmt
	}
	this.Src = images.FullPath(this.Src)
	this.Addtimefmt = carbon.CreateFromTimestamp(this.Addtime).SetTimezone(timezone).Carbon2Time().Format(fmt)
}
