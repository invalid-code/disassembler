package amdx8664

import "fmt"

func DisassembleBytes(data []byte, bitFormat bool) {
	isPrefix := true
	legacePrefixCnt := 0
	isOperandSizeOverride := false
//	isAddressSizeOverride := false
//	isCSSegmentSizeOverride := false
//	isDSSegmentSizeOverride := false
//	isESSegmentSizeOverride := false
//	isFSSegmentSizeOverride := false
//	isGSSegmentSizeOverride := false
//	isSSSegmentSizeOverride := false
	isLock := false
	isRep1 := false
	isRep0 := false
	isRexPrefix := false
//	isRexW := false
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
	for _, curByte := range data {
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
//				isAddressSizeOverride = true
				legacePrefixCnt += 1
			case 0x2E:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
//				isCSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x3E:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
//				isDSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x26:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
//				isESSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x64:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
//				isFSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x65:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
//				isGSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0x36:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
//				isSSSegmentSizeOverride = true
				legacePrefixCnt += 1
			case 0xF0:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isLock = true
				legacePrefixCnt += 1
			case 0xF3:
				if legacePrefixCnt == 4 {
					panic("Error: There can only be 4 legacy prefixes in 1 instruction")
				}
				isRep1 = true
				legacePrefixCnt += 1
			case 0xF2:
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
//				isRexW = false
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
//				isRexW = false
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
//				isRexW = false
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
//				isRexW = false
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
//				isRexW = false
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
//				isRexW = false
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
//				isRexW = false
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
//				isRexW = false
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
//				isRexW = true
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
//				isRexW = true
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
//				isRexW = true
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
//				isRexW = true
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
//				isRexW = true
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
//				isRexW = true
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
//				isRexW = true
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
//				isRexW = true
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
						instruction = vexOpcodeMap1(curByte)
					} else if mapSelect[3] == true && mapSelect[4] == false {
						instruction = vexOpcodeMap2(curByte)
					} else if mapSelect[3] == true && mapSelect[4] == true {
						instruction = vexOpcodeMap3(curByte)
					}
				} else {
					vexOpcodeMap1(curByte)
				}
			} else if isXop {
				if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == false {
					instruction = xopOpcodeMap8(curByte)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == true {
					instruction = xopOpcodeMap9(curByte)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == true && mapSelect[4] == false {
					instruction = xopOpcodeMap10(curByte)
				}
			} else if is3dNow {
				if isLock {
				}
				if bitFormat {
					instruction = opcodeMap3dnow64(curByte)
					if isRexPrefix {
					}
				} else {
					instruction = opcodeMap3dnow32(curByte)
				}
			} else if is3A {
				if isLock {
				}
				if bitFormat {
					instruction = opcodeMap3A(curByte, isOperandSizeOverride)
					if isRexPrefix {
					}
				} else {
					instruction = opcodeMap3A(curByte, isOperandSizeOverride)
				}
			} else if is38 {
				if isLock {
				}
				if bitFormat {
					instruction = opcodeMap3864(curByte, isRep0, isRep1, isOperandSizeOverride)
					if isRexPrefix {
					}
				} else {
					instruction = opcodeMap3832(curByte, isRep0, isRep1, isOperandSizeOverride)
				}
			} else if isSecondaryMap {
				if isLock {
				}
				if bitFormat {
					instruction = secondaryOpcodeMap64(curByte, isRep0, isRep1, isOperandSizeOverride)
					if isRexPrefix {
					}
				} else {
					instruction = secondaryOpcodeMap32(curByte, isRep0, isRep1, isOperandSizeOverride)
				}
			} else {
				if isLock {
				}
				if bitFormat {
					instruction = primaryOpcode64(curByte)
					if isRexPrefix {
					}
				} else {
					instruction = primaryOpcode32(curByte, isOperandSizeOverride)
				}
			}
			isModRM = true
			isOpcode = false
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
			isDisplacement = false
			isPrefix = true
		}
		fmt.Println(instruction.String())
	}
}
