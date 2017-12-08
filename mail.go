package main

import (
	"net/mail"
	"net/smtp"
	"github.com/scorredoira/email"
	"time"
)

var Body string

func SendingEmail() {
	m := email.NewHTMLMessage(t.EmailConfig.Subject + time.Now().Format("2006-01-02"), Body)
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
