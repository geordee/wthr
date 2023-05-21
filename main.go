package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	// api root
	app.Get("/api", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"name": "wthr",
		})
	})

	app.Listen(":8080")
}
