package main

import (
	"math/rand"
	"sync"
	"time"
)

func player(in <-chan string, out chan<- string, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	var m = map[string]string{
		"Ping": "Pong",
		"Pong": "Ping",
	}

	for res := range in {
		if res == "Stop" {
			println("Player", num, "lose!")
			return
		}

		rep := goal(m[res])
		println("===>>>  Player", num, "see:", res, "reply", rep)
		// Sleep добавлено исключительно для красоты отображение. Без него все работает так же
		time.Sleep(time.Second)
		out <- rep
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

func start(ch1 chan<- string, ch2 chan<- string) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	if (rand.Intn(max-min+1) + min) > 50 {
		ch1 <- "Ping"
	} else {
		ch2 <- "Ping"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go player(ch1, ch2, 1, &wg)
	go player(ch2, ch1, 2, &wg)

	start(ch1, ch2)
	wg.Wait()
}
