package main

import (
	"fmt"
	"net/http"
	"thinknetica_golang_core/Lesson_20-final-project/3-Cache/pkg/api"
	"thinknetica_golang_core/Lesson_20-final-project/3-Cache/pkg/queue"
	"thinknetica_golang_core/Lesson_20-final-project/3-Cache/pkg/storage"

	"github.com/gorilla/mux"
)

const (
	queueConnectionString = "amqp://guest:guest@localhost:5672/"
	queueName             = "UrlsApp"
	webAddr               = ":8082"
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
	// Закрыть сторедж тут

	api := api.New(router, q, storage)
	api.Endpoints()
	go http.ListenAndServe(webAddr, router)
	q.Consume()

}
