package adapter

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/lucasschilin/s5n-auth-service/internal/port"
)

type mailerSMTP struct {
	host         string
	port         string
	username     string
	password     string
	defaultFrom  string
	emailMessage emailMessage
}

type emailMessage struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewSMTPMailer(
	host *string, port *string, username *string,
	password *string, defaultFrom *string,
) port.Mailer {
	return &mailerSMTP{
		host:        *host,
		port:        *port,
		username:    *username,
		password:    *password,
		defaultFrom: *defaultFrom,
	}
}

func (m *mailerSMTP) NewMessage() port.Mailer {
	m.emailMessage = emailMessage{
		from: m.defaultFrom,
	}

	return m
}

func (m *mailerSMTP) From(from *string) port.Mailer {
	m.emailMessage.from = *from

	return m
}

func (m *mailerSMTP) To(to *[]string) port.Mailer {
	m.emailMessage.to = *to

	return m
}

func (m *mailerSMTP) Subject(subject *string) port.Mailer {
	m.emailMessage.subject = *subject

	return m
}

func (m *mailerSMTP) Body(body *string) port.Mailer {
	m.emailMessage.body = *body

	return m
}

func (m *mailerSMTP) Send() error {
	if err := m.validateMessage(); err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"", m.username,
		m.password, m.host,
	)

	header := "" +
		fmt.Sprintf("Subject: %s\n", m.emailMessage.subject) +
		fmt.Sprintf("To: %s\n", strings.Join(m.emailMessage.to, ",")) +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body := m.emailMessage.body

	msg := []byte(header + body)
	if err := smtp.SendMail(
		m.host+":"+m.port, auth, m.emailMessage.from, m.emailMessage.to, msg,
	); err != nil {
		return err
	}

	return nil
}

func (m *mailerSMTP) validateMessage() error {
	if m.emailMessage.from == "" || len(m.emailMessage.to) == 0 ||
		m.emailMessage.subject == "" || m.emailMessage.body == "" {
		return errors.New("missing required parameters")
	}

	return nil
}
