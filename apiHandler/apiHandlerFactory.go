package apiHandler

type (
	apiHandlerFactory struct{}
)

// ApiHandlerFactory use for create api handlers
var APIHandlerFactory apiHandlerFactory

var commits = map[string]HandlerBase{
	"menu":   new(MenuAPIHandler),
	"column": new(ColumnsAPIHandler),
}

// GenerateAPIHandler generate handler base on the name
func (factory *apiHandlerFactory) GenerateAPIHandler(name string) *HandlerBase {
	var apiHandler HandlerBase
	apiHandler, hasHandler := commits[name]
	if hasHandler {
		return &apiHandler
	}

	apiHandler = CreateContentHandler(name)

	return &apiHandler
}
