package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozonva/ova-joke-api/internal/app/hellower"
	api "github.com/ozonva/ova-joke-api/internal/ova-joke-api"
	"github.com/ozonva/ova-joke-api/internal/repo"
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
)

func init() {
	flag.StringVar(&grpcPort, "port", "0.0.0.0:9090", "port for gRPC api server")
	flag.StringVar(&dbHost, "db-host", "localhost", "port for gRPC api server")

	flag.UintVar(&dbPort, "db-port", 5432, "database port")
	flag.StringVar(&dbName, "db-name", "postgres", "database name")
	flag.StringVar(&dbUser, "db-user", "postgres", "database user name")
	flag.StringVar(&dbPass, "db-pass", "postgres", "database users' password")
}

func run(dbConn *sqlx.DB) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	desc.RegisterJokeServiceServer(s, api.NewJokeAPI(repo.NewJokePgRepo(dbConn)))

	if err := s.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	flag.Parse()

	if dbPort > math.MaxUint16 {
		log.Fatal().Msg(fmt.Sprintf("invalid dbConn port given %d, must be compatible with uint16", dbPort))
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	dbConn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			panic(fmt.Sprintf("unable to close dbConn connection %v", err))
		}
	}()

	if err := hellower.SayHelloFrom(os.Stdout, serviceName); err != nil {
		panic(err)
	}

	if err := run(dbConn); err != nil {
		panic(err)
	}
}
