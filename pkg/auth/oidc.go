package auth // import "github.com/echoroaster/roaster/pkg/auth"

import (
	"context"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func NewProvider(ctx context.Context) (*oidc.Provider, error) {
	return oidc.NewProvider(ctx, os.Getenv("OIDC_ISSUER"))
}

func NewTokenSource(ctx context.Context, provider *oidc.Provider) oauth2.TokenSource {
	endpoint := provider.Endpoint()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("OIDC_CLIENT_ID"),
		ClientSecret: os.Getenv("OIDC_CLIENT_SECRET"),
		TokenURL:     endpoint.TokenURL,
		AuthStyle:    endpoint.AuthStyle,
	}
	return config.TokenSource(context.Background())
}
