package flusher

import (
	"github.com/ozonva/ova-joke-api/internal/domain/joke"
	"github.com/ozonva/ova-joke-api/internal/repo"
	"github.com/ozonva/ova-joke-api/internal/utils"
)

// Flusher interface to store jokes into repository.
type Flusher interface {
	Flush(entities []joke.Joke) []joke.Joke
}

// NewFlusher returns Flusher with bulk persist support.
func NewFlusher(chunkSize int, entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (f flusher) Flush(entities []joke.Joke) []joke.Joke {
	var failed []joke.Joke
	for _, chunk := range utils.SplitToBulks(entities, f.chunkSize) {
		if err := f.entityRepo.AddEntities(chunk); err != nil {
			failed = append(failed, chunk...)
		}
	}

	return failed
}
