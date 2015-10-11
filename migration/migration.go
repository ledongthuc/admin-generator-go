package migration

import (
	"fmt"

	"github.com/ledongthuc/admin-generator-go/dataAccess"
	"github.com/ledongthuc/admin-generator-go/entity"
)

// Run migration of admin generator. It will get tables and their columns, build configuration file and APIs
func Run() {
	// tables := getTables()
	// columns := getColumns(tables)
	// tables = fillColumnsIntoTables(tables, columns)
	return
}

func getTables() []entity.Table {
	tableDataAccess := dataAccess.NewTableDataAccess()
	tables := tableDataAccess.GetAll()
	fmt.Println(tables)
	return tables
}

func getColumns(tables []entity.Table) []entity.Column {
	columnDataAccess := dataAccess.NewColumnDataAccess()
	columns := columnDataAccess.GetByTables(tables)
	fmt.Println(columns)
	return columns
}
