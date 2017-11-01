package main

import (
	"net/mail"
	"net/smtp"
	"github.com/scorredoira/email"
)

func SendingEmail() {
	// compose the message
	m := email.NewMessage("Anginx notification", "This is the analysis result")
	m.From = mail.Address{Name: "From", Address: t.EmailConfig.Sender}
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
