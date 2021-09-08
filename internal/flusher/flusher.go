package flusher

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/ozonva/ova-joke-api/internal/models"
	"github.com/ozonva/ova-joke-api/internal/utils"
)

// Repo to persist entities on Flush.
type Repo interface {
	AddJokes(entities []models.Joke) error
}

type JokeFlusher struct {
	chunkSize  int
	entityRepo Repo
}

// NewJokeFlusher returns Flusher with bulk persist support.
func NewJokeFlusher(chunkSize int, entityRepo Repo) *JokeFlusher {
	return &JokeFlusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

func (f JokeFlusher) Flush(ctx context.Context, entities []models.Joke) []models.Joke {
	var failed []models.Joke

	for _, chunk := range utils.SplitToBulks(entities, f.chunkSize) {
		span, _ := opentracing.StartSpanFromContext(ctx, "flush_joke_chunk")
		span.LogFields(
			log.String("flush-batch-size", fmt.Sprintf("%d", len(chunk))),
		)
		if err := f.entityRepo.AddJokes(chunk); err != nil {
			failed = append(failed, chunk...)
		}
		span.Finish()
	}

	return failed
}
