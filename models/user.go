package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"ucenter/app/safety/passwordhash"
	"ucenter/app/safety/rsautil"

	"github.com/tidwall/gjson"
)

type UserModel struct {
	Id        int64   `json:"id"`
	Pid       int64   `json:"pid"`
	Account   string  `json:"account"`
	Mail      string  `json:"mail"`
	Phone     string  `json:"phone"`
	Mailvery  int     `json:"mailvery"`
	Phonevery int     `json:"phonevery"`
	Pwd       string  `json:"pwd"`
	Nickname  string  `json:"nickname"`
	Avatar    string  `json:"avatar"`
	Addtime   int64   `json:"addtime"`
	Status    int     `json:"status"`
	Sex       int     `json:"sex"`
	Height    int     `json:"height"`
	Weight    float32 `json:"weight"`
	Birth     int64   `json:"birth"`
	Age       int     `json:"age"`
	Job       string  `json:"job"`
	Income    string  `json:"income"`
	Emotion   int     `json:"emotion"`
	Star      int     `json:"star"`
	Ip        int64   `json:"ip"`
	Edinfo    Editers
}

//账号修改器
type Editers int64

func (this Editers) Account(args ...interface{}) {
	fmt.Println(args)
	fmt.Println(this)
}

//创建新用户
func MakeUser(account, email, phone, pwd, ip string) (user *UserModel, err error) {
	hadUser := new(UserModel)
	insertData := make(map[string]interface{})
	if account != "" {
		DB.Table("users").Where("account = ?", account).First(hadUser)
		insertData["account"] = account
	} else if email != "" {
		DB.Table("users").Where("mail = ?", email).First(hadUser)
		insertData["mail"] = email
	} else if phone != "" {
		DB.Table("users").Where("phone = ?", phone).First(hadUser)
		insertData["phone"] = phone
	}
	if hadUser.Id > 0 {
		err = errors.New("User already exists, please login")
		return
	}
	if pwd != "" {
		insertData["pwd"], err = passwordhash.PasswordHash(pwd)
		if err != nil {
			return
		}
	}
	insertData["status"] = 1
	insertData["ip"] = InetAtoN(ip)
	insertData["addtime"] = time.Now().Unix()
	rs := DB.Table("users").Create(insertData)
	if rs.Error != nil {
		err = rs.Error
	}
	user = new(UserModel)
	user.Edinfo = Editers(user.Id)
	DB.Table("users").Where(insertData).First(user)
	return
}

//查找单个用户
func GetUser(id int64, account, email, phone string) *UserModel {
	user := new(UserModel)
	if id > 0 {
		DB.Table("users").Where("id = ?", id).First(user)
	} else if account != "" {
		DB.Table("users").Where("account = ?", account).First(user)
	} else if email != "" {
		DB.Table("users").Where("mail = ? and mailvery = 1", email).First(user)
	} else if phone != "" {
		DB.Table("users").Where("phone = ? and phonevery = 1", phone).First(user)
	}
	if user.Id < 1 {
		user = nil
	}

	user.Edinfo = Editers(user.Id)
	return user
}

//生成用户token
func (this *UserModel) Token() string {
	if this.Id == 0 {
		log.Println("UserModel Token - 用户实例id为0")
		return ""
	}
	token, err := rsautil.RsaEncrypt(fmt.Sprintf("%d", this.Id))
	if err != nil {
		log.Println("UserModel Token - ", err)
		return ""
	}
	return token
}

//通过 token 生成 user model
func UnToken(token string) *UserModel {
	idstr, err := rsautil.RsaDecrypt(token)
	if err != nil {
		return nil
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil
	}
	return GetUser(int64(id), "", "", "")
}

//返回用户信息
func (this *UserModel) Info() map[string]interface{} {
	if this.Id < 1 {
		return nil
	}
	data := make(map[string]interface{})
	b, _ := json.Marshal(this)
	for k, v := range gjson.ParseBytes(b).Map() {
		if k == "pwd" || k == "status" {
			continue
		}
		data[k] = v.String()
	}
	return data
}
