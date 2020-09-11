package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/echoroaster/roaster/pkg/auth"
	"github.com/echoroaster/roaster/pkg/common"
	"go.opencensus.io/trace"
)

var TokenCtxKey common.ContextKey

func NewAuthMiddleware(validator *auth.Validator) Middleware {
	return MiddlewareFunc(func(next http.Handler) http.Handler {
		return &authHandler{Validator: validator, Next: next}
	})
}

type authHandler struct {
	Validator *auth.Validator
	Next      http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "middleware.Auth")
	defer span.End()
	token, validate := h.validateToken(ctx, w, r)
	if !validate {
		return
	}
	h.Next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), TokenCtxKey, token)))
}

func (h *authHandler) validateToken(ctx context.Context, w http.ResponseWriter, r *http.Request) (token *auth.KeycloakToken, validate bool) {
	headerValue := r.Header.Get("Authorization")
	if headerValue == "" {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Error: ErrorMessage{
				Code:    "UNAUTHORIZATION",
				Message: "authorization is required",
			},
		})
		return
	}

	pieces := strings.SplitN(headerValue, " ", 2)
	if len(pieces) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Error: ErrorMessage{
				Code:    "UNAUTHORIZATION",
				Message: "invalid authorization format",
			},
		})
		return
	}

	if strings.ToLower(pieces[0]) != "bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Error: ErrorMessage{
				Code:    "UNAUTHORIZATION",
				Message: "invalid authorization type",
			},
		})
		return
	}

	token, err := h.Validator.Verify(ctx, pieces[1])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Error: ErrorMessage{
				Code:    "UNAUTHORIZATION",
				Message: "token verify fail",
			},
		})
		return
	}

	return token, true
}
