# ova-joke-api

[![codecov](https://codecov.io/gh/ozonva/ova-joke-api/branch/issue-10/graph/badge.svg)](https://codecov.io/gh/ozonva/ova-joke-api)
![example workflow](https://github.com/ozonva/ova-joke-api/actions/workflows/main.yml/badge.svg)
![gomod version](https://img.shields.io/github/go-mod/go-version/ozonva/ova-joke-api/master)

## Ozon Golang school project

  This application provides API for Joke entity management system. All data stored in OLTP database 
[PostgreSQL](https://www.postgresql.org). Service provides:
- creating new jokes,
- multiple creating,
- updating them,
- removing,
- reading specific joke,
- reading joke ranges specified by offset and limit.

## Used components and ports

| Component    | Description                                          | Port  |
|--------------|------------------------------------------------------|-------|
| db           | PostgreSQL OLTP database, used as persistent storage | 5432  |
| zookeeper    | ZooKeeper- configuration management system for kafka | 2181  |
| kafka        | Apache Kafka- message broker                         | 9092  |
| kafkaui      | Kowl- web ui for kafka, allow observe broker state   | 8087  |
| jaeger       | Jaeger- tracing system                               | 16686 |
| prometheus   | Prometheus- metric management system                 | 9091  |
| grafana      | Project metrics visualization                        | 3000  |
| godoc        | Documentation pages about project components and deps| 9099  |
| gRPC API     | Application API endpoint                             | 9090  |
| metrics API  | Application metrics endpoint for Prometheus          | 9093  |

## Configuration:

| Flag                | Config path       | Env                   | Default value  | Description                       |
|---------------------|-------------------|-----------------------|----------------|-----------------------------------|
| --grpc.addr         | grpc.addr         | OVA_GRPC_ADDR         | 0.0.0.0.9090   | addr:port of gRPC server endpoint |
| --flusher.chunksize | flusher.chunksize | OVA_FLUSHER_CHUNKSIZE | 3              | insert database batch size        |
| --db.host           | db.host           | OVA_DB_HOST           | localhost      | host for database                 |
| --db.port           | db.port           | OVA_DB_PORT           | 5432           | database port                     |
| --db.name           | db.name           | OVA_DB_NAME           | postgres       | database name                     |
| --db.user           | db.user           | OVA_DB_USER           | postgres       | database user name                |
| --db.pass           | db.pass           | OVA_DB_PASS           | postgres       | database user's password          |
| --metrics.addr      | metrics.addr      | OVA_METRICS_ADDR      | 0.0.0.0:9093   | addr of metrics exporter api      |
| --broker.addrs      | broker.addrs      | OVA_BROKER_ADDRS      | [0.0.0.0:9092] | addr of metrics exporter api      |

Settings appied in priority:
- flag
- env
- config

  Passing several values for broker.addrs, see [viper](https://github.com/spf13/viper) and
[pflags](https://github.com/spf13/pflag/) for more information:
```bash
# using options:
ova-joke-api --broker.addrs=127.0.0.1:9092 --broker.addrs=127.0.0.1:9093
# using envs:
OVA_BROKER_ADDRS=127.0.0.1:9092,127.0.0.1:9093 ova-joke-api
```

## Run service:

1. launch docker env:
```bash
docker-compose up
```
2.1 run api server using golang:
```bash
make run
# or call directly:
# go run cmd/ova-joke-api/main.go
```
2.2. or using binary:
```bash
# compile binary
make build
# run binary
./bin/ova-joke-api
```

## Development:

1. Linters- used [golangci-lint](https://github.com/golangci/golangci-lint). For more information about linter 
   configuration see [.golangci.yml](./.golangci.yml).
```bash
# run using make:
make lint
# or directly:
golangci-lint run ./...
```
2. Tests
```bash
# run using make:
make test
# or directly:
go test -tags=test_unit -v -count=1 -race -timeout=1m ./...
```
3. Code coverage level and change you can see at [https://codecov.io](https://codecov.io/gh/ozonva/ova-joke-api)
4. Regenerate protobuf dependencies use
```bash
make generate
```
