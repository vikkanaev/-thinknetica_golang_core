package rpc

import (
	"io/ioutil"
	"net/http"
)

type RPC struct {
	config Config
}

type Config struct {
	Shortner Shortner
	Cache    Cache
}

func New(c Config) *RPC {
	r := RPC{
		config: c,
	}
	return &r
}

// Выполняет GET запрос на указанный url
func get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Выполняет POST запрос на указанный url
func post(url string, r *http.Request) (string, error) {
	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
