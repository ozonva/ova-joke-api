package saver

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/benbjohnson/clock"

	"github.com/ozonva/ova-joke-api/internal/domain/joke"
	"github.com/ozonva/ova-joke-api/internal/flusher"
)

var _ Saver = (*jokeSaver)(nil)

var ErrInvalidArgument = errors.New("invalid argument")

type Saver interface {
	Save(entity joke.Joke)
	Close()
}

type jokeSaver struct {
	mx     *sync.Mutex
	fl     flusher.Flusher
	buffer []joke.Joke
	ctx    context.Context

	once            sync.Once
	ticker          *clock.Ticker
	tickerCtxCancel context.CancelFunc
	tickerWg        *sync.WaitGroup
}

func (s *jokeSaver) Save(entity joke.Joke) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.once.Do(s.run)

	s.buffer = append(s.buffer, entity)
	if len(s.buffer) == cap(s.buffer) {
		s.flushNoLock()
	}
}

func (s *jokeSaver) run() {
	tickerCtx, tickerCtxCancel := context.WithCancel(s.ctx)

	s.tickerCtxCancel = tickerCtxCancel
	s.tickerWg.Add(1)
	go func() {
		defer s.tickerWg.Done()
		defer s.flush()
		defer s.ticker.Stop()

		for {
			select {
			case <-s.ticker.C:
				s.flush()
			case <-tickerCtx.Done():
				return
			}
		}
	}()
}

func (s *jokeSaver) Close() {
	// we can run Close without any Save call and ticker can be not initialised
	if s.tickerCtxCancel != nil {
		s.tickerCtxCancel()
	}

	s.tickerWg.Wait()
}

func (s *jokeSaver) flush() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if len(s.buffer) > 0 {
		s.flushNoLock()
	}
}

func (s *jokeSaver) flushNoLock() {
	failed := s.fl.Flush(s.buffer)
	s.buffer = s.buffer[:0]
	s.buffer = append(s.buffer, failed...)
}

// NewSaver returns Saver with periodic data persist.
func NewSaver(ctx context.Context, capacity uint, flusher flusher.Flusher, ticker *clock.Ticker) (Saver, error) {
	if capacity < 1 {
		return nil, fmt.Errorf("invalid buffer capacity: %d, %w", capacity, ErrInvalidArgument)
	}

	s := &jokeSaver{
		mx:       &sync.Mutex{},
		fl:       flusher,
		buffer:   make([]joke.Joke, 0, capacity),
		ctx:      ctx,
		ticker:   ticker,
		tickerWg: &sync.WaitGroup{},
	}

	return s, nil
}
