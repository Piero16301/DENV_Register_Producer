package configs

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func ConnectProducer() sarama.SyncProducer {
	fmt.Println("Conectando a Kafka...")

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	kfProperties, err := GetKafkaProperties()
	if err != nil {
		log.Fatal(err)
	}

	producer, err := sarama.NewSyncProducer([]string{kfProperties.GetDSN()}, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado a Kafka en " + kfProperties.GetDSN())

	return producer
}

// Producer Instancia de Cliente
var Producer = ConnectProducer()
