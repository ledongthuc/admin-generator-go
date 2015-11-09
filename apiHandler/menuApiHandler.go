package apiHandler

import (
	"net/http"

	"github.com/jbrodriguez/mlog"

	"github.com/ledongthuc/admin-generator-go/dataAccess"
	"github.com/ledongthuc/admin-generator-go/entity"
)

// MenuAPIHandler use to handle API request
type MenuAPIHandler struct {
	HandlerBase
}

// Get logic of Menu Handler
func (handler *MenuAPIHandler) List(request *http.Request, param map[string]string) (int, interface{}) {
	tables := dataAccess.Table.GetAll()

	mappings, err := dataAccess.TableMapping.Load()
	if err != nil {
		mlog.Error(err)
		return 400, "Can't load mappings"
	}

	result := []entity.Table{}
	for _, table := range tables {
		tableMapping, existed := mappings[table.Name]
		if !existed || tableMapping.IsShow {
			result = append(result, table)
		}
	}
	return 200, result
}
