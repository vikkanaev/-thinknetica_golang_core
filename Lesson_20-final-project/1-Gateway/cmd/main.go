package main

import (
	"net/http"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/api"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/rpc"

	"github.com/gorilla/mux"
)

const (
	webAddr = ":3000"
)

func main() {
	router := mux.NewRouter()

	// Вот тут прям просится парсинг yaml. Отдельный пакет configReader?
	config := rpc.Config{
		Shortner: rpc.Shortner{
			NewUrl: "http://localhost:8080/api/v1/url",
			Urls:   "http://localhost:8080/api/v1/urls",
			Url:    "http://localhost:8080/api/v1/urls",
		},
	}
	rpc := rpc.New(config)

	api := api.New(router, rpc)
	api.Endpoints()
	http.ListenAndServe(webAddr, router)
}
