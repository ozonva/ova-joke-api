package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

// RemoveJoke delete joke from storage.
func (j *JokeAPI) RemoveJoke(_ context.Context, req *pb.RemoveJokeRequest) (*pb.RemoveJokeResponse, error) {
	log.Info().Msg(fmt.Sprintf("remove: %s", req.String()))

	j.jokes.mx.Lock()
	defer j.jokes.mx.Unlock()

	if _, ok := j.jokes.data[req.GetId()]; !ok {
		msg := fmt.Sprintf("joke with id=%d not found", req.Id)
		log.Warn().Msg(fmt.Sprintf("remove: %s", msg))
		return nil, nil
	}

	delete(j.jokes.data, req.GetId())

	resp := &pb.RemoveJokeResponse{}
	log.Info().Msg(fmt.Sprintf("joke with id=%d removed", req.GetId()))
	return resp, nil
}
