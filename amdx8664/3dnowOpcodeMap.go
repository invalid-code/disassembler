package amdx8664

func opcodeMap3dnow(curByte byte, is64Bit bool, isRexB bool) (Instruction, []Operand) {
	switch curByte {
	case 0xC:
		return PI2FW, []Operand{}
	case 0xD:
		return PI2FD, []Operand{}
	case 0x1C:
		return PF2IW, []Operand{}
	case 0x1D:
		return PF2ID, []Operand{}
	case 0x8A:
		return PFNACC, []Operand{}
	case 0x8E:
		return PFPNACC, []Operand{}
	case 0x90:
		return PFCMPGE, []Operand{}
	case 0x94:
		return PFMIN, []Operand{}
	case 0x96:
		return PFRCP, []Operand{}
	case 0x97:
		return PFRSQRT, []Operand{}
	case 0x9A:
		return PFSUB, []Operand{}
	case 0x9E:
		return PFADD, []Operand{}
	case 0xA0:
		return PFCMPGT, []Operand{}
	case 0xA4:
		return PFMAX, []Operand{}
	case 0xA6:
		return PFRCPIT1, []Operand{}
	case 0xA7:
		return PFRSQIT1, []Operand{}
	case 0xAA:
		return PFSUBR, []Operand{}
	case 0xAE:
		return PFACC, []Operand{}
	case 0xB0:
		return PFCMPEQ, []Operand{}
	case 0xB4:
		return PFMUL, []Operand{}
	case 0xB6:
		return PFRCPIT2, []Operand{}
	case 0xB7:
		return PMULHRW, []Operand{}
	case 0xBB:
		return PSWAPD, []Operand{}
	case 0xBF:
		return PAVGUSB, []Operand{}
	default:
		panic("Error: Unknown opcode")
	}
}
