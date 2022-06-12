package queue

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
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
func New(cred string, name string) (*Queue, error) {
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
		conn: conn,
		ch:   ch,
		q:    q,
	}
	return &queue, nil
}

// Отправка в аналитику событие "Новый url"
func (queue *Queue) NewUrl(url string) error {
	m := Message{Event: newUrl, Args: url}
	return queue.publish(m)
}

// Отправка в аналитику событие "Обнули статистику"
func (queue *Queue) PruneStat() error {
	m := Message{Event: pruneStat, Args: ""}
	return queue.publish(m)
}

// Публикация события в очередь
func (queue *Queue) publish(m Message) error {
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = queue.ch.Publish(
		"",           // exchange
		queue.q.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s\n", body)
	return nil
}

// Функция закрывает открытые ресурсы
func (queue *Queue) Close() {
	queue.conn.Close()
	queue.ch.Close()
}
