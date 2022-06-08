package api

import (
	"encoding/json"
	"math"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// Набор символов короткого URL
	shortChars = "abcdefghijklmnopqrstuvwxyz123456789"
	// Длинна короткого URL
	urlLen = 6
)

var (
	// Максимально возможное число url для заданного набора чимволов и длинны короткой ссылки
	// При использовании 9 цифр и 26 букы имеем для длинны 6
	// (9+26)**6 = 1_838_265_625 (1.8 млрд) вариантов
	maxUrls = int(math.Pow(float64(len([]byte(shortChars))), urlLen))
)

// Возвращает список всех пар сокращение - ссылка
func (api *API) urls(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	responseOk(w, api.data, http.StatusOK)
}

// Сохраняет новую ссылку и возвращает для нее сокращение
func (api *API) newUrl(w http.ResponseWriter, r *http.Request) {
	// Не создаем новую запись, если мы достигли предела по сохраненным уникальным комбинациям
	if len(api.data) >= maxUrls {
		responseErr(w, http.StatusInternalServerError, "To many urls in memory")
	}

	var d struct{ Url string }
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		responseErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Генерируем случайный ключ и проверяем, что он не занят.
	// Если занят - повторяем заново
	// Этот алгоритм явно будет работать тем медленнее, чем ближе мы к максимальному числу записей
	shortUrl := ""
	api.mu.Lock()
	for {
		shortUrl = randSeq(urlLen)
		res := api.data[shortUrl]
		if res == "" {
			break
		}
	}
	api.data[shortUrl] = d.Url
	api.mu.Unlock()

	responseOk(w, shortUrl, http.StatusOK)
}

// Возвращает ссылку для данного сокращения
func (api *API) url(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	url := api.data[key]
	if url == "" {
		responseErr(w, http.StatusNotFound, nil)
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// Генерирует случайную последовательность заданной динны из фиксированного набора символов
func randSeq(n int) string {
	letters := []rune(shortChars)

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
