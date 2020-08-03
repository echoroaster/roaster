module github.com/echoroaster/roaster/pkg/cloud/aws

go 1.14

replace (
	github.com/echoroaster/roaster/pkg/auth => ../../auth
	github.com/echoroaster/roaster/pkg/httpclient => ../../httpclient
	github.com/echoroaster/roaster/pkg/logging => ../../logging
)

require (
	cloud.google.com/go v0.62.0 // indirect
	github.com/aws/aws-sdk-go v1.33.17
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/echoroaster/roaster/pkg/httpclient v0.0.0-20200802182826-62af7de36742
	github.com/echoroaster/roaster/pkg/logging v0.0.0-20200802182826-62af7de36742 // indirect
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/prometheus/statsd_exporter v0.17.0 // indirect
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	google.golang.org/grpc v1.31.0 // indirect
)
