package messaging

import (
	"fmt"

	"github.com/sdkopen/sdkopen-go/common/env"
	"github.com/sdkopen/sdkopen-go/logging"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	defaultConnectionURI string = "amqp://%s:%s@%s:%d/%s"

	rabbitConnectionSuccessMsg string = "rabbitmq connected"
	rabbitConnectionErrorMsg   string = "an error occurred while trying to connect to rabbitmq: %+v"
)

type RabbitMQConnector struct {
	host     string
	port     int
	username string
	password string
	vhost    string
}

func NewDefaultRabbitMQConnector() *RabbitMQConnector {
	return &RabbitMQConnector{
		host:     env.RABBITMQ_URL,
		port:     env.RABBITMQ_PORT,
		username: env.RABBITMQ_USERNAME,
		password: env.RABBITMQ_PASSWORD,
		vhost:    env.RABBITMQ_VHOST,
	}
}

func (c *RabbitMQConnector) Connect() *amqp.Connection {
	conn, err := amqp.Dial(c.getConnectionURI())
	if err != nil {
		logging.Fatal(rabbitConnectionErrorMsg, err)
	}

	logging.Info(rabbitConnectionSuccessMsg)
	return conn
}

func (c *RabbitMQConnector) getConnectionURI() string {
	return fmt.Sprintf(defaultConnectionURI,
		c.username,
		c.password,
		c.host,
		c.port,
		c.vhost,
	)
}
