package storage

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type url struct {
	Long  string `bson:"long"`
	Short string `bson:"short"`
}

type Storage struct {
	mu sync.Mutex

	databaseName   string
	collectionName string
	client         *mongo.Client
}

const (
	shortChars = "abcdefghijklmnopqrstuvwxyz123456789" // Набор символов короткого URL
	urlLen     = 6                                     // Длинна короткого URL
)

var (
	// Максимально возможное число url для заданного набора чимволов и длинны короткой ссылки
	// При использовании 9 цифр и 26 букы имеем для длинны 6
	// (9+26)**6 = 1_838_265_625 (1.8 млрд) вариантов
	maxUrls = int(math.Pow(float64(len([]byte(shortChars))), urlLen))
)

func New(conn string, dbName string, collName string) (*Storage, error) {
	// подключение к СУБД MongoDB
	mongoOpts := options.Client().ApplyURI(conn)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}
	s := Storage{
		client:         client,
		databaseName:   dbName,
		collectionName: collName,
	}
	return &s, nil
}

// Возвращает все документы из базы
func (s *Storage) Urls() ([]url, error) {
	filter := bson.D{}
	return s.find(filter)
}

// Создает новое сокращение для данного url
func (s *Storage) NewUrl(longUrl string) (string, error) {
	// Не создаем новую запись, если мы достигли предела по сохраненным уникальным комбинациям
	count, err := s.countDocs(context.Background())
	if err != nil {
		return "", err
	}

	if count >= int64(maxUrls) {
		err := errors.New("to many urls in memory")
		return "", err
	}

	// Генерируем случайный ключ и проверяем, что он не занят.
	// Если занят - повторяем заново
	// Этот алгоритм явно будет работать тем медленнее, чем ближе мы к максимальному числу записей
	shortUrl := ""
	for {
		shortUrl = randSeq(urlLen)
		filter := bson.D{primitive.E{Key: "short", Value: shortUrl}}
		resp, err := s.findOne(filter)
		if err != nil {
			return "", err
		}
		if resp.Short == "" {
			break
		}
	}
	ctx := context.Background()
	data := []url{{Long: longUrl, Short: shortUrl}}
	s.insertUrls(ctx, data)

	return shortUrl, nil
}

// Возвращает полный url для заданного сокращения
func (s *Storage) Url(shortUrl string) (string, error) {
	filter := bson.D{primitive.E{Key: "short", Value: shortUrl}}
	u, err := s.findOne(filter)
	if err != nil {
		return "", err
	}

	return u.Long, nil
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

// Возвращает массив документов из Монго
func (s *Storage) find(filter bson.D) ([]url, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	collection := s.client.Database(s.databaseName).Collection(s.collectionName)
	ctx := context.Background()

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var data []url
	for cur.Next(ctx) {
		var l url
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}

	return data, cur.Err()
}

// Возвращает один документ из Монги
func (s *Storage) findOne(filter bson.D) (u url, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	collection := s.client.Database(s.databaseName).Collection(s.collectionName)
	ctx := context.Background()

	err = collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection.
		if err == mongo.ErrNoDocuments {
			return u, nil
		}
	}
	return u, nil
}

// Возвращает число сохраненных документов
func (s *Storage) countDocs(ctx context.Context) (int64, error) {
	collection := s.client.Database(s.databaseName).Collection(s.collectionName)
	s.mu.Lock()
	defer s.mu.Unlock()

	return collection.EstimatedDocumentCount(ctx)
}

// Вставляет в Монгу массив документов.
func (s *Storage) insertUrls(ctx context.Context, data []url) error {
	collection := s.client.Database(s.databaseName).Collection(s.collectionName)
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, doc := range data {
		_, err := collection.InsertOne(ctx, doc)
		if err != nil {
			return err
		}
	}
	return nil
}
