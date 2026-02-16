# sdkopen-go

SDK modular em Go para construcao de servicos. Cada modulo e opcional — voce inicializa apenas o que precisa via `SdkOpenOptions`.

## Instalacao

```bash
go get github.com/sdkopen/sdkopen-go
```

## Inicializacao

A funcao `Initialize` recebe um `*SdkOpenOptions` onde cada campo e opcional. Apenas os modulos preenchidos serao inicializados:

```go
package main

import (
    sdkopen "github.com/sdkopen/sdkopen-go"
    "github.com/sdkopen/sdkopen-go/database"
    "github.com/sdkopen/sdkopen-go/messaging"
    "github.com/sdkopen/sdkopen-go/restserver"
)

func main() {
    sdkopen.Initialize(&sdkopen.SdkOpenOptions{
        Database:   database.Postgresql,
        Messaging:  messaging.RabbitMQ(),
        RestServer: restserver.CreateChiServer,
    })
}
```

## Modulos

### Database

Conexao com banco de dados relacional via provider pattern.

| Provider | Factory |
|----------|---------|
| PostgreSQL | `database.Postgresql` |

```go
sdkopen.Initialize(&sdkopen.SdkOpenOptions{
    Database: database.Postgresql,
})
```

Variaveis de ambiente:

```env
SQL_DB_URL=localhost
SQL_DB_PORT=5432
SQL_DB_NAME=mydb
SQL_DB_USERNAME=postgres
SQL_DB_PASSWORD=secret
SQL_DB_SSL_MODE=disable
SQL_DB_DRIVER=postgres
```

Uso:

```go
stmt := database.NewStatement(ctx, "INSERT INTO users (name) VALUES ($1)", "Alice")
err := stmt.Execute()
```

Documentacao completa: [database/README.md](database/README.md)

### Messaging

Mensageria com publisher e consumer via provider pattern. Um unico provider configura ambos.

| Provider | Factory |
|----------|---------|
| RabbitMQ | `messaging.RabbitMQ()` |

```go
sdkopen.Initialize(&sdkopen.SdkOpenOptions{
    Messaging: messaging.RabbitMQ(),
})
```

Variaveis de ambiente:

```env
RABBITMQ_URL=localhost
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_VHOST=/
```

Uso:

```go
// Publicar
err := messaging.Publish(ctx, "order.created", body)

// Consumir
messaging.Subscribe("order.created", handler)
messaging.StartConsumer()
```

Documentacao completa: [messaging/README.md](messaging/README.md)

### REST Server

Servidor HTTP com suporte a controllers e middlewares.

| Provider | Factory |
|----------|---------|
| Chi | `restserver.CreateChiServer` |

```go
// Registra controllers e middlewares antes de inicializar
restserver.RegisterController(myController)
restserver.RegisterMiddleware(myMiddleware)

sdkopen.Initialize(&sdkopen.SdkOpenOptions{
    RestServer: restserver.CreateChiServer,
})
```

## Exemplos

### Servico HTTP com banco de dados

```go
func main() {
    env.Load()

    restserver.RegisterController(userController)

    sdkopen.Initialize(&sdkopen.SdkOpenOptions{
        Database:   database.Postgresql,
        RestServer: restserver.CreateChiServer,
    })
}
```

### Worker de mensageria com banco de dados

```go
func main() {
    env.Load()

    sdkopen.Initialize(&sdkopen.SdkOpenOptions{
        Database:  database.Postgresql,
        Messaging: messaging.RabbitMQ(),
    })

    messaging.Subscribe("order.created", handleOrder)
    messaging.StartConsumer()
}
```

### Apenas banco de dados

```go
func main() {
    env.Load()

    sdkopen.Initialize(&sdkopen.SdkOpenOptions{
        Database: database.Postgresql,
    })
}
```

## Modulos automaticos

Os seguintes modulos sao inicializados automaticamente (via `init()`), sem necessidade de configuracao:

- **validator** — validacao de structs
- **observer** — graceful shutdown (todos os modulos se registram automaticamente)
