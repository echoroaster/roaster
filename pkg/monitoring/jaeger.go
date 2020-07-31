package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
)

func newJaegerExporter() (*jaeger.Exporter, func(), error) {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: os.Getenv("JAEGER_AGENT"),
		Process: jaeger.Process{
			ServiceName: serviceName,
		},
	})

	if err != nil {
		return nil, nil, err
	}

	return exporter, func() {
		exporter.Flush()
	}, nil
}
