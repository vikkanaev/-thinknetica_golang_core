package main

import "fmt"

type Employee struct {
	age  int
	Name string
}

type Customer struct {
	age  int
	Name string
}

type person interface{}

func oldest(persons []person) (p person) {
	var age, curAge = 0, 0
	for _, curPers := range persons {
		switch e := curPers.(type) {
		case *Customer:
			curAge = e.age
		case *Employee:
			curAge = e.age
		default:
			curAge = -1
		}

		if curAge > age {
			p = curPers
			age = curAge
		}
	}
	return p
}

func main() {
	persons := []person{
		&Customer{age: 10, Name: "Bob"},
		&Employee{age: 40, Name: "Yoe"},
		&Employee{age: 20, Name: "Sam"}}
	oPerson := oldest(persons)

	switch op := oPerson.(type) {
	case *Customer:
		fmt.Println("Oldest person is customer", op.Name)
	case *Employee:
		fmt.Println("Oldest person is emploee", op.Name)
	}
}
