package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"thinknetica_golang_core/pkg/crawler"
	"thinknetica_golang_core/pkg/crawler/spider"
)

// Сервер Интернет-поисковика GoSearch.
type gosearch struct {
	scanner crawler.Interface

	sites []string
	depth int
	addr  string
}

func main() {
	req := flags()
	sercher := new()

	res, err := serch(*sercher)
	if err != nil {
		log.Println(err)
		return
	}

	if req != "" {
		output(req, res)
	}
}

func flags() string {
	var request string
	flag.StringVar(&request, "s", "", "serch keyword")
	flag.Parse()
	return strings.ToLower(request)
}

// new создаёт объект и поисковика и возвращает указатель на него.
func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.depth = 2
	gs.addr = ":80"
	return &gs
}

func serch(s gosearch) (data []crawler.Document, err error) {
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

func output(req string, data []crawler.Document) {
	fmt.Println("Show results for:", req)
	for _, d := range data {
		if strings.Contains(strings.ToLower(d.Title), req) {
			fmt.Printf("Document: '%s' (%s)\n", d.Title, d.URL)
		}
	}
}
