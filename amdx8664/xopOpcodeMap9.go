package amdx8664

func xopOpcodeMap9(curByte byte, opcodeExt [2]bool, isRexW bool) (Instruction, bool, bool, MemSegment, Register, Register, Register, int) {
	switch curByte {
	case 0x1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x12:
		switch opcodeExt {
		case [2]bool{false, false}:
			return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x80:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x81:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x82:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x83:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x90:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x91:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x92:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x93:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x94:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x95:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x96:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x97:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x98:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x99:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9A:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9B:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDBD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDBQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC7:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDWQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xCB:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUBWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUBD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUBQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD7:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUWQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDB:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHSUBBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHSUBWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHSUBDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
