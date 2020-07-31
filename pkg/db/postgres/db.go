package postgres // import "github.com/echoroaster/roaster/pkg/db/postgres"

import (
	"database/sql"
	"database/sql/driver"
	"os"

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"go.opencensus.io/stats/view"
)

var StatementBuilderType = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func init() {
	_ = view.Register(ocsql.DefaultViews...)
}

func New() (db *sql.DB, cleanup func(), err error) {
	var connector driver.Connector
	connector, err = pq.NewConnector(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, nil, err
	}
	connector = ocsql.WrapConnector(connector, ocsql.WithAllTraceOptions())
	db = sql.OpenDB(connector)
	cleanup = func() {
		_ = db.Close()
	}
	return
}
