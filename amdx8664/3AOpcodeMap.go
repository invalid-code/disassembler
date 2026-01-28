package amdx8664

func opcodeMap3A(curByte byte, is64Bit bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x8:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDPS, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x9:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDPD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xA:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDSS, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xB:
		if isOperandSizeOverride {
			// 0x66
			return ROUNDSD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xC:
		if isOperandSizeOverride {
			// 0x66
			return BLENDPS, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xD:
		if isOperandSizeOverride {
			// 0x66
			return BLENDPD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xE:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xF:
		if isOperandSizeOverride {
			// 0x66
			return PALIGNR, []Operand{}
		} else {
			return PALIGNR, []Operand{}
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
			return DPPS, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x41:
		if isOperandSizeOverride {
			// 0x66
			return DPPD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x42:
		if isOperandSizeOverride {
			// 0x66
			return MPSADBW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x44:
		if isOperandSizeOverride {
			// 0x66
			return PCLMULQDQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x60:
		if isOperandSizeOverride {
			// 0x66
			return PCMPESTRM, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x61:
		if isOperandSizeOverride {
			// 0x66
			return PCMPESTRI, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x62:
		if isOperandSizeOverride {
			// 0x66
			return PCMPISTRM, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x63:
		if isOperandSizeOverride {
			// 0x66
			return PCMPISTRI, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return AESKEYGENASSIST, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	default:
		panic("Error: Unknown opcode")
	}
}
