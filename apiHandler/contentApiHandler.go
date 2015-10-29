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

	var responseCode = 200
	var result interface{}
	if param["id"] != "" {
		result = handler.getByID(param["id"])
	} else {
		result = dataAccess.Content.GetAll(handler.TableName)
		if result == nil {
			responseCode = 400
		}
	}

	return responseCode, result
}

func (handler *ContentAPIHandler) getByID(id string) interface{} {
	if handler.TableName == "" {
		return "Don't have any name"
	}

	key := dataAccess.Table.GetKeyByTableName(handler.TableName)
	if key == "" {
		return nil
	}

	result := dataAccess.Content.GetById(handler.TableName, key, id)
	return result
}
