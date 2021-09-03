package ova_joke_api //nolint:revive,stylecheck

import (
	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

//go:generate mockgen -source service.go -destination ./../mocks/service/service.go internal/mocks/service Repo
type Repo interface {
	AddJokes(entities []models.Joke) error
	ListJokes(limit, offset uint64) ([]*models.Joke, error)
	DescribeJoke(jokeID models.JokeID) (*models.Joke, error)
	RemoveJoke(jokeID models.JokeID) error
}

type JokeAPI struct {
	pb.UnimplementedJokeServiceServer
	repo Repo
}

func NewJokeAPI(r Repo) pb.JokeServiceServer {
	return &JokeAPI{
		repo: r,
	}
}

func jokeToPbJoke(j *models.Joke) *pb.Joke {
	return &pb.Joke{
		Id:       j.ID,
		Text:     j.Text,
		AuthorId: j.AuthorID,
	}
}
