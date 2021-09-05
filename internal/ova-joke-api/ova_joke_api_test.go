package ova_joke_api_test

import (
	"context"
	"database/sql"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"

	mocks "github.com/ozonva/ova-joke-api/internal/mocks/service"
	"github.com/ozonva/ova-joke-api/internal/models"
	ova_joke_api "github.com/ozonva/ova-joke-api/internal/ova-joke-api"
	pb "github.com/ozonva/ova-joke-api/pkg/ova-joke-api"
)

var errTestService = errors.New("some-service-error")

var _ = Describe("OvaJokeApi", func() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	var (
		ctrl         *gomock.Controller
		mockRepo     *mocks.MockRepo
		mockFlusher  *mocks.MockFlusher
		mockMetrics  *mocks.MockMetrics
		mockProducer *mocks.MockProducer

		srv    pb.JokeServiceServer
		ctx    context.Context
		jokes  []*models.Joke
		jokeID uint64
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		mockFlusher = mocks.NewMockFlusher(ctrl)
		mockMetrics = mocks.NewMockMetrics(ctrl)
		mockProducer = mocks.NewMockProducer(ctrl)

		srv = ova_joke_api.NewJokeAPI(mockRepo, mockFlusher, mockMetrics, mockProducer)
		ctx = context.TODO()
		jokes = []models.Joke{{ID: 3, Text: "joke #3", AuthorID: 33}}
		jokeID = 3
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Create joke", func() {
		It("successfully done", func() {
			mockRepo.EXPECT().AddJokes([]models.Joke{jokes[0]}).Times(1)
			mockProducer.EXPECT().SendJokeCreatedMsg(ctx, jokes[0].ID).Times(1)
			mockMetrics.EXPECT().CreateJokeCounterInc().Times(1)

			resp, err := srv.CreateJoke(ctx, &pb.CreateJokeRequest{
				Id:       jokes[0].ID,
				Text:     jokes[0].Text,
				AuthorId: jokes[0].AuthorID,
			})

			Expect(err).Should(Succeed())
			Expect(resp).To(BeEquivalentTo(&pb.CreateJokeResponse{}))
		})

		It("failed", func() {
			mockRepo.EXPECT().AddJokes([]models.Joke{jokes[0]}).Return(errTestService).Times(1)

			resp, err := srv.CreateJoke(ctx, &pb.CreateJokeRequest{
				Id:       jokes[0].ID,
				Text:     jokes[0].Text,
				AuthorId: jokes[0].AuthorID,
			})

			Expect(err).ShouldNot(Succeed())
			Expect(resp).To(BeNil())
		})
	})

	Context("Describe joke", func() {
		It("successfully done", func() {
			mockRepo.EXPECT().DescribeJoke(jokeID).Return(&jokes[0], nil).Times(1)
			mockMetrics.EXPECT().DescribeJokeCounterInc().Times(1)

			resp, err := srv.DescribeJoke(ctx, &pb.DescribeJokeRequest{
				Id: jokeID,
			})

			Expect(err).Should(Succeed())
			Expect(resp.Id).To(BeEquivalentTo(jokes[0].ID))
			Expect(resp.Text).To(BeEquivalentTo(jokes[0].Text))
			Expect(resp.AuthorId).To(BeEquivalentTo(jokes[0].AuthorID))
		})

		It("joke not exists", func() {
			mockRepo.EXPECT().DescribeJoke(jokeID).Return(nil, sql.ErrNoRows).Times(1)

			resp, err := srv.DescribeJoke(ctx, &pb.DescribeJokeRequest{
				Id: jokeID,
			})

			Expect(err).ShouldNot(Succeed())
			Expect(resp).To(BeNil())
		})

		It("failed", func() {
			mockRepo.EXPECT().DescribeJoke(jokeID).Return(nil, errTestService).Times(1)

			resp, err := srv.DescribeJoke(ctx, &pb.DescribeJokeRequest{
				Id: jokeID,
			})

			Expect(err).ShouldNot(Succeed())
			Expect(resp).To(BeNil())
		})
	})

	Context("List jokes", func() {
		BeforeEach(func() {
			jokes = []models.Joke{
				{ID: 3, Text: "joke #3", AuthorID: 33},
				{ID: 4, Text: "joke #4", AuthorID: 44},
				{ID: 5, Text: "joke #5", AuthorID: 55},
			}
		})
		It("successfully done", func() {
			mockRepo.EXPECT().ListJokes(uint64(3), uint64(5)).Return(jokes, nil).Times(1)
			mockMetrics.EXPECT().ListJokeCounterInc().Times(1)

			resp, err := srv.ListJoke(ctx, &pb.ListJokeRequest{
				Limit:  uint64(3),
				Offset: uint64(5),
			})

			Expect(err).Should(Succeed())

			for i := range resp.Jokes {
				Expect(resp.Jokes[i].Id).To(BeEquivalentTo(jokes[i].ID))
				Expect(resp.Jokes[i].Text).To(BeEquivalentTo(jokes[i].Text))
				Expect(resp.Jokes[i].AuthorId).To(BeEquivalentTo(jokes[i].AuthorID))
			}
		})

		It("failed", func() {
			mockRepo.EXPECT().ListJokes(uint64(3), uint64(5)).Return(nil, errTestService).Times(1)

			resp, err := srv.ListJoke(ctx, &pb.ListJokeRequest{
				Limit:  uint64(3),
				Offset: uint64(5),
			})

			Expect(err).ShouldNot(Succeed())
			Expect(resp).To(BeNil())
		})
	})

	Context("Remove jokes", func() {
		It("successfully done", func() {
			mockRepo.EXPECT().RemoveJoke(models.JokeID(3)).Return(nil).Times(1)
			mockProducer.EXPECT().SendJokeDeletedMsg(ctx, jokes[0].ID).Times(1)
			mockMetrics.EXPECT().RemoveJokeCounterInc().Times(1)

			resp, err := srv.RemoveJoke(ctx, &pb.RemoveJokeRequest{
				Id: jokeID,
			})

			Expect(err).Should(Succeed())
			Expect(resp).To(BeEquivalentTo(&pb.RemoveJokeResponse{}))
		})

		It("failed", func() {
			mockRepo.EXPECT().RemoveJoke(models.JokeID(3)).Return(errTestService).Times(1)

			resp, err := srv.RemoveJoke(ctx, &pb.RemoveJokeRequest{
				Id: jokeID,
			})

			Expect(err).ShouldNot(Succeed())
			Expect(resp).To(BeEquivalentTo(&pb.RemoveJokeResponse{}))
		})
	})

	Context("Update joke", func() {
		It("successfully done", func() {
			mockRepo.EXPECT().UpdateJoke(*jokes[0]).Return(nil).Times(1)
			mockProducer.EXPECT().SendJokeUpdatedMsg(ctx, jokes[0].ID).Times(1)
			mockMetrics.EXPECT().UpdateJokeCounterInc().Times(1)

			resp, err := srv.UpdateJoke(ctx, &pb.UpdateJokeRequest{
				Id:       jokes[0].ID,
				Text:     jokes[0].Text,
				AuthorId: jokes[0].AuthorID,
			})

			Expect(err).Should(Succeed())
			Expect(resp).To(BeEquivalentTo(&pb.UpdateJokeResponse{}))
		})

		It("failed", func() {
			mockRepo.EXPECT().UpdateJoke(*jokes[0]).Return(errTestService).Times(1)

			resp, err := srv.UpdateJoke(ctx, &pb.UpdateJokeRequest{
				Id:       jokes[0].ID,
				Text:     jokes[0].Text,
				AuthorId: jokes[0].AuthorID,
			})

			Expect(err).ShouldNot(Succeed())
			Expect(resp).To(BeEquivalentTo(&pb.UpdateJokeResponse{}))
		})
	})
})
