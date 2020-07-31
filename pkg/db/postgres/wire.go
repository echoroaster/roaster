//+build wireinject

package postgres

import "github.com/google/wire"

var Set = wire.NewSet(New, wire.Value(StatementBuilderType))
