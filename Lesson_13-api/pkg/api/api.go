package api

import (
	"encoding/json"
	"net/http"
	"sync"
	"thinknetica_golang_core/Lesson_13-api/pkg/crawler"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
	data   []crawler.Document
	mu     sync.Mutex
}

// New создаёт объект API.
func New(r *mux.Router) *API {
	api := API{
		router: r,
	}
	return &api
}

// Endpoints регистрирует конечные точки API.
func (api *API) Endpoints(d []crawler.Document) {
	api.data = d
	api.router.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.doc).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.delDoc).Methods(http.MethodDelete, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs", api.newDoc).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.updateDoc).Methods(http.MethodPatch, http.MethodOptions)
}

// Возвращает успешный ответ
func responseOk(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Возвращает сообщение об ошибке
func responseErr(w http.ResponseWriter, code int, v any) {
	w.WriteHeader(code)
	if v != nil {
		err := json.NewEncoder(w).Encode(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
