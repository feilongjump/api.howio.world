package mail

import (
	"fmt"
	"github.com/feilongjump/api.howio.world/internal/config"
	SimpleMail "github.com/xhit/go-simple-mail/v2"
	"time"
)

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

// NewMail 创建 Mail 发送内容
func NewMail(to string, subject string, body string) Mail {
	return Mail{
		From:    config.GetString("mail.from_name"),
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

// SendEmail 发送邮件
func (mail Mail) SendEmail() bool {
	smtpClient, err := newMailClient()
	if err != nil {
		fmt.Println("Expected nil, got", err, "connecting to client")
		return false
	}

	// Create the email message
	email := SimpleMail.NewMSG()

	email.SetFrom(mail.From).
		AddTo(mail.To).
		SetSubject(mail.Subject).
		SetBody(SimpleMail.TextHTML, mail.Body)

	// Send with high priority
	email.SetPriority(SimpleMail.PriorityHigh)

	// always check error after send
	if email.Error != nil {
		fmt.Println(email.Error)
		return false
	}

	// Pass the client to the email message to send it
	if err = email.Send(smtpClient); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// newMailClient 创建 Mail 发送客户端
func newMailClient() (*SimpleMail.SMTPClient, error) {
	server := SimpleMail.NewSMTPClient()

	// SMTP Client
	server.Host = config.GetString("mail.host")
	server.Port = config.GetInt("mail.port")
	server.Username = config.GetString("mail.username")
	server.Password = config.GetString("mail.password")
	// port 587
	//server.Encryption = SimpleMail.EncryptionSTARTTLS
	// port 465
	server.Encryption = SimpleMail.EncryptionSSLTLS
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	// Connect to client
	return server.Connect()
}
