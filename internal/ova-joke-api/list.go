package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

// ListJoke show list of jokes.
func (j *JokeAPI) ListJoke(_ context.Context, req *pb.ListJokeRequest) (*pb.ListJokeResponse, error) {
	log.Info().Msgf("list: %s", req.String())

	jokes, err := j.repo.ListJokes(req.GetLimit(), req.GetOffset())
	if err != nil {
		msg := fmt.Sprintf("show list failed: %v", err)
		log.Error().Msg(msg)
		return nil, status.Error(codes.Internal, msg)
	}

	respJokes := make([]*pb.Joke, 0, len(jokes))
	for i := range jokes {
		respJokes = append(respJokes, jokeToPbJoke(&jokes[i]))
	}

	resp := &pb.ListJokeResponse{
		Jokes: respJokes,
	}

	log.Info().Msgf("list of %d element showed", len(resp.Jokes))
	j.metrics.ListJokeCounterInc()
	return resp, nil
}
