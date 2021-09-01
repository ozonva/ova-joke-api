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

func jokeToDescribeJokeResponseV1(j *models.Joke) *pb.DescribeJokeResponseV1 {
	return &pb.DescribeJokeResponseV1{
		Id:       j.ID,
		Text:     j.Text,
		AuthorId: j.AuthorID,
	}
}

// DescribeJokeV1 show full information about Joke entity.
func (j *JokeAPI) DescribeJokeV1(_ context.Context, req *pb.DescribeJokeRequestV1) (*pb.DescribeJokeResponseV1, error) {
	log.Info().Msg(fmt.Sprintf("describe: %s", req.String()))

	j.jokes.mx.RLock()
	defer j.jokes.mx.RUnlock()

	jk, ok := j.jokes.data[req.GetId()]

	if !ok {
		msg := fmt.Sprintf("joke with id=%d not found", req.Id)
		log.Error().Msg(fmt.Sprintf("describe: %s", msg))
		return nil, status.Error(codes.NotFound, msg)
	}

	resp := jokeToDescribeJokeResponseV1(jk)
	log.Info().Msg(fmt.Sprintf("described: %s", resp.String()))
	return resp, nil
}
