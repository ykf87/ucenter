package models

type Version struct {
	Id          int64  `json:"id"`
	VersionStr  string `json:"version_str"`
	VersionCode int    `json:"version_code"`
	Remark      string `json:"remark"`
	Url         string `json:"url"`
	Uptime      int64  `json:"-"`
	Addtime     int64  `json:"-"`
	Must        int    `json:"-"`
	Canpay      int    `json:"-"`
	Platform    int    `json:"-"`
}

func GetVersions(platform string) []*Version {
	var lists []*Version
	DB.Model(&Version{}).Where("platform = ? or platform = 0", platform).Order("id Desc").Find(&lists)
	return lists
}
