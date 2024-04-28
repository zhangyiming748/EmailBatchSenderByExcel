package email

import (
	"gopkg.in/gomail.v2"
	"log/slog"
	"time"
)

const (
	Host     = "smtp.163.com" //你自己的邮件服务器地址
	Port     = 25             // "你自己的邮件服务器端口"
	Username = ""             // "你自己的邮件服务器用户名"
	Password = ""             //你自己的邮件服务器密码或授权码
)

func init() {
	initLocal()
}

func initLocal() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func Send(from, to, subject, body, image string) {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("发送邮件发生错误", slog.Any("错误原文", err))
		}
	}()
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if image != "" {
		m.Attach(image)
	}
	d := gomail.NewDialer(Host, Port, Username, Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	slog.Info("发送邮件")
}
