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
	ADDPD
	ADDPS
	ADDSD
	ADDSS
	ADDSUBPD
	ADDSUBPS
	ADOX
	AESDEC
	AESDECLAST
	AESENC
	AESENCLAST
	AESIMC
	AND
	ANDNPD
	ANDNPS
	ANDPD
	ANDPS
	ARPL
	BLENDVPS
	BOUND
	BSF
	BSR
	BSWAP
	BTR
	BT
	BTC
	BTS
	CALL
	CLC
	CLD
	CLI
	CLTS
	CMC
	CMOVB
	CMOVBE
	CMOVL
	CMOVLE
	CMOVNB
	CMOVNBE
	CMOVNL
	CMOVNLE
	CMOVNO
	CMOVNP
	CMOVNS
	CMOVNZ
	CMOVO
	CMOVP
	CMOVS
	CMOVZ
	CMP
	CMPPD
	CMPPS
	CMPSB
	CMPSD
	CMPXCHG
	COMISD
	COMISS
	CPUID
	CRC32
	CVTDQ2PD
	CVTDQ2PS
	CVTPD2DQ
	CVTPD2PI
	CVTPD2PS
	CVTPI2PD
	CVTPI2PS
	CVTPS2DQ
	CVTPS2PD
	CVTPS2PI
	CVTSD2SI
	CVTSD2SS
	CVTSI2SD
	CVTSI2SS
	CVTSS2SD
	CVTSS2SI
	CVTTPD2DQ
	CVTTPD2PI
	CVTTPS2DQ
	CVTTPS2PI
	CVTTSD2SI
	CVTTSS2SI
	DAA
	DAS
	DEC
	DIV
	DIVPD
	DIVPS
	DIVSD
	DIVSS
	DPPD
	DPPS
	EMMS
	ENTER
	EXTRACTPS
	EXTRQ
	FEMMS
	HADDPD
	HADDPS
	HLT
	HRESET
	HSUBPD
	HSUBPS
	IMUL
	IN
	INC
	INSB
	INSERTPS
	INSERTQ
	INT
	INT1
	INT3
	INTO
	INVD
	JMP
	JB
	JBE
	JL
	JLE
	JNB
	JNBE
	JNL
	JNLE
	JNO
	JNP
	JNS
	JNZ
	JO
	JP
	JrCXZ
	JS
	JZ
	LAHF
	LAR
	LDDQU
	LDS
	LEA
	LEAVE
	LES
	LFS
	LGDT
	LGS
	LODSB
	LOOP
	LSL
	LSS
	LZCNT
	MASKMOVDQU
	MASKMOVQ
	MAXPD
	MAXPS
	MAXSD
	MAXSS
	MINPD
	MINPS
	MINSD
	MINSS
	MOV
	MOVAPD
	MOVAPS
	MOVBE
	MOVD
	MOVDDUP
	MOVDQ2Q
	MOVDQA
	MOVDQU
	MOVHPD
	MOVHPS
	MOVLPD
	MOVLPS
	MOVMSKPD
	MOVMSKPS
	MOVNTDQ
	MOVNTDQA
	MOVNTI
	MOVNTPD
	MOVNTPS
	MOVNTQ
	MOVNTSS
	MOVNTSD
	MOVQ
	MOVQ2DQ
	MOVSB
	MOVSD
	MOVSHDUP
	MOVSLDUP
	MOVSS
	MOVSX
	MOVUPD
	MOVUPS
	MOVZX
	MPSADBW
	MULPD
	MULPS
	MULSD
	MULSS
	NOP
	OR
	ORPD
	ORPS
	OUT
	OUTSB
	PABSB
	PABSD
	PABSQ
	PABSW
	PACKSSDW
	PACKSSWB
	PACKUSDW
	PACKUSWB
	PADDB
	PADDD
	PADDQ
	PADDSB
	PADDSW
	PADDUSB
	PADDUSW
	PADDW
	PALIGNR
	PAND
	PANDN
	PAUSE
	PAVGB
	PAVGW
	PBLENDVB
	PBLENDW
	PCLMULQDQ
	PCMPEQB
	PCMPEQD
	PCMPEQQ
	PCMPEQW
	PCMPGTB
	PCMPGTD
	PCMPGTQ
	PCMPGTW
	PEXTRB
	PEXTRD
	PEXTRW
	PHADDD
	PHADDSW
	PHADDW
	PHMINPOSUW
	PHSUBD
	PHSUBSW
	PHSUBW
	PINSRB
	PINSRW
	PI2FD
	PFRSQIT1
	PMADDWD
	PMADDUBSW
	PMAXSB
	PMAXSD
	PMAXSW
	PMAXUB
	PMAXUD
	PMAXUW
	PMINSB
	PMINSD
	PMINSW
	PMINUB
	PMINUD
	PMINUW
	PMOVMSKB
	PMOVSXBW
	PMOVSXBD
	PMOVSXBQ
	PMOVSXWD
	PMOVSXWQ
	PMOVSXDQ
	PMOVZXBD
	PMOVZXBQ
	PI2FW
	PF2IW
	PF2ID
	PFNACC
	PFPNACC
	PFCMPGE
	PFMIN
	PFRCP
	PFRSQRT
	PFSUB
	PFADD
	PFCMPGT
	PFMAX
	PFRCPIT1
	PFSUBR
	PFACC
	PFCMPEQ
	PFMUL
	PFRCPIT2
	PSWAPD
	PAVGUSB
	POPD
	PMULHRW
	PMOVZXBW
	PMOVZXDQ
	PMOVZXWD
	PMOVZXWQ
	PMULDQ
	PMULHRSW
	PMULHUW
	PMULHW
	PMULLD
	PMULLQ
	PMULLW
	PMULUDQ
	POP
	POPA
	POPAD
	POPCNT
	POR
	PSADBW
	PSHUFB
	PSHUFD
	PSHUFHW
	PSHUFLW
	PSHUFW
	PSIGNB
	PSIGND
	PSIGNW
	PSLLD
	PSLLDQ
	PSLLQ
	PSLLW
	PSRAD
	PSRAQ
	PSRAW
	PSRLD
	PSRLDQ
	PSRLQ
	PSRLW
	PSUBB
	PSUBD
	PSUBQ
	PSUBSB
	PSUBSW
	PSUBUSB
	PSUBUSW
	PSUBW
	PTEST
	PTWRITE
	PUNPCKHBW
	PUNPCKHDQ
	PUNPCKHQDQ
	PUNPCKHWD
	PUNPCKLBW
	PUNPCKLDQ
	PUNPCKLQDQ
	PUNPCKLWD
	PUSH
	PUSHA
	PUSHAD
	PUSHD
	PXOR
	RCPPS
	RCPSS
	RDMSR
	RDPMC
	RDTSC
	RET
	RSM
	RSQRTPS
	RSQRTSS
	SAHF
	SALC
	SBB
	SCASB
	SHLD
	SHRD
	SHUFPD
	SETO
	SETB
	SETNB
	SETZ
	SETBE
	SETNBE
	SETS
	SETNS
	SETP
	SETNP
	SETL
	SETNL
	SETLE
	SETNLE
	SETNZ
	SETNO
	SHUFPS
	SQRTPD
	SQRTPS
	SQRTSD
	SQRTSS
	STC
	STD
	STI
	STOSB
	SUB
	SUBPD
	SUBPS
	SUBSD
	SUBSS
	SYSCALL
	SYSENTER
	SYSEXIT
	SYSRET
	TEST
	TZCNT
	UCOMISD
	UD0
	UD2
	UCOMISS
	UNPCKHPD
	UNPCKHPS
	UNPCKLPD
	UNPCKLPS
	WRMSR
	WRUSS
	XADD
	XCHG
	XLAT
	XOR
	XORPD
	XORPS
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
	case ADCX:
		return "ADCX"
	case ADC:
		return "ADC"
	case ADD:
		return "ADD"
	case ADDPD:
		return "ADDPD"
	case ADDPS:
		return "ADDPS"
	case ADDSD:
		return "ADDSD"
	case ADDSS:
		return "ADDSS"
	case ADDSUBPD:
		return "ADDSUBPD"
	case ADDSUBPS:
		return "ADDSUBPS"
	case ADOX:
		return "ADOX"
	case AESDEC:
		return "AESDEC"
	case AESDECLAST:
		return "AESDECLAST"
	case AESENC:
		return "AESENC"
	case AESENCLAST:
		return "AESENCLAST"
	case AESIMC:
		return "AESIMC"
	case AND:
		return "AND"
	case ANDNPD:
		return "ANDPD"
	case ANDNPS:
		return "ANDNPS"
	case ANDPD:
		return "ANDPD"
	case ANDPS:
		return "ANDPS"
	case ARPL:
		return "ARPL"
	case BLENDVPS:
		return "BLENDVPS"
	case BOUND:
		return "BOUND"
	case BSF:
		return "BSF"
	case BSR:
		return "BSR"
	case BSWAP:
		return "BSWAP"
	case BTR:
		return "BTR"
	case BT:
		return "BT"
	case BTC:
		return "BTC"
	case BTS:
		return "BTS"
	case CALL:
		return "CALL"
	case CLC:
		return "CLC"
	case CLD:
		return "CLD"
	case CLI:
		return "CLI"
	case CLTS:
		return "CLTS"
	case CMC:
		return "CMC"
	case CMOVB:
		return "CMOVB"
	case CMOVBE:
		return "CMOVBE"
	case CMOVL:
		return "CMOVL"
	case CMOVLE:
		return "CMOVLE"
	case CMOVNB:
		return "CMOVNB"
	case CMOVNBE:
		return "CMOVNBE"
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
	case CMOVS:
		return "CMOVS"
	case CMOVZ:
		return "CMOVZ"
	case CMP:
		return "CMP"
	case CMPPD:
		return "CMPPD"
	case CMPPS:
		return "CMPPS"
	case CMPSB:
		return "CMPSB"
	case CMPSD:
		return "CMPSD"
	case CMPXCHG:
		return "CMPXCHG"
	case COMISD:
		return "COMISD"
	case COMISS:
		return "COMISS"
	case CPUID:
		return "CPUID"
	case CRC32:
		return "CRC32"
	case CVTDQ2PD:
		return "CVTDQ2PD"
	case CVTDQ2PS:
		return "CVTDQ2PS"
	case CVTPD2DQ:
		return "CVTPD2DQ"
	case CVTPD2PI:
		return "CVTPD2PI"
	case CVTPD2PS:
		return "CVTPD2PS"
	case CVTPI2PD:
		return "CVTPI2PD"
	case CVTPI2PS:
		return "CVTPI2PS"
	case CVTPS2DQ:
		return "CVTPS2DQ"
	case CVTPS2PD:
		return "CVTPS2PD"
	case CVTPS2PI:
		return "CVTPS2PI"
	case CVTSD2SI:
		return "CVTSD2SI"
	case CVTSD2SS:
		return "CVTSD2SS"
	case CVTSI2SD:
		return "CVTSI2SD"
	case CVTSI2SS:
		return "CVTSI2SS"
	case CVTSS2SD:
		return "CVTSS2SD"
	case CVTSS2SI:
		return "CVTSS2SI"
	case CVTTPD2DQ:
		return "CVTTPD2DQ"
	case CVTTPD2PI:
		return "CVTTPD2PI"
	case CVTTPS2DQ:
		return "CVTTPS2DQ"
	case CVTTPS2PI:
		return "CVTTPS2PI"
	case CVTTSD2SI:
		return "CVTTSD2SI"
	case CVTTSS2SI:
		return "CVTTSS2SI"
	case DAA:
		return "DAA"
	case DAS:
		return "DAS"
	case DEC:
		return "DEC"
	case DIV:
		return "DIV"
	case DIVPD:
		return "DIVPD"
	case DIVPS:
		return "DIVPS"
	case DIVSD:
		return "DIVSD"
	case DIVSS:
		return "DIVSS"
	case DPPD:
		return "DPPD"
	case DPPS:
		return "DPPS"
	case EMMS:
		return "EMMS"
	case ENTER:
		return "ENTER"
	case EXTRACTPS:
		return "EXTRACTPS"
	case EXTRQ:
		return "EXTRQ"
	case FEMMS:
		return "FEMMS"
	case HADDPD:
		return "HADDPD"
	case HADDPS:
		return "HADDPS"
	case HLT:
		return "HLT"
	case HRESET:
		return "HRESET"
	case HSUBPD:
		return "HSUBPD"
	case HSUBPS:
		return "HSUBPD"
	case IMUL:
		return "HSUBPD"
	case IN:
		return "IN"
	case INC:
		return "INC"
	case INSB:
		return "INSB"
	case INSERTPS:
		return "INSERTPS"
	case INSERTQ:
		return "INSERTQ"
	case INT:
		return "INT"
	case INT1:
		return "INT1"
	case INT3:
		return "INT3"
	case INTO:
		return "INTO"
	case INVD:
		return "INVD"
	case JMP:
		return "JMP"
	case JB:
		return "JB"
	case JBE:
		return "JBE"
	case JL:
		return "JL"
	case JLE:
		return "JLE"
	case JNB:
		return "JNB"
	case JNBE:
		return "JNBE"
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
	case JrCXZ:
		return "JrCXZ"
	case JS:
		return "JS"
	case JZ:
		return "JZ"
	case LAHF:
		return "LAHF"
	case LAR:
		return "LAR"
	case LDDQU:
		return "LDDQU"
	case LDS:
		return "LDS"
	case LEA:
		return "LEA"
	case LEAVE:
		return "LEAVE"
	case LES:
		return "LES"
	case LFS:
		return "LFS"
	case LGDT:
		return "LGDT"
	case LGS:
		return "LGS"
	case LODSB:
		return "LODSB"
	case LOOP:
		return "LOOP"
	case LSL:
		return "LSL"
	case LSS:
		return "LSS"
	case LZCNT:
		return "LZCNT"
	case MASKMOVDQU:
		return "MASKMOVDQU"
	case MASKMOVQ:
		return "MASKMOVQ"
	case MAXPD:
		return "MASKMOVQ"
	case MAXPS:
		return "MASKMOVQ"
	case MAXSD:
		return "MASKMOVQ"
	case MAXSS:
		return "MAXSS"
	case MINPD:
		return "MAXSS"
	case MINPS:
		return "MAXSS"
	case MINSD:
		return "MINSD"
	case MINSS:
		return "MINSD"
	case MOV:
		return "MINSD"
	case MOVAPD:
		return "MOVAPD"
	case MOVAPS:
		return "MOVAPS"
	case MOVBE:
		return "MOVBE"
	case MOVD:
		return "MOVD"
	case MOVDDUP:
		return "MOVDUP"
	case MOVDQ2Q:
		return "MOVDQ2Q"
	case MOVDQA:
		return "MOVDQ2Q"
	case MOVDQU:
		return "MOVDQ2Q"
	case MOVHPD:
		return "MOVDQ2Q"
	case MOVHPS:
		return "MOVHPS"
	case MOVLPD:
		return "MOVLPD"
	case MOVLPS:
		return "MOVLPD"
	case MOVMSKPD:
		return "MOVMSKPD"
	case MOVMSKPS:
		return "MOVMSKPS"
	case MOVNTDQ:
		return "MOVNTDQ"
	case MOVNTDQA:
		return "MOVNTDQA"
	case MOVNTI:
		return "MOVNTDQA"
	case MOVNTPD:
		return "MOVNTPD"
	case MOVNTPS:
		return "MOVNTPS"
	case MOVNTQ:
		return "MOVNTQ"
	case MOVNTSS:
		return "MOVNTSS"
	case MOVNTSD:
		return "MOVNTSD"
	case MOVQ:
		return "MOVQ"
	case MOVQ2DQ:
		return "MOVQ2DQ"
	case MOVSB:
		return "MOVSB"
	case MOVSD:
		return "MOVSD"
	case MOVSHDUP:
		return "MOVSHDUP"
	case MOVSLDUP:
		return "MOVSLDUP"
	case MOVSS:
		return "MOVSS"
	case MOVSX:
		return "MOVSX"
	case MOVUPD:
		return "MOVUPD"
	case MOVUPS:
		return "MOVUPS"
	case MOVZX:
		return "MOVZX"
	case MPSADBW:
		return "MPSADBW"
	case MULPD:
		return "MULPD"
	case MULPS:
		return "MULPS"
	case MULSD:
		return "MULSD"
	case MULSS:
		return "MULSS"
	case NOP:
		return "NOP"
	case OR:
		return "OR"
	case ORPD:
		return "ORPD"
	case ORPS:
		return "ORPS"
	case OUT:
		return "OUT"
	case OUTSB:
		return "OUTSB"
	case PABSB:
		return "PABSB"
	case PABSD:
		return "PABSD"
	case PABSQ:
		return "PABSQ"
	case PABSW:
		return "PABSW"
	case PACKSSDW:
		return "PACKSSDW"
	case PACKSSWB:
		return "PACKSSWB"
	case PACKUSDW:
		return "PACKUSDW"
	case PACKUSWB:
		return "PACKUSWB"
	case PADDB:
		return "PADDB"
	case PADDD:
		return "PADDD"
	case PADDQ:
		return "PADDQ"
	case PADDSB:
		return "PADDSB"
	case PADDSW:
		return "PADDSW"
	case PADDUSB:
		return "PADDUSB"
	case PADDUSW:
		return "PADDUSW"
	case PADDW:
		return "PADDW"
	case PALIGNR:
		return "PALIGNR"
	case PAND:
		return "PAND"
	case PANDN:
		return "PANDN"
	case PAUSE:
		return "PAUSE"
	case PAVGB:
		return "PAVGB"
	case PAVGW:
		return "PAVGW"
	case PBLENDVB:
		return "PBLENDVB"
	case PBLENDW:
		return "PBLENDW"
	case PCLMULQDQ:
		return "PCLMULQDQ"
	case PCMPEQB:
		return "PCMPEQB"
	case PCMPEQD:
		return "PCMPEQD"
	case PCMPEQQ:
		return "PCMPEQQ"
	case PCMPEQW:
		return "PCMPEQW"
	case PCMPGTB:
		return "PCMPGTB"
	case PCMPGTD:
		return "PCMPGTD"
	case PCMPGTQ:
		return "PCMPGTQ"
	case PCMPGTW:
		return "PCMPGTW"
	case PEXTRB:
		return "PEXTRB"
	case PEXTRD:
		return "PEXTRD"
	case PEXTRW:
		return "PEXTRW"
	case PHADDD:
		return "PHADDD"
	case PHADDSW:
		return "PHADDSW"
	case PHADDW:
		return "PHADDW"
	case PHMINPOSUW:
		return "PHMINPOSUW"
	case PHSUBD:
		return "PHSUBD"
	case PHSUBSW:
		return "PHSUBSW"
	case PHSUBW:
		return "PHSUBW"
	case PINSRB:
		return "PINSRB"
	case PINSRW:
		return "PINSRW"
	case PI2FD:
		return "PI2FD"
	case PFRSQIT1:
		return "PFRSQIT1"
	case PMADDWD:
		return "PMADDWD"
	case PMADDUBSW:
		return "PMADDUBSW"
	case PMAXSB:
		return "PMAXSB"
	case PMAXSD:
		return "PMAXSD"
	case PMAXSW:
		return "PMAXSW"
	case PMAXUB:
		return "PMAXUB"
	case PMAXUD:
		return "PMAXUD"
	case PMAXUW:
		return "PMAXUW"
	case PMINSB:
		return "PMINSB"
	case PMINSD:
		return "PMINSD"
	case PMINSW:
		return "PMINSW"
	case PMINUB:
		return "PMINUB"
	case PMINUD:
		return "PMINUD"
	case PMINUW:
		return "PMINUW"
	case PMOVMSKB:
		return "PMOVMSKB"
	case PMOVSXBW:
		return "PMOVSXBW"
	case PMOVSXBD:
		return "PMOVSXBD"
	case PMOVSXBQ:
		return "PMOVSXBQ"
	case PMOVSXWD:
		return "PMOVSXWD"
	case PMOVSXWQ:
		return "PMOVSXWQ"
	case PMOVSXDQ:
		return "PMOVSXDQ"
	case PMOVZXBD:
		return "PMOVZXBD"
	case PMOVZXBQ:
		return "PMOVZXBQ"
	case PI2FW:
		return "PI2FW"
	case PF2IW:
		return "PF2IW"
	case PF2ID:
		return "PF2ID"
	case PFNACC:
		return "PFNACC"
	case PFPNACC:
		return "PFPNACC"
	case PFCMPGE:
		return "PFCMPGE"
	case PFMIN:
		return "PFMIN"
	case PFRCP:
		return "PFRCP"
	case PFRSQRT:
		return "PFRSQRT"
	case PFSUB:
		return "PFSUB"
	case PFADD:
		return "PFADD"
	case PFCMPGT:
		return "PFCMPGT"
	case PFMAX:
		return "PFMAX"
	case PFRCPIT1:
		return "PFRCPIT1"
	case PFSUBR:
		return "PFSUBR"
	case PFACC:
		return "PFACC"
	case PFCMPEQ:
		return "PFCMPEQ"
	case PFMUL:
		return "PFMUL"
	case PFRCPIT2:
		return "PFRCPIT2"
	case PSWAPD:
		return "PSWAPD"
	case PAVGUSB:
		return "PAVGUSB"
	case POPD:
		return "POPD"
	case PMULHRW:
		return "PMULHRW"
	case PMOVZXBW:
		return "PMOVZXBW"
	case PMOVZXDQ:
		return "PMOVZXDQ"
	case PMOVZXWD:
		return "PMOVZXWD"
	case PMOVZXWQ:
		return "PMOVZXWQ"
	case PMULDQ:
		return "PMULDQ"
	case PMULHRSW:
		return "PMULHRSW"
	case PMULHUW:
		return "PMULHUW"
	case PMULHW:
		return "PMULHW"
	case PMULLD:
		return "PMULLD"
	case PMULLQ:
		return "PMULLQ"
	case PMULLW:
		return "PMULLW"
	case PMULUDQ:
		return "PMULUDQ"
	case POP:
		return "POP"
	case POPA:
		return "POPA"
	case POPAD:
		return "POPAD"
	case POPCNT:
		return "POPCNT"
	case POR:
		return "POR"
	case PSADBW:
		return "PSADBW"
	case PSHUFB:
		return "PSHUFB"
	case PSHUFD:
		return "PSHUFD"
	case PSHUFHW:
		return "PSHUFHW"
	case PSHUFLW:
		return "PSHUFLW"
	case PSHUFW:
		return "PSHUFW"
	case PSIGNB:
		return "PSIGNB"
	case PSIGND:
		return "PSIGND"
	case PSIGNW:
		return "PSIGNW"
	case PSLLD:
		return "PSLLD"
	case PSLLDQ:
		return "PSLLDQ"
	case PSLLQ:
		return "PSLLQ	"
	case PSLLW:
		return "PSLLW"
	case PSRAD:
		return "PSRAD"
	case PSRAQ:
		return "PSRAQ"
	case PSRAW:
		return "PSRAW"
	case PSRLD:
		return "PSRLD"
	case PSRLDQ:
		return "PSRLDQ"
	case PSRLQ:
		return "PSRLQ"
	case PSRLW:
		return "PSRLW"
	case PSUBB:
		return "PSUBB"
	case PSUBD:
		return "PSUBD"
	case PSUBQ:
		return "PSUBQ"
	case PSUBSB:
		return "PSUBSB"
	case PSUBSW:
		return "PSUBSW"
	case PSUBUSB:
		return "PSUBUSB"
	case PSUBUSW:
		return "PSUBUSW"
	case PSUBW:
		return "PSUBW"
	case PTEST:
		return "PTEST"
	case PTWRITE:
		return "PTWRITE"
	case PUNPCKHBW:
		return "PUNPCKHBW"
	case PUNPCKHDQ:
		return "PUNPCKHDQ"
	case PUNPCKHQDQ:
		return "PUNPCKHQDQ"
	case PUNPCKHWD:
		return "PUNPCKHWD"
	case PUNPCKLBW:
		return "PUNPCKLBW"
	case PUNPCKLDQ:
		return "PUNPCKLDQ"
	case PUNPCKLQDQ:
		return "PUNPCKLQDQ"
	case PUNPCKLWD:
		return "PUNPCKLWD"
	case PUSH:
		return "PUSH"
	case PUSHA:
		return "PUSHA"
	case PUSHAD:
		return "PUSHAD"
	case PUSHD:
		return "PUSHD"
	case PXOR:
		return "PXOR"
	case RCPPS:
		return "RCPPS"
	case RCPSS:
		return "RCPSS"
	case RDMSR:
		return "RDMSR"
	case RDPMC:
		return "RDPMC"
	case RDTSC:
		return "RDTSC"
	case RET:
		return "RET"
	case RSM:
		return "RSM"
	case RSQRTPS:
		return "RSQRTPS"
	case RSQRTSS:
		return "RSQRTSS"
	case SAHF:
		return "SAHF"
	case SALC:
		return "SALC"
	case SBB:
		return "SBB"
	case SCASB:
		return "SCASB"
	case SHLD:
		return "SHLD"
	case SHRD:
		return "SHRD"
	case SHUFPD:
		return "SHUFPD"
	case SETO:
		return "SETO"
	case SETB:
		return "SETB"
	case SETNB:
		return "SETNB"
	case SETZ:
		return "SETZ"
	case SETBE:
		return "SETBE"
	case SETNBE:
		return "SETNBE"
	case SETS:
		return "SETS"
	case SETNS:
		return "SETNS"
	case SETP:
		return "SETP"
	case SETNP:
		return "SETNP"
	case SETL:
		return "SETL"
	case SETNL:
		return "SETNL"
	case SETLE:
		return "SETLE"
	case SETNLE:
		return "SETNLE"
	case SETNZ:
		return "SETNZ"
	case SETNO:
		return "SETNO"
	case SHUFPS:
		return "SHUFPS"
	case SQRTPD:
		return "SQRTPD"
	case SQRTPS:
		return "SQRTPS"
	case SQRTSD:
		return "SQRTSD"
	case SQRTSS:
		return "SQRTSS"
	case STC:
		return "STC"
	case STD:
		return "STD"
	case STI:
		return "STI"
	case STOSB:
		return "STOSB"
	case SUB:
		return "SUB"
	case SUBPD:
		return "SUBPD"
	case SUBPS:
		return "SUBPS"
	case SUBSD:
		return "SUBSD"
	case SUBSS:
		return "SUBSS"
	case SYSCALL:
		return "SYSCALL"
	case SYSENTER:
		return "SYSENTER"
	case SYSEXIT:
		return "SYSEXIT"
	case SYSRET:
		return "SYSRET"
	case TEST:
		return "TEST"
	case TZCNT:
		return "TZCNT"
	case UCOMISD:
		return "UCOMISD"
	case UD0:
		return "UD0"
	case UD2:
		return "UD2"
	case UCOMISS:
		return "UCOMISS"
	case UNPCKHPD:
		return "UNPCKHPD"
	case UNPCKHPS:
		return "UNPCKHPS"
	case UNPCKLPD:
		return "UNPCKLPD"
	case UNPCKLPS:
		return "UNPCKLPS"
	case WRMSR:
		return "WRMSR"
	case WRUSS:
		return "WRUSS"
	case XADD:
		return "XADD"
	case XCHG:
		return "XCHG"
	case XLAT:
		return "XLAT"
	case XOR:
		return "XOR"
	case XORPD:
		return "XORPD"
	case XORPS:
		return "XORPS"
	default:
		return "Unknown instruction"
	}
}
