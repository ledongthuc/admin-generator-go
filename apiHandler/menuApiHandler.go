package apiHandler

import (
	"net/http"

	"github.com/ledongthuc/admin-generator-go/dataAccess"
)

// MenuAPIHandler use to handle API request
type MenuAPIHandler struct {
	HandlerBase
}

// Get logic of Menu Handler
func (handler *MenuAPIHandler) List(request *http.Request, param map[string]string) (int, interface{}) {
	tables := dataAccess.Table.GetAll()
	return 200, tables
}
