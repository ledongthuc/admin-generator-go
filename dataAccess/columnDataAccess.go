package dataAccess

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ledongthuc/admin-generator-go/entity"
	"github.com/ledongthuc/admin-generator-go/helpers"
)

// ColumnDataAccess contains methods that used for access to `information_schema.columns`.
type columnDataAccess struct {
}

// Column is the instance of ColumnDataAccess
var Column columnDataAccess

// GetAll use to select all column from `information_schema.columns`.
func (columnDataAccess *columnDataAccess) GetAll() []entity.Column {
	var columns []entity.Column
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	dbx.Close()
	if err != nil {
		log.Println(err)
		return columns
	}

	err = dbx.Select(&columns,
		`SELECT
            column_name,
            is_nullable,
            data_type,
            table_schema,
            table_name
        FROM
            information_schema.columns`)

	if err != nil {
		log.Println(err)
	}

	if len(columns) <= 0 {
		log.Println("Don't have any tables in database")
	}

	return columns
}

// GetByTable use to select columns from `information_schema.tables` of inputed tableName.
func (columnDataAccess *columnDataAccess) GetByTable(tableName string) []entity.Column {
	var columns []entity.Column
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		log.Println(err)
		return columns
	}

	queryString := fmt.Sprintf(`SELECT
        column_name,
        is_nullable,
        data_type,
        table_schema,
        table_name
    FROM
        information_schema.columns
    WHERE table_name = '%s'`, tableName)
	err = dbx.Select(&columns, queryString)
	dbx.Close()
	if err != nil {
		log.Println(err)
		return nil
	}

	if len(columns) <= 0 {
		log.Println("Don't have any columns in database")
		return nil
	}

	return columns
}
