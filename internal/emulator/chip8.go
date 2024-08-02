package emulator

import (
	"fmt"
	"github.com/Lunarisnia/chip-8-go/internal/emulator/fontset"
)

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
	chip8 := Chip8{
		ProgramCounter: 0x200,
		I:              0,
		StackPointer:   0,
		OpCode:         0,
	}
	for i, b := range fontset.FontSet {
		chip8.Memory[i] = b
	}
	return &chip8
}

func (c *Chip8) PrintFontSet() {
	for i := 0; i < len(fontset.FontSet); i += 5 {
		font := c.Memory[i : i+5]
		for _, d := range font {
			fmt.Printf("%08b\n", d)
		}
		fmt.Println("")
	}
}
