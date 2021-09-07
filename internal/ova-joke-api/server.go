package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	traceLog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"net"
	"sync"

	"github.com/ozonva/ova-joke-api/internal/configs"
	log "github.com/ozonva/ova-joke-api/internal/logger"
	"github.com/ozonva/ova-joke-api/internal/metrics"
	"github.com/ozonva/ova-joke-api/internal/repo"
	desc "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func Run(
	wg *sync.WaitGroup,
	config configs.GRPCServerConfig,
	r *repo.JokePgRepo,
	f Flusher,
	m *metrics.Metrics,
	pr Producer,
) (*grpc.Server, error) {
	listen, err := net.Listen("tcp", config.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	grpcSrv := grpc.NewServer(grpc.ChainUnaryInterceptor(newInterceptorWithTrace()))
	desc.RegisterJokeServiceServer(grpcSrv, NewJokeAPI(r, f, m, pr))

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Infof("start listen gRPC API on %s", config.Addr)
		if err := grpcSrv.Serve(listen); err != nil {
			panic(fmt.Errorf("failed to serve: %w", err))
		}
	}()

	return grpcSrv, nil
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
			traceLog.String("request", fmt.Sprintf("%v", req)),
		)
		defer span.Finish()
		return handler(opentracing.ContextWithSpan(ctx, span), req)
	}
}
