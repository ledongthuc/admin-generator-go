package apiHandler

import "log"
import "github.com/ledongthuc/admin-generator-go/dataAccess"

// ColumnsAPIHandler use to handle API request
type ColumnsAPIHandler struct {
	HandlerBase
}

// Get logic of Column Handler
func (handler *ColumnsAPIHandler) Get(param map[string]string) (int, interface{}) {
	tableName := param["id"]
	log.Println(param)
	if tableName == "" {
		return 400, "Missing table_name"
	}
	columns := dataAccess.Column.GetByTable(tableName)
	return 200, columns
}
