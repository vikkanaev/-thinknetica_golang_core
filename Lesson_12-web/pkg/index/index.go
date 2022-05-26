package index

import (
	"fmt"
	"strings"
	"thinknetica_golang_core/Lesson_12-web/pkg/crawler"
)

// Index - индекс на основе хэш-таблицы.
type Index struct {
	data map[string][]int
}

// New - конструктор.
func New() *Index {
	var i Index
	i.data = make(map[string][]int)
	return &i
}

// Add добавляет данные из переданных документов в индекс.
//
// Сначала происходит выделение лексем как ключей словаря из данных документа.
// Потом проверяется наличие номера документа в занчении словаря для лексемы.
// Если номер документа не найден, то он добавляется в значение словаря.
func (i *Index) Add(docs []crawler.Document) {
	for _, doc := range docs {
		for _, word := range tokens(doc.Title) {
			if !exists(i.data[word], doc.ID) {
				i.data[word] = append(i.data[word], doc.ID)
			}
		}
	}
}

// Разделение строки на лексемы.
func tokens(s string) []string {
	words := strings.Split(s, " ")
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return words
}

// Проверка наличия элемента в массиве.
func exists(ids []int, item int) bool {
	for _, id := range ids {
		if id == item {
			return true
		}
	}
	return false
}

// Search возвращает номера документов, где встречается данная лексема.
func (i *Index) Search(req string) []int {
	return i.data[req]
}

// Форматирует индех в массив строк
func (i *Index) ToStrs() (res []string) {
	for k, v := range i.data {
		res = append(res, fmt.Sprintln(k, " => ", v))
	}
	return res
}
