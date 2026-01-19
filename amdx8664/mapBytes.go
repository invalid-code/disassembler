package amdx8664

import (
	"fmt"

	"github.com/invalid-code/disassembler/executableFeatures"
	"github.com/invalid-code/disassembler/util"
)

func DisassembleBytes(data []byte, bitFormat bool, endianness bool, execFeatures executableFeatures.BinSecFeatures) {
	isPrefix, isEscapeSequence, isOpcode, isModRM, isSib, isDisplacement, isImmediate := true, false, false, false, false, false, false
	legacePrefixCnt := 0
	isOperandSizeOverride := false
	//	isAddressSizeOverride := false
	//	isCSSegmentSizeOverride := false
	//	isDSSegmentSizeOverride := false
	//	isESSegmentSizeOverride := false
	//	isFSSegmentSizeOverride := false
	//	isGSSegmentSizeOverride := false
	//	isSSSegmentSizeOverride := false
	isRep1 := false
	isRep0 := false
	// isRexPrefix := false
	isRexW := false
	isRexR := false
	// isRexX := false
	isRexB := false
	isVex := false
	isXop := false
	isVex3Byte := false
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
	modrmMod := [2]bool{false, false}
	modrmReg := [3]bool{false, false, false}
	modrmRM := [4]bool{false, false, false}
	sibScale := [2]bool{false, false}
	sibIndex := [3]bool{false, false, false}
	sibBase := [3]bool{false, false, false}
	instruction := AAA
	memSegment := NoSegment
	regOperand1, regOperand2, regOperand3 := NoRegister, NoRegister, NoRegister
	instructionEncodedRegOperand := 0
	isGpr := true
	opcode := byte(0)
	isCet := false
	regOperand1ModRMReg := false
	noDisplacementBytes := 0
	displacementBytesI := 1
	displacementBytesVal := []byte{}
	noImmediateBytes := 0
	immediateBytesI := 1
	immediateBytesVal := []byte{}
	displacementBytes := 0
	wasDisplacement := false
	resetVars := func() {
		isPrefix = true
		legacePrefixCnt = 0
		isOperandSizeOverride = false
		//	sAddressSizeOverride := false
		//	sCSSegmentSizeOverride := false
		//	sDSSegmentSizeOverride := false
		//	sESSegmentSizeOverride := false
		//	sFSSegmentSizeOverride := false
		//	sGSSegmentSizeOverride := false
		//	sSSSegmentSizeOverride := false
		isRep1 = false
		isRep0 = false
		// sRexPrefix := false
		isRexW = false
		isRexR = false
		// sRexX := false
		isRexB = false
		isVex = false
		isXop = false
		isVex3Byte = false
		isEscapeSequence = false
		is3dNow = false
		is38 = false
		is3A = false
		isSecondaryMap = false
		fieldR = false
		fieldVvvv = [4]bool{false, false, false, false}
		fieldL = false
		fieldPp = [2]bool{false, false}
		isRXB = true
		r3B = false
		x3B = false
		b3B = false
		mapSelect = [5]bool{false, false, false, false, false}
		isModRM = false
		isOpcode = false
		isSib = false
		isImmediate = false
		isDisplacement = false
		modrmMod = [2]bool{false, false}
		modrmReg = [3]bool{false, false, false}
		modrmRM = [4]bool{false, false, false}
		sibScale = [2]bool{false, false}
		sibIndex = [3]bool{false, false, false}
		sibBase = [3]bool{false, false, false}
		instruction = AAA
		memSegment = NoSegment
		regOperand1, regOperand2, regOperand3 = NoRegister, NoRegister, NoRegister
		instructionEncodedRegOperand = 0
		isGpr = true
		opcode = byte(0)
		isCet = false
		regOperand1ModRMReg = false
		noDisplacementBytes = 0
		displacementBytesI = 1
		displacementBytesVal = []byte{}
		noImmediateBytes = 0
		immediateBytesI = 1
		immediateBytesVal = []byte{}
		displacementBytes = 0
		wasDisplacement = false
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
				// isRexX = false
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
				// isRexX = false
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
				// isRexX = true
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
				// isRexX = true
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
				// isRexX = false
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
				// isRexX = false
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
				// isRexX = true
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
				// isRexX = true
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
				// isRexX = false
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
				// isRexX = false
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
				// isRexX = true
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
				// isRexX = true
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
				// isRexX = false
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
				// isRexX = false
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
				// isRexX = true
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
				// isRexX = true
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
						// isEscapeSequence = false
						// isOpcode = true
					} else {
						isSecondaryMap = true
					}
				case 0x38:
					if isSecondaryMap {
						is38 = true
						// isEscapeSequence = false
						// isOpcode = true
					}
				case 0x3A:
					if isSecondaryMap {
						is3A = true
						// isEscapeSequence = false
						// isOpcode = true
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
						instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap1(curByte, fieldPp, isRexW)
					} else if mapSelect[3] == true && mapSelect[4] == false {
						instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap2(curByte, fieldPp, isRexW)
					} else if mapSelect[3] == true && mapSelect[4] == true {
						instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap3(curByte, fieldPp, isRexW)
					}
				} else {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap1(curByte, fieldPp, isRexW)
				}
				isGpr = false
			} else if isXop {
				if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == false {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap8(curByte, fieldPp, isRexW)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == true {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap9(curByte, fieldPp, isRexW)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == true && mapSelect[4] == false {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap10(curByte, fieldPp, isRexW)
				}
				isGpr = false
			} else if is3dNow {
				panic("todo handle 3dnow opcode map differently")
			} else if is3A {
				instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = opcodeMap3A(curByte, isOperandSizeOverride, isRexB)
			} else if is38 {
				instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = opcodeMap38(curByte, bitFormat, isRep0, isOperandSizeOverride, isRexB)
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
				instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeMap(curByte, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
			} else {
				instruction, isModRM, isDisplacement, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand, regOperand1ModRMReg, noImmediateBytes, noDisplacementBytes = primaryOpcode(curByte, bitFormat, isOperandSizeOverride, isRexW)
				fmt.Println(instruction, isModRM, isDisplacement, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand, regOperand1ModRMReg, noImmediateBytes)
			}
			if instruction == NoInstruction {
				opcode = curByte
			}
			if !(isModRM || isImmediate || isDisplacement) {
				fmt.Println(instruction, regOperand1, regOperand2)
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
			modrmReg[0] = curByte/32 == 1
			if modrmReg[0] {
				curByte -= 32
			}
			modrmReg[1] = curByte/16 == 1
			if modrmReg[1] {
				curByte -= 16
			}
			modrmReg[2] = curByte/8 == 1
			if modrmReg[2] {
				curByte -= 8
			}
			if isRexR && r3B {
			}
			if isRexB && b3B {
			}
			modrmRM[0] = curByte/4 == 1
			if modrmRM[0] {
				curByte -= 4
			}
			modrmRM[1] = curByte/2 == 1
			if modrmRM[1] {
				curByte -= 4
			}
			modrmRM[2] = curByte == 1
			if modrmRM[2] {
				curByte -= 1
			}
			isModRM = false
			if instruction == NoInstruction {
				if isVex {
					if isVex3Byte {
						if !mapSelect[3] && mapSelect[4] {
							instruction, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap1ModRMG1(opcode, fieldPp, modrmReg)
						} else {
							panic("Error: Unknown opcode modrm extension group")
						}
					} else {
						instruction, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap1ModRMG1(opcode, fieldPp, modrmReg)
					}
				} else if isXop {
					if mapSelect[1] && !mapSelect[2] && !mapSelect[3] && mapSelect[4] {
						switch opcode {
						case 0x1:
							instruction, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap9ModRMG1(opcode, modrmReg)
						case 0x2:
							instruction, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap9ModRMG2(opcode, modrmReg)
						case 0x12:
							instruction, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap9ModRMG3(opcode, modrmReg)
						default:
							panic("Error: Unknown opcode modrm extension group")
						}
					} else if mapSelect[1] && !mapSelect[2] && mapSelect[3] && !mapSelect[4] {
						switch opcode {
						case 0x12:
							instruction, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap9ModRMG4(opcode, modrmReg)
						default:
							panic("Error: Unknown opcode modrm extension group")
						}
					} else {
						panic("Error: Unknown opcode modrm extension group")
					}
				} else if isSecondaryMap {
					switch opcode {
					case 0x0:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG6(opcode, modrmReg)
					case 0x1:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG7(opcode, modrmReg)
					case 0xBA:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG8(opcode, modrmReg)
					case 0xC7:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG9(opcode, modrmReg, isRep0, isRep1, isOperandSizeOverride)
					case 0xB9:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG10(opcode, modrmReg)
					case 0x71:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG12(opcode, modrmReg, isRep0, isRep1, isOperandSizeOverride)
					case 0x72:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG13(opcode, modrmReg, isRep0, isRep1, isOperandSizeOverride)
					case 0x73:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG14(opcode, modrmReg, isRep0, isRep1, isOperandSizeOverride)
					case 0xAE:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG15(opcode, modrmReg, isRep0, isRep1, isOperandSizeOverride)
					case 0x18:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG16(opcode, modrmReg, isRep0, isRep1, isOperandSizeOverride)
					case 0x78:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMG17(opcode, modrmReg, isOperandSizeOverride)
					case 0x0D:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeModRMGP(opcode, modrmReg, isRep0, isRep1, isOperandSizeOverride)
					default:
						panic("Error: Unknown opcode modrm extension group")
					}
				} else {
					switch opcode {
					case 0x80, 0x81, 0x82, 0x83:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand, noImmediateBytes = primaryOpcodeModRMG1(opcode, modrmReg, bitFormat)
					case 0x8F:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG1a(opcode, modrmReg)
					case 0xC0, 0xC1, 0xD0, 0xD1, 0xD2, 0xD3:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand, noImmediateBytes = primaryOpcodeModRMG2(opcode, modrmReg)
					case 0xF6, 0xF7:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG3(opcode, modrmReg, bitFormat)
					case 0xFE:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG4(opcode, modrmReg, bitFormat)
					case 0xFF:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG5(opcode, modrmReg, bitFormat)
					case 0xC6, 0xC7:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG11(opcode, modrmReg, bitFormat)
					default:
						panic("Error: Unknown opcode modrm extension group")
					}
				}
				switch modrmMod {
				case [2]bool{false, false}:
					switch modrmRM {
					case [4]bool{false, false, false, false}:
					case [4]bool{false, false, true, false}:
					case [4]bool{false, true, false, false}:
					case [4]bool{false, true, true, false}:
					case [4]bool{true, false, false, false}:
					case [4]bool{true, false, true, false}:
						// 64 bit - rip addressing
						// 32 bit - absolute (displacement-only) addressing
						isDisplacement = true
						noDisplacementBytes = 4
					case [4]bool{true, true, false, false}:
					case [4]bool{true, true, true, false}:
					}
				case [2]bool{false, true}:
				case [2]bool{true, false}:
				case [2]bool{true, true}:
					switch modrmRM {
					case [4]bool{false, false, false, false}:
						regOperand1 = RAX
					case [4]bool{false, false, true, false}:
						regOperand1 = RCX
					case [4]bool{false, true, false, false}:
						regOperand1 = RDX
					case [4]bool{false, true, true, false}:
						regOperand1 = RBX
					case [4]bool{true, false, false, false}:
						regOperand1 = RSP
					case [4]bool{true, false, true, false}:
						regOperand1 = RBP
					case [4]bool{true, true, false, false}:
						regOperand1 = RSI
					case [4]bool{true, true, true, false}:
						regOperand1 = RDI
					}
					if !isImmediate {
						fmt.Println(instruction, regOperand1)
						resetVars()
						continue
					}
				}
			} else {
				if isGpr {
					// figure out operand size
					switch modrmMod {
					case [2]bool{false, false}:
						switch modrmReg {
						case [3]bool{false, false, false}:
							if regOperand1ModRMReg {
								regOperand1 = RAX
							} else {
								regOperand2 = RAX
							}
						case [3]bool{false, false, true}:
							if regOperand1ModRMReg {
								regOperand1 = RCX
							} else {
								regOperand2 = RCX
							}
						case [3]bool{false, true, false}:
							if regOperand1ModRMReg {
								regOperand1 = RDX
							} else {
								regOperand2 = RDX
							}
						case [3]bool{false, true, true}:
							if regOperand1ModRMReg {
								regOperand1 = RBX
							} else {
								regOperand2 = RBX
							}
						case [3]bool{true, false, false}:
							if regOperand1ModRMReg {
								regOperand1 = RSP
							} else {
								regOperand2 = RSP
							}
						case [3]bool{true, false, true}:
							if regOperand1ModRMReg {
								regOperand1 = RBP
							} else {
								regOperand2 = RBP
							}
						case [3]bool{true, true, false}:
							if regOperand1ModRMReg {
								regOperand1 = RSI
							} else {
								regOperand2 = RSI
							}
						case [3]bool{true, true, true}:
							if regOperand1ModRMReg {
								regOperand1 = RDI
							} else {
								regOperand2 = RDI
							}
						}
						switch modrmRM {
						case [4]bool{false, false, false, false}:
						case [4]bool{false, false, true, false}:
							if bitFormat {
								// 64 bit - rip addressing
								isDisplacement = true
							}
						case [4]bool{false, true, false, false}:
						case [4]bool{false, true, true, false}:
						case [4]bool{true, false, false, false}:
						case [4]bool{true, false, true, false}:
							// 64 bit - rip addressing
							// 32 bit - absolute (displacement-only) addressing
							isDisplacement = true
							noDisplacementBytes = 4
						case [4]bool{true, true, false, false}:
						case [4]bool{true, true, true, false}:
						}
						// if !(isSib || isDisplacement || isImmediate) {
						// 	fmt.Println(instruction, regOperand1, regOperand2, regOperand3)
						// 	resetVars()
						// 	continue
						// }

					case [2]bool{false, true}:
						isDisplacement = true
						noDisplacementBytes = 1
					case [2]bool{true, false}:
						isDisplacement = true
						noDisplacementBytes = 4
					case [2]bool{true, true}:
						switch modrmReg {
						case [3]bool{false, false, false}:
							if regOperand1ModRMReg {
								regOperand1 = RAX
							} else {
								regOperand2 = RAX
							}
						case [3]bool{false, false, true}:
							if regOperand1ModRMReg {
								regOperand1 = RCX
							} else {
								regOperand2 = RCX
							}
						case [3]bool{false, true, false}:
							if regOperand1ModRMReg {
								regOperand1 = RDX
							} else {
								regOperand2 = RDX
							}
						case [3]bool{false, true, true}:
							if regOperand1ModRMReg {
								regOperand1 = RBX
							} else {
								regOperand2 = RBX
							}
						case [3]bool{true, false, false}:
							if regOperand1ModRMReg {
								regOperand1 = RSP
							} else {
								regOperand2 = RSP
							}
						case [3]bool{true, false, true}:
							if regOperand1ModRMReg {
								regOperand1 = RBP
							} else {
								regOperand2 = RBP
							}
						case [3]bool{true, true, false}:
							if regOperand1ModRMReg {
								regOperand1 = RSI
							} else {
								regOperand2 = RSI
							}
						case [3]bool{true, true, true}:
							if regOperand1ModRMReg {
								regOperand1 = RDI
							} else {
								regOperand2 = RDI
							}
						}
						switch modrmRM {
						case [4]bool{false, false, false, false}:
							if regOperand1ModRMReg {
								regOperand2 = RAX
							} else {
								regOperand1 = RAX
							}
						case [4]bool{false, false, true, false}:
							if regOperand1ModRMReg {
								regOperand2 = RCX
							} else {
								regOperand1 = RCX
							}
						case [4]bool{false, true, false, false}:
							if regOperand1ModRMReg {
								regOperand2 = RDX
							} else {
								regOperand1 = RDX
							}
						case [4]bool{false, true, true, false}:
							if regOperand1ModRMReg {
								regOperand2 = RBX
							} else {
								regOperand1 = RBX
							}
						case [4]bool{true, false, false, false}:
							if regOperand1ModRMReg {
								regOperand2 = RSP
							} else {
								regOperand1 = RSP
							}
						case [4]bool{true, false, true, false}:
							if regOperand1ModRMReg {
								regOperand2 = RBP
							} else {
								regOperand1 = RBP
							}
						case [4]bool{true, true, false, false}:
							if regOperand1ModRMReg {
								regOperand2 = RSI
							} else {
								regOperand1 = RSI
							}
						case [4]bool{true, true, true, false}:
							if regOperand1ModRMReg {
								regOperand2 = RDI
							} else {
								regOperand1 = RDI
							}
						}
						// fmt.Printf("%v %v %v\n", instruction, regOperand1, regOperand2)
						// resetVars()
						// continue
					}
				} else {
					// xmm, ymm regs
				}
			}
			if !(modrmMod[0] && modrmMod[1]) {
				if modrmRM[0] && !modrmRM[1] && !modrmRM[2] {
					isSib = true
				}
			}
			if !(isSib || isDisplacement || isImmediate) {
				fmt.Println(instruction, regOperand1, regOperand2)
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
			case [3]bool{false, false, false}:
				sibIndexVal = RAX
			case [3]bool{false, false, true}:
				sibIndexVal = RCX
			case [3]bool{false, true, false}:
				sibIndexVal = RDX
			case [3]bool{false, true, true}:
				sibIndexVal = RBX
			case [3]bool{true, false, false}:
			case [3]bool{true, false, true}:
				sibIndexVal = RBP
			case [3]bool{true, true, false}:
				sibIndexVal = RSI
			case [3]bool{true, true, true}:
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
			if displacementBytesI == noDisplacementBytes {
				displacementBytes = util.ConvMultiByteToSingleByte(displacementBytesVal, endianness)
				if !isImmediate {
					fmt.Printf("%v %v %v %X\n", instruction, regOperand1, regOperand2, displacementBytes)
					resetVars()
				}
				wasDisplacement = true
				isDisplacement = false
				continue
			}
			displacementBytesI += 1
			continue
		}

		// immediate
		if isImmediate {
			immediateBytesVal = append(immediateBytesVal, curByte)
			if immediateBytesI == noImmediateBytes {
				immediateBytes := util.ConvMultiByteToSingleByte(immediateBytesVal, endianness)
				fmt.Printf("%v %v %v ", instruction, regOperand1, regOperand2)
				if wasDisplacement {
					fmt.Printf("%X ", displacementBytes)
				}
				fmt.Printf("%X\n", immediateBytes)
				resetVars()
				continue
			}
			immediateBytesI += 1
		}
		fmt.Sprint(memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand)
	}
}
