package main

import (
	"github.com/geordee/wthr/routes"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.HandleDir("/static", iris.Dir("./static"))

	// routes
	app.Get("/ping", routes.Ping)
	app.Get("/weather", routes.Weather)

	app.Listen(":8080")
}
