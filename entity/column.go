package entity

// Column entity in table `information_schema.columns`
type Column struct {
	Name        string `db:"column_name" json:"name"`
	Nullable    string `db:"is_nullable" json:"-"`
	DataType    string `db:"udt_name" json:"data_type"`
	TableName   string `db:"table_schema" json:"-"`
	TableSchema string `db:"table_name" json:"table_name"`
	PrimaryKey  bool   `db:"primary_key" json:"primary"`
}

// NewColumn create an instance of `Column` entity
func NewColumn(name string, nullable string, dataType string) Column {
	return Column{Name: name, Nullable: nullable, DataType: dataType}
}
