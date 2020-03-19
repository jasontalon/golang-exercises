package handlers

import (
	"net/http"
	//"fmt"
	"encoding/json"

	"github.com/jasontalon/golang-exercises/easy"
)

type AcronymRequest struct {
	Text string `json:"text"`
}

type AcronymResponse struct {
	Text string `json:"text"`
}

//Acronym handler
func Acronym(w http.ResponseWriter, r *http.Request) {
	request := AcronymRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AcronymResponse{Text: easy.Acronym(request.Text)})
}
