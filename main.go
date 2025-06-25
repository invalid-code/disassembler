package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
)

var elfMagicNumber []byte = []uint8{127, 69, 76, 70}

type InstructionSet int

const (
	NoInstruction InstructionSet = iota
	We32100
	Sparc
	x86
	M68k
	M88k
	Imcu
	I80860
	Mips
	IbmSystem370
	Rs3000Le
	HpPaRisc
	I80960
	PowerPc
	PowerPc64b
	S390
	IbmSpu
	NecV800
	FujitsoFr20
	TrwRh32
	MotorolaRce
	Arm
	DigitalAlpha
	SuperH
	SparceV9
	STriCodeEmbedProc
	ArgonautRiscCore
	HitachiH8300
	HitachiH8300H
	HitachiH8S
	HitachiH8500
	Ia64
	StanfordMips
	MotorolaColdFire
	M68GC12
	MmaMMAccel
	SiemensPcp
	SonynCPURiscProcessor
	DensoNDR1MicroProcessor
	StarCoreProcessor
	Me16
	StmeSt100
	AdvLogCorp
	Amdx8664
	SonyDsp
	Pdp10
	Pdp11
	Fx66
	St9816MicroController
	St8MicroController
	MC68HC16
	MC68HC11
	MC68HC08
	MC68HC05
	Svx
	St198Microcontroller
	DVax
	AxisComms32Embed
	InfineonTechs32Embed
	Elem14Dsp
	Lsi16Dsp
	Tms320C6000
	McstElbrusE2k
	Arm64bits
	ZilogZ80
	RiscV
	BerkleyPacketFilter
	WDC65C816
	LoongArch
)

func convMultiByteToSingleByte(data []byte, endianness bool) int {
	res := 0
	if endianness {
		// big endian
		slices.Reverse(data)
	}
	for dataI, dataByte := range data {
		res += int(dataByte) * int(math.Pow(16, float64(dataI*2)))
	}
	return res
}

func main() {
	data, err := os.ReadFile("hello")
	if err != nil {
		panic("Error: Couldn't read file")
	}
	isElf := false
	if bytes.Equal(data[0:4], elfMagicNumber) {
		isElf = true
	}
	// true - big endian
	// false - small endian
	bitFormat := false
	// true - 64
	// false - 32
	endianness := false
	instructionSet := NoInstruction
	entryPoint := 0
	sectionHeaderOffset := 0
	sectionHeaderLen := 0
	sectionHeaderSize := 0
	textSectionSize := 0
	if isElf {
		bitFormat = data[4] == 2
		endianness = data[5] == 2
		switch convMultiByteToSingleByte(data[18:20], endianness) {
		case 0:
			// 0x0
			instructionSet = NoInstruction
		case 1:
			// 0x1
			instructionSet = We32100
		case 2:
			// 0x2
			instructionSet = Sparc
		case 3:
			// 0x3
			instructionSet = x86
		case 4:
			// 0x4
			instructionSet = M68k
		case 5:
			// 0x5
			instructionSet = M88k
		case 6:
			// 0x6
			instructionSet = Imcu
		case 7:
			// 0x7
			instructionSet = I80860
		case 8:
			// 0x8
			instructionSet = Mips
		case 9:
			// 0x9
			instructionSet = IbmSystem370
		case 10:
			// 0xA
			instructionSet = Rs3000Le
		case 15:
			// 0xF
			instructionSet = HpPaRisc
		case 19:
			// 0x13
			instructionSet = I80960
		case 20:
			// 0x14
			instructionSet = PowerPc
		case 21:
			// 0x15
			instructionSet = PowerPc64b
		case 22:
			// 0x16
			instructionSet = S390
		case 23:
			// 0x17
			instructionSet = IbmSpu
		case 36:
			// 0x24
			instructionSet = NecV800
		case 37:
			// 0x25
			instructionSet = FujitsoFr20
		case 38:
			// 0x26
			instructionSet = TrwRh32
		case 39:
			// 0x27
			instructionSet = MotorolaRce
		case 40:
			// 0x28
			instructionSet = Arm
		case 41:
			// 0x29
			instructionSet = DigitalAlpha
		case 42:
			// 0x2A
			instructionSet = SuperH
		case 43:
			// 0x2B
			instructionSet = SparceV9
		case 44:
			// 0x2C
			instructionSet = STriCodeEmbedProc
		case 45:
			// 0x2D
			instructionSet = ArgonautRiscCore
		case 46:
			// 0x2E
			instructionSet = HitachiH8300
		case 47:
			// 0x2F
			instructionSet = HitachiH8300H
		case 48:
			// 0x30
			instructionSet = HitachiH8S
		case 49:
			// 0x31
			instructionSet = HitachiH8500
		case 50:
			// 0x32
			instructionSet = Ia64
		case 51:
			// 0x33
			instructionSet = StanfordMips
		case 52:
			// 0x34
			instructionSet = MotorolaColdFire
		case 53:
			// 0x35
			instructionSet = M68GC12
		case 54:
			// 0x36
			instructionSet = MmaMMAccel
		case 55:
			// 0x37
			instructionSet = SiemensPcp
		case 56:
			// 0x38
			instructionSet = SonynCPURiscProcessor
		case 57:
			// 0x39
			instructionSet = DensoNDR1MicroProcessor
		case 58:
			// 0x3A
			instructionSet = StarCoreProcessor
		case 59:
			// 0x3B
			instructionSet = Me16
		case 60:
			// 0x3C
			instructionSet = StmeSt100
		case 61:
			// 0x3D
			instructionSet = AdvLogCorp
		case 62:
			// 0x3E
			instructionSet = Amdx8664
		case 63:
			// 0x3F
			instructionSet = SonyDsp
		case 64:
			// 0x40
			instructionSet = Pdp10
		case 65:
			// 0x41
			instructionSet = Pdp11
		case 66:
			// 0x42
			instructionSet = Fx66
		case 67:
			// 0x43
			instructionSet = St9816MicroController
		case 68:
			// 0x44
			instructionSet = St8MicroController
		case 69:
			// 0x45
			instructionSet = MC68HC16
		case 70:
			// 0x46
			instructionSet = MC68HC11
		case 71:
			// 0x47
			instructionSet = MC68HC08
		case 72:
			// 0x48
			instructionSet = MC68HC05
		case 73:
			// 0x49
			instructionSet = Svx
		case 74:
			// 0x4A
			instructionSet = St198Microcontroller
		case 75:
			// 0x4B
			instructionSet = DVax
		case 76:
			// 0x4C
			instructionSet = AxisComms32Embed
		case 77:
			// 0x4D
			instructionSet = InfineonTechs32Embed
		case 78:
			// 0x4E
			instructionSet = Elem14Dsp
		case 79:
			// 0x4F
			instructionSet = Lsi16Dsp
		case 140:
			// 0x8C
			instructionSet = Tms320C6000
		case 175:
			// 0xAF
			instructionSet = McstElbrusE2k
		case 183:
			// 0xB7
			instructionSet = Arm64bits
		case 220:
			// 0xDC
			instructionSet = ZilogZ80
		case 243:
			// 0xF3
			instructionSet = RiscV
		case 247:
			// 0xF7
			instructionSet = BerkleyPacketFilter
		case 257:
			// 0x101
			instructionSet = WDC65C816
		case 258:
			// 0x102
			instructionSet = LoongArch
		}
		if bitFormat {
			// 64-bit
			entryPoint = convMultiByteToSingleByte(data[24:32], endianness)
			sectionHeaderOffset = convMultiByteToSingleByte(data[40:48], endianness)
			sectionHeaderLen = convMultiByteToSingleByte(data[60:62], endianness)
			sectionHeaderSize = convMultiByteToSingleByte(data[52:54], endianness)
		} else {
			// 32-bit
			entryPoint = convMultiByteToSingleByte(data[24:28], endianness)
			sectionHeaderOffset = convMultiByteToSingleByte(data[32:36], endianness)
			sectionHeaderLen = convMultiByteToSingleByte(data[48:50], endianness)
			sectionHeaderSize = convMultiByteToSingleByte(data[40:42], endianness)
		}
		intialSectionHeaderOffset := sectionHeaderOffset
		for range sectionHeaderLen {
			sectionOffset := 0
			if bitFormat {
				sectionOffset = convMultiByteToSingleByte(data[intialSectionHeaderOffset+16:intialSectionHeaderOffset+20], endianness)
			} else {
				sectionOffset = convMultiByteToSingleByte(data[intialSectionHeaderOffset+24:intialSectionHeaderOffset+32], endianness)
			}
			if sectionOffset == entryPoint {
				if bitFormat {
					textSectionSize = convMultiByteToSingleByte(data[intialSectionHeaderOffset+32:intialSectionHeaderOffset+40], endianness)
				} else {
					textSectionSize = convMultiByteToSingleByte(data[intialSectionHeaderOffset+20:intialSectionHeaderOffset+24], endianness)
				}
				break
			}
			intialSectionHeaderOffset += sectionHeaderSize
		}
		switch instructionSet {
		case NoInstruction:
		case We32100:
		case Sparc:
		case x86:
		case M68k:
		case M88k:
		case Imcu:
		case I80860:
		case Mips:
		case IbmSystem370:
		case Rs3000Le:
		case HpPaRisc:
		case I80960:
		case PowerPc:
		case PowerPc64b:
		case S390:
		case IbmSpu:
		case NecV800:
		case FujitsoFr20:
		case TrwRh32:
		case MotorolaRce:
		case Arm:
		case DigitalAlpha:
		case SuperH:
		case SparceV9:
		case STriCodeEmbedProc:
		case ArgonautRiscCore:
		case HitachiH8300:
		case HitachiH8300H:
		case HitachiH8S:
		case HitachiH8500:
		case Ia64:
		case StanfordMips:
		case MotorolaColdFire:
		case M68GC12:
		case MmaMMAccel:
		case SiemensPcp:
		case SonynCPURiscProcessor:
		case DensoNDR1MicroProcessor:
		case StarCoreProcessor:
		case Me16:
		case StmeSt100:
		case AdvLogCorp:
		case Amdx8664:
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
			for curByte := range data[entryPoint:textSectionSize] {
				// 1-15 bytes
				// prefix(optional)
				if isPrefix {
					// legacy prefix (up to 4 times)
					// rex prefix (0x40-0x4F) ! only 1 in a instruction
					// legacy escape sequences
					// vex prefix
					// xop prefix
					switch curByte {
					case 102:
						// 0x66
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isOperandSizeOverride = true
						legacePrefixCnt += 1
					case 103:
						// 0x67
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isAddressSizeOverride = true
						legacePrefixCnt += 1
					case 46:
						// 0x2E
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isCSSegmentSizeOverride = true
						legacePrefixCnt += 1
					case 62:
						// 0x3E
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isDSSegmentSizeOverride = true
						legacePrefixCnt += 1
					case 38:
						// 0x26
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isESSegmentSizeOverride = true
						legacePrefixCnt += 1
					case 100:
						// 0x64
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isFSSegmentSizeOverride = true
						legacePrefixCnt += 1
					case 101:
						// 0x65
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isGSSegmentSizeOverride = true
						legacePrefixCnt += 1
					case 54:
						// 0x36
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isSSSegmentSizeOverride = true
						legacePrefixCnt += 1
					case 240:
						// 0xf0
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isLock = true
						legacePrefixCnt += 1
					case 243:
						// 0xf3
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isRep1 = true
						legacePrefixCnt += 1
					case 242:
						// 0xf2
						if legacePrefixCnt == 4 {
							panic("Error: There can only be 4 legacy prefixes in 1 instruction")
						}
						isRep0 = true
						legacePrefixCnt += 1
					case 64:
						// 0x40
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
					case 65:
						// 0x41
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
					case 66:
						// 0x42
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
					case 67:
						// 0x43
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
					case 68:
						// 0x44
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
					case 69:
						// 0x45
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
					case 70:
						// 0x46
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
					case 71:
						// 0x47
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
					case 72:
						// 0x48
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
					case 73:
						// 0x49
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
					case 74:
						// 0x4A
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
					case 75:
						// 0x4B
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
					case 76:
						// 0x4C
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
					case 77:
						// 0x4D
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
					case 78:
						// 0x4E
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
					case 79:
						// 0x4F
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
					case 197:
						// 0xC5
						isVex = true
						isVex3Byte = false
						isPrefix = false
						isEscapeSequence = true
					case 196:
						// 0xC4
						isVex = true
						isVex3Byte = true
						isPrefix = false
						isEscapeSequence = true
					case 143:
						// 0x8F
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
						case 15:
							// 0xF
							if isSecondaryMap {
								is3dNow = true
								isSecondaryMap = false
								isEscapeSequence = false
								isOpcode = true
							} else {
								isSecondaryMap = true
							}
						case 56:
							// 0x38
							if isSecondaryMap {
								is38 = true
								isEscapeSequence = false
								isOpcode = true
							}
						case 58:
							// 0x3A
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
							} else if mapSelect[3] == true && mapSelect[4] == false {
								// 0x10
							} else if mapSelect[3] == true && mapSelect[4] == true {
								// 0x11
							}
						} else {
						}
					} else if isXop {
						if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == false {
							// 0x01
						} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == false && mapSelect[4] == true {
							// 0x10
						} else if mapSelect[1] == true && mapSelect[2] == false && mapSelect[3] == true && mapSelect[4] == false {
							// 0x11
						}
					} else if isSecondaryMap {
						if is3dNow {
						} else if is38 {
						} else if is3A {
						} else {
						}
					} else {
						if isLock {
						}
						if bitFormat {
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
					isDisplacement = false
				}
			}
		case SonyDsp:
		case Pdp10:
		case Pdp11:
		case Fx66:
		case St9816MicroController:
		case St8MicroController:
		case MC68HC16:
		case MC68HC11:
		case MC68HC08:
		case MC68HC05:
		case Svx:
		case St198Microcontroller:
		case DVax:
		case AxisComms32Embed:
		case InfineonTechs32Embed:
		case Elem14Dsp:
		case Lsi16Dsp:
		case Tms320C6000:
		case McstElbrusE2k:
		case Arm64bits:
		case ZilogZ80:
		case RiscV:
		case BerkleyPacketFilter:
		case WDC65C816:
		case LoongArch:
		}
	}
}

type Instruction int

const (
	AAA Instruction = iota
	AAD
	AAM
	AAS
	ADC
	ADCX
	ADD
	ADOX
	AND
	ANDN
	BEXTR
	BLCI
	BLCIC
	BLCMSK
	BLSR
	BOUND
	BSF
	BSR
	BSWAP
	BT
	BTC
	BTR
	BTS
	BZHI
	CALL
	CBW // 16 bit, CWDE 32 bit, CDQE 64 bit
	CWD // 16 bit, CDQ 32 bit, CQO 64 bit
	CLC
	CLD
	CLFLUSH
	CLFLUSHOPT
	CLWB
	CLZERO
	CMC
	CMOVcc
	CMP
	CMPS // 8 bit,  16 bit, CDQ 32 bit, CQO 64 bit
	CMPXCHG
	CMPXCHH8B // 16 - 32 bit, CMPXCHH16B 64 bit
	CPUID
	CRC32
	DAA
	DAS
	DEC
	DIV
	ENTER
	IDIV
	IMUL
	IN
	INC
	INS //
	INT
	INTO
	Jcc
	JCXZ // 16 bit, JECXZ 32 bit, JRCXZ 64 bit
	JMP
	LAHF
	LDS //
	LEA
	LEAVE
	LFENCE
	LLWPCB
	LODS
	LOOP
	LWPINS
	LWPVAL
	LZCNT
	MCOMMIT
	MFENCE
	MONITORX
	MOV
	MOVBE
	MOVD
	MOVMSKPD
	MOVMSKPS
	MOVNTI
	MOVS
	MOVSX
	MOVSXD
	MOVZX
	MUL
	MULX
	MWAITX
	NEG
	NOP
	NOT
	OR
	OUT
	OUTS
	PAUSE
	PDEP
	PEXT
	POP
	POPA
	POPCNT
	POPF
	PREFETCH
	PUSH
	PUSHA
	PUSHF
	RCL
	RCR
	RDFSBASE
	RDPID
	RDPRU
	RDRAND
	RDSEED
	RET
	ROL
	ROR
	RORX
	SAHF
	SAL
	SAR
	SARX
	SBB
	SCAS
	SETcc
	SFENCE
	SHL
	SHLD
	SHLX
	SHR
	SHRD
	SHRX
	SLWPCB
	STC
	STD
	STOS
	SUB
	T1MSKC
	TEST
	TZCNT
	TZMSK
	UD0
	WRFSBASE
	XADD
	XCHG
	XLAT
	XLATB
	XOR
	ARPL
	CLAC
	CLGI
	CLI
	CLTS
	CLRSSBSY
	HLT
	INCSSP
	INT3
	INVD
	INVLPG
	INVLPGA
	INVLPGB
	INVPCID
	IRET
	LAR
	LGDT
	LIDT
	LLDT
	LMSW
	LSL
	LTR
	MONITOR
	MWAIT
	PSMASH
	PVALIDATE
	RDMSR
	RDPKRU
	RDPMC
	RDSSP
	RDTSC
	RMPADJUST
	RMPQUERY
	RMPREAD
	RMPUPDATE
	RSM
	RSTORSSP
	SAVEPREVSSP
	SETSSBY
	SGDT
	SIDT
	SKINIT
	SLDT
	SMSW
	STAC
	STI
	STGI
	STR
	SWAPGS
	SYSCALL
	SYSENTER
	SYSEXIT
	SYSRET
	TLBSYNC
	VERR
	VERW
	VMLOAD
	VMMCALL
	VMGEXIT
	VMRUN
	VMSAVE
	WBINVD
	WBNOINVD
	WRMSR
	WRPKRU
	WRSS
	WRUSS
)

type OpcodeMap3B int

const (
	VexOM1 OpcodeMap3B = iota
	VexOM2
	VexOM3
	XopOM1
	XopOm2
	XopOm3
)
