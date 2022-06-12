package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) url(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	url := api.storage.Url(key)
	log.Println("storage.Url", url)
	if url == "" {
		responseErr(w, http.StatusNotFound, "")
		return
	}
	responseOk(w, url, http.StatusOK)
}
