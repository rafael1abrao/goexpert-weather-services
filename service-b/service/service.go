package service

import (
	"context"
	"errors"

	model "github.com/rafael1abrao/goexpert-weather-services/service-b/model"
	tracer "github.com/rafael1abrao/goexpert-weather-services/service-b/tracer"
	viacep "github.com/rafael1abrao/goexpert-weather-services/service-b/viacep"
	weather "github.com/rafael1abrao/goexpert-weather-services/service-b/weather"
)

var ErrZipcodeNotFound = errors.New("zipcode not found")

func GetWeatherByCEP(ctx context.Context, cep string) (*model.WeatherResponse, error) {
	ctx, span := tracer.Tracer().Start(ctx, "GetWeatherByCEP")
	defer span.End()

	location, err := viacep.FetchCityByCEP(ctx, cep)
	if err != nil {
		return nil, ErrZipcodeNotFound
	}

	weatherData, err := weather.FetchWeatherByCity(ctx, location.City)
	if err != nil {
		return nil, err
	}

	tempF := weatherData.TempC*1.8 + 32
	tempK := weatherData.TempC + 273

	return &model.WeatherResponse{
		City:  location.City,
		TempC: weatherData.TempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}
