package repo

import (
	"github.com/ozonva/ova-joke-api/internal/domain/joke"
)

//go:generate mockgen -source ./internal/repo/repo.go -destination ./internal/repo/generated/repomock.go Repo
// Repo interface to store Jokes.
type Repo interface {
	AddEntities(entities []joke.Joke) error
	ListEntities(limit, offset uint64) ([]Repo, error)
	DescribeEntity(jokeID joke.ID) (*Repo, error)
}
