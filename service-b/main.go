package main

import (
	"log"
	"net/http"

	handler "github.com/rafael1abrao/goexpert-weather-services/service-b/handler"
	tracer "github.com/rafael1abrao/goexpert-weather-services/service-b/tracer"
)

func main() {
	shutdown := tracer.InitProvider()
	defer shutdown()

	http.HandleFunc("/weather", handler.HandleWeather)
	log.Println("Serviço B rodando na porta 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
