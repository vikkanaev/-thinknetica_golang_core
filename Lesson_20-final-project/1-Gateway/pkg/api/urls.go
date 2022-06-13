package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Сохраняет новую ссылку и возвращает для нее сокращение
func (api *API) newUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl, err := api.shortner.NewUrl(r)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	responseOk(w, shortUrl, http.StatusOK)
}

// Возвращает ссылку для данного сокращения
func (api *API) url(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	// Проверяем в кеше
	// Ошибку не обрабатываем потому, что даже если кеш прилег мы сходим в сторедж
	url, _ := api.cache.Url(key)

	if url != "" {
		// Печатаем 😎 в лог
		log.Printf("%s Url %v got from cache", unquoteCodePoint("\\U0001f60e"), key)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	// Если не нашли в кеше проверяем в хранилище
	url, err := api.shortner.Url(key)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	// Печатаем 👿 в лог
	log.Printf("%s Url %v got from storage", unquoteCodePoint("\\U0001f47f"), key)

	if url == "" {
		responseErr(w, http.StatusNotFound, nil)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// Выводит Юникод emoji
func unquoteCodePoint(s string) string {
	r, _ := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	return string(r)
}
