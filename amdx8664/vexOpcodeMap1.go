package amdx8664

func vexOpcodeMap1(curByte byte, opcodeExt [2]bool, isRexB bool) (Instruction, bool, bool, MemSegment, Register, Register, int) {
	switch curByte {
	case 0x0:
		panic("Error: No opcode")
	case 0x1:
		panic("Error: No opcode")
	case 0x2:
		panic("Error: No opcode")
	case 0x3:
		panic("Error: No opcode")
	case 0x4:
		panic("Error: No opcode")
	case 0x5:
		panic("Error: No opcode")
	case 0x6:
		panic("Error: No opcode")
	case 0x7:
		panic("Error: No opcode")
	case 0x8:
		panic("Error: No opcode")
	case 0x9:
		panic("Error: No opcode")
	case 0xA:
		panic("Error: No opcode")
	case 0xB:
		panic("Error: No opcode")
	case 0xC:
		panic("Error: No opcode")
	case 0xD:
		panic("Error: No opcode")
	case 0xE:
		panic("Error: No opcode")
	case 0xF:
		panic("Error: No opcode")
	case 0x10:
		if isOperandSizeOverride {
			// 0x66
			return MOVUPD, true, false, NoSegment, NoRegister
		} else if isRep1 {
			// 0xF3
			return MOVSS, true, false, NoSegment, NoRegister
		} else if isRep0 {
			// 0xF2
			return MOVSD, true, false, NoSegment, NoRegister
		} else {
			return MOVUPS, true, false, NoSegment, NoRegister
		}
	case 0x11:
		return ADC
	case 0x12:
		return ADC
	case 0x13:
		return ADC
	case 0x14:
		return ADC
	case 0x15:
		return ADC
	case 0x16:
		return PUSH
	case 0x17:
		return POP
	case 0x18:
		panic("Error: No opcode")
	case 0x19:
		panic("Error: No opcode")
	case 0x1A:
		panic("Error: No opcode")
	case 0x1B:
		panic("Error: No opcode")
	case 0x1C:
		panic("Error: No opcode")
	case 0x1D:
		panic("Error: No opcode")
	case 0x1E:
		panic("Error: No opcode")
	case 0x1F:
		panic("Error: No opcode")
	case 0x20:
		return AND
	case 0x21:
		return AND
	case 0x22:
		return AND
	case 0x23:
		return AND
	case 0x24:
		return AND
	case 0x25:
		return AND
	case 0x26:
		// return ADD
	case 0x27:
		return DAA
	case 0x28:
		return SUB
	case 0x29:
		return SUB
	case 0x2A:
		return SUB
	case 0x2B:
		return SUB
	case 0x2C:
		return SUB
	case 0x2D:
		return SUB
	case 0x2E:
		// return ADD
	case 0x2F:
		return DAS
	case 0x30:
		return XOR
	case 0x31:
		return XOR
	case 0x32:
		return XOR
	case 0x33:
		return XOR
	case 0x34:
		return XOR
	case 0x35:
		return XOR
	case 0x36:
		// return ADD
	case 0x37:
		return AAA
	case 0x38:
		return CMP
	case 0x39:
		return CMP
	case 0x3A:
		return CMP
	case 0x3B:
		return CMP
	case 0x3C:
		return CMP
	case 0x3D:
		return CMP
	case 0x3E:
		// return ADD
	case 0x3F:
		return AAS
	case 0x40:
		// return ADD
	case 0x41:
		// return ADD
	case 0x42:
		// return ADD
	case 0x43:
		// return ADD
	case 0x44:
		// return ADD
	case 0x45:
		// return ADD
	case 0x46:
		// return ADD
	case 0x47:
		// return ADD
	case 0x48:
		// return ADD
	case 0x49:
		// return ADD
	case 0x4A:
		// return ADD
	case 0x4B:
		// return ADD
	case 0x4C:
		// return ADD
	case 0x4D:
		// return ADD
	case 0x4E:
		// return ADD
	case 0x4F:
		// return ADD
	case 0x50:
		return PUSH
	case 0x51:
		return PUSH
	case 0x52:
		return PUSH
	case 0x53:
		return PUSH
	case 0x54:
		return PUSH
	case 0x55:
		return PUSH
	case 0x56:
		return PUSH
	case 0x57:
		return PUSH
	case 0x58:
		return POP
	case 0x59:
		return POP
	case 0x5A:
		return POP
	case 0x5B:
		return POP
	case 0x5C:
		return POP
	case 0x5D:
		return POP
	case 0x5E:
		return POP
	case 0x5F:
		return POP
	case 0x60:
		// return ADD operand size override
	case 0x61:
		// return ADD operand size override
	case 0x62:
		return BOUND
	case 0x63:
		return ADD
	case 0x64:
		return ADD
	case 0x65:
		return ADD
	case 0x66:
		return ADD
	case 0x67:
		return ADD
	case 0x68:
		return ADD
	case 0x69:
		return ADD
	case 0x6A:
		return ADD
	case 0x6B:
		return ADD
	case 0x6C:
		return ADD
	case 0x6D:
		return ADD
	case 0x6E:
		return ADD
	case 0x6F:
		return ADD
	case 0x70:
		return ADD
	case 0x71:
		return ADD
	case 0x72:
		return ADD
	case 0x73:
		return ADD
	case 0x74:
		return ADD
	case 0x75:
		return ADD
	case 0x76:
		return ADD
	case 0x77:
		return ADD
	case 0x78:
		return ADD
	case 0x79:
		return ADD
	case 0x7A:
		return ADD
	case 0x7B:
		return ADD
	case 0x7C:
		return ADD
	case 0x7D:
		return ADD
	case 0x7E:
		return ADD
	case 0x7F:
		return ADD
	case 0x80:
		return ADD
	case 0x81:
		return ADD
	case 0x82:
		return ADD
	case 0x83:
		return ADD
	case 0x84:
		return ADD
	case 0x85:
		return ADD
	case 0x86:
		return ADD
	case 0x87:
		return ADD
	case 0x88:
		return ADD
	case 0x89:
		return ADD
	case 0x8A:
		return ADD
	case 0x8B:
		return ADD
	case 0x8C:
		return ADD
	case 0x8D:
		return ADD
	case 0x8E:
		return ADD
	case 0x8F:
		return ADD
	case 0x90:
		return ADD
	case 0x91:
		return ADD
	case 0x92:
		return ADD
	case 0x93:
		return ADD
	case 0x94:
		return ADD
	case 0x95:
		return ADD
	case 0x96:
		return ADD
	case 0x97:
		return ADD
	case 0x98:
		return ADD
	case 0x99:
		return ADD
	case 0x9A:
		return ADD
	case 0x9B:
		return ADD
	case 0x9C:
		return ADD
	case 0x9D:
		return ADD
	case 0x9E:
		return ADD
	case 0x9F:
		return ADD
	case 0xA0:
		return ADD
	case 0xA1:
		return ADD
	case 0xA2:
		return ADD
	case 0xA3:
		return ADD
	case 0xA4:
		return ADD
	case 0xA5:
		return ADD
	case 0xA6:
		return ADD
	case 0xA7:
		return ADD
	case 0xA8:
		return ADD
	case 0xA9:
		return ADD
	case 0xAA:
		return ADD
	case 0xAB:
		return ADD
	case 0xAC:
		return ADD
	case 0xAD:
		return ADD
	case 0xAE:
		return ADD
	case 0xAF:
		return ADD
	case 0xB0:
		return ADD
	case 0xB1:
		return ADD
	case 0xB2:
		return ADD
	case 0xB3:
		return ADD
	case 0xB4:
		return ADD
	case 0xB5:
		return ADD
	case 0xB6:
		return ADD
	case 0xB7:
		return ADD
	case 0xB8:
		return ADD
	case 0xB9:
		return ADD
	case 0xBA:
		return ADD
	case 0xBB:
		return ADD
	case 0xBC:
		return ADD
	case 0xBD:
		return ADD
	case 0xBE:
		return ADD
	case 0xBF:
		return ADD
	case 0xC0:
		return ADD
	case 0xC1:
		return ADD
	case 0xC2:
		return ADD
	case 0xC3:
		return ADD
	case 0xC4:
		return ADD
	case 0xC5:
		return ADD
	case 0xC6:
		return ADD
	case 0xC7:
		return ADD
	case 0xC8:
		return ADD
	case 0xC9:
		return ADD
	case 0xCA:
		return ADD
	case 0xCB:
		return ADD
	case 0xCC:
		return ADD
	case 0xCD:
		return ADD
	case 0xCE:
		return ADD
	case 0xCF:
		return ADD
	case 0xD0:
		return ADD
	case 0xD1:
		return ADD
	case 0xD2:
		return ADD
	case 0xD3:
		return ADD
	case 0xD4:
		return ADD
	case 0xD5:
		return ADD
	case 0xD6:
		return ADD
	case 0xD7:
		return ADD
	case 0xD8:
		return ADD
	case 0xD9:
		return ADD
	case 0xDA:
		return ADD
	case 0xDB:
		return ADD
	case 0xDC:
		return ADD
	case 0xDD:
		return ADD
	case 0xDE:
		return ADD
	case 0xDF:
		return ADD
	case 0xE0:
		return ADD
	case 0xE1:
		return ADD
	case 0xE2:
		return ADD
	case 0xE3:
		return ADD
	case 0xE4:
		return ADD
	case 0xE5:
		return ADD
	case 0xE6:
		return ADD
	case 0xE7:
		return ADD
	case 0xE8:
		return ADD
	case 0xE9:
		return ADD
	case 0xEA:
		return ADD
	case 0xEB:
		return ADD
	case 0xEC:
		return ADD
	case 0xED:
		return ADD
	case 0xEE:
		return ADD
	case 0xEF:
		return ADD
	case 0xF0:
		return ADD
	case 0xF1:
		return ADD
	case 0xF2:
		return ADD
	case 0xF3:
		return ADD
	case 0xF4:
		return ADD
	case 0xF5:
		return ADD
	case 0xF6:
		return ADD
	case 0xF7:
		return ADD
	case 0xF8:
		return ADD
	case 0xF9:
		return ADD
	case 0xFA:
		return ADD
	case 0xFB:
		return ADD
	case 0xFC:
		return ADD
	case 0xFD:
		return ADD
	case 0xFE:
		return ADD
	case 0xFF:
		return ADD
	}
	panic("Error: todo")
}
