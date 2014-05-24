// from "cmd/ld/textflag.h"
// but see go vet bug: https://code.google.com/p/go/issues/detail?id=6695
#define NOSPLIT 4

TEXT ·Asm1(SB),4,$0-16
	MOVQ x+0(FP),AX
	INCQ AX
	MOVQ AX,ret+8(FP)
	RET

// see also 
//   http://golang.org/doc/asm or Japanese translation: http://godocjp.herokuapp.com/doc/asm
//   http://golang.org/src/pkg/hash/crc32/crc32_amd64.s
//   http://golang.org/src/pkg/runtime/memmove_amd64.s
//   http://comments.gmane.org/gmane.os.plan9.general/70287
TEXT ·Thresh(SB),4,$0-25
	// +0: pointer, +8: len, +16: cap, +24: th
	MOVQ buf_base+0(FP), SI
	MOVQ buf_len+8(FP), CX
	MOVB th+24(FP), AX

	// load th to X0 as [th,th,th,th,th,th,th,th]
	MOVD AX, X0
	PXOR X7, X7
	PSHUFB X7, X0

	// load mask to X7 as [1,1,1,1,1,1,1,1]
	MOVW $257, AX
	MOVD AX, X7
	PSHUFLW $0, X7, X7
	PSHUFD  $0, X7, X7

loop:
	MOVOU   (SI), X1 // MOVDQU to XMM1
	PSUBUSB X0, X1    
	PMINUB  X7, X1
	MOVOU   X1, (SI)

	MOVOU   16(SI), X2
	PSUBUSB X0, X2
	PMINUB  X7, X2
	MOVOU   X2, 16(SI)

	MOVOU   32(SI), X3
	PSUBUSB X0, X3
	PMINUB  X7, X3
	MOVOU   X3, 32(SI)

	MOVOU   48(SI), X4
	PSUBUSB X0, X4
	PMINUB  X7, X4
	MOVOU   X4, 48(SI)

	MOVOU   64(SI), X5
	PSUBUSB X0, X5
	PMINUB  X7, X5
	MOVOU   X5, 64(SI)

	MOVOU   80(SI), X6
	PSUBUSB X0, X6
	PMINUB  X7, X6
	MOVOU   X6, 80(SI)

	ADDQ $96, SI
	SUBQ $96, CX
	JG loop

	RET
