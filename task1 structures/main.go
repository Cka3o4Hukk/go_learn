package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) Person {

	return Person{
		name: name,
		age:  age,
	}
}

func main() {
	alex := newPerson("Alex", 20)
	fmt.Println("Имя:", alex.name)
	fmt.Printf("Возраст: %d", alex.age)
}