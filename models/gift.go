package models

type Gift struct {
	Id     int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"not null"`
	Icon   string `json:"icon" gorm:"not null"`
	Bi     int64  `json:"bi" gorm:"defult 1"`
	Sort   int    `json:"-" gorm:"default 0"`
	Status int    `json:"-" gorm:"default 1"`
}

//获取礼物列表
func GetGiftList(page, limit int) ([]*Gift, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	var l []*Gift
	rs := DB.Where("status = 1").Order("sort DESC").Offset((page - 1) * limit).Limit(limit).Find(&l)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return l, nil
}

//获取单个礼物配置
func GetGiftRow(id int) (*Gift, error) {
	r := new(Gift)
	rs := DB.Where("id = ?", id).First(r)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return r, nil
}
