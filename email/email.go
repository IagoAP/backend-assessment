package email

import (
	"github.com/sirupsen/logrus"
	"net/smtp"
	"psT10/database"
	"psT10/environment"
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
	from := environment.GetEnvVariables("EMAIL_FROM")
	password := environment.GetEnvVariables("EMAIL_PASSWORD")

	smtpHost := environment.GetEnvVariables("EMAIL_HOST")
	smtpPort := environment.GetEnvVariables("EMAIL_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, emailTo, message)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return
}