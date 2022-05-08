package main

import (
	"fmt"
	"log"
	"thinknetica_golang_core/Lesson_6-OOP/task_1/pkg/distance"
)

func main() {
	p1, err := distance.NewPoint(1, 1)
	if err != nil {
		log.Fatal(err)
	}

	p2, err := distance.NewPoint(2, 2)
	if err != nil {
		log.Fatal(err)
	}

	d := distance.Calculate(p1, p2)
	fmt.Println(d)
}
