package main

import (
	"github.com/dimkasta/goemail"
	"github.com/dimkasta/gologger"
	"github.com/dimkasta/goteplate"
)

func main() {
	logger := gologger.NewLoggerService()
	logger.Info("Initialized")

	repository := goteplate.NewSqliteTemplateRepository("templates.db", logger)
	templater := goteplate.NewTemplateService(logger, repository)

	data := make(map[string]string)
	data["test"] = "value new"

	output, err := templater.Get("test", data)

	if nil != err {
		logger.Error(err.Error())
	}

	message := goemail.NewHtmlMail()
	message.SetFrom("dimkasta@yahoo.gr", "Dimitris")
	message.AddTo("d.kastaniotis@iconic.gr", "Giorgos")
	message.SetSubject("Subject goes Here")
	message.SetBody(output)

	mailer := goemail.NewMailer(logger, "localhost:1025")
	mailer.Send(message)
}
