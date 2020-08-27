package email

import (
	"github.com/sirupsen/logrus"
	"net/smtp"
	"psT10/database"
)

func SendEmailProductCreated(message []byte) {
	result := database.ConvertProductMessage(message)
	var emailTo []string
	emailTo = append(emailTo, result.CustomerEmail)
	sendEmail(emailTo, message)
}

func SendEmailProductEvaluation(message []byte) {
	result := database.ConvertActivationMessage(message)
	var emailTo []string
	email, err := database.GetEmail(result)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	emailTo = append(emailTo, email)
	sendEmail(emailTo, message)
}

func sendEmail(emailTo []string, message []byte) {
	from := "iagopst10@gmail.com" // ENV
	password := "pst10iago" // ENV

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, emailTo, message)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return
}