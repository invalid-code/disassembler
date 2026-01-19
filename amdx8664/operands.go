package amdx8664

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
	AL Register = iota
	BL
	CL
	DL
	AH
	BH
	CH
	DH
	AX
	BX
	CX
	DX
	EAX
	EBX
	ECX
	EDX
	ESI
	EDI
	ESP
	EBP
	RAX
	RBX
	RCX
	RDX
	RSI
	RDI
	RSP
	RBP
	R8W
	R9W
	R10W
	R11W
	R12W
	R13W
	R14W
	R15W
	R8D
	R9D
	R10D
	R11D
	R12D
	R13D
	R14D
	R15D
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
	case AL:
		return "AL"
	case BL:
		return "BL"
	case CL:
		return "CL"
	case DL:
		return "DL"
	case AX:
		return "AX"
	case BX:
		return "BX"
	case CX:
		return "CX"
	case DX:
		return "DX"
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
