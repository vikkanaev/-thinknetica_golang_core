package storage

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	mu sync.Mutex

	client *redis.Client
	ttl    time.Duration
}

type Url struct {
	Long  string
	Short string
}

func New(conn string, ttl time.Duration) *Storage {
	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Password: "", // без пароля
		DB:       0,  // БД по умолчанию
	})
	s := Storage{
		client: client,
		ttl:    ttl,
	}
	return &s
}

// Возвращает полную ссылку по данному сокращению
func (s *Storage) Url(shortUrl string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := "urls:" + shortUrl
	url, err := s.client.Get(context.Background(), key).Result()
	log.Printf("Redis returns  %v", url)
	if err != nil {
		return ""
	}

	return url
}

// UpdateCache обновляет данные в кэше Redis.
func (s *Storage) UpdateCache(urls []Url) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range urls {
		log.Printf("Add to Redis %v", u)

		key := "urls:" + u.Short
		err := s.client.Set(context.Background(), key, u.Long, s.ttl).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

// Удаляет все ключи из Redis
func (s *Storage) PruneHandler() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("FlushAll in Redis")
	return s.client.FlushAll(context.Background()).Err()
}

// Закрываем соединение с Редисом
func (s *Storage) Close() error {
	return s.client.Close()
}
