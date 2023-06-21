package metrics

import (
	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/gin-gonic/gin"
	"log"
)

// RegisterPrometheusMetricsHandler accepts a gin.RouterGroup and will set up
// a prometheus metrics endpoint on the provided route
func RegisterPrometheusMetricsHandler(metricsGroup *gin.RouterGroup) {
	prometheusExporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "sherlock",
	})
	if err != nil {
		log.Fatalf("error creating prometheus exporter")
	}

	metricsGroup.GET("", gin.WrapH(prometheusExporter))
}
