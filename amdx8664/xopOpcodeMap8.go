package amdx8664

func xopOpcodeMap8(curByte byte, opcodeExt [2]bool, isRexW bool) (Instruction, bool, bool, MemSegment, Register, Register, Register, int) {
	switch curByte {
	case 0x85:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSWW, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x86:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSWD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x87:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSDQL, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x8E:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSDD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x8F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSSDQH, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x95:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSWW, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x96:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSWD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x97:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSDQL, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x9E:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSDD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0x9F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMACSDQH, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xA2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCMOV, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xA3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPPERM, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xA6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMADCSSWD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xB6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPMADCSWD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC0:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTB, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC1:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTW, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xC3:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPROTQ, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCC:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccB, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCD:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccW, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCE:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xCF:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccQ, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xEC:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUB, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xED:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUW, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xEE:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUD, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	case 0xEF:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VPCOMccUQ, true, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown opcode")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
