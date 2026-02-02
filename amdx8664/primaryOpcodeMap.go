package amdx8664

func primaryOpcode(curByte byte, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, bool, bool, bool, MemSegment, Register, Register, int, bool, int, int) {
	noImmediateBytes, noDisplacementBytes := 4, 4
	switch curByte {
	case 0x00, 0x01:
		return ADD, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x02, 0x03:
		return ADD, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x04:
		return ADD, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x05:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return ADD, true, false, false, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0x06:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, false, false, false, ES, NoRegister, NoRegister, 0, false, 0, 0
	case 0x07:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return POP, false, false, false, ES, NoRegister, NoRegister, 0, false, 0, 0
	case 0x08, 0x09:
		return OR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x0A, 0x0B:
		return OR, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x0C:
		return OR, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x0D:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return OR, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0x0E:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, false, false, false, ES, NoRegister, NoRegister, 0, false, 0, 0
	case 0x0F:
		panic("Error: 2-byte opcodes")
	case 0x10, 0x11:
		return ADC, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x12, 0x13:
		return ADC, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x14:
		return ADC, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x15:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return ADC, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0x16:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, false, false, false, SS, NoRegister, NoRegister, 0, false, 0, 0
	case 0x17:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return POP, false, false, false, SS, NoRegister, NoRegister, 0, false, 0, 0
	case 0x18, 0x19:
		return SBB, true, false, false, ES, NoRegister, NoRegister, 0, false, 0, 0
	case 0x1A, 0x1B:
		return SBB, true, false, false, ES, NoRegister, NoRegister, 0, true, 0, 0
	case 0x1C:
		return SBB, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x1D:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return SBB, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0x1E:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, false, false, false, DS, NoRegister, NoRegister, 0, false, 0, 0
	case 0x1F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return POP, false, false, false, DS, NoRegister, NoRegister, 0, false, 0, 0
	case 0x20, 0x21:
		return AND, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x22, 0x23:
		return AND, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x24:
		return AND, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x25:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return AND, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0x26:
		panic("Error: this is the ES segment")
	case 0x27:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return DAA, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x28, 0x29:
		return SUB, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x2A, 0x2B:
		return SUB, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x2C:
		return SUB, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x2D:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return SUB, false, false, true, NoSegment, RAX, NoRegister, 0, false, 0, 0
	case 0x2E:
		panic("Error: this is the CS segment")
	case 0x2F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return DAS, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x30, 0x31:
		return XOR, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x32, 0x33:
		return XOR, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x34:
		return XOR, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x35:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return XOR, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0x36:
		panic("Error: this is the SS segment")
	case 0x37:
		if is64Bit {
			panic("Error: instruction invalid in 64-bit mode")
		}
		return AAA, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x38, 0x39:
		return CMP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x3A, 0x3B:
		return CMP, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x3C:
		return CMP, false, false, true, NoSegment, AL, NoRegister, 0, false, 0, 1
	case 0x3D:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return CMP, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0x3E:
		panic("Error: this is the DS segment")
	case 0x3F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return AAS, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x40:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, EAX, NoRegister, 0, false, 0, 0
	case 0x41:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, ECX, NoRegister, 0, false, 0, 0
	case 0x42:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, EDX, NoRegister, 0, false, 0, 0
	case 0x43:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, EBX, NoRegister, 0, false, 0, 0
	case 0x44:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, ESP, NoRegister, 0, false, 0, 0
	case 0x45:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, EBP, NoRegister, 0, false, 0, 0
	case 0x46:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, ESI, NoRegister, 0, false, 0, 0
	case 0x47:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, false, false, false, NoSegment, EDI, NoRegister, 0, false, 0, 0
	case 0x48:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, false, false, false, NoSegment, EAX, NoRegister, 0, false, 0, 0
	case 0x49:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, false, false, false, NoSegment, ECX, NoRegister, 0, false, 0, 0
	case 0x4A:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, false, false, false, NoSegment, EDX, NoRegister, 0, false, 0, 0
	case 0x4B:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, false, false, false, NoSegment, EBX, NoRegister, 0, false, 0, 0
	case 0x4C:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, false, false, false, NoSegment, ESP, NoRegister, 0, false, 0, 0
	case 0x4D:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, false, false, false, NoSegment, EBP, NoRegister, 0, false, 0, 0
	case 0x4E:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, false, false, false, NoSegment, ESI, NoRegister, 0, false, 0, 0
		}
	case 0x4F:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, false, false, false, NoSegment, EDI, NoRegister, 0, false, 0, 0
	case 0x50:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R8, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, EAX, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RAX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, EAX, NoRegister, 0, false, 0, 0
	case 0x51:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R9, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, ECX, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RCX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, RCX, NoRegister, 0, false, 0, 0
	case 0x52:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R10, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, EDX, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RDX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, RDX, NoRegister, 0, false, 0, 0
	case 0x53:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R11, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, EBX, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RBX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, RBX, NoRegister, 0, false, 0, 0
	case 0x54:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R12, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, ESP, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RSP, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, RSP, NoRegister, 0, false, 0, 0
	case 0x55:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R13, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, EBP, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RBP, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, RBP, NoRegister, 0, false, 0, 0
	case 0x56:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R14, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, ESI, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RSI, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, RSI, NoRegister, 0, false, 0, 0
	case 0x57:
		if is64Bit {
			if isRexW {
				return PUSH, false, false, false, NoSegment, R15, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return PUSH, false, false, false, NoSegment, EDI, NoRegister, 0, false, 0, 0
			}
			return PUSH, false, false, false, NoSegment, RDI, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return PUSH, false, false, false, NoSegment, RDI, NoRegister, 0, false, 0, 0
	case 0x58:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, R8, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, EAX, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RAX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return POP, false, false, false, NoSegment, RAX, NoRegister, 0, false, 0, 0
	case 0x59:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, R9, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, ECX, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RCX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return POP, false, false, false, NoSegment, RCX, NoRegister, 0, false, 0, 0
	case 0x5A:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, R10, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, EDX, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RDX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return POP, false, false, false, NoSegment, RDX, NoRegister, 0, false, 0, 0
	case 0x5B:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, EBX, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, R11, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RBX, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return POP, false, false, false, NoSegment, RBX, NoRegister, 0, false, 0, 0
	case 0x5C:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, R12, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, ESP, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RSP, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return POP, false, false, false, NoSegment, RSP, NoRegister, 0, false, 0, 0
	case 0x5D:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, R13, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, EBP, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RBP, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return POP, false, false, false, NoSegment, RBP, NoRegister, 0, false, 0, 0
	case 0x5E:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, R14, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, ESI, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RSI, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return POP, false, false, false, NoSegment, RSI, NoRegister, 0, false, 0, 0
	case 0x5F:
		if is64Bit {
			if isRexW {
				return POP, false, false, false, NoSegment, R15, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return POP, false, false, false, NoSegment, EDI, NoRegister, 0, false, 0, 0
			}
			return POP, false, false, false, NoSegment, RDI, NoRegister, 0, false, 0, 0
		}
		return POP, false, false, false, NoSegment, RDI, NoRegister, 0, false, 0, 0
	case 0x60:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		if isOperandSizeOveride {
			return PUSHA, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return PUSHD, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x61:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		if isOperandSizeOveride {
			return POPA, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		return POPD, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x62:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return BOUND, false, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x63:
		if is64Bit {
			return MOVSXD, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
		}
		return ARPL, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x64:
		panic("Error: this is the FS segment")
	case 0x65:
		panic("Error: this is the GS segment")
	case 0x66:
		panic("Error: this is the operand size override prefix")
	case 0x67:
		panic("Error: this is the address size override prefix")
	case 0x68:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return PUSH, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, noImmediateBytes, 0
	case 0x69, 0x6B:
		panic("Error: todo has 3 operands")
	case 0x6A:
		return PUSH, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 1, 0
	case 0x6C:
		panic("todo dont know how to deal with Y operand syntax notation")
	case 0x6D:
		panic("todo multiple instructions in 1 byte")
	case 0x6E:
		panic("todo dont know how to deal with X operand syntax notation")
	case 0x6F:
		panic("todo multiple instructions in 1 byte")
	case 0x70:
		return JO, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x71:
		return JNO, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x72:
		return JB, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x73:
		return JNB, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x74:
		return JZ, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x75:
		return JNZ, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x76:
		return JBE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x77:
		return JNBE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x78:
		return JS, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x79:
		return JNS, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x7A:
		return JP, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x7B:
		return JNP, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x7C:
		return JL, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x7D:
		return JNL, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x7E:
		return JLE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x7F:
		return JNLE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0x80, 0x81, 0x82, 0x83:
		return NoInstruction, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0 // todo need to figure out how many immediate bytes
	case 0x84, 0x85:
		return TEST, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x86, 0x87:
		return XCHG, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x88, 0x89:
		return MOV, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x8A, 0x8B:
		return MOV, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x8C:
		panic("todo multiple operand types")
	case 0x8D:
		return LEA, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0x8E:
		panic("Error: todo don't know segment registers")
	case 0x8F:
		panic("Error: todo don't know")
	case 0x90:
		panic("todo could be XCHG or NOP instruction")
	case 0x91:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, false, NoSegment, R9, RAX, 0, false, 0, 0
			}
			return XCHG, false, false, false, NoSegment, RCX, RAX, 0, false, 0, 0
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x92:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, false, NoSegment, R10, RAX, 0, false, 0, 0
			}
			return XCHG, false, false, false, NoSegment, RDX, RAX, 0, false, 0, 0
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x93:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, false, NoSegment, R11, RAX, 0, false, 0, 0
			}
			return XCHG, false, false, false, NoSegment, RBX, RAX, 0, false, 0, 0
		} else {
			panic("todo possibly a different operand size if 32 bit")
		}
	case 0x94:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, false, NoSegment, R12, RAX, 0, false, 0, 0
			}
			return XCHG, false, false, false, NoSegment, RSP, RAX, 0, false, 0, 0
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x95:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, false, NoSegment, R13, RAX, 0, false, 0, 0
			}
			return XCHG, false, false, false, NoSegment, RBP, RAX, 0, false, 0, 0
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x96:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, false, NoSegment, R14, RAX, 0, false, 0, 0
			}
			return XCHG, false, false, false, NoSegment, RSI, RAX, 0, false, 0, 0
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x97:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, false, false, false, NoSegment, R15, RAX, 0, false, 0, 0
			}
			return XCHG, false, false, false, NoSegment, RDI, RAX, 0, false, 0, 0
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x98:
		panic("todo multiple instructions in 1 byte")
	case 0x99:
		panic("todo multiple instructions in 1 byte")
	case 0x9A:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return CALL, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9B:
		panic("todo multiple instructions in 1 byte")
	case 0x9C:
		panic("todo multiple instructions in 1 byte")
	case 0x9D:
		panic("todo multiple instructions in 1 byte")
	case 0x9E:
		return SAHF, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0x9F:
		return LAHF, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xA0:
		return MOV, false, false, true, NoSegment, AL, NoRegister, 0, false, 1, 0
	case 0xA1:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return MOV, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
	case 0xA2:
		return MOV, false, false, true, NoSegment, NoRegister, AL, 0, false, 1, 0 // todo need to figure out how many immediate bytes
	case 0xA3:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return MOV, false, false, true, NoSegment, NoRegister, RAX, 0, false, noImmediateBytes, 0
	case 0xA4:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xA5:
		panic("todo multiple instructions in 1 byte")
	case 0xA6:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xA7:
		panic("todo multiple instructions in 1 byte")
	case 0xA8:
		return TEST, false, false, true, NoSegment, AL, NoRegister, 0, false, 1, 0
	case 0xA9:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return TEST, false, false, true, NoSegment, RAX, NoRegister, 0, false, noImmediateBytes, 0
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
	case 0xB0:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R8B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, AL, NoRegister, 0, false, 1, 0
	case 0xB1:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R9B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, CL, NoRegister, 0, false, 1, 0
	case 0xB2:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R10B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, DL, NoRegister, 0, false, 1, 0
	case 0xB3:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R11B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, BL, NoRegister, 0, false, 1, 0
	case 0xB4:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R12B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, AH, NoRegister, 0, false, 1, 0
	case 0xB5:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R13B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, CH, NoRegister, 0, false, 1, 0
	case 0xB6:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R14B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, DH, NoRegister, 0, false, 1, 0
	case 0xB7:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R15B, NoRegister, 0, false, 1, 0
		}
		return MOV, false, false, true, NoSegment, BH, NoRegister, 0, false, 1, 0
	case 0xB8:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R8, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, EAX, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RAX, NoRegister, 0, false, 4, 0
	case 0xB9:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R9, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, ECX, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RCX, NoRegister, 0, false, 4, 0
	case 0xBA:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R10, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, EDX, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RDX, NoRegister, 0, false, 4, 0
	case 0xBB:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R11, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, EBX, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RBX, NoRegister, 0, false, 4, 0
	case 0xBC:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R12, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, ESP, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RSP, NoRegister, 0, false, 4, 0
	case 0xBD:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R13, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, EBP, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RBP, NoRegister, 0, false, 4, 0
	case 0xBE:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R14, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, ESI, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RSI, NoRegister, 0, false, 4, 0
	case 0xBF:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, false, false, true, NoSegment, R15, NoRegister, 0, false, 8, 0
		}
		if isOperandSizeOveride {
			return MOV, false, false, true, NoSegment, EDI, NoRegister, 0, false, 2, 0
		}
		return MOV, false, false, true, NoSegment, RDI, NoRegister, 0, false, 4, 0
	case 0xC0, 0xC1:
		return NoInstruction, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 1, 0
	case 0xC2:
		// near
		return RET, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 2, 0
	case 0xC3:
		// near
		return RET, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xC4:
		if is64Bit {
			panic("Error: this is the vex escape prefix")
		}
		return LES, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0xC5:
		if is64Bit {
			panic("Error: this is the vex escape prefix")
		}
		return LDS, true, false, false, NoSegment, NoRegister, NoRegister, 0, true, 0, 0
	case 0xC6:
		return NoInstruction, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, 1, 0
	case 0xC7:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return NoInstruction, true, false, true, NoSegment, NoRegister, NoRegister, 0, false, noImmediateBytes, 0
	case 0xC8:
		return ENTER, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0 // todo need to figure out how many immediate bytes
	case 0xC9:
		return LEAVE, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xCA:
		// far
		return RET, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 2, 0
	case 0xCB:
		// far
		return RET, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xCC:
		return INT3, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xCD:
		return INT, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 1, 0
	case 0xCE:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return INTO, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xCF:
		if is64Bit {
			if isRexW {
				return IRETQ, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
			}
			if isOperandSizeOveride {
				return IRET, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
			}
			return IRETD, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		}
		// todo need to figure out defaut operand size
		return IRETD, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD0, 0xD1:
		return NoInstruction, true, false, true, NoSegment, NoRegister, NoRegister, 1, false, 0, 0 // todo need to figure out how many immediate bytes
	case 0xD2, 0xD3:
		return NoInstruction, true, false, true, NoSegment, NoRegister, CL, 1, false, 0, 0
	case 0xD4:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return AAM, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 1, 0
	case 0xD5:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return AAD, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 1, 0
	case 0xD6:
		return XLAT, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xD7:
		panic("todo multiple instructions in 1 byte")
	case 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF:
		panic("Error: this is x87 opcodes")
	case 0xE0:
		return LOOPNE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0xE1:
		return LOOPE, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0xE2:
		return LOOP, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0xE3:
		if is64Bit {
			if isRexW {
				return JRCXZ, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
			}
			if isOperandSizeOveride {
				return JCXZ, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
			}
			return JECXZ, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
		}
		// todo need to figure out defaut operand size
		return JECXZ, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 1, 0
	case 0xE4:
		return IN, false, false, true, NoSegment, AL, NoRegister, 0, false, 1, 0
	case 0xE5:
		return IN, false, false, true, NoSegment, EAX, NoRegister, 0, false, 1, 0
	case 0xE6:
		return OUT, false, false, true, NoSegment, NoRegister, AL, 0, false, 1, 0
	case 0xE7:
		return OUT, false, false, true, NoSegment, NoRegister, EAX, 0, false, 1, 0
	case 0xE8:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noDisplacementBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return CALL, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0xE9:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noDisplacementBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return JMP, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, noDisplacementBytes
	case 0xEA:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return JMP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0 // dont know if this is correct and need to figure out how many immediate bytes
	case 0xEB:
		return JMP, false, true, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 1
	case 0xEC:
		return IN, false, false, false, NoSegment, AL, DX, 0, false, 0, 0
	case 0xED:
		return IN, false, false, false, NoSegment, EAX, DX, 0, false, 0, 0
	case 0xEE:
		return OUT, false, false, false, NoSegment, DX, AL, 0, false, 0, 0
	case 0xEF:
		return OUT, false, false, false, NoSegment, DX, EAX, 0, false, 0, 0
	case 0xF0:
		panic("Error: this is the lock prefix")
	case 0xF1:
		return INT1, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF2:
		panic("Error: this is repne prefix")
	case 0xF3:
		panic("Error: this is rep/e prefix")
	case 0xF4:
		return HLT, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF5:
		return CMC, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF6:
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF7:
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF8:
		return CLC, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xF9:
		return STC, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFA:
		return CLI, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFB:
		return STI, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFC:
		return CLD, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFD:
		return STD, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFE:
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	case 0xFF:
		return NoInstruction, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG1(opcode byte, modrmReg [3]bool, is64bit bool) (Instruction, bool, MemSegment, Register, Register, int, int) {
	switch opcode {
	case 0x80:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, false, true}:
			return OR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, false}:
			return ADC, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, true}:
			return SBB, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, false, false}:
			return AND, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, false, true}:
			return SUB, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, true, false}:
			return XOR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, true, true}:
			return CMP, true, NoSegment, NoRegister, NoRegister, 0, 1
		default:
			panic("Error: Unknown instruction")
		}
	case 0x81:
		switch modrmReg {
		case [3]bool{false, false, false}:
			// todo need to figure out how many immediate bytes
			return ADD, true, NoSegment, NoRegister, NoRegister, 0, 0
		case [3]bool{false, false, true}:
			// todo need to figure out how many immediate bytes
			return OR, true, NoSegment, NoRegister, NoRegister, 0, 0
		case [3]bool{false, true, false}:
			// todo need to figure out how many immediate bytes
			return ADC, true, NoSegment, NoRegister, NoRegister, 0, 0
		case [3]bool{false, true, true}:
			// todo need to figure out how many immediate bytes
			return SBB, true, NoSegment, NoRegister, NoRegister, 0, 0
		case [3]bool{true, false, false}:
			// todo need to figure out how many immediate bytes
			return AND, true, NoSegment, NoRegister, NoRegister, 0, 0
		case [3]bool{true, false, true}:
			// todo need to figure out how many immediate bytes
			return SUB, true, NoSegment, NoRegister, NoRegister, 0, 0
		case [3]bool{true, true, false}:
			// todo need to figure out how many immediate bytes
			return XOR, true, NoSegment, NoRegister, NoRegister, 0, 0
		case [3]bool{true, true, true}:
			// todo need to figure out how many immediate bytes
			return CMP, true, NoSegment, NoRegister, NoRegister, 0, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0x82:
		panic("todo")
	case 0x83:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, false, true}:
			return OR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, false}:
			return ADC, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, true}:
			return SBB, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, false, false}:
			return AND, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, false, true}:
			return SUB, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, true, false}:
			return XOR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, true, true}:
			return CMP, true, NoSegment, NoRegister, NoRegister, 0, 1
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

func primaryOpcodeModRMG2(opcode byte, modrmReg [3]bool) (Instruction, bool, MemSegment, Register, Register, int, int) {
	switch opcode {
	case 0xC0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, false, true}:
			return ROR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, false}:
			return RCL, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, true}:
			return RCR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, true, NoSegment, NoRegister, NoRegister, 0, 1
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, false, true}:
			return ROR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, false}:
			return RCL, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{false, true, true}:
			return RCR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, true, NoSegment, NoRegister, NoRegister, 0, 1
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, true, NoSegment, NoRegister, NoRegister, 0, 1
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{false, false, true}:
			return ROR, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{false, true, false}:
			return RCL, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{false, true, true}:
			return RCR, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, false, NoSegment, NoRegister, NoRegister, 1, 0
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{false, false, true}:
			return ROR, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{false, true, false}:
			return RCL, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{false, true, true}:
			return RCR, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, false, NoSegment, NoRegister, NoRegister, 1, 0
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, false, NoSegment, NoRegister, NoRegister, 1, 0
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
			return INC, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, false, true}:
			return DEC, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, false}:
			return CALL, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{false, true, true}:
			return CALL, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, false}:
			return JMP, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, false, true}:
			return JMP, false, NoSegment, NoRegister, NoRegister, 0
		case [3]bool{true, true, false}:
			return PUSH, false, NoSegment, NoRegister, NoRegister, 0
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG11(opcode byte, modrmReg [3]bool, is64bit bool, isRexW bool, isOperandSizeOveride bool) (Instruction, bool, MemSegment, Register, Register, int, int) {
	switch opcode {
	case 0xC6:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return MOV, true, NoSegment, NoRegister, NoRegister, 0, 1
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC7:
		switch modrmReg {
		case [3]bool{false, false, false}:
			noImmediateBytes := 4
			if isOperandSizeOveride && !isRexW {
				noImmediateBytes = 2
			}
			return MOV, true, NoSegment, NoRegister, NoRegister, 0, noImmediateBytes
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
