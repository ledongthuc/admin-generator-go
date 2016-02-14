package main

import (
	"github.com/jbrodriguez/mlog"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"

	"github.com/ledongthuc/admin-generator-go/server"
)

func main() {
	mlog.Start(mlog.LevelInfo, "logs/app.log")
	martiniRunner := martini.Classic()
	martiniRunner.Use(render.Renderer())
	martiniRunner.Use(martini.Static("ui"))
	martiniRunner.Use(auth.BasicFunc(server.Authentication.Authenticate))
	server.Router.RoutingSetup(martiniRunner)
	martiniRunner.Run()
}
