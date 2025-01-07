package main

import (
	"crypto/tls"
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	// 设置SMTP服务器
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "13138892570@163.com") // 发件人
	mailer.SetHeader("To", "1029612787@qq.com")     // 收件人
	mailer.SetHeader("Subject", "邮件主题")             // 邮件主题
	mailer.SetBody("text/html", "邮件正文内容")           // 邮件内容

	// 163邮箱SMTP服务器配置
	smtpHost := "smtp.163.com"
	smtpPort := 25
	smtpUser := "13138892570@163.com"
	smtpPass := "xxxxxx" // 这不是你的邮箱密码，而是开启SMTP服务后获得的授权码

	// 构建SMTP客户端
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true, // 忽略证书校验，仅用于测试环境
	}

	// 发送邮件
	if err := dialer.DialAndSend(mailer); err != nil {
		log.Fatal(err)
	}

	log.Println("邮件发送成功!")
}
