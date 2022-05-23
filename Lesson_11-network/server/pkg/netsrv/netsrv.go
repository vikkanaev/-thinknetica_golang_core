package netsrv

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"thinknetica_golang_core/Lesson_11-network/server/pkg/crawler"
)

type srv struct {
	data     []crawler.Document
	listener net.Listener
}
type Interface interface {
	Start([]crawler.Document)
}

// Создает новый web сервер
func New() *srv {
	s := srv{}
	var err error

	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		log.Println(err)
		return &s
	}

	s.listener = listener
	return &s
}

// Запускает на web сервере обработку клиентских подключений
func (s *srv) Start(data []crawler.Document) {
	s.data = data
	log.Println("Ready for clients")

	// цикл обработки клиентских подключений
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go s.handler(conn)
	}
}

// Обрабатывает подключение. Принимает поисковые запросы клиента и возврашает найденые результаты
func (s *srv) handler(conn net.Conn) {
	defer conn.Close()

	for {
		reply(conn, []string{"Enter new query or type 'exit' for finish"})
		q, err := readQuery(conn)
		if err != nil {
			return
		}

		if q == "exit" {
			reply(conn, []string{"Goodbuy!"})
			log.Println("Close session")
			return
		}
		rep := search(q, s.data)
		reply(conn, []string{"Results for: ", q})
		reply(conn, rep)

	}
}

// Читает запросы пользователя
func readQuery(conn net.Conn) (res string, err error) {
	r := bufio.NewReader(conn)
	msg, _, err := r.ReadLine()
	if err != nil {
		return res, err
	}
	res = strings.ToLower(string(msg))
	return res, err
}

// Отправляет ответы пользователю
func reply(conn net.Conn, text []string) {
	for _, str := range text {
		msg := []byte(str)
		_, err := conn.Write(msg)
		if err != nil {
			return
		}
	}
	_, err := conn.Write([]byte("\n"))
	if err != nil {
		return
	}
}

// Производит поиск по индексу для конкретного запроса
func search(req string, data []crawler.Document) (res []string) {
	for _, d := range data {
		if strings.Contains(strings.ToLower(d.Title), req) {
			res = append(res, fmt.Sprintf("Document: '%s' (%s)\n", d.Title, d.URL))
		}
	}
	if len(res) == 0 {
		res = append(res, "Nothing found")
	}
	return res
}
