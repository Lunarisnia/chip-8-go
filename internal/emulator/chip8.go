package emulator

import (
	"fmt"
	"github.com/Lunarisnia/chip-8-go/internal/emulator/fontset"
	"os"
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
		DelayTimer:     0,
		SoundTimer:     0,
	}
	for i, b := range fontset.FontSet {
		chip8.Memory[i] = b
	}
	return &chip8
}

func (c *Chip8) LoadROM(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for i, b := range file {
		c.Memory[512+i] = b
	}
	return nil
}

func (c *Chip8) prefixZero() {
	switch c.OpCode & 0x000F {
	case 0x00E0:
		fmt.Println("Clear the screen")
	case 0x00EE:
		fmt.Println("Return from a subroutine")
	default:
		fmt.Println("Execute machine language subroutine")
	}
}

func (c *Chip8) prefixEight() {
	switch c.OpCode & 0x000F {
	case 0x0000:
		fmt.Println("Store the value of register V[0x00Y0] in register V[0x0X00]")
	case 0x0001:
		fmt.Println("Set VX to (VX || VY)")
	case 0x0002:
		fmt.Println("Set VX to (VX && VY)")
	case 0x0003:
		fmt.Println("Set VX to (VX ^ VY)")
	case 0x0004:
		fmt.Println("Add the value of VY to VX set VF to 1 if a carry occur or set VF to 0 if a carry does not occur")
	case 0x0005:
		fmt.Println("Subtract the value of VY from VX set VF to 0 if a borrow occur or set VF to 1 if a borrow does not occur")
	case 0x0006:
		fmt.Println("Store the value of VY shifted by one bit to VX set VF to the least significant bit prior to shift, VY in unchanged")
	case 0x0007:
		fmt.Println("Set VX to VY - VX, VF to 0 if a borrow occurs, to 1 if doesn't")
	case 0x000E:
		fmt.Println("Store the value of VY shifted by one bit to VX set VF to the most significant bit prior to shift, VY in unchanged")
	default:
		fmt.Println("unrecognized instruction")
	}
}

func (c *Chip8) prefixE() {
	switch c.OpCode & 0x000F {
	case 0x000E:
		fmt.Println("Skip the following instruction if the key corresponding to the hex value currently stored in register VX is pressed")
	case 0x0001:
		fmt.Println("Skip the following instruction if the key corresponding to the hex value currently stored in register VX is not pressed")
	default:
		fmt.Println("unrecognized instruction")
	}
}

func (c *Chip8) prefixF() {
	switch c.OpCode & 0x00FF {
	case 0x0007:
		fmt.Println("Store the current value of the delay timer in register VX")
	case 0x000A:
		fmt.Println("Wait for a keypress and store the result in register VX")
	case 0x0015:
		fmt.Println("Set the delay timer to the value of register VX")
	case 0x0018:
		fmt.Println("Set the sound timer to the value of register VX")
	case 0x001E:
		fmt.Println("Add the value stored in register VX to register I")
	case 0x0029:
		fmt.Println("Set I to the memory address of the sprite data corresponding to the hexadecimal digit stored in register VX")
	case 0x0033:
		fmt.Println("Store the binary-coded decimal equivalent of the value stored in register VX at addresses I, I + 1, and I + 2")
	case 0x0055:
		fmt.Println("Store the values of registers V0 to VX inclusive in memory starting at address I, I is set to I + X + 1 after operation")
	case 0x0065:
		fmt.Println("Fill registers V0 to VX inclusive with the values stored in memory starting at address I, I is set to I + X + 1 after operation")
	default:
		fmt.Println("unrecognized instruction")
	}
}

func (c *Chip8) EmulateCycle() {
	// Fetch Opcode
	c.OpCode = uint16(c.Memory[c.ProgramCounter])<<8 | uint16(c.Memory[c.ProgramCounter+1])
	c.ProgramCounter += 2
	// Decode Opcode
	switch c.OpCode & 0xF000 {
	case 0x0000:
		c.prefixZero()
	case 0x1000:
		fmt.Println("Jump to address 0x0NNN")
	case 0x2000:
		fmt.Println("Execute subroutine at address 0x0NNN")
	case 0x3000:
		fmt.Println("Skip the following instruction if the value of register V[0x0X00] equal to 0x00NN")
	case 0x4000:
		fmt.Println("Skip the following instruction if the value of register V[0x0X00] is not equal to 0x00NN")
	case 0x5000:
		fmt.Println("Skip the following instruction if the value of register V[0x0X00] is equal to V[0x00Y0]")
	case 0x6000:
		fmt.Println("Store number 0x00NN in register V[0x0X00]")
	case 0x7000:
		fmt.Println("Add value 0x00NN to register V[0x0X00]")
	case 0x8000:
		c.prefixEight()
	case 0x9000:
		fmt.Println("Skip the following instruction if the value of register VX is not equal to the value of register VY")
	case 0xA000:
		fmt.Println("Store memory address NNN in register I")
	case 0xB000:
		fmt.Println("Jump to address NNN + V0")
	case 0xC000:
		fmt.Println("Set VX to a random number with a mask of NN")
	case 0xD000:
		fmt.Println("Draw a sprite at position VX, VY with N bytes of sprite data starting at the address stored in I\nSet VF to 01 if any set pixels are changed to unset, and 00 otherwise")
	case 0xE000:
		c.prefixE()
	case 0xF000:
		c.prefixF()
	default:
		fmt.Println("unrecognized instruction")
	}
	// Execute Opcode

	// Update timers
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
