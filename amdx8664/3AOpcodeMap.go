package amdx8664

func opcodeMap3A(curByte byte, isOperandSizeOverride bool) (Instruction, bool, bool) {
	switch curByte {
	case 0x14:
		if isOperandSizeOverride {
			return PEXTRB
		} else {
			return NOP
		}
	case 0x15:
		if isOperandSizeOverride {
			return PEXTRW
		} else {
			return NOP
		}
	case 0x16:
		if isOperandSizeOverride {
			return PEXTRD
		} else {
			return NOP
		}
	case 0x17:
		if isOperandSizeOverride {
			return EXTRACTPS
		} else {
			return NOP
		}
	case 0x20:
		if isOperandSizeOverride {
			return PINSRB
		} else {
			return NOP
		}
	case 0x21:
		if isOperandSizeOverride {
			return INSERTPS
		} else {
			return NOP
		}
	case 0x22:
		if isOperandSizeOverride {
			// todo
			return NOP
		} else {
			return NOP
		}
	case 0x28:
		if isOperandSizeOverride {
			return PINSRB
		} else {
			return NOP
		}
	case 0x29:
		if isOperandSizeOverride {
			return INSERTPS
		} else {
			return NOP
		}
	case 0x2A:
		if isOperandSizeOverride {
			// todo
			return NOP
		} else {
			return NOP
		}
	case 0x48:
		if isOperandSizeOverride {
			return DPPS
		} else {
			return NOP
		}
	case 0x49:
		if isOperandSizeOverride {
			return DPPD
		} else {
			return NOP
		}
	case 0x4A:
		if isOperandSizeOverride {
			return MPSADBW
		} else {
			return NOP
		}
	case 0x4C:
		if isOperandSizeOverride {
			return PCLMULQDQ
		} else {
			return NOP
		}
	}
	panic("Error: todo")
}
