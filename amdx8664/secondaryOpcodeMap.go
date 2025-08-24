package amdx8664

func secondaryOpcodeMap(curByte byte, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, bool, bool, MemSegment, Register, Register, int) {
	switch curByte {
	case 0x0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LAR, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LSL, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x4:
		panic("Error: No opcode")
	case 0x5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSCALL, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return CLTS, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSRET, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return INVD, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		panic("todo multiple instructions")
	case 0xA:
		panic("Error: No opcode")
	case 0xB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return UD2, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0xC:
		panic("Error: No opcode")
	case 0xD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		panic("todo multiple instructions")
	case 0xE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return FEMMS, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0xF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		panic("Error: This is the 3dnow opcode map escape sequence")
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MOVSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MOVSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MOVUPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x11:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MOVSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MOVSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MOVUPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x12:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MOVSLDUP, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MOVDDUP, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("todo multiple instructions in opcode")
		}
	case 0x13:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVLPD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKLPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return UNPCKLPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKHPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return UNPCKHPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x16:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MOVSHDUP, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			panic("todo multiple instructions in opcode")
		}
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVHPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1E, 0x1F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		return NoInstruction, true, true, NoSegment, NoRegister, NoRegister, 0
	case 0x1D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		panic("todo weird")
	case 0x20, 0x21, 0x22, 0x23:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x24, 0x25, 0x26, 0x27:
		panic("Error: Unkown opcode")
	case 0x28, 0x29:
		if isOperandSizeOverride {
			// 0x66
			return MOVAPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVAPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x2A:
		if isOperandSizeOverride {
			// 0x66
			return CVTSI2SS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return CVTPI2PD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CVTSI2SD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			// todo
			return CVTPI2PS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x2B:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MOVNTSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MOVNTSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			// todo
			return MOVNTPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x2C:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2PI, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return CVTTSS2SI, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CVTTSD2SI, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return CVTTPS2PI, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x2D:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PI, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return CVTSS2SI, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CVTSD2SI, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return CVTPS2PI, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x2E:
		if isOperandSizeOverride {
			// 0x66
			return UCOMISD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return UCOMISS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x2F:
		if isOperandSizeOverride {
			// 0x66
			return COMISD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return COMISS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x30:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return WRMSR, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x31:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDTSC, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x32:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDMSR, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x33:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDPMC, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x34:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		if !is64Bit {
			panic("Error: Invalid in 32-bit")
		}
		return SYSENTER, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x35:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		if !is64Bit {
			panic("Error: Invalid in 32-bit")
		}
		return SYSEXIT, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x36, 0x37:
		panic("Error: Unkown opcode")
	case 0x38:
		panic("Error: This is the 0x38 escape byte")
	case 0x39:
		panic("Error: Unkown opcode")
	case 0x3A:
		panic("Error: This is the 0x3A escape byte")
	case 0x3B, 0x3C, 0x3D, 0x3E, 0x3F:
		panic("Error: Unkown opcode")
	case 0x40:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVO, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x41:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNO, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x42:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVB, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x43:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNB, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x44:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVZ, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x45:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNZ, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x46:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVBE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x47:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNBE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x48:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x49:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x4A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVP, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x4B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNP, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x4C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVL, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x4D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNL, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x4E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVLE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x4F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNLE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x50:
		if isOperandSizeOverride {
			// 0x66
			return MOVMSKPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVMSKPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x51:
		if isOperandSizeOverride {
			// 0x66
			return SQRTPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return SQRTSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return SQRTSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return SQRTPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x52:
		if isOperandSizeOverride || isRep0 {
			// 0x66
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return RSQRTSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return RSQRTPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x53:
		if isOperandSizeOverride || isRep0 {
			// 0x66
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return RCPSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return RCPPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x54:
		if isOperandSizeOverride {
			// 0x66
			return ANDPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return ANDPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x55:
		if isOperandSizeOverride {
			// 0x66
			return ANDNPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return ANDNPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x56:
		if isOperandSizeOverride {
			// 0x66
			return ORPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return ORPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x57:
		if isOperandSizeOverride {
			// 0x66
			return XORPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return XORPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x58:
		if isOperandSizeOverride {
			// 0x66
			return ADDPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return ADDSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return ADDSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return ADDPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x59:
		if isOperandSizeOverride {
			// 0x66
			return MULPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MULSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MULSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MULPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x5A:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return CVTSS2SD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CVTSD2SS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return CVTPS2PD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x5B:
		if isOperandSizeOverride {
			// 0x66
			return CVTPS2DQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return CVTTPS2DQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return CVTDQ2PS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x5C:
		if isOperandSizeOverride {
			// 0x66
			return SUBPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return SUBSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return SUBSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return SUBPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x5D:
		if isOperandSizeOverride {
			// 0x66
			return MINPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MINSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MINSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MINPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x5E:
		if isOperandSizeOverride {
			// 0x66
			return DIVPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return DIVSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return DIVSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return DIVPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x5F:
		if isOperandSizeOverride {
			// 0x66
			return MAXPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MAXSS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MAXSD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MAXPS, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x60:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLBW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKLBW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x61:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLWD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKLWD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x62:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKLDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x63:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSWB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PACKSSWB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x64:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPGTB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x65:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPGTW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x66:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPGTD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x67:
		if isOperandSizeOverride {
			// 0x66
			return PACKUSWB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PACKUSWB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x68:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHBW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKHBW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x69:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHWD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKHWD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x6A:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKHDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x6B:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSDW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PACKSSDW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x6C:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLQDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x6D:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHQDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x6E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			return MOVD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x6F:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MOVDQU, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVDQA, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x70:
		panic("todo 3 operand types")
		// if isOperandSizeOverride {
		// 	// 0x66
		// 	return PSHUFD
		// } else if isRep1 {
		// 	// 0xF3
		// 	return PSHUFHW
		// } else if isRep0 {
		// 	// 0xF2
		// 	return PSHUFLW
		// } else {
		// 	return PSHUFW
		// }
	case 0x71, 0x72, 0x73:
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x74:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPEQB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x75:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPEQW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x76:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPEQD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x77:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x77")
		}
		return EMMS, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0x78:
		if isOperandSizeOverride {
			// 0x66
			return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			// return INSERTQ, true, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo multiple operands")
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x79:
		if isOperandSizeOverride {
			// 0x66
			return EXTRQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return INSERTQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x7A, 0x7B:
		panic("Error: Unkown opcode")
	case 0x7C:
		if isOperandSizeOverride {
			// 0x66
			return HADDPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return HADDPS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x7D:
		if isOperandSizeOverride {
			// 0x66
			return HSUBPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return HSUBPS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x7E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return MOVQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MOVD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x7F:
		if isOperandSizeOverride {
			// 0x66
			return MOVDQA, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return MOVDQU, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return MOVQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0x80:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x81:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x82:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x83:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x84:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x85:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x86:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x87:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x88:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x89:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x90:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETO, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x91:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNO, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x92:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETB, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x93:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNB, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x94:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETZ, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x95:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNZ, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x96:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETBE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x97:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNBE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x98:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x99:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETP, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNP, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETL, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNL, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETLE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0x9F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNLE, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xA0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH, false, false, FS, NoRegister, NoRegister, 0
	case 0xA1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP, false, false, FS, NoRegister, NoRegister, 0
	case 0xA2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return CPUID, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0xA3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BT, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0xA4:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 operands")
	case 0xA5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 operands")
	case 0xA6, 0xA7:
		panic("Error: Unkown opcode")
	case 0xA8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH, false, false, GS, NoRegister, NoRegister, 0
	case 0xA9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP, false, false, GS, NoRegister, NoRegister, 0
	case 0xAA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return RSM, false, false, NoSegment, NoRegister, NoRegister, 0
	case 0xAB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BTS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xAC:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 opcodes")
	case 0xAD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 opcodes")
	case 0xAE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xAF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return IMUL, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LSS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BTR, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB4:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LFS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LGS, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xB8:
		if isRep1 {
			// 0xF3
			return POPCNT, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: 0xF3 prefix only allowed")
		}
	case 0xB9:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xBA:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xBB:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return BTC, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xBC:
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return TZCNT, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return BSF, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xBD:
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return LZCNT, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			return BSR, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xBE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xBF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xC0:
		return XADD, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xC1:
		return XADD, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xC2:
		if isOperandSizeOverride {
			// 0x66
			return CMPPD, true, true, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return CMPPS, true, true, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CMPSD, true, true, NoSegment, NoRegister, NoRegister, 0
		} else {
			return CMPPS, true, true, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xC3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: Unkown opcode")
		}
		return MOVNTI, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xC4:
		if isOperandSizeOverride {
			// 0x66
			// return PINSRW, true, true, NoSegment, NoRegister, NoRegister, 0
			panic("todo multiple operands")
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			// return PINSRW, true, true, NoSegment, NoRegister, NoRegister, 0
			panic("todo multiple operands")
		}
	case 0xC5:
		if isOperandSizeOverride {
			// 0x66
			return PEXTRW, true, true, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PEXTRW, true, true, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xC6:
		if isOperandSizeOverride {
			// 0x66
			return SHUFPD, true, true, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return SHUFPS, true, true, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xC7:
		// todo modrm byte
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xC8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R8, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RAX, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, EAX, NoRegister, 0
		}
	case 0xC9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R9, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RCX, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, ECX, NoRegister, 0
		}
	case 0xCA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R10, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RDX, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, EDX, NoRegister, 0
		}
	case 0xCB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R11, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RBX, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, EBX, NoRegister, 0
		}
	case 0xCC:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R12, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RSP, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, ESP, NoRegister, 0
		}
	case 0xCD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R13, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RBP, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, EBP, NoRegister, 0
		}
	case 0xCE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R14, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RSI, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, ESI, NoRegister, 0
		}
	case 0xCF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, NoSegment, R15, NoRegister, 0
			} else {
				return BSWAP, false, false, NoSegment, RDI, NoRegister, 0
			}
		} else {
			return BSWAP, false, false, NoSegment, EDI, NoRegister, 0
		}
	case 0xD0:
		if isOperandSizeOverride {
			// 0x66
			return ADDSUBPD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return ADDSUBPS, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown Opcode")
		}
	case 0xD1:
		if isOperandSizeOverride {
			// 0x66
			return PSRLW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSRLW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xD2:
		if isOperandSizeOverride {
			// 0x66
			return PSRLD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSRLD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xD3:
		if isOperandSizeOverride {
			// 0x66
			return PSRLQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSRLQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xD4:
		if isOperandSizeOverride {
			// 0x66
			return PADDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xD5:
		if isOperandSizeOverride {
			// 0x66
			return PMULLW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMULLW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xD6:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return MOVQ2DQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return MOVDQ2Q, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown Opcode")
		}
	case 0xD7:
		if isOperandSizeOverride {
			// 0x66
			return PMOVMSKB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMOVMSKB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xD8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBUSB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xD9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBUSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xDA:
		if isOperandSizeOverride {
			// 0x66
			return PMINUB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMINUB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xDB:
		if isOperandSizeOverride {
			// 0x66
			return PAND, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PAND, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xDC:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDUSB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xDD:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDUSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xDE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMAXUB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return PANDN, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PANDN, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE0:
		if isOperandSizeOverride {
			// 0x66
			return PAVGB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PAVGB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE1:
		if isOperandSizeOverride {
			// 0x66
			return PSRAW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSRAW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE2:
		if isOperandSizeOverride {
			// 0x66
			return PSRAD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSRAD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE3:
		if isOperandSizeOverride {
			// 0x66
			return PAVGW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PAVGW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE4:
		if isOperandSizeOverride {
			// 0x66
			return PMULHUW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMULHUW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE5:
		if isOperandSizeOverride {
			// 0x66
			return PMULHW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMULHW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE6:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2DQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 {
			// 0xF3
			return CVTDQ2PD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep0 {
			// 0xF2
			return CVTPD2DQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else {
			panic("Error: Unknown Opcode")
		}
	case 0xE7:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return MOVNTQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBSB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xE9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xEA:
		if isOperandSizeOverride {
			// 0x66
			return PMINSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMINSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xEB:
		if isOperandSizeOverride {
			// 0x66
			return POR, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return POR, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xEC:
		if isOperandSizeOverride {
			// 0x66
			return PADDSB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDSB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xED:
		if isOperandSizeOverride {
			// 0x66
			return PADDSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xEE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMAXSW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xEF:
		if isOperandSizeOverride {
			// 0x66
			return PXOR, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PXOR, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xF0")
		}
		return LDDQU, true, false, NoSegment, NoRegister, NoRegister, 0
	case 0xF1:
		if isOperandSizeOverride {
			// 0x66
			return PSLLW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSLLW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF2:
		if isOperandSizeOverride {
			// 0x66
			return PSLLD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSLLD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF3:
		if isOperandSizeOverride {
			// 0x66
			return PSLLQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSLLQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF4:
		if isOperandSizeOverride {
			// 0x66
			return PMULUDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMULUDQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF5:
		if isOperandSizeOverride {
			// 0x66
			return PMADDWD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PMADDWD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF6:
		if isOperandSizeOverride {
			// 0x66
			return PSADBW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSADBW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF7:
		if isOperandSizeOverride {
			// 0x66
			return MASKMOVDQU, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return MASKMOVQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xF9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xFA:
		if isOperandSizeOverride {
			// 0x66
			return PSUBD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xFB:
		if isOperandSizeOverride {
			// 0x66
			return PSUBQ, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PSUBQ, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xFC:
		if isOperandSizeOverride {
			// 0x66
			return PADDB, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDB, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xFD:
		if isOperandSizeOverride {
			// 0x66
			return PADDW, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDW, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xFE:
		if isOperandSizeOverride {
			// 0x66
			return PADDD, true, false, NoSegment, NoRegister, NoRegister, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown Opcode")
		} else {
			return PADDD, true, false, NoSegment, NoRegister, NoRegister, 0
		}
	case 0xFF:
		return UD0, false, false, NoSegment, NoRegister, NoRegister, 0
	default:
		panic("Error: Unknown Opcode")
	}
}
