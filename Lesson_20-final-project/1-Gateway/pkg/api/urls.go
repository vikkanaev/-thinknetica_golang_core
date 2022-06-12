package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Возвращает список всех пар сокращение - ссылка
func (api *API) urls(w http.ResponseWriter, r *http.Request) {
	data, err := api.rpc.Urls()
	if err != nil {
		responseErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOk(w, data, http.StatusOK)
}

// Сохраняет новую ссылку и возвращает для нее сокращение
func (api *API) newUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl, err := api.rpc.NewUrl(r)
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
	log.Println("vars returns", key)

	url, err := api.rpc.Url(key)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	log.Println("rpc.Url returns", url)

	if url == "" {
		responseErr(w, http.StatusNotFound, nil)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}
