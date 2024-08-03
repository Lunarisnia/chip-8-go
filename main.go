package main

import (
	"github.com/Lunarisnia/chip-8-go/internal/emulator"
)

func main() {
	chip8 := emulator.NewChip8()
	err := chip8.LoadROM("./roms/br8kout.ch8")
	if err != nil {
		panic(err)
	}
	for chip8.ProgramCounter < 4096 {
		chip8.EmulateCycle()
	}
}
