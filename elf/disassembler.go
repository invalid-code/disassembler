package elf

import (
	"bytes"
	"github.com/invalid-code/disassembler/amdx8664"
	"math"
	"slices"
)

var elfMagicNumber []byte = []uint8{0x7F, 0x45, 0x4C, 0x46}

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

func IsElf(data []byte) bool {
	return bytes.Equal(data, elfMagicNumber)
}

func Disassemble(data []byte) {
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
	bitFormat = data[0x4] == 2
	endianness = data[0x5] == 2
	switch convMultiByteToSingleByte(data[0x12:0x14], endianness) {
	case 0x0:
		instructionSet = NoInstruction
	case 0x1:
		instructionSet = We32100
	case 0x2:
		instructionSet = Sparc
	case 0x3:
		instructionSet = x86
	case 0x4:
		instructionSet = M68k
	case 0x5:
		instructionSet = M88k
	case 0x6:
		instructionSet = Imcu
	case 0x7:
		instructionSet = I80860
	case 0x8:
		instructionSet = Mips
	case 0x9:
		instructionSet = IbmSystem370
	case 0xA:
		instructionSet = Rs3000Le
	case 0xF:
		instructionSet = HpPaRisc
	case 0x13:
		instructionSet = I80960
	case 0x14:
		instructionSet = PowerPc
	case 0x15:
		instructionSet = PowerPc64b
	case 0x16:
		instructionSet = S390
	case 0x17:
		instructionSet = IbmSpu
	case 0x24:
		instructionSet = NecV800
	case 0x25:
		instructionSet = FujitsoFr20
	case 0x26:
		instructionSet = TrwRh32
	case 0x27:
		instructionSet = MotorolaRce
	case 0x28:
		instructionSet = Arm
	case 0x29:
		instructionSet = DigitalAlpha
	case 0x2A:
		instructionSet = SuperH
	case 0x2B:
		instructionSet = SparceV9
	case 0x2C:
		instructionSet = STriCodeEmbedProc
	case 0x2D:
		instructionSet = ArgonautRiscCore
	case 0x2E:
		instructionSet = HitachiH8300
	case 0x2F:
		instructionSet = HitachiH8300H
	case 0x30:
		instructionSet = HitachiH8S
	case 0x31:
		instructionSet = HitachiH8500
	case 0x32:
		instructionSet = Ia64
	case 0x33:
		instructionSet = StanfordMips
	case 0x34:
		instructionSet = MotorolaColdFire
	case 0x35:
		instructionSet = M68GC12
	case 0x36:
		instructionSet = MmaMMAccel
	case 0x37:
		instructionSet = SiemensPcp
	case 0x38:
		instructionSet = SonynCPURiscProcessor
	case 0x39:
		instructionSet = DensoNDR1MicroProcessor
	case 0x3A:
		instructionSet = StarCoreProcessor
	case 0x3B:
		instructionSet = Me16
	case 0x3C:
		instructionSet = StmeSt100
	case 0x3D:
		instructionSet = AdvLogCorp
	case 0x3E:
		instructionSet = Amdx8664
	case 0x3F:
		instructionSet = SonyDsp
	case 0x40:
		instructionSet = Pdp10
	case 0x41:
		instructionSet = Pdp11
	case 0x42:
		instructionSet = Fx66
	case 0x43:
		instructionSet = St9816MicroController
	case 0x44:
		instructionSet = St8MicroController
	case 0x45:
		instructionSet = MC68HC16
	case 0x46:
		instructionSet = MC68HC11
	case 0x47:
		instructionSet = MC68HC08
	case 0x48:
		instructionSet = MC68HC05
	case 0x49:
		instructionSet = Svx
	case 0x4A:
		instructionSet = St198Microcontroller
	case 0x4B:
		instructionSet = DVax
	case 0x4C:
		instructionSet = AxisComms32Embed
	case 0x4D:
		instructionSet = InfineonTechs32Embed
	case 0x4E:
		instructionSet = Elem14Dsp
	case 0x4F:
		instructionSet = Lsi16Dsp
	case 0x8C:
		instructionSet = Tms320C6000
	case 0xAF:
		instructionSet = McstElbrusE2k
	case 0xB7:
		instructionSet = Arm64bits
	case 0xDC:
		instructionSet = ZilogZ80
	case 0xF3:
		instructionSet = RiscV
	case 0xF7:
		instructionSet = BerkleyPacketFilter
	case 0x101:
		instructionSet = WDC65C816
	case 0x102:
		instructionSet = LoongArch
	}
	if bitFormat {
		// 64-bit
		entryPoint = convMultiByteToSingleByte(data[0x18:0x20], endianness)
		sectionHeaderOffset = convMultiByteToSingleByte(data[0x28:0x30], endianness)
		sectionHeaderSize = convMultiByteToSingleByte(data[0x34:0x36], endianness)
		sectionHeaderLen = convMultiByteToSingleByte(data[0x3C:0x3E], endianness)
	} else {
		// 32-bit
		entryPoint = convMultiByteToSingleByte(data[0x18:0x1C], endianness)
		sectionHeaderOffset = convMultiByteToSingleByte(data[0x20:0x24], endianness)
		sectionHeaderSize = convMultiByteToSingleByte(data[0x28:0x2A], endianness)
		sectionHeaderLen = convMultiByteToSingleByte(data[0x30:0x32], endianness)
	}
	intialSectionHeaderOffset := sectionHeaderOffset
	for range sectionHeaderLen {
		sectionOffset := 0
		if bitFormat {
			sectionOffset = convMultiByteToSingleByte(data[intialSectionHeaderOffset+0x10:intialSectionHeaderOffset+0x14], endianness)
		} else {
			sectionOffset = convMultiByteToSingleByte(data[intialSectionHeaderOffset+0x18:intialSectionHeaderOffset+0x20], endianness)
		}
		if sectionOffset == entryPoint {
			if bitFormat {
				textSectionSize = convMultiByteToSingleByte(data[intialSectionHeaderOffset+0x20:intialSectionHeaderOffset+0x28], endianness)
			} else {
				textSectionSize = convMultiByteToSingleByte(data[intialSectionHeaderOffset+0x14:intialSectionHeaderOffset+0x18], endianness)
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
		amdx8664.DisassembleBytes(data[entryPoint:entryPoint+textSectionSize], bitFormat)
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
