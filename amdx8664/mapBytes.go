package amdx8664

func DisassembleBytes(data []byte, bitFormat bool) {
	isPrefix := true
	legacePrefixCnt := 0
	isOperandSizeOverride := false
	isAddressSizeOverride := false
	isCSSegmentSizeOverride := false
	isDSSegmentSizeOverride := false
	isESSegmentSizeOverride := false
	isFSSegmentSizeOverride := false
	isGSSegmentSizeOverride := false
	isSSSegmentSizeOverride := false
	isLock := false
	isRep1 := false
	isRep0 := false
	isRexPrefix := false
	isRexW := false
	isRexR := false
	isRexX := false
	isRexB := false
	isVex := false
	isXop := false
	isVex3Byte := false
	isEscapeSequence := false
	is3dNow := false
	is38 := false
	is3A := false
	isSecondaryMap := false
	fieldR := false
	fieldVvvv := [4]bool{false, false, false, false}
	fieldL := false
	fieldPp := [2]bool{false, false}
	isRXB := true
	r3B := false
	x3B := false
	b3B := false
	mapSelect := [5]bool{false, false, false, false, false}
	isModRM := false
	isOpcode := false
	isSib := false
	isImmediate := false
	isDisplacement := false
	modrmMod := [2]bool{false, false}
	modrmReg := [4]bool{false, false, false}
	modrmRM := [4]bool{false, false, false}
	sibScale := [2]bool{false, false}
	sibIndex := [3]bool{false, false, false}
	sibBase := [3]bool{false, false, false}
	instruction := AAA
	for curByte := range data {
		// 1-15 bytes
		// prefix(optional)
		if isPrefix {
			// legacy prefix (up to 4 times)
			// rex prefix (0x40-0x4F) ! only 1 in a instruction
			// legacy escape sequences
			// vex prefix
			// xop prefix
			switch curByte {
			case 0x66:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isOperandSizeOverride = true
				legacePrefixCnt += 1
			case 0x67:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isAddressSizeOverride = true
				legacePrefixCnt += 1
			case 0x2E:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isCSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x3E:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isDSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x26:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isESSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x64:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isFSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x65:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isGSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x36:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isSSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0xf0:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isLock = true
				legacePrefixCnt += 1
			case 0xf3:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isRep1 = true
				legacePrefixCnt += 1
			case 0xf2:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isRep0 = true
				legacePrefixCnt += 1
			case 0x40:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = false
				isRexB = false
			case 0x41:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = false
				isRexB = true
			case 0x42:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = true
				isRexB = false
			case 0x43:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = true
				isRexB = true
			case 0x44:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = false
				isRexB = false
			case 0x45:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = false
				isRexB = true
			case 0x46:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = true
				isRexB = false
			case 0x47:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = true
				isRexB = true
			case 0x48:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = false
				isRexB = false
			case 0x49:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = false
				isRexB = true
			case 0x4A:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = true
				isRexB = false
			case 0x4B:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = true
				isRexB = true
			case 0x4C:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = false
				isRexB = false
			case 0x4D:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = false
				isRexB = true
			case 0x4E:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = true
				isRexB = false
			case 0x4F:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isRexPrefix = true
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = true
				isRexB = true
			case 0xC5:
				isVex = true
				isVex3Byte = false
				isPrefix = false
				isEscapeSequence = true
			case 0xC4:
				isVex = true
				isVex3Byte = true
				isPrefix = false
				isEscapeSequence = true
			case 0x8F:
				isXop = true
				isPrefix = false
				isEscapeSequence = true
			default:
				isPrefix = false
				isEscapeSequence = true
			}
		}
		if isEscapeSequence {
			field := curByte
			if isVex {
				if isVex3Byte {
					if isRXB {
						r3B = field/128 == 1
						if r3B {
							field -= 128
						}
						x3B = field/64 == 1
						if x3B {
							field -= 64
						}
						b3B = field/32 == 1
						if b3B {
							field -= 32
						}
						mapSelect[0] = field/16 == 1
						if mapSelect[0] {
							field -= 16
						}
						mapSelect[1] = field/8 == 1
						if mapSelect[1] {
							field -= 8
						}
						mapSelect[2] = field/4 == 1
						if mapSelect[2] {
							field -= 4
						}
						mapSelect[3] = field/2 == 1
						if mapSelect[3] {
							field -= 2
						}
						mapSelect[4] = field == 1
						if mapSelect[4] {
							field -= 1
						}
						isRXB = false
					} else {
						fieldR = field/128 == 1
						if fieldR {
							field -= 128
						}
						fieldVvvv[0] = field/64 == 1
						if fieldVvvv[0] {
							field -= 64
						}
						fieldVvvv[1] = field/32 == 1
						if fieldVvvv[1] {
							field -= 32
						}
						fieldVvvv[2] = field/16 == 1
						if fieldVvvv[2] {
							field -= 16
						}
						fieldVvvv[3] = field/8 == 1
						if fieldVvvv[3] {
							field -= 8
						}
						fieldL = field/4 == 1
						if fieldL {
							field -= 4
						}
						fieldPp[0] = field/2 == 1
						if fieldPp[0] {
							field -= 2
						}
						fieldPp[1] = field == 1
						if fieldPp[1] {
							field -= 1
						}
						isEscapeSequence = false
						isOpcode = true
					}
				} else {
					fieldR = field/128 == 1
					if fieldR {
						field -= 128
					}
					fieldVvvv[0] = field/64 == 1
					if fieldVvvv[0] {
						field -= 64
					}
					fieldVvvv[1] = field/32 == 1
					if fieldVvvv[1] {
						field -= 32
					}
					fieldVvvv[2] = field/16 == 1
					if fieldVvvv[2] {
						field -= 16
					}
					fieldVvvv[3] = field/8 == 1
					if fieldVvvv[3] {
						field -= 8
					}
					fieldL = field/4 == 1
					if fieldL {
						field -= 4
					}
					fieldPp[0] = field/2 == 1
					if fieldPp[0] {
						field -= 2
					}
					fieldPp[1] = field == 1
					if fieldPp[1] {
						field -= 1
					}
					isEscapeSequence = false
					isOpcode = true
				}
			} else if isXop {
				if isRXB {
					r3B = field/128 == 1
					if r3B {
						field -= 128
					}
					x3B = field/64 == 1
					if x3B {
						field -= 64
					}
					b3B = field/32 == 1
					if b3B {
						field -= 32
					}
					mapSelect[0] = field/16 == 1
					if mapSelect[0] {
						field -= 16
					}
					mapSelect[1] = field/8 == 1
					if mapSelect[1] {
						field -= 8
					}
					mapSelect[2] = field/4 == 1
					if mapSelect[2] {
						field -= 4
					}
					mapSelect[3] = field/2 == 1
					if mapSelect[3] {
						field -= 2
					}
					mapSelect[4] = field == 1
					if mapSelect[4] {
						field -= 1
					}
					isRXB = false
				} else {
					fieldR = field/128 == 1
					if fieldR {
						field -= 128
					}
					fieldVvvv[0] = field/64 == 1
					if fieldVvvv[0] {
						field -= 64
					}
					fieldVvvv[1] = field/32 == 1
					if fieldVvvv[1] {
						field -= 32
					}
					fieldVvvv[2] = field/16 == 1
					if fieldVvvv[2] {
						field -= 16
					}
					fieldVvvv[3] = field/8 == 1
					if fieldVvvv[3] {
						field -= 8
					}
					fieldL = field/4 == 1
					if fieldL {
						field -= 4
					}
					fieldPp[0] = field/2 == 1
					if fieldPp[0] {
						field -= 2
					}
					fieldPp[1] = field == 1
					if fieldPp[1] {
						field -= 1
					}
					isEscapeSequence = false
					isOpcode = true
				}
			} else {
				switch curByte {
				case 0xF:
					if isSecondaryMap {
						is3dNow = true
						isSecondaryMap = false
						isEscapeSequence = false
						isOpcode = true
					} else {
						isSecondaryMap = true
					}
				case 0x38:
					if isSecondaryMap {
						is38 = true
						isEscapeSequence = false
						isOpcode = true
					}
				case 0x3A:
					if isSecondaryMap {
						is3A = true
						isEscapeSequence = false
						isOpcode = true
					}
				default:
					isEscapeSequence = false
					isOpcode = true
				}
			}
		}
		// opcode
		if isOpcode {
			if isVex {
				if isVex3Byte {
					if mapSelect[3] == false && mapSelect[4] == true {
						// 0x01
						switch curByte {
						case 0x0:
							instruction = ADD
						case 0x1:
							instruction = ADD
						case 0x2:
							instruction = ADD
						case 0x3:
							instruction = ADD
						case 0x4:
							instruction = ADD
						case 0x5:
							instruction = ADD
						case 0x6:
							instruction = PUSH
						case 0x7:
							instruction = POP
						case 0x8:
							instruction = OR
						case 0x9:
							instruction = OR
						case 0xA:
							instruction = OR
						case 0xB:
							instruction = OR
						case 0xC:
							instruction = OR
						case 0xD:
							instruction = OR
						case 0xE:
							instruction = PUSH
						case 0xF:
							// escape to 2nd opcode map
						case 0x10:
							instruction = ADC
						case 0x11:
							instruction = ADC
						case 0x12:
							instruction = ADC
						case 0x13:
							instruction = ADC
						case 0x14:
							instruction = ADC
						case 0x15:
							instruction = ADC
						case 0x16:
							instruction = PUSH
						case 0x17:
							instruction = POP
						case 0x18:
							instruction = SBB
						case 0x19:
							instruction = SBB
						case 0x1A:
							instruction = SBB
						case 0x1B:
							instruction = SBB
						case 0x1C:
							instruction = SBB
						case 0x1D:
							instruction = SBB
						case 0x1E:
							instruction = PUSH
						case 0x1F:
							instruction = POP
						case 0x20:
							instruction = AND
						case 0x21:
							instruction = AND
						case 0x22:
							instruction = AND
						case 0x23:
							instruction = AND
						case 0x24:
							instruction = AND
						case 0x25:
							instruction = AND
						case 0x26:
							// instruction = ADD
						case 0x27:
							instruction = DAA
						case 0x28:
							instruction = SUB
						case 0x29:
							instruction = SUB
						case 0x2A:
							instruction = SUB
						case 0x2B:
							instruction = SUB
						case 0x2C:
							instruction = SUB
						case 0x2D:
							instruction = SUB
						case 0x2E:
							// instruction = ADD
						case 0x2F:
							instruction = DAS
						case 0x30:
							instruction = XOR
						case 0x31:
							instruction = XOR
						case 0x32:
							instruction = XOR
						case 0x33:
							instruction = XOR
						case 0x34:
							instruction = XOR
						case 0x35:
							instruction = XOR
						case 0x36:
							// instruction = ADD
						case 0x37:
							instruction = AAA
						case 0x38:
							instruction = CMP
						case 0x39:
							instruction = CMP
						case 0x3A:
							instruction = CMP
						case 0x3B:
							instruction = CMP
						case 0x3C:
							instruction = CMP
						case 0x3D:
							instruction = CMP
						case 0x3E:
							// instruction = ADD
						case 0x3F:
							instruction = AAS
						case 0x40:
							// instruction = ADD
						case 0x41:
							// instruction = ADD
						case 0x42:
							// instruction = ADD
						case 0x43:
							// instruction = ADD
						case 0x44:
							// instruction = ADD
						case 0x45:
							// instruction = ADD
						case 0x46:
							// instruction = ADD
						case 0x47:
							// instruction = ADD
						case 0x48:
							// instruction = ADD
						case 0x49:
							// instruction = ADD
						case 0x4A:
							// instruction = ADD
						case 0x4B:
							// instruction = ADD
						case 0x4C:
							// instruction = ADD
						case 0x4D:
							// instruction = ADD
						case 0x4E:
							// instruction = ADD
						case 0x4F:
							// instruction = ADD
						case 0x50:
							instruction = PUSH
						case 0x51:
							instruction = PUSH
						case 0x52:
							instruction = PUSH
						case 0x53:
							instruction = PUSH
						case 0x54:
							instruction = PUSH
						case 0x55:
							instruction = PUSH
						case 0x56:
							instruction = PUSH
						case 0x57:
							instruction = PUSH
						case 0x58:
							instruction = POP
						case 0x59:
							instruction = POP
						case 0x5A:
							instruction = POP
						case 0x5B:
							instruction = POP
						case 0x5C:
							instruction = POP
						case 0x5D:
							instruction = POP
						case 0x5E:
							instruction = POP
						case 0x5F:
							instruction = POP
						case 0x60:
							// instruction = ADD operand size override
						case 0x61:
							// instruction = ADD operand size override
						case 0x62:
							instruction = BOUND
						case 0x63:
							instruction = ADD
						case 0x64:
							// instruction = ADD
						case 0x65:
							// instruction = ADD
						case 0x66:
							// instruction = ADD
						case 0x67:
							// instruction = ADD
						case 0x68:
							instruction = PUSH
						case 0x69:
							instruction = IMUL
						case 0x6A:
							instruction = PUSH
						case 0x6B:
							instruction = IMUL
						case 0x6C:
							// instruction = ADD
						case 0x6D:
							// instruction = ADD
						case 0x6E:
							// instruction = ADD
						case 0x6F:
							// instruction = ADD
						case 0x70:
							instruction = ADD
						case 0x71:
							instruction = ADD
						case 0x72:
							instruction = ADD
						case 0x73:
							instruction = ADD
						case 0x74:
							instruction = ADD
						case 0x75:
							instruction = ADD
						case 0x76:
							instruction = ADD
						case 0x77:
							instruction = ADD
						case 0x78:
							instruction = ADD
						case 0x79:
							instruction = ADD
						case 0x7A:
							instruction = ADD
						case 0x7B:
							instruction = ADD
						case 0x7C:
							instruction = ADD
						case 0x7D:
							instruction = ADD
						case 0x7E:
							instruction = ADD
						case 0x7F:
							instruction = ADD
						case 0x80:
							instruction = ADD
						case 0x81:
							instruction = ADD
						case 0x82:
							instruction = ADD
						case 0x83:
							instruction = ADD
						case 0x84:
							instruction = ADD
						case 0x85:
							instruction = ADD
						case 0x86:
							instruction = ADD
						case 0x87:
							instruction = ADD
						case 0x88:
							instruction = ADD
						case 0x89:
							instruction = ADD
						case 0x8A:
							instruction = ADD
						case 0x8B:
							instruction = ADD
						case 0x8C:
							instruction = ADD
						case 0x8D:
							instruction = ADD
						case 0x8E:
							instruction = ADD
						case 0x8F:
							instruction = ADD
						case 0x90:
							instruction = ADD
						case 0x91:
							instruction = ADD
						case 0x92:
							instruction = ADD
						case 0x93:
							instruction = ADD
						case 0x94:
							instruction = ADD
						case 0x95:
							instruction = ADD
						case 0x96:
							instruction = ADD
						case 0x97:
							instruction = ADD
						case 0x98:
							instruction = ADD
						case 0x99:
							instruction = ADD
						case 0x9A:
							instruction = ADD
						case 0x9B:
							instruction = ADD
						case 0x9C:
							instruction = ADD
						case 0x9D:
							instruction = ADD
						case 0x9E:
							instruction = ADD
						case 0x9F:
							instruction = ADD
						case 0xA0:
							instruction = ADD
						case 0xA1:
							instruction = ADD
						case 0xA2:
							instruction = ADD
						case 0xA3:
							instruction = ADD
						case 0xA4:
							instruction = ADD
						case 0xA5:
							instruction = ADD
						case 0xA6:
							instruction = ADD
						case 0xA7:
							instruction = ADD
						case 0xA8:
							instruction = ADD
						case 0xA9:
							instruction = ADD
						case 0xAA:
							instruction = ADD
						case 0xAB:
							instruction = ADD
						case 0xAC:
							instruction = ADD
						case 0xAD:
							instruction = ADD
						case 0xAE:
							instruction = ADD
						case 0xAF:
							instruction = ADD
						case 0xB0:
							instruction = ADD
						case 0xB1:
							instruction = ADD
						case 0xB2:
							instruction = ADD
						case 0xB3:
							instruction = ADD
						case 0xB4:
							instruction = ADD
						case 0xB5:
							instruction = ADD
						case 0xB6:
							instruction = ADD
						case 0xB7:
							instruction = ADD
						case 0xB8:
							instruction = ADD
						case 0xB9:
							instruction = ADD
						case 0xBA:
							instruction = ADD
						case 0xBB:
							instruction = ADD
						case 0xBC:
							instruction = ADD
						case 0xBD:
							instruction = ADD
						case 0xBE:
							instruction = ADD
						case 0xBF:
							instruction = ADD
						case 0xC0:
							instruction = ADD
						case 0xC1:
							instruction = ADD
						case 0xC2:
							instruction = ADD
						case 0xC3:
							instruction = ADD
						case 0xC4:
							instruction = ADD
						case 0xC5:
							instruction = ADD
						case 0xC6:
							instruction = ADD
						case 0xC7:
							instruction = ADD
						case 0xC8:
							instruction = ADD
						case 0xC9:
							instruction = ADD
						case 0xCA:
							instruction = ADD
						case 0xCB:
							instruction = ADD
						case 0xCC:
							instruction = ADD
						case 0xCD:
							instruction = ADD
						case 0xCE:
							instruction = ADD
						case 0xCF:
							instruction = ADD
						case 0xD0:
							instruction = ADD
						case 0xD1:
							instruction = ADD
						case 0xD2:
							instruction = ADD
						case 0xD3:
							instruction = ADD
						case 0xD4:
							instruction = ADD
						case 0xD5:
							instruction = ADD
						case 0xD6:
							instruction = ADD
						case 0xD7:
							instruction = ADD
						case 0xD8:
							instruction = ADD
						case 0xD9:
							instruction = ADD
						case 0xDA:
							instruction = ADD
						case 0xDB:
							instruction = ADD
						case 0xDC:
							instruction = ADD
						case 0xDD:
							instruction = ADD
						case 0xDE:
							instruction = ADD
						case 0xDF:
							instruction = ADD
						case 0xE0:
							instruction = ADD
						case 0xE1:
							instruction = ADD
						case 0xE2:
							instruction = ADD
						case 0xE3:
							instruction = ADD
						case 0xE4:
							instruction = ADD
						case 0xE5:
							instruction = ADD
						case 0xE6:
							instruction = ADD
						case 0xE7:
							instruction = ADD
						case 0xE8:
							instruction = ADD
						case 0xE9:
							instruction = ADD
						case 0xEA:
							instruction = ADD
						case 0xEB:
							instruction = ADD
						case 0xEC:
							instruction = ADD
						case 0xED:
							instruction = ADD
						case 0xEE:
							instruction = ADD
						case 0xEF:
							instruction = ADD
						case 0xF0:
							instruction = ADD
						case 0xF1:
							instruction = ADD
						case 0xF2:
							instruction = ADD
						case 0xF3:
							instruction = ADD
						case 0xF4:
							instruction = ADD
						case 0xF5:
							instruction = ADD
						case 0xF6:
							instruction = ADD
						case 0xF7:
							instruction = ADD
						case 0xF8:
							instruction = ADD
						case 0xF9:
							instruction = ADD
						case 0xFA:
							instruction = ADD
						case 0xFB:
							instruction = ADD
						case 0xFC:
							instruction = ADD
						case 0xFD:
							instruction = ADD
						case 0xFE:
							instruction = ADD
						case 0xFF:
							instruction = ADD
						}
					} else if mapSelect[3] == true && mapSelect[4] == false {
						// 0x10
						switch curByte {
						case 0x0:
							instruction = ADD
						case 0x1:
							instruction = ADD
						case 0x2:
							instruction = ADD
						case 0x3:
							instruction = ADD
						case 0x4:
							instruction = ADD
						case 0x5:
							instruction = ADD
						case 0x6:
							instruction = PUSH
						case 0x7:
							instruction = POP
						case 0x8:
							instruction = OR
						case 0x9:
							instruction = OR
						case 0xA:
							instruction = OR
						case 0xB:
							instruction = OR
						case 0xC:
							instruction = OR
						case 0xD:
							instruction = OR
						case 0xE:
							instruction = PUSH
						case 0xF:
							// escape to 2nd opcode map
						case 0x10:
							instruction = ADC
						case 0x11:
							instruction = ADC
						case 0x12:
							instruction = ADC
						case 0x13:
							instruction = ADC
						case 0x14:
							instruction = ADC
						case 0x15:
							instruction = ADC
						case 0x16:
							instruction = PUSH
						case 0x17:
							instruction = POP
						case 0x18:
							instruction = SBB
						case 0x19:
							instruction = SBB
						case 0x1A:
							instruction = SBB
						case 0x1B:
							instruction = SBB
						case 0x1C:
							instruction = SBB
						case 0x1D:
							instruction = SBB
						case 0x1E:
							instruction = PUSH
						case 0x1F:
							instruction = POP
						case 0x20:
							instruction = AND
						case 0x21:
							instruction = AND
						case 0x22:
							instruction = AND
						case 0x23:
							instruction = AND
						case 0x24:
							instruction = AND
						case 0x25:
							instruction = AND
						case 0x26:
							// instruction = ADD
						case 0x27:
							instruction = DAA
						case 0x28:
							instruction = SUB
						case 0x29:
							instruction = SUB
						case 0x2A:
							instruction = SUB
						case 0x2B:
							instruction = SUB
						case 0x2C:
							instruction = SUB
						case 0x2D:
							instruction = SUB
						case 0x2E:
							// instruction = ADD
						case 0x2F:
							instruction = DAS
						case 0x30:
							instruction = XOR
						case 0x31:
							instruction = XOR
						case 0x32:
							instruction = XOR
						case 0x33:
							instruction = XOR
						case 0x34:
							instruction = XOR
						case 0x35:
							instruction = XOR
						case 0x36:
							// instruction = ADD
						case 0x37:
							instruction = AAA
						case 0x38:
							instruction = CMP
						case 0x39:
							instruction = CMP
						case 0x3A:
							instruction = CMP
						case 0x3B:
							instruction = CMP
						case 0x3C:
							instruction = CMP
						case 0x3D:
							instruction = CMP
						case 0x3E:
							// instruction = ADD
						case 0x3F:
							instruction = AAS
						case 0x40:
							// instruction = ADD
						case 0x41:
							// instruction = ADD
						case 0x42:
							// instruction = ADD
						case 0x43:
							// instruction = ADD
						case 0x44:
							// instruction = ADD
						case 0x45:
							// instruction = ADD
						case 0x46:
							// instruction = ADD
						case 0x47:
							// instruction = ADD
						case 0x48:
							// instruction = ADD
						case 0x49:
							// instruction = ADD
						case 0x4A:
							// instruction = ADD
						case 0x4B:
							// instruction = ADD
						case 0x4C:
							// instruction = ADD
						case 0x4D:
							// instruction = ADD
						case 0x4E:
							// instruction = ADD
						case 0x4F:
							// instruction = ADD
						case 0x50:
							instruction = PUSH
						case 0x51:
							instruction = PUSH
						case 0x52:
							instruction = PUSH
						case 0x53:
							instruction = PUSH
						case 0x54:
							instruction = PUSH
						case 0x55:
							instruction = PUSH
						case 0x56:
							instruction = PUSH
						case 0x57:
							instruction = PUSH
						case 0x58:
							instruction = POP
						case 0x59:
							instruction = POP
						case 0x5A:
							instruction = POP
						case 0x5B:
							instruction = POP
						case 0x5C:
							instruction = POP
						case 0x5D:
							instruction = POP
						case 0x5E:
							instruction = POP
						case 0x5F:
							instruction = POP
						case 0x60:
							// instruction = ADD operand size override
						case 0x61:
							// instruction = ADD operand size override
						case 0x62:
							instruction = BOUND
						case 0x63:
							instruction = ADD
						case 0x64:
							instruction = ADD
						case 0x65:
							instruction = ADD
						case 0x66:
							instruction = ADD
						case 0x67:
							instruction = ADD
						case 0x68:
							instruction = ADD
						case 0x69:
							instruction = ADD
						case 0x6A:
							instruction = ADD
						case 0x6B:
							instruction = ADD
						case 0x6C:
							instruction = ADD
						case 0x6D:
							instruction = ADD
						case 0x6E:
							instruction = ADD
						case 0x6F:
							instruction = ADD
						case 0x70:
							instruction = ADD
						case 0x71:
							instruction = ADD
						case 0x72:
							instruction = ADD
						case 0x73:
							instruction = ADD
						case 0x74:
							instruction = ADD
						case 0x75:
							instruction = ADD
						case 0x76:
							instruction = ADD
						case 0x77:
							instruction = ADD
						case 0x78:
							instruction = ADD
						case 0x79:
							instruction = ADD
						case 0x7A:
							instruction = ADD
						case 0x7B:
							instruction = ADD
						case 0x7C:
							instruction = ADD
						case 0x7D:
							instruction = ADD
						case 0x7E:
							instruction = ADD
						case 0x7F:
							instruction = ADD
						case 0x80:
							instruction = ADD
						case 0x81:
							instruction = ADD
						case 0x82:
							instruction = ADD
						case 0x83:
							instruction = ADD
						case 0x84:
							instruction = ADD
						case 0x85:
							instruction = ADD
						case 0x86:
							instruction = ADD
						case 0x87:
							instruction = ADD
						case 0x88:
							instruction = ADD
						case 0x89:
							instruction = ADD
						case 0x8A:
							instruction = ADD
						case 0x8B:
							instruction = ADD
						case 0x8C:
							instruction = ADD
						case 0x8D:
							instruction = ADD
						case 0x8E:
							instruction = ADD
						case 0x8F:
							instruction = ADD
						case 0x90:
							instruction = ADD
						case 0x91:
							instruction = ADD
						case 0x92:
							instruction = ADD
						case 0x93:
							instruction = ADD
						case 0x94:
							instruction = ADD
						case 0x95:
							instruction = ADD
						case 0x96:
							instruction = ADD
						case 0x97:
							instruction = ADD
						case 0x98:
							instruction = ADD
						case 0x99:
							instruction = ADD
						case 0x9A:
							instruction = ADD
						case 0x9B:
							instruction = ADD
						case 0x9C:
							instruction = ADD
						case 0x9D:
							instruction = ADD
						case 0x9E:
							instruction = ADD
						case 0x9F:
							instruction = ADD
						case 0xA0:
							instruction = ADD
						case 0xA1:
							instruction = ADD
						case 0xA2:
							instruction = ADD
						case 0xA3:
							instruction = ADD
						case 0xA4:
							instruction = ADD
						case 0xA5:
							instruction = ADD
						case 0xA6:
							instruction = ADD
						case 0xA7:
							instruction = ADD
						case 0xA8:
							instruction = ADD
						case 0xA9:
							instruction = ADD
						case 0xAA:
							instruction = ADD
						case 0xAB:
							instruction = ADD
						case 0xAC:
							instruction = ADD
						case 0xAD:
							instruction = ADD
						case 0xAE:
							instruction = ADD
						case 0xAF:
							instruction = ADD
						case 0xB0:
							instruction = ADD
						case 0xB1:
							instruction = ADD
						case 0xB2:
							instruction = ADD
						case 0xB3:
							instruction = ADD
						case 0xB4:
							instruction = ADD
						case 0xB5:
							instruction = ADD
						case 0xB6:
							instruction = ADD
						case 0xB7:
							instruction = ADD
						case 0xB8:
							instruction = ADD
						case 0xB9:
							instruction = ADD
						case 0xBA:
							instruction = ADD
						case 0xBB:
							instruction = ADD
						case 0xBC:
							instruction = ADD
						case 0xBD:
							instruction = ADD
						case 0xBE:
							instruction = ADD
						case 0xBF:
							instruction = ADD
						case 0xC0:
							instruction = ADD
						case 0xC1:
							instruction = ADD
						case 0xC2:
							instruction = ADD
						case 0xC3:
							instruction = ADD
						case 0xC4:
							instruction = ADD
						case 0xC5:
							instruction = ADD
						case 0xC6:
							instruction = ADD
						case 0xC7:
							instruction = ADD
						case 0xC8:
							instruction = ADD
						case 0xC9:
							instruction = ADD
						case 0xCA:
							instruction = ADD
						case 0xCB:
							instruction = ADD
						case 0xCC:
							instruction = ADD
						case 0xCD:
							instruction = ADD
						case 0xCE:
							instruction = ADD
						case 0xCF:
							instruction = ADD
						case 0xD0:
							instruction = ADD
						case 0xD1:
							instruction = ADD
						case 0xD2:
							instruction = ADD
						case 0xD3:
							instruction = ADD
						case 0xD4:
							instruction = ADD
						case 0xD5:
							instruction = ADD
						case 0xD6:
							instruction = ADD
						case 0xD7:
							instruction = ADD
						case 0xD8:
							instruction = ADD
						case 0xD9:
							instruction = ADD
						case 0xDA:
							instruction = ADD
						case 0xDB:
							instruction = ADD
						case 0xDC:
							instruction = ADD
						case 0xDD:
							instruction = ADD
						case 0xDE:
							instruction = ADD
						case 0xDF:
							instruction = ADD
						case 0xE0:
							instruction = ADD
						case 0xE1:
							instruction = ADD
						case 0xE2:
							instruction = ADD
						case 0xE3:
							instruction = ADD
						case 0xE4:
							instruction = ADD
						case 0xE5:
							instruction = ADD
						case 0xE6:
							instruction = ADD
						case 0xE7:
							instruction = ADD
						case 0xE8:
							instruction = ADD
						case 0xE9:
							instruction = ADD
						case 0xEA:
							instruction = ADD
						case 0xEB:
							instruction = ADD
						case 0xEC:
							instruction = ADD
						case 0xED:
							instruction = ADD
						case 0xEE:
							instruction = ADD
						case 0xEF:
							instruction = ADD
						case 0xF0:
							instruction = ADD
						case 0xF1:
							instruction = ADD
						case 0xF2:
							instruction = ADD
						case 0xF3:
							instruction = ADD
						case 0xF4:
							instruction = ADD
						case 0xF5:
							instruction = ADD
						case 0xF6:
							instruction = ADD
						case 0xF7:
							instruction = ADD
						case 0xF8:
							instruction = ADD
						case 0xF9:
							instruction = ADD
						case 0xFA:
							instruction = ADD
						case 0xFB:
							instruction = ADD
						case 0xFC:
							instruction = ADD
						case 0xFD:
							instruction = ADD
						case 0xFE:
							instruction = ADD
						case 0xFF:
							instruction = ADD
						}
					} else if mapSelect[3] == true && mapSelect[4] == true {
						// 0x11
						switch curByte {
						case 0x0:
							instruction = ADD
						case 0x1:
							instruction = ADD
						case 0x2:
							instruction = ADD
						case 0x3:
							instruction = ADD
						case 0x4:
							instruction = ADD
						case 0x5:
							instruction = ADD
						case 0x6:
							instruction = PUSH
						case 0x7:
							instruction = POP
						case 0x8:
							instruction = OR
						case 0x9:
							instruction = OR
						case 0xA:
							instruction = OR
						case 0xB:
							instruction = OR
						case 0xC:
							instruction = OR
						case 0xD:
							instruction = OR
						case 0xE:
							instruction = PUSH
						case 0xF:
							// escape to 2nd opcode map
						case 0x10:
							instruction = ADC
						case 0x11:
							instruction = ADC
						case 0x12:
							instruction = ADC
						case 0x13:
							instruction = ADC
						case 0x14:
							instruction = ADC
						case 0x15:
							instruction = ADC
						case 0x16:
							instruction = PUSH
						case 0x17:
							instruction = POP
						case 0x18:
							instruction = SBB
						case 0x19:
							instruction = SBB
						case 0x1A:
							instruction = SBB
						case 0x1B:
							instruction = SBB
						case 0x1C:
							instruction = SBB
						case 0x1D:
							instruction = SBB
						case 0x1E:
							instruction = PUSH
						case 0x1F:
							instruction = POP
						case 0x20:
							instruction = AND
						case 0x21:
							instruction = AND
						case 0x22:
							instruction = AND
						case 0x23:
							instruction = AND
						case 0x24:
							instruction = AND
						case 0x25:
							instruction = AND
						case 0x26:
							// instruction = ADD
						case 0x27:
							instruction = DAA
						case 0x28:
							instruction = SUB
						case 0x29:
							instruction = SUB
						case 0x2A:
							instruction = SUB
						case 0x2B:
							instruction = SUB
						case 0x2C:
							instruction = SUB
						case 0x2D:
							instruction = SUB
						case 0x2E:
							// instruction = ADD
						case 0x2F:
							instruction = DAS
						case 0x30:
							instruction = XOR
						case 0x31:
							instruction = XOR
						case 0x32:
							instruction = XOR
						case 0x33:
							instruction = XOR
						case 0x34:
							instruction = XOR
						case 0x35:
							instruction = XOR
						case 0x36:
							// instruction = ADD
						case 0x37:
							instruction = AAA
						case 0x38:
							instruction = CMP
						case 0x39:
							instruction = CMP
						case 0x3A:
							instruction = CMP
						case 0x3B:
							instruction = CMP
						case 0x3C:
							instruction = CMP
						case 0x3D:
							instruction = CMP
						case 0x3E:
							// instruction = ADD
						case 0x3F:
							instruction = AAS
						case 0x40:
							// instruction = ADD
						case 0x41:
							// instruction = ADD
						case 0x42:
							// instruction = ADD
						case 0x43:
							// instruction = ADD
						case 0x44:
							// instruction = ADD
						case 0x45:
							// instruction = ADD
						case 0x46:
							// instruction = ADD
						case 0x47:
							// instruction = ADD
						case 0x48:
							// instruction = ADD
						case 0x49:
							// instruction = ADD
						case 0x4A:
							// instruction = ADD
						case 0x4B:
							// instruction = ADD
						case 0x4C:
							// instruction = ADD
						case 0x4D:
							// instruction = ADD
						case 0x4E:
							// instruction = ADD
						case 0x4F:
							// instruction = ADD
						case 0x50:
							instruction = PUSH
						case 0x51:
							instruction = PUSH
						case 0x52:
							instruction = PUSH
						case 0x53:
							instruction = PUSH
						case 0x54:
							instruction = PUSH
						case 0x55:
							instruction = PUSH
						case 0x56:
							instruction = PUSH
						case 0x57:
							instruction = PUSH
						case 0x58:
							instruction = POP
						case 0x59:
							instruction = POP
						case 0x5A:
							instruction = POP
						case 0x5B:
							instruction = POP
						case 0x5C:
							instruction = POP
						case 0x5D:
							instruction = POP
						case 0x5E:
							instruction = POP
						case 0x5F:
							instruction = POP
						case 0x60:
							// instruction = ADD operand size override
						case 0x61:
							// instruction = ADD operand size override
						case 0x62:
							instruction = BOUND
						case 0x63:
							instruction = ADD
						case 0x64:
							instruction = ADD
						case 0x65:
							instruction = ADD
						case 0x66:
							instruction = ADD
						case 0x67:
							instruction = ADD
						case 0x68:
							instruction = ADD
						case 0x69:
							instruction = ADD
						case 0x6A:
							instruction = ADD
						case 0x6B:
							instruction = ADD
						case 0x6C:
							instruction = ADD
						case 0x6D:
							instruction = ADD
						case 0x6E:
							instruction = ADD
						case 0x6F:
							instruction = ADD
						case 0x70:
							instruction = ADD
						case 0x71:
							instruction = ADD
						case 0x72:
							instruction = ADD
						case 0x73:
							instruction = ADD
						case 0x74:
							instruction = ADD
						case 0x75:
							instruction = ADD
						case 0x76:
							instruction = ADD
						case 0x77:
							instruction = ADD
						case 0x78:
							instruction = ADD
						case 0x79:
							instruction = ADD
						case 0x7A:
							instruction = ADD
						case 0x7B:
							instruction = ADD
						case 0x7C:
							instruction = ADD
						case 0x7D:
							instruction = ADD
						case 0x7E:
							instruction = ADD
						case 0x7F:
							instruction = ADD
						case 0x80:
							instruction = ADD
						case 0x81:
							instruction = ADD
						case 0x82:
							instruction = ADD
						case 0x83:
							instruction = ADD
						case 0x84:
							instruction = ADD
						case 0x85:
							instruction = ADD
						case 0x86:
							instruction = ADD
						case 0x87:
							instruction = ADD
						case 0x88:
							instruction = ADD
						case 0x89:
							instruction = ADD
						case 0x8A:
							instruction = ADD
						case 0x8B:
							instruction = ADD
						case 0x8C:
							instruction = ADD
						case 0x8D:
							instruction = ADD
						case 0x8E:
							instruction = ADD
						case 0x8F:
							instruction = ADD
						case 0x90:
							instruction = ADD
						case 0x91:
							instruction = ADD
						case 0x92:
							instruction = ADD
						case 0x93:
							instruction = ADD
						case 0x94:
							instruction = ADD
						case 0x95:
							instruction = ADD
						case 0x96:
							instruction = ADD
						case 0x97:
							instruction = ADD
						case 0x98:
							instruction = ADD
						case 0x99:
							instruction = ADD
						case 0x9A:
							instruction = ADD
						case 0x9B:
							instruction = ADD
						case 0x9C:
							instruction = ADD
						case 0x9D:
							instruction = ADD
						case 0x9E:
							instruction = ADD
						case 0x9F:
							instruction = ADD
						case 0xA0:
							instruction = ADD
						case 0xA1:
							instruction = ADD
						case 0xA2:
							instruction = ADD
						case 0xA3:
							instruction = ADD
						case 0xA4:
							instruction = ADD
						case 0xA5:
							instruction = ADD
						case 0xA6:
							instruction = ADD
						case 0xA7:
							instruction = ADD
						case 0xA8:
							instruction = ADD
						case 0xA9:
							instruction = ADD
						case 0xAA:
							instruction = ADD
						case 0xAB:
							instruction = ADD
						case 0xAC:
							instruction = ADD
						case 0xAD:
							instruction = ADD
						case 0xAE:
							instruction = ADD
						case 0xAF:
							instruction = ADD
						case 0xB0:
							instruction = ADD
						case 0xB1:
							instruction = ADD
						case 0xB2:
							instruction = ADD
						case 0xB3:
							instruction = ADD
						case 0xB4:
							instruction = ADD
						case 0xB5:
							instruction = ADD
						case 0xB6:
							instruction = ADD
						case 0xB7:
							instruction = ADD
						case 0xB8:
							instruction = ADD
						case 0xB9:
							instruction = ADD
						case 0xBA:
							instruction = ADD
						case 0xBB:
							instruction = ADD
						case 0xBC:
							instruction = ADD
						case 0xBD:
							instruction = ADD
						case 0xBE:
							instruction = ADD
						case 0xBF:
							instruction = ADD
						case 0xC0:
							instruction = ADD
						case 0xC1:
							instruction = ADD
						case 0xC2:
							instruction = ADD
						case 0xC3:
							instruction = ADD
						case 0xC4:
							instruction = ADD
						case 0xC5:
							instruction = ADD
						case 0xC6:
							instruction = ADD
						case 0xC7:
							instruction = ADD
						case 0xC8:
							instruction = ADD
						case 0xC9:
							instruction = ADD
						case 0xCA:
							instruction = ADD
						case 0xCB:
							instruction = ADD
						case 0xCC:
							instruction = ADD
						case 0xCD:
							instruction = ADD
						case 0xCE:
							instruction = ADD
						case 0xCF:
							instruction = ADD
						case 0xD0:
							instruction = ADD
						case 0xD1:
							instruction = ADD
						case 0xD2:
							instruction = ADD
						case 0xD3:
							instruction = ADD
						case 0xD4:
							instruction = ADD
						case 0xD5:
							instruction = ADD
						case 0xD6:
							instruction = ADD
						case 0xD7:
							instruction = ADD
						case 0xD8:
							instruction = ADD
						case 0xD9:
							instruction = ADD
						case 0xDA:
							instruction = ADD
						case 0xDB:
							instruction = ADD
						case 0xDC:
							instruction = ADD
						case 0xDD:
							instruction = ADD
						case 0xDE:
							instruction = ADD
						case 0xDF:
							instruction = ADD
						case 0xE0:
							instruction = ADD
						case 0xE1:
							instruction = ADD
						case 0xE2:
							instruction = ADD
						case 0xE3:
							instruction = ADD
						case 0xE4:
							instruction = ADD
						case 0xE5:
							instruction = ADD
						case 0xE6:
							instruction = ADD
						case 0xE7:
							instruction = ADD
						case 0xE8:
							instruction = ADD
						case 0xE9:
							instruction = ADD
						case 0xEA:
							instruction = ADD
						case 0xEB:
							instruction = ADD
						case 0xEC:
							instruction = ADD
						case 0xED:
							instruction = ADD
						case 0xEE:
							instruction = ADD
						case 0xEF:
							instruction = ADD
						case 0xF0:
							instruction = ADD
						case 0xF1:
							instruction = ADD
						case 0xF2:
							instruction = ADD
						case 0xF3:
							instruction = ADD
						case 0xF4:
							instruction = ADD
						case 0xF5:
							instruction = ADD
						case 0xF6:
							instruction = ADD
						case 0xF7:
							instruction = ADD
						case 0xF8:
							instruction = ADD
						case 0xF9:
							instruction = ADD
						case 0xFA:
							instruction = ADD
						case 0xFB:
							instruction = ADD
						case 0xFC:
							instruction = ADD
						case 0xFD:
							instruction = ADD
						case 0xFE:
							instruction = ADD
						case 0xFF:
							instruction = ADD
						}
					}
				} else {
				}
			} else if isXop {
				if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == false {
					// 0x01
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						instruction = ADD
					case 0x65:
						instruction = ADD
					case 0x66:
						instruction = ADD
					case 0x67:
						instruction = ADD
					case 0x68:
						instruction = ADD
					case 0x69:
						instruction = ADD
					case 0x6A:
						instruction = ADD
					case 0x6B:
						instruction = ADD
					case 0x6C:
						instruction = ADD
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == true {
					// 0x10
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						instruction = ADD
					case 0x65:
						instruction = ADD
					case 0x66:
						instruction = ADD
					case 0x67:
						instruction = ADD
					case 0x68:
						instruction = ADD
					case 0x69:
						instruction = ADD
					case 0x6A:
						instruction = ADD
					case 0x6B:
						instruction = ADD
					case 0x6C:
						instruction = ADD
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == true && mapSelect[4] == false {
					// 0x11
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						instruction = ADD
					case 0x65:
						instruction = ADD
					case 0x66:
						instruction = ADD
					case 0x67:
						instruction = ADD
					case 0x68:
						instruction = ADD
					case 0x69:
						instruction = ADD
					case 0x6A:
						instruction = ADD
					case 0x6B:
						instruction = ADD
					case 0x6C:
						instruction = ADD
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
				}
			} else if isSecondaryMap {
				if is3dNow {
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						instruction = ADD
					case 0x65:
						instruction = ADD
					case 0x66:
						instruction = ADD
					case 0x67:
						instruction = ADD
					case 0x68:
						instruction = ADD
					case 0x69:
						instruction = ADD
					case 0x6A:
						instruction = ADD
					case 0x6B:
						instruction = ADD
					case 0x6C:
						instruction = ADD
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
				} else if is38 {
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						instruction = ADD
					case 0x65:
						instruction = ADD
					case 0x66:
						instruction = ADD
					case 0x67:
						instruction = ADD
					case 0x68:
						instruction = ADD
					case 0x69:
						instruction = ADD
					case 0x6A:
						instruction = ADD
					case 0x6B:
						instruction = ADD
					case 0x6C:
						instruction = ADD
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
				} else if is3A {
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						instruction = ADD
					case 0x65:
						instruction = ADD
					case 0x66:
						instruction = ADD
					case 0x67:
						instruction = ADD
					case 0x68:
						instruction = ADD
					case 0x69:
						instruction = ADD
					case 0x6A:
						instruction = ADD
					case 0x6B:
						instruction = ADD
					case 0x6C:
						instruction = ADD
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
				} else {
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						// instruction = ADD
					case 0x65:
						// instruction = ADD
					case 0x66:
						// instruction = ADD
					case 0x67:
						// instruction = ADD
					case 0x68:
						instruction = PUSH
					case 0x69:
						instruction = IMUL
					case 0x6A:
						instruction = PUSH
					case 0x6B:
						instruction = IMUL
					case 0x6C:
						instruction = INS
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
				}
			} else {
				if isLock {
				}
				if bitFormat {
					switch curByte {
					case 0x0:
						instruction = ADD
					case 0x1:
						instruction = ADD
					case 0x2:
						instruction = ADD
					case 0x3:
						instruction = ADD
					case 0x4:
						instruction = ADD
					case 0x5:
						instruction = ADD
					case 0x6:
						instruction = PUSH
					case 0x7:
						instruction = POP
					case 0x8:
						instruction = OR
					case 0x9:
						instruction = OR
					case 0xA:
						instruction = OR
					case 0xB:
						instruction = OR
					case 0xC:
						instruction = OR
					case 0xD:
						instruction = OR
					case 0xE:
						instruction = PUSH
					case 0xF:
						// escape to 2nd opcode map
					case 0x10:
						instruction = ADC
					case 0x11:
						instruction = ADC
					case 0x12:
						instruction = ADC
					case 0x13:
						instruction = ADC
					case 0x14:
						instruction = ADC
					case 0x15:
						instruction = ADC
					case 0x16:
						instruction = PUSH
					case 0x17:
						instruction = POP
					case 0x18:
						instruction = SBB
					case 0x19:
						instruction = SBB
					case 0x1A:
						instruction = SBB
					case 0x1B:
						instruction = SBB
					case 0x1C:
						instruction = SBB
					case 0x1D:
						instruction = SBB
					case 0x1E:
						instruction = PUSH
					case 0x1F:
						instruction = POP
					case 0x20:
						instruction = AND
					case 0x21:
						instruction = AND
					case 0x22:
						instruction = AND
					case 0x23:
						instruction = AND
					case 0x24:
						instruction = AND
					case 0x25:
						instruction = AND
					case 0x26:
						// instruction = ADD
					case 0x27:
						instruction = DAA
					case 0x28:
						instruction = SUB
					case 0x29:
						instruction = SUB
					case 0x2A:
						instruction = SUB
					case 0x2B:
						instruction = SUB
					case 0x2C:
						instruction = SUB
					case 0x2D:
						instruction = SUB
					case 0x2E:
						// instruction = ADD
					case 0x2F:
						instruction = DAS
					case 0x30:
						instruction = XOR
					case 0x31:
						instruction = XOR
					case 0x32:
						instruction = XOR
					case 0x33:
						instruction = XOR
					case 0x34:
						instruction = XOR
					case 0x35:
						instruction = XOR
					case 0x36:
						// instruction = ADD
					case 0x37:
						instruction = AAA
					case 0x38:
						instruction = CMP
					case 0x39:
						instruction = CMP
					case 0x3A:
						instruction = CMP
					case 0x3B:
						instruction = CMP
					case 0x3C:
						instruction = CMP
					case 0x3D:
						instruction = CMP
					case 0x3E:
						// instruction = ADD
					case 0x3F:
						instruction = AAS
					case 0x40:
						// instruction = ADD
					case 0x41:
						// instruction = ADD
					case 0x42:
						// instruction = ADD
					case 0x43:
						// instruction = ADD
					case 0x44:
						// instruction = ADD
					case 0x45:
						// instruction = ADD
					case 0x46:
						// instruction = ADD
					case 0x47:
						// instruction = ADD
					case 0x48:
						// instruction = ADD
					case 0x49:
						// instruction = ADD
					case 0x4A:
						// instruction = ADD
					case 0x4B:
						// instruction = ADD
					case 0x4C:
						// instruction = ADD
					case 0x4D:
						// instruction = ADD
					case 0x4E:
						// instruction = ADD
					case 0x4F:
						// instruction = ADD
					case 0x50:
						instruction = PUSH
					case 0x51:
						instruction = PUSH
					case 0x52:
						instruction = PUSH
					case 0x53:
						instruction = PUSH
					case 0x54:
						instruction = PUSH
					case 0x55:
						instruction = PUSH
					case 0x56:
						instruction = PUSH
					case 0x57:
						instruction = PUSH
					case 0x58:
						instruction = POP
					case 0x59:
						instruction = POP
					case 0x5A:
						instruction = POP
					case 0x5B:
						instruction = POP
					case 0x5C:
						instruction = POP
					case 0x5D:
						instruction = POP
					case 0x5E:
						instruction = POP
					case 0x5F:
						instruction = POP
					case 0x60:
						// instruction = ADD operand size override
					case 0x61:
						// instruction = ADD operand size override
					case 0x62:
						instruction = BOUND
					case 0x63:
						instruction = ADD
					case 0x64:
						instruction = ADD
					case 0x65:
						instruction = ADD
					case 0x66:
						instruction = ADD
					case 0x67:
						instruction = ADD
					case 0x68:
						instruction = ADD
					case 0x69:
						instruction = ADD
					case 0x6A:
						instruction = ADD
					case 0x6B:
						instruction = ADD
					case 0x6C:
						instruction = ADD
					case 0x6D:
						instruction = ADD
					case 0x6E:
						instruction = ADD
					case 0x6F:
						instruction = ADD
					case 0x70:
						instruction = ADD
					case 0x71:
						instruction = ADD
					case 0x72:
						instruction = ADD
					case 0x73:
						instruction = ADD
					case 0x74:
						instruction = ADD
					case 0x75:
						instruction = ADD
					case 0x76:
						instruction = ADD
					case 0x77:
						instruction = ADD
					case 0x78:
						instruction = ADD
					case 0x79:
						instruction = ADD
					case 0x7A:
						instruction = ADD
					case 0x7B:
						instruction = ADD
					case 0x7C:
						instruction = ADD
					case 0x7D:
						instruction = ADD
					case 0x7E:
						instruction = ADD
					case 0x7F:
						instruction = ADD
					case 0x80:
						instruction = ADD
					case 0x81:
						instruction = ADD
					case 0x82:
						instruction = ADD
					case 0x83:
						instruction = ADD
					case 0x84:
						instruction = ADD
					case 0x85:
						instruction = ADD
					case 0x86:
						instruction = ADD
					case 0x87:
						instruction = ADD
					case 0x88:
						instruction = ADD
					case 0x89:
						instruction = ADD
					case 0x8A:
						instruction = ADD
					case 0x8B:
						instruction = ADD
					case 0x8C:
						instruction = ADD
					case 0x8D:
						instruction = ADD
					case 0x8E:
						instruction = ADD
					case 0x8F:
						instruction = ADD
					case 0x90:
						instruction = ADD
					case 0x91:
						instruction = ADD
					case 0x92:
						instruction = ADD
					case 0x93:
						instruction = ADD
					case 0x94:
						instruction = ADD
					case 0x95:
						instruction = ADD
					case 0x96:
						instruction = ADD
					case 0x97:
						instruction = ADD
					case 0x98:
						instruction = ADD
					case 0x99:
						instruction = ADD
					case 0x9A:
						instruction = ADD
					case 0x9B:
						instruction = ADD
					case 0x9C:
						instruction = ADD
					case 0x9D:
						instruction = ADD
					case 0x9E:
						instruction = ADD
					case 0x9F:
						instruction = ADD
					case 0xA0:
						instruction = ADD
					case 0xA1:
						instruction = ADD
					case 0xA2:
						instruction = ADD
					case 0xA3:
						instruction = ADD
					case 0xA4:
						instruction = ADD
					case 0xA5:
						instruction = ADD
					case 0xA6:
						instruction = ADD
					case 0xA7:
						instruction = ADD
					case 0xA8:
						instruction = ADD
					case 0xA9:
						instruction = ADD
					case 0xAA:
						instruction = ADD
					case 0xAB:
						instruction = ADD
					case 0xAC:
						instruction = ADD
					case 0xAD:
						instruction = ADD
					case 0xAE:
						instruction = ADD
					case 0xAF:
						instruction = ADD
					case 0xB0:
						instruction = ADD
					case 0xB1:
						instruction = ADD
					case 0xB2:
						instruction = ADD
					case 0xB3:
						instruction = ADD
					case 0xB4:
						instruction = ADD
					case 0xB5:
						instruction = ADD
					case 0xB6:
						instruction = ADD
					case 0xB7:
						instruction = ADD
					case 0xB8:
						instruction = ADD
					case 0xB9:
						instruction = ADD
					case 0xBA:
						instruction = ADD
					case 0xBB:
						instruction = ADD
					case 0xBC:
						instruction = ADD
					case 0xBD:
						instruction = ADD
					case 0xBE:
						instruction = ADD
					case 0xBF:
						instruction = ADD
					case 0xC0:
						instruction = ADD
					case 0xC1:
						instruction = ADD
					case 0xC2:
						instruction = ADD
					case 0xC3:
						instruction = ADD
					case 0xC4:
						instruction = ADD
					case 0xC5:
						instruction = ADD
					case 0xC6:
						instruction = ADD
					case 0xC7:
						instruction = ADD
					case 0xC8:
						instruction = ADD
					case 0xC9:
						instruction = ADD
					case 0xCA:
						instruction = ADD
					case 0xCB:
						instruction = ADD
					case 0xCC:
						instruction = ADD
					case 0xCD:
						instruction = ADD
					case 0xCE:
						instruction = ADD
					case 0xCF:
						instruction = ADD
					case 0xD0:
						instruction = ADD
					case 0xD1:
						instruction = ADD
					case 0xD2:
						instruction = ADD
					case 0xD3:
						instruction = ADD
					case 0xD4:
						instruction = ADD
					case 0xD5:
						instruction = ADD
					case 0xD6:
						instruction = ADD
					case 0xD7:
						instruction = ADD
					case 0xD8:
						instruction = ADD
					case 0xD9:
						instruction = ADD
					case 0xDA:
						instruction = ADD
					case 0xDB:
						instruction = ADD
					case 0xDC:
						instruction = ADD
					case 0xDD:
						instruction = ADD
					case 0xDE:
						instruction = ADD
					case 0xDF:
						instruction = ADD
					case 0xE0:
						instruction = ADD
					case 0xE1:
						instruction = ADD
					case 0xE2:
						instruction = ADD
					case 0xE3:
						instruction = ADD
					case 0xE4:
						instruction = ADD
					case 0xE5:
						instruction = ADD
					case 0xE6:
						instruction = ADD
					case 0xE7:
						instruction = ADD
					case 0xE8:
						instruction = ADD
					case 0xE9:
						instruction = ADD
					case 0xEA:
						instruction = ADD
					case 0xEB:
						instruction = ADD
					case 0xEC:
						instruction = ADD
					case 0xED:
						instruction = ADD
					case 0xEE:
						instruction = ADD
					case 0xEF:
						instruction = ADD
					case 0xF0:
						instruction = ADD
					case 0xF1:
						instruction = ADD
					case 0xF2:
						instruction = ADD
					case 0xF3:
						instruction = ADD
					case 0xF4:
						instruction = ADD
					case 0xF5:
						instruction = ADD
					case 0xF6:
						instruction = ADD
					case 0xF7:
						instruction = ADD
					case 0xF8:
						instruction = ADD
					case 0xF9:
						instruction = ADD
					case 0xFA:
						instruction = ADD
					case 0xFB:
						instruction = ADD
					case 0xFC:
						instruction = ADD
					case 0xFD:
						instruction = ADD
					case 0xFE:
						instruction = ADD
					case 0xFF:
						instruction = ADD
					}
					if isRexPrefix {
					}
				} else {
					
				}
			}
			isOpcode = false
			isModRM = true
		}
		// modr/m
		if isModRM {
			modRMByte := curByte
			modrmMod[0] = modRMByte/128 == 1
			if modrmMod[0] {
				modRMByte -= 128
			}
			modrmMod[1] = modRMByte/64 == 1
			if modrmMod[1] {
				modRMByte -= 64
			}
			modrmReg[0] = modRMByte/32 == 1
			if modrmReg[0] {
				modRMByte -= 32
			}
			modrmReg[1] = modRMByte/16 == 1
			if modrmReg[1] {
				modRMByte -= 16
			}
			modrmReg[2] = modRMByte/8 == 1
			if modrmReg[2] {
				modRMByte -= 8
			}
			if isRexR && r3B {
			}
			if isRexB && b3B {
			}
			modrmRM[0] = modRMByte/4 == 1
			if modrmRM[0] {
				modRMByte -= 4
			}
			modrmRM[1] = modRMByte/2 == 1
			if modrmRM[1] {
				modRMByte -= 4
			}
			modrmRM[2] = modRMByte == 1
			if modrmRM[2] {
				modRMByte -= 1
			}
			isModRM = false
			isSib = true
		}
		// sib
		if isSib {
			sibByte := curByte
			sibScale[0] = sibByte/128 == 1
			if sibScale[0] {
				sibByte -= 128
			}
			sibScale[1] = sibByte/64 == 1
			if sibScale[1] {
				sibByte -= 64
			}
			sibIndex[0] = sibByte/32 == 1
			if sibIndex[0] {
				sibByte -= 32
			}
			sibIndex[1] = sibByte/16 == 1
			if sibIndex[1] {
				sibByte -= 16
			}
			sibIndex[2] = sibByte/8 == 1
			if sibIndex[2] {
				sibByte -= 8
			}
			if isRexX {
			}
			if isRexB {
			}
			sibBase[0] = sibByte/4 == 1
			if sibBase[0] {
				sibByte -= 4
			}
			sibBase[1] = sibByte/2 == 1
			if sibBase[1] {
				sibByte -= 4
			}
			sibBase[2] = sibByte == 1
			if sibBase[2] {
				sibByte -= 1
			}
			isSib = false
			isImmediate = true
		}
		// immediate
		if isImmediate {
			isImmediate = false
			isDisplacement = true
		}
		// displacemet
		if isDisplacement {
			if curByte != 0 {
				isDisplacement = false
			}
		}
	}
}
