package amdx8664

func opcodeMap3A(curByte byte, isOperandSizeOverride bool, isRexB bool) (Instruction, bool, bool, MemSegment, Register, Register, int) {
	switch curByte {
	case 0x8:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDPS, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x9:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDPD, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xA:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDSS, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xB:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDSD, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xC:
		if isOperandSizeOverride {
			// 0x66
			return BLENDPS, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xD:
		if isOperandSizeOverride {
			// 0x66
			return BLENDPD, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xE:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDW, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xF:
		if isOperandSizeOverride {
			// 0x66
			return PALIGNR, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PALIGNR, true, true, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			panic("todo multiple instruction")
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			panic("todo multiple instruction")
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x16:
		if isOperandSizeOverride {
			// 0x66
			panic("todo multiple instruction")
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			panic("todo multiple instruction")
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x20:
		if isOperandSizeOverride {
			// 0x66
			panic("todo multiple instruction")
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x21:
		if isOperandSizeOverride {
			// 0x66
			panic("todo multiple instruction")
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x22:
		if isOperandSizeOverride {
			// 0x66
			panic("todo multiple instruction")
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x40:
		if isOperandSizeOverride {
			// 0x66
			return DPPS, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x41:
		if isOperandSizeOverride {
			// 0x66
			return DPPD, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x42:
		if isOperandSizeOverride {
			// 0x66
			return MPSADBW, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x44:
		if isOperandSizeOverride {
			// 0x66
			return PCLMULQDQ, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x60:
		if isOperandSizeOverride {
			// 0x66
			return PCMPESTRM, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x61:
		if isOperandSizeOverride {
			// 0x66
			return PCMPESTRI, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x62:
		if isOperandSizeOverride {
			// 0x66
			return PCMPISTRM, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x63:
		if isOperandSizeOverride {
			// 0x66
			return PCMPISTRI, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return AESKEYGENASSIST, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	default:
		panic("Error: Unknown opcode")
	}
}
