# Database

Modulo de banco de dados do sdkopen-go. Utiliza o provider pattern para desacoplar a inicializacao do driver concreto, permitindo trocar o banco de dados sem alterar o codigo de negocio.

## Arquitetura

```
database/
├── database.go                 # Initialize(factory) e variavel dbInstance
├── observer.go                 # Graceful shutdown via observer pattern
├── statement.go                # Statement para execucao de queries
└── postgresql_connector.go     # Implementacao PostgreSQL + factory Postgresql()
```

## Configuracao

Defina as variaveis de ambiente (ou use um arquivo `.env`):

```env
SQL_DB_URL=localhost
SQL_DB_PORT=5432
SQL_DB_NAME=mydb
SQL_DB_USERNAME=postgres
SQL_DB_PASSWORD=secret
SQL_DB_SSL_MODE=disable
SQL_DB_DRIVER=postgres
```

As variaveis sao carregadas automaticamente pelo `env.Load()` na inicializacao da aplicacao.

## Inicializacao

```go
package main

import (
    "github.com/sdkopen/sdkopen-go/common/env"
    "github.com/sdkopen/sdkopen-go/database"
)

func main() {
    env.Load()

    // Inicializa o banco com o provider PostgreSQL
    database.Initialize(database.Postgresql)
}
```

`Initialize` recebe uma factory function `func() *sql.DB`, executa a factory para criar a conexao e registra o observer para graceful shutdown automaticamente.

## Executando queries

Use `Statement` para executar queries com prepared statements:

```go
ctx := context.Background()

// INSERT
stmt := database.NewStatement(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", "Alice", "alice@example.com")
if err := stmt.Execute(); err != nil {
    log.Fatal(err)
}

// UPDATE
stmt = database.NewStatement(ctx, "UPDATE users SET name = $1 WHERE id = $2", "Bob", 1)
if err := stmt.Execute(); err != nil {
    log.Fatal(err)
}

// DELETE
stmt = database.NewStatement(ctx, "DELETE FROM users WHERE id = $1", 1)
if err := stmt.Execute(); err != nil {
    log.Fatal(err)
}
```

### Executando em uma instancia especifica

Se precisar executar em uma instancia diferente da global, use `ExecuteInInstance`:

```go
stmt := database.NewStatement(ctx, "INSERT INTO logs (message) VALUES ($1)", "test")
if err := stmt.ExecuteInInstance(customDB); err != nil {
    log.Fatal(err)
}
```

### Transacoes

O `Statement` suporta transacoes via context. Basta passar um context com a transacao:

```go
tx, _ := db.BeginTx(ctx, nil)
txCtx := context.WithValue(ctx, "SqlTxContext", tx)

stmt := database.NewStatement(txCtx, "INSERT INTO orders (total) VALUES ($1)", 99.90)
if err := stmt.Execute(); err != nil {
    tx.Rollback()
    log.Fatal(err)
}

tx.Commit()
```

## Graceful Shutdown

O modulo se integra automaticamente com o `observer` para shutdown graceful:

1. O observer aguarda as operacoes em andamento terminarem (via WaitGroup)
2. Se o timeout for atingido, forca o encerramento
3. Fecha a conexao com o banco de dados

Isso acontece automaticamente ao usar `Initialize` — nao e necessaria nenhuma configuracao adicional.

## Implementando um novo provider

Para criar um novo provider (ex: MySQL, SQLite), basta criar uma factory function que retorne `*sql.DB`:

```go
func MySQL() *sql.DB {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
    if err != nil {
        logging.Fatal("could not connect to mysql: %v", err)
    }
    if err = db.Ping(); err != nil {
        logging.Fatal("could not ping mysql: %v", err)
    }
    return db
}
```

Uso:

```go
database.Initialize(MySQL)
```

## Exemplo completo

```go
package main

import (
    "context"
    "log"

    "github.com/sdkopen/sdkopen-go/common/env"
    "github.com/sdkopen/sdkopen-go/database"
)

func main() {
    if err := env.Load(); err != nil {
        log.Fatal(err)
    }

    database.Initialize(database.Postgresql)

    ctx := context.Background()

    stmt := database.NewStatement(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", "Alice", "alice@example.com")
    if err := stmt.Execute(); err != nil {
        log.Fatal(err)
    }
}
```
