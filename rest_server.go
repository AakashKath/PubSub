package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AakashKath/PubSub/settings"
	"github.com/AakashKath/PubSub/views"
	"github.com/gorilla/mux"
)

//IndexPage for Microservice
func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := views.Response{Message: "Welcome to PubSub."}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		return
	}
}

// RunRESTServer starts a http listener server
func RunRESTServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexPage).Methods(http.MethodGet)
	router.HandleFunc("/create_topic", views.CreateTopic).Methods(http.MethodPost)
	router.HandleFunc("/delete_topic", views.DeleteTopic).Methods(http.MethodDelete)
	router.HandleFunc("/add_subscription", views.AddSubscription).Methods(http.MethodPost)
	router.HandleFunc("/delete_subscription", views.DeleteSubscription).Methods(http.MethodDelete)
	router.HandleFunc("/publish", views.Publish).Methods(http.MethodPost)
	router.HandleFunc("/subscribe", views.Subscribe).Methods(http.MethodPost)
	router.HandleFunc("/unsubscribe", views.UnSubscribe).Methods(http.MethodPatch)
	serverHost := fmt.Sprintf("0.0.0.0:%d", settings.GetSettings().Generic.Port)
	if err := http.ListenAndServe(serverHost, router); err != nil {
		fmt.Println("Failed to start HTTP server")
	}
}
