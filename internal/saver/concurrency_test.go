package saver

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/ozonva/ova-joke-api/internal/models"
)

type MockFlusher struct {
	FlushCnt int32
}

func (f *MockFlusher) Flush(_ context.Context, _ []models.Joke) []models.Joke {
	atomic.AddInt32(&(f.FlushCnt), 1)
	return nil
}

func TestServer(t *testing.T) {
	defer goleak.VerifyNone(t)

	defaultDuration := 200 * time.Millisecond

	t.Run("Invalid cap", func(t *testing.T) {
		fl := &MockFlusher{}

		require.Panics(t, func() {
			_ = NewSaver(context.TODO(), 0, fl, defaultDuration)
		})
	})

	t.Run("no goroutine leaks on close", func(t *testing.T) {
		fl := &MockFlusher{}
		s := NewSaver(context.TODO(), 42, fl, defaultDuration)
		s.Save(models.Joke{})
		s.Close()
	})

	t.Run("correct closes on ctx close", func(t *testing.T) {
		newCtx, cancel := context.WithCancel(context.TODO())
		fl := &MockFlusher{}
		s := NewSaver(newCtx, 2, fl, defaultDuration)

		defer s.Close()

		for i := 0; i < 10; i++ {
			s.Save(models.Joke{})
		}

		// flush calls at least once
		require.Greater(t, fl.FlushCnt, int32(1))

		cancel()
	})

	t.Run("concurrent save", func(t *testing.T) {
		fl := &MockFlusher{}
		s := NewSaver(context.TODO(), 20, fl, defaultDuration)
		defer s.Close()

		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 100; i++ {
					s.Save(models.Joke{})
				}
			}()
		}
		wg.Wait()

		// 10 * 100 = 1000 elements at all
		// cap = 20 => 1000 / 20 = 50 flushes + 1 on close in defer + N because ticker starts on NewSaver
		require.GreaterOrEqual(t, fl.FlushCnt, int32(50))
	})
}
