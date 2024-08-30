package email1

import (
	"log"

	"gopkg.in/gomail.v2"
)

func sendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "shaxbozakramovic@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "shaxbozakramovic@gmail.com", "vakhaboff2502")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func Sent(email, body string) error {
	to := email
	subject := "New Notification"

	if err := sendEmail(to, subject, body); err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}

