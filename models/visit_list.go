package models

type VisitList struct {
	Id      int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Uid     int64 `json:"uid" gorm:"index;not null"`
	VisitId int64 `json:"visit_id" gorm:"index;not null"`
	Addtime int64 `json:"addtime"`
}
