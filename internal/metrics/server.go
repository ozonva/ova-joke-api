package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricServer struct{}

func NewServer() *MetricServer {
	return &MetricServer{}
}

func (srv *MetricServer) Run(addr string) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			panic(err)
		}
	}()
}
