package main

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//// BuildVersion is intended for use with Go's LDFlags compiler option, to
//// set this value at compile time
//var BuildVersion = "development"
//
//func init() {
//	version.BuildVersion = BuildVersion
//}

func main() {
	// We'll handle SIGINT and SIGTERM ourselves
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	application := &boot.Application{}
	go application.Start()
	<-ctx.Done()

	// Our Kubernetes infra typically sleeps app containers for 10 seconds before propagating the SIGTERM upon the pod
	// terminating. Sherlock is built with a distro-less image, though, so there's no /bin/sleep present. Instead,
	// we just catch the signal ourselves and wait the 10 seconds before shutting down.
	// We gate this behavior on the presence of KUBERNETES_SERVICE_HOST because that's an easy hack to detect if we're
	// running on Kubernetes (so we exit quickly locally)
	if _, present := os.LookupEnv("KUBERNETES_SERVICE_HOST"); present {
		for seconds := 10; seconds > 0; seconds-- {
			log.Info().Msgf("MAIN | running on kubernetes, waiting %d seconds before shutting down...", seconds)
			time.Sleep(time.Second)
		}
	}
	application.Stop()
}
