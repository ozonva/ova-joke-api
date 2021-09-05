package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	createCounter      prometheus.Counter
	multiCreateCounter prometheus.Counter
	listCounter        prometheus.Counter
	describeCounter    prometheus.Counter
	updateCounter      prometheus.Counter
	removeCounter      prometheus.Counter
	totalCounter       prometheus.Counter
}

func (m Metrics) CreateJokeCounterInc() {
	m.createCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) MultiCreateJokeCounterInc() {
	m.multiCreateCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) ListJokeCounterInc() {
	m.listCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) DescribeJokeCounterInc() {
	m.describeCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) UpdateJokeCounterInc() {
	m.updateCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) RemoveJokeCounterInc() {
	m.removeCounter.Inc()
	m.totalCounter.Inc()
}

func NewMetrics() *Metrics {
	m := &Metrics{
		createCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "joke_api_create_total",
			Help: "The total number create requests in joke api",
		}),
		multiCreateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "joke_api_multi_create_total",
			Help: "The total number multi create requests in joke api",
		}),
		listCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "joke_api_list_total",
			Help: "The total number list requests in joke api",
		}),
		describeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "joke_api_describe_total",
			Help: "The total number describe requests in joke api",
		}),
		updateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "joke_api_update_total",
			Help: "The total number update requests in joke api",
		}),
		removeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "joke_api_remove_total",
			Help: "The total number remove requests in joke api",
		}),
		totalCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "joke_api_total",
			Help: "The total number requests in joke api",
		}),
	}

	return m
}
