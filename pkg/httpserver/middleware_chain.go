package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Middleware interface {
	Middleware(next http.Handler) http.Handler
}

type MiddlewareFunc func(next http.Handler) http.Handler

func (mf MiddlewareFunc) Middleware(next http.Handler) http.Handler {
	return mf(next)
}

type MiddlewareChain []Middleware

func (mc MiddlewareChain) Build(mux *mux.Router) (h http.Handler) {
	h = mux

	for i := len(mc) - 1; i >= 0; i -= 1 {
		h = mc[i].Middleware(h)
	}
	return h
}
