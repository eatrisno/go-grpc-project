package utils

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, data interface{}, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), smtpPort, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))
	err := d.DialAndSend(m)
	if err != nil {
		log.Println(err)
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
		log.Printf("send email %s to %s to success", subject, to)
	} else {
		log.Printf("send email %s to %s to failed: %s", subject, to, err)
	}
}
