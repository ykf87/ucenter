package models

type Version struct {
	Id          int64  `json:"id"`
	VersionStr  string `json:"version_str"`
	VersionCode int    `json:"version_code"`
	Reamrk      string `json:"remark"`
	Uptime      int64  `json:"-"`
	Addtime     int64  `json:"-"`
	Must        int    `json:"must"`
	Canpay      int    `json:"canpay"`
	Platform    int    `json:"-"`
}

func GetVersions(platform string) []*Version {
	var lists []*Version
	DB.Model(&Version{}).Where("platform = ?", platform).Order("id Desc").Find(&lists)
	return lists
}
