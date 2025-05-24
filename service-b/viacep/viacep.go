package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/rafael1abrao/goexpert-weather-services/service-b/model"
	tracer "github.com/rafael1abrao/goexpert-weather-services/service-b/tracer"
)

func FetchCityByCEP(ctx context.Context, cep string) (*model.Location, error) {
	ctx, span := tracer.Tracer().Start(ctx, "viacep.fetch")
	defer span.End()

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
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
		return nil, fmt.Errorf("viacep status: %d", resp.StatusCode)
	}

	var location model.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	return &location, nil
}
