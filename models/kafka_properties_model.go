package models

type KafkaProperties struct {
	Host string
	Port string
}

func (k KafkaProperties) GetDSN() string {
	return k.Host + ":" + k.Port
}
