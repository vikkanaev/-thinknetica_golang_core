package main

import (
	"fmt"
	"net/http"
	"os"
	"thinknetica_golang_core/Lesson_20-final-project/4-Analytics/pkg/api"
	"thinknetica_golang_core/Lesson_20-final-project/4-Analytics/pkg/queue"
	"thinknetica_golang_core/Lesson_20-final-project/4-Analytics/pkg/storage"

	"github.com/gorilla/mux"
)

var (
	queueConnectionString = getEnv("QUEUE_CONN", "amqp://guest:guest@localhost:5672/")
	queueName             = getEnv("QUEUE_NAME", "UrlsApp")
	webAddr               = getEnv("WEB_ADDR", ":8081")
)

func main() {
	router := mux.NewRouter()
	storage := storage.New()
	q, err := queue.New(queueConnectionString, queueName, storage)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Не забываем закрывать ресурсы
	defer q.Close()

	api := api.New(router, q, storage)
	api.Endpoints()
	go http.ListenAndServe(webAddr, router)
	q.Consume()
}

// Читаем переменную окружения или значение по умолчанию
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
