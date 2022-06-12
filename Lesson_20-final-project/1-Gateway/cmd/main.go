package main

import (
	"fmt"
	"net/http"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/api"
	"thinknetica_golang_core/Lesson_20-final-project/1-Gateway/pkg/rpc"

	"github.com/gorilla/mux"
)

const (
	webAddr  = ":3000"
	confFile = "./Lesson_20-final-project/1-Gateway/conf.yaml"
)

func main() {
	router := mux.NewRouter()
	// Читаем конфиг из yaml файла
	config, err := rpc.ReadConfig(confFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc := rpc.New(config)

	api := api.New(router, rpc)
	api.Endpoints()
	http.ListenAndServe(webAddr, router)
}
