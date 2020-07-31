//+build wireinject

package httpclient

import (
	"net/http"

	"github.com/google/wire"
	"golang.org/x/oauth2"
)

func NewClient(
	tokenSource oauth2.TokenSource,
) *http.Client {
	wire.Build(newTransport, newClient)
	return nil
}
