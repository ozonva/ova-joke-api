package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-joke-api/internal/domain/joke"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func createJokeRequestV1ToJoke(r *pb.CreateJokeRequestV1, id joke.ID) *joke.Joke {
	return joke.New(
		id,
		r.Text,
		pbAuthorToAuthor(r.GetAuthor()),
	)
}

func jokeToCreateJokeResponseV1(j *joke.Joke) *pb.CreateJokeResponseV1 {
	return &pb.CreateJokeResponseV1{
		Id: int64(j.ID),
	}
}

// CreateJokeV1 create new joke entity.
func (j *JokeAPI) CreateJokeV1(_ context.Context, req *pb.CreateJokeRequestV1) (*pb.CreateJokeResponseV1, error) {
	log.Info().Msg(fmt.Sprintf("create: %s", req.String()))

	stor.mx.Lock()
	defer stor.mx.Unlock()

	stor.seq++
	newJoke := createJokeRequestV1ToJoke(req, joke.ID(stor.seq))
	stor.data[newJoke.ID] = newJoke

	resp := jokeToCreateJokeResponseV1(newJoke)
	log.Info().Msg(fmt.Sprintf("created: %s", resp.String()))
	return resp, nil
}
