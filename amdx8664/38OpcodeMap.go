package amdx8664

func opcodeMap38(curByte byte, is64Bit bool, isRep0 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x0:
		if isOperandSizeOverride {
			// 0x66
			return PSHUFB, []Operand{}
		} else {
			return PSHUFB, []Operand{}
		}
	case 0x1:
		if isOperandSizeOverride {
			// 0x66
			return PHADDW, []Operand{}
		} else {
			return PHADDW, []Operand{}
		}
	case 0x2:
		if isOperandSizeOverride {
			// 0x66
			return PHADDD, []Operand{}
		} else {
			return PHADDD, []Operand{}
		}
	case 0x3:
		if isOperandSizeOverride {
			// 0x66
			return PHADDSW, []Operand{}
		} else {
			return PHADDSW, []Operand{}
		}
	case 0x4:
		if isOperandSizeOverride {
			// 0x66
			return PMADDUBSW, []Operand{}
		} else {
			return PMADDUBSW, []Operand{}
		}
	case 0x5:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBW, []Operand{}
		} else {
			return PHSUBW, []Operand{}
		}
	case 0x6:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBD, []Operand{}
		} else {
			return PHSUBD, []Operand{}
		}
	case 0x7:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBSW, []Operand{}
		} else {
			return PHSUBSW, []Operand{}
		}
	case 0x8:
		if isOperandSizeOverride {
			// 0x66
			return PSIGNB, []Operand{}
		} else {
			return PSIGNB, []Operand{}
		}
	case 0x9:
		if isOperandSizeOverride {
			// 0x66
			return PSIGNW, []Operand{}
		} else {
			return PSIGNW, []Operand{}
		}
	case 0xA:
		if isOperandSizeOverride {
			// 0x66
			return PSIGND, []Operand{}
		} else {
			return PSIGND, []Operand{}
		}
	case 0xB:
		if isOperandSizeOverride {
			// 0x66
			return PMULHRSW, []Operand{}
		} else {
			return PMULHRSW, []Operand{}
		}
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDVB, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			return BLENDVPS, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDVB, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			return PTEST, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x1C:
		if isOperandSizeOverride {
			// 0x66
			return PABSB, []Operand{}
		} else {
			return PABSB, []Operand{}
		}
	case 0x1D:
		if isOperandSizeOverride {
			// 0x66
			return PABSW, []Operand{}
		} else {
			return PABSW, []Operand{}
		}
	case 0x1E:
		if isOperandSizeOverride {
			// 0x66
			return PABSD, []Operand{}
		} else {
			return PABSD, []Operand{}
		}
	case 0x20:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x21:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x22:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x23:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXWD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x24:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXWQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x25:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXDQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x28:
		if isOperandSizeOverride {
			// 0x66
			return PMULDQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x29:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x2A:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTDQA, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x2B:
		if isOperandSizeOverride {
			// 0x66
			return PACKUSDW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x30:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x31:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x32:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x33:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXWD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x34:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXWQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x35:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXDQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x37:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTQ, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x38:
		if isOperandSizeOverride {
			// 0x66
			return PMINSB, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x39:
		if isOperandSizeOverride {
			// 0x66
			return PMINSD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3A:
		if isOperandSizeOverride {
			// 0x66
			return PMINUW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3B:
		if isOperandSizeOverride {
			// 0x66
			return PMINUD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3C:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSB, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3D:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3E:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3F:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x40:
		if isOperandSizeOverride {
			// 0x66
			return PMULLD, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x41:
		if isOperandSizeOverride {
			// 0x66
			return PHMINPOSUW, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDB:
		if isOperandSizeOverride {
			// 0x66
			return AESIMC, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDC:
		if isOperandSizeOverride {
			// 0x66
			return AESENC, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDD:
		if isOperandSizeOverride {
			// 0x66
			return AESENCLAST, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDE:
		if isOperandSizeOverride {
			// 0x66
			return AESDEC, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return AESDECLAST, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xF0:
		if isOperandSizeOverride && isRep0 {
			// 0x66 0xF2
			return CRC32, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CRC32, []Operand{}
		} else if isOperandSizeOverride {
			// 0x66
			return MOVBE, []Operand{}
		} else {
			return MOVBE, []Operand{}
		}
	case 0xF1:
		if isOperandSizeOverride && isRep0 {
			// 0x66 0xF2
			return CRC32, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CRC32, []Operand{}
		} else if isOperandSizeOverride {
			// 0x66
			return MOVBE, []Operand{}
		} else {
			return MOVBE, []Operand{}
		}
	case 0xF5:
		if isOperandSizeOverride {
			// 0x66
			return WRUSS, []Operand{}
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xF6:
		if isOperandSizeOverride || isRep0 {
			panic("Error: Unknown opcode")
		}
		return WRSS, []Operand{}
	default:
		panic("Error: Unknown opcode")
	}
}
