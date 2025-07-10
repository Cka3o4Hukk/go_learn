package main

import "fmt"

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main(){
	rect1 := Rectangle{width:5.5, height:6.5}
	rect2 := Rectangle{width:5.5, height:6.5}
	fmt.Printf("Площадь: %.2f\nПериметр: %.2f", rect1.Area(), rect2.Perimeter())
}