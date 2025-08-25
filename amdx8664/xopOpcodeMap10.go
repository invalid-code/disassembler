package amdx8664

func xopOpcodeMap10(curByte byte, opcodeExt [2]bool, isRexW bool) (Instruction, bool, bool, MemSegment, Register, Register, Register, int) {
	switch curByte {
	case 0x10:
		switch opcodeExt {
		case [2]bool{false, false}:
			return BEXTR, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x12:
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
	default:
		panic("Error: Unknown instruction")
	}
}
