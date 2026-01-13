package amdx8664

func primaryOpcode(curByte byte, is64Bit bool, isOperandSizeOveride bool) (Instruction, bool, bool, MemSegment, Register, Register, int, bool) {
	switch curByte {
	case 0x0, 0x1:
		return ADD, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x2, 0x3:
		return ADD, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x4, 0xC, 0x14, 0x1C, 0x24, 0x2C, 0x34, 0x3C:
		panic("todo dont know what AL opcode syntax notation")
	case 0x5:
		return ADD, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x6:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, ES, NoRegister, NoRegister, 0, false
		}
	case 0x7:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return POP, false, false, ES, NoRegister, NoRegister, 0, false
		}
	case 0x8, 0x9:
		return OR, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xA, 0xB:
		return OR, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0xD:
		return OR, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0xE:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, ES, NoRegister, NoRegister, 0, false
		}
	case 0xF:
		panic("Error: 2-byte opcodes")
	case 0x10, 0x11:
		return ADC, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x12, 0x13:
		return ADC, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x15:
		return ADC, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x16:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, SS, NoRegister, NoRegister, 0, false
		}
	case 0x17:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return POP, false, false, SS, NoRegister, NoRegister, 0, false
		}
	case 0x18, 0x19:
		return SBB, true, false, ES, NoRegister, NoRegister, 0, false
	case 0x1A, 0x1B:
		return SBB, true, false, ES, NoRegister, NoRegister, 0, true
	case 0x1D:
		return SBB, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x1E:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return PUSH, false, false, DS, NoRegister, NoRegister, 0, false
		}
	case 0x1F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return POP, false, false, DS, NoRegister, NoRegister, 0, false
		}
	case 0x20, 0x21:
		return AND, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x22, 0x23:
		return AND, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x25:
		return AND, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x26:
		panic("Error: this is the ES segment")
	case 0x27:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return DAA, false, false, NoSegment, NoRegister, NoRegister, 0, false
		}
	case 0x28, 0x29:
		return SUB, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x2A, 0x2B:
		return SUB, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x2D:
		return SUB, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x2E:
		panic("Error: this is the CS segment")
	case 0x2F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return DAS, false, false, NoSegment, NoRegister, NoRegister, 0, false
		}
	case 0x30, 0x31:
		return XOR, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x32, 0x33:
		return XOR, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x35:
		return XOR, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x36:
		panic("Error: this is the SS segment")
	case 0x37:
		return AAA, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x38, 0x39:
		return CMP, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x3A, 0x3B:
		return CMP, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x3D:
		return CMP, false, true, NoSegment, RAX, NoRegister, 0, false // todo need to figure out operand addressing size if using eax, ax, rax
	case 0x3E:
		panic("Error: this is the DS segment")
	case 0x3F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return AAS, false, false, NoSegment, NoRegister, NoRegister, 0, false
		}
	case 0x40:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EAX, NoRegister, 0, false
		}
	case 0x41:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, ECX, NoRegister, 0, false
		}
	case 0x42:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EDX, NoRegister, 0, false
		}
	case 0x43:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EBX, NoRegister, 0, false
		}
	case 0x44:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, ESP, NoRegister, 0, false
		}
	case 0x45:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EBP, NoRegister, 0, false
		}
	case 0x46:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, ESI, NoRegister, 0, false
		}
	case 0x47:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return INC, false, false, NoSegment, EDI, NoRegister, 0, false
		}
	case 0x48:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EAX, NoRegister, 0, false
		}
	case 0x49:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, ECX, NoRegister, 0, false
		}
	case 0x4A:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EDX, NoRegister, 0, false
		}
	case 0x4B:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EBX, NoRegister, 0, false
		}
	case 0x4C:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, ESP, NoRegister, 0, false
		}
	case 0x4D:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EBP, NoRegister, 0, false
		}
	case 0x4E:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, ESI, NoRegister, 0, false
		}
	case 0x4F:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, NoSegment, EDI, NoRegister, 0, false
		}
	case 0x50:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EAX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R8, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RAX, NoRegister, 0, false
			}
		}
	case 0x51:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, ECX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R9, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RCX, NoRegister, 0, false
			}
		}
	case 0x52:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EDX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R10, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RDX, NoRegister, 0, false
			}
		}
	case 0x53:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EBX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R11, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RBX, NoRegister, 0, false
			}
		}
	case 0x54:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, ESP, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R12, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RSP, NoRegister, 0, false
			}
		}
	case 0x55:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EBP, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R13, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RBP, NoRegister, 0, false
			}
		}
	case 0x56:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, ESI, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R14, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RSI, NoRegister, 0, false
			}
		}
	case 0x57:
		if is64Bit {
			// todo operand size override check
			return PUSH, false, false, NoSegment, EDI, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return PUSH, false, false, NoSegment, R15, NoRegister, 0, false
			} else {
				return PUSH, false, false, NoSegment, RDI, NoRegister, 0, false
			}
		}
	case 0x58:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EAX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R8, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RAX, NoRegister, 0, false
			}
		}
	case 0x59:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, ECX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R9, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RCX, NoRegister, 0, false
			}
		}
	case 0x5A:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EDX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R10, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RDX, NoRegister, 0, false
			}
		}
	case 0x5B:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EBX, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R11, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RBX, NoRegister, 0, false
			}
		}
	case 0x5C:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, ESP, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R12, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RSP, NoRegister, 0, false
			}
		}
	case 0x5D:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EBP, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R13, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RBP, NoRegister, 0, false
			}
		}
	case 0x5E:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, ESI, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R14, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RSI, NoRegister, 0, false
			}
		}
	case 0x5F:
		if is64Bit {
			// todo operand size override check
			return POP, false, false, NoSegment, EDI, NoRegister, 0, false
		} else {
			if isOperandSizeOveride {
				return POP, false, false, NoSegment, R15, NoRegister, 0, false
			} else {
				return POP, false, false, NoSegment, RDI, NoRegister, 0, false
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
			return BOUND, true, false, NoSegment, NoRegister, NoRegister, 0, true
		}
	case 0x63:
		if is64Bit {
			return MOVSXD, true, false, NoSegment, NoRegister, NoRegister, 0, true
		} else {
			return ARPL, true, false, NoSegment, NoRegister, NoRegister, 0, false
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
		return PUSH, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0x69, 0x6B:
		panic("Error: todo has 3 operands")
	case 0x6A:
		return PUSH, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0x6C:
		panic("todo dont know how to deal with Y operand syntax notation")
	case 0x6D:
		panic("todo multiple instructions in 1 byte")
	case 0x6E:
		panic("todo dont know how to deal with X operand syntax notation")
	case 0x6F:
		panic("todo multiple instructions in 1 byte")
	case 0x70:
		// panic("todo dont know how to deal with rip addressing")
		return JO, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x71:
		return JNO, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x72:
		return JB, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x73:
		// panic("todo dont know how to deal with rip addressing")
		return JNB, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x74:
		// panic("todo dont know how to deal with rip addressing")
		return JZ, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x75:
		// panic("todo dont know how to deal with rip addressing")
		return JNZ, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x76:
		// panic("todo dont know how to deal with rip addressing")
		return JBE, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x77:
		// panic("todo dont know how to deal with rip addressing")
		return JNBE, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x78:
		return JS, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x79:
		return JNS, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x7A:
		return JP, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x7B:
		return JNP, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x7C:
		return JL, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x7D:
		return JNL, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x7E:
		return JLE, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x7F:
		return JNLE, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x80, 0x81, 0x82, 0x83:
		return NoInstruction, true, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0x84, 0x85:
		return TEST, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x86, 0x87:
		return XCHG, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x88, 0x89:
		return MOV, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x8A, 0x8B:
		return MOV, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x8E:
		panic("Error: todo don't know segment registers")
	case 0x8C:
		panic("todo multiple operand types")
	case 0x8D:
		return LEA, true, false, NoSegment, NoRegister, NoRegister, 0, true
	case 0x8F:
		panic("Error: todo don't know")
	case 0x90:
		panic("todo could be XCHG or NOP instruction")
	case 0x91:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, NoSegment, R9, RAX, 0, false
			} else {
				return XCHG, false, false, NoSegment, RCX, RAX, 0, false
			}
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x92:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, NoSegment, R10, RAX, 0, false
			} else {
				return XCHG, false, false, NoSegment, RDX, RAX, 0, false
			}
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x93:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, NoSegment, R11, RAX, 0, false
			} else {
				return XCHG, false, false, NoSegment, RBX, RAX, 0, false
			}
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x94:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, NoSegment, R12, RAX, 0, false
			} else {
				return XCHG, false, false, NoSegment, RSP, RAX, 0, false
			}
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x95:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, NoSegment, R13, RAX, 0, false
			} else {
				return XCHG, false, false, NoSegment, RBP, RAX, 0, false
			}
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x96:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, NoSegment, R14, RAX, 0, false
			} else {
				return XCHG, false, false, NoSegment, RSI, RAX, 0, false
			}
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x97:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, NoSegment, R15, RAX, 0, false
			} else {
				return XCHG, false, false, NoSegment, RDI, RAX, 0, false
			}
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x98:
		panic("todo multiple instructions in 1 byte")
	case 0x99:
		panic("todo multiple instructions in 1 byte")
	case 0x9A:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return CALL, false, false, NoSegment, NoRegister, NoRegister, 0, false
		}
	case 0x9B:
		panic("todo multiple instructions in 1 byte")
	case 0x9C:
		panic("todo multiple instructions in 1 byte")
	case 0x9D:
		panic("todo multiple instructions in 1 byte")
	case 0x9E:
		return SAHF, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0x9F:
		return LAHF, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xA0, 0xA2:
		return MOV, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xA1:
		return MOV, false, true, NoSegment, RAX, NoRegister, 0, false
	case 0xA3:
		return MOV, false, true, NoSegment, NoRegister, RAX, 0, false
	case 0xA4:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xA5:
		panic("todo multiple instructions in 1 byte")
	case 0xA6:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xA7:
		panic("todo multiple instructions in 1 byte")
	case 0xA8:
		return TEST, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xA9:
		return TEST, false, true, NoSegment, RAX, NoRegister, 0, false
	case 0xAA:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xAB:
		panic("todo multiple instructions in 1 byte")
	case 0xAC:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xAD:
		panic("todo multiple instructions in 1 byte")
	case 0xAE:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xAF:
		panic("todo multiple instructions in 1 byte")
	case 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF:
		panic("todo multiple operands")
	case 0xC0, 0xC1:
		return NoInstruction, true, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xC2:
		return RET, false, true, NoSegment, NoRegister, NoRegister, 0, false // near
	case 0xC3:
		return RET, false, false, NoSegment, NoRegister, NoRegister, 0, false // near
	case 0xC4:
		if is64Bit {
			panic("Error: this is the vex escape prefix")
		} else {
			return LES, true, false, NoSegment, NoRegister, NoRegister, 0, true
		}
	case 0xC5:
		if is64Bit {
			panic("Error: this is the vex escape prefix")
		} else {
			return LDS, true, false, NoSegment, NoRegister, NoRegister, 0, true
		}
	case 0xC6, 0xC7:
		return NoInstruction, true, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xC8:
		return ENTER, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xC9:
		return LEAVE, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xCA:
		return RET, false, true, NoSegment, NoRegister, NoRegister, 0, false // far
	case 0xCB:
		return RET, false, false, NoSegment, NoRegister, NoRegister, 0, false // far
	case 0xCC:
		return INT3, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xCD:
		return INT, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xCE:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return INTO, false, false, NoSegment, NoRegister, NoRegister, 0, false
		}
	case 0xCF:
		panic("todo multiple instructions in 1 byte")
	case 0xD0:
		return NoInstruction, true, true, NoSegment, NoRegister, NoRegister, 1, false
	case 0xD1:
		return NoInstruction, true, true, NoSegment, NoRegister, NoRegister, 1, false
	case 0xD2:
		panic("todo dont know how to deal with DX")
	case 0xD3:
		panic("todo dont know how to deal with DX")
	case 0xD4:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return AAM, false, true, NoSegment, NoRegister, NoRegister, 0, false
		}
	case 0xD5:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return AAD, false, true, NoSegment, NoRegister, NoRegister, 0, false
		}
	case 0xD6:
		panic("Error: Invalid")
	case 0xD7:
		panic("todo multiple instructions in 1 byte")
	case 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF:
		panic("Error: this is x87 opcodes")
	case 0xE0:
		panic("todo multiple instructions in 1 byte")
	case 0xE1:
		panic("todo multiple instructions in 1 byte")
	case 0xE2:
		panic("todo dont know how to deal with rip addressing")
	case 0xE3:
		panic("todo dont know how to deal with rip addressing")
	case 0xE4:
		return IN, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xE5:
		return IN, false, true, NoSegment, EAX, NoRegister, 0, false
	case 0xE6:
		return OUT, false, true, NoSegment, NoRegister, NoRegister, 0, false
	case 0xE7:
		return OUT, false, true, NoSegment, NoRegister, EAX, 0, false
	case 0xE8:
		panic("todo dont know how to deal with rip addressing")
	case 0xE9:
		panic("todo dont know how to deal with rip addressing")
	case 0xEA:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		} else {
			return JMP, false, true, NoSegment, NoRegister, NoRegister, 0, false // dont know if this is correct
		}
	case 0xEB:
		panic("todo dont know how to deal with rip addressing")
	case 0xEC:
		panic("todo dont know how to deal with DX")
	case 0xED:
		panic("todo dont know how to deal with DX")
	case 0xEE:
		panic("todo dont know how to deal with DX")
	case 0xEF:
		panic("todo dont know how to deal with DX")
	case 0xF0:
		panic("Error: this is the lock prefix")
	case 0xF1:
		return INT1, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xF2:
		panic("Error: this is repne prefix")
	case 0xF3:
		panic("Error: this is rep/e prefix")
	case 0xF4:
		return HLT, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xF5:
		return CMC, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xF6:
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xF7:
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xF8:
		return CLC, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xF9:
		return STC, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xFA:
		return CLI, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xFB:
		return STI, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xFC:
		return CLD, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xFD:
		return STD, false, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xFE:
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0, false
	case 0xFF:
		return NoInstruction, true, false, NoSegment, NoRegister, NoRegister, 0, false
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG1(opcode byte, modrmReg [3]bool, is64bit bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x80:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return OR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return ADC, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return SBB, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return AND, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return SUB, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return XOR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return CMP, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x81:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return OR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return ADC, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return SBB, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return AND, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return SUB, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return XOR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return CMP, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x82:
		panic("todo")
	case 0x83:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return OR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return ADC, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return SBB, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return AND, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return SUB, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return XOR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return CMP, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG1a(opcode byte, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0x8F:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return POP, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}, [3]bool{false, true, false}, [3]bool{false, true, true}, [3]bool{true, false, false}, [3]bool{true, false, true}, [3]bool{true, true, false}, [3]bool{true, true, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG2(opcode byte, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xC0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return ROR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return RCL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return RCR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return ROR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return RCL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return RCR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{false, false, true}:
			return ROR, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{false, true, false}:
			return RCL, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{false, true, true}:
			return RCR, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, false, NoSegment, NoRegister, NoRegister, 1
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{false, false, true}:
			return ROR, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{false, true, false}:
			return RCL, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{false, true, true}:
			return RCR, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, false, NoSegment, NoRegister, NoRegister, 1
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, false, NoSegment, NoRegister, NoRegister, 1
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD2:
		panic("todo")
	case 0xD3:
		panic("todo")
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG3(opcode byte, modrmReg [3]bool, is64bit bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xF6:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return TEST, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return TEST, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return NOT, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return NEG, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return MUL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return IMUL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return DIV, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return IDIV, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF7:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return TEST, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return TEST, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return NOT, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return NEG, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return MUL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return IMUL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return DIV, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, true}:
			return IDIV, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG4(opcode byte, modrmReg [3]bool, is64bit bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xFE:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return INC, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return DEC, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG5(opcode byte, modrmReg [3]bool, is64bit bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xFF:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return INC, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return DEC, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return CALL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return CALL, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return JMP, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return JMP, true, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return PUSH, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG11(opcode byte, modrmReg [3]bool, is64bit bool) (Instruction, bool, MemSegment, Register, Register, int) {
	switch opcode {
	case 0xC6:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return MOV, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC7:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return MOV, true, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
