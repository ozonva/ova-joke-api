package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonva/ova-joke-api/internal/domain/joke"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func jokeToDescribeJokeResponseV1(j *joke.Joke) *pb.DescribeJokeResponseV1 {
	return &pb.DescribeJokeResponseV1{
		Id:     int64(j.ID),
		Text:   j.Text,
		Author: authorToPbAuthor(j.Author),
	}
}

// DescribeJokeV1 show full information about Joke entity.
func (j *JokeAPI) DescribeJokeV1(_ context.Context, req *pb.DescribeJokeRequestV1) (*pb.DescribeJokeResponseV1, error) {
	log.Info().Msg(fmt.Sprintf("describe: %s", req.String()))

	stor.mx.RLock()
	defer stor.mx.RUnlock()

	jk, ok := stor.data[joke.ID(req.GetId())]

	if !ok {
		msg := fmt.Sprintf("joke with id=%d not found", req.Id)
		log.Error().Msg(fmt.Sprintf("describe: %s", msg))
		return nil, status.Error(codes.NotFound, msg)
	}

	resp := jokeToDescribeJokeResponseV1(jk)
	log.Info().Msg(fmt.Sprintf("described: %s", resp.String()))
	return resp, nil
}
