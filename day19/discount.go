package main

import "fmt"

type Discounter interface {
	Apply(price float64) float64
}

type Percent struct {
	Value float64
}
type Flat struct {
	Value float64
}

func (p *Percent) Apply(price float64) float64 {
	return price - price*p.Value/100
}
func (f *Flat) Apply(price float64) float64 {
	return price - f.Value
}

func FinalPrice(discounts []Discounter, price float64) float64 {
	for _, val := range discounts {
		price = val.Apply(price)
	}
	return price
}

func main() {
	sale := []Discounter{&Percent{10}, &Flat{50}}
	fmt.Println(FinalPrice(sale, 1000))
}
