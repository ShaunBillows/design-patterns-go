package main

import "fmt"

// Transporter interface
type Transporter interface {
	Deliver(loc string)
	Load(products []string)
}

// Truck type
type Truck struct {
	products []string
}

func NewTruck() *Truck {
	return &Truck{}
}

func (t *Truck) Load(products []string) {
	t.products = products
}

func (t *Truck) Deliver(loc string) {
	fmt.Printf("truck delivering %v to %v\n", t.products, loc)
}

// Ship type
type Ship struct {
	products []string
}

func NewShip() *Ship {
	return &Ship{}
}

func (s *Ship) Load(products []string) {
	s.products = products
}
func (s *Ship) Deliver(loc string) {
	fmt.Printf("ship delivering %v to %v\n", s.products, loc)
}

// Factory function
func getTransporter(transportType string) Transporter {
	if transportType == "truck" {
		return NewTruck()
	}
	if transportType == "ship" {
		return NewShip()
	}
	return nil
}

// Implementation
func main() {
	truck := getTransporter("truck")
	ship := getTransporter("ship")
	truck.Load([]string{"product 1", "product 2", "product 3"})
	ship.Load([]string{"product 4", "product 5", "product 6"})
	truck.Deliver("London")
	ship.Deliver("Spain")
}
