package coder

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
	"ucenter/app/config"
	"ucenter/app/i18n"
	"ucenter/app/mails/smtp"

	"github.com/matcornic/hermes/v2"
)

type MailCodeStruct struct { //一个已经发送验证码并维护验证码的结构体
	Code     string //验证码内容
	Errtimes int    //输入错误次数
	Sendtime int64  //过期时间
}

type RWMap struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map字段
	m            map[string]*MailCodeStruct
}

var Timeout int64 = 300   //timeout 验证码有效期,单位秒
var ResendTime int64 = 60 //重新发送验证码间隔
var Maps *RWMap

func init() {
	Maps = &RWMap{
		m: make(map[string]*MailCodeStruct),
	}
	go Maps.checkMap(60)
}

func (m *RWMap) checkMap(timer int) { //定时检查邮箱过期验证码
	for {
		nowTime := time.Now().Unix()
		m.Each(func(k string, v *MailCodeStruct) bool {
			if v.Sendtime+Timeout <= nowTime {
				m.Delete(k)
			}
			return true
		})
		time.Sleep(time.Minute * time.Duration(timer))
	}
}

func (m *RWMap) Get(k string) (*MailCodeStruct, bool) { //从map中读取一个值
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k] // 在锁的保护下从map中读取
	return v, existed
}

func (m *RWMap) Set(k string, v *MailCodeStruct) { // 设置一个键值对
	m.Lock() // 锁保护
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k string) { //删除一个键
	m.Lock() // 锁保护
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int { // map的长度
	m.RLock() // 锁保护
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k string, v *MailCodeStruct) bool) { // 遍历map
	m.RLock() //遍历期间一直持有读锁
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}

func Verify(mail, code string) error {
	if code == "" {
		return errors.New("Please input your Captcha")
	}
	if config.Config.Universalcaptcha != "" && code == config.Config.Universalcaptcha {
		return nil
	}
	rs, has := Maps.Get(mail)
	if !has {
		return errors.New("Incorrect Captcha, please resend the Captcha")
	}
	if rs.Sendtime+Timeout <= time.Now().Unix() {
		return errors.New("Incorrect Captcha, please resend the Captcha")
	}
	if rs.Errtimes >= 5 {
		Maps.Delete(mail)
		return errors.New("Too many times the Captcha error, please resend the Captcha")
	}
	if rs.Code != code {
		rs.Errtimes = rs.Errtimes + 1
		return errors.New("Incorrect Captcha")
	}
	Maps.Delete(mail)
	return nil
}

// 发送验证码
func Send(mail, lang string) error {
	if mail == "" {
		return errors.New("Please set the Email")
	}
	mc, ok := Maps.Get(mail)
	if ok {
		if mc.Sendtime+ResendTime > time.Now().Unix() {
			return errors.New("Your verification code is still valid")
		} else {
			Maps.Delete(mail)
		}
	}

	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	s, err := smtp.Client("coder")
	if err != nil {
		return err
	}

	email := hermes.Email{
		Body: hermes.Body{
			Actions: []hermes.Action{
				{
					InviteCode: code,
				},
			},
			Outros: []string{
				string(i18n.T(lang, "Need help? Please contact us in the app, this email will not accept any reply.")),
			},
			Title:     string(i18n.T(lang, "Captcha is valid for 10 minutes, and expires after use or after 5 times of wrong input, please keep it safe!")),
			Signature: string(i18n.T(lang, "Have a good day")),
		},
	}
	h := smtp.SmtpModel()
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		log.Println("Send Email Code Err(coder): ", err)
		return errors.New("Captcha sending failure")
	}

	sub := i18n.T(lang, "{{$1}} verify the authenticity of your email", config.Config.APPName)
	r := s.SetGeter(mail).SetMessage(emailBody).SetSubject(string(sub)).Send()
	if r != nil {
		log.Println("Send Email Code Err: ", r)
		return errors.New("Captcha sending failure")
	}

	Maps.Set(mail, &MailCodeStruct{Code: code, Sendtime: time.Now().Unix(), Errtimes: 0})
	return nil
}
