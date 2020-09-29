package main

/*************************************************************************
Builder Parameters example

*/
import (
	"fmt"
	"strings"
)

type Email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email Email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

type build func(*EmailBuilder)

func sendMailImpl(email *Email) {
	// actually ends the email
	fmt.Println("Sender", email.from)
	fmt.Println("Receiver", email.to)
	fmt.Println("Subject", email.subject)
	fmt.Println("Body", email.body)
}

// SendEmail accepts method with *EmailBuilder as an argument
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}
