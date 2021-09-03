package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"database/sql"
	"errors"
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
	log.Info().Msgf("describe: %s", req.String())

	joke, err := j.repo.DescribeJoke(req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			msg := fmt.Sprintf("joke with id=%d not found", req.GetId())
			log.Warn().Msg(msg)
			return nil, status.Error(codes.NotFound, msg)
		}
		msg := fmt.Sprintf("described failed: %v", err)
		log.Error().Msg(msg)
		return nil, status.Error(codes.Internal, msg)
	}

	if joke == nil {
		msg := fmt.Sprintf("joke with id=%d not found", req.GetId())
		log.Warn().Msg(msg)
		return nil, status.Error(codes.NotFound, msg)
	}

	resp := jokeToDescribeJokeResponse(joke)
	log.Info().Msgf("described: %s", resp.String())
	return resp, nil
}
