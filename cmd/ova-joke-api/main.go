package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/ozonva/ova-joke-api/internal/app/hellower"
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
	grpcPort string
	dbHost   string

	dbPort uint
	dbName string
	dbUser string
	dbPass string

	metricsAddr string

	brokerAddrsArg string
	brokerAddrs    []string
)

func init() {
	flag.StringVar(&grpcPort, "port", "0.0.0.0:9090", "port for gRPC api server")
	flag.StringVar(&dbHost, "db-host", "localhost", "port for gRPC api server")

	flag.UintVar(&dbPort, "db-port", 5432, "database port")
	flag.StringVar(&dbName, "db-name", "postgres", "database name")
	flag.StringVar(&dbUser, "db-user", "postgres", "database user name")
	flag.StringVar(&dbPass, "db-pass", "postgres", "database users' password")

	flag.StringVar(&metricsAddr, "metrics-addr", "0.0.0.0:9093", "addr of metrics exporter api")

	flag.StringVar(&brokerAddrsArg, "broker-addr", "0.0.0.0:9092", "coma separated list of brokers addrs")
}

func run(_ context.Context, r *repo.JokePgRepo, f api.Flusher, m *metrics.Metrics, pr api.Producer) error {
	listen, err := net.Listen("tcp", grpcPort)
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

func createDB() *sqlx.DB {
	if dbPort > math.MaxUint16 {
		panic(fmt.Sprintf("invalid dbConn port given %d, must be compatible with uint16", dbPort))
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	dbConn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	return dbConn
}

func createRepo(db *sqlx.DB) *repo.JokePgRepo {
	return repo.NewJokePgRepo(db)
}

func createJokeFlusher(r *repo.JokePgRepo) *flusher.JokeFlusher {
	flusherCap := 3
	return flusher.NewJokeFlusher(flusherCap, r)
}

func main() {
	// parse cli arguments
	flag.Parse()
	brokerAddrs = strings.Split(brokerAddrsArg, ",")

	// configure logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// create database client
	db := createDB()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(fmt.Sprintf("unable to close db connection %v", err))
		}
	}()

	// create repository (using postgres connection) and flusher (batch wrapper)
	dbRepo := createRepo(db)
	fl := createJokeFlusher(dbRepo)

	// create metric storage and run handler for prometheus
	counters := metrics.NewMetrics()
	metricsSrv := metrics.NewServer()
	metricsSrv.Run(metricsAddr)

	// create opentracing client (jaeger)
	tr, closer := tracer.Init(serviceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tr)

	// create kafka sync producer
	pr, err := producer.NewProducer(brokerAddrs)
	if err != nil {
		panic(err)
	}
	defer pr.Close()

	// write greeting string
	if err := hellower.SayHelloFrom(os.Stdout, serviceName); err != nil {
		panic(err)
	}

	// start api grpc server
	if err := run(context.Background(), dbRepo, fl, counters, pr); err != nil {
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
