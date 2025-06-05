package port

type Mailer interface {
	NewMessage() Mailer
	From(from *string) Mailer
	To(to *[]string) Mailer
	Subject(subject *string) Mailer
	Body(body *string) Mailer
	Send() error
}
