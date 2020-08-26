package server

import (
	"github.com/Shopify/sarama"
	"psT10/database"
	"psT10/database2"
)

func HandleProductCreate (message *sarama.ConsumerMessage, kafka *Kafka) {
	productRequest := database.ConvertProductMessage(message.Value)
	database.CreateProduct(productRequest)
	kafka.SendMessage(productRequest, "ProductCreateReadDB")
}

func HandleProductCreateReadDB (message *sarama.ConsumerMessage, kafka *Kafka) {
	database2.CreateRow(database2.ConvertReadModel(message.Value))
}

func HandleProductActivation (message *sarama.ConsumerMessage, kafka *Kafka) {
	activationRequest := database.ConvertActivationMessage(message.Value)
	database.ActivateProduct(activationRequest)
	kafka.SendMessage(activationRequest, "ProductActivationReadDB")
}

func HandleProductActivationReadDB (message *sarama.ConsumerMessage, kafka *Kafka) {
	database2.UpdateRow(database2.ConvertReadModel(message.Value))
}