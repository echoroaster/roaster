package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	"os"

	"github.com/echoroaster/roaster/pkg/logging"
)

func BootstrapStats() (func(), error) {
	cleanupFunc := make([]func(), 0)
	logger := logging.RootLogger.WithField("component", "stats")

	if os.Getenv("STATS_STACKDRIVER") != "" {
		cleanup, err := newStackdriverStatsExporter()
		if err != nil {
			return nil, err
		}
		cleanupFunc = append(cleanupFunc, cleanup)
		logger.Info("register Stackdriver")
	}

	if os.Getenv("STATS_PROMETHEUS") != "" {
		cleanup, err := newPrometheusExporter()
		if err != nil {
			return nil, err
		}
		cleanupFunc = append(cleanupFunc, cleanup)
		logger.Info("register Prometheus")
	}

	return func() {
		for _, fn := range cleanupFunc {
			fn()
		}
	}, nil
}
