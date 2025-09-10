package amdx8664

func vexOpcodeMap1(curByte byte, opcodeExt [2]bool, isRexW bool) (Instruction, bool, bool, MemSegment, Register, Register, Register, int) {
	switch curByte {
	case 0x10:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVUPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVUPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("todo")
		case [2]bool{true, true}:
			panic("todo")
		}
	case 0x11:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVUPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVUPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VMOVLPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VMOVSLDUP, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VMOVDDUP, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x13:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVLPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVLPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x14:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VUNPCKLPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VUNPCKLPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x15:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VUNPCKHPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VUNPCKHPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VMOVHPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VMOVSHDUP, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x17:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVHPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVHPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x28:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVAPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVAPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x29:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVAPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVAPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VCVTSI2SS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VCVTSI2SD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x2B:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVNTPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVNTPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VUCOMISS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VUCOMISD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x2F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VCOMISS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VCOMISD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return KUNPCKBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x50:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMOVMSKPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMOVMSKPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x51:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VSQRTPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VSQRTPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VSQRTSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VSQRTSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x52:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VRSQRTPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			panic("Error: Unknown opcode")
		case [2]bool{true, false}:
			return VRSQRTSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x53:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VRCPPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			panic("Error: Unknown opcode")
		case [2]bool{true, false}:
			return VRCPPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x54:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VANDPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VANDPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x55:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VANDNPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VANDNPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x56:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VORPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VORPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x57:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VXORPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VXORPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			panic("Error: Unknown opcode")
		case [2]bool{true, true}:
			panic("Error: Unknown opcode")
		}
	case 0x58:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VADDPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VADDPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VADDSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VADDSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x59:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMULPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMULPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VMULSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VMULSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x5A:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VCVTPS2PD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VCVTPD2PS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VCVTSS2SD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VCVTSD2SS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x5B:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VCVTPS2PD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VCVTPD2PS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VCVTSS2SD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			panic("Error: Unknown instruction")
		}
	case 0x5C:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VSUBPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VSUBPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VSUBSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VSUBSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x5D:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMINPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMINPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VMINSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VMINSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x5E:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VDIVPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VDIVPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VDIVSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VDIVSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x5F:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VMAXPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VMAXPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VMAXSS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VMAXSD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0x60:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x61:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x62:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x63:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKSSWB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x64:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x65:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x66:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPGTD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x67:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKUSWB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x68:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x69:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6A:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6B:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPACKSSDW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKLQDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x6D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPUNPCKHQDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VMOVDQA, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VMOVDQU, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x70:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSHUFD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VPSHUFHW, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VPSHUFLW, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x71:
		switch opcodeExt {
		case [2]bool{false, true}:
			return NoInstruction, false, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x72:
		switch opcodeExt {
		case [2]bool{false, true}:
			return NoInstruction, false, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x73:
		switch opcodeExt {
		case [2]bool{false, true}:
			return NoInstruction, false, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x74:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x75:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x76:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPCMPEQD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7C:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VHADDPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VHADDPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x7D:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VHSUBPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VHSUBPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VMOVDQA, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VMOVDQU, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return KMOVW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return KMOVB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			panic("todo")
		default:
			panic("Error: Unknonw instruction")
		}
	case 0x93:
		switch opcodeExt {
		case [2]bool{false, false}:
			return KMOVW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return KMOVB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
			return VCMPPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VCMPPD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VCMPSS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VCMPSD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		}
	case 0xC4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPINSRW, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPEXTRW, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC6:
		switch opcodeExt {
		case [2]bool{false, false}:
			return VSHUFPS, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{false, true}:
			return VSHUFPD, true, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VADDSUBPD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VADDSUBPS, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRLW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRLD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRLQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULLW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVMSKB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMOVMSKB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBUSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINUB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPAND, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDUSB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDUSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXUB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xDF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPANDN, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE0:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPAVGB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRAW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSRAD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPAVGW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULHUW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULHW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VCVTTPD2DQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, false}:
			return VCVTDQ2PD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		case [2]bool{true, true}:
			return VCVTPD2DQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMOVNTDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBSB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xE9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMINSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPOR, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDSB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xED:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMAXSW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xEF:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPXOR, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF0:
		switch opcodeExt {
		case [2]bool{true, true}:
			return VLDDQU, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF1:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSLLW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF2:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSLLD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF3:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSLLQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF4:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMULUDQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF5:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPMADDWD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF6:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSADBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF7:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VMASKMOVDQU, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF8:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF9:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFA:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFB:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPSUBQ, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFC:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDB, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFD:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDW, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xFE:
		switch opcodeExt {
		case [2]bool{false, true}:
			return VPADDD, true, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
	panic("")
}

func vexOpcodeMap1ModRMG1(opcode byte, opcodeExt [2]bool, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, Register, int) {
	switch opcode {
	case 0x71:
		switch opcodeExt {
		case [2]bool{false, true}:
			switch modrmReg {
			case [3]bool{false, true, false}:
				return VPSRLW, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{true, false, false}:
				return VPSRAW, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{true, true, false}:
				return VPSLLW, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
				return VPSRLD, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{true, false, false}:
				return VPSRAD, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{true, true, false}:
				return VPSLLD, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
				return VPSRLQ, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{false, true, true}:
				return VPSRLDQ, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{true, true, false}:
				return VPSLLQ, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{true, true, true}:
				return VPSLLDQ, true, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
				return VLDMXCSR, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{false, true, true}:
				return VSTMXCSR, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
				return BLSR, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{false, true, false}:
				return BLSMSK, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
			case [3]bool{false, true, true}:
				return BLSI, false, NoSegment, NoRegister, NoRegister, NoRegister, 0
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
