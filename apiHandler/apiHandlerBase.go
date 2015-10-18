package apiHandler

// HandlerBase base class for all api handlers
type HandlerBase interface {
	Get(param map[string]string) (int, interface{})
}
