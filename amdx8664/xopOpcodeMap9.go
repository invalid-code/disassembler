package amdx8664

func xopOpcodeMap9(curByte byte, opcodeExt [2]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return NoInstruction, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return NoInstruction, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x12:
		switch opcodeExt {
		case [2]bool{false, false}:
			return NoInstruction, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x80:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x81:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x82:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZSS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x83:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VFRCZSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x90:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x91:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x92:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x93:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x94:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x95:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x96:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x97:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHLQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x98:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x99:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9A:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9B:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPSHAQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDBD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDBQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC7:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDWQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xCB:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUBWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUBD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUBQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD7:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUWQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDB:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHADDUDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHSUBBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHSUBWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPHSUBDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func xopOpcodeMap9ModRMG1(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x1:
		switch modrmReg {
		case [3]bool{false, false, true}:
			return BLCFILL, []Operand{}
		case [3]bool{false, true, false}:
			return BLSFILL, []Operand{}
		case [3]bool{false, true, true}:
			return BLCS, []Operand{}
		case [3]bool{true, false, false}:
			return TZMSK, []Operand{}
		case [3]bool{true, false, true}:
			return BLCIC, []Operand{}
		case [3]bool{true, true, false}:
			return BLSIC, []Operand{}
		case [3]bool{true, false, true}:
			return T1MSKC, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}

	default:
		panic("Error: Unknown instruction")
	}
}

func xopOpcodeMap9ModRMG2(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x2:
		switch modrmReg {
		case [3]bool{false, false, true}:
			return BLCMSK, []Operand{}
		case [3]bool{true, true, false}:
			return BLCI, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func xopOpcodeMap9ModRMG3(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x12:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return LLWPCB, []Operand{}
		case [3]bool{false, false, true}:
			return SLWPCB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
