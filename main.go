package main

import (
	"net/http"

	"github.com/jbrodriguez/mlog"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/ledongthuc/admin-generator-go/apiHandler"
)

func main() {
	mlog.Start(mlog.LevelInfo, "logs/app.log")

	martiniRunner := martini.Classic()
	martiniRunner.Use(render.Renderer())
	martiniRunner.Use(martini.Static("ui"))
	routingSetup(martiniRunner)
	martiniRunner.Run()
}

func routingSetup(martiniRunner *martini.ClassicMartini) {
	martiniRunner.Get("/api/**/:id", routingDetailFunc)
	martiniRunner.Get("/api/**", routingListFunc)
	martiniRunner.Post("/api/**", routingCreateFunc)
	martiniRunner.Delete("/api/**/:id", routingDeleteFunc)
}

func routingListFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).List(request, params))
}

func routingDetailFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Detail(request, params["id"]))
}

func routingCreateFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Create(request, params))
}

func routingDeleteFunc(params martini.Params, r render.Render, request *http.Request) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Delete(request, params["id"]))
}
