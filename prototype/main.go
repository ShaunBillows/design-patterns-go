package main

import "fmt"

func main()  {
	cat := NewAnimal("Cat")
	catClone := cat.Clone()
	fmt.Println(cat.WhatAmI())
	fmt.Println(catClone.WhatAmI())
}

// Define an animal interface, now the animal struct can be swapped out 
// for another type which implements the interface, increasing flexibility. 
// Often only contains the Clone field.
type IAnimal interface {
	Clone() IAnimal
	WhatAmI() string
}

// Define animal constructor
func NewAnimal(kind string) IAnimal {
	return &Animal{
		kind: kind,
	}
}

// Define concrete animal type that implements IAnimal
type Animal struct {
	kind string
}

func (a *Animal) WhatAmI() string {
	return a.kind
}

// In the prototype pattern we return a new struct instance with the same fields.	
func (a *Animal) Clone() IAnimal {
	clone := &Animal{kind: a.kind}
	return clone
}
