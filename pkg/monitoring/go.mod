module github.com/echoroaster/roaster/pkg/monitoring

go 1.14

replace github.com/echoroaster/roaster/pkg/logging => ../logging

require (
	contrib.go.opencensus.io/exporter/aws v0.0.0-20200617204711-c478e41e60e9
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/prometheus v0.2.0
	contrib.go.opencensus.io/exporter/stackdriver v0.13.2
	github.com/echoroaster/roaster/pkg/logging v0.0.0-00010101000000-000000000000
	go.opencensus.io v0.22.4
)
