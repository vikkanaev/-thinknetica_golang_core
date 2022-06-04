package main

import (
	"context"
	"fmt"
	"thinknetica_golang_core/Lesson_16-DB-App/pkg/storage"
	"thinknetica_golang_core/Lesson_16-DB-App/pkg/storage/postgres"
)

type app struct {
	storage storage.Interface
	context context.Context
}

func new() *app {
	app := app{}
	app.context = context.Background()
	app.storage = postgres.New(app.context)
	return &app
}

func main() {
	a := new()
	// Не забываем очищать ресурсы и закрывать соединения
	defer a.storage.Close()

	// Добавление фильмов
	films := []storage.Film{{Title: "Generation Pi", Year: 1999, StudioId: 1}}
	err := a.storage.AddFilms(films)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Получение списка фильмов.
	data, err := a.storage.Films(0)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v\n", data)

	// Удаление фильма
	err = a.storage.DelFilm(1)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Изменения параметров фильма
	err = a.storage.UpdateFilm(2, storage.Film{Title: "Okay", Year: 2015, StudioId: 1})
	if err != nil {
		fmt.Print(err)
		return
	}
}
