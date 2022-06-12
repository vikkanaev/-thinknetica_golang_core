package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Возвращает список всех пар сокращение - ссылка
func (api *API) urls(w http.ResponseWriter, r *http.Request) {
	data, err := api.storage.Urls()
	if err != nil {
		responseErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOk(w, data, http.StatusOK)
}

// Сохраняет новую ссылку и возвращает для нее сокращение
func (api *API) newUrl(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Url string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if d.Url == "" {
		responseErr(w, http.StatusUnprocessableEntity, "url can not be empty")
		return
	}

	shortUrl, err := api.storage.NewUrl(d.Url)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	api.queue.NewUrl(d.Url, shortUrl)

	responseOk(w, shortUrl, http.StatusOK)
}

// Возвращает ссылку для данного сокращения
func (api *API) url(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	url, err := api.storage.Url(key)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if url == "" {
		responseErr(w, http.StatusNotFound, nil)
		return
	}
	responseOk(w, url, http.StatusOK)
}
