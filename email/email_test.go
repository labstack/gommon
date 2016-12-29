package email

import (
	"net/smtp"
	"testing"
)

func TestSend(t *testing.T) {
	// e := New("smtp.gmail.com:465")
	e := New("smtp.elasticemail.com:2525")
	// e.Auth = smtp.PlainAuth("", "vr@labstack.com", "Dream1980", "smtp.gmail.com")
	e.Auth = smtp.PlainAuth("", "vr@labstack.com", "54643b5b-3284-4f33-89bf-d228951c527f", "smtp.elasticemail.com")
	// err := smtp.SendMail("smtp.elasticemail.com:2525", e.Auth, "vr@labstack.com", []string{"ranavishal@gmail.com"}, []byte("body"))
	// fmt.Println(err)
	e.Send(&Message{
		From:    "no-reply@labstack.com",
		To:      "ranavishal@gmail.com",
		Subject: "test",
		Text:    "xxxxxxxxxxxxxxxxxxxxxxxxxx",
	})
}
