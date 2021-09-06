// Package ova-joke-api provides gRPC API for Jokes management
//
// Command options:
//
//  -port, default "0.0.0.0:9090"- port for gRPC api server
//  -db-host, default "localhost"- host for database
//  -db-port, default "5432"- database port
//  -db-name, default "postgres"- database name
//  -db-user, default "postgres"- database user name
//  -db-pass, default "postgres"- database users' password
//  -metrics-addr, default "0.0.0.0:9093"- addr of metrics exporter api& used by prometheus
//  -broker-addr, default "0.0.0.0:9092"- comma separated list of brokers addrs
//
// Launch gRPC service:
//
// Start dev env:
//  docker-compose up
//
// Compile joke api binary:
//  go build -o joke-api ./cmd/ova-joke-api/main.go
//
// Launch joke api server (you can pass come options):
//  ./joke-api
//
// Enjoy!
package main
