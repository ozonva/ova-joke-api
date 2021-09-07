package ova_joke_api //nolint:revive,stylecheck

import (
	"context"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

//go:generate mockgen -source service.go -destination ./../mocks/service/service.go internal/mocks/service
type Repo interface {
	AddJokes(entities []models.Joke) error
	ListJokes(limit, offset uint64) ([]models.Joke, error)
	DescribeJoke(jokeID models.JokeID) (*models.Joke, error)
	UpdateJoke(joke models.Joke) error
	RemoveJoke(jokeID models.JokeID) error
	HealthCheckJoke() error
}

// Flusher interface to store jokes into repository.
type Flusher interface {
	Flush(ctx context.Context, entities []models.Joke) []models.Joke
}

type Metrics interface {
	CreateJokeCounterInc()
	MultiCreateJokeCounterInc()
	MultiCreateJokeFailedCounterInc()
	ListJokeCounterInc()
	DescribeJokeCounterInc()
	DescribeJokeNotExistsCounterInc()
	UpdateJokeCounterInc()
	RemoveJokeCounterInc()
}

type Producer interface {
	SendJokeCreatedMsg(ctx context.Context, id models.JokeID) (int32, int64, error)
	SendJokeUpdatedMsg(ctx context.Context, id models.JokeID) (int32, int64, error)
	SendJokeDeletedMsg(ctx context.Context, id models.JokeID) (int32, int64, error)
}

type JokeAPI struct {
	pb.UnimplementedJokeServiceServer
	repo     Repo
	flusher  Flusher
	metrics  Metrics
	producer Producer
}

func NewJokeAPI(r Repo, fl Flusher, m Metrics, pr Producer) pb.JokeServiceServer {
	return &JokeAPI{
		repo:     r,
		flusher:  fl,
		metrics:  m,
		producer: pr,
	}
}

func jokeToPbJoke(j *models.Joke) *pb.Joke {
	return &pb.Joke{
		Id:       j.ID,
		Text:     j.Text,
		AuthorId: j.AuthorID,
	}
}
