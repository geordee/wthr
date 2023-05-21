package weather

type Weather struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentWeather       struct {
		Temperature   float64 `json:"temperature,omitempty"`
		Windspeed     float64 `json:"windspeed,omitempty"`
		Winddirection float64 `json:"winddirection,omitempty"`
		Weathercode   int     `json:"weathercode,omitempty"`
		IsDay         int     `json:"is_day,omitempty"`
		Time          string  `json:"time,omitempty"`
	} `json:"current_weather"`
	HourlyUnits struct {
		Time               string `json:"time,omitempty"`
		Temperature2M      string `json:"temperature_2m,omitempty"`
		Relativehumidity2M string `json:"relativehumidity_2m,omitempty"`
		Windspeed10M       string `json:"windspeed_10m,omitempty"`
	} `json:"hourly_units,omitempty"`
	Hourly struct {
		Time               []string  `json:"time,omitempty"`
		Temperature2M      []float64 `json:"temperature_2m,omitempty"`
		Relativehumidity2M []int     `json:"relativehumidity_2m,omitempty"`
		Windspeed10M       []float64 `json:"windspeed_10m,omitempty"`
	} `json:"hourly,omitempty"`
}
