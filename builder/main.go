package main

import (
	"design-patterns-go/builder/computer"
	"fmt"
)

func main() {
	builder := computer.NewComputerBuilder()
	director := computer.NewComputerDirector(builder)

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

	defaultComputer := builder.Build()
	fmt.Println("Default computer:")
	println("CPU:", defaultComputer.CPU)
	println("GPU:", defaultComputer.GPU)
	println("RAM:", defaultComputer.RAM)
}
