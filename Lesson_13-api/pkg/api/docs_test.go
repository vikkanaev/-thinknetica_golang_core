package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"thinknetica_golang_core/Lesson_12-web/pkg/crawler"
)

func TestAPI_docs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	//=========================================================

	want := "[{\"ID\":1,\"URL\":\"link1\",\"Title\":\"test\"},{\"ID\":2,\"URL\":\"link2\",\"Title\":\"words\"}]\n"
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}
}

func TestAPI_doc(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs/1", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	//=========================================================

	want := "{\"ID\":1,\"URL\":\"link1\",\"Title\":\"test\"}\n"
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//=========================================================

	req = httptest.NewRequest(http.MethodGet, "/api/v1/docs/100", nil)
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusNotFound) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusNotFound)
	}
	t.Log("Response: ", rr.Body)

	//=========================================================

	req = httptest.NewRequest(http.MethodGet, "/api/v1/docs/foobar", nil)
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusUnprocessableEntity) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusUnprocessableEntity)
	}
	t.Log("Response: ", rr.Body)

	//=========================================================

	want = "\"Invalid id foobar\"\n"
	data, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got = string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}
}

func TestAPI_delDoc(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/docs/1", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	//=========================================================

	want := "null\n"
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//=========================================================
	req = httptest.NewRequest(http.MethodDelete, "/api/v1/docs/100500", nil)
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusNotFound) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)
}

func TestAPI_newDoc(t *testing.T) {
	doc := crawler.Document{
		ID:    3,
		URL:   "http://ya.ru",
		Title: "Yandex",
	}
	payload, _ := json.Marshal(doc)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusCreated) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusCreated)
	}
	t.Log("Response: ", rr.Body)

	//=========================================================

	want := "{\"ID\":3,\"URL\":\"http://ya.ru\",\"Title\":\"Yandex\"}\n"
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//=========================================================

	j := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	payload, _ = json.Marshal(j)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(payload))
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusUnprocessableEntity) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusUnprocessableEntity)
	}
	t.Log("Response: ", rr.Body)
}

func TestAPI_updateDoc(t *testing.T) {
	doc := crawler.Document{
		ID:    1,
		URL:   "http://ya.ru",
		Title: "Yandex",
	}
	payload, _ := json.Marshal(doc)
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/docs/1", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	//=========================================================

	want := "{\"ID\":1,\"URL\":\"http://ya.ru\",\"Title\":\"Yandex\"}\n"
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//=========================================================

	j := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	payload, _ = json.Marshal(j)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(payload))
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusUnprocessableEntity) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusUnprocessableEntity)
	}
	t.Log("Response: ", rr.Body)
}
