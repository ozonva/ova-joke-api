package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozonva/ova-joke-api/internal/app/hellower"
	"github.com/ozonva/ova-joke-api/internal/configs"
	"github.com/ozonva/ova-joke-api/internal/flusher"
	log "github.com/ozonva/ova-joke-api/internal/logger"
	"github.com/ozonva/ova-joke-api/internal/metrics"
	api "github.com/ozonva/ova-joke-api/internal/ova-joke-api"
	"github.com/ozonva/ova-joke-api/internal/producer"
	"github.com/ozonva/ova-joke-api/internal/repo"
	"github.com/ozonva/ova-joke-api/internal/tracer"
)

const serviceName = "ova-joke-api"

var (
	grpcAddr    string
	flChunkSize int

	dbHost string
	dbPort uint16
	dbName string
	dbUser string
	dbPass string

	metricsAddr string

	brokerAddrs []string
)

func init() {
	pflag.StringVar(&grpcAddr, "grpc.addr", "0.0.0.0:9090", "port for gRPC api server")
	pflag.IntVar(&flChunkSize, "flusher.chunksize", 3, "storage insert batch size")

	pflag.StringVar(&dbHost, "db.host", "localhost", "host for database")
	pflag.Uint16Var(&dbPort, "db.port", 5432, "database port")
	pflag.StringVar(&dbName, "db.name", "postgres", "database name")
	pflag.StringVar(&dbUser, "db.user", "postgres", "database user name")
	pflag.StringVar(&dbPass, "db.pass", "postgres", "database users' password")

	pflag.StringVar(&metricsAddr, "metrics.addr", "0.0.0.0:9093", "addr of metrics exporter api")

	pflag.StringSliceVar(&brokerAddrs, "broker.addrs", []string{"0.0.0.0:9092"}, "comma separated list of brokers addrs")
}

func createDB(config configs.DBConfig) *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.GetDSN())
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	return db
}

func createRepo(db *sqlx.DB) *repo.JokePgRepo {
	return repo.NewJokePgRepo(db)
}

func createJokeFlusher(r *repo.JokePgRepo, config configs.FlusherConfig) *flusher.JokeFlusher {
	return flusher.NewJokeFlusher(config.ChunkSize, r)
}

func initSignalHandler() <-chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return done
}

func main() {
	// read configs, env variables and flags
	cfg, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	// create database client
	db := createDB(cfg.DB)
	defer func() {
		err := db.Close()
		if err != nil {
			panic(fmt.Sprintf("unable to close db connection %v", err))
		}
	}()

	// create repository (using postgres connection) and flusher (batch wrapper)
	dbRepo := createRepo(db)
	fl := createJokeFlusher(dbRepo, cfg.Flusher)

	// create metric storage and run handler for prometheus
	counters := metrics.NewMetrics()
	metrics.Run(cfg.Metrics)

	// create opentracing client (jaeger)
	tr, closer := tracer.Init(serviceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tr)

	// create kafka sync producer
	pr, err := producer.NewProducer(cfg.Broker)
	if err != nil {
		panic(err)
	}
	defer pr.Close()

	// write greeting string
	if err := hellower.SayHelloFrom(os.Stdout, serviceName); err != nil {
		panic(err)
	}

	// start api grpc server
	grpcSrv := grpc.NewServer(grpc.ChainUnaryInterceptor(api.NewInterceptorWithTrace()))

	go handleTermination(grpcSrv)

	if err := api.Run(grpcSrv, cfg.GRPC, dbRepo, fl, counters, pr); err != nil {
		panic(err)
	}
}

func handleTermination(
	grpcSrv *grpc.Server,
) {
	sigCh := initSignalHandler()
	for range sigCh {
		log.Infof("terminate signal received, gracefully terminate")
		grpcSrv.GracefulStop()
		log.Infof("done")
	}
}
