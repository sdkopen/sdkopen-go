package messaging

import (
	"testing"

	"github.com/sdkopen/sdkopen-go/common/env"
)

func TestNewDefaultRabbitMQConnector(t *testing.T) {
	env.RABBITMQ_URL = "rabbit-host"
	env.RABBITMQ_PORT = 5672
	env.RABBITMQ_USERNAME = "guest"
	env.RABBITMQ_PASSWORD = "guest"
	env.RABBITMQ_VHOST = "/"

	connector := NewDefaultRabbitMQConnector()

	if connector.host != "rabbit-host" {
		t.Fatalf("expected host=rabbit-host, got %s", connector.host)
	}
	if connector.port != 5672 {
		t.Fatalf("expected port=5672, got %d", connector.port)
	}
	if connector.username != "guest" {
		t.Fatalf("expected username=guest, got %s", connector.username)
	}
	if connector.password != "guest" {
		t.Fatalf("expected password=guest, got %s", connector.password)
	}
	if connector.vhost != "/" {
		t.Fatalf("expected vhost=/, got %s", connector.vhost)
	}
}

func TestRabbitMQConnector_GetConnectionURI(t *testing.T) {
	connector := &RabbitMQConnector{
		host:     "mq.example.com",
		port:     5673,
		username: "admin",
		password: "secret",
		vhost:    "/prod",
	}

	uri := connector.getConnectionURI()
	expected := "amqp://admin:secret@mq.example.com:5673//prod"

	if uri != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, uri)
	}
}

func TestRabbitMQConnector_GetConnectionURI_DefaultValues(t *testing.T) {
	connector := &RabbitMQConnector{
		host:     "localhost",
		port:     5672,
		username: "guest",
		password: "guest",
		vhost:    "/",
	}

	uri := connector.getConnectionURI()
	expected := "amqp://guest:guest@localhost:5672//"

	if uri != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, uri)
	}
}

func TestRabbitMQ_ReturnsProvider(t *testing.T) {
	provider := RabbitMQ()

	if provider == nil {
		t.Fatal("expected non-nil provider")
	}
	if provider.CreatePublisher == nil {
		t.Fatal("expected non-nil CreatePublisher")
	}
	if provider.CreateConsumer == nil {
		t.Fatal("expected non-nil CreateConsumer")
	}
}
