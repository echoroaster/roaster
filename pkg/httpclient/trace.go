package httpclient // import "github.com/echoroaster/roaster/pkg/httpclient"

import (
	"net/http"

	"github.com/echoroaster/roaster/pkg/monitoring"
	"go.opencensus.io/plugin/ochttp"
)

func WrapTraceTransport(base http.RoundTripper) *ochttp.Transport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &ochttp.Transport{
		Base:        base,
		Propagation: monitoring.HTTPFormat,
	}
}
