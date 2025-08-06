package amdx8664

func primaryOpcode(curByte byte, is64Bit bool) (Instruction, bool, bool, MemSegment, Register) {
	switch curByte {
	case 0x0, 0x1, 0x2, 0x3:
		return ADD, true, false, NoSegment, NoRegister
	case 0x4, 0x14:
		panic("todo dont know what AL opcode syntax notation")
	case 0x5:
		return ADD, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x6:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, ES, NoRegister
		}
	case 0x7:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return POP, false, false, ES, NoRegister
		}
	case 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1F:
		return POP, false, false, ES, 
	case 0x8, 0x9, 0xA, 0xB, 0xC, 0xD:
		return OR, true, false, NoSegment
	case 0xE, 0x1E, 0x27, 0x37, 0x62, 0x9A, 0x9E, 0x9F, 0xC4, 0xC5, 0xCE:
		return PUSH, false, false, CS
	case 0xF:
		panic("Error: 2-byte opcodes")
	case 0x10, 0x11, 0x12, 0x13:
		return ADC, true, false, NoSegment
	case 0x15:
		return ADC, false, true, NoSegment
	case 0x16:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, SS
		}
	case 0x17:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return POP, false, false, SS
		}
	case 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F, 0x68, 0x6A:
		return PUSH, true
	case 0x20, 0x21, 0x22, 0x23, 0x24, 0x25:
		return AND
	case 0x26:
		panic("Error: this is the ES segment")
	case 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D:
		return SUB
	case 0x2E:
		panic("Error: this is the CS segment")
	case 0x2F:
		return DAS
	case 0x30, 0x31, 0x32, 0x33, 0x34, 0x35:
		return XOR
	case 0x36:
		panic("Error: this is the SS segment")
	case 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D:
		return CMP
	case 0x3E:
		panic("Error: this is the DS segment")
	case 0x3F:
		return AAS
	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x4E, 0x4F:
		panic("Error: Used as REX prefix in 64-bit mode")
	case 0x60:
		return PUSHA
	case 0x61:
		return POPA
	case 0x63:
		return PUSH
	case 0x64:
		panic("Error: this is the FS segment")
	case 0x65:
		panic("Error: this is the GS segment")
	case 0x66:
		panic("Error: this is the operand size override prefix")
	case 0x67:
		panic("Error: this is the address size override prefix")
	case 0x69, 0x6B:
		return IMUL
	case 0x6C:
		return INSB
	case 0x6D:
		panic("todo multiple instructions in 1 byte")
	case 0x6E:
		return OUTSB
	case 0x6F:
		panic("todo multiple instructions in 1 byte")
	case 0x70:
		return JO
	case 0x71:
		return JNO
	case 0x72:
		return JB
	case 0x73:
		return JNB
	case 0x74:
		return JZ
	case 0x75:
		return JNZ
	case 0x76:
		return JBE
	case 0x77:
		return JNBE
	case 0x78:
		return JS
	case 0x79:
		return JNS
	case 0x7A:
		return JP
	case 0x7B:
		return JNP
	case 0x7C:
		return JL
	case 0x7D:
		return JNL
	case 0x7E:
		return JLE
	case 0x7F:
		return JNLE
	case 0x80:
		panic("todo modr/m opcode extension")
	case 0x81:
		panic("todo modr/m opcode extension")
	case 0x82:
		panic("todo modr/m opcode extension")
	case 0x83:
		panic("todo modr/m opcode extension")
	case 0x84, 0x85, 0xA8, 0xA9:
		return TEST
	case 0x86, 0x87, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97:
		return XCHG
	case 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8E, 0xA0, 0xA1, 0xA2, 0xA3, 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF:
		return MOV
	case 0x8D:
		return LEA
	case 0x8F:
		panic("todo modr/m opcode extension")
	case 0x90:
		panic("todo could be XCHG or NOP instruction")
	case 0x98:
		panic("todo multiple instructions in 1 byte")
	case 0x99:
		panic("todo multiple instructions in 1 byte")
	case 0x9B:
		panic("todo multiple instructions in 1 byte")
	case 0x9C:
		panic("todo multiple instructions in 1 byte")
	case 0x9D:
		panic("todo multiple instructions in 1 byte")
	case 0xA4:
		return MOVSB
	case 0xA5:
		panic("todo multiple instructions in 1 byte")
	case 0xA6:
		return CMPSB
	case 0xA7:
		panic("todo multiple instructions in 1 byte")
	case 0xAA:
		return STOSB
	case 0xAB:
		panic("todo multiple instructions in 1 byte")
	case 0xAC:
		return LODSB
	case 0xAD:
		panic("todo multiple instructions in 1 byte")
	case 0xAE:
		return SCASB
	case 0xAF:
		panic("todo multiple instructions in 1 byte")
	case 0xC0, 0xC1:
		panic("todo modr/m opcode extension")
	case 0xC2, 0xC3:
		return RET // near
	case 0xC6:
		panic("todo modr/m opcode extension")
	case 0xC7:
		panic("todo modr/m opcode extension")
	case 0xC8:
		return ENTER
	case 0xC9:
		return LEAVE
	case 0xCA:
		return RET // far
	case 0xCB:
		return RET // far
	case 0xCC:
		return INT3
	case 0xCD:
		return INT
	case 0xCF:
		panic("todo multiple instructions in 1 byte")
	case 0xD0, 0xD1, 0xD2, 0xD3:
		panic("todo modr/m opcode extension")
	case 0xD4:
		return AAM
	case 0xD5:
		return AAD
	case 0xD6:
		return SALC
	case 0xD7:
		return XLAT
	case 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF:
		panic("Error: this is x87 opcodes")
	case 0xE0:
		panic("todo: ? LOOPNE/NZ prob opperand size diff")
	case 0xE1:
		panic("todo: ? LOOPE/Z prob opperand size diff")
	case 0xE2:
		return LOOP
	case 0xE3:
		return JrCXZ
	case 0xE4, 0xE5:
		return IN
	case 0xE6, 0xE7, 0xEE, 0xEF:
		return OUT
	case 0xE8:
		return CALL
	case 0xE9, 0xEA, 0xEB:
		return JMP
	case 0xEC, 0xED:
		return IN
	case 0xF0:
		panic("Error: this is the lock prefix")
	case 0xF1:
		return INT1
	case 0xF2:
		panic("Error: this is repne prefix")
	case 0xF3:
		panic("Error: this is rep/e prefix")
	case 0xF4:
		return HLT
	case 0xF5:
		return CMC
	case 0xF6, 0xF7:
		panic("todo modr/m opcode extension")
	case 0xF8:
		return CLC
	case 0xF9:
		return STC
	case 0xFA:
		return CLI
	case 0xFB:
		return STI
	case 0xFC:
		return CLD
	case 0xFD:
		return STD
	case 0xFE:
		panic("todo modr/m opcode extension")
	case 0xFF:
		panic("todo modr/m opcode extension")
	}
	panic("Error: todo")
}
