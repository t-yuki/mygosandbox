// from "cmd/ld/textflag.h"
// but see go vet bug: https://code.google.com/p/go/issues/detail?id=6695
#define NOSPLIT 4

// see also
//   http://golang.org/doc/asm or Japanese translation: http://godocjp.herokuapp.com/doc/asm
//   http://golang.org/src/pkg/hash/crc32/crc32_amd64.s
//   http://golang.org/src/pkg/runtime/memmove_amd64.s
//   http://comments.gmane.org/gmane.os.plan9.general/70287
//   http://www.intel.com/content/www/us/en/processors/architectures-software-developer-manuals.html?iid=tech_vt_tech+64-32_manuals
TEXT ·BytesToBits(SB),4,$0-48
	MOVQ dst_base+0(FP), DI
	MOVQ dst_len+8(FP), CX
	MOVQ src_base+24(FP), SI

loop:
	// 32 -> 4px
	MOVOU    (SI), X1
	PMOVMSKB X1, AX
	SHLL     $16, AX
	MOVOU    16(SI), X2
	PMOVMSKB X2, BX
	ORL      BX, AX
	MOVL     AX, (DI)

	ADDQ $32, SI
	ADDQ $4, DI
	SUBQ $4, CX
	JG loop

	RET

TEXT ·bitsToBytes(SB),4,$0-48
	MOVQ dst_base+0(FP), DI
	MOVQ dst_len+8(FP), CX
	MOVQ src_base+24(FP), SI

	// load mask to X0 as [1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0]
	MOVD $4294967295, AX
	MOVD AX, X0
	PSHUFD $5, X0, X0

	// load mask to X7 as [1,1,1,1,1,1,1,1]
	MOVW $257, AX
	MOVD AX, X7
	PSHUFLW $0, X7, X7
	PSHUFD  $0, X7, X7

loop2:
	MOVOU   (SI), X1

	// first 2px -> 16px
	MOVO    X1, X2
	PSHUFB  X0, X2
	PMINUB  X7, X2
	MOVOU   X2, (DI)
	MOVB   (SI), AX
	MOVB   AX, (DI)

	BYTE $0x66; BYTE $0x0f; BYTE $0x3a; BYTE $0x0f; BYTE $0xd9; BYTE $0x02; // PALIGNR $2, X1, X3
	PSHUFB  X0, X3
	PMINUB  X7, X3
	MOVOU   X3, 16(DI)

	BYTE $0x66; BYTE $0x0f; BYTE $0x3a; BYTE $0x0f; BYTE $0xe1; BYTE $0x04; // PALIGNR $4, X1, X4
	PSHUFB  X0, X4
	PMINUB  X7, X4
	MOVOU   X4, 32(DI)

	BYTE $0x66; BYTE $0x0f; BYTE $0x3a; BYTE $0x0f; BYTE $0xe9; BYTE $0x06; // PALIGNR $6, X1, X5
	PSHUFB  X0, X5
	PMINUB  X7, X5
	MOVOU   X5, 48(DI)

	BYTE $0x66; BYTE $0x0f; BYTE $0x3a; BYTE $0x0f; BYTE $0xf1; BYTE $0x08; // PALIGNR $8, X1, X6
	PSHUFB  X0, X6
	PMINUB  X7, X6
	MOVOU   X6, 64(DI)

	BYTE $0x66; BYTE $0x0f; BYTE $0x3a; BYTE $0x0f; BYTE $0xd9; BYTE $0x0a; // PALIGNR $10, X1, X2
	PSHUFB  X0, X2
	PMINUB  X7, X2
	MOVOU   X2, 80(DI)

	BYTE $0x66; BYTE $0x0f; BYTE $0x3a; BYTE $0x0f; BYTE $0xe1; BYTE $0x0c; // PALIGNR $12, X1, X3
	PSHUFB  X0, X3
	PMINUB  X7, X3
	MOVOU   X3, 96(DI)

	BYTE $0x66; BYTE $0x0f; BYTE $0x3a; BYTE $0x0f; BYTE $0xe9; BYTE $0x0e; // PALIGNR $14, X1, X4
	PSHUFB  X0, X4
	PMINUB  X7, X4
	MOVOU   X4, 112(DI)

	ADDQ $16, SI
	ADDQ $128, DI
	SUBQ $16, CX
	JG loop2

	RET
