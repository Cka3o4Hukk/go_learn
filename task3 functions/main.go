package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i:=2; i<=int(math.Sqrt(float64(n))); i++{
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main(){
	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil{
		fmt.Print("Ошибка с числом")
	}
	if isPrime(number){
		fmt.Printf("%d - простое число", number)
	} else {
		fmt.Printf("%d - сложное число", number)
	}	 
}