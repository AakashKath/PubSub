package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Response struct
type Response struct {
	Message string `json:"message,omitempty"`
}

//IndexPage for Microservice
func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Welcome to PubSub."}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		return
	}
}

// RunRESTServer starts a http listener server
func RunRESTServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexPage).Methods("GET")
	serverHost := fmt.Sprintf("0.0.0.0:9000")
	if err := http.ListenAndServe(serverHost, router); err != nil {
		fmt.Println("Failed to start HTTP server")
	}
}
