package main

import (
	"fmt"
	//"math"
)

func square (num int) int{
	return num*num
}

func main(){
	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil{
		fmt.Print("Ошибка с числом")
	}
	fmt.Printf("Квадрат числа %d равен %d", number, square(number))
}