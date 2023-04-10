package computer

// Computer director

func NewComputerDirector(builder IComputerBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

type Director struct {
	builder IComputerBuilder
}

func (cd *Director) BuildGamingComputer() (c Computer) {
	c = cd.builder.SetCPU("High-end CPU").SetGPU("High-end GPU").SetRAM(32).Build()
	cd.builder.Reset()
	return
}

func (cd *Director) BuildWorkstation() (c Computer) {
	c = cd.builder.SetCPU("Medium-end CPU").SetGPU("Medium-end GPU").SetRAM(16).Build()
	cd.builder.Reset()
	return
}
