package saver

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/ozonva/ova-joke-api/internal/domain/joke"
)

type MockFlusher struct {
	FlushCnt int32
}

func (f *MockFlusher) Flush([]joke.Joke) []joke.Joke {
	atomic.AddInt32(&f.FlushCnt, 1)
	return nil
}

//nolint:funlen
func TestServer(t *testing.T) {
	defer goleak.VerifyNone(t)

	clc := clock.NewMock()

	t.Run("Invalid cap", func(t *testing.T) {
		fl := &MockFlusher{}
		_, err := NewSaver(context.TODO(), 0, fl, clc.Ticker(200*time.Microsecond))
		require.ErrorIs(t, err, ErrInvalidArgument)
	})

	t.Run("no goroutine leaks on close", func(t *testing.T) {
		fl := &MockFlusher{}
		s, _ := NewSaver(context.TODO(), 42, fl, clc.Ticker(200*time.Microsecond))
		s.Save(joke.Joke{})
		s.Close()
	})

	t.Run("when Step() not called, tickers not running", func(t *testing.T) {
		fl := &MockFlusher{}
		s, _ := NewSaver(context.TODO(), 42, fl, clc.Ticker(200*time.Microsecond))
		defer s.Close()

		before := runtime.NumGoroutine()
		s.Save(joke.Joke{})
		after := runtime.NumGoroutine()
		require.Less(t, before, after)
	})

	t.Run("correct closes on ctx close", func(t *testing.T) {
		newCtx, cancel := context.WithCancel(context.TODO())
		fl := &MockFlusher{}
		s, _ := NewSaver(newCtx, 2, fl, clc.Ticker(200*time.Microsecond))

		defer s.Close()

		for i := 0; i < 10; i++ {
			s.Save(joke.Joke{})
		}

		// flush calls at least once
		require.Greater(t, fl.FlushCnt, int32(1))

		cancel()
	})

	t.Run("concurrent save", func(t *testing.T) {
		fl := &MockFlusher{}
		s, _ := NewSaver(context.TODO(), 20, fl, clc.Ticker(200*time.Microsecond))

		defer s.Close()

		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 100; i++ {
					s.Save(joke.Joke{})
				}
			}()
		}
		wg.Wait()

		// 10 * 100 = 1000 elements at all
		// cap = 20 => 1000 / 20 = 50 flushes + 1 on close in defer
		require.Equal(t, int32(50), fl.FlushCnt)
	})
}
