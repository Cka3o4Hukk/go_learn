package main

import "fmt"

type Notifier interface {
	Notify() string
}

type EmailNotifier struct {
	message string
}

func (e *EmailNotifier) Notify() string {
	return fmt.Sprintf("Уведомление: %s", e.message)
}

func main() {
	test_mess := EmailNotifier{"Письмо отправлено"}
	fmt.Println(test_mess.Notify())
}
