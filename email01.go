package main

import (
	"log"
	"net/smtp"
)

func main() {

	from := "example@notreal.com"
	password := "password123"
	to := []string{"recipient@notreal.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("My secret message")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
