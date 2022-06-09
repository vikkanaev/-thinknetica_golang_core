package main

import (
	"fmt"
	"net/http"
	"thinknetica_golang_core/Lesson_19-queue/1-Shortener/pkg/api"
	"thinknetica_golang_core/Lesson_19-queue/1-Shortener/pkg/queue"

	"github.com/gorilla/mux"
)

const (
	queueConnectionString = "amqp://guest:guest@localhost:5672/"
	queueName             = "UrlsApp"
	webAddr               = ":8080"
)

func main() {
	router := mux.NewRouter()
	q, err := queue.New(queueConnectionString, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Не забываем закрывать ресурсы
	defer q.Close()

	api := api.New(router, q)
	api.Endpoints()
	http.ListenAndServe(webAddr, router)
}
