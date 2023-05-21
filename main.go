package main

import (
	"github.com/geordee/wthr/api"
	"github.com/geordee/wthr/pages"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	// api root
	app.Get("/api", api.Root)

	// Weather API
	app.Get("/api/weather", api.Weather)

	// Weather Page
	app.Get("/weather", pages.Weather)

	app.Listen(":8080")
}
