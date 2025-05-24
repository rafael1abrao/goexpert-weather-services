package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	service "github.com/rafael1abrao/goexpert-weather-services/service-a/service"
)

type Input struct {
	CEP string `json:"cep"`
}

func HandleInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	var input Input
	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	matched := regexp.MustCompile(`^\d{8}$`).MatchString(input.CEP)
	if !matched {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	result, err := service.SendToServicoB(r.Context(), input.CEP)
	if err != nil {
		http.Error(w, "failed to contact service B", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
