package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	"net/http"

	xray "contrib.go.opencensus.io/exporter/aws"
	jaeger "contrib.go.opencensus.io/exporter/jaeger/propagation"
	stackdriver "contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

var HttpFormat = httpFormat{
	"X-Amzn-Trace-Id":       new(xray.HTTPFormat),
	"uber-trace-id":         new(jaeger.HTTPFormat),
	"X-Cloud-Trace-Context": new(stackdriver.HTTPFormat),
}

type httpFormat map[string]propagation.HTTPFormat

func (f httpFormat) SpanContextFromRequest(req *http.Request) (sc trace.SpanContext, ok bool) {
	for header, format := range f {
		if req.Header.Get(header) != "" {
			return format.SpanContextFromRequest(req)
		}
	}
	return trace.SpanContext{}, false
}

func (f httpFormat) SpanContextToRequest(sc trace.SpanContext, req *http.Request) {
	for _, format := range f {
		format.SpanContextToRequest(sc, req)
	}
}
