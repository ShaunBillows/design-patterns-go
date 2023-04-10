package computer

// Interface declaration

type IComputerBuilder interface {
	SetCPU(cpu CPU) IComputerBuilder
	SetGPU(gpu GPU) IComputerBuilder
	SetRAM(ram RAM) IComputerBuilder
	Build() Computer
	Reset()
}

// Computer builder

func NewComputerBuilder() *Builder {
	return &Builder{
		computer: Computer{
			CPU: DefaultCPU,
			GPU: DefaultGPU,
			RAM: DefaultRAM,
		},
	}
}

type Builder struct {
	computer Computer
}

func (cb *Builder) SetCPU(cpu CPU) IComputerBuilder {
	cb.computer.CPU = cpu
	return cb
}

func (cb *Builder) SetGPU(gpu GPU) IComputerBuilder {
	cb.computer.GPU = gpu
	return cb
}

func (cb *Builder) SetRAM(ram RAM) IComputerBuilder {
	cb.computer.RAM = ram
	return cb
}

func (cb *Builder) Build() Computer {
	c := cb.computer
	cb.Reset()
	return c
}
func (cb *Builder) Reset() {
	cb.computer = Computer{
		CPU: DefaultCPU,
		GPU: DefaultGPU,
		RAM: DefaultRAM,
	}
}
