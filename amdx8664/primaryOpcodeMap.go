package amdx8664

func primaryOpcode(curByte byte, is64Bit bool, isOperandSizeOveride bool) (Instruction, bool, bool, MemSegment, Register, Register) {
	switch curByte {
	case 0x0, 0x1, 0x2, 0x3:
		return ADD, true, false, NoSegment, NoRegister
	case 0x4, 0xC, 0x14, 0x1C, 0x24, 0x2C, 0x34, 0x3C:
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
	case 0x8, 0x9, 0xA, 0xB:
		return OR, true, false, NoSegment, NoRegister
	case 0xD:
		return OR, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0xE:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, ES, NoRegister
		}
	case 0xF:
		panic("Error: 2-byte opcodes")
	case 0x10, 0x11, 0x12, 0x13:
		return ADC, true, false, NoSegment, NoRegister
	case 0x15:
		return ADC, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x16:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, SS, NoRegister
		}
	case 0x17:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return POP, false, false, SS, NoRegister
		}
	case 0x18, 0x19, 0x1A, 0x1B:
		return SBB, true, false, ES, NoRegister
	case 0x1D:
		return SBB, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x1E:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, DS, NoRegister
		}
	case 0x1F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return POP, false, false, DS, NoRegister
		}
	case 0x20, 0x21, 0x22, 0x23:
		return AND, true, false, NoSegment, NoRegister
	case 0x25:
		return AND, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x26:
		panic("Error: this is the ES segment")
	case 0x27:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return DAA, false, false, NoSegment, NoRegister
		}
	case 0x28, 0x29, 0x2A, 0x2B:
		return SUB, true, false, NoSegment, NoRegister
	case 0x2D:
		return SUB, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x2E:
		panic("Error: this is the CS segment")
	case 0x2F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return DAS, false, false, NoSegment, NoRegister
		}
	case 0x30, 0x31, 0x32, 0x33:
		return XOR, true, false, NoSegment, NoRegister
	case 0x35:
		return XOR, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x36:
		panic("Error: this is the SS segment")
	case 0x37:
		return AAA, false, false, NoSegment, NoRegister
	case 0x38, 0x39, 0x3A, 0x3B:
		return CMP, true, false, NoSegment, NoRegister
	case 0x3D:
		return CMP, false, true, NoSegment, RAX // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x3E:
		panic("Error: this is the DS segment")
	case 0x3F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return AAS, false, false, NoSegment, NoRegister
		}
	case 0x40:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EAX
		}
	case 0x41:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, ECX
		}
	case 0x42:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EDX
		}
	case 0x43:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EBX
		}
	case 0x44:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, ESP
		}
	case 0x45:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EBP
		}
	case 0x46:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, ESI
		}
	case 0x47:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EDI
		}
	case 0x48:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EAX
		}
	case 0x49:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, ECX
		}
	case 0x4A:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EDX
		}
	case 0x4B:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EBX
		}
	case 0x4C:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, ESP
		}
	case 0x4D:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EBP
		}
	case 0x4E:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, ESI
		}
	case 0x4F:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EDI
		}
	case 0x50:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EAX
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R8
			} else {
				return PUSH, false, false, NoSegment, RAX
			}
		}
	case 0x51:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, ECX
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R9
			} else {
				return PUSH, false, false, NoSegment, RCX
			}
		}
	case 0x52:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EDX
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R10
			} else {
				return PUSH, false, false, NoSegment, RDX
			}
		}
	case 0x53:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EBX
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R11
			} else {
				return PUSH, false, false, NoSegment, RBX
			}
		}
	case 0x54:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, ESP
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R12
			} else {
				return PUSH, false, false, NoSegment, RSP
			}
		}
	case 0x55:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EBP
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R13
			} else {
				return PUSH, false, false, NoSegment, RBP
			}
		}
	case 0x56:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, ESI
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R14
			} else {
				return PUSH, false, false, NoSegment, RSI
			}
		}
	case 0x57:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EDI
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R15
			} else {
				return PUSH, false, false, NoSegment, RDI
			}
		}
	case 0x58:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EAX
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R8
			} else {
				return POP, false, false, NoSegment, RAX
			}
		}
	case 0x59:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, ECX
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R9
			} else {
				return POP, false, false, NoSegment, RCX
			}
		}
	case 0x5A:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EDX
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R10
			} else {
				return POP, false, false, NoSegment, RDX
			}
		}
	case 0x5B:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EBX
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R11
			} else {
				return POP, false, false, NoSegment, RBX
			}
		}
	case 0x5C:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, ESP
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R12
			} else {
				return POP, false, false, NoSegment, RSP
			}
		}
	case 0x5D:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EBP
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R13
			} else {
				return POP, false, false, NoSegment, RBP
			}
		}
	case 0x5E:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, ESI
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R14
			} else {
				return POP, false, false, NoSegment, RSI
			}
		}
	case 0x5F:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EDI
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R15
			} else {
				return POP, false, false, NoSegment, RDI
			}
		}
	case 0x60:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			panic("todo multiple instruction in 1 opcode")
		}
	case 0x61:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			panic("todo multiple instruction in 1 opcode")
		}
	case 0x62:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return BOUND, true, false, NoSegment, NoRegister
		}
	case 0x63:
		if is64Bit {
			return MOVSXD, true, false, NoSegment, NoRegister
		} else {
			return ARPL, true, false, NoSegment, NoRegister
		}
	case 0x64:
		panic("Error: this is the FS segment")
	case 0x65:
		panic("Error: this is the GS segment")
	case 0x66:
		panic("Error: this is the operand size override prefix")
	case 0x67:
		panic("Error: this is the address size override prefix")
	case 0x68:
		return PUSH, false, true, NoSegment, NoRegister
	case 0x69, 0x6B:
		return IMUL, true, true, NoSegment, NoRegister
	case 0x6A:
		return PUSH, false, true, NoSegment, NoRegister
	case 0x6C:
		panic("todo dont know how to deal with Y operand syntax notation")
	case 0x6D:
		panic("todo multiple instructions in 1 byte")
	case 0x6E:
		panic("todo dont know how to deal with X operand syntax notation")
	case 0x6F:
		panic("todo multiple instructions in 1 byte")
	case 0x70:
		panic("todo dont know how to deal with rip addressing")
	case 0x71:
		panic("todo dont know how to deal with rip addressing")
	case 0x72:
		panic("todo dont know how to deal with rip addressing")
	case 0x73:
		panic("todo dont know how to deal with rip addressing")
	case 0x74:
		panic("todo dont know how to deal with rip addressing")
	case 0x75:
		panic("todo dont know how to deal with rip addressing")
	case 0x76:
		panic("todo dont know how to deal with rip addressing")
	case 0x77:
		panic("todo dont know how to deal with rip addressing")
	case 0x78:
		panic("todo dont know how to deal with rip addressing")
	case 0x79:
		panic("todo dont know how to deal with rip addressing")
	case 0x7A:
		panic("todo dont know how to deal with rip addressing")
	case 0x7B:
		panic("todo dont know how to deal with rip addressing")
	case 0x7C:
		panic("todo dont know how to deal with rip addressing")
	case 0x7D:
		panic("todo dont know how to deal with rip addressing")
	case 0x7E:
		panic("todo dont know how to deal with rip addressing")
	case 0x7F:
		panic("todo dont know how to deal with rip addressing")
	case 0x80:
		panic("todo modr/m opcode extension")
	case 0x81:
		panic("todo modr/m opcode extension")
	case 0x82:
		panic("todo modr/m opcode extension")
	case 0x83:
		panic("todo modr/m opcode extension")
	case 0x84, 0x85:
		return TEST, true, false, NoSegment, NoRegister
	case 0x86, 0x87:
		return XCHG, true, false, NoSegment, NoRegister
	case 0x88, 0x89, 0x8A, 0x8B, 0x8E:
		return MOV, true, false, NoSegment, NoRegister
	case 0x8C:
		panic("todo multiple operand types")
	case 0x8D:
		return LEA, true, false, NoSegment, NoRegister
	case 0x8F:
		panic("todo modr/m opcode extension")
	case 0x90:
		panic("todo could be XCHG or NOP instruction")
	case 0x91:
		return XCHG, false, false, 
	case 0x92:
	case 0x93:
	case 0x94:
	case 0x95:
	case 0x96:
	case 0x97:
	case 0x98:
		panic("todo multiple instructions in 1 byte")
	case 0x99:
		panic("todo multiple instructions in 1 byte")
	case 0x9A, 0x9E, 0x9F, 0xC4, 0xC5, 0xCE:
		panic("todo")
	case 0x9B:
		panic("todo multiple instructions in 1 byte")
	case 0x9C:
		panic("todo multiple instructions in 1 byte")
	case 0x9D:
		panic("todo multiple instructions in 1 byte")
	case 0xA0, 0xA1, 0xA2, 0xA3, 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF:
		panic("")
	case 0xA4:
		return MOVSB
	case 0xA5:
		panic("todo multiple instructions in 1 byte")
	case 0xA6:
		return CMPSB
	case 0xA7:
		panic("todo multiple instructions in 1 byte")
	case 0xA8, 0xA9:
		panic("")
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
