package api

import (
	"encoding/json"
	"net/http"
)

// Возвращает статистику
func (api *API) stat(w http.ResponseWriter, r *http.Request) {
	data := api.storage.Stat()
	body, err := json.Marshal(data)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err)
	}
	responseOk(w, body, http.StatusOK)
}
