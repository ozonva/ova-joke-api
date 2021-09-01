package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func createJokeRequestV1ToJoke(r *pb.CreateJokeRequestV1, id models.JokeID) *models.Joke {
	return models.NewJoke(
		id,
		r.Text,
		r.GetAuthorId(),
	)
}

func jokeToCreateJokeResponseV1(j *models.Joke) *pb.CreateJokeResponseV1 {
	return &pb.CreateJokeResponseV1{
		Id: j.ID,
	}
}

// CreateJokeV1 create new joke entity.
func (j *JokeAPI) CreateJokeV1(_ context.Context, req *pb.CreateJokeRequestV1) (*pb.CreateJokeResponseV1, error) {
	log.Info().Msg(fmt.Sprintf("create: %s", req.String()))

	j.jokes.mx.Lock()
	defer j.jokes.mx.Unlock()

	j.jokes.seq++
	newJoke := createJokeRequestV1ToJoke(req, models.JokeID(j.jokes.seq))
	j.jokes.data[newJoke.ID] = newJoke

	resp := jokeToCreateJokeResponseV1(newJoke)
	log.Info().Msg(fmt.Sprintf("created: %s", resp.String()))
	return resp, nil
}
