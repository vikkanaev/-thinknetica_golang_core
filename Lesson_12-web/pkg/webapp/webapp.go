package webapp

import (
	"fmt"
	"net/http"
	"thinknetica_golang_core/Lesson_12-web/pkg/crawler"
	"thinknetica_golang_core/Lesson_12-web/pkg/index"

	"github.com/gorilla/mux"
)

type Webapp struct {
	i    index.Index
	docs []crawler.Document
}

// Запускает веб-сервер
func New(r *mux.Router, ind *index.Index, d []crawler.Document) {
	var wa Webapp
	wa.i = *ind
	wa.docs = d
	r.HandleFunc("/index", wa.indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/docs", wa.docsHandler).Methods(http.MethodGet)
}

// HTTP-обработчик для эндпоинта /index
func (wa *Webapp) indexHandler(w http.ResponseWriter, r *http.Request) {
	responser(w, "Index", formatArr(wa.i.ToStrs()))
}

// HTTP-обработчик для эндпоинта /docs
func (wa *Webapp) docsHandler(w http.ResponseWriter, r *http.Request) {
	var arr []string
	for _, d := range wa.docs {
		arr = append(arr, d.ToStr())
	}
	responser(w, "Docs", formatArr(arr))
}

// Формирует ответ пользователю
func responser(w http.ResponseWriter, title string, data string) {
	fmt.Fprintf(w, "<html><body><h2>%v</h2><div>%v</div></body></html>", title, data)
}

// форматирует массив строк в HTML-строку
func formatArr(arr []string) (res string) {
	for _, text := range arr {
		res = res + "<div>" + text + "</div>"
	}
	return res
}
