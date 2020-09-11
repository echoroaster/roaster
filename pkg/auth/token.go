package auth // import "github.com/echoroaster/roaster/pkg/auth"

import (
	"context"
	"encoding/json"

	"github.com/coreos/go-oidc"
	"go.opencensus.io/trace"
	"gopkg.in/square/go-jose.v2/jwt"
)

type KeycloakToken struct {
	jwt.Claims
	AuthorizedParty string                    `json:"azp,omitempty"`
	RealmAccess     *ResourceAccess           `json:"realm_access,omitempty"`
	ResourceAccess  map[string]ResourceAccess `json:"resource_access,omitempty"`
	ClientId        string                    `json:"clientId,omitempty"`
	Scope           string                    `json:"scope,omitempty"`
}

type ResourceAccess struct {
	Roles []string `json:"roles"`
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

func (v *Validator) Verify(ctx context.Context, token string) (t *KeycloakToken, err error) {
	ctx, span := trace.StartSpan(ctx, "auth.Validator.Verify")
	defer span.End()
	payload, err := v.KeySet.VerifySignature(ctx, token)
	if err != nil {
		return nil, err
	}
	t = new(KeycloakToken)
	err = json.Unmarshal(payload, t)
	return
}
