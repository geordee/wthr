package api

import "github.com/kataras/iris/v12"

func Root(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"name": "wthr",
	})
}
