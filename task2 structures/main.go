package main

import "fmt"

type Person struct {
	name string
	age  int
	address string
}

func newPerson(name string, age int, address string) Person {

	return Person{
		name: name,
		age:  age,
		address: address,
	}
}

func (p Person)String() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d, Адрес: %s", p.name, p.age, p.address)
}

func main() {
	fmt.Print(newPerson("Alex", 30, "Moscow"))
}