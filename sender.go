package mail_sender

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

// NewSender configures SMTP sender.
func NewSender(cfg Configuration) *Sender {
	return &Sender{
		smtpServer: cfg.SMTPServer,
		smtpPort:   cfg.SMTPPort,
		user:       cfg.MailUser,
		password:   cfg.MailPassword,
	}
}

// WritePlainEmail sends a plain text email.
func (sender *Sender) WritePlainEmail(dest []string, subject, message string) error {
	return sender.writeEmail(&Message{
		Recipients:  dest,
		ContentType: contentTypePlain,
		Subject:     subject,
		Body:        []byte(message),
	})
}

// WriteHTMLEmail sends an HTML email based on your HTML template.
func (sender *Sender) WriteHTMLEmail(dest []string, subject, tmplName string, data interface{}) error {
	temp, err := template.ParseFiles(tmplName)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = temp.Execute(&body, data)
	if err != nil {
		return err
	}

	return sender.writeEmail(&Message{
		Recipients:  dest,
		ContentType: contentTypeHTML,
		Subject:     subject,
		Body:        body.Bytes(),
	})
}

func (sender Sender) writeEmail(m *Message) error {
	header := make(map[string]string)
	header["MIME-Version"] = mimeVersion
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", m.ContentType)
	header["Subject"] = m.Subject

	var headers string
	for key, value := range header {
		headers += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	msg := make([]byte, len(headers)+len(m.Body))
	i := copy(msg, headers)
	copy(msg[i:], m.Body)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", sender.smtpServer, sender.smtpPort),
		smtp.PlainAuth("", sender.user, sender.password, sender.smtpServer),
		sender.user,
		m.Recipients,
		msg,
	)
}
