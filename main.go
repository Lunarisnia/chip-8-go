package main

import (
	"fmt"
	"github.com/Lunarisnia/chip-8-go/internal/emulator"
)

func main() {
	chip8 := emulator.NewChip8()
	fmt.Println(chip8.Memory[0] == byte('c'))
}
