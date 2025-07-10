package main

import "fmt"

type Messanger interface {
	Message() string
}

type BasicMessenger struct{}

func (m *BasicMessenger) Message() string {
	return "Я сообщение!"
}

func main() {
	telega := BasicMessenger{}
	fmt.Println(telega.Message())
}
