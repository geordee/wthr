package api

import (
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/kataras/iris/v12"
)

func Weather(ctx iris.Context) {
	var (
		err   error
		valid bool
	)

	latitude := ctx.URLParam("latitude")
	valid, _ = regexp.MatchString(`^(\+|-)?(?:90(?:(?:\.0{1,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\.[0-9]{1,6})?))$`, latitude)
	if !valid {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"error": "invalid_latitude",
		})
		return
	}

	longitude := ctx.URLParam("longitude")
	valid, _ = regexp.MatchString(`^(\+|-)?(?:180(?:(?:\.0{1,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\.[0-9]{1,6})?))$`, longitude)
	if !valid {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"error": "invalid_longitude",
		})
		return
	}

	req, err := http.NewRequest("GET", "https://api.open-meteo.com/v1/forecast", nil)
	if err != nil {
		log.Println(err)
	}

	q := req.URL.Query()
	q.Add("latitude", latitude)
	q.Add("longitude", longitude)
	q.Add("current_weather", "true")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	ctx.StatusCode(resp.StatusCode)
	ctx.Write(body)
}
