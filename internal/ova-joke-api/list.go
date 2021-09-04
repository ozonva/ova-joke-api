package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

// ListJoke show list of jokes.
func (j *JokeAPI) ListJoke(_ context.Context, req *pb.ListJokeRequest) (*pb.ListJokeResponse, error) {
	log.Info().Msg(fmt.Sprintf("list: %s", req.String()))

	j.jokes.mx.RLock()
	defer j.jokes.mx.RUnlock()

	resp := &pb.ListJokeResponse{}

	resp.Jokes = make([]*pb.Joke, 0, len(j.jokes.data))
	for _, v := range j.jokes.data {
		resp.Jokes = append(resp.Jokes, jokeToPbJoke(v))
	}
	log.Info().Msg(fmt.Sprintf("list of %d element showed", len(resp.Jokes)))
	return resp, nil
}
