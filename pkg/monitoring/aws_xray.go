package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	xray "contrib.go.opencensus.io/exporter/aws"
)

func newXRayExporter() (*xray.Exporter, func(), error) {
	exporter, err := xray.NewExporter(
		xray.WithServiceName(serviceName),
		xray.WithOrigin(xray.OriginECS),
	)
	if err != nil {
		return nil, nil, err
	}
	return exporter, func() {
		exporter.Close()
	}, nil
}
