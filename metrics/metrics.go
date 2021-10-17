package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func CollectPrometheusMetrics() error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":2112", nil)
}
