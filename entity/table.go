package entity

import ()

type Table struct {
    Schema string `db:"table_schema"`
    Name   string `db:"table_name"`
}

func NewTable(schema string, name string) Table {
    return Table{Schema: schema, Name: name}
}
