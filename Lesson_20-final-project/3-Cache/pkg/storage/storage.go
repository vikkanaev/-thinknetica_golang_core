package storage

import (
	"sync"
)

type Storage struct {
	mu sync.Mutex
	// data Stat
	// urls []string
}

type Url struct {
	Long  string `bson:"long"`
	Short string `bson:"short"`
}

func New() *Storage {
	s := Storage{}
	return &s
}

func (s *Storage) Url(shortUrl string) string {
	var str string
	s.mu.Lock()
	defer s.mu.Unlock()
	if shortUrl == "123" {
		return "http://google.com"
	}

	return str
}
