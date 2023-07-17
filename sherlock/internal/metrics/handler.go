package metrics

import (
	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/gin-gonic/gin"
	"log"
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
