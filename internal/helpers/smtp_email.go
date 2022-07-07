package helpers

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

var CONFIG_SMTP_HOST = os.Getenv("CONFIG_SMTP_HOST")
var CONFIG_SMTP_PORT = os.Getenv("CONFIG_SMTP_PORT")
var CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
var CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
var CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

func SMTPEmail(send_to string, subject_email string, message_email string) error {

	to := []string{send_to}
	// to := []string{"kap21kap@gmail.com", "emaillain@gmail.com"}
	cc := []string{"krisnaanggapamungkas@gmail.com"}
	subject := subject_email
	message := message_email

	err := sendMail(to, cc, subject, message)
	if err != nil {
		return err
	}

	log.Println("Mail sent!")
	return nil
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
