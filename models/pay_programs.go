package models

type PayProgram struct {
	Id      int64   `json:"id" gorm:"primaryKey"`
	Price   float64 `json:"price" gorm:"not null"`
	Bi      int     `json:"bi" gorm:"not null"`
	Remark  string  `json:"remark"`
	Status  int     `json:"status" gorm:"default:1"`
	Used    int     `json:"used" gorm:"default:0"`
	Pin     int     `json:"pin" gorm:"default:0"`
	PinTime int64   `json:"pin_time" gorm:"default:0"`
}

//func init() {
//	DB.AutoMigrate(&PayProgram{})
//}
func GetPayPriceLists() []*PayProgram {
	var ll []*PayProgram
	DB.Select("id", "price", "bi", "remark", "used").Where("status = 1").Order("price asc").Find(&ll)
	return ll
}
