package producer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
)

const (
	topicJokeCreated = "joke_created"
	topicJokeUpdated = "joke_updated"
	topicJokeDeleted = "joke_deleted"

	eventJokeCreated = "joke_created"
	eventJokeUpdated = "joke_updated"
	eventJokeDeleted = "joke_deleted"
)

type message struct {
	ID    uint64 `json:"id"`
	Event string `json:"event"`
}

func prepareMessage(topic string, m message) (*sarama.ProducerMessage, error) {
	payload, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(payload),
	}

	return msg, nil
}
