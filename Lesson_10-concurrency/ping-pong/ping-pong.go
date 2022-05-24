package main

import (
	"math/rand"
	"sync"
	"time"
)

func player(ch chan string, num int, wg *sync.WaitGroup) {
	defer wg.Done()

	var m = map[string]string{
		"Ping": "Pong",
		"Pong": "Ping",
	}

	for res := range ch {
		if res == "Stop" {
			println("Player", num, "lose!")
			close(ch)
			return
		}

		rep := goal(m[res])
		println("===>>>  Player", num, "see:", res, "reply", rep)
		// Sleep добавлено исключительно для красоты отображение. Без него все работает так же
		time.Sleep(time.Second)
		ch <- rep
	}
	println("Player", num, "win!")
}

func goal(str string) string {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	if (rand.Intn(max-min+1) + min) > 80 {
		return "Stop"
	}
	return str
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go player(ch, 1, &wg)
	go player(ch, 2, &wg)

	ch <- "Ping"
	wg.Wait()
}
