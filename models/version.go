package models

import (
	"time"
)

type Version struct {
	Id          int64  `json:"id"`
	VersionStr  string `json:"version_str"`
	VersionCode int    `json:"version_code"`
	Remark      string `json:"remark"`
	Url         string `json:"url"`
	Uptime      int64  `json:"-"`
	Addtime     int64  `json:"-"`
	Must        int    `json:"must"`
	Canpay      int    `json:"canpay"`
	Platform    int    `json:"-"`
	Audit       int    `json:"audit"`
}

func GetVersions(platform string, version_code int) []*Version {
	var lists []*Version
	now := time.Now().Unix()
	DB.Model(&Version{}).Where("platform = ?", platform).Where("version_code >= ?", version_code).Where("uptime <= ?", now).Order("id Desc").Find(&lists)
	return lists
}
