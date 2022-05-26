package webapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"thinknetica_golang_core/Lesson_12-web/pkg/crawler"
	"thinknetica_golang_core/Lesson_12-web/pkg/index"

	"github.com/gorilla/mux"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	// создаем документы, индекс и производим индексиование
	docs := []crawler.Document{{ID: 1, Title: "test"}}
	i := *index.New()
	i.Add(docs)

	// создаем веб-сервер
	var wa Webapp
	wa.docs = docs
	wa.i = i

	// создаем тестовый маршрутизатор
	testMux = mux.NewRouter()
	testMux.HandleFunc("/docs", wa.docsHandler).Methods(http.MethodGet)
	testMux.HandleFunc("/index", wa.indexHandler).Methods(http.MethodGet)
	m.Run()
}

func TestWebapp_docsHandler(t *testing.T) {
	want := "<html><body><h2>Docs</h2><div><div>1  =>  test\n</div></div></body></html>"

	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("content-type", "plain/text")

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	//=========================================================

	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}
}

func TestWebapp_indexHandler(t *testing.T) {
	want := "<html><body><h2>Index</h2><div><div>test  =>  [1]\n</div></div></body></html>"

	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Add("content-type", "plain/text")

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	//=========================================================

	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}
}
