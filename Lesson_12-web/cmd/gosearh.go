package main

import (
	"log"

	"thinknetica_golang_core/Lesson_12-web/pkg/crawler"
	"thinknetica_golang_core/Lesson_12-web/pkg/crawler/spider"
	"thinknetica_golang_core/Lesson_12-web/pkg/index"
	"thinknetica_golang_core/Lesson_12-web/pkg/webapp"
)

// Сервер Интернет-поисковика GoSearch.
type gosearch struct {
	scanner crawler.Interface
	index   index.Index
	web     webapp.Interface

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
	searcher.index.Add(data)

	searcher.web.Start(&searcher.index, searcher.data)
}

// new создаёт объект и поисковика и возвращает указатель на него.
func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.index = *index.New()
	gs.web = webapp.New()
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
