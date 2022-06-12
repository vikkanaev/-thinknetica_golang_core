package queue

import (
	"encoding/json"
	"log"
	"sync"
	"thinknetica_golang_core/Lesson_19-queue/2-Analytics/pkg/storage"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	q       amqp.Queue
	storage *storage.Storage
}

// Сообщение для обмена между сервисами Shortner и Analytics
type Message struct {
	Event string
	Args  string
}

// Название эвентов поожим в константы
const (
	// Создание нового url
	newUrl = "NewUrl"
	// Обнулить статистику
	pruneStat = "Prune"
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

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue", err)
		return nil, err
	}

	queue := Queue{
		conn:    conn,
		ch:      ch,
		q:       q,
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
		queue.storage.NewUrlHandler(m.Args)
	case pruneStat:
		queue.storage.PruneStatHandler()
	default:
		log.Printf("Unknown event %v", m.Event)
	}
}

// Функция закрывает открытые ресурсы
func (queue *Queue) Close() {
	queue.conn.Close()
	queue.ch.Close()
}
