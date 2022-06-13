package cache

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Структура хранит ссылки на эндпоинты сервиса Cache
type Cache struct {
	urlEndpoint string // Получить полный url для данного сокращения
}

// Принимает адреса эндпоинтов микросервиса Cache
func New(urlEndpoint string) *Cache {
	s := Cache{
		urlEndpoint: urlEndpoint,
	}
	return &s
}

// Запрашивает полную ссылку по короткой из кеша
func (c *Cache) Url(shortUrl string) (string, error) {
	link := c.urlEndpoint + "/" + shortUrl

	resp, err := get(link)
	if err != nil {
		return "", err
	}

	// удаляем кавычки, которые есть внутри строки
	data := strings.Replace(string(resp), "\"", "", -1)

	if data == "" {
		return data, errors.New("url not found")
	}

	return data, nil
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
