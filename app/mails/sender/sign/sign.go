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
	return
	if mail == "" {
		log.Println("Send sign mail error, the email is empty!")
		return
	}

	s, err := smtp.Client("sign")
	if err != nil {
		log.Println("Send sign mail error:", err)
		return
	}

	email := hermes.Email{
		Body: hermes.Body{
			Actions: []hermes.Action{
				{
					InviteCode: "",
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
		log.Println("Send sign mail error, when generateHtml:", err)
		return
		// return errors.New("Mail delivery failure")
	}

	sub := i18n.T(lang, "Welcome", config.Config.APPName)
	r := s.SetGeter(mail).SetMessage(emailBody).SetSubject(string(sub)).Send()
	if r != nil {
		log.Println("Send sign mail error, when do send:", err)
		return
		// return errors.New("Mail delivery failure")
	}

	return
}
