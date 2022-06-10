package models

type UserModel struct {
	Id       int64   `json:"id"`
	Pid      int64   `json:"pid"`
	Account  string  `json:"account"`
	Mail     string  `json:"mail"`
	Phone    string  `json:"phone"`
	Pwd      string  `json:"pwd"`
	Nickname string  `json:"nickname"`
	Avatar   string  `json:"avatar"`
	Addtime  int64   `json:"addtime"`
	Status   int     `json:"status"`
	Sex      int     `json:"sex"`
	Height   int     `json:"height"`
	Weight   float32 `json:"weight"`
	Birth    int64   `json:"birth"`
	Age      int     `json:"age"`
	Job      string  `json:"job"`
	Income   string  `json:"income"`
	Emotion  int     `json:"emotion"`
	Star     int     `json:"star"`
	Ip       int64   `json:"ip"`
}
