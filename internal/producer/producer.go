package producer

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/ozonva/ova-joke-api/internal/configs"
	"github.com/ozonva/ova-joke-api/internal/models"
)

type Producer struct {
	sarama.SyncProducer
}

func NewProducer(config configs.BrokerConfig) (*Producer, error) {
	producer, err := sarama.NewSyncProducer(config.Addrs, nil)
	if err != nil {
		return nil, err
	}
	return &Producer{producer}, err
}

func (p *Producer) sendMsg(topic string, m message) (int32, int64, error) {
	payload, err := prepareMessage(topic, m)
	if err != nil {
		return 0, 0, err
	}

	return p.SendMessage(payload)
}

func (p *Producer) SendJokeCreatedMsg(ctx context.Context, id models.JokeID) (int32, int64, error) {
	ev := eventJokeCreated
	msg := message{
		ID:    id,
		Event: ev,
	}

	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("kafka_%s", ev))
	span.LogFields(log.String("msg", fmt.Sprintf("%v", msg)))
	defer span.Finish()

	return p.sendMsg(topicJokeCreated, msg)
}

func (p *Producer) SendJokeUpdatedMsg(ctx context.Context, id models.JokeID) (int32, int64, error) {
	ev := eventJokeUpdated
	msg := message{
		ID:    id,
		Event: ev,
	}

	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("kafka_%s", ev))
	span.LogFields(log.String("msg", fmt.Sprintf("%v", msg)))
	defer span.Finish()

	return p.sendMsg(topicJokeUpdated, msg)
}

func (p *Producer) SendJokeDeletedMsg(ctx context.Context, id models.JokeID) (int32, int64, error) {
	ev := eventJokeDeleted
	msg := message{
		ID:    id,
		Event: ev,
	}

	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("kafka_%s", ev))
	span.LogFields(log.String("msg", fmt.Sprintf("%v", msg)))
	defer span.Finish()

	return p.sendMsg(topicJokeDeleted, msg)
}
