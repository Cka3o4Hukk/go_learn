package main

import "fmt"

func main(){
	var number int
	fmt.Println("Введите целое число: ")
	_, err := fmt.Scan(&number)
	if err != nil{
		fmt.Print("Ошибка с числом")
	}
	sum := 0
	for i:=1; i<=number; i++{
		sum += i
	}
	fmt.Printf("Сумма чисел: %d", sum)
}