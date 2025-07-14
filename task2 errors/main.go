package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivideBy0 = errors.New("деление на ноль недопустимо")
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideBy0
	} else {
		return a / b, nil
	}
}

func main() {
	a, b := 10, 0
	//a, b := 10, 2
	result, err := Divide(a, b)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(result)
	}
}
