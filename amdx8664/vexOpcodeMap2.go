package amdx8664

func vexOpcodeMap2(curByte byte, opcodeExt [2]bool, isRexW bool) (Instruction, bool, bool, MemSegment, Register, Register, Register, int) {
	switch curByte {
	case 0x0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSHUFB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHADDW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHADDD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHADDSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMADDUBSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHSUBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHSUBD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHSUBSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGNB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGNW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGND, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGND, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VTESTPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VTESTPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x13:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VCVTPH2PS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x16:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x17:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPTEST, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x18:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x19:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTF128, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPABSB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPABSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPABSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x20:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x21:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXBD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x22:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXBQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x23:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x24:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXWQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x25:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x28:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x29:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVNTDQA, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKUSDW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x30:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x31:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXBD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x32:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXBQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x33:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x34:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXWQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x35:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x36:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x37:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x38:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINSB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x39:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINUW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINUD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXSB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXUW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXUD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x40:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULLD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x41:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHMINPOSUW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x45:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x46:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRAVD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x47:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x50:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPDPBUSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x51:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPDPBUSDS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x52:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPDPWSSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x53:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPDPWSSDS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x58:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x59:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTI128, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x78:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x79:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x8C:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x8E:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x90:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x91:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x92:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x93:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x96:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x97:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x98:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x99:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9A:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9B:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9C:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9D:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9E:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9F:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA6:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA7:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA8:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA9:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAA:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAB:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAC:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAD:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAE:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAF:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB6:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB7:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB8:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB9:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xBA:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xBB:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xBC:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xBD:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xBE:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xBF:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xCF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VGF2P8MULB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESIMC, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESENC, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESENCLAST, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESDEC, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESDECLAST, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF2:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF3:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF5:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{true, false}:
			panic("todo")
		case [2]bool{true, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF6:
		switch opcodeExt {
		case [2]bool{true, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF7:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		case [2]bool{true, false}:
			panic("todo")
		case [2]bool{true, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	}
	panic("Error: todo")
}
