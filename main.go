package main

import (
	"github.com/sirupsen/logrus"
	"psT10/server"
)

func main() { // ENV
	kafka := server.Kafka{
		Brokers: []string{"localhost:9092"},
		GroupID: "T10",
		Version: "0.11.0.2",
	}
	kafkaClose, err := kafka.GetDefaultConfig()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	defer kafkaClose()
	s := server.Server{Kafka: kafka}
	go s.Kafka.CreateConsumers([]string{"ProductCreate", "ProductCreateReadDB", "ProductActivation", "ProductActivationReadDB"}) // Colocar no env
	s.Run()
}
