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

// RemoveJokeV1 delete joke from storage.
func (j *JokeAPI) RemoveJokeV1(_ context.Context, req *pb.RemoveJokeRequestV1) (*pb.RemoveJokeResponseV1, error) {
	log.Info().Msg(fmt.Sprintf("remove: %s", req.String()))

	stor.mx.Lock()
	defer stor.mx.Unlock()

	if _, ok := stor.data[models.JokeID(req.GetId())]; !ok {
		msg := fmt.Sprintf("joke with id=%d not found", req.Id)
		log.Error().Msg(fmt.Sprintf("remove: %s", msg))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("joke with id=%d not found", req.Id))
	}

	delete(stor.data, models.JokeID(req.GetId()))

	resp := &pb.RemoveJokeResponseV1{}
	log.Info().Msg(fmt.Sprintf("joke with id=%d removed", req.GetId()))
	return resp, nil
}
