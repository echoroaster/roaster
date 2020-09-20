package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

import (
	"net/http"

	"github.com/echoroaster/roaster/pkg/logging"

	"github.com/echoroaster/roaster/pkg/auth"
	"github.com/gorilla/mux"
	"github.com/samsarahq/thunder/graphql"
)

func MakeHandler(
	m *mux.Router,
	validator *auth.Validator,
	logger logging.Logger,
) http.Handler {
	return append(
		NewDefaultMiddlewareChain(logger),
		NewAuthMiddleware(validator),
	).Build(m)
}

func MakeHandlerWithGraphQL(
	m *mux.Router,
	schema *graphql.Schema,
	opt *GraphQLHandlerOptions,
	validator *auth.Validator,
	logger logging.Logger,
) http.Handler {
	m = addGraphQLEndpointToMux(schema, opt, m)
	return MakeHandler(
		m,
		validator,
		logger,
	)
}
