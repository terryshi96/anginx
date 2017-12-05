package main

import (
	"net/mail"
	"net/smtp"
	"github.com/scorredoira/email"
	"time"
)

func SendingEmail() {
	// compose the message
	body := t.EmailConfig.Body + "\n\n"
	len := len(data.LongestReq)
	for i:=0;i<len;i++ {
		for j:=0;j<7;j++ {
			body = body + data.LongestReq[i][j] + " "
		}
		body = body + "\n\n"
	}
	m := email.NewMessage(t.EmailConfig.Subject + time.Now().Format("2006-01-02"), body)
	m.From = mail.Address{Name: t.EmailConfig.SenderName, Address: t.EmailConfig.Sender}
	m.To = t.EmailConfig.Receivers
	m.Cc = t.EmailConfig.Cc

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
