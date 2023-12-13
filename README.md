# weather forecast server

This is a simple weather forecast app that uses the Openmeteo API to get the weather forecast for the next 7 days for a specific place. Every 5 minutes the worker will make a call to Openmeteo API and update the database with the new forecast. The server will return the forecast from the database.

## Requirements

- [Go](https://go.dev/doc/install) >= v1.21
- [Task](https://taskfile.dev/#/) >= 3.10
- [Docker](https://docs.docker.com/get-docker/)

## How to run the server

### Postgres

Run the following command to start a Postgres container:

```bash
task docker:postgres-up env=.local.env
```

Then run the following command to create the database and run the migrations:

```bash
task db:migrate-up-all env=.local.env
```

Or you can just run all at once:

```bash
task db:init env=.local.env
```

The `db:init` starts the Postgres container, creates the database and runs the migrations.

### Server

Run the following command to start the server:

```bash
task run:server -- local
```

By default, the server will run in port `8080`, but you can change it by setting the `PORT` environment variable.

## How to run the tests

Run the following command to run the tests:

```bash
task test:all
```

## How to stop Postgres

Run the following command to stop the Postgres container:

```bash
task docker:postgres-down
```

## APIs

Once the server is running you can access the following APIs:

- [/forecast](http://localhost:8080/forecast)
- [/previsao](http://localhost:8080/previsao)
- [/alerts](http://localhost:8080/alerts)
- [/alerta](http://localhost:8080/alerta)

## Linters

The following linters are available, but it needs to be installed locally:

- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [gosec](https://github.com/securego/gosec)
- [govulncheck](https://go.dev/blog/govulncheck)

To run each linter, run the following commands:

```bash
task linter:golangci-lint
task linter:gosec
task linter:govulncheck
```

Or to run all at once:

```bash
task linter:all
```
