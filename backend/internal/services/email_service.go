package services

import (
	"fmt"
	"net/smtp"
	"net/url"
)

type EmailConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	From        string
	FrontendURL string
}

type EmailService struct {
	config EmailConfig
}

func NewEmailService(
	config EmailConfig,
) *EmailService {
	return &EmailService{
		config: config,
	}
}

func (s *EmailService) SendVerificationEmail(
	to string,
	token string,
) error {
	verifyURL := fmt.Sprintf(
		"%s/verify-email?token=%s",
		s.config.FrontendURL,
		url.QueryEscape(token),
	)

	subject := "ยืนยันอีเมลของคุณ"

	body := fmt.Sprintf(
		`สวัสดี

กรุณายืนยันอีเมลของคุณโดยกดลิงก์ด้านล่าง:

%s

ลิงก์นี้จะหมดอายุภายใน 15 นาที
หากคุณไม่ได้สมัครสมาชิก สามารถละเว้นอีเมลนี้ได้
`,
		verifyURL,
	)

	message := []byte(
		"From: " + s.config.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n" +
			"\r\n" +
			body,
	)

	auth := smtp.PlainAuth(
		"",
		s.config.Username,
		s.config.Password,
		s.config.Host,
	)

	return smtp.SendMail(
		s.config.Host+":"+s.config.Port,
		auth,
		s.config.From,
		[]string{to},
		message,
	)
}
