package entity

// Column entity in table `information_schema.columns`
type Column struct {
	Name        string `db:"column_name"`
	Nullable    string `db:"is_nullable"`
	DataType    string `db:"data_type"`
	TableName   string `db:"table_schema"`
	TableSchema string `db:"table_name"`
}

// NewColumn create an instance of `Column` entity
func NewColumn(name string, nullable string, dataType string) Column {
	return Column{Name: name, Nullable: nullable, DataType: dataType}
}
