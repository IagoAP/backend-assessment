package server

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"psT10/database"
	"psT10/database2"
	"psT10/email"
)

func HandleProductCreate (message *sarama.ConsumerMessage, kafka *Kafka) error {
	productRequest, err := database.ConvertProductMessage(message.Value)
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = database.CreateProduct(productRequest)
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = kafka.SendMessage(productRequest, "ProductCreateReadDB")
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = email.SendEmailProductCreated(message.Value)
	if err != nil {
		logrus.Infof(err.Error())
	}
	return err
}

func HandleProductCreateReadDB (message *sarama.ConsumerMessage, kafka *Kafka) error {
	readModel, err := database2.ConvertReadModel(message.Value)
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = database2.CreateRow(readModel)
	if err != nil {
		logrus.Infof(err.Error())
	}
	return err
}

func HandleProductActivation (message *sarama.ConsumerMessage, kafka *Kafka) error {
	activationRequest, err := database.ConvertActivationMessage(message.Value)
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = database.ActivateProduct(activationRequest)
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = kafka.SendMessage(activationRequest, "ProductActivationReadDB")
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = email.SendEmailProductEvaluation(message.Value)
	if err != nil {
		logrus.Infof(err.Error())
	}
	return err
}

func HandleProductActivationReadDB (message *sarama.ConsumerMessage, kafka *Kafka) error {
	readModel, err := database2.ConvertReadModel(message.Value)
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	err = database2.UpdateRow(readModel)
	if err != nil {
		logrus.Infof(err.Error())
	}
	return err
}