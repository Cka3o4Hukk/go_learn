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
// Создайте программу, которая запрашивает у пользователя температуру в градусах Цельсия (тип float64), сохраняет её в переменную, а затем конвертирует в градусы Фаренгейта по формуле: F = C * 9/5 + 32. Выведите результат с точностью до двух знаков после запятой.