# Messaging

Modulo de mensageria do sdkopen-go. Fornece interfaces abstratas (`Publisher`, `Consumer`) e uma implementacao concreta para **RabbitMQ** usando a lib [amqp091-go](https://github.com/rabbitmq/amqp091-go).

O modulo utiliza o provider pattern — um unico provider configura tanto o publisher quanto o consumer, garantindo que ambos usem o mesmo backend de mensageria.

## Arquitetura

```
messaging/
├── messaging.go              # Initialize(provider), Provider struct
├── publisher.go              # Interface Publisher e funcao Publish
├── consumer.go               # Interface Consumer, Subscribe e StartConsumer
├── message.go                # Struct Message e PublishOption (functional options)
├── observer.go               # Graceful shutdown via observer pattern
├── rabbitmq_connector.go     # Conexao AMQP (RabbitMQConnector) + factory RabbitMQ()
├── rabbitmq_publisher.go     # Implementacao Publisher para RabbitMQ
└── rabbitmq_consumer.go      # Implementacao Consumer para RabbitMQ
```

## Configuracao

Defina as variaveis de ambiente (ou use um arquivo `.env`):

```env
RABBITMQ_URL=localhost
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_VHOST=/
```

As variaveis sao carregadas automaticamente pelo `env.Load()` na inicializacao da aplicacao.

## Inicializacao

```go
package main

import (
    "github.com/sdkopen/sdkopen-go/messaging"
)

func main() {
    // Inicializa publisher e consumer com o provider RabbitMQ
    messaging.Initialize(messaging.RabbitMQ())
}
```

`Initialize` recebe um `*Provider` que contem as factories de publisher e consumer. O publisher e criado imediatamente; o consumer fica disponivel para ser iniciado via `StartConsumer()`.

## Publisher

### Publicando mensagens

```go
ctx := context.Background()

// Publicacao simples
err := messaging.Publish(ctx, "order.created", orderBytes)

// Com headers customizados
err := messaging.Publish(ctx, "order.created", orderBytes,
    messaging.WithHeaders(map[string]string{
        "correlation-id": "abc-123",
        "source":         "order-service",
    }),
)

// Com delay (requer plugin rabbitmq_delayed_message_exchange)
err := messaging.Publish(ctx, "order.retry", orderBytes,
    messaging.WithDelay(30), // 30 segundos
)

// Combinando opcoes
err := messaging.Publish(ctx, "order.created", orderBytes,
    messaging.WithHeaders(map[string]string{"source": "api"}),
    messaging.WithDelay(10),
)
```

### Comportamento do Publisher

- Declara automaticamente o exchange do tipo `topic` (durable)
- Envia mensagens com `ContentType: application/json`
- Headers sao mapeados para `amqp.Table`
- `DelaySeconds` e convertido para o header `x-delay` em milissegundos

## Consumer

### Registrando handlers

Registre os handlers **antes** de chamar `StartConsumer()`:

```go
messaging.Subscribe("order.created", handleOrderCreated)
messaging.Subscribe("payment.confirmed", handlePayment)

// Inicia o consumer (bloqueia a goroutine atual)
messaging.StartConsumer()
```

### Comportamento do Consumer

- Para cada subscription, declara automaticamente: exchange (topic, durable), queue (durable) e binding
- Cada mensagem e processada em uma goroutine separada
- Usa `observer.GetWaitGroup()` para garantir graceful shutdown
- **Sucesso**: handler retorna `nil` -> mensagem recebe `Ack`
- **Erro**: handler retorna `error` -> mensagem recebe `Nack` com requeue (volta para a fila)
- `Start()` e bloqueante — mantem o consumer rodando ate `Close()` ser chamado

### Struct Message

O handler recebe um `messaging.Message` com os seguintes campos:

```go
type Message struct {
    ID        string            // MessageId do AMQP
    Topic     string            // Nome do topico/exchange
    Body      []byte            // Corpo da mensagem
    Headers   map[string]string // Headers extraidos da mensagem
    Timestamp time.Time         // Timestamp da mensagem
}
```

## Graceful Shutdown

O modulo se integra automaticamente com o `observer` para shutdown graceful:

1. O observer aguarda as mensagens em processamento terminarem (via WaitGroup)
2. Se o timeout for atingido, forca o encerramento
3. Fecha o channel e a conexao AMQP

Isso acontece automaticamente ao usar `Initialize` e `StartConsumer` — nao e necessaria nenhuma configuracao adicional.

## Exemplo completo

```go
package main

import (
    "context"
    "encoding/json"
    "log"

    "github.com/sdkopen/sdkopen-go/common/env"
    "github.com/sdkopen/sdkopen-go/messaging"
)

type Order struct {
    ID    string  `json:"id"`
    Total float64 `json:"total"`
}

func main() {
    if err := env.Load(); err != nil {
        log.Fatal(err)
    }

    // Inicializa o provider RabbitMQ (publisher + consumer)
    messaging.Initialize(messaging.RabbitMQ())

    // --- Publicando eventos ---
    order := Order{ID: "123", Total: 99.90}
    body, _ := json.Marshal(order)

    err := messaging.Publish(context.Background(), "order.created", body)
    if err != nil {
        log.Fatal(err)
    }

    // --- Consumindo eventos ---
    messaging.Subscribe("order.created", func(ctx context.Context, msg messaging.Message) error {
        var o Order
        if err := json.Unmarshal(msg.Body, &o); err != nil {
            return err
        }
        log.Printf("Pedido recebido: %s - R$%.2f", o.ID, o.Total)
        return nil
    })

    messaging.StartConsumer()
}
```

## Implementando um novo provider

Para criar um novo provider (ex: Kafka, SQS), implemente as interfaces `Publisher` e `Consumer`:

```go
type Publisher interface {
    Publish(ctx context.Context, topic string, body []byte, opts ...PublishOption) error
    Close() error
}

type Consumer interface {
    Subscribe(subscription Subscription)
    Start() error
    Close() error
}
```

E crie uma funcao que retorne o `*Provider` com as factories:

```go
func Kafka() *messaging.Provider {
    return &messaging.Provider{
        CreatePublisher: func() messaging.Publisher { ... },
        CreateConsumer:  func() messaging.Consumer { ... },
    }
}
```

Uso:

```go
messaging.Initialize(Kafka())
```
