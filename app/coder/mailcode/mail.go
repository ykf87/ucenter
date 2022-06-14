package mailcode

import (
	"errors"
	"sync"
	"time"
)

type MailCodeStruct struct { //一个已经发送验证码并维护验证码的结构体
	Code     string //验证码内容
	Errtimes int    //输入错误次数
	Expire   int64  //过期时间
}

type RWMap struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map字段
	m            map[string]*MailCodeStruct
}

var Timeout = 600 // timeout 验证码有效期,单位秒
var Maps *RWMap

func init() {
	Maps = &RWMap{
		m: make(map[string]*MailCodeStruct),
	}
	go Maps.checkMap(60)
}

func (m *RWMap) checkMap(timer int) { //定时检查邮箱过期二维码
	for {
		nowTime := time.Now().Unix()
		m.Each(func(k string, v *MailCodeStruct) bool {
			if v.Expire <= nowTime {
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
	rs, has := Maps.Get(mail)
	if !has {
		return errors.New("Incorrect Captcha, please resend the Captcha")
	}
	if rs.Expire <= time.Now().Unix() {
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
