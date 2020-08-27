package email

import (
	"github.com/sirupsen/logrus"
	"net/smtp"
	"psT10/database"
)

func SendEmailProductCreated(message []byte) error {
	var err error = nil
	result, err := database.ConvertProductMessage(message)
	if err != nil {
		return err
	}
	var emailTo []string
	emailTo = append(emailTo, result.CustomerEmail)
	sendEmail(emailTo, message)
	return err
}

func SendEmailProductEvaluation(message []byte) error {
	result, err := database.ConvertActivationMessage(message)
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	var emailTo []string
	email, err := database.GetEmail(result)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	emailTo = append(emailTo, email)
	sendEmail(emailTo, message)
	return err
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