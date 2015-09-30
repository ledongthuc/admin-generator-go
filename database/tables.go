package database

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"

    . "github.com/ledongthuc/admin-generator-go/entity"
    "github.com/ledongthuc/admin-generator-go/helpers"
)

type TableDataAccess struct {
}

func NewTableDataAccess() TableDataAccess {
    return TableDataAccess{}
}

func (this *TableDataAccess) GetAll() []Table {
    var tables []Table
    configuration := helpers.LoadConfiguration()

    dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
    if err != nil {
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

    return tables
}
