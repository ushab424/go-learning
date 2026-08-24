package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

type Inventory struct {
	Items []string
}

func (i *Inventory) Add(item string) {
	i.Items = append(i.Items, item)
}

func (i *Inventory) Remove(item string) {
	for z, val := range i.Items {
		if val == item {
			i.Items = append(i.Items[:z], i.Items[z+1:]...)
			return
		}
	}
}

func (i *Inventory) Count() int {
	return len(i.Items)
}

func (i *Inventory) Print() {
	for _, value := range i.Items {
		fmt.Println(value)
	}
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance >= amount {
		b.Balance -= amount
	} else {
		fmt.Println("Insufficient funds")
	}
}

func (b *BankAccount) Info() {
	fmt.Println(b.Owner, b.Balance)
}

func main() {
	Account := BankAccount{Owner: "Ivan"}
	Account.Deposit(1000.00)
	Account.Deposit(265.20)
	Account.Withdraw(328.50)
	Account.Info()
	Products := Inventory{}
	Products.Add("Bread")
	Products.Add("Milk")
	Products.Add("Cofee")
	Products.Add("Oil")
	Products.Remove("Oil")
	fmt.Println(Products.Count())
	Products.Print()
}
