package dataAccess

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ledongthuc/admin-generator-go/entity"
	"github.com/ledongthuc/admin-generator-go/helpers"
)

// ColumnDataAccess contains methods that used for access to `information_schema.columns`.
type ColumnDataAccess struct {
}

// NewColumnDataAccess use create the instance of ColumnDataAccess
func NewColumnDataAccess() ColumnDataAccess {
	return ColumnDataAccess{}
}

// GetAll use to select all column from `information_schema.columns`.
func (columnDataAccess *ColumnDataAccess) GetAll() []entity.Column {
	var columns []entity.Column
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	if len(columns) <= 0 {
		log.Fatal("Don't have any tables in database")
	}

	return columns
}

// GetByTable use to select colums from `information_schema.tables` of inputed tables.
func (columnDataAccess *ColumnDataAccess) GetByTable(table entity.Table) []entity.Column {
	var columns []entity.Column
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		log.Fatal(err)
		return columns
	}

	queryString := `SELECT
        column_name,
        is_nullable,
        data_type,
        table_schema,
        table_name
    FROM
        information_schema.columns
    WHERE `
	queryString = queryString + columnDataAccess.buildWhereStringFromTables([]entity.Table{table})
	err = dbx.Select(&columns, queryString)

	if err != nil {
		log.Fatal(err)
	}

	if len(columns) <= 0 {
		log.Fatal("Don't have any tables in database")
	}

	return columns
}

// GetByTables use to select colums from `information_schema.tables` of inputed tables.
func (columnDataAccess *ColumnDataAccess) GetByTables(tables []entity.Table) []entity.Column {
	var columns []entity.Column
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		log.Fatal(err)
		return columns
	}

	queryString := `SELECT
        column_name,
        is_nullable,
        data_type,
        table_schema,
        table_name
    FROM
        information_schema.columns
    WHERE `
	queryString = queryString + columnDataAccess.buildWhereStringFromTables(tables)
	err = dbx.Select(&columns, queryString)

	if err != nil {
		log.Fatal(err)
	}

	if len(columns) <= 0 {
		log.Fatal("Don't have any tables in database")
	}

	return columns
}

func (columnDataAccess *ColumnDataAccess) buildWhereStringFromTables(tables []entity.Table) string {
	var result []string
	for _, table := range tables {
		result = append(result, fmt.Sprintf("(table_schema = '%s' AND table_name = '%s')", table.Schema, table.Name))

	}
	return strings.Join(result, " OR ")
}
