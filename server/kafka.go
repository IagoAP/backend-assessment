package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"time"
)

type Kafka struct {
	Brokers  []string
	GroupID  string
	Version  string
	Client   sarama.Client
	Consumer sarama.ConsumerGroup
	Producer sarama.SyncProducer
}

func (k *Kafka) GetDefaultConfig() (func(), error) {
	kfVersion, err := sarama.ParseKafkaVersion(k.Version) // kafkaVersion is the version of kafka server like 0.11.0.2
	if err != nil {
		return nil, err
	}
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = kfVersion
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	k.Client, err = sarama.NewClient(k.Brokers, kafkaConfig)

	if err != nil {
		return nil, err
	}

	k.Consumer, err = sarama.NewConsumerGroupFromClient(k.GroupID, k.Client)
	if err != nil {
		return nil, err
	}

	k.Producer, err = sarama.NewSyncProducerFromClient(k.Client)

	if err != nil {
		return nil, err
	}

	def := func() {
		go func() {
			if err := k.Client.Close(); err != nil {
				logrus.Fatal(err)
			}
		}()
		go func() {
			if err := k.Consumer.Close(); err != nil {
				logrus.Fatal(err)
			}
		}()
		go func() {
			if err := k.Producer.Close(); err != nil {
				logrus.Fatal(err)
			}
		}()
	}
	return def, nil
}

func (k *Kafka) SendMessage(message interface{}, topic string) error {
	object, err := json.Marshal(message)
	if err != nil {
		logrus.Errorf("error unsmarshalling message: %v", err)
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(object),
		Timestamp: time.Now(),
	}
	partition, offset, err := k.Producer.SendMessage(msg)
	if err != nil {
		logrus.Errorf("error sending message to kafka: %s", err.Error())
		return err
	}
	logrus.Infof("Message is stored in topic %s/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func messageReceived(message *sarama.ConsumerMessage, kafka *Kafka) error {
	switch message.Topic {
	case "ProductCreate":
		HandleProductCreate(message, kafka)
	case "ProductCreateReadDB":

	case "ProductActivation":
		HandleProductActivation(message, kafka)
	case "ProductActivationReadDB":
		//database2.CompleteRow(database.ConvertActivationMessage(message.Value))
	}
	return nil
}

func (k *Kafka) CreateConsumers(readingTopics []string) {
	for {
		topics := readingTopics
		handler := &ConsumerHandler{k}
		err := k.Consumer.Consume(context.Background(), topics, handler)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

// ConsumerGroupHandler instances are used to handle individual topic/partition claims.
// It also provides hooks for your consumer group session life-cycle and allow you to
// trigger logic before or after the consume loop(s).
//
// PLEASE NOTE that handlers are likely be called from several goroutines concurrently,
// ensure that all state is safely protected against race conditions.
type ConsumerHandler struct {
	kafka *Kafka
	// Setup is run at the beginning of a new session, before ConsumeClaim.

	// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
	// but before the offsets are committed for the very last time.

	// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
	// Once the Messages() channel is closed, the Handler must finish its processing
	// loop and exit.
}

func (ch *ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (ch *ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (ch ConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		logrus.Info("Processing Message topic:" + fmt.Sprint(msg.Topic) + "partition:" + fmt.Sprint(msg.Partition) + " offset:" + fmt.Sprint(msg.Offset))
		err := messageReceived(msg, ch.kafka)
		if err != nil {
			return err
		}
		sess.MarkMessage(msg, "")
		logrus.Info("Message processed topic:" + fmt.Sprint(msg.Topic) + "partition:" + fmt.Sprint(msg.Partition) + " offset:" + fmt.Sprint(msg.Offset))
	}
	return nil
}
