package model

type Location struct {
	City  string `json:"localidade"`
	State string `json:"uf"`
}

type Weather struct {
	TempC float64 `json:"temp_c"`
}

type WeatherResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}
