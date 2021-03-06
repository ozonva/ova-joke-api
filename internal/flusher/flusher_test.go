package flusher_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"

	"github.com/ozonva/ova-joke-api/internal/flusher"
	mock "github.com/ozonva/ova-joke-api/internal/mocks/flusher"
	"github.com/ozonva/ova-joke-api/internal/models"
)

func makeJokeCollection(sz int) []models.Joke {
	jokes := make([]models.Joke, 0, sz)
	for i := 1; i < sz+1; i++ {
		jokes = append(jokes, *models.NewJoke(models.JokeID(i), "joke#"+strconv.Itoa(i), 1))
	}

	return jokes
}

var _ = Describe("When Flusher", func() {
	var (
		ctrl                    *gomock.Controller
		jokes                   []models.Joke
		repo                    *mock.MockRepo
		fl                      *flusher.JokeFlusher
		errTestingRepoAddFailed = fmt.Errorf("failed Repo.AddJokes()")
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		repo = mock.NewMockRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("calls Flush()", func() {
		When("JokeFlusher sz is > 0", func() {
			BeforeEach(func() {
				fl = flusher.NewJokeFlusher(5, repo)
			})

			When("passed collection is empty", func() {
				BeforeEach(func() {
					jokes = []models.Joke{}
				})

				It("not call Repo.AddJokes() at all", func() {
					repo.EXPECT().AddJokes(jokes).Times(0)

					Expect(fl.Flush(context.TODO(), jokes)).To(BeNil())
				})
			})

			When("passed collection is smaller then JokeFlusher sz", func() {
				BeforeEach(func() {
					jokes = makeJokeCollection(3)
				})

				It("call Repo.AddJokes() once and return nil", func() {
					repo.EXPECT().AddJokes(jokes).Return(nil)
					Expect(fl.Flush(context.TODO(), jokes)).To(BeNil())
				})

				It("call Repo.AddJokes() once and fail", func() {
					repo.EXPECT().AddJokes(jokes).Return(errTestingRepoAddFailed)
					Expect(fl.Flush(context.TODO(), jokes)).To(Equal(jokes))
				})
			})

			When("passed collection is larger then JokeFlusher sz", func() {
				BeforeEach(func() {
					jokes = makeJokeCollection(7)
				})

				It("call Repo.AddJokes() twice and return nil", func() {
					repo.EXPECT().AddJokes(jokes[:5]).Return(nil)
					repo.EXPECT().AddJokes(jokes[5:]).Return(nil)

					Expect(fl.Flush(context.TODO(), jokes)).To(BeNil())
				})

				It("call Repo.AddJokes() twice and fail first", func() {
					repo.EXPECT().AddJokes(jokes[:5]).Return(errTestingRepoAddFailed)
					repo.EXPECT().AddJokes(jokes[5:]).Return(nil)

					Expect(fl.Flush(context.TODO(), jokes)).To(Equal(jokes[:5]))
				})

				It("call Repo.AddJokes() twice and fail second", func() {
					repo.EXPECT().AddJokes(jokes[:5]).Return(nil)
					repo.EXPECT().AddJokes(jokes[5:]).Return(errTestingRepoAddFailed)

					Expect(fl.Flush(context.TODO(), jokes)).To(Equal(jokes[5:]))
				})

				It("call Repo.AddJokes() twice and fail both", func() {
					repo.EXPECT().AddJokes(jokes[:5]).Return(errTestingRepoAddFailed)
					repo.EXPECT().AddJokes(jokes[5:]).Return(errTestingRepoAddFailed)

					Expect(fl.Flush(context.TODO(), jokes)).To(Equal(jokes))
				})
			})
		})

		When("JokeFlusher sz is zero", func() {
			BeforeEach(func() {
				fl = flusher.NewJokeFlusher(0, repo)
			})

			When("passed collection is empty", func() {
				BeforeEach(func() {
					jokes = []models.Joke{}
				})

				It("not call Repo.AddJokes() at all", func() {
					repo.EXPECT().AddJokes(jokes).Times(0)

					Expect(fl.Flush(context.TODO(), jokes)).To(BeNil())
				})
			})

			When("passed collection is larger then JokeFlusher sz", func() {
				BeforeEach(func() {
					jokes = makeJokeCollection(3)
				})

				It("call Repo.AddJokes() once and return nil", func() {
					repo.EXPECT().AddJokes(jokes).Return(nil)

					Expect(fl.Flush(context.TODO(), jokes)).To(BeNil())
				})

				It("call Repo.AddJokes() once and fail first", func() {
					repo.EXPECT().AddJokes(jokes).Return(errTestingRepoAddFailed)

					Expect(fl.Flush(context.TODO(), jokes)).To(Equal(jokes))
				})
			})
		})

		When("JokeFlusher sz is negative", func() {
			BeforeEach(func() {
				fl = flusher.NewJokeFlusher(-2, repo)
			})

			When("passed collection is empty", func() {
				BeforeEach(func() {
					jokes = []models.Joke{}
				})

				It("not call Repo.AddJokes() at all", func() {
					repo.EXPECT().AddJokes(jokes).Times(0)

					Expect(fl.Flush(context.TODO(), jokes)).To(BeNil())
				})
			})

			When("passed collection is larger then JokeFlusher sz", func() {
				BeforeEach(func() {
					jokes = makeJokeCollection(3)
				})

				It("call Repo.AddJokes() once and return nil", func() {
					repo.EXPECT().AddJokes(jokes).Return(nil)

					Expect(fl.Flush(context.TODO(), jokes)).To(BeNil())
				})

				It("call Repo.AddJokes() once and fail first", func() {
					repo.EXPECT().AddJokes(jokes).Return(errTestingRepoAddFailed)

					Expect(fl.Flush(context.TODO(), jokes)).To(Equal(jokes))
				})
			})
		})
	})
})
