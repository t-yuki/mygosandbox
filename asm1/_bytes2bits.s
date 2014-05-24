// as --64 -o _bytes2bits.o _bytes2bits.s
// objdump -D -b binary -m i386:x86-64 _bytes2bits.o

.section .text
.globl _start
_start:
	PALIGNR $2, %XMM1, %XMM3;
	PALIGNR $4, %XMM1, %XMM4;
	PALIGNR $6, %XMM1, %XMM5;
	PALIGNR $8, %XMM1, %XMM6;
	PALIGNR $10, %XMM1, %XMM3;
	PALIGNR $12, %XMM1, %XMM4;
	PALIGNR $14, %XMM1, %XMM5;
	ret
