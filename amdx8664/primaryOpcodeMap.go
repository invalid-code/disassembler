package amdx8664

func primaryOpcode(curByte byte, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	noImmediateBytes, noDisplacementBytes := 4, 4
	switch curByte {
	case 0x00, 0x01:
		return ADD, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x02, 0x03:
		return ADD, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x04:
		return ADD, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x05:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return ADD, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x06:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, []Operand{memSegmentOperand(ES)}
	case 0x07:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return POP, []Operand{memSegmentOperand(ES)}
	case 0x08, 0x09:
		return OR, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x0A, 0x0B:
		return OR, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x0C:
		return OR, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x0D:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return OR, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x0E:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, []Operand{memSegmentOperand(CS)}
	case 0x0F:
		panic("Error: 2-byte opcodes")
	case 0x10, 0x11:
		return ADC, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x12, 0x13:
		return ADC, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x14:
		return ADC, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x15:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return ADC, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x16:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, []Operand{memSegmentOperand(SS)}
	case 0x17:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return POP, []Operand{memSegmentOperand(SS)}
	case 0x18, 0x19:
		return SBB, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x1A, 0x1B:
		return SBB, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x1C:
		return SBB, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x1D:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return SBB, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x1E:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return PUSH, []Operand{memSegmentOperand(DS)}
	case 0x1F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return POP, []Operand{memSegmentOperand(DS)}
	case 0x20, 0x21:
		return AND, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x22, 0x23:
		return AND, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x24:
		return AND, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x25:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return AND, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x26:
		panic("Error: this is the ES segment")
	case 0x27:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return DAA, []Operand{}
	case 0x28, 0x29:
		return SUB, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x2A, 0x2B:
		return SUB, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x2C:
		return SUB, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x2D:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return SUB, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x2E:
		panic("Error: this is the CS segment")
	case 0x2F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return DAS, []Operand{}
	case 0x30, 0x31:
		return XOR, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x32, 0x33:
		return XOR, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x34:
		return XOR, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x35:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return XOR, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x36:
		panic("Error: this is the SS segment")
	case 0x37:
		if is64Bit {
			panic("Error: instruction invalid in 64-bit mode")
		}
		return AAA, []Operand{}
	case 0x38, 0x39:
		return CMP, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x3A, 0x3B:
		return CMP, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x3C:
		return CMP, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0x3D:
		// todo need to figure out operand addressing size if using eax, ax, rax
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return CMP, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0x3E:
		panic("Error: this is the DS segment")
	case 0x3F:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return AAS, []Operand{}
	case 0x40:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(EAX)}
	case 0x41:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(ECX)}
	case 0x42:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(EDX)}
	case 0x43:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(EBX)}
	case 0x44:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(ESP)}
	case 0x45:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(EBP)}
	case 0x46:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(ESI)}
	case 0x47:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return INC, []Operand{instructionEncodedRegOperand(EDI)}
	case 0x48:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, []Operand{instructionEncodedRegOperand(EAX)}
	case 0x49:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, []Operand{instructionEncodedRegOperand(ECX)}
	case 0x4A:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, []Operand{instructionEncodedRegOperand(EDX)}
	case 0x4B:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, []Operand{instructionEncodedRegOperand(EBX)}
	case 0x4C:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, []Operand{instructionEncodedRegOperand(ESP)}
	case 0x4D:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, []Operand{instructionEncodedRegOperand(EBP)}
	case 0x4E:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		} else {
			return DEC, []Operand{instructionEncodedRegOperand(ESI)}
		}
	case 0x4F:
		if is64Bit {
			panic("Error: Used as REX prefix in 64-bit mode")
		}
		return DEC, []Operand{instructionEncodedRegOperand(EDI)}
	case 0x50:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R8)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(EAX)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(RAX)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RAX)}
	case 0x51:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R9)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(ECX)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(ECX)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RCX)}
	case 0x52:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R10)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(EDX)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(RDX)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RDX)}
	case 0x53:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R11)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(EBX)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(RBX)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RBX)}
	case 0x54:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R12)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(ESP)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(RSP)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RSP)}
	case 0x55:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R13)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(EBP)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(RBP)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RBP)}
	case 0x56:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R14)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(ESI)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(RSI)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RSI)}
	case 0x57:
		if is64Bit {
			if isRexW {
				return PUSH, []Operand{instructionEncodedRegOperand(R15)}
			}
			if isOperandSizeOveride {
				return PUSH, []Operand{instructionEncodedRegOperand(RDI)}
			}
			return PUSH, []Operand{instructionEncodedRegOperand(RDI)}
		}
		// todo need to figure out defaut operand size
		return PUSH, []Operand{instructionEncodedRegOperand(RDI)}
	case 0x58:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R8)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(EAX)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RAX)}
		}
		// todo need to figure out defaut operand size
		return POP, []Operand{instructionEncodedRegOperand(RAX)}
	case 0x59:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R9)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(ECX)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RCX)}
		}
		// todo need to figure out defaut operand size
		return POP, []Operand{instructionEncodedRegOperand(RCX)}
	case 0x5A:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R10)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(EDX)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RDX)}
		}
		// todo need to figure out defaut operand size
		return POP, []Operand{instructionEncodedRegOperand(RDX)}
	case 0x5B:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R11)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(EBX)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RBX)}
		}
		// todo need to figure out defaut operand size
		return POP, []Operand{instructionEncodedRegOperand(RBX)}
	case 0x5C:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R12)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(ESP)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RSP)}
		}
		// todo need to figure out defaut operand size
		return POP, []Operand{instructionEncodedRegOperand(RSP)}
	case 0x5D:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R13)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(EBP)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RBP)}
		}
		// todo need to figure out defaut operand size
		return POP, []Operand{instructionEncodedRegOperand(RBP)}
	case 0x5E:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R14)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(ESI)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RSI)}
		}
		// todo need to figure out defaut operand size
		return POP, []Operand{instructionEncodedRegOperand(RSI)}
	case 0x5F:
		if is64Bit {
			if isRexW {
				return POP, []Operand{instructionEncodedRegOperand(R15)}
			}
			if isOperandSizeOveride {
				return POP, []Operand{instructionEncodedRegOperand(EDI)}
			}
			return POP, []Operand{instructionEncodedRegOperand(RDI)}
		}
		return POP, []Operand{instructionEncodedRegOperand(RDI)}
	case 0x60:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		if isOperandSizeOveride {
			return PUSHA, []Operand{}
		}
		return PUSHD, []Operand{}
	case 0x61:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		if isOperandSizeOveride {
			return POPA, []Operand{}
		}
		return POPD, []Operand{}
	case 0x62:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return BOUND, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x63:
		if is64Bit {
			return MOVSXD, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
		}
		return ARPL, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
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
		return PUSH, []Operand{immediateOperand(noImmediateBytes)}
	case 0x69:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return PUSH, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
	case 0x6A:
		return PUSH, []Operand{immediateOperand(1)}
	case 0x6B:
		return PUSH, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true), immediateOperand(1)}
	case 0x6C:
		panic("todo dont know how to deal with Y operand syntax notation")
	case 0x6D:
		panic("todo multiple instructions in 1 byte")
	case 0x6E:
		panic("todo dont know how to deal with X operand syntax notation")
	case 0x6F:
		panic("todo multiple instructions in 1 byte")
	case 0x70:
		return JO, []Operand{displacementOperand(1)}
	case 0x71:
		return JNO, []Operand{displacementOperand(1)}
	case 0x72:
		return JB, []Operand{displacementOperand(1)}
	case 0x73:
		return JNB, []Operand{displacementOperand(1)}
	case 0x74:
		return JZ, []Operand{displacementOperand(1)}
	case 0x75:
		return JNZ, []Operand{displacementOperand(1)}
	case 0x76:
		return JBE, []Operand{displacementOperand(1)}
	case 0x77:
		return JNBE, []Operand{displacementOperand(1)}
	case 0x78:
		return JS, []Operand{displacementOperand(1)}
	case 0x79:
		return JNS, []Operand{displacementOperand(1)}
	case 0x7A:
		return JP, []Operand{displacementOperand(1)}
	case 0x7B:
		return JNP, []Operand{displacementOperand(1)}
	case 0x7C:
		return JL, []Operand{displacementOperand(1)}
	case 0x7D:
		return JNL, []Operand{displacementOperand(1)}
	case 0x7E:
		return JLE, []Operand{displacementOperand(1)}
	case 0x7F:
		return JNLE, []Operand{displacementOperand(1)}
	case 0x80, 0x82, 0x83:
		return NoInstruction, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
	case 0x81:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return NoInstruction, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
	case 0x84, 0x85:
		return TEST, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x86, 0x87:
		return XCHG, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x88, 0x89:
		return MOV, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0x8A, 0x8B:
		return MOV, []Operand{modRMRegOperand(true, true), modRMRegOperand(false, true)}
	case 0x8C:
		panic("todo multiple operand types")
	case 0x8D:
		return LEA, []Operand{modRMRegOperand(true, true), modRMMemOperand(false, true)}
	case 0x8E:
		panic("Error: todo don't know segment registers")
	case 0x8F:
		panic("Error: todo don't know")
	case 0x90:
		panic("todo could be XCHG or NOP instruction")
	case 0x91:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, []Operand{instructionEncodedRegOperand(R9), instructionEncodedRegOperand(RAX)}
			}
			return XCHG, []Operand{instructionEncodedRegOperand(RCX), instructionEncodedRegOperand(RAX)}
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x92:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, []Operand{instructionEncodedRegOperand(R10), instructionEncodedRegOperand(RAX)}
			}
			return XCHG, []Operand{instructionEncodedRegOperand(RDX), instructionEncodedRegOperand(RAX)}
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x93:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, []Operand{instructionEncodedRegOperand(R11), instructionEncodedRegOperand(RAX)}
			}
			return XCHG, []Operand{instructionEncodedRegOperand(RBX), instructionEncodedRegOperand(RAX)}
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x94:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, []Operand{instructionEncodedRegOperand(R12), instructionEncodedRegOperand(RAX)}
			}
			return XCHG, []Operand{instructionEncodedRegOperand(RSP), instructionEncodedRegOperand(RAX)}
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x95:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, []Operand{instructionEncodedRegOperand(R13), instructionEncodedRegOperand(RAX)}
			}
			return XCHG, []Operand{instructionEncodedRegOperand(RBP), instructionEncodedRegOperand(RAX)}
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x96:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, []Operand{instructionEncodedRegOperand(R14), instructionEncodedRegOperand(RAX)}
			}
			return XCHG, []Operand{instructionEncodedRegOperand(RSI), instructionEncodedRegOperand(RAX)}
		}
		panic("todo possibly a different operand size if 32 bit")
	case 0x97:
		if is64Bit {
			if isOperandSizeOveride {
				return XCHG, []Operand{instructionEncodedRegOperand(R15), instructionEncodedRegOperand(RAX)}
			}
			return XCHG, []Operand{instructionEncodedRegOperand(RDI), instructionEncodedRegOperand(RAX)}
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
		// return CALL, false, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0
		panic("todo don't know how to disassemble far pointer operands")
	case 0x9B:
		panic("todo multiple instructions in 1 byte")
	case 0x9C:
		panic("todo multiple instructions in 1 byte")
	case 0x9D:
		panic("todo multiple instructions in 1 byte")
	case 0x9E:
		return SAHF, []Operand{}
	case 0x9F:
		return LAHF, []Operand{}
	case 0xA0:
		return MOV, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0xA1:
		if is64Bit {
			if isRexW {
				noImmediateBytes = 8
			} else if isOperandSizeOveride {
				noImmediateBytes = 4
			} else {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return MOV, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
	case 0xA2:
		return MOV, []Operand{immediateOperand(1), instructionEncodedRegOperand(AL)}
	case 0xA3:
		if is64Bit {
			if isRexW {
				noImmediateBytes = 8
			} else if isOperandSizeOveride {
				noImmediateBytes = 4
			} else {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return MOV, []Operand{immediateOperand(noImmediateBytes), instructionEncodedRegOperand(RAX)}
	case 0xA4:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xA5:
		panic("todo multiple instructions in 1 byte")
	case 0xA6:
		panic("todo dont know how to deal with X, Y operand notation")
	case 0xA7:
		panic("todo multiple instructions in 1 byte")
	case 0xA8:
		return TEST, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0xA9:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return TEST, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(noImmediateBytes)}
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
			return MOV, []Operand{instructionEncodedRegOperand(R8B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0xB1:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R9B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(CL), immediateOperand(1)}
	case 0xB2:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R10B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(DL), immediateOperand(1)}
	case 0xB3:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R11B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(BL), immediateOperand(1)}
	case 0xB4:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R12B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(AH), immediateOperand(1)}
	case 0xB5:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R13B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(CH), immediateOperand(1)}
	case 0xB6:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R14B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(DH), immediateOperand(1)}
	case 0xB7:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R15B), immediateOperand(1)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(BH), immediateOperand(1)}
	case 0xB8:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R8), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(EAX), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RAX), immediateOperand(4)}
	case 0xB9:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R9), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(ECX), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RCX), immediateOperand(4)}
	case 0xBA:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R10), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(EDX), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RDX), immediateOperand(4)}
	case 0xBB:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R11), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(EBX), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RBX), immediateOperand(4)}
	case 0xBC:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R12), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(ESP), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RSP), immediateOperand(4)}
	case 0xBD:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R13), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(EBP), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RBP), immediateOperand(4)}
	case 0xBE:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R14), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(ESI), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RSI), immediateOperand(4)}
	case 0xBF:
		if !is64Bit {
			panic("Error: applicable only in 64-bite mode")
		}
		// todo handle 32-bit
		if isRexW {
			return MOV, []Operand{instructionEncodedRegOperand(R15), immediateOperand(8)}
		}
		if isOperandSizeOveride {
			return MOV, []Operand{instructionEncodedRegOperand(EDI), immediateOperand(2)}
		}
		return MOV, []Operand{instructionEncodedRegOperand(RDI), immediateOperand(4)}
	case 0xC0, 0xC1:
		return NoInstruction, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
	case 0xC2:
		// near
		return RET, []Operand{immediateOperand(2)}
	case 0xC3:
		// near
		return RET, []Operand{}
	case 0xC4:
		if is64Bit {
			panic("Error: this is the vex escape prefix")
		}
		return LES, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0xC5:
		if is64Bit {
			panic("Error: this is the vex escape prefix")
		}
		return LDS, []Operand{modRMRegOperand(false, true), modRMRegOperand(true, true)}
	case 0xC6:
		return NoInstruction, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
	case 0xC7:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return NoInstruction, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
	case 0xC8:
		// return ENTER, false, false, true, NoSegment, NoRegister, NoRegister, 0, false, 0, 0 // todo need to figure out how many immediate bytes
		panic("todo don't know. This takes 2 immediates")
	case 0xC9:
		return LEAVE, []Operand{}
	case 0xCA:
		// far
		return RET, []Operand{immediateOperand(2)}
	case 0xCB:
		// far
		return RET, []Operand{}
	case 0xCC:
		return INT3, []Operand{}
	case 0xCD:
		return INT, []Operand{immediateOperand(1)}
	case 0xCE:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return INTO, []Operand{}
	case 0xCF:
		if is64Bit {
			if isRexW {
				return IRETQ, []Operand{}
			}
			if isOperandSizeOveride {
				return IRET, []Operand{}
			}
			return IRETD, []Operand{}
		}
		// todo need to figure out defaut operand size
		return IRETD, []Operand{}
	case 0xD0, 0xD1:
		return NoInstruction, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
	case 0xD2, 0xD3:
		return NoInstruction, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
	case 0xD4:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return AAM, []Operand{immediateOperand(1)}
	case 0xD5:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		return AAD, []Operand{immediateOperand(1)}
	case 0xD6:
		return XLAT, []Operand{}
	case 0xD7:
		panic("todo multiple instructions in 1 byte")
	case 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF:
		panic("Error: this is x87 opcodes")
	case 0xE0:
		return LOOPNE, []Operand{displacementOperand(1)}
	case 0xE1:
		return LOOPE, []Operand{displacementOperand(1)}
	case 0xE2:
		return LOOP, []Operand{displacementOperand(1)}
	case 0xE3:
		if is64Bit {
			if isRexW {
				return JRCXZ, []Operand{displacementOperand(1)}
			}
			if isOperandSizeOveride {
				return JCXZ, []Operand{displacementOperand(1)}
			}
			return JECXZ, []Operand{displacementOperand(1)}
		}
		// todo need to figure out defaut operand size
		return JECXZ, []Operand{displacementOperand(1)}
	case 0xE4:
		return IN, []Operand{instructionEncodedRegOperand(AL), immediateOperand(1)}
	case 0xE5:
		return IN, []Operand{instructionEncodedRegOperand(EAX), immediateOperand(1)}
	case 0xE6:
		return OUT, []Operand{immediateOperand(1), instructionEncodedRegOperand(AL)}
	case 0xE7:
		return OUT, []Operand{immediateOperand(1), instructionEncodedRegOperand(EAX)}
	case 0xE8:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noDisplacementBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return CALL, []Operand{displacementOperand(noDisplacementBytes)}
	case 0xE9:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noDisplacementBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		return JMP, []Operand{displacementOperand(noDisplacementBytes)}
	case 0xEA:
		if is64Bit {
			panic("Error: Invalid in 64-bit mode")
		}
		// return JMP, true, false, false, NoSegment, NoRegister, NoRegister, 0, false, 0, 0 // dont know if this is correct and need to figure out how many immediate bytes
		panic("todo don't know how to deal with far pointer operands")
	case 0xEB:
		return JMP, []Operand{displacementOperand(1)}
	case 0xEC:
		return IN, []Operand{instructionEncodedRegOperand(AL), instructionEncodedRegOperand(DX)}
	case 0xED:
		return IN, []Operand{instructionEncodedRegOperand(EAX), instructionEncodedRegOperand(DX)}
	case 0xEE:
		return OUT, []Operand{instructionEncodedRegOperand(DX), instructionEncodedRegOperand(AL)}
	case 0xEF:
		return OUT, []Operand{instructionEncodedRegOperand(DX), instructionEncodedRegOperand(EAX)}
	case 0xF0:
		panic("Error: this is the lock prefix")
	case 0xF1:
		return INT1, []Operand{}
	case 0xF2:
		panic("Error: this is repne prefix")
	case 0xF3:
		panic("Error: this is rep/e prefix")
	case 0xF4:
		return HLT, []Operand{}
	case 0xF5:
		return CMC, []Operand{}
	case 0xF6:
		return NoInstruction, []Operand{modRMRegOperand(false, true)}
	case 0xF7:
		return NoInstruction, []Operand{modRMRegOperand(false, true)}
	case 0xF8:
		return CLC, []Operand{}
	case 0xF9:
		return STC, []Operand{}
	case 0xFA:
		return CLI, []Operand{}
	case 0xFB:
		return STI, []Operand{}
	case 0xFC:
		return CLD, []Operand{}
	case 0xFD:
		return STD, []Operand{}
	case 0xFE:
		return NoInstruction, []Operand{modRMRegOperand(false, true)}
	case 0xFF:
		return NoInstruction, []Operand{}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG1(opcode byte, modrmReg [3]bool, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	noImmediateBytes := 4
	switch opcode {
	case 0x80:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, false, true}:
			return OR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, false}:
			return ADC, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, true}:
			return SBB, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, false, false}:
			return AND, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, false, true}:
			return SUB, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, false}:
			return XOR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, true}:
			return CMP, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x81:
		if is64Bit {
			if !isRexW && isOperandSizeOveride {
				noImmediateBytes = 2
			}
		} else {
			// todo need to figure out defaut operand size
		}
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{false, false, true}:
			return OR, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{false, true, false}:
			return ADC, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{false, true, true}:
			return SBB, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{true, false, false}:
			return AND, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{true, false, true}:
			return SUB, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{true, true, false}:
			return XOR, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{true, true, true}:
			return CMP, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0x82:
		panic("todo")
	case 0x83:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ADD, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, false, true}:
			return OR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, false}:
			return ADC, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, true}:
			return SBB, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, false, false}:
			return AND, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, false, true}:
			return SUB, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, false}:
			return XOR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, true}:
			return CMP, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG1a(opcode byte, modrmReg [3]bool, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0x8F:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return POP, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, false, true}, [3]bool{false, true, false}, [3]bool{false, true, true}, [3]bool{true, false, false}, [3]bool{true, false, true}, [3]bool{true, true, false}, [3]bool{true, true, true}:
			panic("todo")
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG2(opcode byte, modrmReg [3]bool, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xC0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, false, true}:
			return ROR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, false}:
			return RCL, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, true}:
			return RCR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, false, true}:
			return ROR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, false}:
			return RCL, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{false, true, true}:
			return RCR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD0:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{false, false, true}:
			return ROR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{false, true, false}:
			return RCL, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{false, true, true}:
			return RCR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD1:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{false, false, true}:
			return ROR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{false, true, false}:
			return RCL, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{false, true, true}:
			return RCR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, []Operand{modRMRegOperand(false, true), instructionEncodedIntOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD2:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{false, false, true}:
			return ROR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{false, true, false}:
			return RCL, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{false, true, true}:
			return RCR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xD3:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return ROL, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{false, false, true}:
			return ROR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{false, true, false}:
			return RCL, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{false, true, true}:
			return RCR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{true, false, false}:
			panic("todo")
		case [3]bool{true, false, true}:
			return SHR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		case [3]bool{true, true, false}:
			panic("todo")
		case [3]bool{true, true, true}:
			return SAR, []Operand{modRMRegOperand(false, true), instructionEncodedRegOperand(CL)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG3(opcode byte, modrmReg [3]bool, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	noImmediateBytes := 4
	switch opcode {
	case 0xF6:
		switch modrmReg {
		case [3]bool{false, false, false}:
			if is64Bit {
				if !isRexW && isOperandSizeOveride {
					noImmediateBytes = 2
				}
			} else {
				// todo need to figure out defaut operand size
			}
			return TEST, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{false, false, true}:
			if is64Bit {
				if !isRexW && isOperandSizeOveride {
					noImmediateBytes = 2
				}
			} else {
				// todo need to figure out defaut operand size
			}
			return TEST, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{false, true, false}:
			return NOT, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, true, true}:
			return NEG, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, false}:
			return MUL, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, true}:
			return IMUL, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, true, false}:
			return DIV, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, true, true}:
			return IDIV, []Operand{modRMRegOperand(false, true)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xF7:
		switch modrmReg {
		case [3]bool{false, false, false}:
			if is64Bit {
				if !isRexW && isOperandSizeOveride {
					noImmediateBytes = 2
				}
			} else {
				// todo need to figure out defaut operand size
			}
			return TEST, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{false, false, true}:
			if is64Bit {
				if !isRexW && isOperandSizeOveride {
					noImmediateBytes = 2
				}
			} else {
				// todo need to figure out defaut operand size
			}
			return TEST, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		case [3]bool{false, true, false}:
			return NOT, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, true, true}:
			return NEG, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, false}:
			return MUL, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, true}:
			return IMUL, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, true, false}:
			return DIV, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, true, true}:
			return IDIV, []Operand{modRMRegOperand(false, true)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG4(opcode byte, modrmReg [3]bool, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xFE:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return INC, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, false, true}:
			return DEC, []Operand{modRMRegOperand(false, true)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG5(opcode byte, modrmReg [3]bool, is64bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	switch opcode {
	case 0xFF:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return INC, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, false, true}:
			return DEC, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, true, false}:
			return CALL, []Operand{modRMRegOperand(false, true)}
		case [3]bool{false, true, true}:
			return CALL, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, false}:
			return JMP, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, false, true}:
			return JMP, []Operand{modRMRegOperand(false, true)}
		case [3]bool{true, true, false}:
			return PUSH, []Operand{modRMRegOperand(false, true)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}

func primaryOpcodeModRMG11(opcode byte, modrmReg [3]bool, is64Bit bool, isOperandSizeOveride bool, isRexW bool) (Instruction, []Operand) {
	noImmediateBytes := 4
	switch opcode {
	case 0xC6:
		switch modrmReg {
		case [3]bool{false, false, false}:
			return MOV, []Operand{modRMRegOperand(false, true), immediateOperand(1)}
		default:
			panic("Error: Unknown instruction")
		}
	case 0xC7:
		switch modrmReg {
		case [3]bool{false, false, false}:
			if is64Bit {
				if !isRexW && isOperandSizeOveride {
					noImmediateBytes = 2
				}
			} else {
				// todo need to figure out defaut operand size
			}
			return MOV, []Operand{modRMRegOperand(false, true), immediateOperand(noImmediateBytes)}
		default:
			panic("Error: Unknown instruction")
		}
	default:
		panic("Error: Unknown instruction")
	}
}
