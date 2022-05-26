package webapp

import (
	"fmt"
	"net/http"
	"thinknetica_golang_core/Lesson_12-web/pkg/crawler"
	"thinknetica_golang_core/Lesson_12-web/pkg/index"

	"github.com/gorilla/mux"
)

type Interface interface {
	Start(*index.Index, []crawler.Document)
}

type Webapp struct {
	addr string
	i    index.Index
	docs []crawler.Document
}

// Создает новый объект веб-сервера и возвращает указатель на него
func New() *Webapp {
	var w Webapp
	w.addr = ":80"
	return &w
}

// Запускает веб-сервер
func (wa *Webapp) Start(ind *index.Index, d []crawler.Document) {
	wa.i = *ind
	wa.docs = d
	mux := mux.NewRouter()
	mux.HandleFunc("/index", wa.indexHandler).Methods(http.MethodGet)
	mux.HandleFunc("/docs", wa.docsHandler).Methods(http.MethodGet)

	http.ListenAndServe(wa.addr, mux)
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><h2>%v</h2><div>%v</div></body></html>", title, data)
}

// форматирует массив строк в HTML-строку
func formatArr(arr []string) (res string) {
	for _, text := range arr {
		res = res + "<div>" + text + "</div>"
	}
	return res
}
