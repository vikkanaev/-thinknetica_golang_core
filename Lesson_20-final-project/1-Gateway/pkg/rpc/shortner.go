package rpc

import (
	"net/http"
	"strings"
)

type Shortner struct {
	Urls   string `yaml:"urls"`
	NewUrl string `yaml:"newUrl"`
	Url    string `yaml:"url"`
}

// Запрашивает все пары ссылка-сокращение из хранилища ссылок
func (r *RPC) Urls() (string, error) {
	data, err := get(r.config.Shortner.Urls)
	if err != nil {
		return "", err
	}
	return data, nil
}

// Запрашивает полную ссылку по короткой
func (r *RPC) Url(shortUrl string) (string, error) {
	link := r.config.Shortner.Url + "/" + shortUrl

	resp, err := get(link)
	if err != nil {
		return "", err
	}

	// удаляем кавычки, которые есть внутри строки
	data := strings.Replace(resp, "\"", "", -1)
	return data, nil
}

// Создает новое сокращение для данной ссылки
func (r *RPC) NewUrl(req *http.Request) (string, error) {
	link := r.config.Shortner.NewUrl
	resp, err := post(link, req)
	if err != nil {
		return "", err
	}
	return resp, nil
}
