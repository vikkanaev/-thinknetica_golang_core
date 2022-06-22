package main

import "fmt"

type Employee struct {
	age int
}

type Customer struct {
	age int
}

type person interface {
	get_age() int
}

func (e *Employee) get_age() int {
	return e.age
}

func (c *Customer) get_age() int {
	return c.age
}

func oldest(persons []person) (age int) {
	for _, e := range persons {
		if e.get_age() > age {
			age = e.get_age()
		}
	}
	return age
}

func main() {
	persons := []person{
		&Customer{age: 10},
		&Employee{age: 40},
		&Employee{age: 20}}

	age := oldest(persons)
	fmt.Println("Oldest person has age:", age)
}
