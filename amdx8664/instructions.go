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
	AESKEYGENASSIST
	AND
	ANDNPD
	ANDNPS
	ANDPD
	ANDPS
	ARPL
	BEXTR
	BLCFILL
	BLCI
	BLCIC
	BLCMSK
	BLCS
	BLENDPD
	BLENDPS
	BLENDVPS
	BLSFILL
	BLSI
	BLSIC
	BLSMSK
	BLSR
	BOUND
	BSF
	BSR
	BSWAP
	BT
	BTC
	BTR
	BTS
	CALL
	CLC
	CLD
	CLI
	CLRSSBSY
	CLTS
	CLWB
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
	ENDBR32
	ENDBR64
	ENTER
	EXTRACTPS
	EXTRQ
	FEMMS
	FXRSTOR
	FXSAVE
	HADDPD
	HADDPS
	HLT
	HRESET
	HSUBPD
	HSUBPS
	IDIV
	IMUL
	IN
	INC
	INCSSP
	INSB
	INSERTPS
	INSERTQ
	INT
	INT1
	INT3
	INTO
	INVD
	JB
	JBE
	JL
	JLE
	JMP
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
	KMOVB
	KMOVW
	KUNPCKBW
	LAHF
	LAR
	LDDQU
	LDMXCSR
	LDS
	LEA
	LEAVE
	LES
	LFS
	LGDT
	LGS
	LLDT
	LLWPCB
	LODSB
	LOOP
	LSL
	LSS
	LTR
	LWPINS
	LWPVAL
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
	MOVNTSD
	MOVNTSS
	MOVQ
	MOVQ2DQ
	MOVSB
	MOVSD
	MOVSHDUP
	MOVSLDUP
	MOVSS
	MOVSX
	MOVSXD
	MOVUPD
	MOVUPS
	MOVZX
	MPSADBW
	MUL
	MULPD
	MULPS
	MULSD
	MULSS
	NEG
	NOP
	NOT
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
	PAVGUSB
	PAVGW
	PBLENDVB
	PBLENDW
	PCLMULQDQ
	PCMPEQB
	PCMPEQD
	PCMPEQQ
	PCMPEQW
	PCMPESTRI
	PCMPESTRM
	PCMPGTB
	PCMPGTD
	PCMPGTQ
	PCMPGTW
	PCMPISTRI
	PCMPISTRM
	PEXTRB
	PEXTRD
	PEXTRW
	PF2ID
	PF2IW
	PFACC
	PFADD
	PFCMPEQ
	PFCMPGE
	PFCMPGT
	PFMAX
	PFMIN
	PFMUL
	PFNACC
	PFPNACC
	PFRCP
	PFRCPIT1
	PFRCPIT2
	PFRSQIT1
	PFRSQRT
	PFSUB
	PFSUBR
	PHADDD
	PHADDSW
	PHADDW
	PHMINPOSUW
	PHSUBD
	PHSUBSW
	PHSUBW
	PI2FD
	PI2FW
	PINSRB
	PINSRW
	PMADDUBSW
	PMADDWD
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
	PMOVSXBD
	PMOVSXBQ
	PMOVSXBW
	PMOVSXDQ
	PMOVSXWD
	PMOVSXWQ
	PMOVZXBD
	PMOVZXBQ
	PMOVZXBW
	PMOVZXDQ
	PMOVZXWD
	PMOVZXWQ
	PMULDQ
	PMULHRSW
	PMULHRW
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
	POPD
	POR
	PREFETCH
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
	PSWAPD
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
	RCL
	RCPPS
	RCPSS
	RCR
	RDFSBASE
	RDGSBASE
	RDMSR
	RDPMC
	RDRAND
	RDSEED
	RDTSC
	RET
	ROL
	ROR
	ROUNDPD
	ROUNDPS
	ROUNDSD
	ROUNDSS
	RSM
	RSQRTPS
	RSQRTSS
	SAHF
	SALC
	SAR
	SBB
	SCASB
	SETB
	SETBE
	SETL
	SETLE
	SETNB
	SETNBE
	SETNL
	SETNLE
	SETNO
	SETNP
	SETNS
	SETNZ
	SETO
	SETP
	SETS
	SETZ
	SGDT
	SHLD
	SHR
	SHRD
	SHUFPD
	SHUFPS
	SLWPCB
	SQRTPD
	SQRTPS
	SQRTSD
	SQRTSS
	STC
	STD
	STI
	STMXCSR
	STOSB
	STR
	SUB
	SUBPD
	SUBPS
	SUBSD
	SUBSS
	SYSCALL
	SYSENTER
	SYSEXIT
	SYSRET
	T1MSKC
	TEST
	TZCNT
	TZMSK
	UCOMISD
	UCOMISS
	UD0
	UD1
	UD2
	UNPCKHPD
	UNPCKHPS
	UNPCKLPD
	UNPCKLPS
	VADDPD
	VADDPS
	VADDSD
	VADDSS
	VADDSUBPD
	VADDSUBPS
	VAESDEC
	VAESDECLAST
	VAESENC
	VAESENCLAST
	VAESIMC
	VAESKEYGENASSIST
	VANDNPD
	VANDNPS
	VANDPD
	VANDPS
	VBLENDPD
	VBLENDPS
	VBROADCASTF128
	VBROADCASTI128
	VBROADCASTSD
	VBROADCASTSS
	VCMPPD
	VCMPPS
	VCMPSD
	VCMPSS
	VCOMISD
	VCOMISS
	VCVTDQ2PD
	VCVTPD2DQ
	VCVTPD2PS
	VCVTPH2PS
	VCVTPS2PD
	VCVTPS2PH
	VCVTSD2SS
	VCVTSI2SD
	VCVTSI2SS
	VCVTSS2SD
	VCVTTPD2DQ
	VDIVPD
	VDIVPS
	VDIVSD
	VDIVSS
	VDPPD
	VDPPS
	VERR
	VERW
	VEXTRACTF128
	VEXTRACTI128
	VEXTRACTPS
	VFMADDPD
	VFMADDPS
	VFMADDSD
	VFMADDSS
	VFMADDSUBPD
	VFMADDSUBPS
	VFMSUBADDPD
	VFMSUBADDPS
	VFMSUBPD
	VFMSUBPS
	VFMSUBSD
	VFMSUBSS
	VFNMADDPD
	VFNMADDPS
	VFNMADDSD
	VFNMADDSS
	VFNMSUBPD
	VFNMSUBPS
	VFNMSUBSD
	VFNMSUBSS
	VFRCZPD
	VFRCZPS
	VFRCZSD
	VFRCZSS
	VGF2P8AFFINEINVQB
	VGF2P8AFFINEQB
	VGF2P8MULB
	VHADDPD
	VHADDPS
	VHSUBPD
	VHSUBPS
	VINSERTF128
	VINSERTI128
	VINSERTPS
	VLDDQU
	VLDMXCSR
	VMASKMOVDQU
	VMASKMOVPD
	VMASKMOVPS
	VMAXPD
	VMAXPS
	VMAXSD
	VMAXSS
	VMINPD
	VMINPS
	VMINSD
	VMINSS
	VMOVAPD
	VMOVAPS
	VMOVDDUP
	VMOVDQA
	VMOVDQU
	VMOVHPD
	VMOVHPS
	VMOVLPD
	VMOVLPS
	VMOVMSKPD
	VMOVMSKPS
	VMOVNTDQ
	VMOVNTDQA
	VMOVNTPD
	VMOVNTPS
	VMOVQ
	VMOVSHDUP
	VMOVSLDUP
	VMOVUPD
	VMOVUPS
	VMPSADBW
	VMULPD
	VMULPS
	VMULSD
	VMULSS
	VORPD
	VORPS
	VPABSB
	VPABSD
	VPABSW
	VPACKSSDW
	VPACKSSWB
	VPACKUSDW
	VPACKUSWB
	VPADDB
	VPADDD
	VPADDQ
	VPADDSB
	VPADDSW
	VPADDUSB
	VPADDUSW
	VPADDW
	VPALIGNR
	VPAND
	VPANDN
	VPAVGB
	VPAVGW
	VPBLENDD
	VPBLENDW
	VPBROADCASTB
	VPBROADCASTD
	VPBROADCASTQ
	VPBROADCASTW
	VPCLMULQDQ
	VPCMOV
	VPCMPEQB
	VPCMPEQD
	VPCMPEQQ
	VPCMPEQW
	VPCMPESTRI
	VPCMPESTRM
	VPCMPGTB
	VPCMPGTD
	VPCMPGTQ
	VPCMPGTW
	VPCMPISTRI
	VPCMPISTRM
	VPCOMccB
	VPCOMccD
	VPCOMccQ
	VPCOMccUB
	VPCOMccUD
	VPCOMccUQ
	VPCOMccUW
	VPCOMccW
	VPDPBUSD
	VPDPBUSDS
	VPDPWSSD
	VPDPWSSDS
	VPERM2F128
	VPERM2I128
	VPERMD
	VPERMILPD
	VPERMILPS
	VPERMPD
	VPERMPS
	VPERMQ
	VPEXTRB
	VPEXTRW
	VPHADDBD
	VPHADDBQ
	VPHADDBW
	VPHADDD
	VPHADDDQ
	VPHADDSW
	VPHADDUBD
	VPHADDUBQ
	VPHADDUBWD
	VPHADDUDQ
	VPHADDUWD
	VPHADDUWQ
	VPHADDW
	VPHADDWD
	VPHADDWQ
	VPHMINPOSUW
	VPHSUBBW
	VPHSUBD
	VPHSUBDQ
	VPHSUBSW
	VPHSUBW
	VPHSUBWD
	VPINSRB
	VPINSRW
	VPMACSDD
	VPMACSDQH
	VPMACSDQL
	VPMACSSDD
	VPMACSSDQH
	VPMACSSDQL
	VPMACSSWD
	VPMACSSWW
	VPMACSWD
	VPMACSWW
	VPMADCSSWD
	VPMADCSWD
	VPMADDUBSW
	VPMADDWD
	VPMAXSB
	VPMAXSD
	VPMAXSW
	VPMAXUB
	VPMAXUD
	VPMAXUW
	VPMINSB
	VPMINSD
	VPMINSW
	VPMINUB
	VPMINUD
	VPMINUW
	VPMOVMSKB
	VPMOVSXBD
	VPMOVSXBQ
	VPMOVSXBW
	VPMOVSXDQ
	VPMOVSXWD
	VPMOVSXWQ
	VPMOVZXBD
	VPMOVZXBQ
	VPMOVZXBW
	VPMOVZXDQ
	VPMOVZXWD
	VPMOVZXWQ
	VPMULDQ
	VPMULHUW
	VPMULHW
	VPMULLD
	VPMULLW
	VPMULUDQ
	VPOR
	VPPERM
	VPROTB
	VPROTD
	VPROTQ
	VPROTW
	VPSADBW
	VPSHAB
	VPSHAD
	VPSHAQ
	VPSHAW
	VPSHLB
	VPSHLD
	VPSHLQ
	VPSHLW
	VPSHUFB
	VPSHUFD
	VPSHUFHW
	VPSHUFLW
	VPSIGNB
	VPSIGND
	VPSIGNW
	VPSLLD
	VPSLLDQ
	VPSLLQ
	VPSLLW
	VPSRAD
	VPSRAVD
	VPSRAW
	VPSRLD
	VPSRLDQ
	VPSRLQ
	VPSRLW
	VPSUBB
	VPSUBD
	VPSUBQ
	VPSUBSB
	VPSUBSW
	VPSUBUSW
	VPSUBW
	VPTEST
	VPUNPCKHBW
	VPUNPCKHDQ
	VPUNPCKHQDQ
	VPUNPCKHWD
	VPUNPCKLBW
	VPUNPCKLDQ
	VPUNPCKLQDQ
	VPUNPCKLWD
	VPXOR
	VRCPPS
	VROUNDPD
	VROUNDPS
	VROUNDSD
	VROUNDSS
	VRSQRTPS
	VRSQRTSS
	VSHUFPD
	VSHUFPS
	VSQRTPD
	VSQRTPS
	VSQRTSD
	VSQRTSS
	VSTMXCSR
	VSUBPD
	VSUBPS
	VSUBSD
	VSUBSS
	VTESTPD
	VTESTPS
	VUCOMISD
	VUCOMISS
	VUNPCKHPD
	VUNPCKHPS
	VUNPCKLPD
	VUNPCKLPS
	VXORPD
	VXORPS
	WRFSBASE
	WRMSR
	WRSS
	WRUSS
	XADD
	XCHG
	XLAT
	XOR
	XORPD
	XORPS
	XSAVE
	NoInstruction
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
	case AESKEYGENASSIST:
		return "AESKEYGENASSIST"
	case AND:
		return "AND"
	case ANDNPD:
		return "ANDNPD"
	case ANDNPS:
		return "ANDNPS"
	case ANDPD:
		return "ANDPD"
	case ANDPS:
		return "ANDPS"
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
	case BLENDPD:
		return "BLENDPD"
	case BLENDPS:
		return "BLENDPS"
	case BLENDVPS:
		return "BLENDVPS"
	case BLSFILL:
		return "BLSFILL"
	case BLSI:
		return "BLSI"
	case BLSIC:
		return "BLSIC"
	case BLSMSK:
		return "BLSMSK"
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
	case CALL:
		return "CALL"
	case CLC:
		return "CLC"
	case CLD:
		return "CLD"
	case CLI:
		return "CLI"
	case CLRSSBSY:
		return "CLRSSBSY"
	case CLTS:
		return "CLTS"
	case CLWB:
		return "CLWB"
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
	case ENDBR32:
		return "ENDBR32"
	case ENDBR64:
		return "ENDBR64"
	case ENTER:
		return "ENTER"
	case EXTRACTPS:
		return "EXTRACTPS"
	case EXTRQ:
		return "EXTRQ"
	case FEMMS:
		return "FEMMS"
	case FXRSTOR:
		return "FXRSTOR"
	case FXSAVE:
		return "FXSAVE"
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
		return "HSUBPS"
	case IDIV:
		return "IDIV"
	case IMUL:
		return "IMUL"
	case IN:
		return "IN"
	case INC:
		return "INC"
	case INCSSP:
		return "INCSSP"
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
	case JB:
		return "JB"
	case JBE:
		return "JBE"
	case JL:
		return "JL"
	case JLE:
		return "JLE"
	case JMP:
		return "JMP"
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
	case KMOVB:
		return "KMOVB"
	case KMOVW:
		return "KMOVW"
	case KUNPCKBW:
		return "KUNPCKBW"
	case LAHF:
		return "LAHF"
	case LAR:
		return "LAR"
	case LDDQU:
		return "LDDQU"
	case LDMXCSR:
		return "LDMXCSR"
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
	case LLDT:
		return "LLDT"
	case LLWPCB:
		return "LLWPCB"
	case LODSB:
		return "LODSB"
	case LOOP:
		return "LOOP"
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
	case MASKMOVDQU:
		return "MASKMOVDQU"
	case MASKMOVQ:
		return "MASKMOVQ"
	case MAXPD:
		return "MAXPD"
	case MAXPS:
		return "MAXPS"
	case MAXSD:
		return "MAXSD"
	case MAXSS:
		return "MAXSS"
	case MINPD:
		return "MINPD"
	case MINPS:
		return "MINPS"
	case MINSD:
		return "MINSD"
	case MINSS:
		return "MINSS"
	case MOV:
		return "MOV"
	case MOVAPD:
		return "MOVAPD"
	case MOVAPS:
		return "MOVAPS"
	case MOVBE:
		return "MOVBE"
	case MOVD:
		return "MOVD"
	case MOVDDUP:
		return "MOVDDUP"
	case MOVDQ2Q:
		return "MOVDQ2Q"
	case MOVDQA:
		return "MOVDQA"
	case MOVDQU:
		return "MOVDQU"
	case MOVHPD:
		return "MOVHPD"
	case MOVHPS:
		return "MOVHPS"
	case MOVLPD:
		return "MOVLPD"
	case MOVLPS:
		return "MOVLPS"
	case MOVMSKPD:
		return "MOVMSKPD"
	case MOVMSKPS:
		return "MOVMSKPS"
	case MOVNTDQ:
		return "MOVNTDQ"
	case MOVNTDQA:
		return "MOVNTDQA"
	case MOVNTI:
		return "MOVNTI"
	case MOVNTPD:
		return "MOVNTPD"
	case MOVNTPS:
		return "MOVNTPS"
	case MOVNTQ:
		return "MOVNTQ"
	case MOVNTSD:
		return "MOVNTSD"
	case MOVNTSS:
		return "MOVNTSS"
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
	case MOVSXD:
		return "MOVSXD"
	case MOVUPD:
		return "MOVUPD"
	case MOVUPS:
		return "MOVUPS"
	case MOVZX:
		return "MOVZX"
	case MPSADBW:
		return "MPSADBW"
	case MUL:
		return "MUL"
	case MULPD:
		return "MULPD"
	case MULPS:
		return "MULPS"
	case MULSD:
		return "MULSD"
	case MULSS:
		return "MULSS"
	case NEG:
		return "NEG"
	case NOP:
		return "NOP"
	case NOT:
		return "NOT"
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
	case PAVGUSB:
		return "PAVGUSB"
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
	case PCMPESTRI:
		return "PCMPESTRI"
	case PCMPESTRM:
		return "PCMPESTRM"
	case PCMPGTB:
		return "PCMPGTB"
	case PCMPGTD:
		return "PCMPGTD"
	case PCMPGTQ:
		return "PCMPGTQ"
	case PCMPGTW:
		return "PCMPGTW"
	case PCMPISTRI:
		return "PCMPISTRI"
	case PCMPISTRM:
		return "PCMPISTRM"
	case PEXTRB:
		return "PEXTRB"
	case PEXTRD:
		return "PEXTRD"
	case PEXTRW:
		return "PEXTRW"
	case PF2ID:
		return "PF2ID"
	case PF2IW:
		return "PF2IW"
	case PFACC:
		return "PFACC"
	case PFADD:
		return "PFADD"
	case PFCMPEQ:
		return "PFCMPEQ"
	case PFCMPGE:
		return "PFCMPGE"
	case PFCMPGT:
		return "PFCMPGT"
	case PFMAX:
		return "PFMAX"
	case PFMIN:
		return "PFMIN"
	case PFMUL:
		return "PFMUL"
	case PFNACC:
		return "PFNACC"
	case PFPNACC:
		return "PFPNACC"
	case PFRCP:
		return "PFRCP"
	case PFRCPIT1:
		return "PFRCPIT1"
	case PFRCPIT2:
		return "PFRCPIT2"
	case PFRSQIT1:
		return "PFRSQIT1"
	case PFRSQRT:
		return "PFRSQRT"
	case PFSUB:
		return "PFSUB"
	case PFSUBR:
		return "PFSUBR"
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
	case PI2FD:
		return "PI2FD"
	case PI2FW:
		return "PI2FW"
	case PINSRB:
		return "PINSRB"
	case PINSRW:
		return "PINSRW"
	case PMADDUBSW:
		return "PMADDUBSW"
	case PMADDWD:
		return "PMADDWD"
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
	case PMOVSXBD:
		return "PMOVSXBD"
	case PMOVSXBQ:
		return "PMOVSXBQ"
	case PMOVSXBW:
		return "PMOVSXBW"
	case PMOVSXDQ:
		return "PMOVSXDQ"
	case PMOVSXWD:
		return "PMOVSXWD"
	case PMOVSXWQ:
		return "PMOVSXWQ"
	case PMOVZXBD:
		return "PMOVZXBD"
	case PMOVZXBQ:
		return "PMOVZXBQ"
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
	case PMULHRW:
		return "PMULHRW"
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
	case POPD:
		return "POPD"
	case POR:
		return "POR"
	case PREFETCH:
		return "PREFETCH"
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
		return "PSLLQ"
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
	case PSWAPD:
		return "PSWAPD"
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
	case RCL:
		return "RCL"
	case RCPPS:
		return "RCPPS"
	case RCPSS:
		return "RCPSS"
	case RCR:
		return "RCR"
	case RDFSBASE:
		return "RDFSBASE"
	case RDGSBASE:
		return "RDGSBASE"
	case RDMSR:
		return "RDMSR"
	case RDPMC:
		return "RDPMC"
	case RDRAND:
		return "RDRAND"
	case RDSEED:
		return "RDSEED"
	case RDTSC:
		return "RDTSC"
	case RET:
		return "RET"
	case ROL:
		return "ROL"
	case ROR:
		return "ROR"
	case ROUNDPD:
		return "ROUNDPD"
	case ROUNDPS:
		return "ROUNDPS"
	case ROUNDSD:
		return "ROUNDSD"
	case ROUNDSS:
		return "ROUNDSS"
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
	case SAR:
		return "SAR"
	case SBB:
		return "SBB"
	case SCASB:
		return "SCASB"
	case SETB:
		return "SETB"
	case SETBE:
		return "SETBE"
	case SETL:
		return "SETL"
	case SETLE:
		return "SETLE"
	case SETNB:
		return "SETNB"
	case SETNBE:
		return "SETNBE"
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
	case SETS:
		return "SETS"
	case SETZ:
		return "SETZ"
	case SGDT:
		return "SGDT"
	case SHLD:
		return "SHLD"
	case SHR:
		return "SHR"
	case SHRD:
		return "SHRD"
	case SHUFPD:
		return "SHUFPD"
	case SHUFPS:
		return "SHUFPS"
	case SLWPCB:
		return "SLWPCB"
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
	case STMXCSR:
		return "STMXCSR"
	case STOSB:
		return "STOSB"
	case STR:
		return "STR"
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
	case T1MSKC:
		return "T1MSKC"
	case TEST:
		return "TEST"
	case TZCNT:
		return "TZCNT"
	case TZMSK:
		return "TZMSK"
	case UCOMISD:
		return "UCOMISD"
	case UCOMISS:
		return "UCOMISS"
	case UD0:
		return "UD0"
	case UD1:
		return "UD1"
	case UD2:
		return "UD2"
	case UNPCKHPD:
		return "UNPCKHPD"
	case UNPCKHPS:
		return "UNPCKHPS"
	case UNPCKLPD:
		return "UNPCKLPD"
	case UNPCKLPS:
		return "UNPCKLPS"
	case VADDPD:
		return "VADDPD"
	case VADDPS:
		return "VADDPS"
	case VADDSD:
		return "VADDSD"
	case VADDSS:
		return "VADDSS"
	case VADDSUBPD:
		return "VADDSUBPD"
	case VADDSUBPS:
		return "VADDSUBPS"
	case VAESDEC:
		return "VAESDEC"
	case VAESDECLAST:
		return "VAESDECLAST"
	case VAESENC:
		return "VAESENC"
	case VAESENCLAST:
		return "VAESENCLAST"
	case VAESIMC:
		return "VAESIMC"
	case VAESKEYGENASSIST:
		return "VAESKEYGENASSIST"
	case VANDNPD:
		return "VANDNPD"
	case VANDNPS:
		return "VANDNPS"
	case VANDPD:
		return "VANDPD"
	case VANDPS:
		return "VANDPS"
	case VBLENDPD:
		return "VBLENDPD"
	case VBLENDPS:
		return "VBLENDPS"
	case VBROADCASTF128:
		return "VBROADCASTF128"
	case VBROADCASTI128:
		return "VBROADCASTI128"
	case VBROADCASTSD:
		return "VBROADCASTSD"
	case VBROADCASTSS:
		return "VBROADCASTSS"
	case VCMPPD:
		return "VCMPPD"
	case VCMPPS:
		return "VCMPPS"
	case VCMPSD:
		return "VCMPSD"
	case VCMPSS:
		return "VCMPSS"
	case VCOMISD:
		return "VCOMISD"
	case VCOMISS:
		return "VCOMISS"
	case VCVTDQ2PD:
		return "VCVTDQ2PD"
	case VCVTPD2DQ:
		return "VCVTPD2DQ"
	case VCVTPD2PS:
		return "VCVTPD2PS"
	case VCVTPH2PS:
		return "VCVTPH2PS"
	case VCVTPS2PD:
		return "VCVTPS2PD"
	case VCVTPS2PH:
		return "VCVTPS2PH"
	case VCVTSD2SS:
		return "VCVTSD2SS"
	case VCVTSI2SD:
		return "VCVTSI2SD"
	case VCVTSI2SS:
		return "VCVTSI2SS"
	case VCVTSS2SD:
		return "VCVTSS2SD"
	case VCVTTPD2DQ:
		return "VCVTTPD2DQ"
	case VDIVPD:
		return "VDIVPD"
	case VDIVPS:
		return "VDIVPS"
	case VDIVSD:
		return "VDIVSD"
	case VDIVSS:
		return "VDIVSS"
	case VDPPD:
		return "VDPPD"
	case VDPPS:
		return "VDPPS"
	case VERR:
		return "VERR"
	case VERW:
		return "VERW"
	case VEXTRACTF128:
		return "VEXTRACTF128"
	case VEXTRACTI128:
		return "VEXTRACTI128"
	case VEXTRACTPS:
		return "VEXTRACTPS"
	case VFMADDPD:
		return "VFMADDPD"
	case VFMADDPS:
		return "VFMADDPS"
	case VFMADDSD:
		return "VFMADDSD"
	case VFMADDSS:
		return "VFMADDSS"
	case VFMADDSUBPD:
		return "VFMADDSUBPD"
	case VFMADDSUBPS:
		return "VFMADDSUBPS"
	case VFMSUBADDPD:
		return "VFMSUBADDPD"
	case VFMSUBADDPS:
		return "VFMSUBADDPS"
	case VFMSUBPD:
		return "VFMSUBPD"
	case VFMSUBPS:
		return "VFMSUBPS"
	case VFMSUBSD:
		return "VFMSUBSD"
	case VFMSUBSS:
		return "VFMSUBSS"
	case VFNMADDPD:
		return "VFNMADDPD"
	case VFNMADDPS:
		return "VFNMADDPS"
	case VFNMADDSD:
		return "VFNMADDSD"
	case VFNMADDSS:
		return "VFNMADDSS"
	case VFNMSUBPD:
		return "VFNMSUBPD"
	case VFNMSUBPS:
		return "VFNMSUBPS"
	case VFNMSUBSD:
		return "VFNMSUBSD"
	case VFNMSUBSS:
		return "VFNMSUBSS"
	case VFRCZPD:
		return "VFRCZPD"
	case VFRCZPS:
		return "VFRCZPS"
	case VFRCZSD:
		return "VFRCZSD"
	case VFRCZSS:
		return "VFRCZSS"
	case VGF2P8AFFINEINVQB:
		return "VGF2P8AFFINEINVQB"
	case VGF2P8AFFINEQB:
		return "VGF2P8AFFINEQB"
	case VGF2P8MULB:
		return "VGF2P8MULB"
	case VHADDPD:
		return "VHADDPD"
	case VHADDPS:
		return "VHADDPS"
	case VHSUBPD:
		return "VHSUBPD"
	case VHSUBPS:
		return "VHSUBPS"
	case VINSERTF128:
		return "VINSERTF128"
	case VINSERTI128:
		return "VINSERTI128"
	case VINSERTPS:
		return "VINSERTPS"
	case VLDDQU:
		return "VLDDQU"
	case VLDMXCSR:
		return "VLDMXCSR"
	case VMASKMOVDQU:
		return "VMASKMOVDQU"
	case VMASKMOVPD:
		return "VMASKMOVPD"
	case VMASKMOVPS:
		return "VMASKMOVPS"
	case VMAXPD:
		return "VMAXPD"
	case VMAXPS:
		return "VMAXPS"
	case VMAXSD:
		return "VMAXSD"
	case VMAXSS:
		return "VMAXSS"
	case VMINPD:
		return "VMINPD"
	case VMINPS:
		return "VMINPS"
	case VMINSD:
		return "VMINSD"
	case VMINSS:
		return "VMINSS"
	case VMOVAPD:
		return "VMOVAPD"
	case VMOVAPS:
		return "VMOVAPS"
	case VMOVDDUP:
		return "VMOVDDUP"
	case VMOVDQA:
		return "VMOVDQA"
	case VMOVDQU:
		return "VMOVDQU"
	case VMOVHPD:
		return "VMOVHPD"
	case VMOVHPS:
		return "VMOVHPS"
	case VMOVLPD:
		return "VMOVLPD"
	case VMOVLPS:
		return "VMOVLPS"
	case VMOVMSKPD:
		return "VMOVMSKPD"
	case VMOVMSKPS:
		return "VMOVMSKPS"
	case VMOVNTDQ:
		return "VMOVNTDQ"
	case VMOVNTDQA:
		return "VMOVNTDQA"
	case VMOVNTPD:
		return "VMOVNTPD"
	case VMOVNTPS:
		return "VMOVNTPS"
	case VMOVQ:
		return "VMOVQ"
	case VMOVSHDUP:
		return "VMOVSHDUP"
	case VMOVSLDUP:
		return "VMOVSLDUP"
	case VMOVUPD:
		return "VMOVUPD"
	case VMOVUPS:
		return "VMOVUPS"
	case VMPSADBW:
		return "VMPSADBW"
	case VMULPD:
		return "VMULPD"
	case VMULPS:
		return "VMULPS"
	case VMULSD:
		return "VMULSD"
	case VMULSS:
		return "VMULSS"
	case VORPD:
		return "VORPD"
	case VORPS:
		return "VORPS"
	case VPABSB:
		return "VPABSB"
	case VPABSD:
		return "VPABSD"
	case VPABSW:
		return "VPABSW"
	case VPACKSSDW:
		return "VPACKSSDW"
	case VPACKSSWB:
		return "VPACKSSWB"
	case VPACKUSDW:
		return "VPACKUSDW"
	case VPACKUSWB:
		return "VPACKUSWB"
	case VPADDB:
		return "VPADDB"
	case VPADDD:
		return "VPADDD"
	case VPADDQ:
		return "VPADDQ"
	case VPADDSB:
		return "VPADDSB"
	case VPADDSW:
		return "VPADDSW"
	case VPADDUSB:
		return "VPADDUSB"
	case VPADDUSW:
		return "VPADDUSW"
	case VPADDW:
		return "VPADDW"
	case VPALIGNR:
		return "VPALIGNR"
	case VPAND:
		return "VPAND"
	case VPANDN:
		return "VPANDN"
	case VPAVGB:
		return "VPAVGB"
	case VPAVGW:
		return "VPAVGW"
	case VPBLENDD:
		return "VPBLENDD"
	case VPBLENDW:
		return "VPBLENDW"
	case VPBROADCASTB:
		return "VPBROADCASTB"
	case VPBROADCASTD:
		return "VPBROADCASTD"
	case VPBROADCASTQ:
		return "VPBROADCASTQ"
	case VPBROADCASTW:
		return "VPBROADCASTW"
	case VPCLMULQDQ:
		return "VPCLMULQDQ"
	case VPCMOV:
		return "VPCMOV"
	case VPCMPEQB:
		return "VPCMPEQB"
	case VPCMPEQD:
		return "VPCMPEQD"
	case VPCMPEQQ:
		return "VPCMPEQQ"
	case VPCMPEQW:
		return "VPCMPEQW"
	case VPCMPESTRI:
		return "VPCMPESTRI"
	case VPCMPESTRM:
		return "VPCMPESTRM"
	case VPCMPGTB:
		return "VPCMPGTB"
	case VPCMPGTD:
		return "VPCMPGTD"
	case VPCMPGTQ:
		return "VPCMPGTQ"
	case VPCMPGTW:
		return "VPCMPGTW"
	case VPCMPISTRI:
		return "VPCMPISTRI"
	case VPCMPISTRM:
		return "VPCMPISTRM"
	case VPCOMccB:
		return "VPCOMccB"
	case VPCOMccD:
		return "VPCOMccD"
	case VPCOMccQ:
		return "VPCOMccQ"
	case VPCOMccUB:
		return "VPCOMccUB"
	case VPCOMccUD:
		return "VPCOMccUD"
	case VPCOMccUQ:
		return "VPCOMccUQ"
	case VPCOMccUW:
		return "VPCOMccUW"
	case VPCOMccW:
		return "VPCOMccW"
	case VPDPBUSD:
		return "VPDPBUSD"
	case VPDPBUSDS:
		return "VPDPBUSDS"
	case VPDPWSSD:
		return "VPDPWSSD"
	case VPDPWSSDS:
		return "VPDPWSSDS"
	case VPERM2F128:
		return "VPERM2F128"
	case VPERM2I128:
		return "VPERM2I128"
	case VPERMD:
		return "VPERMD"
	case VPERMILPD:
		return "VPERMILPD"
	case VPERMILPS:
		return "VPERMILPS"
	case VPERMPD:
		return "VPERMPD"
	case VPERMPS:
		return "VPERMPS"
	case VPERMQ:
		return "VPERMQ"
	case VPEXTRB:
		return "VPEXTRB"
	case VPEXTRW:
		return "VPEXTRW"
	case VPHADDBD:
		return "VPHADDBD"
	case VPHADDBQ:
		return "VPHADDBQ"
	case VPHADDBW:
		return "VPHADDBW"
	case VPHADDD:
		return "VPHADDD"
	case VPHADDDQ:
		return "VPHADDDQ"
	case VPHADDSW:
		return "VPHADDSW"
	case VPHADDUBD:
		return "VPHADDUBD"
	case VPHADDUBQ:
		return "VPHADDUBQ"
	case VPHADDUBWD:
		return "VPHADDUBWD"
	case VPHADDUDQ:
		return "VPHADDUDQ"
	case VPHADDUWD:
		return "VPHADDUWD"
	case VPHADDUWQ:
		return "VPHADDUWQ"
	case VPHADDW:
		return "VPHADDW"
	case VPHADDWD:
		return "VPHADDWD"
	case VPHADDWQ:
		return "VPHADDWQ"
	case VPHMINPOSUW:
		return "VPHMINPOSUW"
	case VPHSUBBW:
		return "VPHSUBBW"
	case VPHSUBD:
		return "VPHSUBD"
	case VPHSUBDQ:
		return "VPHSUBDQ"
	case VPHSUBSW:
		return "VPHSUBSW"
	case VPHSUBW:
		return "VPHSUBW"
	case VPHSUBWD:
		return "VPHSUBWD"
	case VPINSRB:
		return "VPINSRB"
	case VPINSRW:
		return "VPINSRW"
	case VPMACSDD:
		return "VPMACSDD"
	case VPMACSDQH:
		return "VPMACSDQH"
	case VPMACSDQL:
		return "VPMACSDQL"
	case VPMACSSDD:
		return "VPMACSSDD"
	case VPMACSSDQH:
		return "VPMACSSDQH"
	case VPMACSSDQL:
		return "VPMACSSDQL"
	case VPMACSSWD:
		return "VPMACSSWD"
	case VPMACSSWW:
		return "VPMACSSWW"
	case VPMACSWD:
		return "VPMACSWD"
	case VPMACSWW:
		return "VPMACSWW"
	case VPMADCSSWD:
		return "VPMADCSSWD"
	case VPMADCSWD:
		return "VPMADCSWD"
	case VPMADDUBSW:
		return "VPMADDUBSW"
	case VPMADDWD:
		return "VPMADDWD"
	case VPMAXSB:
		return "VPMAXSB"
	case VPMAXSD:
		return "VPMAXSD"
	case VPMAXSW:
		return "VPMAXSW"
	case VPMAXUB:
		return "VPMAXUB"
	case VPMAXUD:
		return "VPMAXUD"
	case VPMAXUW:
		return "VPMAXUW"
	case VPMINSB:
		return "VPMINSB"
	case VPMINSD:
		return "VPMINSD"
	case VPMINSW:
		return "VPMINSW"
	case VPMINUB:
		return "VPMINUB"
	case VPMINUD:
		return "VPMINUD"
	case VPMINUW:
		return "VPMINUW"
	case VPMOVMSKB:
		return "VPMOVMSKB"
	case VPMOVSXBD:
		return "VPMOVSXBD"
	case VPMOVSXBQ:
		return "VPMOVSXBQ"
	case VPMOVSXBW:
		return "VPMOVSXBW"
	case VPMOVSXDQ:
		return "VPMOVSXDQ"
	case VPMOVSXWD:
		return "VPMOVSXWD"
	case VPMOVSXWQ:
		return "VPMOVSXWQ"
	case VPMOVZXBD:
		return "VPMOVZXBD"
	case VPMOVZXBQ:
		return "VPMOVZXBQ"
	case VPMOVZXBW:
		return "VPMOVZXBW"
	case VPMOVZXDQ:
		return "VPMOVZXDQ"
	case VPMOVZXWD:
		return "VPMOVZXWD"
	case VPMOVZXWQ:
		return "VPMOVZXWQ"
	case VPMULDQ:
		return "VPMULDQ"
	case VPMULHUW:
		return "VPMULHUW"
	case VPMULHW:
		return "VPMULHW"
	case VPMULLD:
		return "VPMULLD"
	case VPMULLW:
		return "VPMULLW"
	case VPMULUDQ:
		return "VPMULUDQ"
	case VPOR:
		return "VPOR"
	case VPPERM:
		return "VPPERM"
	case VPROTB:
		return "VPROTB"
	case VPROTD:
		return "VPROTD"
	case VPROTQ:
		return "VPROTQ"
	case VPROTW:
		return "VPROTW"
	case VPSADBW:
		return "VPSADBW"
	case VPSHAB:
		return "VPSHAB"
	case VPSHAD:
		return "VPSHAD"
	case VPSHAQ:
		return "VPSHAQ"
	case VPSHAW:
		return "VPSHAW"
	case VPSHLB:
		return "VPSHLB"
	case VPSHLD:
		return "VPSHLD"
	case VPSHLQ:
		return "VPSHLQ"
	case VPSHLW:
		return "VPSHLW"
	case VPSHUFB:
		return "VPSHUFB"
	case VPSHUFD:
		return "VPSHUFD"
	case VPSHUFHW:
		return "VPSHUFHW"
	case VPSHUFLW:
		return "VPSHUFLW"
	case VPSIGNB:
		return "VPSIGNB"
	case VPSIGND:
		return "VPSIGND"
	case VPSIGNW:
		return "VPSIGNW"
	case VPSLLD:
		return "VPSLLD"
	case VPSLLDQ:
		return "VPSLLDQ"
	case VPSLLQ:
		return "VPSLLQ"
	case VPSLLW:
		return "VPSLLW"
	case VPSRAD:
		return "VPSRAD"
	case VPSRAVD:
		return "VPSRAVD"
	case VPSRAW:
		return "VPSRAW"
	case VPSRLD:
		return "VPSRLD"
	case VPSRLDQ:
		return "VPSRLDQ"
	case VPSRLQ:
		return "VPSRLQ"
	case VPSRLW:
		return "VPSRLW"
	case VPSUBB:
		return "VPSUBB"
	case VPSUBD:
		return "VPSUBD"
	case VPSUBQ:
		return "VPSUBQ"
	case VPSUBSB:
		return "VPSUBSB"
	case VPSUBSW:
		return "VPSUBSW"
	case VPSUBUSW:
		return "VPSUBUSW"
	case VPSUBW:
		return "VPSUBW"
	case VPTEST:
		return "VPTEST"
	case VPUNPCKHBW:
		return "VPUNPCKHBW"
	case VPUNPCKHDQ:
		return "VPUNPCKHDQ"
	case VPUNPCKHQDQ:
		return "VPUNPCKHQDQ"
	case VPUNPCKHWD:
		return "VPUNPCKHWD"
	case VPUNPCKLBW:
		return "VPUNPCKLBW"
	case VPUNPCKLDQ:
		return "VPUNPCKLDQ"
	case VPUNPCKLQDQ:
		return "VPUNPCKLQDQ"
	case VPUNPCKLWD:
		return "VPUNPCKLWD"
	case VPXOR:
		return "VPXOR"
	case VRCPPS:
		return "VRCPPS"
	case VROUNDPD:
		return "VROUNDPD"
	case VROUNDPS:
		return "VROUNDPS"
	case VROUNDSD:
		return "VROUNDSD"
	case VROUNDSS:
		return "VROUNDSS"
	case VRSQRTPS:
		return "VRSQRTPS"
	case VRSQRTSS:
		return "VRSQRTSS"
	case VSHUFPD:
		return "VSHUFPD"
	case VSHUFPS:
		return "VSHUFPS"
	case VSQRTPD:
		return "VSQRTPD"
	case VSQRTPS:
		return "VSQRTPS"
	case VSQRTSD:
		return "VSQRTSD"
	case VSQRTSS:
		return "VSQRTSS"
	case VSTMXCSR:
		return "VSTMXCSR"
	case VSUBPD:
		return "VSUBPD"
	case VSUBPS:
		return "VSUBPS"
	case VSUBSD:
		return "VSUBSD"
	case VSUBSS:
		return "VSUBSS"
	case VTESTPD:
		return "VTESTPD"
	case VTESTPS:
		return "VTESTPS"
	case VUCOMISD:
		return "VUCOMISD"
	case VUCOMISS:
		return "VUCOMISS"
	case VUNPCKHPD:
		return "VUNPCKHPD"
	case VUNPCKHPS:
		return "VUNPCKHPS"
	case VUNPCKLPD:
		return "VUNPCKLPD"
	case VUNPCKLPS:
		return "VUNPCKLPS"
	case VXORPD:
		return "VXORPD"
	case VXORPS:
		return "VXORPS"
	case WRFSBASE:
		return "WRFSBASE"
	case WRMSR:
		return "WRMSR"
	case WRSS:
		return "WRSS"
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
	case XSAVE:
		return "XSAVE"
	case NoInstruction:
		return "NoInstruction"
	default:
		return "Unknown instruction"
	}
}
