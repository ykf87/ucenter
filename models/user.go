package models

//uid 从39350开始,邀请码后4位有效,前方2位补0
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"ucenter/app/safety/base34"
	"ucenter/app/safety/passwordhash"
	"ucenter/app/safety/rsautil"

	carbon "github.com/golang-module/carbon/v2"
	"github.com/tidwall/gjson"
)

type UserModel struct {
	Id        int64   `json:"id"`
	Pid       int64   `json:"pid"`
	Invite    string  `json:"invite"`
	Chain     string  `json:"chain"`
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
	Country   int64   `json:"country"`
	City      int64   `json:"city"`
	Singleid  int64
	Edinfo    Editers
}

//账号修改器
type Editers int64

//创建新用户
func MakeUser(account, email, phone, pwd, invite, ip string) (user *UserModel, err error) {
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
	if invite != "" {
		inviteUser := new(UserModel)
		DB.Table("users").Where("invite = ?", invite).First(inviteUser)
		if inviteUser.Id > 0 {
			insertData["pid"] = inviteUser.Id
			var addChain string
			if inviteUser.Chain != "" {
				addChain = inviteUser.Chain + ","
			}
			insertData["chain"] = addChain + fmt.Sprintf("%d", inviteUser.Id)
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
	if DB.Table("users").Where(insertData).First(user).Error == nil && user.Id > 0 {
		DB.Table("users").Where("id = ?", user.Id).Update("invite", string(base34.Base34(uint64(user.Id))))
	}

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
		DB.Table("users").Where("mail = ?", email).First(user) // and mailvery = 1
	} else if phone != "" {
		DB.Table("users").Where("phone = ?", phone).First(user) // and phonevery = 1
	}
	if user.Id < 1 {
		user = nil
	} else {
		user.Edinfo = Editers(user.Id)
	}

	return user
}

//生成用户token
func (this *UserModel) Token() string {
	if this.Id == 0 {
		log.Println("UserModel Token - 用户实例id为0")
		return ""
	}
	token, err := rsautil.RsaEncrypt(fmt.Sprintf(`{"id":%d,"time":%d,"sid":%d}`, this.Id, time.Now().Unix(), this.Singleid))
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
	ts := gjson.Get(idstr, "time").Int()
	id := gjson.Get(idstr, "id").Int()
	sid := gjson.Get(idstr, "sid").Int()
	if id < 1 || ts < 1 {
		return nil
	}
	if time.Now().Unix()-ts >= 86400*365 {
		return nil
	}
	user := GetUser(int64(id), "", "", "")
	if sid != user.Singleid {
		return nil
	}
	return user
}

//返回用户信息
func (this *UserModel) Info() map[string]interface{} {
	if this.Id < 1 {
		return nil
	}
	data := make(map[string]interface{})
	b, _ := json.Marshal(this)
	for k, v := range gjson.ParseBytes(b).Map() {
		if k == "pwd" || k == "status" || k == "singleid" || k == "chain" {
			continue
		}
		data[k] = v.String()
	}
	return data
}

//账号名称修改
func (this Editers) SetAccount(user *UserModel, args ...interface{}) error {
	newAccount := args[0].(string)
	if newAccount == "" {
		return errors.New("Please set the account name")
	}
	if user.Account == newAccount {
		return errors.New("No account changes")
	}
	u := new(UserModel)
	DB.Table("users").Where("account = ?", newAccount).First(u)
	if u.Id > 0 {
		return errors.New("Account name already exists, please use another one")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("account", newAccount)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//邮箱修改,未接入邮箱验证码
func (this Editers) SetEmail(user *UserModel, args ...interface{}) error {
	newAccount := args[0].(string)
	if newAccount == "" {
		return errors.New("Please set the Email")
	}
	if user.Mail == newAccount {
		return errors.New("No email changes")
	}
	u := new(UserModel)
	DB.Table("users").Where("mail = ?", newAccount).First(u)
	if u.Id > 0 {
		return errors.New("Email already exists, please use another one")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("mail", newAccount)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//手机号修改,未接入手机证码
func (this Editers) SetPhone(user *UserModel, args ...interface{}) error {
	newAccount := args[0].(string)
	if newAccount == "" {
		return errors.New("Please set the Phone")
	}
	if user.Mail == newAccount {
		return errors.New("No phone changes")
	}
	u := new(UserModel)
	DB.Table("users").Where("phone = ?", newAccount).First(u)
	if u.Id > 0 {
		return errors.New("Phone already exists, please use another one")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("phone", newAccount)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//修改密码
func (this Editers) SetPassword(user *UserModel, args ...interface{}) (error, map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		return errors.New("Please set the content to be modified"), nil
	}
	npwd, _ := passwordhash.PasswordHash(changeto)
	dt := &map[string]interface{}{
		"pwd":      npwd,
		"singleid": user.Singleid + 1,
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Updates(dt)
	if rs.Error != nil {
		return rs.Error, nil
	}
	return nil, map[string]interface{}{"token": user.Token()}
}

//修改昵称
func (this Editers) SetNickname(user *UserModel, args ...interface{}) error {
	changeto := args[0].(string)
	if changeto == "" {
		return errors.New("Please set the content to be modified")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("nickname", changeto)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//修改头像
func (this Editers) SetAvatar(user *UserModel, args ...interface{}) error {
	changeto := args[0].(string)
	if changeto == "" {
		return errors.New("Please set the content to be modified")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("avatar", changeto)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//修改性别
func (this Editers) SetSex(user *UserModel, args ...interface{}) error {
	changeto := args[0].(string)
	if changeto == "" {
		return errors.New("Please set the content to be modified")
	}
	if changeto != "0" && changeto != "1" && changeto != "2" {
		changeto = "0"
	}
	cgt, _ := strconv.Atoi(changeto)
	rs := DB.Table("users").Where("id = ?", user.Id).Update("sex", cgt)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//修改身高
func (this Editers) SetHeight(user *UserModel, args ...interface{}) error {
	changeto := args[0].(string)
	if changeto == "" {
		return errors.New("Please set the content to be modified")
	}
	cgt, _ := strconv.Atoi(changeto)
	if cgt < 1 || cgt >= 300 {
		return errors.New("Please set the height reasonably")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("height", cgt)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//修改体重
func (this Editers) SetWeight(user *UserModel, args ...interface{}) error {
	changeto := args[0].(string)
	if changeto == "" {
		return errors.New("Please set the content to be modified")
	}
	cgt, _ := strconv.ParseFloat(changeto, 64)
	if cgt < 1.0 || cgt >= 500000.0 {
		return errors.New("Please set your weight wisely")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("weight", cgt)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

//修改生日
func (this Editers) SetBirth(user *UserModel, args ...interface{}) error {
	changeto := args[0].(string)
	if changeto == "" {
		return errors.New("Please set the content to be modified")
	}
	cgt := carbon.Parse(changeto).Timestamp()
	if cgt == 0 {
		return errors.New("Wrong date format")
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("birth", cgt)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}
