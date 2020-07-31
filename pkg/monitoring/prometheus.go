package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	"context"
	"net/http"
	"os"
	"strings"

	"contrib.go.opencensus.io/exporter/prometheus"
)

func newPrometheusExporter() (func(), error) {
	exporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: strings.ReplaceAll(serviceName, "-", "_"),
	})
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", exporter)
	port := os.Getenv("PROMETHEUS_PORT")
	if port == "" {
		port = "8090"
	}
	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	return func() {
		srv.Shutdown(context.Background())
	}, nil
}
