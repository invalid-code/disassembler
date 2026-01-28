package amdx8664

func xopOpcodeMap8(curByte byte, opcodeExt [2]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x85:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSWW, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x86:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSWD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x87:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSDQL, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x8E:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSDD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x8F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSDQH, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x95:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSWW, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x96:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSWD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x97:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSDQL, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x9E:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSDD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0x9F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSDQH, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xA2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCMOV, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xA3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPPERM, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xA6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMADCSSWD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xB6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMADCSWD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC0:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTB, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTW, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTQ, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCC:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccB, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCD:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccW, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCE:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCF:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccQ, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xEC:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUB, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xED:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUW, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xEE:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUD, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	case 0xEF:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUQ, []Operand{}
		default:
			panic("Error: Unknown opcode")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
