package amdx8664

func secondaryOpcodeMap32(curByte byte, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) Instruction {
	switch curByte {
	case 0x0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0x1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0x2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LAR
	case 0x3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LSL
	case 0x5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSCALL
	case 0x6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return CLTS
	case 0x7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSRET
	case 0x8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return INVD
	case 0x9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0xB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return UD2
	case 0xD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0xE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return FEMMS
	case 0xF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		// escape to 3dnow opcode map
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD
		} else if isRep1 {
			// 0xF3
			return MOVSS
		} else if isRep0 {
			// 0xF2
			return MOVSD
		} else {
			return MOVUPS
		}
	case 0x11:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD
		} else if isRep1 {
			// 0xF3
			return MOVSS
		} else if isRep0 {
			// 0xF2
			return MOVSD
		} else {
			return MOVUPS
		}
	case 0x12:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPD
		} else if isRep1 {
			// 0xF3
			return MOVSLDUP
		} else if isRep0 {
			// 0xF2
			return MOVDDUP
		} else {
			// todo
			return NOP
		}
	case 0x13:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPS
		} else {
			return MOVLPD
		}
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKLPD
		} else {
			return UNPCKLPS
		}
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKHPD
		} else {
			return UNPCKHPS
		}
	case 0x16:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD
		} else if isRep1 {
			// 0xF3
			return MOVSHDUP
		} else {
			// todo
			return NOP
		}
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD
		} else {
			return MOVHPS
		}
	case 0x18:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x19:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x20:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x21:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x22:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x23:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x28:
		if isOperandSizeOverride {
			// 0x66
			return MOVAPD
		} else {
			return MOVAPS
		}
	case 0x29:
		if isOperandSizeOverride {
			// 0x66
			return MOVAPD
		} else {
			return MOVAPS
		}
	case 0x2A:
		if isOperandSizeOverride {
			// 0x66
			return CVTSI2SS
		} else if isRep1 {
			// 0xF3
			return CVTPI2PD
		} else if isRep0 {
			// 0xF2
			return CVTSI2SD
		} else {
			// todo
			return CVTPI2PS
		}
	case 0x2B:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTPD
		} else if isRep1 {
			// 0xF3
			return MOVNTSS
		} else if isRep0 {
			// 0xF2
			return MOVNTSD
		} else {
			// todo
			return MOVNTPS
		}
	case 0x2C:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2PI
		} else if isRep1 {
			// 0xF3
			return CVTTSS2SI
		} else if isRep0 {
			// 0xF2
			return CVTTSD2SI
		} else {
			return CVTTPS2PI
		}
	case 0x2D:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PI
		} else if isRep1 {
			// 0xF3
			return CVTSS2SI
		} else if isRep0 {
			// 0xF2
			return CVTSD2SI
		} else {
			return CVTPS2PI
		}
	case 0x2E:
		if isOperandSizeOverride {
			// 0x66
			return UCOMISD
		} else {
			return UCOMISS
		}
	case 0x2F:
		if isOperandSizeOverride {
			// 0x66
			return COMISD
		} else {
			return COMISS
		}
	case 0x30:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return WRMSR
	case 0x31:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDTSC
	case 0x32:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDMSR
	case 0x33:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDPMC
	case 0x34:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return SYSENTER
	case 0x35:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return SYSEXIT
	case 0x38:
		panic("Error: This is the 0x38 escape byte")
	case 0x3A:
		panic("Error: This is the 0x3A escape byte")
	case 0x40:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVO
	case 0x41:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNO
	case 0x42:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVB
	case 0x43:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNB
	case 0x44:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVZ
	case 0x45:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNZ
	case 0x46:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVBE
	case 0x47:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNBE
	case 0x48:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVS
	case 0x49:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNS
	case 0x4A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVP
	case 0x4B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNP
	case 0x4C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVL
	case 0x4D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNL
	case 0x4E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVLE
	case 0x4F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNLE
	case 0x50:
		if isOperandSizeOverride {
			// 0x66
			return MOVMSKPD
		} else {
			return MOVMSKPS
		}
	case 0x51:
		if isOperandSizeOverride {
			// 0x66
			return SQRTPD
		} else if isRep1 {
			// 0xF3
			return SQRTSS
		} else if isRep0 {
			// 0xF2
			return SQRTSD
		} else {
			return SQRTPS
		}
	case 0x52:
		if isRep1 {
			// 0xF3
			return RSQRTSS
		} else {
			return RSQRTPS
		}
	case 0x53:
		if isRep1 {
			// 0xF3
			return RCPSS
		} else {
			return RCPPS
		}
	case 0x54:
		if isOperandSizeOverride {
			// 0x66
			return ANDPD
		} else {
			return ANDPS
		}
	case 0x55:
		if isOperandSizeOverride {
			// 0x66
			return ANDNPD
		} else {
			return ANDNPS
		}
	case 0x56:
		if isOperandSizeOverride {
			// 0x66
			return ORPD
		} else {
			return ORPS
		}
	case 0x57:
		if isOperandSizeOverride {
			// 0x66
			return XORPD
		} else {
			return XORPS
		}
	case 0x58:
		if isOperandSizeOverride {
			// 0x66
			return ADDPD
		} else if isRep1 {
			// 0xF3
			return ADDSS
		} else if isRep0 {
			// 0xF2
			return ADDSD
		} else {
			return ADDPS
		}
	case 0x59:
		if isOperandSizeOverride {
			// 0x66
			return MULPD
		} else if isRep1 {
			// 0xF3
			return MULSS
		} else if isRep0 {
			// 0xF2
			return MULSD
		} else {
			return MULPS
		}
	case 0x5A:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PS
		} else if isRep1 {
			// 0xF3
			return CVTSS2SD
		} else if isRep0 {
			// 0xF2
			return CVTSD2SS
		} else {
			return CVTPS2PD
		}
	case 0x5B:
		if isOperandSizeOverride {
			// 0x66
			return CVTPS2DQ
		} else if isRep1 {
			// 0xF3
			return CVTTPS2DQ
		} else {
			return CVTDQ2PS
		}
	case 0x5C:
		if isOperandSizeOverride {
			// 0x66
			return SUBPD
		} else if isRep1 {
			// 0xF3
			return SUBSS
		} else if isRep0 {
			// 0xF2
			return SUBSD
		} else {
			return SUBPS
		}
	case 0x5D:
		if isOperandSizeOverride {
			// 0x66
			return MINPD
		} else if isRep1 {
			// 0xF3
			return MINSS
		} else if isRep0 {
			// 0xF2
			return MINSD
		} else {
			return MINPS
		}
	case 0x5E:
		if isOperandSizeOverride {
			// 0x66
			return DIVPD
		} else if isRep1 {
			// 0xF3
			return DIVSS
		} else if isRep0 {
			// 0xF2
			return DIVSD
		} else {
			return DIVPS
		}
	case 0x5F:
		if isOperandSizeOverride {
			// 0x66
			return MAXPD
		} else if isRep1 {
			// 0xF3
			return MAXSS
		} else if isRep0 {
			// 0xF2
			return MAXSD
		} else {
			return MAXPS
		}
	case 0x60:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLBW
		} else {
			return PUNPCKLBW
		}
	case 0x61:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLWD
		} else {
			return PUNPCKLWD
		}
	case 0x62:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLDQ
		} else {
			return PUNPCKLDQ
		}
	case 0x63:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSWB
		} else {
			return PACKSSWB
		}
	case 0x64:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTB
		} else {
			return PCMPGTB
		}
	case 0x65:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTW
		} else {
			return PCMPGTW
		}
	case 0x66:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTD
		} else {
			return PCMPGTD
		}
	case 0x67:
		if isOperandSizeOverride {
			// 0x66
			return PACKUSWB
		} else {
			return PACKUSWB
		}
	case 0x68:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHBW
		} else {
			return PUNPCKHBW
		}
	case 0x69:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHWD
		} else {
			return PUNPCKHWD
		}
	case 0x6A:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHDQ
		} else {
			return PUNPCKHDQ
		}
	case 0x6B:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSDW
		} else {
			return PACKSSDW
		}
	case 0x6C:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLQDQ
		}
	case 0x6D:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHQDQ
		}
	case 0x6E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD
		} else {
			return MOVD
		}
	case 0x6F:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ
		} else if isRep1 {
			// 0xF3
			return MOVDQU
		} else {
			return MOVDQA
		}
	case 0x70:
		if isOperandSizeOverride {
			// 0x66
			return PSHUFD
		} else if isRep1 {
			// 0xF3
			return PSHUFHW
		} else if isRep0 {
			// 0xF2
			return PSHUFLW
		} else {
			return PSHUFW
		}
	case 0x71:
		// todo modrm byte
		if isOperandSizeOverride {
			// 0x66
			return NOP
		} else if isRep1 {
			// 0xF3
			return NOP
		} else if isRep0 {
			// 0xF2
			return NOP
		} else {
			return NOP
		}
	case 0x72:
		// todo modrm byte
		if isOperandSizeOverride {
			// 0x66
			return NOP
		} else if isRep1 {
			// 0xF3
			return NOP
		} else if isRep0 {
			// 0xF2
			return NOP
		} else {
			return NOP
		}
	case 0x73:
		// todo modrm byte
		if isOperandSizeOverride {
			// 0x66
			return NOP
		} else if isRep1 {
			// 0xF3
			return NOP
		} else if isRep0 {
			// 0xF2
			return NOP
		} else {
			return NOP
		}
	case 0x74:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQB
		} else {
			return PCMPEQB
		}
	case 0x75:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQW
		} else {
			return PCMPEQW
		}
	case 0x76:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQD
		} else {
			return PCMPEQD
		}
	case 0x77:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x77")
		}
		return EMMS
	case 0x78:
		if isOperandSizeOverride {
			// 0x66
			// todo modrm byte
			return NOP
		} else if isRep0 {
			// 0xF2
			return INSERTQ
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x79:
		if isOperandSizeOverride {
			// 0x66
			return EXTRQ
		} else if isRep0 {
			// 0xF2
			return INSERTQ
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x7C:
		if isOperandSizeOverride {
			// 0x66
			return HADDPD
		} else if isRep0 {
			// 0xF2
			return HADDPS
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x7D:
		if isOperandSizeOverride {
			// 0x66
			return HSUBPD
		} else if isRep0 {
			// 0xF2
			return HSUBPS
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x7E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD
		} else if isRep1 {
			// 0xF3
			return MOVQ
		} else {
			return MOVD
		}
	case 0x7F:
		if isOperandSizeOverride {
			// 0x66
			return MOVDQA
		} else if isRep1 {
			// 0xF3
			return MOVDQU
		} else {
			return MOVQ
		}
	case 0x80:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JO
	case 0x81:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNO
	case 0x82:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JB
	case 0x83:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNB
	case 0x84:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JZ
	case 0x85:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNZ
	case 0x86:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JBE
	case 0x87:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNBE
	case 0x88:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JS
	case 0x89:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNS
	case 0x8A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JP
	case 0x8B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNP
	case 0x8C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JL
	case 0x8D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNL
	case 0x8E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JLE
	case 0x8F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNLE
	case 0x90:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETO
	case 0x91:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNO
	case 0x92:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETB
	case 0x93:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNB
	case 0x94:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETZ
	case 0x95:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNZ
	case 0x96:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETBE
	case 0x97:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNBE
	case 0x98:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETS
	case 0x99:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNS
	case 0x9A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETP
	case 0x9B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNP
	case 0x9C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETL
	case 0x9D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNL
	case 0x9E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETLE
	case 0x9F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNLE
	case 0xA0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH
	case 0xA1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP
	case 0xA2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return CPUID
	case 0xA3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BT
	case 0xA4:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHLD
	case 0xA5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHLD
	case 0xA8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH
	case 0xA9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP
	case 0xAA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return RSM
	case 0xAB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BTS
	case 0xAC:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHRD
	case 0xAD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHRD
	case 0xAE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		// todo modrm
		return NOP
	case 0xAF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return IMUL
	case 0xB0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG
	case 0xB1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG
	case 0xB2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LSS
	case 0xB3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BTR
	case 0xB4:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LFS
	case 0xB5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LGS
	case 0xB6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX
	case 0xB7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX
	case 0xB8:
		if isRep1 {
			// 0xF3
			return POPCNT
		} else {
			panic("Error: 0xF3 prefix only allowed")
		}
	case 0xB9:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			// todo modrm
			return NOP
		}
		return NOP
	case 0xBA:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			// todo modrm
			return NOP
		}
		return NOP
	case 0xBB:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			return BTC
		}
		return NOP
	case 0xBC:
		if isRep1 {
			// 0xF3
			return TZCNT
		} else {
			return BSF
		}
	case 0xBD:
		if isRep1 {
			// 0xF3
			return LZCNT
		} else {
			return BSR
		}
	case 0xBE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX
	case 0xBF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX
	case 0xC0:
		return XADD
	case 0xC1:
		return XADD
	case 0xC2:
		if isOperandSizeOverride {
			// 0x66
			return CMPPD
		} else if isRep1 {
			// 0xF3
			return CMPPS
		} else if isRep0 {
			// 0xF2
			return CMPSD
		} else {
			return CMPPS
		}
	case 0xC3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			return NOP
		}
		return MOVNTI
	case 0xC4:
		if isOperandSizeOverride {
			// 0x66
			return PINSRW
		} else {
			return PINSRW
		}
	case 0xC5:
		if isOperandSizeOverride {
			// 0x66
			return PEXTRW
		} else {
			return PEXTRW
		}
	case 0xC6:
		if isOperandSizeOverride {
			// 0x66
			return SHUFPD
		} else {
			return SHUFPS
		}
	case 0xC7:
		// todo modrm byte
		return NOP
	case 0xC8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xC9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCC:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xD0:
		if isOperandSizeOverride {
			// 0x66
			return ADDSUBPD
		} else if isRep0 {
			// 0xF2
			return ADDSUBPS
		} else {
			return NOP
		}
	case 0xD1:
		if isOperandSizeOverride {
			// 0x66
			return PSRLW
		} else {
			return PSRLW
		}
	case 0xD2:
		if isOperandSizeOverride {
			// 0x66
			return PSRLD
		} else {
			return PSRLD
		}
	case 0xD3:
		if isOperandSizeOverride {
			// 0x66
			return PSRLQ
		} else {
			return PSRLQ
		}
	case 0xD4:
		if isOperandSizeOverride {
			// 0x66
			return PADDQ
		} else {
			return PADDQ
		}
	case 0xD5:
		if isOperandSizeOverride {
			// 0x66
			return PMULLW
		} else {
			return PMULLW
		}
	case 0xD6:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ
		} else if isRep1 {
			// 0xF3
			return MOVQ2DQ
		} else if isRep0 {
			// 0xF2
			return MOVDQ2Q
		} else {
			return NOP
		}
	case 0xD7:
		if isOperandSizeOverride {
			// 0x66
			return PMOVMSKB
		} else {
			return PMOVMSKB
		}
	case 0xD8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSB
		} else {
			return PSUBUSB
		}
	case 0xD9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSW
		} else {
			return PSUBUSW
		}
	case 0xDA:
		if isOperandSizeOverride {
			// 0x66
			return PMINUB
		} else {
			return PMINUB
		}
	case 0xDB:
		if isOperandSizeOverride {
			// 0x66
			return PAND
		} else {
			return PAND
		}
	case 0xDC:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSB
		} else {
			return PADDUSB
		}
	case 0xDD:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSW
		} else {
			return PADDUSW
		}
	case 0xDE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUB
		} else {
			return PMAXUB
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return PANDN
		} else {
			return PANDN
		}
	case 0xE0:
		if isOperandSizeOverride {
			// 0x66
			return PAVGB
		} else {
			return PAVGB
		}
	case 0xE1:
		if isOperandSizeOverride {
			// 0x66
			return PSRAW
		} else {
			return PSRAW
		}
	case 0xE2:
		if isOperandSizeOverride {
			// 0x66
			return PSRAD
		} else {
			return PSRAD
		}
	case 0xE3:
		if isOperandSizeOverride {
			// 0x66
			return PAVGW
		} else {
			return PAVGW
		}
	case 0xE4:
		if isOperandSizeOverride {
			// 0x66
			return PMULHUW
		} else {
			return PMULHUW
		}
	case 0xE5:
		if isOperandSizeOverride {
			// 0x66
			return PMULHW
		} else {
			return PMULHW
		}
	case 0xE6:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2DQ
		} else if isRep1 {
			// 0xF3
			return CVTDQ2PD
		} else if isRep0 {
			// 0xF2
			return CVTPD2DQ
		} else {
			return NOP
		}
	case 0xE7:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTDQ
		} else {
			return MOVNTQ
		}
	case 0xE8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSB
		} else {
			return PSUBSB
		}
	case 0xE9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSW
		} else {
			return PSUBSW
		}
	case 0xEA:
		if isOperandSizeOverride {
			// 0x66
			return PMINSW
		} else {
			return PMINSW
		}
	case 0xEB:
		if isOperandSizeOverride {
			// 0x66
			return POR
		} else {
			return POR
		}
	case 0xEC:
		if isOperandSizeOverride {
			// 0x66
			return PADDSB
		} else {
			return PADDSB
		}
	case 0xED:
		if isOperandSizeOverride {
			// 0x66
			return PADDSW
		} else {
			return PADDSW
		}
	case 0xEE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSW
		} else {
			return PMAXSW
		}
	case 0xEF:
		if isOperandSizeOverride {
			// 0x66
			return PXOR
		} else {
			return PXOR
		}
	case 0xF0:
		if isRep0 {
			// 0xF2
			return LDDQU
		} else {
			return NOP
		}
	case 0xF1:
		if isOperandSizeOverride {
			// 0x66
			return PSLLW
		} else {
			return PSLLW
		}
	case 0xF2:
		if isOperandSizeOverride {
			// 0x66
			return PSLLD
		} else {
			return PSLLD
		}
	case 0xF3:
		if isOperandSizeOverride {
			// 0x66
			return PSLLQ
		} else {
			return PSLLQ
		}
	case 0xF4:
		if isOperandSizeOverride {
			// 0x66
			return PMULUDQ
		} else {
			return PMULUDQ
		}
	case 0xF5:
		if isOperandSizeOverride {
			// 0x66
			return PMADDWD
		} else {
			return PMADDWD
		}
	case 0xF6:
		if isOperandSizeOverride {
			// 0x66
			return PSADBW
		} else {
			return PSADBW
		}
	case 0xF7:
		if isOperandSizeOverride {
			// 0x66
			return MASKMOVDQU
		} else {
			return MASKMOVQ
		}
	case 0xF8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBB
		} else {
			return PSUBB
		}
	case 0xF9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBW
		} else {
			return PSUBW
		}
	case 0xFA:
		if isOperandSizeOverride {
			// 0x66
			return PSUBD
		} else {
			return PSUBD
		}
	case 0xFB:
		if isOperandSizeOverride {
			// 0x66
			return PSUBQ
		} else {
			return PSUBQ
		}
	case 0xFC:
		if isOperandSizeOverride {
			// 0x66
			return PADDB
		} else {
			return PADDB
		}
	case 0xFD:
		if isOperandSizeOverride {
			// 0x66
			return PADDW
		} else {
			return PADDW
		}
	case 0xFE:
		if isOperandSizeOverride {
			// 0x66
			return PADDD
		} else {
			return PADDD
		}
	case 0xFF:
		return UD0
	}
	panic("Error: todo")
}

func secondaryOpcodeMap64(curByte byte, isRep0 bool, isRep1 bool, isOperandSizeOverride bool) Instruction {
	switch curByte {
	case 0x0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0x1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0x2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LAR
	case 0x3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return LSL
	case 0x5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSCALL
	case 0x6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return CLTS
	case 0x7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return SYSRET
	case 0x8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return INVD
	case 0x9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0xB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return UD2
	case 0xD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
	case 0xE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		return FEMMS
	case 0xF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x0 - 0xF")
		}
		// escape to 3dnow opcode map
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD
		} else if isRep1 {
			// 0xF3
			return MOVSS
		} else if isRep0 {
			// 0xF2
			return MOVSD
		} else {
			return MOVUPS
		}
	case 0x11:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD
		} else if isRep1 {
			// 0xF3
			return MOVSS
		} else if isRep0 {
			// 0xF2
			return MOVSD
		} else {
			return MOVUPS
		}
	case 0x12:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPD
		} else if isRep1 {
			// 0xF3
			return MOVSLDUP
		} else if isRep0 {
			// 0xF2
			return MOVDDUP
		} else {
			// todo
			return NOP
		}
	case 0x13:
		if isOperandSizeOverride {
			// 0x66
			return MOVLPS
		} else {
			return MOVLPD
		}
	case 0x14:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKLPD
		} else {
			return UNPCKLPS
		}
	case 0x15:
		if isOperandSizeOverride {
			// 0x66
			return UNPCKHPD
		} else {
			return UNPCKHPS
		}
	case 0x16:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD
		} else if isRep1 {
			// 0xF3
			return MOVSHDUP
		} else {
			// todo
			return NOP
		}
	case 0x17:
		if isOperandSizeOverride {
			// 0x66
			return MOVHPD
		} else {
			return MOVHPS
		}
	case 0x18:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x19:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x1F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x18 - 0x1F")
		}
		// todo
		return NOP
	case 0x20:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x21:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x22:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x23:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x20 - 0x27")
		}
		return MOV
	case 0x28:
		if isOperandSizeOverride {
			// 0x66
			return MOVAPD
		} else {
			return MOVAPS
		}
	case 0x29:
		if isOperandSizeOverride {
			// 0x66
			return MOVAPD
		} else {
			return MOVAPS
		}
	case 0x2A:
		if isOperandSizeOverride {
			// 0x66
			return CVTSI2SS
		} else if isRep1 {
			// 0xF3
			return CVTPI2PD
		} else if isRep0 {
			// 0xF2
			return CVTSI2SD
		} else {
			// todo
			return CVTPI2PS
		}
	case 0x2B:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTPD
		} else if isRep1 {
			// 0xF3
			return MOVNTSS
		} else if isRep0 {
			// 0xF2
			return MOVNTSD
		} else {
			// todo
			return MOVNTPS
		}
	case 0x2C:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2PI
		} else if isRep1 {
			// 0xF3
			return CVTTSS2SI
		} else if isRep0 {
			// 0xF2
			return CVTTSD2SI
		} else {
			return CVTTPS2PI
		}
	case 0x2D:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PI
		} else if isRep1 {
			// 0xF3
			return CVTSS2SI
		} else if isRep0 {
			// 0xF2
			return CVTSD2SI
		} else {
			return CVTPS2PI
		}
	case 0x2E:
		if isOperandSizeOverride {
			// 0x66
			return UCOMISD
		} else {
			return UCOMISS
		}
	case 0x2F:
		if isOperandSizeOverride {
			// 0x66
			return COMISD
		} else {
			return COMISS
		}
	case 0x30:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return WRMSR
	case 0x31:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDTSC
	case 0x32:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDMSR
	case 0x33:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return RDPMC
	case 0x34:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return SYSENTER
	case 0x35:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x30 - 0x37")
		}
		return SYSEXIT
	case 0x38:
		panic("Error: This is the 0x38 escape byte")
	case 0x3A:
		panic("Error: This is the 0x3A escape byte")
	case 0x40:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVO
	case 0x41:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNO
	case 0x42:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVB
	case 0x43:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNB
	case 0x44:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVZ
	case 0x45:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNZ
	case 0x46:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVBE
	case 0x47:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNBE
	case 0x48:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVS
	case 0x49:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNS
	case 0x4A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVP
	case 0x4B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNP
	case 0x4C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVL
	case 0x4D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNL
	case 0x4E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVLE
	case 0x4F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x40 - 0x4F")
		}
		return CMOVNLE
	case 0x50:
		if isOperandSizeOverride {
			// 0x66
			return MOVMSKPD
		} else {
			return MOVMSKPS
		}
	case 0x51:
		if isOperandSizeOverride {
			// 0x66
			return SQRTPD
		} else if isRep1 {
			// 0xF3
			return SQRTSS
		} else if isRep0 {
			// 0xF2
			return SQRTSD
		} else {
			return SQRTPS
		}
	case 0x52:
		if isRep1 {
			// 0xF3
			return RSQRTSS
		} else {
			return RSQRTPS
		}
	case 0x53:
		if isRep1 {
			// 0xF3
			return RCPSS
		} else {
			return RCPPS
		}
	case 0x54:
		if isOperandSizeOverride {
			// 0x66
			return ANDPD
		} else {
			return ANDPS
		}
	case 0x55:
		if isOperandSizeOverride {
			// 0x66
			return ANDNPD
		} else {
			return ANDNPS
		}
	case 0x56:
		if isOperandSizeOverride {
			// 0x66
			return ORPD
		} else {
			return ORPS
		}
	case 0x57:
		if isOperandSizeOverride {
			// 0x66
			return XORPD
		} else {
			return XORPS
		}
	case 0x58:
		if isOperandSizeOverride {
			// 0x66
			return ADDPD
		} else if isRep1 {
			// 0xF3
			return ADDSS
		} else if isRep0 {
			// 0xF2
			return ADDSD
		} else {
			return ADDPS
		}
	case 0x59:
		if isOperandSizeOverride {
			// 0x66
			return MULPD
		} else if isRep1 {
			// 0xF3
			return MULSS
		} else if isRep0 {
			// 0xF2
			return MULSD
		} else {
			return MULPS
		}
	case 0x5A:
		if isOperandSizeOverride {
			// 0x66
			return CVTPD2PS
		} else if isRep1 {
			// 0xF3
			return CVTSS2SD
		} else if isRep0 {
			// 0xF2
			return CVTSD2SS
		} else {
			return CVTPS2PD
		}
	case 0x5B:
		if isOperandSizeOverride {
			// 0x66
			return CVTPS2DQ
		} else if isRep1 {
			// 0xF3
			return CVTTPS2DQ
		} else {
			return CVTDQ2PS
		}
	case 0x5C:
		if isOperandSizeOverride {
			// 0x66
			return SUBPD
		} else if isRep1 {
			// 0xF3
			return SUBSS
		} else if isRep0 {
			// 0xF2
			return SUBSD
		} else {
			return SUBPS
		}
	case 0x5D:
		if isOperandSizeOverride {
			// 0x66
			return MINPD
		} else if isRep1 {
			// 0xF3
			return MINSS
		} else if isRep0 {
			// 0xF2
			return MINSD
		} else {
			return MINPS
		}
	case 0x5E:
		if isOperandSizeOverride {
			// 0x66
			return DIVPD
		} else if isRep1 {
			// 0xF3
			return DIVSS
		} else if isRep0 {
			// 0xF2
			return DIVSD
		} else {
			return DIVPS
		}
	case 0x5F:
		if isOperandSizeOverride {
			// 0x66
			return MAXPD
		} else if isRep1 {
			// 0xF3
			return MAXSS
		} else if isRep0 {
			// 0xF2
			return MAXSD
		} else {
			return MAXPS
		}
	case 0x60:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLBW
		} else {
			return PUNPCKLBW
		}
	case 0x61:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLWD
		} else {
			return PUNPCKLWD
		}
	case 0x62:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLDQ
		} else {
			return PUNPCKLDQ
		}
	case 0x63:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSWB
		} else {
			return PACKSSWB
		}
	case 0x64:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTB
		} else {
			return PCMPGTB
		}
	case 0x65:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTW
		} else {
			return PCMPGTW
		}
	case 0x66:
		if isOperandSizeOverride {
			// 0x66
			return PCMPGTD
		} else {
			return PCMPGTD
		}
	case 0x67:
		if isOperandSizeOverride {
			// 0x66
			return PACKUSWB
		} else {
			return PACKUSWB
		}
	case 0x68:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHBW
		} else {
			return PUNPCKHBW
		}
	case 0x69:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHWD
		} else {
			return PUNPCKHWD
		}
	case 0x6A:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHDQ
		} else {
			return PUNPCKHDQ
		}
	case 0x6B:
		if isOperandSizeOverride {
			// 0x66
			return PACKSSDW
		} else {
			return PACKSSDW
		}
	case 0x6C:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKLQDQ
		}
	case 0x6D:
		if isOperandSizeOverride {
			// 0x66
			return PUNPCKHQDQ
		}
	case 0x6E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD
		} else {
			return MOVD
		}
	case 0x6F:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ
		} else if isRep1 {
			// 0xF3
			return MOVDQU
		} else {
			return MOVDQA
		}
	case 0x70:
		if isOperandSizeOverride {
			// 0x66
			return PSHUFD
		} else if isRep1 {
			// 0xF3
			return PSHUFHW
		} else if isRep0 {
			// 0xF2
			return PSHUFLW
		} else {
			return PSHUFW
		}
	case 0x71:
		// todo modrm byte
		if isOperandSizeOverride {
			// 0x66
			return NOP
		} else if isRep1 {
			// 0xF3
			return NOP
		} else if isRep0 {
			// 0xF2
			return NOP
		} else {
			return NOP
		}
	case 0x72:
		// todo modrm byte
		if isOperandSizeOverride {
			// 0x66
			return NOP
		} else if isRep1 {
			// 0xF3
			return NOP
		} else if isRep0 {
			// 0xF2
			return NOP
		} else {
			return NOP
		}
	case 0x73:
		// todo modrm byte
		if isOperandSizeOverride {
			// 0x66
			return NOP
		} else if isRep1 {
			// 0xF3
			return NOP
		} else if isRep0 {
			// 0xF2
			return NOP
		} else {
			return NOP
		}
	case 0x74:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQB
		} else {
			return PCMPEQB
		}
	case 0x75:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQW
		} else {
			return PCMPEQW
		}
	case 0x76:
		if isOperandSizeOverride {
			// 0x66
			return PCMPEQD
		} else {
			return PCMPEQD
		}
	case 0x77:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x77")
		}
		return EMMS
	case 0x78:
		if isOperandSizeOverride {
			// 0x66
			// todo modrm byte
			return NOP
		} else if isRep0 {
			// 0xF2
			return INSERTQ
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x79:
		if isOperandSizeOverride {
			// 0x66
			return EXTRQ
		} else if isRep0 {
			// 0xF2
			return INSERTQ
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x7C:
		if isOperandSizeOverride {
			// 0x66
			return HADDPD
		} else if isRep0 {
			// 0xF2
			return HADDPS
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x7D:
		if isOperandSizeOverride {
			// 0x66
			return HSUBPD
		} else if isRep0 {
			// 0xF2
			return HSUBPS
		} else {
			panic("Error: No prefix was chosen")
		}
	case 0x7E:
		if isOperandSizeOverride {
			// 0x66
			return MOVD
		} else if isRep1 {
			// 0xF3
			return MOVQ
		} else {
			return MOVD
		}
	case 0x7F:
		if isOperandSizeOverride {
			// 0x66
			return MOVDQA
		} else if isRep1 {
			// 0xF3
			return MOVDQU
		} else {
			return MOVQ
		}
	case 0x80:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JO
	case 0x81:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNO
	case 0x82:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JB
	case 0x83:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNB
	case 0x84:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JZ
	case 0x85:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNZ
	case 0x86:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JBE
	case 0x87:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNBE
	case 0x88:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JS
	case 0x89:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNS
	case 0x8A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JP
	case 0x8B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNP
	case 0x8C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JL
	case 0x8D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNL
	case 0x8E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JLE
	case 0x8F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x80-8F")
		}
		return JNLE
	case 0x90:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETO
	case 0x91:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNO
	case 0x92:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETB
	case 0x93:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNB
	case 0x94:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETZ
	case 0x95:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNZ
	case 0x96:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETBE
	case 0x97:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNBE
	case 0x98:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETS
	case 0x99:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNS
	case 0x9A:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETP
	case 0x9B:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNP
	case 0x9C:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETL
	case 0x9D:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNL
	case 0x9E:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETLE
	case 0x9F:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0x90-9F")
		}
		return SETNLE
	case 0xA0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH
	case 0xA1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP
	case 0xA2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return CPUID
	case 0xA3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BT
	case 0xA4:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHLD
	case 0xA5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHLD
	case 0xA8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return PUSH
	case 0xA9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return POP
	case 0xAA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return RSM
	case 0xAB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return BTS
	case 0xAC:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHRD
	case 0xAD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return SHRD
	case 0xAE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		// todo modrm
		return NOP
	case 0xAF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xA0-A5..A8-AF")
		}
		return IMUL
	case 0xB0:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG
	case 0xB1:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return CMPXCHG
	case 0xB2:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LSS
	case 0xB3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BTR
	case 0xB4:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LFS
	case 0xB5:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return LGS
	case 0xB6:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX
	case 0xB7:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVZX
	case 0xB8:
		if isRep1 {
			// 0xF3
			return POPCNT
		} else {
			panic("Error: 0xF3 prefix only allowed")
		}
	case 0xB9:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			// todo modrm
			return NOP
		}
		return NOP
	case 0xBA:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			// todo modrm
			return NOP
		}
		return NOP
	case 0xBB:
		if !(isRep0 || isRep1 || isOperandSizeOverride) {
			return BTC
		}
		return NOP
	case 0xBC:
		if isRep1 {
			// 0xF3
			return TZCNT
		} else {
			return BSF
		}
	case 0xBD:
		if isRep1 {
			// 0xF3
			return LZCNT
		} else {
			return BSR
		}
	case 0xBE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX
	case 0xBF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return MOVSX
	case 0xC0:
		return XADD
	case 0xC1:
		return XADD
	case 0xC2:
		if isOperandSizeOverride {
			// 0x66
			return CMPPD
		} else if isRep1 {
			// 0xF3
			return CMPPS
		} else if isRep0 {
			// 0xF2
			return CMPSD
		} else {
			return CMPPS
		}
	case 0xC3:
		if isRep0 || isRep1 || isOperandSizeOverride {
			return NOP
		}
		return MOVNTI
	case 0xC4:
		if isOperandSizeOverride {
			// 0x66
			return PINSRW
		} else {
			return PINSRW
		}
	case 0xC5:
		if isOperandSizeOverride {
			// 0x66
			return PEXTRW
		} else {
			return PEXTRW
		}
	case 0xC6:
		if isOperandSizeOverride {
			// 0x66
			return SHUFPD
		} else {
			return SHUFPS
		}
	case 0xC7:
		// todo modrm byte
		return NOP
	case 0xC8:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xC9:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCA:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCB:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCC:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCD:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCE:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xCF:
		if isRep0 || isRep1 || isOperandSizeOverride {
			panic("Error: prefix not allowed for the secondary map opcode 0xB0-B7")
		}
		return BSWAP
	case 0xD0:
		if isOperandSizeOverride {
			// 0x66
			return ADDSUBPD
		} else if isRep0 {
			// 0xF2
			return ADDSUBPS
		} else {
			return NOP
		}
	case 0xD1:
		if isOperandSizeOverride {
			// 0x66
			return PSRLW
		} else {
			return PSRLW
		}
	case 0xD2:
		if isOperandSizeOverride {
			// 0x66
			return PSRLD
		} else {
			return PSRLD
		}
	case 0xD3:
		if isOperandSizeOverride {
			// 0x66
			return PSRLQ
		} else {
			return PSRLQ
		}
	case 0xD4:
		if isOperandSizeOverride {
			// 0x66
			return PADDQ
		} else {
			return PADDQ
		}
	case 0xD5:
		if isOperandSizeOverride {
			// 0x66
			return PMULLW
		} else {
			return PMULLW
		}
	case 0xD6:
		if isOperandSizeOverride {
			// 0x66
			return MOVQ
		} else if isRep1 {
			// 0xF3
			return MOVQ2DQ
		} else if isRep0 {
			// 0xF2
			return MOVDQ2Q
		} else {
			return NOP
		}
	case 0xD7:
		if isOperandSizeOverride {
			// 0x66
			return PMOVMSKB
		} else {
			return PMOVMSKB
		}
	case 0xD8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSB
		} else {
			return PSUBUSB
		}
	case 0xD9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBUSW
		} else {
			return PSUBUSW
		}
	case 0xDA:
		if isOperandSizeOverride {
			// 0x66
			return PMINUB
		} else {
			return PMINUB
		}
	case 0xDB:
		if isOperandSizeOverride {
			// 0x66
			return PAND
		} else {
			return PAND
		}
	case 0xDC:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSB
		} else {
			return PADDUSB
		}
	case 0xDD:
		if isOperandSizeOverride {
			// 0x66
			return PADDUSW
		} else {
			return PADDUSW
		}
	case 0xDE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXUB
		} else {
			return PMAXUB
		}
	case 0xDF:
		if isOperandSizeOverride {
			// 0x66
			return PANDN
		} else {
			return PANDN
		}
	case 0xE0:
		if isOperandSizeOverride {
			// 0x66
			return PAVGB
		} else {
			return PAVGB
		}
	case 0xE1:
		if isOperandSizeOverride {
			// 0x66
			return PSRAW
		} else {
			return PSRAW
		}
	case 0xE2:
		if isOperandSizeOverride {
			// 0x66
			return PSRAD
		} else {
			return PSRAD
		}
	case 0xE3:
		if isOperandSizeOverride {
			// 0x66
			return PAVGW
		} else {
			return PAVGW
		}
	case 0xE4:
		if isOperandSizeOverride {
			// 0x66
			return PMULHUW
		} else {
			return PMULHUW
		}
	case 0xE5:
		if isOperandSizeOverride {
			// 0x66
			return PMULHW
		} else {
			return PMULHW
		}
	case 0xE6:
		if isOperandSizeOverride {
			// 0x66
			return CVTTPD2DQ
		} else if isRep1 {
			// 0xF3
			return CVTDQ2PD
		} else if isRep0 {
			// 0xF2
			return CVTPD2DQ
		} else {
			return NOP
		}
	case 0xE7:
		if isOperandSizeOverride {
			// 0x66
			return MOVNTDQ
		} else {
			return MOVNTQ
		}
	case 0xE8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSB
		} else {
			return PSUBSB
		}
	case 0xE9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBSW
		} else {
			return PSUBSW
		}
	case 0xEA:
		if isOperandSizeOverride {
			// 0x66
			return PMINSW
		} else {
			return PMINSW
		}
	case 0xEB:
		if isOperandSizeOverride {
			// 0x66
			return POR
		} else {
			return POR
		}
	case 0xEC:
		if isOperandSizeOverride {
			// 0x66
			return PADDSB
		} else {
			return PADDSB
		}
	case 0xED:
		if isOperandSizeOverride {
			// 0x66
			return PADDSW
		} else {
			return PADDSW
		}
	case 0xEE:
		if isOperandSizeOverride {
			// 0x66
			return PMAXSW
		} else {
			return PMAXSW
		}
	case 0xEF:
		if isOperandSizeOverride {
			// 0x66
			return PXOR
		} else {
			return PXOR
		}
	case 0xF0:
		if isRep0 {
			// 0xF2
			return LDDQU
		} else {
			return NOP
		}
	case 0xF1:
		if isOperandSizeOverride {
			// 0x66
			return PSLLW
		} else {
			return PSLLW
		}
	case 0xF2:
		if isOperandSizeOverride {
			// 0x66
			return PSLLD
		} else {
			return PSLLD
		}
	case 0xF3:
		if isOperandSizeOverride {
			// 0x66
			return PSLLQ
		} else {
			return PSLLQ
		}
	case 0xF4:
		if isOperandSizeOverride {
			// 0x66
			return PMULUDQ
		} else {
			return PMULUDQ
		}
	case 0xF5:
		if isOperandSizeOverride {
			// 0x66
			return PMADDWD
		} else {
			return PMADDWD
		}
	case 0xF6:
		if isOperandSizeOverride {
			// 0x66
			return PSADBW
		} else {
			return PSADBW
		}
	case 0xF7:
		if isOperandSizeOverride {
			// 0x66
			return MASKMOVDQU
		} else {
			return MASKMOVQ
		}
	case 0xF8:
		if isOperandSizeOverride {
			// 0x66
			return PSUBB
		} else {
			return PSUBB
		}
	case 0xF9:
		if isOperandSizeOverride {
			// 0x66
			return PSUBW
		} else {
			return PSUBW
		}
	case 0xFA:
		if isOperandSizeOverride {
			// 0x66
			return PSUBD
		} else {
			return PSUBD
		}
	case 0xFB:
		if isOperandSizeOverride {
			// 0x66
			return PSUBQ
		} else {
			return PSUBQ
		}
	case 0xFC:
		if isOperandSizeOverride {
			// 0x66
			return PADDB
		} else {
			return PADDB
		}
	case 0xFD:
		if isOperandSizeOverride {
			// 0x66
			return PADDW
		} else {
			return PADDW
		}
	case 0xFE:
		if isOperandSizeOverride {
			// 0x66
			return PADDD
		} else {
			return PADDD
		}
	case 0xFF:
		return UD0
	}
	panic("Error: todo")
}
