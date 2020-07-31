package auth // import "github.com/echoroaster/roaster/pkg/auth"

import (
	"net/http"

	"golang.org/x/oauth2"
)

func WrapTransport(base http.RoundTripper, source oauth2.TokenSource) *oauth2.Transport {
	return &oauth2.Transport{
		Source: source,
		Base:   base,
	}
}
