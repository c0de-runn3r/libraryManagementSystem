package utils

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendVerificationCodeViaEmail(emailAddress string, code string) error {
	emailSMTP := os.Getenv("EMAIL_SMTP")
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	user := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", "LMS authentication code")
	m.SetBody("text/html", "Hello!\nThis is your verification link:\nhttp://localhost:4000/api/users/verifyemail/"+code)

	d := gomail.NewDialer(emailSMTP, emailPort, user, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
