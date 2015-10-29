package main

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/ledongthuc/admin-generator-go/apiHandler"
)

func main() {
	martiniRunner := martini.Classic()
	martiniRunner.Use(render.Renderer())
	martiniRunner.Use(martini.Static("ui"))
	routingSetup(martiniRunner)
	martiniRunner.Run()
}

func routingSetup(martiniRunner *martini.ClassicMartini) {
	martiniRunner.Get("/api/**/:id", routingGetFunc)
	martiniRunner.Get("/api/**", routingGetFunc)
}

func routingGetFunc(params martini.Params, log *log.Logger, r render.Render) {
	apiHandler := apiHandler.APIHandlerFactory.GenerateAPIHandler(params["_1"])
	if apiHandler == nil {
		r.JSON(404, "404 - API is not exist")
	}

	r.JSON((*apiHandler).Get(params))
}
