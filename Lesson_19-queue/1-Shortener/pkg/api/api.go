package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"thinknetica_golang_core/Lesson_19-queue/1-Shortener/pkg/queue"
	"time"

	"github.com/gorilla/mux"
)

type API struct {
	mu   sync.Mutex
	data map[string]string

	router *mux.Router
	queue  *queue.Queue
}

// New создаёт объект API.
func New(r *mux.Router, q *queue.Queue) *API {
	rand.Seed(time.Now().UnixNano())
	// обнуляем статистику при старте сервиса
	q.PruneStat()

	api := API{
		router: r,
		queue:  q,
		data:   make(map[string]string),
	}
	return &api
}

// Endpoints регистрирует конечные точки API.
func (api *API) Endpoints() {
	api.router.HandleFunc("/api/v1/urls", api.urls).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/url", api.newUrl).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/v1/urls/{key}", api.url).Methods(http.MethodGet, http.MethodOptions)
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
