package server

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/ledongthuc/admin-generator-go/apiHandler"
)

type router struct{}

// Routing object
var Router router

func (router *router) RoutingSetup(martiniRunner *martini.ClassicMartini) {
	martiniRunner.Get("/api/**/:id", router.routingDetailFunc)
	martiniRunner.Get("/api/**", router.routingListFunc)
	martiniRunner.Post("/api/**", router.routingCreateFunc)
	martiniRunner.Put("/api/**/:id", router.routingUpdateFunc)
	martiniRunner.Delete("/api/**/:id", router.routingDeleteFunc)
}

func (router *router) routingListFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).List(request, params))
}

func (router *router) routingDetailFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Detail(request, params["id"]))
}

func (router *router) routingCreateFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Create(request, params))
}

func (router *router) routingUpdateFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Update(request, params["id"], params))
}

func (router *router) routingDeleteFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Delete(request, params["id"]))
}
