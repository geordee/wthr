package pages

import (
	"github.com/geordee/wthr/services/weather"
	"github.com/kataras/iris/v12"
)

func Weather(ctx iris.Context) {
	latitude := ctx.URLParam("latitude")
	longitude := ctx.URLParam("longitude")

	weather, code, err := weather.Service(latitude, longitude)
	if err != nil {
		ctx.View("error", iris.Map{
			"Code":  code,
			"Error": err.Error(),
		})
		return
	}

	ctx.View("weather", weather)
}
