package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/ozonva/ova-joke-api/internal/logger"
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
	log.Infof("describe: %s", req.String())

	joke, err := j.repo.DescribeJoke(req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			msg := fmt.Sprintf("joke with id=%d not found", req.GetId())
			log.Warnf(msg)
			j.metrics.DescribeJokeNotExistsCounterInc()
			return nil, status.Error(codes.NotFound, msg)
		}
		msg := fmt.Sprintf("described failed: %v", err)
		log.Errorf(msg)
		return nil, status.Error(codes.Internal, msg)
	}

	resp := jokeToDescribeJokeResponse(joke)
	log.Infof("described: %s", resp.String())
	j.metrics.DescribeJokeCounterInc()
	return resp, nil
}
