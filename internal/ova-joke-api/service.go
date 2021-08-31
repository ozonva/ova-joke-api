package ova_joke_api //nolint:revive,stylecheck

import (
	"sync"

	"github.com/ozonva/ova-joke-api/internal/domain/author"
	"github.com/ozonva/ova-joke-api/internal/domain/joke"
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
	data map[joke.ID]*joke.Joke
	seq  int64
}

var stor = &storage{
	data: make(map[joke.ID]*joke.Joke),
}

func authorToPbAuthor(a *author.Author) *pb.Author {
	return &pb.Author{
		Id:   int64(a.ID),
		Name: a.Name,
	}
}

func jokeToListPbJoke(j *joke.Joke) *pb.Joke {
	return &pb.Joke{
		Id:     int64(j.ID),
		Text:   j.Text,
		Author: authorToPbAuthor(j.Author),
	}
}

func pbAuthorToAuthor(a *pb.Author) *author.Author {
	return author.New(
		author.ID(a.Id),
		a.Name,
	)
}
