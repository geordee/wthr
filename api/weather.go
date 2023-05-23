package api

import (
	"github.com/geordee/wthr/services/weather"
	"github.com/kataras/iris/v12"
)

func Weather(ctx iris.Context) {
	latitude := ctx.URLParam("latitude")
	longitude := ctx.URLParam("longitude")

	weather, code, err := weather.Service(latitude, longitude)
	ctx.StatusCode(code)

	if err != nil {
		ctx.JSON(iris.Map{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(weather)
	}
}
