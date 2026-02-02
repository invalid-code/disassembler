package amdx8664

func secondaryOpcodeMap(curByte byte, is64Bit bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool, isRexW bool) (Instruction, bool, bool, bool, MemSegment, Register, Register, int, bool, int, int) {
	noDisplacementBytes := 4
	switch curByte {
	case 0x0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LAR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LSL, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x4:
		panic("Error: No opcode")
	case 0x5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSCALL, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return CLTS, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSRET, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return INVD, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
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
		return UD2, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
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
		return FEMMS, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		panic("Error: This is the 3dnow opcode map escape sequence")
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MOVSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MOVSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return MOVUPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x11:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MOVSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MOVSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return MOVUPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x12:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MOVSLDUP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MOVDDUP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		panic("todo multiple instructions in opcode")
	case 0x13:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return MOVLPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKLPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return UNPCKLPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKHPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return UNPCKHPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x16:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MOVSHDUP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		panic("todo multiple instructions in opcode")
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return MOVHPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1E:
		if isRep0 || isRep1 || !isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		return NoInstruction, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x1D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		panic("todo weird")
	case 0x1F:
		if isRep0 || isRep1 {
			panic("Error: prefices 0xF2 and 0xF3 not allowed for the secondary map opcode 0x1F")
		}
		return NOP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x20, 0x21, 0x22, 0x23:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x24, 0x25, 0x26, 0x27:
		panic("Error: Unkown opcode")
	case 0x28, 0x29:
		if isOperandSizeOverride {
			// 0x66
			return MOVAPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return MOVAPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2A:
		if isOperandSizeOverride {
			// 0x66
			return CVTSI2SS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return CVTPI2PD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return CVTSI2SD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		// todo
		return CVTPI2PS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2B:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MOVNTSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MOVNTSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		// todo
		return MOVNTPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2C:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2PI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return CVTTSS2SI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return CVTTSD2SI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return CVTTPS2PI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2D:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return CVTSS2SI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return CVTSD2SI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return CVTPS2PI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2E:
		if isOperandSizeOverride {
			// 0x66
			return UCOMISD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return UCOMISS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2F:
		if isOperandSizeOverride {
			// 0x66
			return COMISD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return COMISS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x30:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return WRMSR, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x31:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDTSC, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x32:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDMSR, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x33:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDPMC, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x34:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		if !is64Bit {
			panic("Error: Invalid in 32-bit")
		}
		return SYSENTER, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x35:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		if !is64Bit {
			panic("Error: Invalid in 32-bit")
		}
		return SYSEXIT, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
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
		return CMOVO, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x41:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNO, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x42:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x43:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x44:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVZ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x45:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNZ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x46:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVBE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x47:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNBE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x48:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x49:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x4A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x4B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x4C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVL, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x4D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNL, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x4E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVLE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x4F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNLE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x50:
		if isOperandSizeOverride {
			// 0x66
			return MOVMSKPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return MOVMSKPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x51:
		if isOperandSizeOverride {
			// 0x66
			return SQRTPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return SQRTSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return SQRTSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return SQRTPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x52:
		if isOperandSizeOverride || isRep0 {
			// 0x66 0xF2
			panic("Error: Unkown opcode")
		}
		if isRep1 {
			// 0xF3
			return RSQRTSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return RSQRTPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x53:
		if isOperandSizeOverride || isRep0 {
			// 0x66 0xF2
			panic("Error: Unkown opcode")
		}
		if isRep1 {
			// 0xF3
			return RCPSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return RCPPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x54:
		if isOperandSizeOverride {
			// 0x66
			return ANDPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return ANDPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x55:
		if isOperandSizeOverride {
			// 0x66
			return ANDNPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return ANDNPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x56:
		if isOperandSizeOverride {
			// 0x66
			return ORPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return ORPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x57:
		if isOperandSizeOverride {
			// 0x66
			return XORPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return XORPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x58:
		if isOperandSizeOverride {
			// 0x66
			return ADDPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return ADDSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return ADDSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return ADDPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x59:
		if isOperandSizeOverride {
			// 0x66
			return MULPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MULSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MULSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return MULPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x5A:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return CVTSS2SD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return CVTSD2SS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return CVTPS2PD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x5B:
		if isOperandSizeOverride {
			// 0x66
			return CVTPS2DQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return CVTTPS2DQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		return CVTDQ2PS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x5C:
		if isOperandSizeOverride {
			// 0x66
			return SUBPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return SUBSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return SUBSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return SUBPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x5D:
		if isOperandSizeOverride {
			// 0x66
			return MINPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MINSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MINSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return MINPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x5E:
		if isOperandSizeOverride {
			// 0x66
			return DIVPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return DIVSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return DIVSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return DIVPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x5F:
		if isOperandSizeOverride {
			// 0x66
			return MAXPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MAXSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MAXSD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return MAXPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x60:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PUNPCKLBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x61:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLWD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PUNPCKLWD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x62:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PUNPCKLDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x63:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSWB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PACKSSWB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x64:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PCMPGTB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x65:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PCMPGTW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x66:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PCMPGTD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x67:
		if isOperandSizeOverride {
			// 0x66
			return PACKUSWB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			//
			panic("Error: Unkown opcode")
		}
		return PACKUSWB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x68:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PUNPCKHBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x69:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHWD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PUNPCKHWD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x6A:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PUNPCKHDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x6B:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSDW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PACKSSDW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x6C:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLQDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		panic("Error: Unkown opcode")
	case 0x6D:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHQDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		panic("Error: Unkown opcode")
	case 0x6E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		return MOVD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x6F:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MOVDQU, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		return MOVDQA, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
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
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x74:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PCMPEQB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x75:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PCMPEQW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x76:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unkown opcode")
		}
		return PCMPEQD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x77:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x77")
		}
		return EMMS, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x78:
		if isOperandSizeOverride {
			// 0x66
			return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			// return INSERTQ, true, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo multiple operands")
		}
		if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		panic("Error: Unkown opcode")
	case 0x79:
		if isOperandSizeOverride {
			// 0x66
			return EXTRQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return INSERTQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		panic("Error: Unkown opcode")
	case 0x7A, 0x7B:
		panic("Error: Unkown opcode")
	case 0x7C:
		if isOperandSizeOverride {
			// 0x66
			return HADDPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return HADDPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		panic("Error: Unkown opcode")
	case 0x7D:
		if isOperandSizeOverride {
			// 0x66
			return HSUBPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return HSUBPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		panic("Error: Unkown opcode")
	case 0x7E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		if isRep1 {
			// 0xF3
			return MOVQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return MOVD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x7F:
		if isOperandSizeOverride {
			// 0x66
			return MOVDQA, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		if isRep1 {
			// 0xF3
			return MOVDQU, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return MOVQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x80:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JO, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x81:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNO, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x82:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JB, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x83:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNB, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x84:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JZ, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x85:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNZ, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x86:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JBE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x87:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNBE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x88:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JS, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x89:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNS, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x8A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JP, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x8B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNP, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x8C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JL, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x8D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNL, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x8E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JLE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x8F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNLE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0x90:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETO, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x91:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNO, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x92:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x93:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x94:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETZ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x95:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNZ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x96:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETBE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x97:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNBE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x98:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x99:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETL, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNL, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETLE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNLE, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xA0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH, false, false, false, FS, NoRegister, NoRegister, 0, false, 0, 0
	case 0xA1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP, false, false, false, FS, NoRegister, NoRegister, 0, false, 0, 0
	case 0xA2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return CPUID, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xA3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BT, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
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
		return PUSH, false, false, false, GS, NoRegister, NoRegister, 0, false, 0, 0
	case 0xA9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP, false, false, false, GS, NoRegister, NoRegister, 0, false, 0, 0
	case 0xAA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return RSM, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xAB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BTS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
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
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xAF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return IMUL, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LSS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BTR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB4:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LFS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LGS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xB8:
		if isRep1 {
			// 0xF3
			return POPCNT, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		panic("Error: 0xF3 prefix only allowed")
	case 0xB9:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xBA:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xBB:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			panic("Error: prefix not allowed for the secondary map opcode 0xB9-BB")
		}
		return BTC, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xBC:
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		if isRep1 {
			// 0xF3
			return TZCNT, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return BSF, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xBD:
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		if isRep1 {
			// 0xF3
			return LZCNT, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return BSR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xBE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xBF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC0:
		return XADD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC1:
		return XADD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC2:
		if isOperandSizeOverride {
			// 0x66
			return CMPPD, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return CMPPS, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return CMPSD, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return CMPPS, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: Unkown opcode")
		}
		return MOVNTI, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC4:
		if isOperandSizeOverride {
			// 0x66
			// return PINSRW, true, true, NoSegment, NoRegister, NoRegister, 0
			panic("todo multiple operands")
		}
		if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		// return PINSRW, true, true, NoSegment, NoRegister, NoRegister, 0
		panic("todo multiple operands")
	case 0xC5:
		if isOperandSizeOverride {
			// 0x66
			return PEXTRW, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		return PEXTRW, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC6:
		if isOperandSizeOverride {
			// 0x66
			return SHUFPD, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			panic("Error: Unkown opcode")
		}
		if isRep0 {
			// 0xF2
			panic("Error: Unkown opcode")
		}
		return SHUFPS, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC7:
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R8, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RAX, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, EAX, NoRegister, 0, false, 0, 0
	case 0xC9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R9, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RCX, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, ECX, NoRegister, 0, false, 0, 0
	case 0xCA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R10, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RDX, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, EDX, NoRegister, 0, false, 0, 0
	case 0xCB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R11, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RBX, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, EBX, NoRegister, 0, false, 0, 0
	case 0xCC:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R12, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RSP, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, ESP, NoRegister, 0, false, 0, 0
	case 0xCD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R13, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RBP, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, EBP, NoRegister, 0, false, 0, 0
	case 0xCE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R14, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RSI, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, ESI, NoRegister, 0, false, 0, 0
	case 0xCF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		if is64Bit {
			if isRexW {
				return BSWAP, false, false, false, NoSegment, R15, NoRegister, 0, false, 0, 0
			}
			return BSWAP, false, false, false, NoSegment, RDI, NoRegister, 0, false, 0, 0
		}
		return BSWAP, false, false, false, NoSegment, EDI, NoRegister, 0, false, 0, 0
	case 0xD0:
		if isOperandSizeOverride {
			// 0x66
			return ADDSUBPD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return ADDSUBPS, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		panic("Error: Unknown instruction")
	case 0xD1:
		if isOperandSizeOverride {
			// 0x66
			return PSRLW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSRLW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD2:
		if isOperandSizeOverride {
			// 0x66
			return PSRLD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSRLD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD3:
		if isOperandSizeOverride {
			// 0x66
			return PSRLQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSRLQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD4:
		if isOperandSizeOverride {
			// 0x66
			return PADDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD5:
		if isOperandSizeOverride {
			// 0x66
			return PMULLW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMULLW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD6:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return MOVQ2DQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return MOVDQ2Q, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		panic("Error: Unknown instruction")
	case 0xD7:
		if isOperandSizeOverride {
			// 0x66
			return PMOVMSKB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMOVMSKB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSUBUSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBUSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
	case 0xDA:
		if isOperandSizeOverride {
			// 0x66
			return PMINUB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMINUB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xDB:
		if isOperandSizeOverride {
			// 0x66
			return PAND, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PAND, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xDC:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDUSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xDD:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDUSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xDE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMAXUB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return PANDN, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PANDN, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE0:
		if isOperandSizeOverride {
			// 0x66
			return PAVGB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PAVGB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE1:
		if isOperandSizeOverride {
			// 0x66
			return PSRAW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSRAW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE2:
		if isOperandSizeOverride {
			// 0x66
			return PSRAD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSRAD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE3:
		if isOperandSizeOverride {
			// 0x66
			return PAVGW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PAVGW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE4:
		if isOperandSizeOverride {
			// 0x66
			return PMULHUW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMULHUW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE5:
		if isOperandSizeOverride {
			// 0x66
			return PMULHW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMULHW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE6:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2DQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 {
			// 0xF3
			return CVTDQ2PD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep0 {
			// 0xF2
			return CVTPD2DQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		panic("Error: Unknown instruction")
	case 0xE7:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return MOVNTQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSUBSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xE9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSUBSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xEA:
		if isOperandSizeOverride {
			// 0x66
			return PMINSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMINSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xEB:
		if isOperandSizeOverride {
			// 0x66
			return POR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return POR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xEC:
		if isOperandSizeOverride {
			// 0x66
			return PADDSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDSB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xED:
		if isOperandSizeOverride {
			// 0x66
			return PADDSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xEE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMAXSW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xEF:
		if isOperandSizeOverride {
			// 0x66
			return PXOR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PXOR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xF0")
		}
		return LDDQU, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF1:
		if isOperandSizeOverride {
			// 0x66
			return PSLLW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSLLW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF2:
		if isOperandSizeOverride {
			// 0x66
			return PSLLD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSLLD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF3:
		if isOperandSizeOverride {
			// 0x66
			return PSLLQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSLLQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF4:
		if isOperandSizeOverride {
			// 0x66
			return PMULUDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMULUDQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF5:
		if isOperandSizeOverride {
			// 0x66
			return PMADDWD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PMADDWD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF6:
		if isOperandSizeOverride {
			// 0x66
			return PSADBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSADBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF7:
		if isOperandSizeOverride {
			// 0x66
			return MASKMOVDQU, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return MASKMOVQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		} else if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		} else {
			return PSUBB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
	case 0xF9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSUBW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFA:
		if isOperandSizeOverride {
			// 0x66
			return PSUBD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSUBD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFB:
		if isOperandSizeOverride {
			// 0x66
			return PSUBQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PSUBQ, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFC:
		if isOperandSizeOverride {
			// 0x66
			return PADDB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFD:
		if isOperandSizeOverride {
			// 0x66
			return PADDW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDW, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFE:
		if isOperandSizeOverride {
			// 0x66
			return PADDD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		if isRep1 || isRep0 {
			// 0xF3 0xF2
			panic("Error: Unknown instruction")
		}
		return PADDD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFF:
		return UD0, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG6(opcode byte, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			// return SLDT, false, NoSegment, NoRegister, NoRegister, 0
			panic("todo")
		case [3]bool{false, false, true}:
			return STR, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return LLDT, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return LTR, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return VERR, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return VERW, false, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG7(opcode byte, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return SGDT, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			panic("todo")
		case [3]bool{false, true, false}:
			panic("todo")
		case [3]bool{false, true, true}:
			panic("todo")
		case [3]bool{true, false, false}:
			return MUL, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return IMUL, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return DIV, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return IDIV, false, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG8(opcode byte, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xBA:
		switch modrmReg {
		case [3]bool{true, false, false}:
			return BT, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return BTS, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return BTR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return BTC, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG9(opcode byte, modrmReg [3]bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
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
				return RDRAND, false, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown opcode")
			} else {
				return RDRAND, false, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, true, true}:
			if isOperandSizeOverride {
				// 0x66
				return RDSEED, false, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 {
				// 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				panic("todo")
			} else {
				return RDSEED, false, NoSegment, NoRegister, NoRegister, 0
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG10(opcode byte, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xB9:
		return UD1, false, NoSegment, NoRegister, NoRegister, 0
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG12(opcode byte, modrmReg [3]bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x71:
		switch modrmReg {
		case [3]bool{false, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRLW, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRLW, true, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, false, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRAW, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRAW, true, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLW, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSLLW, true, NoSegment, NoRegister, NoRegister, 0
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG13(opcode byte, modrmReg [3]bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x72:
		switch modrmReg {
		case [3]bool{false, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRLD, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRLD, true, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, false, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRAD, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRAD, true, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLD, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSLLD, true, NoSegment, NoRegister, NoRegister, 0
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG14(opcode byte, modrmReg [3]bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x73:
		switch modrmReg {
		case [3]bool{false, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSRLD, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRLD, true, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{false, true, true}:
			if isOperandSizeOverride {
				// 0x66
				return PSRAD, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSRAD, true, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLD, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSLLD, true, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, true, true}:
			if isOperandSizeOverride {
				// 0x66
				return PSLLD, true, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 || isRep1 {
				// 0xF2 0xF3
				panic("Error: Unkown instruction")
			} else {
				return PSLLD, true, NoSegment, NoRegister, NoRegister, 0
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG15(opcode byte, modrmReg [3]bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xAE:
		switch modrmReg {
		case [3]bool{false, false, false}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return RDFSBASE, false, NoSegment, NoRegister, NoRegister, 0
			} else {
				return FXSAVE, false, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{false, false, true}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return RDGSBASE, false, NoSegment, NoRegister, NoRegister, 0
			} else {
				return FXRSTOR, false, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{false, true, false}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return WRFSBASE, false, NoSegment, NoRegister, NoRegister, 0
			} else {
				return LDMXCSR, false, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{false, true, true}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return WRFSBASE, false, NoSegment, NoRegister, NoRegister, 0
			} else {
				return STMXCSR, false, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, false, false}:
			if isOperandSizeOverride || isRep0 || isRep1 {
				// 0x66 0xF2 0xF3
				panic("Error: Unkown opcode")
			} else {
				return XSAVE, false, NoSegment, NoRegister, NoRegister, 0
			}
		case [3]bool{true, false, true}:
			if isOperandSizeOverride || isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return INCSSP, false, NoSegment, NoRegister, NoRegister, 0
			} else {
				panic("todo")
			}
		case [3]bool{true, true, false}:
			if isOperandSizeOverride {
				return CLWB, false, NoSegment, NoRegister, NoRegister, 0
			} else if isRep0 {
				// 0x66 0xF2
				panic("Error: Unkown opcode")
			} else if isRep1 {
				// 0xF3
				return CLRSSBSY, false, NoSegment, NoRegister, NoRegister, 0
			} else {
				panic("todo")
			}
		case [3]bool{true, true, true}:
			if isOperandSizeOverride || isRep0 || isRep1 {
				// 0x66 0xF2 0xF3
				panic("Error: Unkown opcode")
			} else {
				panic("todo")
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG16(opcode byte, modrmReg [3]bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xAE:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return PREFETCH, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return PREFETCH, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return PREFETCH, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return NOP, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return NOP, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return NOP, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return NOP, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return NOP, false, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func secondaryOpcodeModRMG17(opcode byte, modrmReg [3]bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x78:
		switch modrmReg {
		case [3]bool{false, false, false}:
			if isOperandSizeOverride {
				return EXTRQ, true, NoSegment, NoRegister, NoRegister, 0
			}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
	panic("")
}

func secondaryOpcodeModRMGP(opcode byte, modrmReg [3]bool, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x0D:
		switch modrmReg {
		case [3]bool{false, false, false}:
			panic("todo")
		case [3]bool{false, false, true}:
			panic("todo")
		case [3]bool{false, true, false}:
			panic("todo")
		case [3]bool{false, true, true}:
			panic("todo")
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			panic("todo")
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			panic("todo")
		}
	default:
		panic("Error: Unknown instruction")
	}
	panic("")
}
