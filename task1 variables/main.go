package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Введите 1 целое число: \n")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Ошибка с числом ")
	}
	fmt.Print("Введите 2 целое число: \n")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Ошибка с числом ")
	}
	a, b = b, a
	fmt.Printf("1 число: %d , 2 число: %d", a, b)

}
