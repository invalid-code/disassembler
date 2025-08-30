package amdx8664

import "fmt"

type MemSegment int

const (
	CS MemSegment = iota
	DS
	ES
	FS
	GS
	SS
	NoSegment
)

func (memSegment MemSegment) String() string {
	switch memSegment {
	case CS:
		return "CS"
	case DS:
		return "DS"
	case ES:
		return "ES"
	case FS:
		return "FS"
	case GS:
		return "GS"
	case SS:
		return "SS"
	case NoSegment:
		return "NoSegment"
	default:
		return "Unknown segment"
	}
}

type Register int

const (
	AX Register = iota
	BX
	CX
	DX
	SI
	DI
	BP
	SP
	R8W
	R9W
	R10W
	R11W
	R12W
	R13W
	R14W
	R15W
	EAX
	EBX
	ECX
	EDX
	ESI
	EDI
	EBP
	ESP
	R8D
	R9D
	R10D
	R11D
	R12D
	R13D
	R14D
	R15D
	RAX
	RBX
	RCX
	RDX
	RSI
	RDI
	RBP
	RSP
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	NoRegister
)

func (register Register) String() string {
	switch register {
	case AX:
		return "AX"
	case BX:
		return "BX"
	case CX:
		return "CX"
	case DX:
		return "DX"
	case SI:
		return "SI"
	case DI:
		return "DI"
	case BP:
		return "BP"
	case SP:
		return "SP"
	case R8W:
		return "R8W"
	case R9W:
		return "R9W"
	case R10W:
		return "R10W"
	case R11W:
		return "R11W"
	case R12W:
		return "R12W"
	case R13W:
		return "R13W"
	case R14W:
		return "R14W"
	case R15W:
		return "R15W"
	case EAX:
		return "EAX"
	case EBX:
		return "EBX"
	case ECX:
		return "ECX"
	case EDX:
		return "EDX"
	case ESI:
		return "ESI"
	case EDI:
		return "EDI"
	case EBP:
		return "EBP"
	case ESP:
		return "ESP"
	case R8D:
		return "R8D"
	case R9D:
		return "R9D"
	case R10D:
		return "R10D"
	case R11D:
		return "R11D"
	case R12D:
		return "R12D"
	case R13D:
		return "R13D"
	case R14D:
		return "R14D"
	case R15D:
		return "R15D"
	case RAX:
		return "RAX"
	case RBX:
		return "RBX"
	case RCX:
		return "RCX"
	case RDX:
		return "RDX"
	case RSI:
		return "RSI"
	case RDI:
		return "RDI"
	case RBP:
		return "RBP"
	case RSP:
		return "RSP"
	case R8:
		return "R8"
	case R9:
		return "R9"
	case R10:
		return "R10"
	case R11:
		return "R11"
	case R12:
		return "R12"
	case R13:
		return "R13"
	case R14:
		return "R14"
	case R15:
		return "R15"
	case NoRegister:
		return "NoRegister"
	default:
		return "Unknown register"
	}
}

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
	isRep1 := false
	isRep0 := false
	// isRexPrefix := false
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
	modrmReg := [3]bool{false, false, false}
	modrmRM := [4]bool{false, false, false}
	sibScale := [2]bool{false, false}
	sibIndex := [3]bool{false, false, false}
	sibBase := [3]bool{false, false, false}
	instruction := AAA
	memSegment := NoSegment
	regOperand1, regOperand2, regOperand3 := NoRegister, NoRegister, NoRegister
	instructionEncodedRegOperand := 0
	opcode := byte(0)
	for _, curByte := range data {
		fmt.Println(curByte)
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
			case 0x4F:
				if !bitFormat {
					panic("Error: REX prefix are only allowed in 64-bit mode")
				}
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
						instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap1(curByte, fieldPp, isRexW)
					} else if mapSelect[3] == true && mapSelect[4] == false {
						instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap2(curByte, fieldPp, isRexW)
					} else if mapSelect[3] == true && mapSelect[4] == true {
						instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap3(curByte, fieldPp, isRexW)
					}
				} else {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = vexOpcodeMap1(curByte, fieldPp, isRexW)
				}
			} else if isXop {
				if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == false {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap8(curByte, fieldPp, isRexW)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == true {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap9(curByte, fieldPp, isRexW)
				} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == true && mapSelect[4] == false {
					instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, regOperand3, instructionEncodedRegOperand = xopOpcodeMap10(curByte, fieldPp, isRexW)
				}
			} else if is3dNow {
				panic("todo handle 3dnow opcode map differently")
				// instruction, isModRM, isDisplacement, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = opcodeMap3dnow(curByte, bitFormat, isRexB)
			} else if is3A {
				instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = opcodeMap3A(curByte, isOperandSizeOverride, isRexB)
			} else if is38 {
				instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = opcodeMap38(curByte, bitFormat, isRep0, isOperandSizeOverride, isRexB)
			} else if isSecondaryMap {
				instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = secondaryOpcodeMap(curByte, bitFormat, isRep0, isRep1, isOperandSizeOverride, isRexW)
			} else {
				instruction, isModRM, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcode(curByte, bitFormat, isOperandSizeOverride, isRexW)
			}
			if instruction == NoInstruction {
				opcode = curByte
			}
			isOpcode = false
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
				if isSecondaryMap {
					switch opcode {
					case 0x0:
					case 0x1:
					case 0xBA:
					case 0xC7:
					case 0xB9:
					case 0x71:
					case 0x72:
					case 0x73:
					default:
						panic("Error: Unknown opcode modrm extension group")
					}
				} else {
					switch opcode {
					case 0x80, 0x81, 0x82, 0x83:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG1(opcode, modrmReg, bitFormat)
					case 0x8F:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG1a(opcode, modrmReg, bitFormat)
					case 0xC0, 0xC1, 0xD0, 0xD1, 0xD2, 0xD3:
						instruction, isImmediate, memSegment, regOperand1, regOperand2, instructionEncodedRegOperand = primaryOpcodeModRMG2(opcode, modrmReg, bitFormat)
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
			}
			if !(modrmMod[0] && modrmMod[1]) {
				if modrmRM[0] && !modrmRM[1] && !modrmRM[2] {
					isSib = true
				}
			}
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
	}
}
