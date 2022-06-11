package storage

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"sync"
)

type Urls map[string]string

type Storage struct {
	mu   sync.Mutex
	data Urls
}

const (
	// Набор символов короткого URL
	shortChars = "abcdefghijklmnopqrstuvwxyz123456789"
	// Длинна короткого URL
	urlLen = 6
)

var (
	// Максимально возможное число url для заданного набора чимволов и длинны короткой ссылки
	// При использовании 9 цифр и 26 букы имеем для длинны 6
	// (9+26)**6 = 1_838_265_625 (1.8 млрд) вариантов
	maxUrls = int(math.Pow(float64(len([]byte(shortChars))), urlLen))
)

func New() *Storage {
	s := Storage{
		data: make(map[string]string),
	}
	log.Println("Started storage with", s.data)
	return &s
}

// Возвращает статистику
func (s *Storage) Urls() Urls {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data
}

func (s *Storage) NewUrl(url string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Не создаем новую запись, если мы достигли предела по сохраненным уникальным комбинациям
	if len(s.data) >= maxUrls {
		err := errors.New("To many urls in memory")
		return "", err
	}

	// Генерируем случайный ключ и проверяем, что он не занят.
	// Если занят - повторяем заново
	// Этот алгоритм явно будет работать тем медленнее, чем ближе мы к максимальному числу записей
	shortUrl := ""
	for {
		shortUrl = randSeq(urlLen)
		res := s.data[shortUrl]
		if res == "" {
			break
		}
	}
	s.data[shortUrl] = url

	return shortUrl, nil
}

func (s *Storage) Url(shortUrl string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data[shortUrl]
}

// Генерирует случайную последовательность заданной динны из фиксированного набора символов
func randSeq(n int) string {
	letters := []rune(shortChars)

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
