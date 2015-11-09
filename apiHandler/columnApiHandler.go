package apiHandler

import (
	"net/http"

	"github.com/ledongthuc/admin-generator-go/dataAccess"
)

// ColumnsAPIHandler use to handle API request
type ColumnsAPIHandler struct {
	HandlerBase
}

// Get logic of Column Handler
func (handler *ColumnsAPIHandler) List(request *http.Request, param map[string]string) (int, interface{}) {
	tableName := request.FormValue("table_name")
	if tableName == "" {
		return 400, "Missing table_name"
	}

	isShowed := dataAccess.TableMapping.IsShowTable(tableName)
	if !isShowed {
		return 404, "Doesn't see table " + tableName
	}

	columns := dataAccess.Column.GetByTable(tableName)
	return 200, columns
}
