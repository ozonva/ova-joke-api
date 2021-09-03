package flusher

import (
	"github.com/ozonva/ova-joke-api/internal/models"
	"github.com/ozonva/ova-joke-api/internal/utils"
)

//go:generate mockgen -source flusher.go -destination ./../mocks/flusher/flusher.go internal/mocks/flusher Flusher,Repo
// Flusher interface to store jokes into repository.
type Flusher interface {
	Flush(entities []models.Joke) []models.Joke
}

// Repo to persist entities on Flush.
type Repo interface {
	AddJokes(entities []models.Joke) error
}

type flusher struct {
	chunkSize  int
	entityRepo Repo
}

// NewFlusher returns Flusher with bulk persist support.
func NewFlusher(chunkSize int, entityRepo Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

func (f flusher) Flush(entities []models.Joke) []models.Joke {
	var failed []models.Joke
	for _, chunk := range utils.SplitToBulks(entities, f.chunkSize) {
		if err := f.entityRepo.AddJokes(chunk); err != nil {
			failed = append(failed, chunk...)
		}
	}

	return failed
}
