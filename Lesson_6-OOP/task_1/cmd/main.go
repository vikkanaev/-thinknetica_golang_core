package main

import (
	"fmt"
	"thinknetica_golang_core/Lesson_6-OOP/task_1/pkg/distance"
)

func main() {
	d := distance.CalculateDistance(1, 1, 4, 5)
	fmt.Println(d)
}
