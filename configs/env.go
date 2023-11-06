package configs

import (
	"DENV_Register_Producer/models"
)

func GetKafkaProperties() (models.KafkaProperties, error) {
	kfProperties := models.KafkaProperties{
		Host: "192.168.1.13",
		Port: "9092",
	}

	return kfProperties, nil
}
