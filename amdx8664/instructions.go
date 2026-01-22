package amdx8664

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
	ARPL
	BEXTR
	BLCFILL
	BLCI
	BLCIC
	BLCMSK
	BLCS
	BLSFILL
	BLSI
	BLSIC
	BLSMK
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
	CBW
	CDQ
	CDQE
	CLAC
	CLC
	CLD
	CLFLUSH
	CLFLUSHOPT
	CLGI
	CLI
	CLRSSBSY
	CLTS
	CLWB
	CLZERO
	CMC
	CMOVA
	CMOVAE
	CMOVB
	CMOVBE
	CMOVC
	CMOVE
	CMOVG
	CMOVGE
	CMOVL
	CMOVLE
	CMOVNA
	CMOVNAE
	CMOVNB
	CMOVNBE
	CMOVNC
	CMOVNE
	CMOVNG
	CMOVNGE
	CMOVNL
	CMOVNLE
	CMOVNO
	CMOVNP
	CMOVNS
	CMOVNZ
	CMOVO
	CMOVP
	CMOVPE
	CMOVPO
	CMOVS
	CMOVZ
	CMP
	CMPS
	CMPSB
	CMPSD
	CMPSQ
	CMPSW
	CMPXCHG
	CMPXCHG16B
	CMPXCHG8B
	CPUID
	CQO
	CRC32
	CWD
	CWDE
	DAA
	DAS
	DEC
	DIV
	ENTER
	HLT
	IDIV
	IMUL
	IN
	INC
	INCSSPD
	INCSSPQ
	INS
	INSB
	INSD
	INSW
	INT
	INT3
	INTO
	INVD
	INVLPG
	INVLPGA
	INVLPGB
	INVPCID
	IRET
	IRETD
	IRETQ
	JA
	JAE
	JB
	JBE
	JC
	JCXZ
	JE
	JECXZ
	JG
	JGE
	JL
	JLE
	JMP
	JNAE
	JNB
	JNBE
	JNC
	JNG
	JNGE
	JNL
	JNLE
	JNO
	JNP
	JNS
	JNZ
	JO
	JP
	JPE
	JPO
	JRCXZ
	JS
	JZ
	LAHF
	LAR
	LDS
	LEA
	LEAVE
	LES
	LFENCE
	LFS
	LGDT
	LGS
	LIDT
	LLDT
	LLWPCB
	LMSW
	LODS
	LODSB
	LODSD
	LODSQ
	LODSW
	LOOP
	LOOPE
	LOOPNE
	LOOPNZ
	LOOPZ
	LSL
	LSS
	LTR
	LWPINS
	LWPVAL
	LZCNT
	MCOMMIT
	MFENCE
	MONITOR
	MONITORX
	MOV
	MOVBE
	MOVD
	MOVDIR64B
	MOVDIRI
	MOVMSKPD
	MOVMSKPS
	MOVNTI
	MOVS
	MOVSB
	MOVSD
	MOVSQ
	MOVSW
	MOVSX
	MOVSXD
	MOVZX
	MUL
	MULX
	MWAIT
	MWAITX
	NEG
	NOP
	NOT
	OR
	OUT
	OUTS
	OUTSB
	OUTSD
	OUTSW
	PAUSE
	PDEP
	PEXT
	POP
	POPA
	POPAD
	POPCNT
	POPF
	POPFD
	POPFQ
	PREFETCH
	PREFETCHIT0
	PREFETCHIT1
	PREFETCHNTA
	PREFETCHT0
	PREFETCHT1
	PREFETCHT2
	PREFETCHW
	PSMASH
	PUSH
	PUSHA
	PUSHAD
	PUSHF
	PUSHFD
	PUSHFQ
	PVALIDATE
	RCL
	RCR
	RDFSBASE
	RDMSR
	RDPID
	RDPKRU
	RDPMC
	RDPRU
	RDRAND
	RDSEED
	RDSSPD
	RDSSPQ
	RDTSC
	RDTSCP
	RET
	RMPADJUST
	RMPQUERY
	RMPREAD
	RMPUPDATE
	ROL
	ROR
	RORX
	RSM
	RSTORSSP
	SAHF
	SAL
	SAR
	SARX
	SAVEPREVSSP
	SBB
	SCAS
	SCASB
	SCASD
	SCASQ
	SCASW
	SETA
	SETAE
	SETB
	SETBE
	SETC
	SETE
	SETG
	SETGE
	SETL
	SETLE
	SETNA
	SETNAE
	SETNB
	SETNBE
	SETNC
	SETNE
	SETNG
	SETNGE
	SETNL
	SETNLE
	SETNO
	SETNP
	SETNS
	SETNZ
	SETO
	SETP
	SETPE
	SETPO
	SETS
	SETSSBSY
	SETZ
	SFENCE
	SGDT
	SHL
	SHLD
	SHLX
	SHR
	SHRD
	SHRX
	SIDT
	SKINIT
	SLDT
	SLWPCB
	SMSW
	STAC
	STC
	STD
	STGI
	STI
	STOS
	STOSB
	STOSD
	STOSQ
	STOSW
	STR
	SUB
	SWAPGS
	SYSCALL
	SYSENTER
	SYSEXIT
	SYSRET
	T1MSKC
	TEST
	TLBSYNC
	TZCNT
	TZMSK
	UD0
	UD1
	UD2
	VERR
	VERW
	VMGEXIT
	VMLOAD
	VMMCALL
	VMRUN
	VMSAVE
	WBINVD
	WBNOINVD
	WRFSBASE
	WRMSR
	WRPKRU
	WRSS
	WRSSQ
	WRUSSD
	WRUSSQ
	XADD
	XCHG
	XLAT
	XOR
)

func (instruction Instruction) String() string {
	switch instruction {
	case AAA:
		return "AAA"
	case AAD:
		return "AAD"
	case AAM:
		return "AAM"
	case AAS:
		return "AAS"
	case ADC:
		return "ADC"
	case ADCX:
		return "ADCX"
	case ADD:
		return "ADD"
	case ADOX:
		return "ADOX"
	case AND:
		return "AND"
	case ANDN:
		return "ANDN"
	case ARPL:
		return "ARPL"
	case BEXTR:
		return "BEXTR"
	case BLCFILL:
		return "BLCFILL"
	case BLCI:
		return "BLCI"
	case BLCIC:
		return "BLCIC"
	case BLCMSK:
		return "BLCMSK"
	case BLCS:
		return "BLCS"
	case BLSFILL:
		return "BLSFILL"
	case BLSI:
		return "BLSI"
	case BLSIC:
		return "BLSIC"
	case BLSMK:
		return "BLSMK"
	case BLSR:
		return "BLSR"
	case BOUND:
		return "BOUND"
	case BSF:
		return "BSF"
	case BSR:
		return "BSR"
	case BSWAP:
		return "BSWAP"
	case BT:
		return "BT"
	case BTC:
		return "BTC"
	case BTR:
		return "BTR"
	case BTS:
		return "BTS"
	case BZHI:
		return "BZHI"
	case CALL:
		return "CALL"
	case CBW:
		return "CBW"
	case CDQ:
		return "CDQ"
	case CDQE:
		return "CDQE"
	case CLAC:
		return "CLAC"
	case CLC:
		return "CLC"
	case CLD:
		return "CLD"
	case CLFLUSH:
		return "CLFLUSH"
	case CLFLUSHOPT:
		return "CLFLUSHOPT"
	case CLGI:
		return "CLGI"
	case CLI:
		return "CLI"
	case CLRSSBSY:
		return "CLRSSBSY"
	case CLTS:
		return "CLTS"
	case CLWB:
		return "CLWB"
	case CLZERO:
		return "CLZERO"
	case CMC:
		return "CMC"
	case CMOVA:
		return "CMOVA"
	case CMOVAE:
		return "CMOVAE"
	case CMOVB:
		return "CMOVB"
	case CMOVBE:
		return "CMOVBE"
	case CMOVC:
		return "CMOVC"
	case CMOVE:
		return "CMOVE"
	case CMOVG:
		return "CMOVG"
	case CMOVGE:
		return "CMOVGE"
	case CMOVL:
		return "CMOVL"
	case CMOVLE:
		return "CMOVLE"
	case CMOVNA:
		return "CMOVNA"
	case CMOVNAE:
		return "CMOVNAE"
	case CMOVNB:
		return "CMOVNB"
	case CMOVNBE:
		return "CMOVNBE"
	case CMOVNC:
		return "CMOVNC"
	case CMOVNE:
		return "CMOVNE"
	case CMOVNG:
		return "CMOVNG"
	case CMOVNGE:
		return "CMOVNGE"
	case CMOVNL:
		return "CMOVNL"
	case CMOVNLE:
		return "CMOVNLE"
	case CMOVNO:
		return "CMOVNO"
	case CMOVNP:
		return "CMOVNP"
	case CMOVNS:
		return "CMOVNS"
	case CMOVNZ:
		return "CMOVNZ"
	case CMOVO:
		return "CMOVO"
	case CMOVP:
		return "CMOVP"
	case CMOVPE:
		return "CMOVPE"
	case CMOVPO:
		return "CMOVPO"
	case CMOVS:
		return "CMOVS"
	case CMOVZ:
		return "CMOVZ"
	case CMP:
		return "CMP"
	case CMPS:
		return "CMPS"
	case CMPSB:
		return "CMPSB"
	case CMPSD:
		return "CMPSD"
	case CMPSQ:
		return "CMPSQ"
	case CMPSW:
		return "CMPSW"
	case CMPXCHG:
		return "CMPXCHG"
	case CMPXCHG16B:
		return "CMPXCHG16B"
	case CMPXCHG8B:
		return "CMPXCHG8B"
	case CPUID:
		return "CPUID"
	case CQO:
		return "CQO"
	case CRC32:
		return "CRC32"
	case CWD:
		return "CWD"
	case CWDE:
		return "CWDE"
	case DAA:
		return "DAA"
	case DAS:
		return "DAS"
	case DEC:
		return "DEC"
	case DIV:
		return "DIV"
	case ENTER:
		return "ENTER"
	case HLT:
		return "HLT"
	case IDIV:
		return "IDIV"
	case IMUL:
		return "IMUL"
	case IN:
		return "IN"
	case INC:
		return "INC"
	case INCSSPD:
		return "INCSSPD"
	case INCSSPQ:
		return "INCSSPQ"
	case INS:
		return "INS"
	case INSB:
		return "INSB"
	case INSD:
		return "INSD"
	case INSW:
		return "INSW"
	case INT:
		return "INT"
	case INT3:
		return "INT3"
	case INTO:
		return "INTO"
	case INVD:
		return "INVD"
	case INVLPG:
		return "INVLPG"
	case INVLPGA:
		return "INVLPGA"
	case INVLPGB:
		return "INVLPGB"
	case INVPCID:
		return "INVPCID"
	case IRET:
		return "IRET"
	case IRETD:
		return "IRETD"
	case IRETQ:
		return "IRETQ"
	case JA:
		return "JA"
	case JAE:
		return "JEA"
	case JB:
		return "JB"
	case JBE:
		return "JBE"
	case JC:
		return "JC"
	case JCXZ:
		return "JCXZ"
	case JE:
		return "JE"
	case JECXZ:
		return "JECXZ"
	case JG:
		return "JG"
	case JGE:
		return "JGE"
	case JL:
		return "JL"
	case JLE:
		return "JLE"
	case JMP:
		return "JMP"
	case JNAE:
		return "JNAE"
	case JNB:
		return "JNB"
	case JNBE:
		return "JNBE"
	case JNC:
		return "JNC"
	case JNG:
		return "JNG"
	case JNGE:
		return "JNGE"
	case JNL:
		return "JNL"
	case JNLE:
		return "JNLE"
	case JNO:
		return "JNO"
	case JNP:
		return "JNP"
	case JNS:
		return "JNS"
	case JNZ:
		return "JNZ"
	case JO:
		return "JO"
	case JP:
		return "JP"
	case JPE:
		return "JPE"
	case JPO:
		return "JPO"
	case JRCXZ:
		return "JRCXZ"
	case JS:
		return "JS"
	case JZ:
		return "JZ"
	case LAHF:
		return "LAHF"
	case LAR:
		return "LAR"
	case LDS:
		return "LDS"
	case LEA:
		return "LEA"
	case LEAVE:
		return "LEAVE"
	case LES:
		return "LES"
	case LFENCE:
		return "LFENCE"
	case LFS:
		return "LFS"
	case LGDT:
		return "LGDT"
	case LGS:
		return "LGS"
	case LIDT:
		return "LIDT"
	case LLDT:
		return "LLDT"
	case LLWPCB:
		return "LLWPCB"
	case LMSW:
		return "LMSW"
	case LODS:
		return "LODS"
	case LODSB:
		return "LODSB"
	case LODSD:
		return "LODSD"
	case LODSQ:
		return "LODSQ"
	case LODSW:
		return "LODSW"
	case LOOP:
		return "LOOP"
	case LOOPE:
		return "LOOPE"
	case LOOPNE:
		return "LOOPNE"
	case LOOPNZ:
		return "LOOPNZ"
	case LOOPZ:
		return "LOOPZ"
	case LSL:
		return "LSL"
	case LSS:
		return "LSS"
	case LTR:
		return "LTR"
	case LWPINS:
		return "LWPINS"
	case LWPVAL:
		return "LWPVAL"
	case LZCNT:
		return "LZCNT"
	case MCOMMIT:
		return "MCOMMIT"
	case MFENCE:
		return "MFENCE"
	case MONITOR:
		return "MONITOR"
	case MONITORX:
		return "MONITORX"
	case MOV:
		return "MOV"
	case MOVBE:
		return "MOVBE"
	case MOVD:
		return "MOVD"
	case MOVDIR64B:
		return "MOVDIR64B"
	case MOVDIRI:
		return "MODIRI"
	case MOVMSKPD:
		return "MOVMSKPD"
	case MOVMSKPS:
		return "MOVMSPKS"
	case MOVNTI:
		return "MOVNTI"
	case MOVS:
		return "MOVS"
	case MOVSB:
		return "MOVSB"
	case MOVSD:
		return "MOVSD"
	case MOVSQ:
		return "MOVSQ"
	case MOVSW:
		return "MOVSW"
	case MOVSX:
		return "MOVSX"
	case MOVSXD:
		return "MOVSXD"
	case MOVZX:
		return "MOVZX"
	case MUL:
		return "MUL"
	case MULX:
		return "MULX"
	case MWAIT:
		return "MWAIT"
	case MWAITX:
		return "MWAITX"
	case NEG:
		return "NEG"
	case NOP:
		return "NOP"
	case NOT:
		return "NOT"
	case OR:
		return "OR"
	case OUT:
		return "OUT"
	case OUTS:
		return "OUTS"
	case OUTSB:
		return "OUTSB"
	case OUTSD:
		return "OUTSD"
	case OUTSW:
		return "OUTSW"
	case PAUSE:
		return "PAUSE"
	case PDEP:
		return "PDEP"
	case PEXT:
		return "PEXT"
	case POP:
		return "POP"
	case POPA:
		return "POPA"
	case POPAD:
		return "POPAD"
	case POPCNT:
		return "POPCNT"
	case POPF:
		return "POPF"
	case POPFD:
		return "POPFD"
	case POPFQ:
		return "POPFQ"
	case PREFETCH:
		return "PREFETCH"
	case PREFETCHIT0:
		return "PREFETCH10"
	case PREFETCHIT1:
		return "PREFETCH1"
	case PREFETCHNTA:
		return "PREFETCHNTA"
	case PREFETCHT0:
		return "PREFETCH0"
	case PREFETCHT1:
		return "PREFECTCH1"
	case PREFETCHT2:
		return "PREFETCH2"
	case PREFETCHW:
		return "PREFETCHW"
	case PSMASH:
		return "PSMASH"
	case PUSH:
		return "PUSH"
	case PUSHA:
		return "PUSHA"
	case PUSHAD:
		return "PUSHAD"
	case PUSHF:
		return "PUSHF"
	case PUSHFD:
		return "PUSHFD"
	case PUSHFQ:
		return "PUSHFQ"
	case PVALIDATE:
		return "PVALIDATE"
	case RCL:
		return "RCL"
	case RCR:
		return "RCR"
	case RDFSBASE:
		return "RDFBASE"
	case RDMSR:
		return "RDMSR"
	case RDPID:
		return "RDPID"
	case RDPKRU:
		return "RDPKRU"
	case RDPMC:
		return "RDPMC"
	case RDPRU:
		return "RDPRU"
	case RDRAND:
		return "RDRAND"
	case RDSEED:
		return "RDSEED"
	case RDSSPD:
		return "RDSSPD"
	case RDSSPQ:
		return "RDSSQ"
	case RDTSC:
		return "RDTSC"
	case RDTSCP:
		return "RDTSCP"
	case RET:
		return "RET"
	case RMPADJUST:
		return "RMPADJUST"
	case RMPQUERY:
		return "RMPQUERY"
	case RMPREAD:
		return "RMPREAD"
	case RMPUPDATE:
		return "RMUPDATE"
	case ROL:
		return "ROL"
	case ROR:
		return "ROR"
	case RORX:
		return "RORX"
	case RSM:
		return "RSH"
	case RSTORSSP:
		return "RSTORSSP"
	case SAHF:
		return "SAHF"
	case SAL:
		return "SAL"
	case SAR:
		return "SAR"
	case SARX:
		return "SARX"
	case SAVEPREVSSP:
		return "SAVEPREVSP"
	case SBB:
		return "SBB"
	case SCAS:
		return "SCAS"
	case SCASB:
		return "SCASB"
	case SCASD:
		return "SCASD"
	case SCASQ:
		return "SCASQ"
	case SCASW:
		return "SCASW"
	case SETA:
		return "SETA"
	case SETAE:
		return "SETEA"
	case SETB:
		return "SETB"
	case SETBE:
		return "SETBE"
	case SETC:
		return "SETC"
	case SETE:
		return "SETE"
	case SETG:
		return "SETG"
	case SETGE:
		return "SETGE"
	case SETL:
		return "SETL"
	case SETLE:
		return "SETLE"
	case SETNA:
		return "SETNA"
	case SETNAE:
		return "SETNAE"
	case SETNB:
		return "SETNB"
	case SETNBE:
		return "SETNBE"
	case SETNC:
		return "SETNC"
	case SETNE:
		return "SETNE"
	case SETNG:
		return "SETNG"
	case SETNGE:
		return "SETNGE"
	case SETNL:
		return "SETNL"
	case SETNLE:
		return "SETNLE"
	case SETNO:
		return "SETNO"
	case SETNP:
		return "SETNP"
	case SETNS:
		return "SETNS"
	case SETNZ:
		return "SETNZ"
	case SETO:
		return "SETO"
	case SETP:
		return "SETP"
	case SETPE:
		return "SETPE"
	case SETPO:
		return "SETPO"
	case SETS:
		return "SETS"
	case SETSSBSY:
		return "SETSSBY"
	case SETZ:
		return "SETZ"
	case SFENCE:
		return "SFENCE"
	case SGDT:
		return "SGDT"
	case SHL:
		return "SHL"
	case SHLD:
		return "SHLD"
	case SHLX:
		return "SHLX"
	case SHR:
		return "SHR"
	case SHRD:
		return "SHRD"
	case SHRX:
		return "SHRX"
	case SIDT:
		return "SIDT"
	case SKINIT:
		return "SKINIT"
	case SLDT:
		return "SLDT"
	case SLWPCB:
		return "SLWPCB"
	case SMSW:
		return "SMSW"
	case STAC:
		return "STAC"
	case STC:
		return "STC"
	case STD:
		return "STD"
	case STGI:
		return "STGI"
	case STI:
		return "STI"
	case STOS:
		return "STOS"
	case STOSB:
		return "STOSB"
	case STOSD:
		return "STOSD"
	case STOSQ:
		return "STOSQ"
	case STOSW:
		return "STOSW"
	case STR:
		return "STR"
	case SUB:
		return "SUB"
	case SWAPGS:
		return "SWAPGS"
	case SYSCALL:
		return "SYSCALL"
	case SYSENTER:
		return "SYSENTER"
	case SYSEXIT:
		return "SYSEXIT"
	case SYSRET:
		return "SYSRET"
	case T1MSKC:
		return "T1MSKC"
	case TEST:
		return "TEST"
	case TLBSYNC:
		return "TLBSYNC"
	case TZCNT:
		return "TZCNT"
	case TZMSK:
		return "TZMSK"
	case UD0:
		return "UD0"
	case UD1:
		return "UD1"
	case UD2:
		return "UD2"
	case VERR:
		return "VERR"
	case VERW:
		return "VERW"
	case VMGEXIT:
		return "VMGEXIT"
	case VMLOAD:
		return "VMLOAD"
	case VMMCALL:
		return "VMMCALL"
	case VMRUN:
		return "VMRUN"
	case VMSAVE:
		return "VMSAVE"
	case WBINVD:
		return "VBINVD"
	case WBNOINVD:
		return "WBNOINVD"
	case WRFSBASE:
		return "WRFSBASE"
	case WRMSR:
		return "WRMSR"
	case WRPKRU:
		return "WRPKRU"
	case WRSS:
		return "WRSS"
	case WRSSQ:
		return "WRSSQ"
	case WRUSSD:
		return "WRUSSD"
	case WRUSSQ:
		return "WRUSSQ"
	case XADD:
		return "XADD"
	case XCHG:
		return "XCHG"
	case XLAT:
		return "XLAT"
	case XOR:
		return "XOR"
	default:
		return "Unknown instruction"
	}
}
