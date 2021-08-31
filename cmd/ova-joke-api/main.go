package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozonva/ova-joke-api/internal/app/hellower"
	api "github.com/ozonva/ova-joke-api/internal/ova-joke-api"
	desc "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

const serviceName = "ova-joke-api"

var grpcPort string

func init() {
	flag.StringVar(&grpcPort, "port", "0.0.0.0:9090", "port for gRPC api server")
}

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	desc.RegisterJokeServiceServer(s, api.NewJokeAPI())

	if err := s.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	flag.Parse()

	if err := hellower.SayHelloFrom(os.Stdout, serviceName); err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
