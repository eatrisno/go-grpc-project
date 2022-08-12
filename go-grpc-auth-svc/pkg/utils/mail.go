package utils

import (
	"bytes"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, data interface{}, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "eaprilitrisno@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	senderPort := 587
	d := gomail.NewDialer("smtp.gmail.com", senderPort, "eaprilitrisno@gmail.com", "nenapuvmeovtgrvx")
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return err
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func SendEmailTemplate(to string, subject string, data interface{}, template string) {
	var err error
	body, _ := ParseTemplate(template, data)
	err = SendEmail(to, subject, data, body)
	if err == nil {
		log.Println("send email '" + subject + "' success")
	} else {
		log.Println(err)
	}
}
