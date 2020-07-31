package graphql // import "github.com/echoroaster/roaster/pkg/graphql"

import (
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

type SchemaRegister interface {
	Register(builder *schemabuilder.Schema)
}

func MakeSchema(schemas ...SchemaRegister) (schema *graphql.Schema, err error) {
	builder := schemabuilder.NewSchema()
	for _, register := range schemas {
		register.Register(builder)
	}
	schema, err = builder.Build()
	if err != nil {
		return
	}
	introspection.AddIntrospectionToSchema(schema)
	return
}
