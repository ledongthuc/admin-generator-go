package dataAccess

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ledongthuc/admin-generator-go/entity"
	"github.com/ledongthuc/admin-generator-go/helpers"
)

// TableDataAccess contains methods that used for access to `information_schema.tables`.
type tableDataAccess struct {
}

var Table tableDataAccess

// GetAll use to select all tables from `information_schema.tables`.
func (tableDataAccess *tableDataAccess) GetAll() []entity.Table {
	var tables []entity.Table
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = dbx.Select(&tables,
		`SELECT
            table_schema, table_name
        FROM
            information_schema.tables
        WHERE
            table_schema='public' AND
            table_type='BASE TABLE';`)
	dbx.Close()
	if err != nil {
		log.Println(err)
		return nil
	}

	if len(tables) <= 0 {
		log.Println("Don't have any tables in database")
		return nil
	}

	return tables
}
