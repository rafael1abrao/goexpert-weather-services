package main

import (
	"log"
	"net/http"

	handler "github.com/rafael1abrao/goexpert-weather-services/service-a/handler"
	tracer "github.com/rafael1abrao/goexpert-weather-services/service-a/tracer"
)

func main() {
	shutdown := tracer.InitProvider()
	defer shutdown()

	http.HandleFunc("/input", handler.HandleInput)
	log.Println("Serviço A rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
