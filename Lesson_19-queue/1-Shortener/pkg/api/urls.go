package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Возвращает список всех пар сокращение - ссылка
func (api *API) urls(w http.ResponseWriter, r *http.Request) {
	responseOk(w, api.storage.Urls(), http.StatusOK)
}

// Сохраняет новую ссылку и возвращает для нее сокращение
func (api *API) newUrl(w http.ResponseWriter, r *http.Request) {
	var d struct{ Url string }
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	shortUrl, err := api.storage.NewUrl(d.Url)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	api.queue.NewUrl(d.Url)

	responseOk(w, shortUrl, http.StatusOK)
}

// Возвращает ссылку для данного сокращения
func (api *API) url(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	url := api.storage.Url(key)
	if url == "" {
		responseErr(w, http.StatusNotFound, nil)
	}
	// Переделал на Redirect
	http.Redirect(w, r, url, http.StatusSeeOther)
}
