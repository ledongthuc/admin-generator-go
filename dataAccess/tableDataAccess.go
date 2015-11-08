package dataAccess

import (
	"github.com/jbrodriguez/mlog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ledongthuc/admin-generator-go/entity"
	"github.com/ledongthuc/admin-generator-go/helpers"
)

// TableDataAccess contains methods that used for access to `information_schema.tables`.
type tableDataAccess struct {
}

// Table instance of Table Data Access
var Table tableDataAccess

// GetKeyByTableName primary key by table name
func (tableDataAccess *tableDataAccess) GetKeyByTableName(tableName string) string {
	settings, err := helpers.LoadSettings()
	if err != nil {
		mlog.Error(err)
		return ""
	}

	dbx, err := sqlx.Open(settings.Database.Type, settings.Database.ConnectionString)
	if err != nil {
		mlog.Error(err)
		return ""
	}

	query := `SELECT
        t.table_schema, t.table_name, kcu.column_name as primary_key
    FROM
        information_schema.tables t
    LEFT JOIN INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
         ON tc.table_catalog = t.table_catalog
         AND tc.table_schema = t.table_schema
         AND tc.table_name = t.table_name
         AND tc.constraint_type = 'PRIMARY KEY'
    LEFT JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu
         ON kcu.table_catalog = tc.table_catalog
         AND kcu.table_schema = tc.table_schema
         AND kcu.table_name = tc.table_name
         AND kcu.constraint_name = tc.constraint_name
    WHERE
        t.table_schema NOT IN ('pg_catalog', 'information_schema')
        AND t.table_name = '` + tableName + `'
    LIMIT 1`

	var tables []entity.Table
	err = dbx.Select(&tables, query)
	dbx.Close()
	if err != nil {
		mlog.Error(err)
		return ""
	}

	if len(tables) <= 0 {
		mlog.Warning("Don't have any tables in database")
		return ""
	}

	return tables[0].PrimaryKey
}

// GetAll use to select all tables from `information_schema.tables`.
func (tableDataAccess *tableDataAccess) GetAll() []entity.Table {
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

	var tables []entity.Table
	err = dbx.Select(&tables,
		`SELECT
            t.table_schema, t.table_name, kcu.column_name as primary_key
        FROM
            information_schema.tables t
        LEFT JOIN INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
             ON tc.table_catalog = t.table_catalog
             AND tc.table_schema = t.table_schema
             AND tc.table_name = t.table_name
             AND tc.constraint_type = 'PRIMARY KEY'
        LEFT JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu
             ON kcu.table_catalog = tc.table_catalog
             AND kcu.table_schema = tc.table_schema
             AND kcu.table_name = tc.table_name
             AND kcu.constraint_name = tc.constraint_name
        WHERE
            t.table_schema NOT IN ('pg_catalog', 'information_schema')`)
	dbx.Close()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	if len(tables) <= 0 {
		mlog.Warning("Don't have any tables in database")
		return nil
	}

	return tables
}
