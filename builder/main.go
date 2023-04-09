package main

import "fmt"

func main() {
	builder := NewComputerBuilder()
	director := NewComputerDirector(builder)

	gamingComputer := director.BuildGamingComputer()
	println("Gaming computer:")
	println("CPU:", gamingComputer.CPU)
	println("GPU:", gamingComputer.GPU)
	println("RAM:", gamingComputer.RAM)

	workstation := director.BuildWorkstation()
	println("Workstation:")
	println("CPU:", workstation.CPU)
	println("GPU:", workstation.GPU)
	println("RAM:", workstation.RAM)

	customComputer := builder.SetCPU("Custom CPU").SetGPU("Custom GPU").SetRAM(64).Build()
	fmt.Println("Custom computer:")
	println("CPU:", customComputer.CPU)
	println("GPU:", customComputer.GPU)
	println("RAM:", customComputer.RAM)
}

// Computer type

type Computer struct {
	CPU string
	GPU string
	RAM int
}

// Computer builder interface

type IComputerBuilder interface {
	SetCPU(s string) IComputerBuilder
	SetGPU(s string) IComputerBuilder
	SetRAM(i int) IComputerBuilder
	Build() Computer
}

// Computer builder

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

type ComputerBuilder struct {
	computer Computer
}

func (cb *ComputerBuilder) SetCPU(s string) IComputerBuilder {
	cb.computer.CPU = s
	return cb
}

func (cb *ComputerBuilder) SetGPU(s string) IComputerBuilder {
	cb.computer.GPU = s
	return cb
}

func (cb *ComputerBuilder) SetRAM(i int) IComputerBuilder {
	cb.computer.RAM = i
	return cb
}

func (cb *ComputerBuilder) Build() Computer {
	return cb.computer
}

// Computer director

func NewComputerDirector(builder IComputerBuilder) *ComputerDirector {
	return &ComputerDirector{
		builder: builder,
	}
}

type ComputerDirector struct {
	builder IComputerBuilder
}

func (cd *ComputerDirector) BuildGamingComputer() Computer {
	return cd.builder.SetCPU("High-end CPU").SetGPU("High-end GPU").SetRAM(32).Build()
}

func (cd *ComputerDirector) BuildWorkstation() Computer {
	return cd.builder.SetCPU("Medium-end CPU").SetGPU("Medium-end GPU").SetRAM(16).Build()
}
