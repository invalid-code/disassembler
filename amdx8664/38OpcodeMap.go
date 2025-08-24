package amdx8664

func opcodeMap38(curByte byte, is64Bit bool, isRep0 bool, isOperandSizeOverride bool, isRexB bool) (Instruction, bool, bool, MemSegment, Register, Register, int) {
	switch curByte {
	case 0x0:
		if isOperandSizeOverride {
			// 0x66
			return PSHUFB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PSHUFB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x1:
		if isOperandSizeOverride {
			// 0x66
			return PHADDW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PHADDW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x2:
		if isOperandSizeOverride {
			// 0x66
			return PHADDD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PHADDD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x3:
		if isOperandSizeOverride {
			// 0x66
			return PHADDSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PHADDSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x4:
		if isOperandSizeOverride {
			// 0x66
			return PMADDUBSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PMADDUBSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x5:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PHSUBW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x6:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PHSUBD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x7:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PHSUBSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x8:
		if isOperandSizeOverride {
			// 0x66
			return PSIGNB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PSIGNB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x9:
		if isOperandSizeOverride {
			// 0x66
			return PSIGNW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PSIGNW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xA:
		if isOperandSizeOverride {
			// 0x66
			return PSIGND, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PSIGND, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xB:
		if isOperandSizeOverride {
			// 0x66
			return PMULHRSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PMULHRSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDVB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			return BLENDVPS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDVB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			return PTEST, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x1C:
		if isOperandSizeOverride {
			// 0x66
			return PABSB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PABSB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x1D:
		if isOperandSizeOverride {
			// 0x66
			return PABSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PABSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x1E:
		if isOperandSizeOverride {
			// 0x66
			return PABSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return PABSD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x20:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x21:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x22:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x23:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXWD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x24:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXWQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x25:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x28:
		if isOperandSizeOverride {
			// 0x66
			return PMULDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x29:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x2A:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTDQA, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x2B:
		if isOperandSizeOverride {
			// 0x66
			return PACKUSDW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x30:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x31:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x32:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x33:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXWD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x34:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXWQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x35:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x37:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x38:
		if isOperandSizeOverride {
			// 0x66
			return PMINSB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x39:
		if isOperandSizeOverride {
			// 0x66
			return PMINSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3A:
		if isOperandSizeOverride {
			// 0x66
			return PMINUW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3B:
		if isOperandSizeOverride {
			// 0x66
			return PMINUD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3C:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3D:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3E:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x3F:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x40:
		if isOperandSizeOverride {
			// 0x66
			return PMULLD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0x41:
		if isOperandSizeOverride {
			// 0x66
			return PHMINPOSUW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDB:
		if isOperandSizeOverride {
			// 0x66
			return AESIMC, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDC:
		if isOperandSizeOverride {
			// 0x66
			return AESENC, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDD:
		if isOperandSizeOverride {
			// 0x66
			return AESENCLAST, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDE:
		if isOperandSizeOverride {
			// 0x66
			return AESDEC, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return AESDECLAST, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xF0:
		if isOperandSizeOverride && isRep0 {
			// 0x66 0xF2
			return CRC32, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CRC32, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isOperandSizeOverride {
			// 0x66
			return MOVBE, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MOVBE, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF1:
		if isOperandSizeOverride && isRep0 {
			// 0x66 0xF2
			return CRC32, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CRC32, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isOperandSizeOverride {
			// 0x66
			return MOVBE, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MOVBE, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF5:
		if isOperandSizeOverride {
			// 0x66
			return WRUSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown opcode")
		}
	case 0xF6:
		if isOperandSizeOverride || isRep0 {
			panic("Error: Unknown opcode")
		}
		return WRSS, true, false, NoSegment, NoRegister, NoRegister, 0
	default:
		panic("Error: Unknown opcode")
	}
}
