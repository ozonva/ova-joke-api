package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	tracelog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

func MultiCreateJokeRequestToJokes(req *pb.MultiCreateJokeRequest) []models.Joke {
	jokes := make([]models.Joke, 0, len(req.Jokes))
	for _, joke := range req.Jokes {
		jokes = append(jokes, models.Joke{
			ID:       joke.Id,
			Text:     joke.Text,
			AuthorID: joke.AuthorId,
		})
	}

	return jokes
}

// MultiCreateJoke create many new jokes asynchronously.
func (j *JokeAPI) MultiCreateJoke(
	ctx context.Context,
	req *pb.MultiCreateJokeRequest,
) (*pb.MultiCreateJokeResponse, error) {
	log.Info().Msg("multiple create")

	jokes := MultiCreateJokeRequestToJokes(req)

	span, _ := opentracing.StartSpanFromContext(ctx, "multiple_create")
	span.LogFields(
		tracelog.String("batch_size", fmt.Sprintf("%d", len(jokes))),
	)
	defer span.Finish()

	j.flusher.Flush(opentracing.ContextWithSpan(ctx, span), jokes)

	log.Info().Msg(fmt.Sprintf("multiple created %d jokes", len(jokes)))

	j.metrics.MultiCreateJokeCounterInc()
	return &pb.MultiCreateJokeResponse{}, nil
}
