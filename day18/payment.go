package main

import "fmt"

type Payer interface {
	Pay(amount float64) string
}

type CreditCard struct {
	Number string
}

type Cash struct {
	Amount float64
}

type Crypto struct {
	Wallet string
}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f via card *%s", amount, c.Number)
}

func (c *Cash) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f cash", amount)
}

func (c *Crypto) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f via wallet %s", amount, c.Wallet)
}

func Chekout(p Payer, amount float64) {
	fmt.Println(p.Pay(amount))
}

func main() {
	card := CreditCard{Number: "1234"}
	cash := Cash{Amount: 250.50}
	crypto := Crypto{Wallet: "card"}
	Chekout(&card, 120)
	Chekout(&cash, 115)
	Chekout(&crypto, 110)
}
