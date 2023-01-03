package metrics

import (
	"contrib.go.opencensus.io/exporter/prometheus"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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

func RegisterStackdriverExporter() (*stackdriver.Exporter, error) {
	sdExporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:         "dsp-tools-k8s",
		MetricPrefix:      "sherlock",
		ReportingInterval: 60 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	if err := sdExporter.StartMetricsExporter(); err != nil {
		return nil, fmt.Errorf("error starting stackdriver metrics exporter: %v", err)
	}
	return sdExporter, nil
}
