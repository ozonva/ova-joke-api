package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ozonva/ova-joke-api/internal/configs"
)

type MetricServer struct{}

func NewServer() *MetricServer {
	return &MetricServer{}
}

func (srv *MetricServer) Run(config configs.MetricsServerConfig) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(config.Addr, nil)
		if err != nil {
			panic(err)
		}
	}()
}
