package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"thinknetica_golang_core/Lesson_13-api/pkg/crawler"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
)

// Возвращает список всех документов
func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	responseOk(w, api.data, http.StatusOK)
}

// Возвращает заданный документ
func (api *API) doc(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, fmt.Sprintf("Invalid id %v", vars["id"]))
		return
	}

	_, doc, err := api.takeDoc(int(id))
	if err != nil {
		responseErr(w, http.StatusNotFound, "")
		return
	}

	responseOk(w, doc, http.StatusOK)
}

// Удаляет заданный документ
func (api *API) delDoc(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, fmt.Sprintf("Invalid id %v", vars["id"]))
		return
	}

	idx, _, err := api.takeDoc(int(id))
	if err != nil {
		responseErr(w, http.StatusNotFound, "")
		return
	}

	api.data = append(api.data[:idx], api.data[idx+1:]...)
	responseOk(w, nil, http.StatusOK)
}

// Создает новый документ
func (api *API) newDoc(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	var d crawler.Document

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	api.data = append(api.data, d)
	responseOk(w, d, http.StatusCreated)
}

// Изменяет заданный документ
func (api *API) updateDoc(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	vars := mux.Vars(r)
	var d crawler.Document

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, fmt.Sprintf("Invalid id %v", vars["id"]))
		return
	}

	idx, _, err := api.takeDoc(int(id))
	if err != nil {
		responseErr(w, http.StatusNotFound, "")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		responseErr(w, http.StatusNotFound, err.Error())
		return
	}

	api.data[idx] = d
	responseOk(w, d, http.StatusOK)
}

// Возвращает позицию документа в массике и сам документ
func (api *API) takeDoc(id int) (idx int, doc crawler.Document, err error) {
	idx = slices.IndexFunc(api.data, func(d crawler.Document) bool { return d.ID == int(id) })
	if idx < 0 {
		err := errors.New("NOT FOUND")
		return idx, doc, err
	}

	doc = api.data[idx]
	return idx, doc, err
}
