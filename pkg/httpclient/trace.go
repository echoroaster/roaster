package httpclient // import "github.com/echoroaster/roaster/pkg/httpclient"

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

func wrapTraceTransport(base http.RoundTripper) *ochttp.Transport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &ochttp.Transport{
		Base: base,
	}
}
