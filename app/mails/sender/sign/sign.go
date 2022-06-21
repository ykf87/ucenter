//注册成功邮件发送
package sign

import (
	"log"
	"ucenter/app/config"
	"ucenter/app/i18n"
	"ucenter/app/mails/smtp"

	"github.com/matcornic/hermes/v2"
)

func Send(mail, lang string) {
	if mail == "" {
		log.Println("Send sign mail error, the email is empty!")
		return
	}

	s, err := smtp.Client("sign")
	if err != nil {
		log.Println("Send sign mail error:", err)
		return
	}

	msg := i18n.T(lang, "Hello, I am a Frisky Meets assistant, welcome to join Frisky Meets, I am very happy to have a small partner, I wish you can find like-minded friends in Frisky Meets, start a heart-to-heart dating, welcome your valuable comments on Frisky Meets.")

	email := hermes.Email{
		Body: hermes.Body{
			Title: string(msg),
			Outros: []string{
				string(i18n.T(lang, "Need help? Please contact us in the app, this email will not accept any reply.")),
			},
			Signature: string(i18n.T(lang, "Have a good day")),
		},
	}
	h := smtp.SmtpModel()
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		log.Println("Send sign mail error, when generateHtml:", err)
		return
	}

	sub := i18n.T(lang, "Welcome to {{$1}}", config.Config.APPName)
	r := s.SetGeter(mail).SetMessage(emailBody).SetSubject(string(sub)).Send()
	if r != nil {
		log.Println("Send sign mail error, when do send:", r)
		return
	}

	return
}
