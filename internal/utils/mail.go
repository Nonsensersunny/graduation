package utils

import (
	"gopkg.in/gomail.v2"
	"graduation/internal/config"
	"regexp"
)

func SendMail(subject, content string, tos... string) error {
	dialer := config.GetMailDialer()
	m := gomail.NewMessage()
	m.SetHeader("From", dialer.Username)
	m.SetHeader("Subject", subject)
	m.SetHeader("To", tos...)
	m.SetBody("text/html", "<h1>" + content + "</h1>")
	return dialer.DialAndSend(m)
}

func CheckMail(mail string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mail)
}