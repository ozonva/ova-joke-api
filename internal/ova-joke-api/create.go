package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func createJokeRequestToJoke(r *pb.CreateJokeRequest) *models.Joke {
	return models.NewJoke(
		r.GetId(),
		r.GetText(),
		r.GetAuthorId(),
	)
}

// CreateJoke create new joke entity.
func (j *JokeAPI) CreateJoke(ctx context.Context, req *pb.CreateJokeRequest) (*pb.CreateJokeResponse, error) {
	log.Info().Msgf("create: %s", req.String())

	newJoke := createJokeRequestToJoke(req)
	err := j.repo.AddJokes([]models.Joke{*newJoke})
	if err != nil {
		msg := fmt.Sprintf("create joke failed, reason: %v", err)
		log.Error().Msg(msg)
		return nil, status.Error(codes.Internal, msg)
	}

	log.Info().Msgf("created: %v", newJoke)

	// send message to kafka
	_, _, err = j.producer.SendJokeCreatedMsg(ctx, req.GetId())
	if err != nil {
		log.Warn().Msgf("send create joke event failed, reason: %v", err)
	}

	j.metrics.CreateJokeCounterInc()
	return &pb.CreateJokeResponse{}, nil
}
