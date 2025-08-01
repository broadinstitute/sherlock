package metrics

import (
	"log"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/gin-gonic/gin"
)

// PrometheusHandler exposes registered Prometheus metrics
func PrometheusHandler() gin.HandlerFunc {
	prometheusExporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "sherlock",
	})
	if err != nil {
		log.Fatalf("error creating prometheus exporter")
	}

	return gin.WrapH(prometheusExporter)
}
