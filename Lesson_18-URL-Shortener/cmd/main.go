package main

import (
	"net/http"
	"thinknetica_golang_core/Lesson_18-URL-Shortener/pkg/api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	api := api.New(router)
	api.Endpoints()
	http.ListenAndServe(":8080", router)
}
