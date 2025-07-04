package main

import "fmt"

func main(){
	var start float64
	var end float64
	var step float64
	fmt.Println("Введите начальную температуру: ")
	_, err := fmt.Scan(&start)
	if err != nil{
		fmt.Print("Ошибка с числом")
	}
	fmt.Println("Введите конечную температуру: ")
	_, err = fmt.Scan(&end)
	if err != nil{
		fmt.Print("Ошибка с числом")
	}
	fmt.Println("Введите шаг: ")
	_, err = fmt.Scan(&step)
	if err != nil{
		fmt.Print("Ошибка с числом")
	}
	for i := start; i <= end; i+=step{
		fmt.Printf("%.2f°С = %.2f°F\n", i, i*1.8+32)
	}
}