package main

import (
	"fmt"
	"net/http"
	"thinknetica_golang_core/Lesson_20-final-project/2-Shortener/pkg/api"
	"thinknetica_golang_core/Lesson_20-final-project/2-Shortener/pkg/queue"
	"thinknetica_golang_core/Lesson_20-final-project/2-Shortener/pkg/storage"

	"github.com/gorilla/mux"
)

const (
	queueConnectionString = "amqp://guest:guest@localhost:5672/"
	queueName             = "UrlsApp"
	webAddr               = ":8080"

	mongoConn     = "mongodb://localhost:27017/" // строка подключения к Монго
	mongoDbName   = "shortener"                  // имя БД в Монго
	mongoCollName = "urls"                       // имя коллекции в Монго
)

func main() {
	router := mux.NewRouter()
	storage, err := storage.New(mongoConn, mongoDbName, mongoCollName)
	if err != nil {
		fmt.Println(err)
		return
	}

	q, err := queue.New(queueConnectionString, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// обнуляем статистику при старте сервиса
	q.PruneStat()
	// Не забываем закрывать ресурсы
	defer q.Close()

	api := api.New(router, q, storage)
	api.Endpoints()
	http.ListenAndServe(webAddr, router)
}
