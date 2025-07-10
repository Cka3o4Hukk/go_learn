package main

import "fmt"

const USDTtoRUB = 80

type CreditCard struct {
	OwnerName string
	Cash      float64
	Limit     float64
}

func NewCard(ownerName string, cash, limit float64) CreditCard {
	return CreditCard{
		OwnerName: ownerName,
		Cash:      cash,
		Limit:     limit,
	}
}

func (c CreditCard) ShowRUB() {
	fmt.Printf("%.2f rub\n", c.Cash)
}

func (c CreditCard) ShowUSDT() {
	fmt.Printf("%.2f usdt\n", c.Cash/USDTtoRUB)
}

func (c CreditCard) IsExpired() bool {
	return c.Limit > c.Cash
}

func (c CreditCard) ShowInfo(currentHour int) {
	if c.IsExpired() {
		fmt.Printf("%s, жди коллекторов, сука\n", c.OwnerName)
		return
	}

	fmt.Printf("Hello %s, it's %d o'clock\n", c.OwnerName, currentHour)
}

func main() {
	sber := NewCard("Alex", 1000, 500)
	alfa := NewCard("Tema", 500, 1000)

	sber.ShowUSDT()
	sber.ShowInfo(7)

	alfa.ShowRUB()
	alfa.ShowInfo(3)
}
