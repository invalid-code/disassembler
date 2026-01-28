package amdx8664

import (
	"fmt"

	"github.com/invalid-code/disassembler/executableFeatures"
	"github.com/invalid-code/disassembler/util"
)

func DisassembleBytes(data []byte, bitFormat bool, endianness bool, execFeatures executableFeatures.BinSecFeatures) {
	isPrefix, isEscapeSequence, isOpcode, isModRM, isSib, isDisplacement, isImmediate := true, false, false, false, false, false, false
	legacePrefixCnt := 0
	isOperandSizeOverride, isRep1, isRep0 := false, false, false
	//	isAddressSizeOverride := false
	//	isCSSegmentSizeOverride := false
	//	isDSSegmentSizeOverride := false
	//	isESSegmentSizeOverride := false
	//	isFSSegmentSizeOverride := false
	//	isGSSegmentSizeOverride := false
	//	isSSSegmentSizeOverride := false
	isRexW, isRexR, isRexX, isRexB := false, false, false, false
	isVex := false
	isXop := false
	isVex3Byte := false
	isSecondaryMap, is3dNow, is38, is3A := false, false, false, false
	fieldR := false
	fieldVvvv := [4]bool{false, false, false, false}
	fieldL := false
	fieldPp := [2]bool{false, false}
	isRXB := true
	r3B, x3B, b3B := false, false, false
	mapSelect := [5]bool{false, false, false, false, false}
	modrmMod, modrmReg, modrmRM := [2]bool{false, false}, [4]bool{false, false, false, false}, [4]bool{false, false, false}
	sibScale, sibIndex, sibBase := [2]bool{false, false}, [4]bool{false, false, false, false}, [4]bool{false, false, false, false}
	instruction := AAA
	isGpr := true
	opcode := byte(0)
	isCet := false
	displacementBytesI, displacementBytesVal, displacementBytes, wasDisplacement := 1, []byte{}, 0, false
	immediateBytesI, immediateBytesVal := 1, []byte{}
	operands := []Operand{}
	resetVars := func() {
		isPrefix, isEscapeSequence, isOpcode, isModRM, isSib, isDisplacement, isImmediate = true, false, false, false, false, false, false
		legacePrefixCnt = 0
		isOperandSizeOverride, isRep1, isRep0 = false, false, false
		//	isAddressSizeOverride := false
		//	isCSSegmentSizeOverride := false
		//	isDSSegmentSizeOverride := false
		//	isESSegmentSizeOverride := false
		//	isFSSegmentSizeOverride := false
		//	isGSSegmentSizeOverride := false
		//	isSSSegmentSizeOverride := false
		isRexW, isRexR, isRexX, isRexB = false, false, false, false
		isVex = false
		isXop = false
		isVex3Byte = false
		isSecondaryMap, is3dNow, is38, is3A = false, false, false, false
		fieldR = false
		fieldVvvv = [4]bool{false, false, false, false}
		fieldL = false
		fieldPp = [2]bool{false, false}
		isRXB = true
		r3B, x3B, b3B = false, false, false
		mapSelect = [5]bool{false, false, false, false, false}
		modrmMod, modrmReg, modrmRM = [2]bool{false, false}, [4]bool{false, false, false, false}, [4]bool{false, false, false}
		sibScale, sibIndex, sibBase = [2]bool{false, false}, [4]bool{false, false, false, false}, [4]bool{false, false, false, false}
		instruction = AAA
		isGpr = true
		opcode = byte(0)
		isCet = false
		displacementBytesI, displacementBytesVal, displacementBytes, wasDisplacement = 1, []byte{}, 0, false
		immediateBytesI, immediateBytesVal = 1, []byte{}
		operands = []Operand{}
	}
	for _, curByte := range data {
		fmt.Printf("%X\n", curByte)
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
			// case 0xF0:
			// 	if legacePrefixCnt == 4 {
			// 		panic("Error: There can only be 4 legacy prefixes in 1 instruction")
			// 	}
			// 	isLock = true
			// 	legacePrefixCnt += 1
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
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = false
				isRexB = false
				continue
			case 0x41:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = false
				isRexB = true
				continue
			case 0x42:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = true
				isRexB = false
				continue
			case 0x43:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = false
				isRexX = true
				isRexB = true
				continue
			case 0x44:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = false
				isRexB = false
				continue
			case 0x45:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = false
				isRexB = true
				continue
			case 0x46:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = true
				isRexB = false
				continue
			case 0x47:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = false
				isRexR = true
				isRexX = true
				isRexB = true
				continue
			case 0x48:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = false
				isRexB = false
				continue
			case 0x49:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = false
				isRexB = true
				continue
			case 0x4A:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = true
				isRexB = false
				continue
			case 0x4B:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = false
				isRexX = true
				isRexB = true
				continue
			case 0x4C:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = false
				isRexB = false
				continue
			case 0x4D:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = false
				isRexB = true
				continue
			case 0x4E:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = true
				isRexB = false
				continue
			case 0x4F:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
				isPrefix = false
				isEscapeSequence = true
				isRexW = true
				isRexR = true
				isRexX = true
				isRexB = true
				continue
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
					continue
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
					continue
				}
			} else {
				switch curByte {
				case 0xF:
					if isSecondaryMap {
						is3dNow = true
						isSecondaryMap = false
						isEscapeSequence = false
						isOpcode = true
						continue
					} else {
						isSecondaryMap = true
					}
				case 0x38:
					if isSecondaryMap {
						is38 = true
						isEscapeSequence = false
						isOpcode = true
						continue
					}
				case 0x3A:
					if isSecondaryMap {
						is3A = true
						isEscapeSequence = false
						isOpcode = true
						continue
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
						instruction, operands = vexOpcodeMap1(curByte, fieldPp, bitFormat, isRexW)
					} else if mapSelect[3] == true && mapSelect[4] == false {
						instruction, operands = vexOpcodeMap2(curByte, fieldPp, bitFormat, isRexW)
					} else if mapSelect[3] == true && mapSelect[4] == true {
						instruction, operands = vexOpcodeMap3(curByte, fieldPp, bitFormat, isRexW)
					}
				} else {
					instruction, operands = vexOpcodeMap1(curByte, fieldPp, bitFormat, isRexW)
				}
				isGpr = false
			} else if isXop {
				if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == false {
					instruction, operands = xopOpcodeMap8(curByte, fieldPp, bitFormat, isRexW)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == true {
					instruction, operands = xopOpcodeMap9(curByte, fieldPp, bitFormat, isRexW)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == true && mapSelect[4] == false {
					instruction, operands = xopOpcodeMap10(curByte, fieldPp, bitFormat, isRexW)
				}
				isGpr = false
			} else if is3dNow {
				isModRM = true
				isImmediate = true
			} else if is3A {
				instruction, operands = opcodeMap3A(curByte, bitFormat, isOperandSizeOverride, isRexB)
			} else if is38 {
				instruction, operands = opcodeMap38(curByte, bitFormat, isRep0, isOperandSizeOverride, isRexW)
			} else if isSecondaryMap {
				if isCet {
					switch execFeatures {
					case executableFeatures.IBT, executableFeatures.Both:
						if bitFormat {
							instruction = ENDBR64
						} else {
							instruction = ENDBR32
						}
						fmt.Println(instruction)
						resetVars()
						continue
					}
				}
				if isRep1 {
					isCet = true
					continue
				}
				instruction, operands = secondaryOpcodeMap(curByte, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
			} else {
				instruction, operands = primaryOpcode(curByte, bitFormat, isOperandSizeOverride, isRexW)
			}
			for _, operand := range operands {
				isModRM = isModRM || operand.isModRM
				isDisplacement = isDisplacement || operand.isDisplacement
				isImmediate = isImmediate || operand.isImmediate
				isGpr = isGpr || operand.isGpr
			}
			if instruction == NoInstruction {
				opcode = curByte
			}
			if !isModRM && !isDisplacement && !isImmediate {
				fmt.Print(instruction)
				for _, operand := range operands {
					fmt.Println(operand)
				}
				resetVars()
			}
			isOpcode = false
			continue
		}
		// modr/m
		if isModRM {
			modrmMod[0] = curByte/128 == 1
			if modrmMod[0] {
				curByte -= 128
			}
			modrmMod[1] = curByte/64 == 1
			if modrmMod[1] {
				curByte -= 64
			}
			modrmReg[1] = curByte/32 == 1
			if modrmReg[1] {
				curByte -= 32
			}
			modrmReg[2] = curByte/16 == 1
			if modrmReg[2] {
				curByte -= 16
			}
			modrmReg[3] = curByte/8 == 1
			if modrmReg[3] {
				curByte -= 8
			}
			modrmRM[1] = curByte/4 == 1
			if modrmRM[1] {
				curByte -= 4
			}
			modrmRM[2] = curByte/2 == 1
			if modrmRM[2] {
				curByte -= 4
			}
			modrmRM[3] = curByte == 1
			if modrmRM[3] {
				curByte -= 1
			}
			modrmReg[0] = isRexR || r3B
			if modrmRM != [4]bool{false, true, false, false} {
				modrmRM[0] = isRexB || b3B
			}
			isModRM = false
			if instruction == NoInstruction {
				opcodeSel := [3]bool{modrmReg[1], modrmReg[2], modrmReg[3]}
				if isVex {
					if isVex3Byte {
						if !mapSelect[3] && mapSelect[4] {
							instruction, operands = vexOpcodeMap1ModRMG1(opcode, fieldPp, opcodeSel, bitFormat, isRexW)
						} else {
							panic("Error: Unknown opcode modrm extension group")
						}
					} else {
						instruction, operands = vexOpcodeMap1ModRMG1(opcode, fieldPp, opcodeSel, bitFormat, isRexW)
					}
				} else if isXop {
					if mapSelect[1] && !mapSelect[2] && !mapSelect[3] && mapSelect[4] {
						switch opcode {
						case 0x1:
							instruction, operands = xopOpcodeMap9ModRMG1(opcode, opcodeSel, bitFormat, isRexW)
						case 0x2:
							instruction, operands = xopOpcodeMap9ModRMG2(opcode, opcodeSel, bitFormat, isRexW)
						case 0x12:
							instruction, operands = xopOpcodeMap9ModRMG3(opcode, opcodeSel, bitFormat, isRexW)
						default:
							panic("Error: Unknown opcode modrm extension group")
						}
					} else if mapSelect[1] && !mapSelect[2] && mapSelect[3] && !mapSelect[4] {
						switch opcode {
						case 0x12:
							instruction, operands = xopOpcodeMap9ModRMG4(opcode, opcodeSel, bitFormat, isRexW)
						default:
							panic("Error: Unknown opcode modrm extension group")
						}
					} else {
						panic("Error: Unknown opcode modrm extension group")
					}
				} else if isSecondaryMap {
					switch opcode {
					case 0x0:
						instruction, operands = secondaryOpcodeModRMG6(opcode, opcodeSel, bitFormat, isRexW)
					case 0x1:
						instruction, operands = secondaryOpcodeModRMG7(opcode, opcodeSel, bitFormat, isRexW)
					case 0xBA:
						instruction, operands = secondaryOpcodeModRMG8(opcode, opcodeSel, bitFormat, isRexW)
					case 0xC7:
						instruction, operands = secondaryOpcodeModRMG9(opcode, opcodeSel, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
					case 0xB9:
						instruction, operands = secondaryOpcodeModRMG10(opcode, opcodeSel, bitFormat, isRexW)
					case 0x71:
						instruction, operands = secondaryOpcodeModRMG12(opcode, opcodeSel, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
					case 0x72:
						instruction, operands = secondaryOpcodeModRMG13(opcode, opcodeSel, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
					case 0x73:
						instruction, operands = secondaryOpcodeModRMG14(opcode, opcodeSel, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
					case 0xAE:
						instruction, operands = secondaryOpcodeModRMG15(opcode, opcodeSel, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
					case 0x18:
						instruction, operands = secondaryOpcodeModRMG16(opcode, opcodeSel, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
					case 0x78:
						instruction, operands = secondaryOpcodeModRMG17(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					case 0x0D:
						instruction, operands = secondaryOpcodeModRMGP(opcode, opcodeSel, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
					default:
						panic("Error: Unknown opcode modrm extension group")
					}
				} else {
					switch opcode {
					case 0x80, 0x81, 0x82, 0x83:
						instruction, operands = primaryOpcodeModRMG1(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					case 0x8F:
						instruction, operands = primaryOpcodeModRMG1a(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					case 0xC0, 0xC1, 0xD0, 0xD1, 0xD2, 0xD3:
						instruction, operands = primaryOpcodeModRMG2(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					case 0xF6, 0xF7:
						instruction, operands = primaryOpcodeModRMG3(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					case 0xFE:
						instruction, operands = primaryOpcodeModRMG4(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					case 0xFF:
						instruction, operands = primaryOpcodeModRMG5(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					case 0xC6, 0xC7:
						instruction, operands = primaryOpcodeModRMG11(opcode, opcodeSel, bitFormat, isOperandSizeOverride, isRexW)
					default:
						panic("Error: Unknown opcode modrm extension group")
					}
					for _, operand := range operands {
						isDisplacement = operand.isDisplacement
						isImmediate = operand.isImmediate
					}
				}
				switch modrmMod {
				case [2]bool{false, false}:
					if modrmRM == [4]bool{false, true, false, true} {
						// 32 bit - absolute (displacement-only) addressing
						isDisplacement = true
						for operandI, operand := range operands {
							if operand.operandType == MemoryAddressOT {
								operands[operandI].noDisplacementBytes = 4
							}
						}
					}
				case [2]bool{false, true}:
					isDisplacement = true
					for operandI, operand := range operands {
						if operand.operandType == MemoryAddressOT {
							operands[operandI].noDisplacementBytes = 1
						}
					}
				case [2]bool{true, false}:
					isDisplacement = true
					for operandI, operand := range operands {
						if operand.operandType == MemoryAddressOT {
							operands[operandI].noDisplacementBytes = 4
						}
					}
				case [2]bool{true, true}:
					for operandI, operand := range operands {
						if operand.operandType == ImmediateOT {
							continue
						}
						operands[operandI].operandType = RegisterOT
						if operand.isModRMRegField {
							switch modrmReg {
							case [4]bool{false, false, false, false}:
								operands[operandI].register = RAX
							case [4]bool{false, false, false, true}:
								operands[operandI].register = RCX
							case [4]bool{false, false, true, false}:
								operands[operandI].register = RDX
							case [4]bool{false, false, true, true}:
								operands[operandI].register = RBX
							case [4]bool{false, true, false, false}:
								operands[operandI].register = RSP
							case [4]bool{false, true, false, true}:
								operands[operandI].register = RBP
							case [4]bool{false, true, true, false}:
								operands[operandI].register = RSI
							case [4]bool{false, true, true, true}:
								operands[operandI].register = RDI
							case [4]bool{true, false, false, false}:
								operands[operandI].register = R8
							case [4]bool{true, false, false, true}:
								operands[operandI].register = R9
							case [4]bool{true, false, true, false}:
								operands[operandI].register = R10
							case [4]bool{true, false, true, true}:
								operands[operandI].register = R11
							case [4]bool{true, true, false, false}:
								operands[operandI].register = R12
							case [4]bool{true, true, false, true}:
								operands[operandI].register = R13
							case [4]bool{true, true, true, false}:
								operands[operandI].register = R14
							case [4]bool{true, true, true, true}:
								operands[operandI].register = R15
							}
						} else {
							switch modrmRM {
							case [4]bool{false, false, false, false}:
								operands[operandI].register = RAX
							case [4]bool{false, false, false, true}:
								operands[operandI].register = RCX
							case [4]bool{false, false, true, false}:
								operands[operandI].register = RDX
							case [4]bool{false, false, true, true}:
								operands[operandI].register = RBX
							case [4]bool{false, true, false, false}:
								operands[operandI].register = RSP
							case [4]bool{false, true, false, true}:
								operands[operandI].register = RBP
							case [4]bool{false, true, true, false}:
								operands[operandI].register = RSI
							case [4]bool{false, true, true, true}:
								operands[operandI].register = RDI
							case [4]bool{true, false, false, false}:
								operands[operandI].register = R8
							case [4]bool{true, false, false, true}:
								operands[operandI].register = R9
							case [4]bool{true, false, true, false}:
								operands[operandI].register = R10
							case [4]bool{true, false, true, true}:
								operands[operandI].register = R11
							case [4]bool{true, true, false, false}:
								operands[operandI].register = R12
							case [4]bool{true, true, false, true}:
								operands[operandI].register = R13
							case [4]bool{true, true, true, false}:
								operands[operandI].register = R14
							case [4]bool{true, true, true, true}:
								operands[operandI].register = R15
							}
						}
					}
				}
			} else {
				if isGpr {
					switch modrmMod {
					case [2]bool{false, false}:
						// 32 bit - absolute (displacement-only) addressing
						if modrmRM == [4]bool{false, true, false, true} {
							isDisplacement = true
							for operandI, operand := range operands {
								if operand.operandType == MemoryAddressOT {
									operands[operandI].noDisplacementBytes = 4
								}
							}
						}
					case [2]bool{false, true}:
						isDisplacement = true
						for operandI, operand := range operands {
							if operand.operandType == MemoryAddressOT {
								operands[operandI].noDisplacementBytes = 1
							}
						}
					case [2]bool{true, false}:
						isDisplacement = true
						for operandI, operand := range operands {
							if operand.operandType == MemoryAddressOT {
								operands[operandI].noDisplacementBytes = 4
							}
						}
					case [2]bool{true, true}:
						for operandI, operand := range operands {
							operands[operandI].operandType = RegisterOT
							if operand.operandType == ImmediateOT {
								continue
							}
							if operand.isModRMRegField {
								switch modrmReg {
								case [4]bool{false, false, false, false}:
									operands[operandI].register = RAX
								case [4]bool{false, false, false, true}:
									operands[operandI].register = RCX
								case [4]bool{false, false, true, false}:
									operands[operandI].register = RDX
								case [4]bool{false, false, true, true}:
									operands[operandI].register = RBX
								case [4]bool{false, true, false, false}:
									operands[operandI].register = RSP
								case [4]bool{false, true, false, true}:
									operands[operandI].register = RBP
								case [4]bool{false, true, true, false}:
									operands[operandI].register = RSI
								case [4]bool{false, true, true, true}:
									operands[operandI].register = RDI
								case [4]bool{true, false, false, false}:
									operands[operandI].register = R8
								case [4]bool{true, false, false, true}:
									operands[operandI].register = R9
								case [4]bool{true, false, true, false}:
									operands[operandI].register = R10
								case [4]bool{true, false, true, true}:
									operands[operandI].register = R11
								case [4]bool{true, true, false, false}:
									operands[operandI].register = R12
								case [4]bool{true, true, false, true}:
									operands[operandI].register = R13
								case [4]bool{true, true, true, false}:
									operands[operandI].register = R14
								case [4]bool{true, true, true, true}:
									operands[operandI].register = R15
								}
							} else {
								switch modrmRM {
								case [4]bool{false, false, false, false}:
									operands[operandI].register = RAX
								case [4]bool{false, false, false, true}:
									operands[operandI].register = RCX
								case [4]bool{false, false, true, false}:
									operands[operandI].register = RDX
								case [4]bool{false, false, true, true}:
									operands[operandI].register = RBX
								case [4]bool{false, true, false, false}:
									operands[operandI].register = RSP
								case [4]bool{false, true, false, true}:
									operands[operandI].register = RBP
								case [4]bool{false, true, true, false}:
									operands[operandI].register = RSI
								case [4]bool{false, true, true, true}:
									operands[operandI].register = RDI
								case [4]bool{true, false, false, false}:
									operands[operandI].register = R8
								case [4]bool{true, false, false, true}:
									operands[operandI].register = R9
								case [4]bool{true, false, true, false}:
									operands[operandI].register = R10
								case [4]bool{true, false, true, true}:
									operands[operandI].register = R11
								case [4]bool{true, true, false, false}:
									operands[operandI].register = R12
								case [4]bool{true, true, false, true}:
									operands[operandI].register = R13
								case [4]bool{true, true, true, false}:
									operands[operandI].register = R14
								case [4]bool{true, true, true, true}:
									operands[operandI].register = R15
								}
							}
						}
					}
				} else {
					// xmm, ymm, control regs
				}
			}
			if !(modrmMod[0] && modrmMod[1]) {
				if modrmRM[1] && !modrmRM[2] && !modrmRM[3] {
					isSib = true
				}
			}
			if !(isSib || isDisplacement || isImmediate) {
				fmt.Printf("%v ", instruction)
				for _, operand := range operands {
					switch operand.operandType {
					case ImmediateOT:
					case RegisterOT:
						fmt.Printf("%v ", operand.register)
					case MemoryAddressOT:
						fmt.Printf("%v ", operand.memoryAddress)
					case UnknownOT:
					}
				}
				fmt.Println()
				resetVars()
			}
			continue
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
			sibIndex[1] = sibByte/32 == 1
			if sibIndex[0] {
				sibByte -= 32
			}
			sibIndex[2] = sibByte/16 == 1
			if sibIndex[1] {
				sibByte -= 16
			}
			sibIndex[3] = sibByte/8 == 1
			if sibIndex[2] {
				sibByte -= 8
			}
			sibBase[1] = sibByte/4 == 1
			if sibBase[0] {
				sibByte -= 4
			}
			sibBase[2] = sibByte/2 == 1
			if sibBase[1] {
				sibByte -= 4
			}
			sibBase[3] = sibByte == 1
			if sibBase[2] {
				sibByte -= 1
			}
			sibIndex[0] = isRexX || x3B
			sibBase[0] = isRexB || b3B

			sibScaleVal := 0
			switch sibScale {
			case [2]bool{false, false}:
				sibScaleVal = 1
			case [2]bool{false, true}:
				sibScaleVal = 2
			case [2]bool{true, false}:
				sibScaleVal = 4
			case [2]bool{true, true}:
				sibScaleVal = 8
			}
			sibIndexVal := NoRegister
			switch sibIndex {
			case [4]bool{false, false, false, false}:
				sibIndexVal = RAX
			case [4]bool{false, false, false, true}:
				sibIndexVal = RCX
			case [4]bool{false, false, true, false}:
				sibIndexVal = RDX
			case [4]bool{false, false, true, true}:
				sibIndexVal = RBX
			case [4]bool{false, true, false, false}:
			case [4]bool{false, true, false, true}:
				sibIndexVal = RBP
			case [4]bool{false, true, true, false}:
				sibIndexVal = RSI
			case [4]bool{false, true, true, true}:
				sibIndexVal = RDI
			}
			sibBaseVal := NoRegister
			// `switch modrmRM {
			// case [4]bool{false, false, false, false}:
			// case [4]bool{false, false, true, false}:
			// case [4]bool{false, true, false, false}:
			// case [4]bool{false, false, true, false}:
			// case [4]bool{false, true, true, false}:
			// case [4]bool{true, false, false, false}:
			// 	if modrmMod != [2]bool{true, true} {
			// 		switch sibBase {
			// 		case [3]bool{false, false, false}:
			// 			sibBaseVal = RAX
			// 		case [3]bool{false, false, true}:
			// 			sibBaseVal = RCX
			// 		case [3]bool{false, true, false}:
			// 			sibBaseVal = RDX
			// 		case [3]bool{false, true, true}:
			// 			sibBaseVal = RBX
			// 		case [3]bool{true, false, false}:
			// 			sibBaseVal = RSP
			// 		case [3]bool{true, false, true}:
			// 			isDisplacement = true
			// 			switch modrmMod {
			// 			case [2]bool{false, true}:
			// 				// todo need to specify how many displacement bytes
			// 			case [2]bool{true, false}:
			// 			}
			// 		case [3]bool{true, true, false}:
			// 			sibBaseVal = RSI
			// 		case [3]bool{true, true, true}:
			// 			sibBaseVal = RDI
			// 		}
			// 	}

			// case [4]bool{true, false, true, false}:
			// case [4]bool{true, true, false, false}:
			// case [4]bool{true, true, true, false}:
			// }`
			isSib = false
			fmt.Sprint(sibScaleVal, sibIndexVal, sibBaseVal)
			continue
		}
		// displacemet
		if isDisplacement {
			displacementBytesVal = append(displacementBytesVal, curByte)
			finishedDisp := false
			for _, operand := range operands {
				if displacementBytesI == operand.noDisplacementBytes {
					displacementBytes = util.ConvMultiByteToSingleByte(displacementBytesVal, endianness)
					fmt.Println("hello")
					if !isImmediate {
						resetVars()
					}
					wasDisplacement = true
					isDisplacement = false
					finishedDisp = true
					break
				}
			}
			if finishedDisp {
				continue
			}
			displacementBytesI += 1
			continue
		}

		// immediate
		if isImmediate {
			immediateBytesVal = append(immediateBytesVal, curByte)
			for _, operand := range operands {
				if operand.operandType != ImmediateOT {
					continue
				}
				if immediateBytesI == operand.noImmediateBytes {
					immediateBytes := util.ConvMultiByteToSingleByte(immediateBytesVal, endianness)
					fmt.Printf("%v ", instruction)
					if wasDisplacement {
						fmt.Printf("%X ", displacementBytes)
					}
					fmt.Printf("%X\n", immediateBytes)
					resetVars()
					continue
				}
			}
			immediateBytesI += 1
		}
		// fmt.Sprint(memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand)
	}
}
