package saver

import (
	"context"
	"runtime"
	"strconv"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozonva/ova-joke-api/internal/domain/author"
	"github.com/ozonva/ova-joke-api/internal/domain/joke"
	mock_flusher "github.com/ozonva/ova-joke-api/internal/flusher/generated"
)

func makeJokeCollection(sz int) []joke.Joke {
	a := &author.Author{
		ID:   12,
		Name: "L.Tolstoy",
	}
	jokes := make([]joke.Joke, 0, sz)
	for i := 0; i < sz; i++ {
		jokes = append(jokes, *joke.New(joke.ID(i+1), "joke#"+strconv.Itoa(i+1), a))
	}

	return jokes
}

func callSaveOnJokes(s *JokeSaver, jokes []joke.Joke) {
	for _, jj := range jokes {
		s.Save(jj)
	}
}

// Return internal buffer of JokeSaver for assert.
// Race detector prevent concurrent access and fail tests asserts.
func getBuffer(s *JokeSaver) []joke.Joke {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.buffer
}

// Set internal saver buffer to zero length.
// Race detector prevent concurrent access and fail tests asertation.
func resetBuffer(s *JokeSaver) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.buffer = s.buffer[:0]
}

var _ = Describe("Saver", func() {
	jokeFixtures := makeJokeCollection(10)
	var (
		clc *clock.Mock

		ctrl        *gomock.Controller
		mockFlusher *mock_flusher.MockFlusher

		jSaver *JokeSaver
		jokes  []joke.Joke
	)

	BeforeEach(func() {
		clc = clock.NewMock()
		getClock = func() clock.Clock {
			return clc
		}

		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mock_flusher.NewMockFlusher(ctrl)

		jSaver = NewSaver(context.TODO(), 3, mockFlusher, 200*time.Millisecond)
		jokes = jokeFixtures[:0]
	})

	AfterEach(func() {
		ctrl.Finish()

		// prevent extra flushing on mock when close
		resetBuffer(jSaver)
		jSaver.Close()
	})

	Context("calls Save() on Saver", func() {
		Context("saver's buffer sz + 1 < cap", func() {
			When("shouldn't call flusher at all", func() {
				BeforeEach(func() {
					jokes = jokeFixtures[:1]
					mockFlusher.EXPECT().Flush(jokes).Times(0)
				})

				It("expect one new element in Saver's buffer", func() {
					callSaveOnJokes(jSaver, jokes)
					Expect(getBuffer(jSaver)).To(Equal(jokeFixtures[:1]))
				})
			})
		})

		When("saver's buffer sz + 1 == cap", func() {
			Context("saver flush buffer to prevent overflow", func() {
				When("Flusher success", func() {
					BeforeEach(func() {
						jokes = jokeFixtures[:3]
						mockFlusher.EXPECT().Flush(jokes).Return(nil).Times(1)
					})

					It("expect no elements in Saver's buffer", func() {
						callSaveOnJokes(jSaver, jokes)
						Expect(getBuffer(jSaver)).To(Equal(jokeFixtures[:0]))
					})
				})

				When("Flusher failed", func() {
					BeforeEach(func() {
						jokes = jokeFixtures[:3]
						mockFlusher.EXPECT().Flush(jokes).Return(jokes[1:2]).Times(1)
					})

					It("failed elements stay in Saver's buffer", func() {
						callSaveOnJokes(jSaver, jokes)
						Expect(getBuffer(jSaver)).To(Equal(jokes[1:2]))
					})
				})
			})
		})
	})

	Context("ticker run flush", func() {
		Context("when buffer is empty", func() {
			When("shouldn't call flusher at all", func() {
				BeforeEach(func() {
					jokes = jokeFixtures[:0]
					mockFlusher.EXPECT().Flush(jokes).Times(0)
				})

				It("expect buffer still empty", func() {
					// init run, but clear buffer
					jSaver.Save(joke.Joke{})
					resetBuffer(jSaver)

					callSaveOnJokes(jSaver, jokes)
					// it must trigger flush by ticker once
					clc.Add(300 * time.Millisecond)
					runtime.Gosched()
					Expect(getBuffer(jSaver)).To(Equal(jokeFixtures[:0]))
				})
			})
		})

		Context("saver's buffer sz + 1 < cap", func() {
			Context("saver flush buffer not full buffer", func() {
				When("Flusher success", func() {
					BeforeEach(func() {
						jokes = jokeFixtures[:2]
						mockFlusher.EXPECT().Flush(jokes).Return(nil).Times(1)
					})

					It("expect no elements in Saver's buffer", func() {
						callSaveOnJokes(jSaver, jokes)
						// it must trigger flush by ticker once
						clc.Add(300 * time.Millisecond)
						runtime.Gosched()
						Expect(getBuffer(jSaver)).To(Equal(jokeFixtures[:0]))
					})
				})

				When("Flusher failed", func() {
					BeforeEach(func() {
						jokes = jokeFixtures[:2]
						mockFlusher.EXPECT().Flush(jokes).Return(jokes[1:2]).Times(1)
					})

					It("failed elements stay in Saver's buffer", func() {
						callSaveOnJokes(jSaver, jokes)
						// it must trigger flush by ticker once
						clc.Add(300 * time.Millisecond)
						runtime.Gosched()
						Expect(getBuffer(jSaver)).To(Equal(jokes[1:2]))
					})
				})
			})
		})
	})

	Context("calls Close() on Saver", func() {
		Context("when buffer is empty", func() {
			When("shouldn't call flusher at all", func() {
				BeforeEach(func() {
					jokes = jokeFixtures[:0]
					mockFlusher.EXPECT().Flush(jokes).Times(0)
				})

				It("expect buffer still empty", func() {
					callSaveOnJokes(jSaver, jokes)
					jSaver.Close()
					Expect(getBuffer(jSaver)).To(Equal(jokeFixtures[:0]))
				})
			})
		})

		Context("saver's buffer sz + 1 < cap", func() {
			Context("saver flush buffer not full buffer", func() {
				When("Flusher success", func() {
					BeforeEach(func() {
						jokes = jokeFixtures[:2]
						mockFlusher.EXPECT().Flush(jokes).Return(nil).Times(1)
					})

					It("expect no elements in Saver's buffer", func() {
						callSaveOnJokes(jSaver, jokes)
						jSaver.Close()
						Expect(getBuffer(jSaver)).To(Equal(jokeFixtures[:0]))
					})
				})

				When("Flusher failed", func() {
					BeforeEach(func() {
						jokes = jokeFixtures[:2]
						mockFlusher.EXPECT().Flush(jokes).Return(jokes[1:2]).Times(1)
					})

					It("failed elements stay in Saver's buffer", func() {
						callSaveOnJokes(jSaver, jokes)
						jSaver.Close()
						Expect(getBuffer(jSaver)).To(Equal(jokes[1:2]))
					})
				})
			})
		})
	})
})
