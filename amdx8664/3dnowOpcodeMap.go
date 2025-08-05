package amdx8664

func opcodeMap3dnow(curByte byte, is64Bit bool) (Instruction, bool, bool) {
	switch curByte {
	case 0xC:
		return PI2FW
	case 0xD:
		return PI2FD
	case 0x1C:
		return PF2IW
	case 0x1D:
		return PF2ID
	case 0x8A:
		return PFNACC
	case 0x8E:
		return PFPNACC
	case 0x90:
		return PFCMPGE
	case 0x94:
		return PFMIN
	case 0x96:
		return PFRCP
	case 0x97:
		return PFRSQRT
	case 0x9A:
		return PFSUB
	case 0x9E:
		return PFADD
	case 0xA0:
		return PFCMPGT
	case 0xA4:
		return PFMAX
	case 0xA6:
		return PFRCPIT1
	case 0xA7:
		return PFRSQIT1
	case 0xAA:
		return PFSUBR
	case 0xAE:
		return PFACC
	case 0xB0:
		return PFCMPEQ
	case 0xB4:
		return PFMUL
	case 0xB6:
		return PFRCPIT2
	case 0xB7:
		return PMULHRW
	case 0xBB:
		return PSWAPD
	case 0xBF:
		return PAVGUSB
	}
	panic("Error: todo")
}
