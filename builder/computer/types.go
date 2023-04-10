package computer

// Type and constant declarations

type Computer struct {
	CPU CPU
	GPU GPU
	RAM RAM
}

type RAM int
type CPU string
type GPU string

const (
	DefaultRAM = RAM(32)
	DefaultCPU = CPU("Default CPU")
	DefaultGPU = GPU("Default GPU")
)
