package metrics

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"github.com/ozonva/ova-joke-api/internal/configs"
)

func Run(config configs.MetricsServerConfig) *http.Server {
	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    config.Addr,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	return srv
}
