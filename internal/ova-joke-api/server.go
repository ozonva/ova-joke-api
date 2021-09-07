package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	traceLog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"net"

	"github.com/ozonva/ova-joke-api/internal/configs"
	log "github.com/ozonva/ova-joke-api/internal/logger"
	"github.com/ozonva/ova-joke-api/internal/metrics"
	"github.com/ozonva/ova-joke-api/internal/repo"
	desc "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func Run(
	grpcSrv *grpc.Server,
	config configs.GRPCServerConfig,
	r *repo.JokePgRepo,
	f Flusher,
	m *metrics.Metrics,
	pr Producer,
) error {
	listen, err := net.Listen("tcp", config.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	desc.RegisterJokeServiceServer(grpcSrv, NewJokeAPI(r, f, m, pr))

	log.Infof("start listen gRPC API on %s", config.Addr)
	if err := grpcSrv.Serve(listen); err != nil {
		panic(fmt.Errorf("failed to serve: %w", err))
	}

	return nil
}

// NewInterceptorWithTrace wraps all gRPC calls with tracer's span.
func NewInterceptorWithTrace() grpc.UnaryServerInterceptor {
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
