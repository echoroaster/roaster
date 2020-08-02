package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

import (
	"net/http"
	"os"

	"github.com/echoroaster/roaster/pkg/auth"
	"github.com/echoroaster/roaster/pkg/logging"
	"github.com/echoroaster/roaster/pkg/monitoring"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/samsarahq/thunder/graphql"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

func init() {
	view.Register(ochttp.ServerRequestCountView,
		ochttp.ServerRequestBytesView,
		ochttp.ServerResponseBytesView,
		ochttp.ServerLatencyView,
		ochttp.ServerRequestCountByMethod,
		ochttp.ServerResponseCountByStatusCode,
	)
}

func MakeHandler(
	m *mux.Router,
	validator *auth.Validator,
) http.Handler {
	return MiddlewareChain{
		MiddlewareFunc(func(next http.Handler) http.Handler {
			return &ochttp.Handler{
				Propagation: monitoring.HTTPFormat,
				Handler:     next,
			}
		}),
		MiddlewareFunc(handlers.ProxyHeaders),
		MiddlewareFunc(handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true),
			handlers.RecoveryLogger(logging.RootLogger.WithField("scope", "panic")),
		)),
		MiddlewareFunc(func(next http.Handler) http.Handler {
			return handlers.CombinedLoggingHandler(os.Stdout, next)
		}),
		corsMiddleware,
		NewAuthMiddleware(validator),
	}.Build(m)
}

func MakeHandlerWithGraphQL(
	m *mux.Router,
	schema *graphql.Schema,
	opt *GraphQLHandlerOptions,
	validator *auth.Validator,
) http.Handler {
	return MakeHandler(addGraphQLEndpointToMux(schema, opt, m), validator)
}
