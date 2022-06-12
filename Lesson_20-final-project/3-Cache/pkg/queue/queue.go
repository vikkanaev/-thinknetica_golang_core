package queue

import (
	"encoding/json"
	"log"
	"sync"
	"thinknetica_golang_core/Lesson_20-final-project/3-Cache/pkg/storage"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	q       amqp.Queue
	name    string
	storage *storage.Storage
}

// Сообщение для обмена между сервисами Shortner и Analytics
type Message struct {
	Event    string
	LongUrl  string
	ShortUrl string
}

// Название эвентов поожим в константы
const (
	// Создание нового url
	newUrl = "NewUrl"
	// Обнулить статистику
	pruneStat = "PruneStat"
)

// Подключается к очереди сообщений
func New(cred string, name string, s *storage.Storage) (*Queue, error) {
	conn, err := amqp.Dial(cred)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel", err)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		name,     // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Println("Failed to declare an exchange", err)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue", err)
		return nil, err
	}

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		name,   // exchange
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to bind a queue", err)
		return nil, err
	}

	queue := Queue{
		conn:    conn,
		ch:      ch,
		q:       q,
		name:    name,
		storage: s,
	}
	return &queue, nil
}

// Запуск прослушивания очереди
func (queue *Queue) Consume() error {
	msgs, err := queue.ch.Consume(
		queue.q.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Println("Failed to declare a queue", err)
		return err
	}

	var wg sync.WaitGroup
	var m Message

	wg.Add(1)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &m)
			if err == nil {
				queue.handleEvent(m)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	wg.Wait()
	return nil
}

// Обработчик для входящих эвентов
func (queue *Queue) handleEvent(m Message) {
	switch m.Event {
	case newUrl:
		url := storage.Url{Long: m.LongUrl, Short: m.ShortUrl}
		queue.storage.UpdateCache([]storage.Url{url})
	case pruneStat:
		queue.storage.PruneHandler()
		log.Printf("pruneStat event with %v", m.LongUrl)
	default:
		log.Printf("Unknown event %v", m.Event)
	}
}

// Функция закрывает открытые ресурсы
func (queue *Queue) Close() {
	queue.conn.Close()
	queue.ch.Close()
}
