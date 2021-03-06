package ova_joke_api //nolint:revive,stylecheck

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	tracelog "github.com/opentracing/opentracing-go/log"

	log "github.com/ozonva/ova-joke-api/internal/logger"
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

func failedJokesToMultiCreateJokeResponse(jokes []models.Joke) *pb.MultiCreateJokeResponse {
	responseJokes := make([]*pb.Joke, 0, len(jokes))
	for i := range jokes {
		responseJokes = append(responseJokes, jokeToPbJoke(&jokes[i]))
	}

	return &pb.MultiCreateJokeResponse{
		FailedJokes: responseJokes,
	}
}

// MultiCreateJoke create many new jokes asynchronously.
func (j *JokeAPI) MultiCreateJoke(
	ctx context.Context,
	req *pb.MultiCreateJokeRequest,
) (*pb.MultiCreateJokeResponse, error) {
	log.Infof("multiple create")

	jokes := MultiCreateJokeRequestToJokes(req)

	span, traceCtx := opentracing.StartSpanFromContext(ctx, "multiple_create")
	span.LogFields(
		tracelog.String("batch_size", fmt.Sprintf("%d", len(jokes))),
	)
	defer span.Finish()

	failedJokes := j.flusher.Flush(traceCtx, jokes)
	if len(failedJokes) > 0 {
		log.Warnf("multiple created failed for %d/%d jokes", len(failedJokes), len(jokes))
		j.metrics.MultiCreateJokeFailedCounterInc()
		return failedJokesToMultiCreateJokeResponse(failedJokes), nil
	}

	log.Infof("multiple created %d jokes", len(jokes))
	j.metrics.MultiCreateJokeCounterInc()
	return &pb.MultiCreateJokeResponse{}, nil
}
