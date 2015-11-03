package apiHandler

import (
	"net/http"

	"github.com/jbrodriguez/mlog"

	"github.com/ledongthuc/admin-generator-go/dataAccess"
)

// ContentAPIHandler use to handle API request
type ContentAPIHandler struct {
	TableName string
	HandlerBase
}

// CreateContentHandler create Content API handler with table name
func CreateContentHandler(tableName string) *ContentAPIHandler {
	handler := new(ContentAPIHandler)
	handler.TableName = tableName
	return handler
}

// Get logic of content
func (handler *ContentAPIHandler) List(request *http.Request, param map[string]string) (int, interface{}) {
	mlog.Info("Table name: %s", handler.TableName)
	if handler.TableName == "" {
		return 400, "Don't have any name"
	}

	var responseCode = 200
	var result interface{}
	result = dataAccess.Content.GetAll(handler.TableName)
	if result == nil {
		responseCode = 400
	}

	return responseCode, result
}

func (handler *ContentAPIHandler) Detail(request *http.Request, keyValue string) (int, interface{}) {
	mlog.Info("Table name: %s, keyValue: %s", handler.TableName, keyValue)
	if handler.TableName == "" {
		return 400, "Don't have any name"
	}

	key := dataAccess.Table.GetKeyByTableName(handler.TableName)
	if key == "" {
		return 400, "Don't have primary key"
	}

	result := dataAccess.Content.GetById(handler.TableName, key, keyValue)
	return 200, result
}
