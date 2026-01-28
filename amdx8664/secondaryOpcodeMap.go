package amdx8664

func secondaryOpcodeMap(curByte byte, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	switch curByte {
	case 0x0, 0x1:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return NoInstruction, []Operand{} // todo the opcode map doesn't specify operands but in the mod.r/m opcode extension map specifies the operand
	case 0x2:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LAR, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x3:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LSL, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x4:
		panic("Error: No opcode")
	case 0x5:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSCALL, []Operand{}
	case 0x6:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return CLTS, []Operand{}
	case 0x7:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSRET, []Operand{}
	case 0x8:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return INVD, []Operand{}
	case 0x9:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		panic("todo multiple instructions")
	case 0xA:
		panic("Error: No opcode")
	case 0xB:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return UD2, []Operand{}
	case 0xC:
		panic("Error: No opcode")
	case 0xD:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		panic("todo multiple instructions")
	case 0xE:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return FEMMS, []Operand{}
	case 0xF:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		panic("Error: This is the 3dnow opcode map escape sequence")
	case 0x10:
		if isOperandSizeOveride {
			// 0x66
			return MOVUPD, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep1 {
			// 0xF3
			return MOVSS, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep0 {
			// 0xF2
			return MOVSD, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else {
			return MOVUPS, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		}
	case 0x11:
		if isOperandSizeOveride {
			// 0x66
			return MOVUPD, []Operand{modRMRegOperand(false, false), modRMRegOperand(true, false)}
		} else if isRep1 {
			// 0xF3
			return MOVSS, []Operand{modRMRegOperand(false, false), modRMRegOperand(true, false)}
		} else if isRep0 {
			// 0xF2
			return MOVSD, []Operand{modRMRegOperand(false, false), modRMRegOperand(true, false)}
		} else {
			return MOVUPS, []Operand{modRMRegOperand(false, false), modRMRegOperand(true, false)}
		}
	case 0x12:
		if isOperandSizeOveride {
			// 0x66
			return MOVLPD, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep1 {
			// 0xF3
			return MOVSLDUP, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep0 {
			// 0xF2
			return MOVDDUP, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else {
			panic("todo multiple instructions in opcode")
		}
	case 0x13:
		if isOperandSizeOveride {
			// 0x66
			return MOVLPS, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, false)}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVLPD, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, false)}
		}
	case 0x14:
		if isOperandSizeOveride {
			// 0x66
			return UNPCKLPD, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		} else {
			return UNPCKLPS, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		}
	case 0x15:
		if isOperandSizeOveride {
			// 0x66
			return UNPCKHPD, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		} else {
			return UNPCKHPS, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		}
	case 0x16:
		if isOperandSizeOveride {
			// 0x66
			return MOVHPD, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep1 {
			// 0xF3
			return MOVSHDUP, []Operand{modRMRegOperand(true, false), modRMRegOperand(false, false)}
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			panic("todo multiple instructions in opcode")
		}
	case 0x17:
		if isOperandSizeOveride {
			// 0x66
			return MOVHPD, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, false)}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVHPS, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, false)}
		}
	case 0x18:
		if isRep0 || isRep1 || !isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		return NoInstruction, []Operand{modRMRegOperand(false, true)} // todo the opcode map doesn't specify operands but in the mod.r/m opcode extension map specifies the operand
	case 0x19, 0x1A, 0x1B, 0x1C, 0x1E, 0x1F:
		return NOP, []Operand{}
	case 0x1D:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		panic("todo weird")
	case 0x20, 0x21, 0x22, 0x23:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV, []Operand{modRMRegOperand(false, true), modRMRegOperand(false, false)}
	case 0x24, 0x25, 0x26, 0x27:
		panic("Error: Unkown opcode")
	case 0x28, 0x29:
		if isOperandSizeOveride {
			// 0x66
			return MOVAPD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVAPS, []Operand{}
		}
	case 0x2A:
		if isOperandSizeOveride {
			// 0x66
			return CVTSI2SS, []Operand{}
		} else if isRep1 {
			// 0xF3
			return CVTPI2PD, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CVTSI2SD, []Operand{}
		} else {
			// todo
			return CVTPI2PS, []Operand{}
		}
	case 0x2B:
		if isOperandSizeOveride {
			// 0x66
			return MOVNTPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return MOVNTSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return MOVNTSD, []Operand{}
		} else {
			// todo
			return MOVNTPS, []Operand{}
		}
	case 0x2C:
		if isOperandSizeOveride {
			// 0x66
			return CVTTPD2PI, []Operand{}
		} else if isRep1 {
			// 0xF3
			return CVTTSS2SI, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CVTTSD2SI, []Operand{}
		} else {
			return CVTTPS2PI, []Operand{}
		}
	case 0x2D:
		if isOperandSizeOveride {
			// 0x66
			return CVTPD2PI, []Operand{}
		} else if isRep1 {
			// 0xF3
			return CVTSS2SI, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CVTSD2SI, []Operand{}
		} else {
			return CVTPS2PI, []Operand{}
		}
	case 0x2E:
		if isOperandSizeOveride {
			// 0x66
			return UCOMISD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		} else {
			return UCOMISS, []Operand{}
		}
	case 0x2F:
		if isOperandSizeOveride {
			// 0x66
			return COMISD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		} else {
			return COMISS, []Operand{}
		}
	case 0x30:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return WRMSR, []Operand{}
	case 0x31:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDTSC, []Operand{}
	case 0x32:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDMSR, []Operand{}
	case 0x33:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDPMC, []Operand{}
	case 0x34:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		if !is64Bit {
			panic("Error: Invalid in 32-bit")
		}
		return SYSENTER, []Operand{}
	case 0x35:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		if !is64Bit {
			panic("Error: Invalid in 32-bit")
		}
		return SYSEXIT, []Operand{}
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
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVO, []Operand{}
	case 0x41:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNO, []Operand{}
	case 0x42:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVB, []Operand{}
	case 0x43:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNB, []Operand{}
	case 0x44:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVZ, []Operand{}
	case 0x45:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNZ, []Operand{}
	case 0x46:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVBE, []Operand{}
	case 0x47:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNBE, []Operand{}
	case 0x48:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVS, []Operand{}
	case 0x49:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNS, []Operand{}
	case 0x4A:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVP, []Operand{}
	case 0x4B:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNP, []Operand{}
	case 0x4C:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVL, []Operand{}
	case 0x4D:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNL, []Operand{}
	case 0x4E:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVLE, []Operand{}
	case 0x4F:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNLE, []Operand{}
	case 0x50:
		if isOperandSizeOveride {
			// 0x66
			return MOVMSKPD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVMSKPS, []Operand{}
		}
	case 0x51:
		if isOperandSizeOveride {
			// 0x66
			return SQRTPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return SQRTSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return SQRTSD, []Operand{}
		} else {
			return SQRTPS, []Operand{}
		}
	case 0x52:
		if isOperandSizeOveride || isRep0 {
			// 0x66
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return RSQRTSS, []Operand{}
		} else {
			return RSQRTPS, []Operand{}
		}
	case 0x53:
		if isOperandSizeOveride || isRep0 {
			// 0x66
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return RCPSS, []Operand{}
		} else {
			return RCPPS, []Operand{}
		}
	case 0x54:
		if isOperandSizeOveride {
			// 0x66
			return ANDPD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return ANDPS, []Operand{}
		}
	case 0x55:
		if isOperandSizeOveride {
			// 0x66
			return ANDNPD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return ANDNPS, []Operand{}
		}
	case 0x56:
		if isOperandSizeOveride {
			// 0x66
			return ORPD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return ORPS, []Operand{}
		}
	case 0x57:
		if isOperandSizeOveride {
			// 0x66
			return XORPD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return XORPS, []Operand{}
		}
	case 0x58:
		if isOperandSizeOveride {
			// 0x66
			return ADDPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return ADDSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return ADDSD, []Operand{}
		} else {
			return ADDPS, []Operand{}
		}
	case 0x59:
		if isOperandSizeOveride {
			// 0x66
			return MULPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return MULSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return MULSD, []Operand{}
		} else {
			return MULPS, []Operand{}
		}
	case 0x5A:
		if isOperandSizeOveride {
			// 0x66
			return CVTPD2PS, []Operand{}
		} else if isRep1 {
			// 0xF3
			return CVTSS2SD, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CVTSD2SS, []Operand{}
		} else {
			return CVTPS2PD, []Operand{}
		}
	case 0x5B:
		if isOperandSizeOveride {
			// 0x66
			return CVTPS2DQ, []Operand{}
		} else if isRep1 {
			// 0xF3
			return CVTTPS2DQ, []Operand{}
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return CVTDQ2PS, []Operand{}
		}
	case 0x5C:
		if isOperandSizeOveride {
			// 0x66
			return SUBPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return SUBSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return SUBSD, []Operand{}
		} else {
			return SUBPS, []Operand{}
		}
	case 0x5D:
		if isOperandSizeOveride {
			// 0x66
			return MINPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return MINSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return MINSD, []Operand{}
		} else {
			return MINPS, []Operand{}
		}
	case 0x5E:
		if isOperandSizeOveride {
			// 0x66
			return DIVPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return DIVSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return DIVSD, []Operand{}
		} else {
			return DIVPS, []Operand{}
		}
	case 0x5F:
		if isOperandSizeOveride {
			// 0x66
			return MAXPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return MAXSS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return MAXSD, []Operand{}
		} else {
			return MAXPS, []Operand{}
		}
	case 0x60:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKLBW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKLBW, []Operand{}
		}
	case 0x61:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKLWD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKLWD, []Operand{}
		}
	case 0x62:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKLDQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKLDQ, []Operand{}
		}
	case 0x63:
		if isOperandSizeOveride {
			// 0x66
			return PACKSSWB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PACKSSWB, []Operand{}
		}
	case 0x64:
		if isOperandSizeOveride {
			// 0x66
			return PCMPGTB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPGTB, []Operand{}
		}
	case 0x65:
		if isOperandSizeOveride {
			// 0x66
			return PCMPGTW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPGTW, []Operand{}
		}
	case 0x66:
		if isOperandSizeOveride {
			// 0x66
			return PCMPGTD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPGTD, []Operand{}
		}
	case 0x67:
		if isOperandSizeOveride {
			// 0x66
			return PACKUSWB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PACKUSWB, []Operand{}
		}
	case 0x68:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKHBW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKHBW, []Operand{}
		}
	case 0x69:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKHWD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKHWD, []Operand{}
		}
	case 0x6A:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKHDQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PUNPCKHDQ, []Operand{}
		}
	case 0x6B:
		if isOperandSizeOveride {
			// 0x66
			return PACKSSDW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PACKSSDW, []Operand{}
		}
	case 0x6C:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKLQDQ, []Operand{}
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x6D:
		if isOperandSizeOveride {
			// 0x66
			return PUNPCKHQDQ, []Operand{}
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x6E:
		if isOperandSizeOveride {
			// 0x66
			return MOVD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			return MOVD, []Operand{}
		}
	case 0x6F:
		if isOperandSizeOveride {
			// 0x66
			return MOVQ, []Operand{}
		} else if isRep1 {
			// 0xF3
			return MOVDQU, []Operand{}
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return MOVDQA, []Operand{}
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
		return NoInstruction, []Operand{}
	case 0x74:
		if isOperandSizeOveride {
			// 0x66
			return PCMPEQB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPEQB, []Operand{}
		}
	case 0x75:
		if isOperandSizeOveride {
			// 0x66
			return PCMPEQW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPEQW, []Operand{}
		}
	case 0x76:
		if isOperandSizeOveride {
			// 0x66
			return PCMPEQD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PCMPEQD, []Operand{}
		}
	case 0x77:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x77")
		}
		return EMMS, []Operand{}
	case 0x78:
		if isOperandSizeOveride {
			// 0x66
			return NoInstruction, []Operand{}
		} else if isRep0 {
			// 0xF2
			// return INSERTQ, []Operand{}
			panic("todo multiple operands")
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x79:
		if isOperandSizeOveride {
			// 0x66
			return EXTRQ, []Operand{}
		} else if isRep0 {
			// 0xF2
			return INSERTQ, []Operand{}
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x7A, 0x7B:
		panic("Error: Unkown opcode")
	case 0x7C:
		if isOperandSizeOveride {
			// 0x66
			return HADDPD, []Operand{}
		} else if isRep0 {
			// 0xF2
			return HADDPS, []Operand{}
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x7D:
		if isOperandSizeOveride {
			// 0x66
			return HSUBPD, []Operand{}
		} else if isRep0 {
			// 0xF2
			return HSUBPS, []Operand{}
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else {
			panic("Error: Unkown opcode")
		}
	case 0x7E:
		if isOperandSizeOveride {
			// 0x66
			return MOVD, []Operand{}
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return MOVQ, []Operand{}
		} else {
			return MOVD, []Operand{}
		}
	case 0x7F:
		if isOperandSizeOveride {
			// 0x66
			return MOVDQA, []Operand{}
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return MOVDQU, []Operand{}
		} else {
			return MOVQ, []Operand{}
		}
	case 0x80:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x81:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x82:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x83:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x84:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x85:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x86:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x87:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x88:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x89:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8A:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8B:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8C:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8D:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8E:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x8F:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		panic("todo dont know how to deal with rip addressing")
	case 0x90:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETO, []Operand{}
	case 0x91:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNO, []Operand{}
	case 0x92:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETB, []Operand{}
	case 0x93:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNB, []Operand{}
	case 0x94:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETZ, []Operand{}
	case 0x95:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNZ, []Operand{}
	case 0x96:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETBE, []Operand{}
	case 0x97:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNBE, []Operand{}
	case 0x98:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETS, []Operand{}
	case 0x99:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNS, []Operand{}
	case 0x9A:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETP, []Operand{}
	case 0x9B:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNP, []Operand{}
	case 0x9C:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETL, []Operand{}
	case 0x9D:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNL, []Operand{}
	case 0x9E:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETLE, []Operand{}
	case 0x9F:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNLE, []Operand{}
	case 0xA0:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH, []Operand{}
	case 0xA1:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP, []Operand{}
	case 0xA2:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return CPUID, []Operand{}
	case 0xA3:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BT, []Operand{}
	case 0xA4:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 operands")
	case 0xA5:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 operands")
	case 0xA6, 0xA7:
		panic("Error: Unkown opcode")
	case 0xA8:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH, []Operand{}
	case 0xA9:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP, []Operand{}
	case 0xAA:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return RSM, []Operand{}
	case 0xAB:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BTS, []Operand{}
	case 0xAC:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 opcodes")
	case 0xAD:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		panic("todo 3 opcodes")
	case 0xAE:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return NoInstruction, []Operand{}
	case 0xAF:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return IMUL, []Operand{}
	case 0xB0:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG, []Operand{}
	case 0xB1:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG, []Operand{}
	case 0xB2:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LSS, []Operand{}
	case 0xB3:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BTR, []Operand{}
	case 0xB4:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LFS, []Operand{}
	case 0xB5:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LGS, []Operand{}
	case 0xB6:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX, []Operand{}
	case 0xB7:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX, []Operand{}
	case 0xB8:
		if isRep1 {
			// 0xF3
			return POPCNT, []Operand{}
		} else {
			panic("Error: 0xF3 prefix only allowed")
		}
	case 0xB9:
		if !(isRep0 || isRep1 || isOperandSizeOveride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return NoInstruction, []Operand{}
	case 0xBA:
		if !(isRep0 || isRep1 || isOperandSizeOveride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return NoInstruction, []Operand{}
	case 0xBB:
		if !(isRep0 || isRep1 || isOperandSizeOveride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return BTC, []Operand{}
	case 0xBC:
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return TZCNT, []Operand{}
		} else {
			return BSF, []Operand{}
		}
	case 0xBD:
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else if isRep1 {
			// 0xF3
			return LZCNT, []Operand{}
		} else {
			return BSR, []Operand{}
		}
	case 0xBE:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX, []Operand{}
	case 0xBF:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX, []Operand{}
	case 0xC0:
		return XADD, []Operand{}
	case 0xC1:
		return XADD, []Operand{}
	case 0xC2:
		if isOperandSizeOveride {
			// 0x66
			return CMPPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			return CMPPS, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CMPSD, []Operand{}
		} else {
			return CMPPS, []Operand{}
		}
	case 0xC3:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: Unkown opcode")
		}
		return MOVNTI, []Operand{}
	case 0xC4:
		if isOperandSizeOveride {
			// 0x66
			// return PINSRW, []Operand{}
			panic("todo multiple operands")
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			// return PINSRW, []Operand{}
			panic("todo multiple operands")
		}
	case 0xC5:
		if isOperandSizeOveride {
			// 0x66
			return PEXTRW, []Operand{}
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return PEXTRW, []Operand{}
		}
	case 0xC6:
		if isOperandSizeOveride {
			// 0x66
			return SHUFPD, []Operand{}
		} else if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		} else if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		} else {
			return SHUFPS, []Operand{}
		}
	case 0xC7:
		// todo modrm byte
		return NoInstruction, []Operand{}
	case 0xC8:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xC9:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xCA:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xCB:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xCC:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xCD:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xCE:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xCF:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, []Operand{}
			} else {
				return BSWAP, []Operand{}
			}
		} else {
			return BSWAP, []Operand{}
		}
	case 0xD0:
		if isOperandSizeOveride {
			// 0x66
			return ADDSUBPD, []Operand{}
		} else if isRep0 {
			// 0xF2
			return ADDSUBPS, []Operand{}
		} else {
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		if isOperandSizeOveride {
			// 0x66
			return PSRLW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSRLW, []Operand{}
		}
	case 0xD2:
		if isOperandSizeOveride {
			// 0x66
			return PSRLD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSRLD, []Operand{}
		}
	case 0xD3:
		if isOperandSizeOveride {
			// 0x66
			return PSRLQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSRLQ, []Operand{}
		}
	case 0xD4:
		if isOperandSizeOveride {
			// 0x66
			return PADDQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDQ, []Operand{}
		}
	case 0xD5:
		if isOperandSizeOveride {
			// 0x66
			return PMULLW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMULLW, []Operand{}
		}
	case 0xD6:
		if isOperandSizeOveride {
			// 0x66
			return MOVQ, []Operand{}
		} else if isRep1 {
			// 0xF3
			return MOVQ2DQ, []Operand{}
		} else if isRep0 {
			// 0xF2
			return MOVDQ2Q, []Operand{}
		} else {
			panic("Error: Unknown instruction")
		}
	case 0xD7:
		if isOperandSizeOveride {
			// 0x66
			return PMOVMSKB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMOVMSKB, []Operand{}
		}
	case 0xD8:
		if isOperandSizeOveride {
			// 0x66
			return PSUBUSB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBUSB, []Operand{}
		}
	case 0xD9:
		if isOperandSizeOveride {
			// 0x66
			return PSUBUSW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBUSW, []Operand{}
		}
	case 0xDA:
		if isOperandSizeOveride {
			// 0x66
			return PMINUB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMINUB, []Operand{}
		}
	case 0xDB:
		if isOperandSizeOveride {
			// 0x66
			return PAND, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PAND, []Operand{}
		}
	case 0xDC:
		if isOperandSizeOveride {
			// 0x66
			return PADDUSB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDUSB, []Operand{}
		}
	case 0xDD:
		if isOperandSizeOveride {
			// 0x66
			return PADDUSW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDUSW, []Operand{}
		}
	case 0xDE:
		if isOperandSizeOveride {
			// 0x66
			return PMAXUB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMAXUB, []Operand{}
		}
	case 0xDF:
		if isOperandSizeOveride {
			// 0x66
			return PANDN, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PANDN, []Operand{}
		}
	case 0xE0:
		if isOperandSizeOveride {
			// 0x66
			return PAVGB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PAVGB, []Operand{}
		}
	case 0xE1:
		if isOperandSizeOveride {
			// 0x66
			return PSRAW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSRAW, []Operand{}
		}
	case 0xE2:
		if isOperandSizeOveride {
			// 0x66
			return PSRAD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSRAD, []Operand{}
		}
	case 0xE3:
		if isOperandSizeOveride {
			// 0x66
			return PAVGW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PAVGW, []Operand{}
		}
	case 0xE4:
		if isOperandSizeOveride {
			// 0x66
			return PMULHUW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMULHUW, []Operand{}
		}
	case 0xE5:
		if isOperandSizeOveride {
			// 0x66
			return PMULHW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMULHW, []Operand{}
		}
	case 0xE6:
		if isOperandSizeOveride {
			// 0x66
			return CVTTPD2DQ, []Operand{}
		} else if isRep1 {
			// 0xF3
			return CVTDQ2PD, []Operand{}
		} else if isRep0 {
			// 0xF2
			return CVTPD2DQ, []Operand{}
		} else {
			panic("Error: Unknown instruction")
		}
	case 0xE7:
		if isOperandSizeOveride {
			// 0x66
			return MOVNTDQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return MOVNTQ, []Operand{}
		}
	case 0xE8:
		if isOperandSizeOveride {
			// 0x66
			return PSUBSB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBSB, []Operand{}
		}
	case 0xE9:
		if isOperandSizeOveride {
			// 0x66
			return PSUBSW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBSW, []Operand{}
		}
	case 0xEA:
		if isOperandSizeOveride {
			// 0x66
			return PMINSW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMINSW, []Operand{}
		}
	case 0xEB:
		if isOperandSizeOveride {
			// 0x66
			return POR, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return POR, []Operand{}
		}
	case 0xEC:
		if isOperandSizeOveride {
			// 0x66
			return PADDSB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDSB, []Operand{}
		}
	case 0xED:
		if isOperandSizeOveride {
			// 0x66
			return PADDSW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDSW, []Operand{}
		}
	case 0xEE:
		if isOperandSizeOveride {
			// 0x66
			return PMAXSW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMAXSW, []Operand{}
		}
	case 0xEF:
		if isOperandSizeOveride {
			// 0x66
			return PXOR, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PXOR, []Operand{}
		}
	case 0xF0:
		if isRep0 || isRep1 || isOperandSizeOveride {
			panic("Error: prefix not allowed for the secondary map opcode 0xF0")
		}
		return LDDQU, []Operand{}
	case 0xF1:
		if isOperandSizeOveride {
			// 0x66
			return PSLLW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSLLW, []Operand{}
		}
	case 0xF2:
		if isOperandSizeOveride {
			// 0x66
			return PSLLD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSLLD, []Operand{}
		}
	case 0xF3:
		if isOperandSizeOveride {
			// 0x66
			return PSLLQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSLLQ, []Operand{}
		}
	case 0xF4:
		if isOperandSizeOveride {
			// 0x66
			return PMULUDQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMULUDQ, []Operand{}
		}
	case 0xF5:
		if isOperandSizeOveride {
			// 0x66
			return PMADDWD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PMADDWD, []Operand{}
		}
	case 0xF6:
		if isOperandSizeOveride {
			// 0x66
			return PSADBW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSADBW, []Operand{}
		}
	case 0xF7:
		if isOperandSizeOveride {
			// 0x66
			return MASKMOVDQU, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return MASKMOVQ, []Operand{}
		}
	case 0xF8:
		if isOperandSizeOveride {
			// 0x66
			return PSUBB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBB, []Operand{}
		}
	case 0xF9:
		if isOperandSizeOveride {
			// 0x66
			return PSUBW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBW, []Operand{}
		}
	case 0xFA:
		if isOperandSizeOveride {
			// 0x66
			return PSUBD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBD, []Operand{}
		}
	case 0xFB:
		if isOperandSizeOveride {
			// 0x66
			return PSUBQ, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBQ, []Operand{}
		}
	case 0xFC:
		if isOperandSizeOveride {
			// 0x66
			return PADDB, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDB, []Operand{}
		}
	case 0xFD:
		if isOperandSizeOveride {
			// 0x66
			return PADDW, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDW, []Operand{}
		}
	case 0xFE:
		if isOperandSizeOveride {
			// 0x66
			return PADDD, []Operand{}
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PADDD, []Operand{}
		}
	case 0xFF:
		return UD0, []Operand{}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG6(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			// return SLDT, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo")
		case [3]bool{false, false, true}:
			// return STR, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo")
		case [3]bool{false, true, false}:
			return LLDT, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, true, true}:
			return LTR, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, false}:
			return VERR, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, true}:
			return VERW, []Operand{modRMRegOperand(false, true)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG7(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return SGDT, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, false, true}:
			panic("todo")
		case [3]bool{false, true, false}:
			panic("todo")
		case [3]bool{false, true, true}:
			panic("todo")
		case [3]bool{true, false, false}:
			// return SMSW, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo")
		case [3]bool{true, false, true}:
			// return IMUL, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo")
		case [3]bool{true, true, false}:
			return LMSW, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, true, true}:
			// return IDIV, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG8(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xBA:
		switch modrmReg {
		case [3]bool{true, false, false}:
			return BT, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, false, true}:
			return BTS, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, false}:
			return BTR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, true}:
			return BTC, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG9(opcode byte, modrmReg [3]bool, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xC7:
		switch modrmReg {
		case [3]bool{false, false, true}:
			if isOperandSizeOverride {
				// 0x66
				panic("todo")
			} else if isRep1 || isRep0 {
				// 0xF3 0xF2
				panic("Error: Unknown instruction")
			} else {
				panic("todo")
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return RDRAND, []Operand{modRMRegOperand(false, true)}
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown opcode")
			} else {
				return RDRAND, []Operand{modRMRegOperand(false, true)}
			}
		case [3]bool{true, true, true}:
			if isOperandSizeOverride {
				// 0x66
				return RDSEED, []Operand{modRMRegOperand(false, true)}
			} else if isRep0 {
				// 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				panic("todo")
			} else {
				return RDSEED, []Operand{modRMRegOperand(false, true)}
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG10(opcode byte, modrmReg [3]bool, is64Bit bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xB9:
		return UD1, []Operand{}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG12(opcode byte, modrmReg [3]bool, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x71:
		switch modrmReg {
		case [3]bool{false, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRLW, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRLW, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
		case [3]bool{true, false, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRAW, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRAW, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLW, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSLLW, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG13(opcode byte, modrmReg [3]bool, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x72:
		switch modrmReg {
		case [3]bool{false, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
		case [3]bool{true, false, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRAD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRAD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSLLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG14(opcode byte, modrmReg [3]bool, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x73:
		switch modrmReg {
		case [3]bool{false, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
			if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			}
			return PSRLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
		case [3]bool{false, true, true}:
			if isOperandSizeOverride {
				// 0x66
				return PSRAD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
			if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			}
			return PSRAD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
			if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			}
			return PSLLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
		case [3]bool{true, true, true}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
			}
			if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			}
			return PSLLD, []Operand{modRMRegOperand(false, false), immediateOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG15(opcode byte, modrmReg [3]bool, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xAE:
		switch modrmReg {
		case [3]bool{false, false, false}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return RDFSBASE, []Operand{modRMRegOperand(false, true)}
			} else {
				return FXSAVE, []Operand{modRMRegOperand(true, true)}
			}
		case [3]bool{false, false, true}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return RDGSBASE, []Operand{modRMRegOperand(false, true)}
			} else {
				return FXRSTOR, []Operand{modRMRegOperand(true, true)}
			}
		case [3]bool{false, true, false}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return WRFSBASE, []Operand{modRMRegOperand(false, true)}
			} else {
				return LDMXCSR, []Operand{modRMRegOperand(true, true)}
			}
		case [3]bool{false, true, true}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return WRGSBASE, []Operand{modRMRegOperand(false, true)}
			} else {
				return STMXCSR, []Operand{modRMRegOperand(true, true)}
			}
		case [3]bool{true, false, false}:
			if isOperandSizeOverride || isRep0 || isRep1 {
				// 0x66 0xF2 0xF3
				panic("Error: Unkown opcode")
			} else {
				// return XSAVE, []Operand{modRMOperand(false, true)}
				panic("todo this needs the mod.r/m mod field")
			}
		case [3]bool{true, false, true}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return INCSSP, []Operand{}
			} else {
				panic("todo this needs the mod.r/m mod field")
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// return CLWB, false, NoSegment, NoRegister, NoRegister, 0
				panic("todo this needs the mod.r/m mod field")
			}
			if isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			}
			if isRep1 {
				return CLWB, []Operand{modRMRegOperand(true, true)}
			}
			if isRep1 {
				// 0xF3
				return CLRSSBSY, []Operand{}
			}
			panic("todo")
		case [3]bool{true, true, true}:
			if isOperandSizeOverride || isRep0 || isRep1 {
				// 0x66 0xF2 0xF3
				panic("Error: Unkown opcode")
			} else {
				panic("todo this needs the mod.r/m mod field")
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG16(opcode byte, modrmReg [3]bool, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xAE:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return PREFETCHNTA, []Operand{}
		case [3]bool{false, false, true}:
			return PREFETCHT0, []Operand{}
		case [3]bool{false, true, false}:
			return PREFETCHT1, []Operand{}
		case [3]bool{false, true, true}:
			return PREFETCHT2, []Operand{}
		case [3]bool{true, false, false}:
			return NOP, []Operand{}
		case [3]bool{true, false, true}:
			return NOP, []Operand{}
		case [3]bool{true, true, false}:
			return NOP, []Operand{}
		case [3]bool{true, true, true}:
			return NOP, []Operand{}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG17(opcode byte, modrmReg [3]bool, is64Bit bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x78:
		switch modrmReg {
		case [3]bool{false, false, false}:
			if isOperandSizeOverride {
				// return EXTRQ, []Operand{}
				panic("todo multiple immediate operands")
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
	panic("")
}

func secondaryOpcodeModRMGP(opcode byte, modrmReg [3]bool, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x0D:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return PREFETCHEXCLUSIVE, []Operand{}
		case [3]bool{false, false, true}:
			return PREFETCHMODIFIED, []Operand{}
		case [3]bool{false, true, false}:
			panic("todo this goes back to modrmReg /0")
		case [3]bool{false, true, true}:
			panic("todo this goes back to modrmReg /0")
		case [3]bool{true, false, false}:
			panic("todo this goes back to modrmReg /0")
		case [3]bool{true, false, true}:
			panic("todo this goes back to modrmReg /0")
		case [3]bool{true, true, false}:
			panic("todo this goes back to modrmReg /0")
		case [3]bool{true, true, true}:
			panic("todo this goes back to modrmReg /0")
		}
	default:
		panic("Error: Unknown instruction")
	}
	panic("")
}
