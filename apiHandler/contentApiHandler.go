package apiHandler

import "github.com/ledongthuc/admin-generator-go/dataAccess"

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
func (handler *ContentAPIHandler) Get(param map[string]string) (int, interface{}) {
	if handler.TableName == "" {
		return 400, "Don't have any name"
	}
	result := dataAccess.Content.GetAll(handler.TableName)
	return 200, result
}
