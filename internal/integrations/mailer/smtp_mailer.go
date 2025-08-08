package mailer

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type mailerSMTP struct {
	host     string
	port     string
	username string
	password string
	from     string
}

type message struct {
	to      []string
	subject string
	body    string
}

func NewSMTPMailer(
	host *string, port *string, username *string,
	password *string, from *string,
) Mailer {
	return &mailerSMTP{
		host:     *host,
		port:     *port,
		username: *username,
		password: *password,
		from:     *from,
	}
}

func (m *mailerSMTP) SendMessage(to []string, subject string, body string) error {

	message := message{
		to:      to,
		subject: subject,
		body:    body,
	}

	if err := validateMessage(message); err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"", m.username,
		m.password, m.host,
	)

	header := "" +
		fmt.Sprintf("Subject: %s\n", subject) +
		fmt.Sprintf("To: %s\n", strings.Join(to, ",")) +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte(header + body)
	if err := smtp.SendMail(
		m.host+":"+m.port, auth, m.from, to, msg,
	); err != nil {
		return err
	}

	return nil

}

func validateMessage(m message) error {
	if len(m.to) == 0 ||
		m.subject == "" || m.body == "" {
		return errors.New("missing required parameters")
	}

	return nil
}
