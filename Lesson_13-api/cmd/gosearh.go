package main

import (
	"log"
	"net/http"

	"thinknetica_golang_core/Lesson_13-api/pkg/api"
	"thinknetica_golang_core/Lesson_13-api/pkg/crawler"
	"thinknetica_golang_core/Lesson_13-api/pkg/crawler/spider"

	"github.com/gorilla/mux"
)

// Сервер Интернет-поисковика GoSearch.
type gosearch struct {
	scanner crawler.Interface

	sites []string
	depth int

	data []crawler.Document
}

func main() {
	searcher := new()
	data, err := searcher.scan()
	if err != nil {
		log.Println(err)
		return
	}

	searcher.data = data

	router := mux.NewRouter()
	api := api.New(router, searcher.data)
	api.Endpoints()
	http.ListenAndServe(":8080", router)
}

// new создаёт объект и поисковика и возвращает указатель на него.
func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	// gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.sites = []string{"https://go.dev"}
	gs.depth = 2
	return &gs
}

// сканирует сайты
func (s *gosearch) scan() (data []crawler.Document, err error) {
	for _, site := range s.sites {
		res, e := s.scanner.Scan(site, s.depth)
		if e != nil {
			err = e
			break
		}
		data = append(data, res...)
	}
	return data, err
}
