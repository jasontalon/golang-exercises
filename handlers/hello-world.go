package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jasontalon/golang-exercises/easy"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
}

//HelloWorld handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := easy.HelloWorld()
	fmt.Println(m)
	payload := helloResponse{Message: m}

	json.NewEncoder(w).Encode(payload)
}
