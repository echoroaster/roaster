package auth // import "github.com/echoroaster/roaster/pkg/auth"

import (
	"context"
	"encoding/json"

	"github.com/coreos/go-oidc"
	"go.opencensus.io/trace"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Token struct {
	jwt.Claims
	Roles map[string][]string `json:"roles.omitempty"`
}

func NewRemoteKeySet(ctx context.Context, provider *oidc.Provider) (oidc.KeySet, error) {
	var claim struct {
		JwksUri string `json:"jwks_uri"`
	}
	if err := provider.Claims(&claim); err != nil {
		return nil, err
	}
	return oidc.NewRemoteKeySet(ctx, claim.JwksUri), nil
}

type Validator struct {
	KeySet oidc.KeySet
}

func (v *Validator) Verify(ctx context.Context, token string) (t *Token, err error) {
	ctx, span := trace.StartSpan(ctx, "auth.Validator.Verify")
	defer span.End()
	payload, err := v.KeySet.VerifySignature(ctx, token)
	if err != nil {
		return nil, err
	}
	t = new(Token)
	err = json.Unmarshal(payload, &t)
	return
}
