package ova_joke_api //nolint:revive,stylecheck

import (
	"sync"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

type JokeAPI struct {
	pb.UnimplementedJokeServiceServer
	jokes *jokeStorage
}

func NewJokeAPI() pb.JokeServiceServer {
	return &JokeAPI{
		jokes: &jokeStorage{
			data: make(map[models.JokeID]*models.Joke),
		},
	}
}

type jokeStorage struct {
	mx   sync.RWMutex
	data map[models.JokeID]*models.Joke
	seq  int64
}

func jokeToPbJoke(j *models.Joke) *pb.Joke {
	return &pb.Joke{
		Id:       j.ID,
		Text:     j.Text,
		AuthorId: j.AuthorID,
	}
}
