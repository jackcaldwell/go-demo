package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", Server)
	http.ListenAndServe(":80", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(apiResponse{
		Data: "Response sent from API",
	})
}

type apiResponse struct {
	Data string `json:"data"`
}

