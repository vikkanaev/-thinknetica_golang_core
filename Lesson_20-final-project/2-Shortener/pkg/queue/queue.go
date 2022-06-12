package queue

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	name string
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

	queue := Queue{
		conn: conn,
		ch:   ch,
		name: name,
	}
	return &queue, nil
}

// Отправка в аналитику событие "Новый url"
func (queue *Queue) NewUrl(longUrl string, shortUrl string) error {
	m := Message{Event: newUrl, LongUrl: longUrl, ShortUrl: shortUrl}
	return queue.publish(m)
}

// Отправка в аналитику событие "Обнули статистику"
func (queue *Queue) PruneStat() error {
	m := Message{Event: pruneStat, LongUrl: "", ShortUrl: ""}
	return queue.publish(m)
}

// Публикация события в очередь
func (queue *Queue) publish(m Message) error {
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = queue.ch.Publish(
		queue.name, // exchange
		"",         // routing key
		false,      // mandatory
		false,      // immediate
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
