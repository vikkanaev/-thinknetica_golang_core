package storage

import (
	"log"
	"sync"
)

type Storage struct {
	mu   sync.Mutex
	data Stat
	urls []string
}

// Структура для кранения статистики
type Stat struct {
	UrlsCount int32 `json:"urlsCount"`
	MaxUrlLen int   `json:"maxUrlLen"`
	AvgUrlLen int   `json:"avgUrlLen"`
}

func New() *Storage {
	s := Storage{
		data: Stat{
			UrlsCount: 0,
			MaxUrlLen: 0,
			AvgUrlLen: 0,
		},
		urls: make([]string, 0),
	}
	log.Println("Started storage with", s.data)
	return &s
}

// Возвращает статистику
func (s *Storage) Stat() Stat {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data
}

// Обработчик события "Новый url"
func (s *Storage) NewUrlHandler(url string) Stat {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := s.data
	data.UrlsCount = data.UrlsCount + 1
	s.urls = append(s.urls, url)

	t := 0
	max := 0
	for _, str := range s.urls {
		l := len([]rune(str))
		t = t + l
		if l > max {
			max = l
		}
	}
	avg := t / len(s.urls)

	data.AvgUrlLen = avg
	data.MaxUrlLen = max

	s.data = data
	return s.data
}

// Обработчик события "Очистить статистику"
func (s *Storage) PruneStatHandler() Stat {
	s.mu.Lock()
	defer s.mu.Unlock()

	stat := Stat{
		UrlsCount: 0,
		MaxUrlLen: 0,
		AvgUrlLen: 0,
	}

	s.data = stat
	s.urls = make([]string, 0)
	return s.data
}
