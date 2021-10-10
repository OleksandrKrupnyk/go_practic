// Отправка электронных писем созданных из шаблона
package main

import (
	"bytes"
	"log"
	"net/smtp"
	"strconv"
	"text/template"
)

type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `
From: {{.From}}
To: {{.To}}
Subject {{.Subject}}

{{.Body}}
`

// Переменная для хранения шаблона
var t *template.Template

func init() {
	// Шаблон
	t = template.New("email")
	_, err := t.Parse(emailTemplate)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func main() {
	message := &EmailMessage{
		From:    "me@localhost",
		To:      []string{"you@localhost"},
		Subject: "This is a test",
		Body:    "Just say hi!",
	}
	// Переменная для хранения результатов парсинга письма
	var body bytes.Buffer
	err := t.Execute(&body, message)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	authCreds := &EmailCredentials{
		Username: "u",
		Password: "p",
		Server:   "smtp.mailtrap.io",
		Port:     2525,
	}
	// Настрока smpt клиента
	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
	)
	// Отправка электронного письма
	err = smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		body.Bytes())
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}