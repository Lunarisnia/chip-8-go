package emulator

const windowWidth int32 = 64
const windowHeight int32 = 32

// Chip8 Memory Layout https://multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/
type Chip8 struct {
	Memory         [4096]byte
	OpCode         uint16
	V              [16]byte
	I              uint16
	ProgramCounter uint16
	Graphic        [windowWidth * windowHeight]byte
	DelayTimer     byte
	SoundTimer     byte
	Stack          [16]uint16
	StackPointer   uint16
	Key            [16]byte
}

func NewChip8() *Chip8 {
	chip8 := Chip8{}
	return &chip8
}
