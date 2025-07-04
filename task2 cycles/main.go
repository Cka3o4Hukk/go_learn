package main

import "fmt"

func main(){
	var number int
	fmt.Println("Введите положительное целое число или 0: ")
	_, err := fmt.Scan(&number)
	if err != nil{
		fmt.Print("Ошибка с числом")
	} else if number == 1  {
		fmt.Print("Факториал: 0")
		return
	}
	sum := 1
 	for i:=1; i<=number; i++{
 		sum *= i
 	}
 	fmt.Printf("Факториал: %d", sum)
}