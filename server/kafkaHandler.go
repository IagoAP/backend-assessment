package server

import (
	"github.com/Shopify/sarama"
	"psT10/database"
	"psT10/database2"
)

func HandleProductCreate (message *sarama.ConsumerMessage, kafka *Kafka) {
	database.CreateProduct(database.ConvertProductMessage(message.Value))
	kafka.SendMessage(message, "ProductCreateReadDB")
}

func HandleProductCreateReadDB (message *sarama.ConsumerMessage, kafka *Kafka) {
	database2.CreateRow(database2.ConvertReadModel(message.Value))
}

func HandleProductActivation (message *sarama.ConsumerMessage, kafka *Kafka) {
	database.ActivateProduct(database.ConvertActivationMessage(message.Value))
	kafka.SendMessage(message, "ProductActivationReadDB")
}

func HandleProductActivationReadDB (message *sarama.ConsumerMessage, kafka *Kafka) {
	database2.UpdateRow(database2.ConvertReadModel(message.Value))
}