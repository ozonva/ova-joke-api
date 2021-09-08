package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/ozonva/ova-joke-api/internal/logger"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

// RemoveJoke delete joke from storage.
func (j *JokeAPI) RemoveJoke(ctx context.Context, req *pb.RemoveJokeRequest) (*pb.RemoveJokeResponse, error) {
	log.Infof("remove: %s", req.String())

	resp := &pb.RemoveJokeResponse{}

	err := j.repo.RemoveJoke(req.GetId())
	if err != nil {
		msg := fmt.Sprintf("remove failed joke with id=%d, reason: %v", req.Id, err)
		log.Errorf(msg)
		return resp, status.Error(codes.Internal, msg)
	}

	// send message to kafka
	_, _, err = j.producer.SendJokeDeletedMsg(ctx, req.GetId())
	if err != nil {
		log.Warnf("send remove joke event failed, reason: %v", err)
	}

	log.Infof("joke with id=%d removed", req.GetId())
	j.metrics.RemoveJokeCounterInc()
	return resp, nil
}
