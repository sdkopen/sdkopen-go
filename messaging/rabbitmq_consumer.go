package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	conn          *amqp.Connection
	channel       *amqp.Channel
	subscriptions []Subscription
	done          chan struct{}
}

func CreateRabbitMQConsumer() Consumer {
	connector := NewDefaultRabbitMQConnector()
	conn := connector.Connect()

	ch, err := conn.Channel()
	if err != nil {
		logging.Fatal("failed to open rabbitmq channel: %+v", err)
	}

	return &RabbitMQConsumer{
		conn:    conn,
		channel: ch,
		done:    make(chan struct{}),
	}
}

func (c *RabbitMQConsumer) Subscribe(subscription Subscription) {
	c.subscriptions = append(c.subscriptions, subscription)
}

func (c *RabbitMQConsumer) Start() error {
	for _, sub := range c.subscriptions {
		if err := c.consume(sub); err != nil {
			return fmt.Errorf("failed to start consumer for %s: %w", sub.Topic, err)
		}
	}

	<-c.done
	return nil
}

func (c *RabbitMQConsumer) consume(sub Subscription) error {
	err := c.channel.ExchangeDeclare(
		sub.Topic,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange %s: %w", sub.Topic, err)
	}

	q, err := c.channel.QueueDeclare(
		sub.Topic,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", sub.Topic, err)
	}

	err = c.channel.QueueBind(
		q.Name,
		sub.Topic,
		sub.Topic,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue %s: %w", sub.Topic, err)
	}

	deliveries, err := c.channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume from queue %s: %w", sub.Topic, err)
	}

	go func() {
		for d := range deliveries {
			wg := observer.GetWaitGroup()
			wg.Add(1)

			go func(delivery amqp.Delivery) {
				defer wg.Done()

				msg := Message{
					ID:        delivery.MessageId,
					Topic:     sub.Topic,
					Body:      delivery.Body,
					Headers:   extractHeaders(delivery.Headers),
					Timestamp: delivery.Timestamp,
				}

				if msg.Timestamp.IsZero() {
					msg.Timestamp = time.Now()
				}

				if err := sub.Handler(context.Background(), msg); err != nil {
					logging.Error("error handling message on topic %s: %v", sub.Topic, err)
					_ = delivery.Nack(false, true)
					return
				}

				_ = delivery.Ack(false)
			}(d)
		}
	}()

	logging.Info("consuming messages from topic: %s", sub.Topic)
	return nil
}

func (c *RabbitMQConsumer) Close() error {
	close(c.done)

	if c.channel != nil {
		if err := c.channel.Close(); err != nil {
			logging.Error("error closing rabbitmq consumer channel: %v", err)
		}
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func extractHeaders(table amqp.Table) map[string]string {
	headers := make(map[string]string)
	for k, v := range table {
		if s, ok := v.(string); ok {
			headers[k] = s
		}
	}
	return headers
}
