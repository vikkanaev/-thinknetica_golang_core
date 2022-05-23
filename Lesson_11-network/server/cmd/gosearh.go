package main

import (
	"log"

	"thinknetica_golang_core/Lesson_11-network/server/pkg/crawler"
	"thinknetica_golang_core/Lesson_11-network/server/pkg/crawler/spider"
	"thinknetica_golang_core/Lesson_11-network/server/pkg/netsrv"
)

// Сервер Интернет-поисковика GoSearch.
type gosearch struct {
	scanner crawler.Interface
	web     netsrv.Interface

	sites []string
	depth int
}

func main() {
	log.Println("Server starting")
	searcher := new()

	log.Println("Start site scaning")
	data, err := searcher.search()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Scan complete")

	searcher.web.Start(data)
}

// new создаёт объект поисковика и возвращает указатель на него.
func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.web = netsrv.New()
	gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.depth = 2
	return &gs
}

// Производить сканирование сайтов
func (s *gosearch) search() (data []crawler.Document, err error) {
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
