package helper

import (
	"fmt"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

type serverEmail struct {
	email    *email.Email
	password string
	from     string
}

var ServerEmail *serverEmail

func init() {
	// 这里读配置比较合理
	ServerEmail = &serverEmail{}
	ServerEmail.from = "go_blog2021@163.com"
	ServerEmail.password = "KBDPVPDJXAAYMEKB"
	ServerEmail.email = &email.Email{
		From:        "go_blog2021@163.com",
		Subject:     "go-blog",
		Headers:     textproto.MIMEHeader{},
		Attachments: nil,
		ReadReceipt: nil,
	}
}

func (e *serverEmail) SendCode(to string, code string) error {
	e.email.Text = []byte(fmt.Sprintf("[go-blog]您的验证码是 %s，5分钟内有效，请勿泄漏", code))
	e.email.To = []string{to}
	return e.email.Send("smtp.163.com:25", smtp.PlainAuth("", e.from, e.password, "smtp.163.com"))
}
