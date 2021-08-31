package repo

import (
	"github.com/ozonva/ova-joke-api/internal/models"
)

//go:generate mockgen -source repo.go -package=mocks -destination ./../mocks/repo.go Repo
// Repo interface to store Jokes.
type Repo interface {
	AddEntities(entities []models.Joke) error
	ListEntities(limit, offset uint64) ([]Repo, error)
	DescribeEntity(jokeID models.JokeID) (*Repo, error)
}
