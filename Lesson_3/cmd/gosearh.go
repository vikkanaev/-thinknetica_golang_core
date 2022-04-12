package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	"thinknetica_golang_core/Lesson_3/pkg/crawler"
	"thinknetica_golang_core/Lesson_3/pkg/crawler/spider"
	"thinknetica_golang_core/Lesson_3/pkg/index"
)

// Сервер Интернет-поисковика GoSearch.
type gosearch struct {
	scanner crawler.Interface
	index   index.Index

	sites []string
	depth int

	// data это хранилище документов и его прям просится вынести в отдельный пакет.
	// не делаю это только из-за требования "решить задание максимально просто"
	data []crawler.Document
}

type notFoundErr struct {
	code int
}

func (e *notFoundErr) Error() string {
	return fmt.Sprintf("%d", e.code)
}

func main() {
	searcher := new()
	err := searcher.scan()
	if err != nil {
		log.Println(err)
		return
	}

	q := query()
	if q != "" {
		docs := searcher.search(q)
		output(q, docs)
	}
}

func query() string {
	var request string
	flag.StringVar(&request, "s", "", "serch keyword")
	flag.Parse()
	return strings.ToLower(request)
}

// new создаёт объект и поисковика и возвращает указатель на него.
func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.index = *index.New()
	gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.depth = 2
	return &gs
}

// сканирует сайты и помещает данные в индекс и хранилище документов
func (s *gosearch) scan() (err error) {
	var data []crawler.Document
	for _, site := range s.sites {
		res, e := s.scanner.Scan(site, s.depth)
		if e != nil {
			err = e
			break
		}
		data = append(data, res...)
		s.index.Add(data)
	}
	s.data = data
	return err
}

// ищет заданное слово по индексу и возвращает совпавщие документы
func (s *gosearch) search(req string) (docs []crawler.Document) {
	for _, id := range s.index.Search(req) {
		d, e := s.document(id)
		if e.Error() == "0" {
			docs = append(docs, d)
		}
	}
	return docs
}

// производить поиск документа по id в хранилище документов методом бинарного поиска
func (s *gosearch) document(id int) (crawler.Document, error) {
	var doc crawler.Document
	var err notFoundErr

	i := sort.Search(len(s.data), func(i int) bool { return s.data[i].ID >= id })
	if i < len(s.data) && s.data[i].ID == id {
		doc = s.data[i]
	} else {
		err.code = -1
	}
	return doc, &err
}

// выводит конечный результат
func output(req string, docs []crawler.Document) {
	if len(docs) > 0 {
		fmt.Println("Show results for:", req)
		for _, d := range docs {
			fmt.Printf("Document: '%s' (%s)\n", d.Title, d.URL)
		}
	} else {
		fmt.Println("Nothing found for:", req)
	}
}
