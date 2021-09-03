package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func createJokeRequestToJoke(r *pb.CreateJokeRequest, id models.JokeID) *models.Joke {
	return models.NewJoke(
		id,
		r.Text,
		r.GetAuthorId(),
	)
}

func jokeToCreateJokeResponse(j *models.Joke) *pb.CreateJokeResponse {
	return &pb.CreateJokeResponse{
		Id: j.ID,
	}
}

// CreateJoke create new joke entity.
func (j *JokeAPI) CreateJoke(_ context.Context, req *pb.CreateJokeRequest) (*pb.CreateJokeResponse, error) {
	log.Info().Msg(fmt.Sprintf("create: %s", req.String()))

	j.jokes.mx.Lock()
	defer j.jokes.mx.Unlock()

	j.jokes.seq++
	newJoke := createJokeRequestToJoke(req, models.JokeID(j.jokes.seq))
	j.jokes.data[newJoke.ID] = newJoke

	resp := jokeToCreateJokeResponse(newJoke)
	log.Info().Msg(fmt.Sprintf("created: %s", resp.String()))
	return resp, nil
}
