package amdx8664

func xopOpcodeMap10(curByte byte, opcodeExt [2]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x10:
		switch opcodeExt {
		case [2]bool{false, false}:
			return BEXTR, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x12:
		return NoInstruction, []Operand{}
	default:
		panic("Error: Unknown instruction")
	}
}

func xopOpcodeMap9ModRMG4(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x12:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return LWPINS, []Operand{}
		case [3]bool{false, false, true}:
			return LWPVAL, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
