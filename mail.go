package main

import (
	"net/mail"
	"net/smtp"
	"github.com/scorredoira/email"
	"time"
)

func SendingEmail() {
	// compose the message
	m := email.NewMessage("留扬宝Nginx日志分析统计" + time.Now().Format("2006-01-02") + "(系统自动发送，无需回复!)", "统计结果见附件")
	m.From = mail.Address{Name: "【云深监控系统】", Address: t.EmailConfig.Sender}
	m.To = t.EmailConfig.Receivers

	// add attachments
	if err := m.Attach(result_path); err != nil {
		Check(err)
	}

	// send it
	auth := smtp.PlainAuth("", t.EmailConfig.UserName, t.EmailConfig.Password, t.EmailConfig.SmtpHost)
	if err := email.Send(t.EmailConfig.SmtpHost + ":587", auth, m); err != nil {
		Check(err)
	}
}
