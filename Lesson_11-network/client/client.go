package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		return
	}
	defer conn.Close()

	go reader(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		req := []byte(text)

		_, err := conn.Write(req)
		if err != nil {
			return
		}

		if text == "exit\n" {
			fmt.Println("See you on the other side!")
			return
		}
	}
}

// Читает ответы от сервера
func reader(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		fmt.Print(string(msg), "\n", "> ")
	}
}
