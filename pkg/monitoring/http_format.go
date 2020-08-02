package monitoring // import "github.com/echoroaster/roaster/pkg/monitoring"

import (
	"net/http"

	xray "contrib.go.opencensus.io/exporter/aws"
	jaeger "contrib.go.opencensus.io/exporter/jaeger/propagation"
	stackdriver "contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

var HTTPFormat = CombinedHTTPFormat{
	&b3.HTTPFormat{},
	&jaeger.HTTPFormat{},
	&stackdriver.HTTPFormat{},
	&xray.HTTPFormat{},
}

type CombinedHTTPFormat []propagation.HTTPFormat

func (hf CombinedHTTPFormat) SpanContextFromRequest(req *http.Request) (trace.SpanContext, bool) {
	for _, format := range hf {
		if sc, ok := format.SpanContextFromRequest(req); ok {
			return sc, ok
		}
	}

	return trace.SpanContext{}, false
}

func (hf CombinedHTTPFormat) SpanContextToRequest(sc trace.SpanContext, req *http.Request) {
	for _, format := range hf {
		format.SpanContextToRequest(sc, req)
	}
}
