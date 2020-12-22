package main

import (
	"github.com/sirupsen/logrus"
	"psT10/environment"
	"psT10/server"
)

func main() {
	kafka := server.Kafka{
		Brokers: []string{environment.GetEnvVariables("BROKER_HOST")},
		GroupID: environment.GetEnvVariables("KAFKA_GROUP"),
		Version: environment.GetEnvVariables("KAFKA_VERSION"),
	}
	kafkaClose, err := kafka.GetDefaultConfig()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	defer kafkaClose()
	s := server.Server{Kafka: kafka}
	go s.Kafka.CreateConsumers([]string{environment.GetEnvVariables("CUSTOMER1"), environment.GetEnvVariables("CUSTOMER2"),
		environment.GetEnvVariables("CUSTOMER3"), environment.GetEnvVariables("CUSTOMER4")})
	s.Run()
}
