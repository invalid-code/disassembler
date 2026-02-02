package main

import (
	"github.com/invalid-code/disassembler/elf"
	"os"
)

func main() {
	data, err := os.ReadFile("hello2")
	if err != nil {
		panic("Error: Couldn't read file")
	}
	isElf := elf.IsElf(data[0x0:0x4])
	if isElf {
		elf.Disassemble(data)
	}
}
