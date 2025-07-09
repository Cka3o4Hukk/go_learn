package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64{
	return celsius * 9/5 + 32
}

func main(){
	var number float64
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil{
		fmt.Print("Ошибка с числом")
	}
	fmt.Printf("%.2f°С = %.2f°С", number, celsiusToFahrenheit(number))
}