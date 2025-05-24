package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	model "github.com/rafael1abrao/goexpert-weather-services/service-b/model"
	tracer "github.com/rafael1abrao/goexpert-weather-services/service-b/tracer"
)

func FetchWeatherByCity(ctx context.Context, city string) (*model.Weather, error) {
	ctx, span := tracer.Tracer().Start(ctx, "weather.fetch")
	defer span.End()

	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, url.QueryEscape(city))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weatherapi status: %d", resp.StatusCode)
	}

	type weatherResp struct {
		Current model.Weather `json:"current"`
	}
	var result weatherResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result.Current, nil
}
