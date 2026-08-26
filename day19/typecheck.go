package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct {
	Brand string
}
type Bike struct{}
type Truck struct {
	Tons int
}

func (c *Car) Drive() string {
	return fmt.Sprintf("Car: %s", c.Brand)
}

func (b *Bike) Drive() string {
	return "Bike: light and fast"
}

func (t *Truck) Drive() string {
	return fmt.Sprintf("Truck: carries %d tons", t.Tons)
}

func Info(v Vehicle) {
	switch m := v.(type) {
	case *Car:
		fmt.Println("Car:", m.Brand)
	case *Bike:
		fmt.Println("Bike: light and fast")
	case *Truck:
		fmt.Println("Truck: carries", m.Tons, "tons")
	}
}

func main() {
	transport := []Vehicle{
		&Car{"BMW"},
		&Bike{},
		&Truck{520},
	}
	for _, v := range transport {
		Info(v)
	}
}
