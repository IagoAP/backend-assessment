package server

import (
	"github.com/Shopify/sarama"
	"psT10/database"
)

func HandleProductCreate (message *sarama.ConsumerMessage, kafka *Kafka) {
	database.CreateProduct(database.ConvertProductMessage(message.Value))
	kafka.SendMessage(message, "ProductCreateReadDB")
}

func HandleProductCreateReadDB (message *sarama.ConsumerMessage, kafka *Kafka) {

}

func HandleProductActivation (message *sarama.ConsumerMessage, kafka *Kafka) {
	database.ActivateProduct(database.ConvertActivationMessage(message.Value))
	kafka.SendMessage(message, "ProductActivationReadDB")
}