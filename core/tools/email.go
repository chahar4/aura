package tools

import (
	"crypto/tls"
	"os"

	"gopkg.in/gomail.v2"
)

func NotificationSender(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_NOTIFICATION"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_NOTIFICATION"), os.Getenv("PASSWORD_NOTIFICATION"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return EmailErr
	}

	return nil
}
