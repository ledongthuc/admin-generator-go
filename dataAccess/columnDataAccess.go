package dataAccess

import (
	"fmt"

	"github.com/jbrodriguez/mlog"
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
	settings, err := helpers.LoadSettings()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	dbx, err := sqlx.Open(settings.Database.Type, settings.Database.ConnectionString)
	dbx.Close()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	var columns []entity.Column
	err = dbx.Select(&columns,
		`SELECT
            t.column_name,
            t.is_nullable,
            t.udt_name,
            t.table_schema,
            t.table_name,
            t.column_name = kcu.column_name as primary_key,
            t.column_default IS NOT NULL AND t.column_default LIKE 'nextval%%' as is_sequence
        FROM
            information_schema.columns t
        LEFT JOIN
            INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
        ON
            tc.table_schema = t.table_schema
            AND tc.table_name = t.table_name
            AND tc.constraint_type = 'PRIMARY KEY'
        LEFT JOIN
            INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu
        ON
            kcu.table_schema = tc.table_schema
            AND kcu.table_name = tc.table_name
            AND kcu.constraint_name = tc.constraint_name`)

	if err != nil {
		mlog.Error(err)
	}

	if len(columns) <= 0 {
		mlog.Warning("Don't have any tables in database")
	}

	return columns
}

// GetByTable use to select columns from `information_schema.tables` of inputed tableName.
func (columnDataAccess *columnDataAccess) GetByTable(tableName string) []entity.Column {
	settings, err := helpers.LoadSettings()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	dbx, err := sqlx.Open(settings.Database.Type, settings.Database.ConnectionString)
	if err != nil {
		mlog.Error(err)
		return nil
	}

	var columns []entity.Column
	queryString := fmt.Sprintf(`SELECT
        t.column_name,
        t.is_nullable,
        t.udt_name,
        t.table_schema,
        t.table_name,
        t.column_name = kcu.column_name as primary_key,
        t.column_default IS NOT NULL AND t.column_default LIKE 'nextval%%' as is_sequence
    FROM
        information_schema.columns t
    LEFT JOIN
        INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
    ON
        tc.table_schema = t.table_schema
        AND tc.table_name = t.table_name
        AND tc.constraint_type = 'PRIMARY KEY'
    LEFT JOIN
        INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu
    ON
        kcu.table_schema = tc.table_schema
        AND kcu.table_name = tc.table_name
        AND kcu.constraint_name = tc.constraint_name
    WHERE t.table_name = '%s'`, tableName)
	err = dbx.Select(&columns, queryString)
	dbx.Close()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	if len(columns) <= 0 {
		mlog.Warning("Don't have any columns in database")
		return nil
	}

	return columns
}
