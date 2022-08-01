package smtp

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"ucenter/app/config"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type Stmp struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
	Appname  string
	Nickname string
	Sender   string
	Geter    []string
	Message  string
	Attaches []string
	Subject  string
}

var configs map[string]*config.SmtpConf

func Client(key string) (s *Stmp, err error) {
	if configs == nil {
		err = errors.New("SMTP configuration exception")
		return
	}
	if key == "" {
		key = "default"
	}
	c, ok := configs[key]
	if !ok {
		if key == "default" {
			err = errors.New("SMTP configuration exception")
			return
		}
		c, ok = configs["default"]
		if !ok {
			err = errors.New("SMTP configuration exception")
			return
		}
	}
	s = &Stmp{
		Host:    c.Host,
		Port:    c.Port,
		Email:   c.Email,
		Pass:    c.Pass,
		Sender:  c.Sender,
		Appname: config.Config.APPName,
	}
	return
}

func SetConfig(key string, c *config.SmtpConf) bool {
	if key == "" {
		return false
	}
	if configs == nil {
		configs = make(map[string]*config.SmtpConf)
	}
	configs[key] = c
	return true
}

func (this *Stmp) SetSender(mail string) *Stmp {
	this.Sender = mail
	return this
}

func (this *Stmp) SetGeter(mail string) *Stmp {
	this.Geter = append(this.Geter, mail)
	return this
}
func (this *Stmp) SetAtta(attache string) *Stmp {
	if _, err := os.Stat(attache); err == nil {
		this.Attaches = append(this.Attaches, attache)
	}
	return this
}
func (this *Stmp) SetNickname(nickname string) *Stmp {
	this.Nickname = nickname
	return this
}

func (this *Stmp) SetMessage(content string) *Stmp {
	this.Message = this.Message + content
	return this
}
func (this *Stmp) SetSubject(title string) *Stmp {
	this.Subject = title
	return this
}

func (this *Stmp) Send() error {
	m := gomail.NewMessage()
	if this.Sender == "" {
		this.Sender = this.Email
	}
	if this.Nickname == "" {
		this.Nickname = this.Appname
	}
	if this.Subject == "" {
		tmp := strings.Split(this.Sender, "@")
		this.Subject = "From " + tmp[0]
	}
	m.SetHeader("Subject", this.Subject)
	m.SetHeader("From", fmt.Sprintf("%s<%s>", this.Nickname, this.Sender))
	m.SetHeader("To", this.Geter...)
	m.SetBody("text/html", this.Message)
	if this.Attaches != nil {
		for _, v := range this.Attaches {
			m.Attach(v)
		}
	}

	d := gomail.NewDialer(this.Host, this.Port, this.Email, this.Pass)
	// d.TLSConfig = &tls.Config{}
	err := d.DialAndSend(m)
	return err
}

func SmtpModel() hermes.Hermes {
	copyright := config.Config.Copyright
	if copyright == "" {
		copyright = "Copyright © " + config.Config.APPName + ". All rights reserved."
	}
	mailLogo := config.Config.Mainbanner
	if mailLogo == "" {
		mailLogo = config.Config.Logo
	}
	return hermes.Hermes{
		Product: hermes.Product{
			Name: config.Config.APPName,
			Link: config.Config.Domain,
			// Logo:        "https://www.zhishukongjian.com/reset/images/logo.png",
			Logo:      mailLogo,
			Copyright: copyright,
		},
	}
}
