package routes

import (
	"github.com/geordee/wthr/api"
	"github.com/geordee/wthr/pages"
	"github.com/kataras/iris/v12"
)

func Weather(ctx iris.Context) {
	switch ctx.GetHeader("Accept") {
	case "application/json":
		api.Weather(ctx)
	case "text/html":
		pages.Weather(ctx)
	default:
		pages.Weather(ctx)
	}
}
