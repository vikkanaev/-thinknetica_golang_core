package main

import (
	"net/http"
	"os"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/api"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/cache"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/shortner"

	"github.com/gorilla/mux"
)

var (
	webAddr                = getEnv("WEB_ADDR", ":3000")
	shortnerNewUrlEndpoint = getEnv("SHORTNER_NEW_URL", "http://localhost:8080/api/v1/url")
	shortnerUrlEndpoint    = getEnv("SHORTNER_URL", "http://localhost:8080/api/v1/urls")
	cacheUrlEndpoint       = getEnv("CACHE_URL", "http://localhost:8082/api/v1/urls")
)

func main() {
	router := mux.NewRouter()
	shortner := shortner.New(shortnerNewUrlEndpoint, shortnerUrlEndpoint)
	cache := cache.New(cacheUrlEndpoint)

	api := api.New(router, shortner, cache)
	api.Endpoints()
	http.ListenAndServe(webAddr, router)
}

// Читаем переменную окружения или значение по умолчанию
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
