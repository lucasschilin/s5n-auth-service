package mailer

type Mailer interface {
	SendMessage(to []string, subject string, body string) error
}
