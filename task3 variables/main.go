package main

import "fmt"

func main(){
	var length float64
	var width float64
	fmt.Println("Введите длину прямоугольника: ")
	_, err := fmt.Scan(&length)
	if err != nil {
		fmt.Println("Ошибка с числом")
	}
	fmt.Print("Введите ширину прямоугольника: \n")
	_, err = fmt.Scan(&width)
	if err != nil {
		fmt.Println("Ошибка с числом")
	}
	fmt.Printf("Периметр прямоугольника: \n%.2f\n", (length + width) * 2)
	fmt.Printf("Площадь прямоугольника: \n%.2f", length * width)
}
