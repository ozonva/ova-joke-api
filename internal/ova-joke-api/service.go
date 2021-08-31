package ova_joke_api //nolint:revive,stylecheck

import (
	"sync"

	"github.com/ozonva/ova-joke-api/internal/models"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

type JokeAPI struct {
	pb.UnimplementedJokeServiceServer
}

func NewJokeAPI() pb.JokeServiceServer {
	return &JokeAPI{}
}

type storage struct {
	mx   sync.RWMutex
	data map[models.JokeID]*models.Joke
	seq  int64
}

var stor = &storage{
	data: make(map[models.JokeID]*models.Joke),
}

func authorToPbAuthor(a *models.Author) *pb.Author {
	return &pb.Author{
		Id:   int64(a.ID),
		Name: a.Name,
	}
}

func jokeToListPbJoke(j *models.Joke) *pb.Joke {
	return &pb.Joke{
		Id:     int64(j.ID),
		Text:   j.Text,
		Author: authorToPbAuthor(j.Author),
	}
}

func pbAuthorToAuthor(a *pb.Author) *models.Author {
	return models.NewAuthor(
		models.AuthorID(a.Id),
		a.Name,
	)
}
