// from "cmd/ld/textflag.h"
#define NOSPLIT 4

TEXT ·Asm1(SB),NOSPLIT,$0
	MOVL x+0(FP),AX
	INCL AX
	MOVL AX,ret+8(FP)
	RET
