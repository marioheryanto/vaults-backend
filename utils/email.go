package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmail() {
	// Sender's email configuration
	senderEmail := "marioheryanto.tech@gmail.com"
	senderPassword := "knkfpprpoukfjadc"

	// Recipient's email address
	recipientEmail := "marioheryanto.tech@gmail.com"

	// Email content
	subject := "Test Email from Gomail"
	body := "This is a test email sent from Gomail."

	// Initialize the email message
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Create a new SMTP sender
	d := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
