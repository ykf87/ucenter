package coder

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"ucenter/app/config"
	"ucenter/app/dbs/redis"
	"ucenter/app/i18n"
	"ucenter/app/logs"
	"ucenter/app/mails/smtp"

	"github.com/matcornic/hermes/v2"
)

var timeout int = 300      //timeout 验证码有效期,单位秒
var ResendTime int64 = 60  //重新发送验证码间隔
var prev string = "coder/" //redis前缀

//生成code的key
func k(mail string) string {
	return prev + mail
}

//设置code,如果存在将删除
func SetCode(mail, code string) error {
	val := fmt.Sprintf("%s:%d", code, 0)
	key := k(mail)

	ss, errs := redis.Get(key)
	if errs == nil && ss != "" {
		redis.Del(key)
	}

	err := redis.Set(key, val, timeout)
	if err != nil {
		logs.Logger.Error(err)
	}
	return err
}

//获取内容
func GetCode(mail string) (code string, errtimes int, ttl int, err error) {
	key := k(mail)
	ss, errs := redis.Get(key)
	if errs != nil {
		err = errs
		return
	}
	if ss == "" {
		redis.Del(key)
		err = errors.New("The Captcha has expired")
		return
	}
	sps := strings.Split(ss, ":")
	if len(sps) != 2 {
		redis.Del(key)
		err = errors.New("The Captcha has expired")
		return
	}
	errtimes, _ = strconv.Atoi(sps[1])
	code = sps[0]
	ttls, _ := redis.Ttl(key)
	ttl = ttls
	return
}

//错误次数加1
func Increment(mail string) error {
	key := k(mail)

	code, errtimes, ttl, err := GetCode(mail)
	if err != nil {
		return err
	}

	val := fmt.Sprintf("%s:%d", code, (errtimes + 1))
	redis.Set(key, val, ttl)
	return nil
}

//验证code
func Verify(mail, code string) error {
	if code == "" {
		return errors.New("Please input your Captcha")
	}

	if config.Config.Universalcaptcha != "" && code == config.Config.Universalcaptcha {
		return nil
	}

	key := k(mail)
	coder, errtimes, _, err := GetCode(mail)

	if err != nil {
		return errors.New("Incorrect Captcha, please resend the Captcha")
	}

	if errtimes >= 4 {
		redis.Del(key)
		return errors.New("Too many times the Captcha error, please resend the Captcha")
	}
	if code != coder {
		Increment(mail)
		return errors.New("Incorrect Captcha")
	}

	return nil
}

// 发送验证码
func Send(mail, lang string) (error, int) {
	if mail == "" {
		return errors.New("Please set the Email"), 0
	}
	coder, _, ttl, err := GetCode(mail)
	if coder != "" {
		if timeout-ttl <= 60 {
			return errors.New("Your verification code is still valid"), ttl
		}
	}

	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	err = SetCode(mail, code)
	if err != nil {
		return err, 0
	}
	key := k(mail)

	s, err := smtp.Client("coder")
	if err != nil {
		redis.Del(key)
		return err, 0
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
		redis.Del(key)
		log.Println("Send Email Code Err(coder): ", err)
		logs.Logger.Error("Send Email Code Err(coder: " + mail + "): " + err.Error())
		return errors.New("Captcha sending failure"), 0
	}

	sub := i18n.T(lang, "{{$1}} verify the authenticity of your email", config.Config.APPName)
	r := s.SetGeter(mail).SetMessage(emailBody).SetSubject(string(sub)).Send()
	if r != nil {
		redis.Del(key)
		logs.Logger.Error("Send Email Code Err(coder: " + mail + "): " + r.Error())
		log.Println("Send Email Code Err(coder: "+mail+"): ", r)
		return errors.New("Captcha sending failure"), 0
	}

	return nil, timeout
}
