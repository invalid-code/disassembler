package amdx8664

func vexOpcodeMap3(curByte byte, opcodeExt [2]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBLENDD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERM2F128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDSS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBLENDPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBLENDPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBLENDW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPALIGNR, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x14:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPEXTRB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x15:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPEXTRW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x16:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x17:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VEXTRACTPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x18:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VINSERTF128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x19:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VEXTRACTF128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VCVTPS2PH, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x20:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPINSRB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x21:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VINSERTPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x22:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x30:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x31:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x32:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x33:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x38:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VINSERTI128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x39:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VEXTRACTI128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x40:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VDPPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x41:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VDPPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x42:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMPSADBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x44:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCLMULQDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x46:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERM2I128, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4A:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4B:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4C:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDSUBPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDSUBPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBADDPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBADDPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x60:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPESTRM, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x61:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPESTRI, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x62:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPISTRM, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x63:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPISTRI, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x68:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x69:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDSS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBSS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x78:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x79:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDSS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBSS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBSD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xCE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VGF2P8AFFINEQB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xCF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VGF2P8AFFINEINVQB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESKEYGENASSIST, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF0:
		panic("todo")
	default:
		panic("Error: Unkown instruction")
	}
}
