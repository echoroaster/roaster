package httpclient // import "github.com/echoroaster/roaster/pkg/httpclient"

import (
	"net/http"

	"github.com/echoroaster/roaster/pkg/auth"
	"golang.org/x/oauth2"
)

func newClient(transport http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: transport,
	}
}

func newTransport(
	tokenSource oauth2.TokenSource,
) (t http.RoundTripper) {
	t = http.DefaultTransport
	t = auth.WrapTransport(t, tokenSource)
	t = wrapTraceTransport(t)
	return
}
