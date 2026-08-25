package main

import "fmt"

type Converter interface {
	Convert(value float64) float64
}

type KmToMiles struct{}

type CelsiusToFahrenheit struct{}

func (k *KmToMiles) Convert(value float64) float64 {
	return value * 0.621
}

func (c *CelsiusToFahrenheit) Convert(value float64) float64 {
	return value*1.8 + 32
}

func Show(c Converter, value float64) {
	fmt.Println(c.Convert(value))
}

func main() {
	km := KmToMiles{}
	cels := CelsiusToFahrenheit{}
	Show(&km, 23.3)
	Show(&cels, 36.6)
}
