package main

import "fmt"

type Alerter interface {
	Alert() string
}

type PhoneAlerter struct {
	alertText string
}

type WatchAlerter struct {
	alertText string
}

func (p *PhoneAlerter) Alert() string {
	return fmt.Sprintf("Звонок: %s", p.alertText)
}

func (w *WatchAlerter) Alert() string {
	return fmt.Sprintf("Вибрация: %s", w.alertText)
}

func SendAlert(a Alerter) {
	fmt.Println(a.Alert())
}

func main() {
	iphone13 := PhoneAlerter{alertText: "Пора встать"}
	applewatch := WatchAlerter{alertText: "Напоминание"}

	SendAlert(&iphone13)
	SendAlert(&applewatch)
}
