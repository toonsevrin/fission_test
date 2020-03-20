package main

import (
    "net/http"
    "encoding/json"
)

type Response struct {
	Message string `json:"msg"`
}
// Handler is the entry point for this fission function
func Handler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, world!"}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
