package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	"thinknetica_golang_core/Lesson_5-io/pkg/crawler"
	"thinknetica_golang_core/Lesson_5-io/pkg/crawler/spider"
	"thinknetica_golang_core/Lesson_5-io/pkg/index"
	"thinknetica_golang_core/Lesson_5-io/pkg/storage"
)

// Сервер Интернет-поисковика GoSearch.
type gosearch struct {
	scanner crawler.Interface
	index   index.Index
	storage storage.Interface

	sites []string
	depth int
	data  []crawler.Document
}

type notFoundErr struct {
	code int
	msg  string
}

func (e *notFoundErr) Error() string {
	return fmt.Sprintf("%d", e.code)
}

func main() {
	s, err := new()
	if err != nil {
		log.Println(err)
		return
	}

	docs, err := s.storage.Retrieve()
	if err != nil {
		log.Println(err)
		return
	}

	if len(docs) > 0 {
		s.data = docs
		s.index.Add(docs)
	} else {
		err = s.scan()
		if err != nil {
			log.Println(err)
			return
		}
	}

	q := query()
	if q != "" {
		docs, err := s.search(q)
		if err != nil {
			log.Println(err)
			return
		}

		output(q, docs)
	}
}

// Получает запрос пользователя из командной строки
func query() string {
	var request string
	flag.StringVar(&request, "s", "", "serch keyword")
	flag.Parse()
	return strings.ToLower(request)
}

// Функция new создаёт объект и поисковика и возвращает указатель на него.
func new() (*gosearch, error) {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.index = *index.New()
	gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.depth = 2
	storage, err := storage.New()
	if err != nil {
		return &gs, err
	}
	gs.storage = storage
	return &gs, err
}

// Сканирует сайты и помещает данные в индекс и хранилище документов
func (s *gosearch) scan() (err error) {
	var data []crawler.Document
	for _, site := range s.sites {
		res, e := s.scanner.Scan(site, s.depth)
		if e != nil {
			return e
		}
		data = append(data, res...)
		s.index.Add(data)
	}
	s.data = data
	err = s.storage.Persist(s.data)
	if err != nil {
		log.Println(err)
		return
	}
	return err
}

// Ищет заданное слово по индексу и возвращает совпавщие документы
func (s *gosearch) search(req string) (docs []crawler.Document, err error) {
	for _, id := range s.index.Search(req) {
		d, e := s.document(id)
		if er, ok := e.(*notFoundErr); ok {
			if er.code != 0 {
				return docs, er
			}
		} else {
			return docs, er
		}
		docs = append(docs, d)
	}
	return docs, err
}

// Производить поиск документа по id в хранилище документов методом бинарного поиска
func (s *gosearch) document(id int) (crawler.Document, error) {
	var doc crawler.Document
	var err notFoundErr

	i := sort.Search(len(s.data), func(i int) bool { return s.data[i].ID >= id })
	if i < len(s.data) && s.data[i].ID == id {
		doc = s.data[i]
	} else {
		err.code = -1
		err.msg = "Doc not found"
	}
	return doc, &err
}

// Выводит конечный результат
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
