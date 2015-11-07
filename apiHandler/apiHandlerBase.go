package apiHandler

import "net/http"

// HandlerBase base class for all api handlers
type HandlerBase interface {
	List(request *http.Request, param map[string]string) (int, interface{})
	Detail(request *http.Request, key string) (int, interface{})
	Create(request *http.Request, data map[string]string) (int, interface{})
	Update(request *http.Request, keyValue string, data map[string]string) (int, interface{})
	Delete(request *http.Request, key string) (int, interface{})
}
