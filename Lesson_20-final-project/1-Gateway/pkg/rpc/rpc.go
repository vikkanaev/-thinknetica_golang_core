package rpc

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
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

// Читает конфиг из yaml-файла
func ReadConfig(fileName string) (Config, error) {
	config := Config{}

	yfile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
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
