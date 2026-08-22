package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b BankAccount) Deposit(amount float64) BankAccount {
	b.Balance += amount
	return b
}

func (b BankAccount) Withdraw(amount float64) BankAccount {
	if b.Balance >= amount {
		b.Balance -= amount
	} else {
		fmt.Println("Недостаточно средств!")
	}
	return b
}

func (b BankAccount) Transfer(amount float64, to BankAccount) (BankAccount, BankAccount) {
	User1 := b.Withdraw(amount)
	User2 := to.Deposit(amount)
	return User1, User2
}

func main() {
	Users1 := BankAccount{Owner: "Ivan", Balance: 1000}
	Users2 := BankAccount{Owner: "Olga", Balance: 0}
	Users1, Users2 = Users1.Transfer(500, Users2)
	fmt.Println(Users1.Owner, Users1.Balance)
	fmt.Println(Users2.Owner, Users2.Balance)
}
