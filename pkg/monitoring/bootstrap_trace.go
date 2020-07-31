package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	"os"

	"github.com/echoroaster/roaster/pkg/logging"
	"go.opencensus.io/trace"
)

func BootstrapTrace() (cleanup func(), err error) {
	{
		sampler := trace.ProbabilitySampler(0.2)
		if os.Getenv("DEBUG") != "" {
			sampler = trace.AlwaysSample()
		}
		trace.ApplyConfig(trace.Config{
			DefaultSampler: sampler,
		})
	}
	cleanupFunc := make([]func(), 0)
	logger := logging.RootLogger.WithField("component", "trace")

	if os.Getenv("TRACE_AWS_XRAY") != "" {
		exporter, cleanup, err := newXRayExporter()
		if err != nil {
			return nil, err
		}
		cleanupFunc = append(cleanupFunc, cleanup)
		trace.RegisterExporter(exporter)
		logger.Info("register AWS XRay")
	}

	if os.Getenv("JAEGER_AGENT") != "" {
		exporter, cleanup, err := newJaegerExporter()
		if err != nil {
			return nil, err
		}
		cleanupFunc = append(cleanupFunc, cleanup)
		trace.RegisterExporter(exporter)
		logger.Info("register Jaeger")
	}

	if os.Getenv("TRACE_STACKDRIVER") != "" {
		exporter, cleanup, err := newStackdriverTraceExporter()
		if err != nil {
			return nil, err
		}
		cleanupFunc = append(cleanupFunc, cleanup)
		trace.RegisterExporter(exporter)
		logger.Info("register Stackdriver")
	}

	return func() {
		for _, fn := range cleanupFunc {
			fn()
		}
	}, nil
}
