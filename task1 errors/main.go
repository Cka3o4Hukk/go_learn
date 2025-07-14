package main

import (
	"errors"
	"fmt"
)

func CheckAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("возраст не может быть отрицательным")
	}
	return fmt.Sprintf("Возраст принят: %d", age), nil
}

func main() {
	age, 
	err := CheckAge(-3)

	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(age)
	}
}
