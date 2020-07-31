package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

import (
	"encoding/json"
	"net/http"

	"github.com/echoroaster/roaster/pkg/logging"
	"github.com/gorilla/mux"
	"github.com/samsarahq/thunder/graphql"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
)

const introspectionQuery = `
query IntrospectionQuery {
	__schema {
		queryType { name }
		mutationType { name }
		types {
			...FullType
		}
		directives {
			name
			description
			locations
			args {
				...InputValue
			}
		}
	}
}
fragment FullType on __Type {
	kind
	name
	description
	fields(includeDeprecated: true) {
		name
		description
		args {
			...InputValue
		}
		type {
			...TypeRef
		}
		isDeprecated
		deprecationReason
	}
	inputFields {
		...InputValue
	}
	interfaces {
		...TypeRef
	}
	enumValues(includeDeprecated: true) {
		name
		description
		isDeprecated
		deprecationReason
	}
	possibleTypes {
		...TypeRef
	}
}
fragment InputValue on __InputValue {
	name
	description
	type { ...TypeRef }
	defaultValue
}
fragment TypeRef on __Type {
	kind
	name
	ofType {
		kind
		name
		ofType {
			kind
			name
			ofType {
				kind
				name
				ofType {
					kind
					name
					ofType {
						kind
						name
						ofType {
							kind
							name
							ofType {
								kind
								name
							}
						}
					}
				}
			}
		}
	}
}`

type GraphQLHandlerOptions struct {
	Path           string
	EnableGraphIQL bool
}

func addGraphQLEndpointToMux(schema *graphql.Schema, opt *GraphQLHandlerOptions, m *mux.Router) *mux.Router {
	if opt == nil {
		opt = &GraphQLHandlerOptions{}
	}

	if opt.Path == "" {
		opt.Path = "/graphql"
	}

	graphqlHandler := graphql.HTTPHandler(schema)
	m.Handle(opt.Path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "graphql.Handler")
		defer span.End()
		graphqlHandler.ServeHTTP(w, r.WithContext(ctx))
	}))

	m.Handle(opt.Path+"/schema.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "graphql.Schema")
		defer span.End()
		query, err := graphql.Parse(introspectionQuery, map[string]interface{}{})
		encoder := json.NewEncoder(w)
		logger := logging.RootLogger.WithFields(logrus.Fields{
			"request":   trace.FromContext(r.Context()).SpanContext().TraceID.String(),
			"component": "graphql",
			"action":    "schema",
		})
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			_ = encoder.Encode(ErrorResponse{
				Error: ErrorMessage{
					Code:    "INTERNAL_ERROR",
					Message: "Internal Error",
				},
			})
			return
		}

		if err := graphql.PrepareQuery(schema.Query, query.SelectionSet); err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			_ = encoder.Encode(ErrorResponse{
				Error: ErrorMessage{
					Code:    "INTERNAL_ERROR",
					Message: "Internal Error",
				},
			})
			return
		}

		executor := graphql.Executor{}
		value, err := executor.Execute(ctx, schema.Query, nil, query)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			_ = encoder.Encode(ErrorResponse{
				Error: ErrorMessage{
					Code:    "INTERNAL_ERROR",
					Message: "Internal Error",
				},
			})
			return
		}
		_ = encoder.Encode(value)
	}))

	return m
}
