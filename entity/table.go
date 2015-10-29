package entity

// Table entity in table `information_schema.tables`
type Table struct {
	Schema     string `db:"table_schema" json:"-"`
	Name       string `db:"table_name" json:"name"`
	PrimaryKey string `db:"primary_key" json:"primary_key"`
}

// NewTable create an instance of `Table` entity
func NewTable(schema string, name string) Table {
	return Table{Schema: schema, Name: name}
}
