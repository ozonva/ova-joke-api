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

func jokeToDescribeJokeResponse(j *models.Joke) *pb.DescribeJokeResponse {
	return &pb.DescribeJokeResponse{
		Id:       j.ID,
		Text:     j.Text,
		AuthorId: j.AuthorID,
	}
}

// DescribeJoke show full information about Joke entity.
func (j *JokeAPI) DescribeJoke(_ context.Context, req *pb.DescribeJokeRequest) (*pb.DescribeJokeResponse, error) {
	log.Info().Msg(fmt.Sprintf("describe: %s", req.String()))

	j.jokes.mx.RLock()
	defer j.jokes.mx.RUnlock()

	jk, ok := j.jokes.data[req.GetId()]

	if !ok {
		msg := fmt.Sprintf("joke with id=%d not found", req.Id)
		log.Warn().Msg(fmt.Sprintf("describe: %s", msg))
		return nil, status.Error(codes.NotFound, msg)
	}

	resp := jokeToDescribeJokeResponse(jk)
	log.Info().Msg(fmt.Sprintf("described: %s", resp.String()))
	return resp, nil
}
