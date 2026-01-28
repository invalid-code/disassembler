package amdx8664

func vexOpcodeMap1(curByte byte, opcodeExt [2]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x10:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVUPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVUPD, []Operand{}
		case [2]bool{true, false}:
			panic("todo")
		case [2]bool{true, true}:
			panic("todo")
		}
	case 0x11:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVUPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVUPD, []Operand{}
		case [2]bool{true, false}:
			panic("todo")
		case [2]bool{true, true}:
			panic("todo")
		}
	case 0x12:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			return VMOVLPD, []Operand{}
		case [2]bool{true, false}:
			return VMOVSLDUP, []Operand{}
		case [2]bool{true, true}:
			return VMOVDDUP, []Operand{}
		}
	case 0x13:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVLPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVLPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x14:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VUNPCKLPS, []Operand{}
		case [2]bool{false, true}:
			return VUNPCKLPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x15:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VUNPCKHPS, []Operand{}
		case [2]bool{false, true}:
			return VUNPCKHPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x16:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			return VMOVHPD, []Operand{}
		case [2]bool{true, false}:
			return VMOVSHDUP, []Operand{}
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x17:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVHPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVHPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x28:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVAPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVAPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x29:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVAPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVAPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x2A:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("Error: Unknown opcode")
		case [2]bool{false, true}:
			panic("Error: Unknown opcode")
		case [2]bool{true, false}:
			return VCVTSI2SS, []Operand{}
		case [2]bool{true, true}:
			return VCVTSI2SD, []Operand{}
		}
	case 0x2B:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVNTPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVNTPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x2C:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("Error: Unknown opcode")
		case [2]bool{false, true}:
			panic("Error: Unknown opcode")
		case [2]bool{true, false}:
			panic("todo")
		case [2]bool{true, true}:
			panic("todo")
		}
	case 0x2D:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("Error: Unknown opcode")
		case [2]bool{false, true}:
			panic("Error: Unknown opcode")
		case [2]bool{true, false}:
			panic("todo")
		case [2]bool{true, true}:
			panic("todo")
		}
	case 0x2E:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VUCOMISS, []Operand{}
		case [2]bool{false, true}:
			return VUCOMISD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x2F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VCOMISS, []Operand{}
		case [2]bool{false, true}:
			return VCOMISD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x41:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x42:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x44:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x45:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x46:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x47:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x48:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x49:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4A:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x4B:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			return KUNPCKBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x50:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVMSKPS, []Operand{}
		case [2]bool{false, true}:
			return VMOVMSKPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x51:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VSQRTPS, []Operand{}
		case [2]bool{false, true}:
			return VSQRTPD, []Operand{}
		case [2]bool{true, false}:
			return VSQRTSS, []Operand{}
		case [2]bool{true, true}:
			return VSQRTSD, []Operand{}
		}
	case 0x52:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VRSQRTPS, []Operand{}
		case [2]bool{false, true}:
			panic("Error: Unknown opcode")
		case [2]bool{true, false}:
			return VRSQRTSS, []Operand{}
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x53:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VRCPPS, []Operand{}
		case [2]bool{false, true}:
			panic("Error: Unknown opcode")
		case [2]bool{true, false}:
			return VRCPPS, []Operand{}
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x54:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VANDPS, []Operand{}
		case [2]bool{false, true}:
			return VANDPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x55:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VANDNPS, []Operand{}
		case [2]bool{false, true}:
			return VANDNPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x56:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VORPS, []Operand{}
		case [2]bool{false, true}:
			return VORPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x57:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VXORPS, []Operand{}
		case [2]bool{false, true}:
			return VXORPD, []Operand{}
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x58:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VADDPS, []Operand{}
		case [2]bool{false, true}:
			return VADDPD, []Operand{}
		case [2]bool{true, false}:
			return VADDSS, []Operand{}
		case [2]bool{true, true}:
			return VADDSD, []Operand{}
		}
	case 0x59:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMULPS, []Operand{}
		case [2]bool{false, true}:
			return VMULPD, []Operand{}
		case [2]bool{true, false}:
			return VMULSS, []Operand{}
		case [2]bool{true, true}:
			return VMULSD, []Operand{}
		}
	case 0x5A:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VCVTPS2PD, []Operand{}
		case [2]bool{false, true}:
			return VCVTPD2PS, []Operand{}
		case [2]bool{true, false}:
			return VCVTSS2SD, []Operand{}
		case [2]bool{true, true}:
			return VCVTSD2SS, []Operand{}
		}
	case 0x5B:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VCVTPS2PD, []Operand{}
		case [2]bool{false, true}:
			return VCVTPD2PS, []Operand{}
		case [2]bool{true, false}:
			return VCVTSS2SD, []Operand{}
		case [2]bool{true, true}:
			panic("Error: Unknown instruction")
		}
	case 0x5C:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VSUBPS, []Operand{}
		case [2]bool{false, true}:
			return VSUBPD, []Operand{}
		case [2]bool{true, false}:
			return VSUBSS, []Operand{}
		case [2]bool{true, true}:
			return VSUBSD, []Operand{}
		}
	case 0x5D:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMINPS, []Operand{}
		case [2]bool{false, true}:
			return VMINPD, []Operand{}
		case [2]bool{true, false}:
			return VMINSS, []Operand{}
		case [2]bool{true, true}:
			return VMINSD, []Operand{}
		}
	case 0x5E:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VDIVPS, []Operand{}
		case [2]bool{false, true}:
			return VDIVPD, []Operand{}
		case [2]bool{true, false}:
			return VDIVSS, []Operand{}
		case [2]bool{true, true}:
			return VDIVSD, []Operand{}
		}
	case 0x5F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMAXPS, []Operand{}
		case [2]bool{false, true}:
			return VMAXPD, []Operand{}
		case [2]bool{true, false}:
			return VMAXSS, []Operand{}
		case [2]bool{true, true}:
			return VMAXSD, []Operand{}
		}
	case 0x60:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x61:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x62:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x63:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKSSWB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x64:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x65:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x66:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x67:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKUSWB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x68:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x69:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKSSDW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLQDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHQDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6E:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVDQA, []Operand{}
		case [2]bool{true, false}:
			return VMOVDQU, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x70:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSHUFD, []Operand{}
		case [2]bool{true, false}:
			return VPSHUFHW, []Operand{}
		case [2]bool{true, true}:
			return VPSHUFLW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x71:
		switch opcodeExt {
		case [2]bool{false, true}:
			return NoInstruction, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x72:
		switch opcodeExt {
		case [2]bool{false, true}:
			return NoInstruction, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x73:
		switch opcodeExt {
		case [2]bool{false, true}:
			return NoInstruction, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x74:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x75:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x76:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VHADDPD, []Operand{}
		case [2]bool{true, true}:
			return VHADDPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VHSUBPD, []Operand{}
		case [2]bool{true, true}:
			return VHSUBPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7E:
		switch opcodeExt {
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7F:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVDQA, []Operand{}
		case [2]bool{true, false}:
			return VMOVDQU, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x90:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x91:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0x92:
		switch opcodeExt {
		case [2]bool{false, false}:
			return KMOVW, []Operand{}
		case [2]bool{false, true}:
			return KMOVB, []Operand{}
		case [2]bool{true, true}:
			panic("todo")
		default:
			panic("Error: Unknonw instruction")
		}
	case 0x93:
		switch opcodeExt {
		case [2]bool{false, false}:
			return KMOVW, []Operand{}
		case [2]bool{false, true}:
			return KMOVB, []Operand{}
		case [2]bool{true, true}:
			panic("todo")
		default:
			panic("Error: Unknonw instruction")
		}
	case 0x98:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknonw instruction")
		}
	case 0x99:
		switch opcodeExt {
		case [2]bool{false, false}:
			panic("todo")
		case [2]bool{false, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAE:
		if opcodeExt == [2]bool{false, false} {
			panic("todo")
		} else {
			panic("Error: Unknown instruction")
		}
	case 0xC2:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VCMPPS, []Operand{}
		case [2]bool{false, true}:
			return VCMPPD, []Operand{}
		case [2]bool{true, false}:
			return VCMPSS, []Operand{}
		case [2]bool{true, true}:
			return VCMPSD, []Operand{}
		}
	case 0xC4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPINSRW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPEXTRW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VSHUFPS, []Operand{}
		case [2]bool{false, true}:
			return VSHUFPD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VADDSUBPD, []Operand{}
		case [2]bool{true, true}:
			return VADDSUBPS, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRLW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRLD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRLQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULLW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVMSKB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVMSKB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBUSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINUB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPAND, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDUSB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDUSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXUB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPANDN, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPAVGB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRAW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRAD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPAVGW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULHUW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULHW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VCVTTPD2DQ, []Operand{}
		case [2]bool{true, false}:
			return VCVTDQ2PD, []Operand{}
		case [2]bool{true, true}:
			return VCVTPD2DQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVNTDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBSB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPOR, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDSB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xED:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXSW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPXOR, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF0:
		switch opcodeExt {
		case [2]bool{true, true}:
			return VLDDQU, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSLLW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSLLD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSLLQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULUDQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMADDWD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSADBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVDQU, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBQ, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDB, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDW, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDD, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
	panic("")
}

func vexOpcodeMap1ModRMG1(opcode byte, opcodeExt [2]bool, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x71:
		switch opcodeExt {
		case [2]bool{false, true}:
			switch modrmReg {
			case [3]bool{false, true, false}:
				return VPSRLW, []Operand{}
			case [3]bool{true, false, false}:
				return VPSRAW, []Operand{}
			case [3]bool{true, true, false}:
				return VPSLLW, []Operand{}
			default:
				panic("Error: Unknown instruction")
			}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x72:
		switch opcodeExt {
		case [2]bool{false, true}:
			switch modrmReg {
			case [3]bool{false, true, false}:
				return VPSRLD, []Operand{}
			case [3]bool{true, false, false}:
				return VPSRAD, []Operand{}
			case [3]bool{true, true, false}:
				return VPSLLD, []Operand{}
			default:
				panic("Error: Unknown instruction")
			}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x73:
		switch opcodeExt {
		case [2]bool{false, true}:
			switch modrmReg {
			case [3]bool{false, true, false}:
				return VPSRLQ, []Operand{}
			case [3]bool{false, true, true}:
				return VPSRLDQ, []Operand{}
			case [3]bool{true, true, false}:
				return VPSLLQ, []Operand{}
			case [3]bool{true, true, true}:
				return VPSLLDQ, []Operand{}
			default:
				panic("Error: Unknown instruction")
			}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xAE:
		switch opcodeExt {
		case [2]bool{false, false}:
			switch modrmReg {
			case [3]bool{false, true, false}:
				return VLDMXCSR, []Operand{}
			case [3]bool{false, true, true}:
				return VSTMXCSR, []Operand{}
			default:
				panic("Error: Unknown instruction")
			}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF3:
		switch opcodeExt {
		case [2]bool{false, false}:
			switch modrmReg {
			case [3]bool{false, false, true}:
				return BLSR, []Operand{}
			case [3]bool{false, true, false}:
				return BLSMSK, []Operand{}
			case [3]bool{false, true, true}:
				return BLSI, []Operand{}
			default:
				panic("Error: Unknown instruction")
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
