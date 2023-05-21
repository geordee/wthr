package weather

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
)

func Service(latitude string, longitude string) (*Weather, int, error) {
	var (
		err   error
		valid bool
	)

	valid, _ = regexp.MatchString(`^(\+|-)?(?:90(?:(?:\.0{1,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\.[0-9]{1,6})?))$`, latitude)
	if !valid {
		err = errors.New("invalid_latitude")
		return nil, http.StatusBadRequest, err
	}

	valid, _ = regexp.MatchString(`^(\+|-)?(?:180(?:(?:\.0{1,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\.[0-9]{1,6})?))$`, longitude)
	if !valid {
		err = errors.New("invalid_longitude")
		return nil, http.StatusBadRequest, err
	}

	req, err := http.NewRequest("GET", "https://api.open-meteo.com/v1/forecast", nil)
	if err != nil {
		log.Println(err)
		err = errors.New("http_request_error")
		return nil, http.StatusInternalServerError, err
	}

	q := req.URL.Query()
	q.Add("latitude", latitude)
	q.Add("longitude", longitude)
	q.Add("current_weather", "true")
	q.Add("hourly", "temperature_2m,relativehumidity_2m,windspeed_10m")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		err = errors.New("http_client_error")
		return nil, http.StatusInternalServerError, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		err = errors.New("http_read_error")
		return nil, http.StatusInternalServerError, err
	}

	weather := &Weather{}
	err = json.Unmarshal(body, weather)
	if err != nil {
		log.Println(err)
		err = errors.New("json_unmarshal_error")
		return nil, http.StatusInternalServerError, err
	}

	return weather, http.StatusOK, nil
}
