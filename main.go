package main

import (
	"bytes"
	"github.com/dimkasta/mailer/mail"
	"github.com/dimkasta/mailer/service"
	"text/template"
)

func main() {
	logger := service.NewLoggerService()
    logger.Log.Info("Initialized")

	html := "" +
		"<html><body><h1>Hello World!<small>{{ .test }}</small></h1></body></html>"

	template, _ := template.New("UsersPage").Parse(string(html))

	var body bytes.Buffer

	data := make(map[string]string)
	data["test"] = "value"

	template.Execute(&body, data)

	email := mail.NewHtmlMail()
	email.SetFrom("dimkasta@yahoo.gr", "Dimitris")
	email.AddTo("d.kastaniotis@iconic.gr","Giorgos")
	email.SetSubject("Subject goes Here")
	email.SetBody(body.String())

	mailer := mail.NewMailer(logger, "localhost:1025")
	mailer.Send(email)
}
