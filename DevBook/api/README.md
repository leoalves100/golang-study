# DevBook

## Dependencies use in project

1. [Cobra CLI](https://github.com/spf13/cobra) # CLI command
2. [Migrator](https://github.com/larapulse/migrator) # Migration database schema
3. [godotenv](https://github.com/joho/godotenv) # Variable environment
4. [mysql-driver](https://github.com/go-sql-driver/mysql) # Driver MySQL

## Util

1. [REST API Tutorial](https://restfulapi.net/)
2. [Status Code HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

## Structure project

```bash
├── cmd
│   ├── gen-base64.go
│   ├── migrate
│   │   ├── migration.go
│   │   └── revert.go
│   ├── root.go
│   └── server.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── README.md
└── src
    ├── authentication # Create and validate JWT token
    │   └── token.go
    ├── config # Loads the variables (.env) in the application
    │   └── config.go
    ├── controllers # Application flow control
    │   ├── login.go
    │   └── users.go
    ├── database # Connect to the database
    │   ├── database.go 
    │   └── migrations # Database migrations
    │       └── 2023_03_20_create_users_table.go
    ├── middlewares # Middleware of token and log
    │   └── middlewares.go
    ├── model # Validates, formats and prepares the users model
    │   └── user.go
    ├── repository
    │   └── users.go
    ├── response # API response template
    │   └── response.go
    ├── router
    │   ├── rotas
    │   │   ├── login.go
    │   │   ├── rotas.go
    │   │   └── users.go
    │   └── router.go
    └── sec
        └── sec.go

```

## Deploy

1. Clone project

    ```bash
    git clone git@github.com:leoalves100/golang-study.git
    ```

2. Enter path

    ```bash
    cd DevBook/api
    ```

3. Run docker compose

    ```bash
    docker compose up -d

    #In case it is necessary to recreate the containers
    docker compose up -d --force-recreate

    #Check logs
    docker compose logs -f
    ```

4. Run the migrations

    ```bash
    go run main.go migration
    ```

5. Run the server

    ```bash
    go run main.go server
    ```
