package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/ozonva/ova-joke-api/internal/logger"
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
	log.Infof("create: %s", req.String())

	newJoke := createJokeRequestToJoke(req)
	err := j.repo.AddJokes([]models.Joke{*newJoke})
	if err != nil {
		msg := fmt.Sprintf("create joke failed, reason: %v", err)
		log.Errorf(msg)
		return nil, status.Error(codes.Internal, msg)
	}

	log.Infof("created: %v", newJoke)

	// send message to kafka
	_, _, err = j.producer.SendJokeCreatedMsg(ctx, req.GetId())
	if err != nil {
		log.Warnf("send create joke event failed, reason: %v", err)
	}

	j.metrics.CreateJokeCounterInc()
	return &pb.CreateJokeResponse{}, nil
}
