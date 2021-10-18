package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func CollectPrometheusMetrics() error {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":2112",
		Handler: mux,
	}

	mux.Handle("/metrics", promhttp.Handler())
	return srv.ListenAndServe()
}
