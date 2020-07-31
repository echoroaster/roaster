package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

import (
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
)

var corsMiddleware Middleware

func init() {
	options := []handlers.CORSOption{
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Accept", "Accept-Language", "Content-Language", "Origin"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete}),
	}
	if allowedOrigin := os.Getenv("CORS_ALLOWED_ORIGIN"); allowedOrigin != "" {
		options = append(options, handlers.AllowedOrigins(strings.Split(allowedOrigin, ",")))
	}
	corsMiddleware = MiddlewareFunc(handlers.CORS(options...))
}
