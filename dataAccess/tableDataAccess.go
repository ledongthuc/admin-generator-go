package dataAccess

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ledongthuc/admin-generator-go/entity"
	"github.com/ledongthuc/admin-generator-go/helpers"
)

// TableDataAccess contains methods that used for access to `information_schema.tables`.
type TableDataAccess struct {
}

// NewTableDataAccess use create the instance of TableDataAccess
func NewTableDataAccess() TableDataAccess {
	return TableDataAccess{}
}

// GetAll use to select all tables from `information_schema.tables`.
func (tableDataAccess *TableDataAccess) GetAll() []entity.Table {
	var tables []entity.Table
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		log.Fatal(err)
		return tables
	}

	err = dbx.Select(&tables,
		`SELECT
            table_schema, table_name
        FROM
            information_schema.tables
        WHERE
            table_schema='public' AND
            table_type='BASE TABLE';`)

	if err != nil {
		log.Fatal(err)
	}

	if len(tables) <= 0 {
		log.Fatal("Don't have any tables in database")
	}

	return tables
}
