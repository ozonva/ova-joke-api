package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/ozonva/ova-joke-api/internal/app/hellower"
	"github.com/ozonva/ova-joke-api/internal/configs"
	"github.com/ozonva/ova-joke-api/internal/flusher"
	"github.com/ozonva/ova-joke-api/internal/metrics"
	api "github.com/ozonva/ova-joke-api/internal/ova-joke-api"
	"github.com/ozonva/ova-joke-api/internal/producer"
	"github.com/ozonva/ova-joke-api/internal/repo"
	"github.com/ozonva/ova-joke-api/internal/tracer"
	desc "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
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

func run(
	_ context.Context,
	config configs.GRPCServerConfig,
	r *repo.JokePgRepo,
	f api.Flusher,
	m *metrics.Metrics,
	pr api.Producer,
) error {
	listen, err := net.Listen("tcp", config.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcSrv := grpc.NewServer(grpc.ChainUnaryInterceptor(newInterceptorWithTrace()))
	desc.RegisterJokeServiceServer(grpcSrv, api.NewJokeAPI(r, f, m, pr))

	if err := grpcSrv.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
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

func main() {
	// configure logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

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
	metricsSrv := metrics.NewServer()
	metricsSrv.Run(cfg.Metrics)

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
	if err := run(context.Background(), cfg.GRPC, dbRepo, fl, counters, pr); err != nil {
		panic(err)
	}
}

// newInterceptorWithTrace wraps all gRPC calls with tracer's span.
func newInterceptorWithTrace() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		trace := opentracing.GlobalTracer()
		span := trace.StartSpan(info.FullMethod)
		span.LogFields(
			log.String("request", fmt.Sprintf("%v", req)),
		)
		defer span.Finish()
		return handler(opentracing.ContextWithSpan(ctx, span), req)
	}
}
