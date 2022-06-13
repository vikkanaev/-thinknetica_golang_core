package api

import (
	"encoding/json"
	"net/http"

	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/cache"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/shortner"

	"github.com/gorilla/mux"
)

type API struct {
	router   *mux.Router
	shortner *shortner.Shortner
	cache    *cache.Cache
}

// New создаёт объект API.
func New(r *mux.Router, s *shortner.Shortner, c *cache.Cache) *API {
	api := API{
		router:   r,
		shortner: s,
		cache:    c,
	}
	return &api
}

// Endpoints регистрирует конечные точки API.
func (api *API) Endpoints() {
	api.router.HandleFunc("/api/v1/url", api.newUrl).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/v1/urls/{key}", api.url).Methods(http.MethodGet, http.MethodOptions)
}

// Возвращает успешный ответ
func responseOk(w http.ResponseWriter, v string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(v))
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
