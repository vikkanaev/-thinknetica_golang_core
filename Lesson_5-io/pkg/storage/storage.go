package storage

import (
	"encoding/json"
	"io"
	"os"
	"thinknetica_golang_core/Lesson_5-io/pkg/crawler"
)

type Interface interface {
	Retrieve() ([]crawler.Document, error)
	Persist([]crawler.Document) error
}

type storage struct {
	file io.ReadWriter
}

// Создает новый объект хранилища и возвращает на него указатель
func New() (*storage, error) {
	var st storage
	var err error

	f, err := os.OpenFile("./storage.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return &st, err
	}

	st.file = f
	return &st, err
}

// Считывает раннее сохраненные данные из постоянного хранилиша
func (st *storage) Retrieve() ([]crawler.Document, error) {
	var docs []crawler.Document

	docsJson, err := get(st.file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(docsJson), &docs)

	return docs, err
}

// Сохраният данные в постоянное хранилище
// TODO сделать что бы персист удалял старые данные перед записью новых
func (st *storage) Persist(docs []crawler.Document) error {
	b, err := json.Marshal(docs)
	if err != nil {
		return err
	}

	err = store(st.file, b)
	return err
}

// Процедура записи данных
func store(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}

// Процедура зчтения данных
func get(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
