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

	// Обнуляем статистику при старте сервиса и перезаливаем данные
	err = setupAnalytics(storage, q)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Не забываем закрывать ресурсы
	defer q.Close()
	defer storage.Close()

	api := api.New(router, q, storage)
	api.Endpoints()
	http.ListenAndServe(webAddr, router)
}

// Обнуляем статистику при старте сервиса и перезаливаем данные
func setupAnalytics(s *storage.Storage, q *queue.Queue) error {
	// обнуляем статистику при старте сервиса
	q.PruneStat()

	// отправляем все ссылки в аналитику
	urls, err := s.Urls()
	if err != nil {
		return err
	}

	// Отпрвка отдельным потоком что бы не тормозить старт сервиса.
	go func() {
		for _, doc := range urls {
			q.NewUrl(doc.Long)
		}
	}()

	return nil
}
