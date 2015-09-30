package main

import (
    "log"
    "os"

    "github.com/go-martini/martini"

    "github.com/ledongthuc/admin-generator-go/migration"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "migration" {
        migration.Run()
        return
    }

    martiniRunner := martini.Classic()
    martiniRunner.Group("/lockers", func(router martini.Router) {
        router.Get("/:id", func(logger *log.Logger) (int, string) {
            return 200, "Get!"
        })
        router.Post("/", func() string {
            return "New!"
        })
        router.Put("/", func() string {
            return "Update!"
        })
        router.Delete("/", func() string {
            return "Delete!"
        })
    })
    martiniRunner.Run()
}
