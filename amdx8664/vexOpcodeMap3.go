package amdx8664

func vexOpcodeMap3(curByte byte, opcodeExt [2]bool, isRexW bool) (Instruction, bool, bool, MemSegment, Register, Register, Register, int) {
	switch curByte {
	case 0x0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMQ, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMPD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBLENDD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERMILPD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERM2F128, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDPD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDSS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VROUNDSD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBLENDPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VBLENDPD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPBLENDW, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPALIGNR, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x14:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPEXTRB, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x15:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPEXTRW, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VEXTRACTPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x18:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VINSERTF128, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x19:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VEXTRACTF128, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x1D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VCVTPS2PH, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x20:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPINSRB, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x21:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VINSERTPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VINSERTI128, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x39:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VEXTRACTI128, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x40:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VDPPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x41:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VDPPD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x42:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMPSADBW, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x44:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCLMULQDQ, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x46:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPERM2I128, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VFMADDSUBPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDSUBPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBADDPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x5F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBADDPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x60:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPESTRM, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x61:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPESTRI, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x62:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPISTRM, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x63:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPISTRI, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x68:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x69:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMADDSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFMSUBSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x78:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x79:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMADDSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7E:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VFNMSUBSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xCE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VGF2P8AFFINEQB, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xCF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VGF2P8AFFINEINVQB, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VAESKEYGENASSIST, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF0:
		panic("todo")
	default:
		panic("Error: Unkown instruction")
	}
}
