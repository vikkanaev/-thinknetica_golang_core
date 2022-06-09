package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"thinknetica_golang_core/Lesson_19-queue/2-Analytics/pkg/queue"
	"thinknetica_golang_core/Lesson_19-queue/2-Analytics/pkg/storage"

	"time"

	"github.com/gorilla/mux"
)

type API struct {
	router  *mux.Router
	queue   *queue.Queue
	storage *storage.Storage
}

// New создаёт объект API.
func New(r *mux.Router, q *queue.Queue, s *storage.Storage) *API {
	rand.Seed(time.Now().UnixNano())

	api := API{
		router:  r,
		queue:   q,
		storage: s,
	}
	return &api
}

// Endpoints регистрирует конечные точки API.
func (api *API) Endpoints() {
	api.router.HandleFunc("/api/v1/stat", api.stat).Methods(http.MethodGet, http.MethodOptions)
}

// Возвращает успешный ответ
func responseOk(w http.ResponseWriter, v []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(v)
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
