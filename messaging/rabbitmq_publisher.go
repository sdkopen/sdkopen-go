package messaging

import (
	"context"
	"fmt"

	"github.com/sdkopen/sdkopen-go/logging"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func CreateRabbitMQPublisher() Publisher {
	connector := NewDefaultRabbitMQConnector()
	conn := connector.Connect()

	ch, err := conn.Channel()
	if err != nil {
		logging.Fatal("failed to open rabbitmq channel: %+v", err)
	}

	return &RabbitMQPublisher{
		conn:    conn,
		channel: ch,
	}
}

func (p *RabbitMQPublisher) Publish(ctx context.Context, topic string, body []byte, opts ...PublishOption) error {
	cfg := applyOptions(opts)

	err := p.channel.ExchangeDeclare(
		topic,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange %s: %w", topic, err)
	}

	headers := amqp.Table{}
	for k, v := range cfg.Headers {
		headers[k] = v
	}
	if cfg.DelaySeconds > 0 {
		headers["x-delay"] = int32(cfg.DelaySeconds * 1000)
	}

	err = p.channel.PublishWithContext(
		ctx,
		topic,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Headers:     headers,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message to %s: %w", topic, err)
	}

	return nil
}

func (p *RabbitMQPublisher) Close() error {
	if p.channel != nil {
		if err := p.channel.Close(); err != nil {
			logging.Error("error closing rabbitmq publisher channel: %v", err)
		}
	}
	if p.conn != nil {
		return p.conn.Close()
	}
	return nil
}
