package amdx8664

func opcodeMap38(curByte byte, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, bool) {
	switch curByte {
	case 0x0:
		if isOperandSizeOverride {
			// 0x66
			return PSHUFB
		} else {
			return PSHUFB
		}
	case 0x1:
		if isOperandSizeOverride {
			// 0x66
			return PHADDW
		} else {
			return PHADDW
		}
	case 0x2:
		if isOperandSizeOverride {
			// 0x66
			return PHADDD
		} else {
			return PHADDD
		}
	case 0x3:
		if isOperandSizeOverride {
			// 0x66
			return PHADDSW
		} else {
			return PHADDSW
		}
	case 0x4:
		if isOperandSizeOverride {
			// 0x66
			return PMADDUBSW
		} else {
			return PMADDUBSW
		}
	case 0x5:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBW
		} else {
			return PHSUBW
		}
	case 0x6:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBD
		} else {
			return PHSUBD
		}
	case 0x7:
		if isOperandSizeOverride {
			// 0x66
			return PHSUBSW
		} else {
			return PHSUBSW
		}
	case 0x8:
		if isOperandSizeOverride {
			// 0x66
			return PSIGNB
		} else {
			return PSIGNB
		}
	case 0x9:
		if isOperandSizeOverride {
			// 0x66
			return PSIGNW
		} else {
			return PSIGNW
		}
	case 0xA:
		if isOperandSizeOverride {
			// 0x66
			return PSIGND
		} else {
			return PSIGND
		}
	case 0xB:
		if isOperandSizeOverride {
			// 0x66
			return PMULHRSW
		} else {
			return PMULHRSW
		}
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDVB
		} else {
			return NOP
		}
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			return BLENDVPS
		} else {
			return NOP
		}
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			return PBLENDVB
		} else {
			return NOP
		}
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			return PTEST
		} else {
			return NOP
		}
	case 0x1C:
		if isOperandSizeOverride {
			// 0x66
			return PABSB
		} else {
			return PABSB
		}
	case 0x1D:
		if isOperandSizeOverride {
			// 0x66
			return PABSW
		} else {
			return PABSW
		}
	case 0x1E:
		if isOperandSizeOverride {
			// 0x66
			return PABSD
		} else {
			return PABSD
		}
	case 0x20:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBW
		} else {
			return NOP
		}
	case 0x21:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBD
		} else {
			return NOP
		}
	case 0x22:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXBQ
		} else {
			return NOP
		}
	case 0x23:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXWD
		} else {
			return NOP
		}
	case 0x24:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXWQ
		} else {
			return NOP
		}
	case 0x25:
		if isOperandSizeOverride {
			// 0x66
			return PMOVSXDQ
		} else {
			return NOP
		}
	case 0x28:
		if isOperandSizeOverride {
			// 0x66
			return PMULDQ
		} else {
			return NOP
		}
	case 0x29:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQQ
		} else {
			return NOP
		}
	case 0x2A:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTDQA
		} else {
			return NOP
		}
	case 0x2B:
		if isOperandSizeOverride {
			// 0x66
			return PACKUSDW
		} else {
			return NOP
		}
	case 0x30:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBW
		} else {
			return NOP
		}
	case 0x31:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBD
		} else {
			return NOP
		}
	case 0x32:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXBQ
		} else {
			return NOP
		}
	case 0x33:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXWD
		} else {
			return NOP
		}
	case 0x34:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXWQ
		} else {
			return NOP
		}
	case 0x35:
		if isOperandSizeOverride {
			// 0x66
			return PMOVZXDQ
		} else {
			return NOP
		}
	case 0x37:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTQ
		} else {
			return NOP
		}
	case 0x38:
		if isOperandSizeOverride {
			// 0x66
			return PMINSB
		} else {
			return NOP
		}
	case 0x39:
		if isOperandSizeOverride {
			// 0x66
			return PMINSD
		} else {
			return NOP
		}
	case 0x3A:
		if isOperandSizeOverride {
			// 0x66
			return PMINUW
		} else {
			return NOP
		}
	case 0x3B:
		if isOperandSizeOverride {
			// 0x66
			return PMINUD
		} else {
			return NOP
		}
	case 0x3C:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSB
		} else {
			return NOP
		}
	case 0x3D:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSD
		} else {
			return NOP
		}
	case 0x3E:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUW
		} else {
			return NOP
		}
	case 0x3F:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUD
		} else {
			return NOP
		}
	case 0x40:
		if isOperandSizeOverride {
			// 0x66
			return PMULLD
		} else {
			return NOP
		}
	case 0x41:
		if isOperandSizeOverride {
			// 0x66
			return PHMINPOSUW
		} else {
			return NOP
		}
	case 0xDB:
		if isOperandSizeOverride {
			// 0x66
			return AESIMC
		} else {
			return NOP
		}
	case 0xDC:
		if isOperandSizeOverride {
			// 0x66
			return AESENC
		} else {
			return NOP
		}
	case 0xDD:
		if isOperandSizeOverride {
			// 0x66
			return AESENCLAST
		} else {
			return NOP
		}
	case 0xDE:
		if isOperandSizeOverride {
			// 0x66
			return AESDEC
		} else {
			return NOP
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return AESDECLAST
		} else {
			return NOP
		}
	case 0xF0:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD
		} else if isRep1 {
			// 0xF3
			return MOVSS
		} else if isRep0 {
			// 0xF2
			return MOVSD
		} else {
			return MOVUPS
		}
	case 0xF1:
		if isOperandSizeOverride {
			// 0x66
			return MOVBE
		} else if isRep1 {
			// 0xF3
			return CRC32
		} else if isRep0 {
			// 0xF2
			return CRC32
		} else {
			return MOVBE
		}
	case 0xF5:
		if isOperandSizeOverride {
			// 0x66
			return WRUSS
		} else {
			return NOP
		}
	case 0xF6:
		return ADD
	}
	panic("Error: todo")
}
