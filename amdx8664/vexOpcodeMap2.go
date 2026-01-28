package amdx8664

func vexOpcodeMap2(curByte byte, opcodeExt [2]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSHUFB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHADDW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHADDD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHADDSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMADDUBSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHSUBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHSUBD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHSUBSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGNB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGNW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGND, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSIGND, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VTESTPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VTESTPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x13:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VCVTPH2PS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x16:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x17:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPTEST, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x18:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTSS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x19:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTF128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPABSB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPABSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPABSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x20:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x21:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXBD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x22:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXBQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x23:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x24:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXWQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x25:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVSXDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x28:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x29:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVNTDQA, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKUSDW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x30:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x31:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXBD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x32:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXBQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x33:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x34:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXWQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x35:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVZXDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x36:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x37:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x38:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINSB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x39:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINUW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINUD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXSB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXUW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x3F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXUD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x40:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULLD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x41:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPHMINPOSUW, []Operand{}
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
			return VPSRAVD, []Operand{}
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
			return VPDPBUSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x51:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPDPBUSDS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x52:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPDPWSSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x53:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPDPWSSDS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x58:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x59:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBROADCASTI128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x78:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x79:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBROADCASTW, []Operand{}
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
			return VGF2P8MULB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESIMC, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESENC, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESENCLAST, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESDEC, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESDECLAST, []Operand{}
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
