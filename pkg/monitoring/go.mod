module github.com/echoroaster/roaster/pkg/monitoring

go 1.14

replace github.com/echoroaster/roaster/pkg/logging => ../logging

require (
	cloud.google.com/go v0.62.0 // indirect
	contrib.go.opencensus.io/exporter/aws v0.0.0-20200617204711-c478e41e60e9
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/prometheus v0.2.0
	contrib.go.opencensus.io/exporter/stackdriver v0.13.2
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/echoroaster/roaster/pkg/logging v0.0.0-20200802182826-62af7de36742
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/prometheus/statsd_exporter v0.17.0 // indirect
	go.opencensus.io v0.22.4
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	google.golang.org/grpc v1.31.0 // indirect
)
