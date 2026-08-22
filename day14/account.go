package main

import "fmt"

type Account struct {
	Name    string
	Balance float64
}

func (A Account) Deposit(amount float64) Account {
	A.Balance += amount
	return A
}

func (A Account) Withdraw(amount float64) Account {
	if A.Balance >= amount {
		A.Balance -= amount
	} else {
		fmt.Println("Недостаточно средств")
	}
	return A
}

func main() {
	User := Account{Name: "Julia", Balance: 0}
	User = User.Deposit(1000)
	User = User.Withdraw(250)
	User = User.Withdraw(250)
	fmt.Println(User.Balance)
}
