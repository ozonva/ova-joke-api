package saver

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/benbjohnson/clock"

	"github.com/ozonva/ova-joke-api/internal/flusher"
	"github.com/ozonva/ova-joke-api/internal/models"
)

var ErrInvalidArgument = errors.New("invalid argument")

type JokeSaver struct {
	mx     sync.Mutex
	fl     flusher.Flusher
	buffer []models.Joke

	tickerCtxCancel context.CancelFunc
	tickerWg        sync.WaitGroup
}

func (s *JokeSaver) Save(entity models.Joke) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.buffer = append(s.buffer, entity)
	if len(s.buffer) == cap(s.buffer) {
		s.flushNoLock()
	}
}

func (s *JokeSaver) run(ctx context.Context, dur time.Duration) {
	tickerCtx, tickerCtxCancel := context.WithCancel(ctx)
	s.tickerCtxCancel = tickerCtxCancel

	ticker := getClock().Ticker(dur)

	s.tickerWg.Add(1)
	go func() {
		defer s.tickerWg.Done()
		defer s.flush()
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.flush()
			case <-tickerCtx.Done():
				return
			}
		}
	}()
}

func (s *JokeSaver) Close() {
	s.tickerCtxCancel()
	s.tickerWg.Wait()
}

func (s *JokeSaver) flush() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if len(s.buffer) > 0 {
		s.flushNoLock()
	}
}

func (s *JokeSaver) flushNoLock() {
	failed := s.fl.Flush(s.buffer)
	s.buffer = s.buffer[:0]
	s.buffer = append(s.buffer, failed...)
}

//nolint:gocritic // not to replace on clock.New(), because it redefined in tests.
var getClock = func() clock.Clock {
	return clock.New()
}

// NewSaver returns saver with periodic data persist.
func NewSaver(ctx context.Context, capacity uint, flusher flusher.Flusher, dur time.Duration) *JokeSaver {
	if capacity < 1 {
		panic(fmt.Errorf("invalid buffer capacity: %d, %w", capacity, ErrInvalidArgument))
	}

	s := &JokeSaver{
		fl:     flusher,
		buffer: make([]models.Joke, 0, capacity),
	}

	s.run(ctx, dur)

	return s
}
