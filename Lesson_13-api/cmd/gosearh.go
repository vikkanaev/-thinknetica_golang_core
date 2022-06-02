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
	api     *api.API
	router  *mux.Router

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

	searcher.api.Endpoints(searcher.data)
	http.ListenAndServe(":8080", searcher.router)
}

// new создаёт объект и поисковика и возвращает указатель на него.
func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.router = mux.NewRouter()
	gs.api = api.New(gs.router)
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
