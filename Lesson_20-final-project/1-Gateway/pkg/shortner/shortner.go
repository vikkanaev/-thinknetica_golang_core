package shortner

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Структура хранит ссылки на эндпоинты сервиса Shortner
type Shortner struct {
	newUrlEndpoint string // Создать сокращение для данного url
	urlEndpoint    string // Получить полный url для данного сокращения
}

// Принимает адреса эндпоинтов микросервиса Shortner
func New(newUrlEndpoint string, urlEndpoint string) *Shortner {
	s := Shortner{
		newUrlEndpoint: newUrlEndpoint,
		urlEndpoint:    urlEndpoint,
	}
	return &s
}

// Запрашивает полную ссылку по короткой
func (s *Shortner) Url(shortUrl string) (string, error) {
	link := s.urlEndpoint + "/" + shortUrl

	resp, err := get(link)
	if err != nil {
		return "", err
	}

	// удаляем кавычки, которые есть внутри строки
	data := strings.Replace(string(resp), "\"", "", -1)
	return data, nil
}

// Создает новое сокращение для данной ссылки
func (s *Shortner) NewUrl(req *http.Request) (string, error) {
	resp, err := post(s.newUrlEndpoint, req)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// Выполняет GET запрос на указанный url
func get(url string) ([]byte, error) {
	var body []byte
	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Выполняет POST запрос на указанный url
func post(url string, r *http.Request) ([]byte, error) {
	var body []byte
	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		return body, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
