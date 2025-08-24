package amdx8664

func opcodeMap3dnow(curByte byte, is64Bit bool, isRexB bool) (Instruction, bool, bool, MemSegment, Register, Register, int) {
	switch curByte {
	case 0xC:
		return PI2FW, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xD:
		return PI2FD, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x1C:
		return PF2IW, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x1D:
		return PF2ID, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x8A:
		return PFNACC, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x8E:
		return PFPNACC, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x90:
		return PFCMPGE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x94:
		return PFMIN, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x96:
		return PFRCP, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x97:
		return PFRSQRT, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9A:
		return PFSUB, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9E:
		return PFADD, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xA0:
		return PFCMPGT, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xA4:
		return PFMAX, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xA6:
		return PFRCPIT1, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xA7:
		return PFRSQIT1, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xAA:
		return PFSUBR, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xAE:
		return PFACC, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB0:
		return PFCMPEQ, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB4:
		return PFMUL, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB6:
		return PFRCPIT2, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB7:
		return PMULHRW, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xBB:
		return PSWAPD, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xBF:
		return PAVGUSB, true, false, NoSegment, NoRegister, NoRegister, 0
	default:
		panic("Error: Unknown opcode")
	}
}
