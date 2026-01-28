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
	ANDN
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
	BLENDVPD
	BLENDVPS
	BLSFILL
	BLSI
	BLSIC
	BLSMK
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
	CMPPD
	CMPPS
	CMPS
	CMPSB
	CMPSD
	CMPSQ
	CMPSW
	CMPXCHG
	CMPXCHG16B
	CMPXCHG8B
	COMISD
	COMISS
	CPUID
	CQO
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
	CWD
	CWDE
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
	F2XM1
	FABS
	FADD
	FADDP
	FBLD
	FBSTP
	FCHS
	FCLEX
	FCMOVB
	FCMOVBE
	FCMOVE
	FCMOVNB
	FCMOVNBE
	FCMOVNE
	FCMOVNU
	FCMOVU
	FCOM
	FCOMI
	FCOMIP
	FCOMP
	FCOMPP
	FCOS
	FDECSTP
	FDIV
	FDIVP
	FDIVR
	FDIVRP
	FEMMS
	FFREE
	FIADD
	FICOM
	FICOMP
	FIDIV
	FIDIVR
	FILD
	FIMUL
	FINCSTP
	FINIT
	FIST
	FISTP
	FISTTP
	FISUB
	FISUBR
	FLD
	FLD1
	FLDCW
	FLDENV
	FLDL2E
	FLDL2T
	FLDLG2
	FLDLN2
	FLDPI
	FLDZ
	FMUL
	FMULP
	FNINIT
	FNOP
	FNSAVE
	FPATAN
	FPREM
	FPREM1
	FPTAN
	FRNDINT
	FRSTOR
	FSAVE
	FSCALE
	FSIN
	FSINCOS
	FSQRT
	FST
	FSTCW
	FSTENV
	FSTP
	FSTSW
	FSUB
	FSUBP
	FSUBR
	FSUBRP
	FTST
	FUCOM
	FUCOMI
	FUCOMIP
	FUCOMP
	FUCOMPP
	FWAIT
	FXAM
	FXCH
	FXRSTOR
	FXSAVE
	FXTRACT
	FYL2X
	HADDPD
	HADDPS
	HLT
	HSUBPD
	HSUBPS
	IDIV
	IMUL
	IN
	INC
	INCSSP
	INCSSPD
	INCSSPQ
	INS
	INSB
	INSD
	INSERTPS
	INSERTQ
	INSW
	INT
	INT1
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
	MASKMOVDQU
	MASKMOVQ
	MAXPD
	MAXPS
	MAXSD
	MAXSS
	MCOMMIT
	MFENCE
	MINPD
	MINPS
	MINSD
	MINSS
	MONITOR
	MONITORX
	MOV
	MOVAPD
	MOVAPS
	MOVBE
	MOVD
	MOVDDUP
	MOVDIR64B
	MOVDIRI
	MOVDQ2Q
	MOVDQA
	MOVDQU
	MOVHLPS
	MOVHPD
	MOVHPS
	MOVLHPS
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
	MOVS
	MOVSB
	MOVSD
	MOVSHDUP
	MOVSLDUP
	MOVSQ
	MOVSS
	MOVSW
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
	MULX
	MWAIT
	MWAITX
	NEG
	NOP
	NOT
	OR
	ORPD
	ORPS
	OUT
	OUTS
	OUTSB
	OUTSD
	OUTSW
	PABSB
	PABSD
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
	PDEP
	PEXT
	PEXTRB
	PEXTRD
	PEXTRQ
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
	PINSRD
	PINSRQ
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
	PMULLW
	PMULUDQ
	POP
	POPA
	POPAD
	POPCNT
	POPD
	POPF
	POPFD
	POPFQ
	POR
	PREFETCH
	PREFETCHEXCLUSIVE
	PREFETCHIT0
	PREFETCHIT1
	PREFETCHMODIFIED
	PREFETCHNTA
	PREFETCHT0
	PREFETCHT1
	PREFETCHT2
	PREFETCHW
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
	PSMASH
	PSRAD
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
	PUSHF
	PUSHFD
	PUSHFQ
	PVALIDATE
	PXOR
	RCL
	RCPPS
	RCPSS
	RCR
	RDFSBASE
	RDGSBASE
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
	ROUNDPD
	ROUNDPS
	ROUNDSD
	ROUNDSS
	RSM
	RSQRTPS
	RSQRTSS
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
	SHUFPD
	SHUFPS
	SIDT
	SKINIT
	SLDT
	SLWPCB
	SMSW
	SQRTPD
	SQRTPS
	SQRTSD
	SQRTSS
	STAC
	STC
	STD
	STGI
	STI
	STMXCSR
	STOS
	STOSB
	STOSD
	STOSQ
	STOSW
	STR
	SUB
	SUBPD
	SUBPS
	SUBSD
	SUBSS
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
	VBLENDVPD
	VBLENDVPS
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
	VCVTDQ2PS
	VCVTPD2DQ
	VCVTPD2PS
	VCVTPH2PS
	VCVTPS2DQ
	VCVTPS2PD
	VCVTPS2PH
	VCVTSD2SI
	VCVTSD2SS
	VCVTSI2SD
	VCVTSI2SS
	VCVTSS2SD
	VCVTSS2SI
	VCVTTPD2DQ
	VCVTTPS2DQ
	VCVTTSD2SI
	VCVTTSS2SI
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
	VFMADD132PD
	VFMADD132PS
	VFMADD132SD
	VFMADD132SS
	VFMADD213PD
	VFMADD213PS
	VFMADD213SS
	VFMADD231PD
	VFMADD231SD
	VFMADD231SS
	VFMADDPD
	VFMADDPS
	VFMADDSD
	VFMADDSS
	VFMADDSUB132PD
	VFMADDSUB132PS
	VFMADDSUB213PD
	VFMADDSUB213PS
	VFMADDSUB231PD
	VFMADDSUB231PS
	VFMADDSUBPD
	VFMADDSUBPS
	VFMSUB132PD
	VFMSUB132PS
	VFMSUB132SD
	VFMSUB132SS
	VFMSUB213PD
	VFMSUB213PS
	VFMSUB213SD
	VFMSUB213SS
	VFMSUB231PD
	VFMSUB231PS
	VFMSUB231SD
	VFMSUB231SS
	VFMSUBADD132PD
	VFMSUBADD132PS
	VFMSUBADD213PD
	VFMSUBADD213PS
	VFMSUBADD231PD
	VFMSUBADD231PS
	VFMSUBADDPD
	VFMSUBADDPS
	VFMSUBPD
	VFMSUBPS
	VFMSUBSD
	VFMSUBSS
	VFNMADD132PD
	VFNMADD132PS
	VFNMADD132SD
	VFNMADD132SS
	VFNMADD213PD
	VFNMADD213PS
	VFNMADD213SD
	VFNMADD213SS
	VFNMADD231PD
	VFNMADD231PS
	VFNMADD231SD
	VFNMADD231SS
	VFNMADDPD
	VFNMADDPS
	VFNMADDSD
	VFNMADDSS
	VFNMSUB132PD
	VFNMSUB132PS
	VFNMSUB132SD
	VFNMSUB132SS
	VFNMSUB213PD
	VFNMSUB213PS
	VFNMSUB213SD
	VFNMSUB213SS
	VFNMSUB231PD
	VFNMSUB231PS
	VFNMSUB231SD
	VFNMSUB231SS
	VFNMSUBPD
	VFNMSUBPS
	VFNMSUBSD
	VFNMSUBSS
	VFRCZPD
	VFRCZPS
	VFRCZSD
	VFRCZSS
	VGATHERDPD
	VGATHERDPS
	VGATHERQPD
	VGATHERQPS
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
	VMGEXIT
	VMINPD
	VMINPS
	VMINSD
	VMINSS
	VMLOAD
	VMMCALL
	VMOVAPD
	VMOVAPS
	VMOVD
	VMOVDDUP
	VMOVDQA
	VMOVDQU
	VMOVHLPS
	VMOVHPD
	VMOVHPS
	VMOVLHPS
	VMOVLPD
	VMOVLPS
	VMOVMSKB
	VMOVMSKPD
	VMOVMSKPS
	VMOVNTDQ
	VMOVNTDQA
	VMOVNTPD
	VMOVNTPS
	VMOVQ
	VMOVSD
	VMOVSHDUP
	VMOVSLDUP
	VMOVSS
	VMOVUPD
	VMOVUPS
	VMPSADBW
	VMRUN
	VMSAVE
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
	VPBLENDVB
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
	VPCOMB
	VPCOMccB
	VPCOMccD
	VPCOMccQ
	VPCOMccUB
	VPCOMccUD
	VPCOMccUQ
	VPCOMccUW
	VPCOMccW
	VPCOMD
	VPCOMQ
	VPCOMUB
	VPCOMUD
	VPCOMUQ
	VPCOMUW
	VPCOMW
	VPDPBUSD
	VPDPBUSDS
	VPDPWSSD
	VPDPWSSDS
	VPERM2F128
	VPERM2I128
	VPERMD
	VPERMIL2PD
	VPERMIL2PS
	VPERMILPD
	VPERMILPS
	VPERMPD
	VPERMPS
	VPERMQ
	VPEXTRB
	VPEXTRD
	VPEXTRQ
	VPEXTRW
	VPGATHERDD
	VPGATHERDQ
	VPGATHERQD
	VPGATHERQQ
	VPHADDBD
	VPHADDBQ
	VPHADDBW
	VPHADDD
	VPHADDDQ
	VPHADDSW
	VPHADDUBD
	VPHADDUBQ
	VPHADDUBW
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
	VPINSRD
	VPINSRQ
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
	VPMASKMOVD
	VPMASKMOVQ
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
	VPMULHRSW
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
	VPSLLVD
	VPSLLVQ
	VPSLLW
	VPSRAD
	VPSRAVD
	VPSRAW
	VPSRLD
	VPSRLDQ
	VPSRLQ
	VPSRLVD
	VPSRLVQ
	VPSRLW
	VPSUBB
	VPSUBD
	VPSUBQ
	VPSUBSB
	VPSUBSW
	VPSUBUSB
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
	VRCPSS
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
	VZEROALL
	VZEROUPPER
	WBINVD
	WBNOINVD
	WRFSBASE
	WRGSBASE
	WRMSR
	WRPKRU
	WRSS
	WRSSQ
	WRUSS
	WRUSSD
	WRUSSQ
	XADD
	XCHG
	XGETBV
	XLAT
	XOR
	XORPD
	XORPS
	XRSTOR
	XSAVE
	XSAVEOPT
	XSETBV
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
	case ANDN:
		return "ANDN"
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
	case BLENDVPD:
		return "BLENDVPD"
	case BLENDVPS:
		return "BLENDVPS"
	case BLSFILL:
		return "BLSFILL"
	case BLSI:
		return "BLSI"
	case BLSIC:
		return "BLSIC"
	case BLSMK:
		return "BLSMK"
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
	case CMPPD:
		return "CMPPD"
	case CMPPS:
		return "CMPPS"
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
	case COMISD:
		return "COMISD"
	case COMISS:
		return "COMISS"
	case CPUID:
		return "CPUID"
	case CQO:
		return "CQO"
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
	case F2XM1:
		return "F2XM1"
	case FABS:
		return "FABS"
	case FADD:
		return "FADD"
	case FADDP:
		return "FADDP"
	case FBLD:
		return "FBLD"
	case FBSTP:
		return "FBSTP"
	case FCHS:
		return "FCHS"
	case FCLEX:
		return "FCLEX"
	case FCMOVB:
		return "FCMOVB"
	case FCMOVBE:
		return "FCMOVBE"
	case FCMOVE:
		return "FCMOVE"
	case FCMOVNB:
		return "FCMOVNB"
	case FCMOVNBE:
		return "FCMOVNBE"
	case FCMOVNE:
		return "FCMOVNE"
	case FCMOVNU:
		return "FCMOVNU"
	case FCMOVU:
		return "FCMOVU"
	case FCOM:
		return "FCOM"
	case FCOMI:
		return "FCOMI"
	case FCOMIP:
		return "FCOMIP"
	case FCOMP:
		return "FCOMP"
	case FCOMPP:
		return "FCOMPP"
	case FCOS:
		return "FCOS"
	case FDECSTP:
		return "FDECSTP"
	case FDIV:
		return "FDIV"
	case FDIVP:
		return "FDIVP"
	case FDIVR:
		return "FDIVR"
	case FDIVRP:
		return "FDIVRP"
	case FEMMS:
		return "FEMMS"
	case FFREE:
		return "FFREE"
	case FIADD:
		return "FIADD"
	case FICOM:
		return "FICOM"
	case FICOMP:
		return "FICOMP"
	case FIDIV:
		return "FIDIV"
	case FIDIVR:
		return "FIDIVR"
	case FILD:
		return "FILD"
	case FIMUL:
		return "FIMUL"
	case FINCSTP:
		return "FINCSTP"
	case FINIT:
		return "FINIT"
	case FIST:
		return "FIST"
	case FISTP:
		return "FISTP"
	case FISTTP:
		return "FISTTP"
	case FISUB:
		return "FISUB"
	case FISUBR:
		return "FISUBR"
	case FLD:
		return "FLD"
	case FLD1:
		return "FLD1"
	case FLDCW:
		return "FLDCW"
	case FLDENV:
		return "FLDENV"
	case FLDL2E:
		return "FLDL2E"
	case FLDL2T:
		return "FLDL2T"
	case FLDLG2:
		return "FLDLG2"
	case FLDLN2:
		return "FLDLN2"
	case FLDPI:
		return "FLDPI"
	case FLDZ:
		return "FLDZ"
	case FMUL:
		return "FMUL"
	case FMULP:
		return "FMULP"
	case FNINIT:
		return "FNINIT"
	case FNOP:
		return "FNOP"
	case FNSAVE:
		return "FNSAVE"
	case FPATAN:
		return "FPATAN"
	case FPREM:
		return "FPREM"
	case FPREM1:
		return "FPREM1"
	case FPTAN:
		return "FPTAN"
	case FRNDINT:
		return "FRNDINT"
	case FRSTOR:
		return "FRSTOR"
	case FSAVE:
		return "FSAVE"
	case FSCALE:
		return "FSCALE"
	case FSIN:
		return "FSIN"
	case FSINCOS:
		return "FSINCOS"
	case FSQRT:
		return "FSQRT"
	case FST:
		return "FST"
	case FSTCW:
		return "FSTCW"
	case FSTENV:
		return "FSTENV"
	case FSTP:
		return "FSTP"
	case FSTSW:
		return "FSTSW"
	case FSUB:
		return "FSUB"
	case FSUBP:
		return "FSUBP"
	case FSUBR:
		return "FSUBR"
	case FSUBRP:
		return "FSUBRP"
	case FTST:
		return "FTST"
	case FUCOM:
		return "FUCOM"
	case FUCOMI:
		return "FUCOMI"
	case FUCOMIP:
		return "FUCOMIP"
	case FUCOMP:
		return "FUCOMP"
	case FUCOMPP:
		return "FUCOMPP"
	case FWAIT:
		return "FWAIT"
	case FXAM:
		return "FXAM"
	case FXCH:
		return "FXCH"
	case FXRSTOR:
		return "FXRSTOR"
	case FXSAVE:
		return "FXSAVE"
	case FXTRACT:
		return "FXTRACT"
	case FYL2X:
		return "FYL2X"
	case HADDPD:
		return "HADDPD"
	case HADDPS:
		return "HADDPS"
	case HLT:
		return "HLT"
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
	case INSERTPS:
		return "INSERTPS"
	case INSERTQ:
		return "INSERTQ"
	case INSW:
		return "INSW"
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
		return "JAE"
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
	case MCOMMIT:
		return "MCOMMIT"
	case MFENCE:
		return "MFENCE"
	case MINPD:
		return "MINPD"
	case MINPS:
		return "MINPS"
	case MINSD:
		return "MINSD"
	case MINSS:
		return "MINSS"
	case MONITOR:
		return "MONITOR"
	case MONITORX:
		return "MONITORX"
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
	case MOVDIR64B:
		return "MOVDIR64B"
	case MOVDIRI:
		return "MOVDIRI"
	case MOVDQ2Q:
		return "MOVDQ2Q"
	case MOVDQA:
		return "MOVDQA"
	case MOVDQU:
		return "MOVDQU"
	case MOVHLPS:
		return "MOVHLPS"
	case MOVHPD:
		return "MOVHPD"
	case MOVHPS:
		return "MOVHPS"
	case MOVLHPS:
		return "MOVLHPS"
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
	case MOVS:
		return "MOVS"
	case MOVSB:
		return "MOVSB"
	case MOVSD:
		return "MOVSD"
	case MOVSHDUP:
		return "MOVSHDUP"
	case MOVSLDUP:
		return "MOVSLDUP"
	case MOVSQ:
		return "MOVSQ"
	case MOVSS:
		return "MOVSS"
	case MOVSW:
		return "MOVSW"
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
	case ORPD:
		return "ORPD"
	case ORPS:
		return "ORPS"
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
	case PABSB:
		return "PABSB"
	case PABSD:
		return "PABSD"
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
	case PDEP:
		return "PDEP"
	case PEXT:
		return "PEXT"
	case PEXTRB:
		return "PEXTRB"
	case PEXTRD:
		return "PEXTRD"
	case PEXTRQ:
		return "PEXTRQ"
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
	case PINSRD:
		return "PINSRD"
	case PINSRQ:
		return "PINSRQ"
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
	case POPF:
		return "POPF"
	case POPFD:
		return "POPFD"
	case POPFQ:
		return "POPFQ"
	case POR:
		return "POR"
	case PREFETCH:
		return "PREFETCH"
	case PREFETCHEXCLUSIVE:
		return "PREFETCHEXCLUSIVE"
	case PREFETCHIT0:
		return "PREFETCHIT0"
	case PREFETCHIT1:
		return "PREFETCHIT1"
	case PREFETCHMODIFIED:
		return "PREFETCHMODIFIED"
	case PREFETCHNTA:
		return "PREFETCHNTA"
	case PREFETCHT0:
		return "PREFETCHT0"
	case PREFETCHT1:
		return "PREFETCHT1"
	case PREFETCHT2:
		return "PREFETCHT2"
	case PREFETCHW:
		return "PREFETCHW"
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
	case PSMASH:
		return "PSMASH"
	case PSRAD:
		return "PSRAD"
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
	case PUSHF:
		return "PUSHF"
	case PUSHFD:
		return "PUSHFD"
	case PUSHFQ:
		return "PUSHFQ"
	case PVALIDATE:
		return "PVALIDATE"
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
		return "RDSSPQ"
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
		return "RMPUPDATE"
	case ROL:
		return "ROL"
	case ROR:
		return "ROR"
	case RORX:
		return "RORX"
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
		return "SAVEPREVSSP"
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
		return "SETAE"
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
		return "SETSSBSY"
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
	case SHUFPD:
		return "SHUFPD"
	case SHUFPS:
		return "SHUFPS"
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
	case SQRTPD:
		return "SQRTPD"
	case SQRTPS:
		return "SQRTPS"
	case SQRTSD:
		return "SQRTSD"
	case SQRTSS:
		return "SQRTSS"
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
	case STMXCSR:
		return "STMXCSR"
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
	case SUBPD:
		return "SUBPD"
	case SUBPS:
		return "SUBPS"
	case SUBSD:
		return "SUBSD"
	case SUBSS:
		return "SUBSS"
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
	case VBLENDVPD:
		return "VBLENDVPD"
	case VBLENDVPS:
		return "VBLENDVPS"
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
	case VCVTDQ2PS:
		return "VCVTDQ2PS"
	case VCVTPD2DQ:
		return "VCVTPD2DQ"
	case VCVTPD2PS:
		return "VCVTPD2PS"
	case VCVTPH2PS:
		return "VCVTPH2PS"
	case VCVTPS2DQ:
		return "VCVTPS2DQ"
	case VCVTPS2PD:
		return "VCVTPS2PD"
	case VCVTPS2PH:
		return "VCVTPS2PH"
	case VCVTSD2SI:
		return "VCVTSD2SI"
	case VCVTSD2SS:
		return "VCVTSD2SS"
	case VCVTSI2SD:
		return "VCVTSI2SD"
	case VCVTSI2SS:
		return "VCVTSI2SS"
	case VCVTSS2SD:
		return "VCVTSS2SD"
	case VCVTSS2SI:
		return "VCVTSS2SI"
	case VCVTTPD2DQ:
		return "VCVTTPD2DQ"
	case VCVTTPS2DQ:
		return "VCVTTPS2DQ"
	case VCVTTSD2SI:
		return "VCVTTSD2SI"
	case VCVTTSS2SI:
		return "VCVTTSS2SI"
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
	case VFMADD132PD:
		return "VFMADD132PD"
	case VFMADD132PS:
		return "VFMADD132PS"
	case VFMADD132SD:
		return "VFMADD132SD"
	case VFMADD132SS:
		return "VFMADD132SS"
	case VFMADD213PD:
		return "VFMADD213PD"
	case VFMADD213PS:
		return "VFMADD213PS"
	case VFMADD213SS:
		return "VFMADD213SS"
	case VFMADD231PD:
		return "VFMADD231PD"
	case VFMADD231SD:
		return "VFMADD231SD"
	case VFMADD231SS:
		return "VFMADD231SS"
	case VFMADDPD:
		return "VFMADDPD"
	case VFMADDPS:
		return "VFMADDPS"
	case VFMADDSD:
		return "VFMADDSD"
	case VFMADDSS:
		return "VFMADDSS"
	case VFMADDSUB132PD:
		return "VFMADDSUB132PD"
	case VFMADDSUB132PS:
		return "VFMADDSUB132PS"
	case VFMADDSUB213PD:
		return "VFMADDSUB213PD"
	case VFMADDSUB213PS:
		return "VFMADDSUB213PS"
	case VFMADDSUB231PD:
		return "VFMADDSUB231PD"
	case VFMADDSUB231PS:
		return "VFMADDSUB231PS"
	case VFMADDSUBPD:
		return "VFMADDSUBPD"
	case VFMADDSUBPS:
		return "VFMADDSUBPS"
	case VFMSUB132PD:
		return "VFMSUB132PD"
	case VFMSUB132PS:
		return "VFMSUB132PS"
	case VFMSUB132SD:
		return "VFMSUB132SD"
	case VFMSUB132SS:
		return "VFMSUB132SS"
	case VFMSUB213PD:
		return "VFMSUB213PD"
	case VFMSUB213PS:
		return "VFMSUB213PS"
	case VFMSUB213SD:
		return "VFMSUB213SD"
	case VFMSUB213SS:
		return "VFMSUB213SS"
	case VFMSUB231PD:
		return "VFMSUB231PD"
	case VFMSUB231PS:
		return "VFMSUB231PS"
	case VFMSUB231SD:
		return "VFMSUB231SD"
	case VFMSUB231SS:
		return "VFMSUB231SS"
	case VFMSUBADD132PD:
		return "VFMSUBADD132PD"
	case VFMSUBADD132PS:
		return "VFMSUBADD132PS"
	case VFMSUBADD213PD:
		return "VFMSUBADD213PD"
	case VFMSUBADD213PS:
		return "VFMSUBADD213PS"
	case VFMSUBADD231PD:
		return "VFMSUBADD231PD"
	case VFMSUBADD231PS:
		return "VFMSUBADD231PS"
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
	case VFNMADD132PD:
		return "VFNMADD132PD"
	case VFNMADD132PS:
		return "VFNMADD132PS"
	case VFNMADD132SD:
		return "VFNMADD132SD"
	case VFNMADD132SS:
		return "VFNMADD132SS"
	case VFNMADD213PD:
		return "VFNMADD213PD"
	case VFNMADD213PS:
		return "VFNMADD213PS"
	case VFNMADD213SD:
		return "VFNMADD213SD"
	case VFNMADD213SS:
		return "VFNMADD213SS"
	case VFNMADD231PD:
		return "VFNMADD231PD"
	case VFNMADD231PS:
		return "VFNMADD231PS"
	case VFNMADD231SD:
		return "VFNMADD231SD"
	case VFNMADD231SS:
		return "VFNMADD231SS"
	case VFNMADDPD:
		return "VFNMADDPD"
	case VFNMADDPS:
		return "VFNMADDPS"
	case VFNMADDSD:
		return "VFNMADDSD"
	case VFNMADDSS:
		return "VFNMADDSS"
	case VFNMSUB132PD:
		return "VFNMSUB132PD"
	case VFNMSUB132PS:
		return "VFNMSUB132PS"
	case VFNMSUB132SD:
		return "VFNMSUB132SD"
	case VFNMSUB132SS:
		return "VFNMSUB132SS"
	case VFNMSUB213PD:
		return "VFNMSUB213PD"
	case VFNMSUB213PS:
		return "VFNMSUB213PS"
	case VFNMSUB213SD:
		return "VFNMSUB213SD"
	case VFNMSUB213SS:
		return "VFNMSUB213SS"
	case VFNMSUB231PD:
		return "VFNMSUB231PD"
	case VFNMSUB231PS:
		return "VFNMSUB231PS"
	case VFNMSUB231SD:
		return "VFNMSUB231SD"
	case VFNMSUB231SS:
		return "VFNMSUB231SS"
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
	case VGATHERDPD:
		return "VGATHERDPD"
	case VGATHERDPS:
		return "VGATHERDPS"
	case VGATHERQPD:
		return "VGATHERQPD"
	case VGATHERQPS:
		return "VGATHERQPS"
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
	case VMGEXIT:
		return "VMGEXIT"
	case VMINPD:
		return "VMINPD"
	case VMINPS:
		return "VMINPS"
	case VMINSD:
		return "VMINSD"
	case VMINSS:
		return "VMINSS"
	case VMLOAD:
		return "VMLOAD"
	case VMMCALL:
		return "VMMCALL"
	case VMOVAPD:
		return "VMOVAPD"
	case VMOVAPS:
		return "VMOVAPS"
	case VMOVD:
		return "VMOVD"
	case VMOVDDUP:
		return "VMOVDDUP"
	case VMOVDQA:
		return "VMOVDQA"
	case VMOVDQU:
		return "VMOVDQU"
	case VMOVHLPS:
		return "VMOVHLPS"
	case VMOVHPD:
		return "VMOVHPD"
	case VMOVHPS:
		return "VMOVHPS"
	case VMOVLHPS:
		return "VMOVLHPS"
	case VMOVLPD:
		return "VMOVLPD"
	case VMOVLPS:
		return "VMOVLPS"
	case VMOVMSKB:
		return "VMOVMSKB"
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
	case VMOVSD:
		return "VMOVSD"
	case VMOVSHDUP:
		return "VMOVSHDUP"
	case VMOVSLDUP:
		return "VMOVSLDUP"
	case VMOVSS:
		return "VMOVSS"
	case VMOVUPD:
		return "VMOVUPD"
	case VMOVUPS:
		return "VMOVUPS"
	case VMPSADBW:
		return "VMPSADBW"
	case VMRUN:
		return "VMRUN"
	case VMSAVE:
		return "VMSAVE"
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
	case VPBLENDVB:
		return "VPBLENDVB"
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
	case VPCOMB:
		return "VPCOMB"
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
	case VPCOMD:
		return "VPCOMD"
	case VPCOMQ:
		return "VPCOMQ"
	case VPCOMUB:
		return "VPCOMUB"
	case VPCOMUD:
		return "VPCOMUD"
	case VPCOMUQ:
		return "VPCOMUQ"
	case VPCOMUW:
		return "VPCOMUW"
	case VPCOMW:
		return "VPCOMW"
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
	case VPERMIL2PD:
		return "VPERMIL2PD"
	case VPERMIL2PS:
		return "VPERMIL2PS"
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
	case VPEXTRD:
		return "VPEXTRD"
	case VPEXTRQ:
		return "VPEXTRQ"
	case VPEXTRW:
		return "VPEXTRW"
	case VPGATHERDD:
		return "VPGATHERDD"
	case VPGATHERDQ:
		return "VPGATHERDQ"
	case VPGATHERQD:
		return "VPGATHERQD"
	case VPGATHERQQ:
		return "VPGATHERQQ"
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
	case VPHADDUBW:
		return "VPHADDUBW"
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
	case VPINSRD:
		return "VPINSRD"
	case VPINSRQ:
		return "VPINSRQ"
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
	case VPMASKMOVD:
		return "VPMASKMOVD"
	case VPMASKMOVQ:
		return "VPMASKMOVQ"
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
	case VPMULHRSW:
		return "VPMULHRSW"
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
	case VPSLLVD:
		return "VPSLLVD"
	case VPSLLVQ:
		return "VPSLLVQ"
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
	case VPSRLVD:
		return "VPSRLVD"
	case VPSRLVQ:
		return "VPSRLVQ"
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
	case VPSUBUSB:
		return "VPSUBUSB"
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
	case VRCPSS:
		return "VRCPSS"
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
	case VZEROALL:
		return "VZEROALL"
	case VZEROUPPER:
		return "VZEROUPPER"
	case WBINVD:
		return "WBINVD"
	case WBNOINVD:
		return "WBNOINVD"
	case WRFSBASE:
		return "WRFSBASE"
	case WRGSBASE:
		return "WRGSBASE"
	case WRMSR:
		return "WRMSR"
	case WRPKRU:
		return "WRPKRU"
	case WRSS:
		return "WRSS"
	case WRSSQ:
		return "WRSSQ"
	case WRUSS:
		return "WRUSS"
	case WRUSSD:
		return "WRUSSD"
	case WRUSSQ:
		return "WRUSSQ"
	case XADD:
		return "XADD"
	case XCHG:
		return "XCHG"
	case XGETBV:
		return "XGETBV"
	case XLAT:
		return "XLAT"
	case XOR:
		return "XOR"
	case XORPD:
		return "XORPD"
	case XORPS:
		return "XORPS"
	case XRSTOR:
		return "XRSTOR"
	case XSAVE:
		return "XSAVE"
	case XSAVEOPT:
		return "XSAVEOPT"
	case XSETBV:
		return "XSETBV"
	case NoInstruction:
		return "NoInstruction"
	default:
		return "Unknown instruction"
	}
}
