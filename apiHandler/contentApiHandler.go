package apiHandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jbrodriguez/mlog"

	"github.com/ledongthuc/admin-generator-go/dataAccess"
	"github.com/ledongthuc/admin-generator-go/entity"
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

// Create create content
func (handler *ContentAPIHandler) Create(request *http.Request, data map[string]string) (int, interface{}) {
	mlog.Info("Table name: %s", handler.TableName)
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return 400, "Can't read the request"
	}

	var formParams map[string]string
	err = json.Unmarshal(body, &formParams)
	if err != nil {
		return 400, "Can't read the request"
	}

	if handler.TableName == "" {
		return 400, "Don't have any name"
	}

	columns := dataAccess.Column.GetByTable(handler.TableName)
	for name := range formParams {
		existedColumn := hasColumnName(name, columns)
		if !existedColumn {
			delete(formParams, name)
		}
	}

	if len(formParams) <= 0 {
		return 404, "Data should not empty, please fill something"
	}

	var responseCode = 201
	var result interface{}
	_, err = dataAccess.Content.New(handler.TableName, formParams)
	if err != nil {
		result = err.Error()
		responseCode = 400
	}

	return responseCode, result
}

// Delete content action
func (handler *ContentAPIHandler) Delete(request *http.Request, key string) (int, interface{}) {
	mlog.Info("Table name: %s", handler.TableName)
	if handler.TableName == "" {
		return 400, "Don't have any name"
	}

	keyName := dataAccess.Table.GetKeyByTableName(handler.TableName)
	if key == "" {
		return 404, "Don't have primary key"
	}

	err := dataAccess.Content.Delete(handler.TableName, keyName, key)
	if key == "" {
		mlog.Error(err)
		return 404, "Can't delete item"
	}

	return 200, "Successful"
}

func hasColumnName(key string, columns []entity.Column) bool {
	result := false
	for _, columnName := range columns {
		if columnName.Name == key {
			result = true
			break
		}
	}

	return result
}
