package ova_joke_api //nolint:revive,stylecheck

import (
	"context"

	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

// HealthCheckJoke verify gRPC server's status.
func (j *JokeAPI) HealthCheckJoke(_ context.Context, _ *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	if err := j.repo.HealthCheckJoke(); err != nil {
		//nolint:nilerr // when repository (database) error it's not an error for health check call.
		return &pb.HealthCheckResponse{
			Grpc:     1,
			Database: 0,
		}, nil
	}

	return &pb.HealthCheckResponse{
		Grpc:     1,
		Database: 1,
	}, nil
}
