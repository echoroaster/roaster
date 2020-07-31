package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	"os"
	"sync"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource/gcp"
)

var (
	sdInit     sync.Once
	sdExporter *stackdriver.Exporter
)
var sdOptions = stackdriver.Options{
	MonitoredResource: &gcp.GKEContainer{
		PodID:       os.Getenv("POD_ID"),
		NamespaceID: os.Getenv("NAMESPACE"),
	},
}

func initStackdriverExporter() (err error) {
	sdInit.Do(func() {
		sdExporter, err = stackdriver.NewExporter(stackdriver.Options{
			MonitoredResource: &gcp.GKEContainer{
				PodID:       os.Getenv("POD_ID"),
				NamespaceID: os.Getenv("NAMESPACE"),
			},
		})
	})
	return err
}

func newStackdriverTraceExporter() (*stackdriver.Exporter, func(), error) {
	if err := initStackdriverExporter(); err != nil {
		return nil, nil, err
	}

	return sdExporter, func() {
		sdExporter.Flush()
	}, nil
}

func newStackdriverStatsExporter() (func(), error) {
	if err := initStackdriverExporter(); err != nil {
		return nil, err
	}

	if err := sdExporter.StartMetricsExporter(); err != nil {
		return nil, err
	}

	return func() {
		sdExporter.StopMetricsExporter()
		sdExporter.Flush()
	}, nil
}
