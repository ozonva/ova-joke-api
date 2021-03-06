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

func UpdateJokeRequestToJoke(r *pb.UpdateJokeRequest) *models.Joke {
	return models.NewJoke(
		r.GetId(),
		r.GetText(),
		r.GetAuthorId(),
	)
}

// UpdateJoke update models.Joke with given id.
func (j *JokeAPI) UpdateJoke(ctx context.Context, req *pb.UpdateJokeRequest) (*pb.UpdateJokeResponse, error) {
	log.Infof("update: %s", req.String())

	resp := &pb.UpdateJokeResponse{}
	joke := UpdateJokeRequestToJoke(req)
	err := j.repo.UpdateJoke(*joke)
	if err != nil {
		msg := fmt.Sprintf("remove failed joke with id=%d, reason: %v", req.Id, err)
		log.Errorf(msg)
		return resp, status.Error(codes.Internal, msg)
	}

	// send message to kafka
	_, _, err = j.producer.SendJokeUpdatedMsg(ctx, req.GetId())
	if err != nil {
		log.Warnf("send update joke event failed, reason: %v", err)
	}

	log.Infof("joke %s updated", req.String())
	j.metrics.UpdateJokeCounterInc()
	return resp, nil
}
