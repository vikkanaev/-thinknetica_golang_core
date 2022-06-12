package rpc

import "strings"

type Cache struct {
	Url string `yaml:"url"`
}

// Запрашивает полную ссылку по короткой из кеша
func (r *RPC) CachedUrl(shortUrl string) (string, error) {
	link := r.config.Cache.Url + "/" + shortUrl

	resp, err := get(link)
	if err != nil {
		return "", err
	}

	// удаляем кавычки, которые есть внутри строки
	data := strings.Replace(resp, "\"", "", -1)
	return data, nil
}
