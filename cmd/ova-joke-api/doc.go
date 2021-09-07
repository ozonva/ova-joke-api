// Package ova-joke-api provides gRPC API for Jokes management
//
// Command options:
//
//  --grpc.addr         (OVA_GRPC_ADDR)         default: 0.0.0.0.9090  "addr:port" of gRPC server endpoint
//  --flusher.chunksize (OVA_FLUSHER_CHUNKSIZE) default: 3              insert database batch size
//  --db.host           (OVA_DB_HOST)           default: localhost      host for database
//  --db.port           (OVA_DB_PORT)           default: 5432           database port
//  --db.name           (OVA_DB_NAME)           default: postgres       database name
//  --db.user           (OVA_DB_USER)           default: postgres       database user name
//  --db.pass           (OVA_DB_PASS)           default: postgres       database user's password
//  --metrics.addr      (OVA_METRICS_ADDR)      default: 0.0.0.0:9093   addr of metrics exporter api
//  --broker.addrs      (OVA_BROKER_ADDRS)      default: [0.0.0.0:9092] addr of metrics exporter api
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
