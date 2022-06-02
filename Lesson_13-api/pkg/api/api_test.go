package api

import (
	"os"
	"testing"
	"thinknetica_golang_core/Lesson_13-api/pkg/crawler"

	"github.com/gorilla/mux"
)

var api *API

func TestMain(m *testing.M) {
	router := mux.NewRouter()
	api = New(router)
	docs := []crawler.Document{{ID: 1, URL: "link1", Title: "test"}, {ID: 2, URL: "link2", Title: "words"}}
	api.Endpoints(docs)
	os.Exit(m.Run())
}
