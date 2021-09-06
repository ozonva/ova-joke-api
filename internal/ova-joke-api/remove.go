package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

// RemoveJoke delete joke from storage.
func (j *JokeAPI) RemoveJoke(_ context.Context, req *pb.RemoveJokeRequest) (*pb.RemoveJokeResponse, error) {
	log.Info().Msgf("remove: %s", req.String())

	resp := &pb.RemoveJokeResponse{}

	err := j.repo.RemoveJoke(req.GetId())
	if err != nil {
		msg := fmt.Sprintf("remove failed joke with id=%d, reason: %v", req.Id, err)
		log.Error().Msg(msg)
		return resp, status.Error(codes.Internal, msg)
	}

	log.Info().Msgf("joke with id=%d removed", req.GetId())
	return resp, nil
}
