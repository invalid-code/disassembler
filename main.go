package main

import (
	"github.com/invalid-code/disassembler/elf"
	"os"
)

func main() {
	data, err := os.ReadFile("avx_example")
	if err != nil {
		panic("Error: Couldn't read file")
	}
	isElf := elf.IsElf(data[0x0:0x4])
	if isElf {
		elf.Disassemble(data)
	}
}
