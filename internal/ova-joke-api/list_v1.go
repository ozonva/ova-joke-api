package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

// ListJokeV1 show list of jokes.
func (j *JokeAPI) ListJokeV1(_ context.Context, req *pb.ListJokeRequestV1) (*pb.ListJokeResponseV1, error) {
	log.Info().Msg(fmt.Sprintf("list: %s", req.String()))

	stor.mx.RLock()
	defer stor.mx.RUnlock()

	resp := &pb.ListJokeResponseV1{}

	resp.Jokes = make([]*pb.Joke, 0, len(stor.data))
	for _, v := range stor.data {
		resp.Jokes = append(resp.Jokes, jokeToListPbJoke(v))
	}
	log.Info().Msg(fmt.Sprintf("list of %d element showed", len(resp.Jokes)))
	return resp, nil
}
